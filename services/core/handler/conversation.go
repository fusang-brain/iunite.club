package handler

import (
	"context"
	"fmt"
	"github.com/iron-kit/go-ironic/utils/agent"

	"iunite.club/services/user/proto"

	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	pb "iunite.club/services/core/proto/conversation"
)

type ConversationHandler struct {
	ironic.BaseHandler

	connection monger.Connection
}

type LCMetaData struct {
	UniteConversationID string
	Kind                string
	Name                string
	Avatar              string
	ID                  string
	IsTop               bool
	TopMembers          []string
}

func (self *ConversationHandler) model(ctx context.Context, name string) monger.Model {
	if self.connection == nil {
		conn, err := ironic.MongerConnectionFromContext(ctx)

		if err != nil {
			panic(err.Error())
		}
		self.connection = conn
	}

	return self.connection.M(name)
}

func (self *ConversationHandler) updateConversationMetaDataByID(ctx context.Context, id bson.ObjectId) (bool, *models.ConversationMetaData) {
	ConversationModel := self.model(ctx, "Conversation")
	conversation := new(models.Conversation)
	ConversationModel.
		Where(bson.M{"_id": id}).
		Populate("Members", "Members.User", "User.Profile").
		FindOne(conversation)
	if conversation.IsEmpty() {
		return false, nil
	}
	metaUsers := make([]*models.UserMetaData, 0)
	topMembers := make([]string, 0)

	metaData := LCMetaData{
		UniteConversationID: conversation.ID.Hex(),
		Kind:                conversation.Kind,
		Name:                conversation.Name,
		Avatar:              conversation.Avatar,
		ID:                  conversation.ID.Hex(),
	}

	for _, val := range conversation.Members {
		umd := new(models.UserMetaData)
		if conversation.Kind != "group" {
			if val.User != nil {
				umd.InitByUser(val.User, "", val.Nickname)
				metaUsers = append(metaUsers, umd)
			}
		}

		if val.IsTop {
			topMembers = append(topMembers, val.UserID.Hex())
		}
	}

	metaData.TopMembers = topMembers

	return self.updateConversationMetaData(&metaData, metaUsers)
}

func (self *ConversationHandler) updateConversationMetaData(metaData *LCMetaData, members []*models.UserMetaData) (bool, *models.ConversationMetaData) {
	cmd := &models.ConversationMetaData{
		UniteConversationID: metaData.UniteConversationID,
		ConversationName:    metaData.Name,
		ConversationAvatar:  metaData.Avatar,
		Kind:                metaData.Kind,
		IsTop:               metaData.IsTop,
		TopMembers:          metaData.TopMembers,
	}
	if members != nil {
		userMap := map[string]models.UserMetaData{}
		for _, v := range members {
			userMap[v.ID] = *v
		}

		cmd.MemberMapper = userMap
	}

	headers := map[string]string{
		"X-LC-Id":      "gS4uGBBo0mqYzrU9nqlOOxGC-gzGzoHsz",
		"X-LC-Key":     "iuhkU6AQpA4OuF9ft6WAsAsC",
		"Content-Type": "application/json",
	}

	postData := struct {
		Name          string
		UniteMetaData *models.ConversationMetaData
	}{
		UniteMetaData: cmd,
		Name:          cmd.ConversationName,
	}

	resp, data := agent.Put(fmt.Sprintf("%s/conversations/%s", "https://gs4ugbbo.api.lncld.net/1.2/rtm", metaData.ID), headers, &postData)
	if resp == nil {
		return false, nil
	}

	fmt.Println(string(data))

	return true, cmd
}

