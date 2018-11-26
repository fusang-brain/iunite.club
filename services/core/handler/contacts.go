package handler

import (
	"context"

	"iunite.club/services/organization/proto"

	"iunite.club/services/core/proto/conversation"

	userPB "iunite.club/services/user/proto"

	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	pb "iunite.club/services/core/proto/contacts"
	orgPB "iunite.club/services/organization/proto"
)

type ContactsHandler struct {
	ironic.BaseHandler
	connection monger.Connection
}

func (self *ContactsHandler) model(ctx context.Context, name string) monger.Model {
	if self.connection == nil {
		conn, err := ironic.MongerConnectionFromContext(ctx)

		if err != nil {
			panic(err.Error())
		}
		self.connection = conn
	}

	return self.connection.M(name)
}

func (hdl *ContactsHandler) FindFriendList(ctx context.Context, req *pb.FindFriendListRequest, rsp *pb.FriendListResponse) error {
	// panic("not implemented")
	UserFriendModel := hdl.model(ctx, "UserFriend")
	foundUserFriends := make([]models.UserFriend, 0)
	UserFriendModel.Where(bson.M{"user_id": req.UserID}).Populate("Friend", "Friend.Profile", "Friend.School").FindAll(&foundUserFriends)

	users := make([]*userPB.User, 0, len(foundUserFriends))

	for _, uf := range foundUserFriends {
		users = append(users, uf.Friend.ToPB())
	}

	rsp.Users = users
	return nil
}

func (hdl *ContactsHandler) FindAllGroup(ctx context.Context, req *pb.FindAllGroupRequest, rsp *pb.FindAllGroupResponse) error {
	ConversationModel := hdl.model(ctx, "Conversation")
	conversations := make([]models.Conversation, 0)
	ConversationModel.Where(bson.M{
		"$or": []bson.M{
			{
				"master": req.UserID,
			},
			{
				"members.$.user_id": req.UserID,
			},
		},
	}).FindAll(&conversations)

	conversationPBs := make([]*iunite_club_srv_core_conversation.ConversationPB, 0, len(conversations))

	for _, c := range conversations {
		conversationPBs = append(conversationPBs, c.ToPB())
	}

	rsp.Conversations = conversationPBs

	return nil
}

func (hdl *ContactsHandler) FindContactsList(ctx context.Context, req *pb.FindContactsRequest, rsp *pb.ContactsResponse) error {
	UserClubProfileModel := hdl.model(ctx, "UserClubProfile")
	condition := bson.M{
		"organization_id": req.ClubID,
	}

	if req.Department != "" {
		condition["department_id"] = req.Department
	}

	if req.Search != "" {
		condition["$or"] = []bson.M{
			{
				"user.profile.first_name": bson.RegEx{req.Search, "i"},
			},
			{
				"user.profile.last_name": bson.RegEx{req.Search, "i"},
			},
			{
				"user.profile.mobile": bson.RegEx{req.Search, "i"},
			},
		}
	}

	userClubProfiles := make([]models.UserClubProfile, 0)
	UserClubProfileModel.Where(condition).Populate("User", "User.Profile").FindAll(&userClubProfiles)
	total := UserClubProfileModel.Where(condition).Populate("User", "User.Profile").Count()

	contacts := make([]*orgPB.UserClubProfile, 0)

	for _, ucp := range userClubProfiles {
		contacts = append(contacts, ucp.ToPB())
	}

	rsp.Contacts = contacts
	rsp.Total = int32(total)
	return nil
}

func (hdl *ContactsHandler) FindDepartmentGroupByUserID(ctx context.Context, req *pb.UserIDRequest, rsp *pb.OrganizationsResponse) error {
	UserClubProfileModel := hdl.model(ctx, "UserClubProfile")
	Organization := hdl.model(ctx, "Organization")
	organizations := make([]models.Organization, 0)
	userClubProfiles := make([]models.UserClubProfile, 0)
	UserClubProfileModel.Where(bson.M{"user_id": req.UserID}).Populate("Organization").FindAll(&userClubProfiles)
	clubIDs := make([]bson.ObjectId, 0)
	for _, ucp := range userClubProfiles {
		clubIDs = append(clubIDs, ucp.OrganizationID)
		organizations = append(organizations, *ucp.Organization)
	}
	departments := make([]models.Organization, 0)
	Organization.Where(bson.M{"club_id": bson.M{"$in": clubIDs}}).FindAll(&departments)
	for _, dept := range departments {
		organizations = append(organizations, dept)
	}
	organizationPBs := make([]*iunite_club_srv_organization.Organization, 0)
	for _, org := range organizations {
		organizationPBs = append(organizationPBs, org.ToPB())
	}

	rsp.Organizations = organizationPBs
	return nil
}

func (hdl *ContactsHandler) FindUsersByDepartment(ctx context.Context, req *pb.DepartmentRequest, rsp *pb.UserJobListResponse) error {
	UserClubProfileModel := hdl.model(ctx, "UserClubProfile")
	userClubProfiles := make([]models.UserClubProfile, 0)
	UserClubProfileModel.Where(bson.M{
		"department_id": req.DepartmentID,
	}).Populate("User", "User.Profile", "Job").FindAll(&userClubProfiles)
	userJobs := make([]*pb.UserJob, 0)
	for _, ucp := range userClubProfiles {
		userJobs = append(userJobs, &pb.UserJob{
			User: ucp.User.ToPB(),
			Job:  ucp.Job.ToPB(),
		})
	}
	rsp.UserJobs = userJobs
	return nil
}

func (hdl *ContactsHandler) FindDepartmentsByOrganization(ctx context.Context, req *pb.OrgRequest, rsp *pb.OrganizationsResponse) error {
	OrganizationModel := hdl.model(ctx, "Organization")
	departments := make([]models.Organization, 0)
	OrganizationModel.Where(bson.M{
		"club_id": req.OrgID,
	}).FindAll(&departments)
	departmentpbs := make([]*iunite_club_srv_organization.Organization, 0, len(departments))
	for _, dept := range departments {
		departmentpbs = append(departmentpbs, dept.ToPB())
	}

	rsp.Organizations = departmentpbs
	rsp.Counts = int32(len(departmentpbs))
	return nil
}

func (hdl *ContactsHandler) GetUserCardDetails(ctx context.Context, req *pb.UserIDRequest, rsp *pb.UserCardResponse) error {
	UserModel := hdl.model(ctx, "User")
	UserClubProfileModel := hdl.model(ctx, "UserClubProfile")
	UserFriendModel := hdl.model(ctx, "UserFriend")
	foundUser := new(models.User)
	userClubProfiles := make([]models.UserClubProfile, 0)
	UserModel.Where(bson.M{"_id": req.UserID}).Populate("Profile", "School").FindOne(foundUser)
	UserClubProfileModel.Where(bson.M{"user_id": req.UserID}).Populate("User", "Organization", "Job", "Department").FindAll(&userClubProfiles)
	rsp.BaseInfo = foundUser.ToPB()
	pbUserClubProfiles := make([]*orgPB.UserClubProfile, 0)

	for _, ucp := range userClubProfiles {
		pbUserClubProfiles = append(pbUserClubProfiles, ucp.ToPB())
	}
	rsp.OrganizationUserInfo = pbUserClubProfiles

	foundCount := UserFriendModel.Where(bson.M{
		"user_id":   req.CurrentUserID,
		"friend_id": req.UserID,
	}).Count()

	rsp.IsFriend = foundCount > 0

	currUserClubs := make([]models.UserClubProfile, 0)
	UserClubProfileModel.Where(bson.M{
		"user_id": req.CurrentUserID,
	}).FindAll(&currUserClubs)
	clubs := make([]bson.ObjectId, 0)
	for _, ucp := range currUserClubs {
		clubs = append(clubs, ucp.OrganizationID)
	}

	foundUserClub := UserClubProfileModel.Where(bson.M{
		"user_id": req.CurrentUserID,
		"organization_id": bson.M{
			"$in": clubs,
		},
	}).Count()

	rsp.InSameOrganization = foundUserClub > 0
	return nil
}

func (hdl *ContactsHandler) SearchUsers(ctx context.Context, req *pb.SearchUserRequest, rsp *pb.UsersResponse) error {
	UserModel := hdl.model(ctx, "User")
	users := make([]models.User, 0)
	UserModel.Where(bson.M{"profile.mobile": bson.RegEx{req.Mobile, "i"}}).Populate("Profile").FindAll(&users)

	userpbs := make([]*userPB.User, 0)

	for _, u := range users {
		userpbs = append(userpbs, u.ToPB())
	}
	rsp.Users = userpbs
	return nil
}