func (self *ConversationHandler) CreateConversation(ctx context.Context, req *pb.WithConversationBundle, rsp *pb.CreatedConversationOK) error {
	ConversationModel := self.model(ctx, "Conversation")

	if req.ConversationID == "" || !bson.IsObjectIdHex(req.ConversationID) {
		return self.Error(ctx).BadRequest("Conversation can be empty")
	}

	conversation := new(models.Conversation)
	ConversationModel.Where(bson.M{"_id": req.ConversationID}).FindOne(conversation)
	conversation.Kind = req.Kind
	conversation.Name = req.Name
	conversation.Avatar = req.Avatar
	conversation.Master = bson.ObjectIdHex(req.Master)
	conversation.IsStartValidate = false
	conversation.IsTop = false
	isExist := false

	if conversation.IsEmpty() {
		conversation.ID = bson.ObjectIdHex(req.ConversationID)
	} else {
		isExist = true
	}

	if len(req.Members) > 0 {
		members := make([]models.ConversationMember, 0)
		for _, v := range req.Members {
			members = append(members, models.ConversationMember{
				UserID: bson.ObjectIdHex(v),
			})
		}
		conversation.Members = members
	}

	if isExist {
		if err := ConversationModel.Update(bson.M{"_id": conversation.ID}, conversation); err != nil {
			fmt.Println(err.Error(), "FMT ERROR1")
			return self.Error(ctx).BadRequest(err.Error())
		}
	} else {
		if err := ConversationModel.Create(conversation); err != nil {
			fmt.Println(err.Error(), "FMT ERROR2")
			return self.Error(ctx).BadRequest(err.Error())
		}
	}

	// TODO 更新会话云数据
	self.updateConversationMetaDataByID(ctx, conversation.ID)

	rsp.OK = true
	rsp.ID = conversation.ID.Hex()
	rsp.IsExists = isExist

	return nil
}

func (self *ConversationHandler) GetConversationsByMemberID(ctx context.Context, req *pb.ByUserID, rsp *pb.ConversationsResponse) error {
	ConversationModel := self.model(ctx, "Conversation")
	conversations := make([]models.Conversation, 0)

	err := ConversationModel.
		Where(bson.M{
			"members": bson.M{
				"$elemMatch": bson.M{
					"_id": bson.ObjectIdHex(req.ID),
				},
			},
		}).
		Skip(0).
		Limit(100).
		FindAll(conversations)

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	pbConversations := make([]*pb.ConversationPB, 0)

	for _, v := range conversations {
		pbConversations = append(pbConversations, v.ToPB())
	}
	rsp.Conversations = pbConversations
	rsp.Total = int64(len(rsp.Conversations))
	return nil
}

func (self *ConversationHandler) FindConversationDetails(ctx context.Context, req *pb.ByID, rsp *pb.ConversationDetails) error {
	ConversationModel := self.model(ctx, "Conversation")
	conversation := new(models.Conversation)
	err := ConversationModel.
		Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).
		FindOne(conversation)

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.Conversation = conversation.ToPB()
	if req.GetPut() {
		self.updateConversationMetaDataByID(ctx, conversation.ID)
	}
	return nil
}

func (self *ConversationHandler) ExitGroup(ctx context.Context, req *pb.ByIDWithUserID, rsp *pb.IsOK) error {
	ConversationModel := self.model(ctx, "Conversation")
	err := ConversationModel.Update(bson.M{
		"_id": bson.ObjectIdHex(req.ID),
		"master": bson.M{
			"$neq": bson.ObjectIdHex(req.UserID),
		},
	}, bson.M{
		"$pull": bson.M{
			"members": bson.M{
				"user_id": req.UserID,
			},
		},
	})

	if err != nil {
		self.Error(ctx).BadRequest(err.Error())
	}
	return nil
}

func (self *ConversationHandler) DismissGroup(ctx context.Context, req *pb.ByIDWithUserID, rsp *pb.IsOK) error {
	ConversationModel := self.model(ctx, "Conversation")

	err := ConversationModel.Where(bson.M{
		"_id":    bson.ObjectIdHex(req.ID),
		"master": bson.ObjectIdHex(req.UserID),
	}).Delete()

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	return nil
}