func (hdl *ContactsHandler) AddFriend(ctx context.Context, req *pb.AddFriendRequest, rsp *pb.CreatedResponse) error {
	SocialApplicationModel := hdl.model(ctx, "SocialApplication")

	if err := SocialApplicationModel.Create(&models.SocialApplication{
		SubjectID:  bson.ObjectIdHex(req.FriendID),
		SenderID:   bson.ObjectIdHex(req.UserID),
		ReceiverID: bson.ObjectIdHex(req.FriendID),
		Kind:       "FRIEND",
		Body:       "请求添加您为好友",
	}); err != nil {
		return hdl.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (hdl *ContactsHandler) RemoveFriend(ctx context.Context, req *pb.RemoveFriendRequest, rsp *pb.RemovedResponse) error {
	UserFriendModel := hdl.model(ctx, "UserFriend")

	if err := UserFriendModel.OffSoftDeletes().Where(bson.M{"user_id": req.UserID, "friend_id": req.FriendID}).Delete(); err != nil {
		return hdl.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (hdl *ContactsHandler) FindFriendAcceptList(ctx context.Context, req *pb.UserIDRequest, rsp *pb.FriendAcceptResponse) error {
	SocialApplicationModel := hdl.model(ctx, "SocialApplication")
	foundAccepts := make([]models.SocialApplication, 0)
	SocialApplicationModel.Where(bson.M{
		"receiver_id": req.UserID,
	}).FindAll(&foundAccepts)

	acceptList := make([]*pb.FriendAccept, 0)
	for _, accept := range foundAccepts {
		acceptList = append(acceptList, &pb.FriendAccept{
			ID:         accept.ID.Hex(),
			SenderID:   accept.SenderID.Hex(),
			ReceiverID: accept.ReceiverID.Hex(),
			Body:       accept.Body,
			GroupID:    accept.SubjectID.Hex(),
			State:      accept.State,
			Kind: func() int32 {
				if accept.Kind == "FRIEND" {
					return 0
				}
				return 1
			}(),
		})
	}

	rsp.AcceptList = acceptList
	return nil
}

func (hdl *ContactsHandler) FindFrientAcceptCount(ctx context.Context, req *pb.UserIDRequest, rsp *pb.CountResponse) error {
	SocialApplicationModel := hdl.model(ctx, "SocialApplication")
	count := SocialApplicationModel.Where(bson.M{
		"receiver_id": req.UserID,
	}).Count()

	rsp.Count = int32(count)

	return nil
}

func (hdl *ContactsHandler) AgreeFriendAccept(ctx context.Context, req *pb.AgreeFriendAcceptRequest, rsp *pb.UpdatedResponse) error {
	SocialApplicationModel := hdl.model(ctx, "SocialApplication")

	foundAccept := new(models.SocialApplication)

	SocialApplicationModel.FindByID(bson.ObjectIdHex(req.ID), foundAccept)

	if foundAccept.IsEmpty() {
		return hdl.Error(ctx).NotFound("Not found the accept")
	}

	if foundAccept.State > 0 {
		return hdl.Error(ctx).BadRequest("Accept has executed")
	}

	if foundAccept.Kind != "FRIEND" {
		return hdl.Error(ctx).BadRequest("Error accept type")
	}

	UserFriendModel := hdl.model(ctx, "UserFriend")

	if err := UserFriendModel.Create(&models.UserFriend{
		UserID:   foundAccept.SenderID,
		FriendID: foundAccept.SubjectID,
	}); err != nil {
		return hdl.Error(ctx).InternalServerError(err.Error())
	}

	if err := UserFriendModel.Create(&models.UserFriend{
		UserID:   foundAccept.SubjectID,
		FriendID: foundAccept.SenderID,
	}); err != nil {
		return hdl.Error(ctx).InternalServerError(err.Error())
	}

	// 将申请状态修改为已通过
	SocialApplicationModel.Update(bson.M{"_id": req.ID}, bson.M{"$set": bson.M{"state": 1}})

	rsp.OK = true
	return nil
}