func (self *ConversationHandler) UpdateGroupConversation(ctx context.Context, req *pb.WithUpdateBundle, rsp *pb.IsOK) error {
	foundConversation := new(models.Conversation)
	ConversationModel := self.model(ctx, "Conversation")

	ConversationModel.Where(bson.M{"_id": req.ID}).FindOne(foundConversation)

	if foundConversation.IsEmpty() {
		return self.Error(ctx).NotFound("not found conversation")
	}
	if req.Avatar != "" {
		foundConversation.Avatar = req.Avatar
	}

	if req.GroupName != "" {
		foundConversation.Name = req.GroupName
	}

	if req.IsTop != "" {

		if req.IsTop == "yes" {
			foundConversation.IsTop = true
		} else {
			foundConversation.IsTop = false
		}
	}

	if req.IsStartValidate != "" {
		if req.IsStartValidate == "yes" {
			foundConversation.IsStartValidate = true
		} else {
			foundConversation.IsStartValidate = false
		}
	}

	if req.Master != "" {
		foundConversation.Master = bson.ObjectIdHex(req.Master)
	}

	err := ConversationModel.Update(bson.M{"_id": bson.ObjectIdHex(req.ID)}, foundConversation)
	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}
	self.updateConversationMetaDataByID(ctx, foundConversation.ID)
	return nil
}

func (self *ConversationHandler) GetAllMembersOfConversation(ctx context.Context, req *pb.ByID, rsp *pb.MembersResponse) error {
	ConversationModel := self.model(ctx, "Conversation")
	UserModel := self.model(ctx, "User")
	conversation := new(models.Conversation)

	if err := ConversationModel.
		Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).
		FindAll(conversation); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	userids := make([]string, 0)

	for _, v := range conversation.Members {
		userids = append(userids, v.UserID.Hex())
	}

	users := make([]models.User, 0)

	if err := UserModel.Where(bson.M{
		"_id": bson.M{
			"$in": userids,
		},
	}).Populate("Profile").FindAll(&users); err != nil {
		return self.Error(ctx).InternalServerError(err.Error())
	}

	pbUsers := make([]*iunite_club_srv_user.User, 0)

	for _, v := range users {
		pbUsers = append(pbUsers, v.ToPB())
	}
	rsp.Members = pbUsers
	return nil
}

func (self *ConversationHandler) RemoveConversationNotice(ctx context.Context, req *pb.ByNoticeID, rsp *pb.IsOK) error {
	// self.
	ConversationNoticeModel := self.model(ctx, "ConversationNotice")
	if err := ConversationNoticeModel.Where(bson.M{"_id": req.ID}).Delete(); err != nil {
		return self.Error(ctx).InternalServerError(err.Error())
	}

	return nil
}

func (self *ConversationHandler) GetNoticeList(ctx context.Context, req *pb.ByIDWithPager, rsp *pb.NoticesResponse) error {
	ConversationNoticeModel := self.model(ctx, "ConversationNotice")

	notices := make([]models.ConversationNotice, 0, req.Limit)
	ConversationNoticeModel.
		Where(bson.M{"conversation_id": req.ID}).
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int((req.Page - 1) * req.Limit)).
		FindAll(&notices)

	count := ConversationNoticeModel.Where(bson.M{"conversation_id": req.ID}).Count()
	rsp.Total = int64(count)
	pbNotices := make([]*pb.NoticePB, 0, len(notices))
	for _, val := range notices {
		pbNotices = append(pbNotices, val.ToPB())
	}
	rsp.Notices = pbNotices

	return nil
}

func (self *ConversationHandler) CreateNotice(ctx context.Context, req *pb.WithNoticeBundle, rsp *pb.CreatedOK) error {
	ConversationNoticeModel := self.model(ctx, "ConversationNotice")
	conversationNotice := new(models.ConversationNotice)

	conversationNotice.ConversationID = bson.ObjectIdHex(req.ConversationID)
	conversationNotice.Body = req.Body
	conversationNotice.Title = req.Title

	if err := ConversationNoticeModel.Create(conversationNotice); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	rsp.ID = conversationNotice.ID.Hex()
	return nil
}

func (self *ConversationHandler) MarkedNoticeToHasRead(ctx context.Context, req *pb.WithMarkedBundle, rsp *pb.UpdatedOK) error {
	ConversationNoticeModel := self.model(ctx, "ConversationNotice")

	err := ConversationNoticeModel.Update(bson.M{
		"_id": bson.ObjectIdHex(req.NoticeID),
	}, bson.M{
		"$addToSet": bson.M{
			"readers": bson.ObjectIdHex(req.UserID),
		},
	})

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
	// panic("not implemented")
}

func (self *ConversationHandler) AddGroupMember(ctx context.Context, req *pb.WithIDAndMembers, rsp *pb.UpdatedOK) error {
	ConversationModel := self.model(ctx, "Conversation")

	cmembers := make([]models.ConversationMember, 0)
	members := make([]bson.ObjectId, 0)
	for _, m := range req.Members {
		cmembers = append(cmembers, models.ConversationMember{
			UserID: bson.ObjectIdHex(m),
		})
		members = append(members, bson.ObjectIdHex(m))
	}

	ConversationModel.Update(bson.M{
		"_id": bson.ObjectIdHex(req.ID),
	}, bson.M{
		"$pull": bson.M{
			"members": bson.M{
				"$_id": bson.M{
					"$in": members,
				},
			},
		},
	})
	err := ConversationModel.Update(bson.M{
		"_id": bson.ObjectIdHex(req.ID),
	}, bson.M{
		"$addToSet": bson.M{
			"members": bson.M{
				"$each": cmembers,
			},
		},
	})

	if err != nil {
		self.Error(ctx).BadRequest(err.Error())
	}
	return nil
}

func (self *ConversationHandler) RemoveGroupMember(ctx context.Context, req *pb.WithIDAndMembers, rsp *pb.UpdatedOK) error {
	ConversationModel := self.model(ctx, "Conversation")

	members := make([]bson.ObjectId, 0)
	for _, m := range req.Members {
		members = append(members, bson.ObjectIdHex(m))
	}

	err := ConversationModel.Update(bson.M{
		"_id": bson.ObjectIdHex(req.ID),
	}, bson.M{
		"$pull": bson.M{
			"members": bson.M{
				"$_id": bson.M{
					"$in": members,
				},
			},
		},
	})

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}
	return nil
}

func (self *ConversationHandler) JoinInGroup(ctx context.Context, req *pb.WithIDAndUserID, rsp *pb.IsOK) error {
	ConversationModel := self.model(ctx, "Conversation")
	SocialApplicationModel := self.model(ctx, "SocialApplication")
	conversation := new(models.Conversation)

	ConversationModel.FindByID(bson.ObjectIdHex(req.ID), conversation)
	if conversation.IsEmpty() {
		return self.Error(ctx).NotFound("Not found conversation %v", req.ID)
	}
	application := &models.SocialApplication{
		SenderID:   bson.ObjectIdHex(req.UserID),
		ReceiverID: conversation.Master,
		Kind:       "in.group",
		SubjectID:  bson.ObjectIdHex(req.ID),
		Body:       fmt.Sprintf("请求加入群 '%s'", conversation.Name),
	}

	foundCount := SocialApplicationModel.Count(bson.M{
		"sender_id":   bson.ObjectIdHex(req.UserID),
		"receiver_id": conversation.Master,
		"kind":        "in.group",
		"subject_id":  bson.ObjectIdHex(req.ID),
	})

	if foundCount == 0 {
		SocialApplicationModel.Create(application)
	}
	rsp.OK = true
	return nil
}
