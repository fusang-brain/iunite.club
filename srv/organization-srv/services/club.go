package services

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	orgPB "iunite.club/srv/organization-srv/proto"
	clubPB "iunite.club/srv/organization-srv/proto/club"
	"time"
)

type ClubService struct {
	assistant.Service
	Error *assistant.ErrorManager
}

type AcceptJoinClubBundle struct {
	UserID       string
	ClubID       string
	JobID        string
	DepartmentID string
}

type ExecuteJoinClubAccept struct {
	IsPassed bool
	AcceptID string
}

type PagerBundle struct {
	Page  int
	Limit int
}

type ClubsResult struct {
	Organizations []models.Organization
	Total         int
}

func (c *ClubService) CreateClub(req *clubPB.CreateClubRequest) (*models.Organization, error) {
	ClubModel := c.Connection.M("Organization")
	UserClubProfileModel := c.Connection.M("UserClubProfile")
	newClub := models.Organization{
		Name:     req.Name,
		SchoolID: bson.ObjectIdHex(req.SchoolID),
		ClubProfile: models.ClubProfile{
			Scale: req.Scale,
			// Paperworks
		},
		Kind: "club",
	}

	if err := ClubModel.Create(&newClub); err != nil {
		return nil, c.Error.InternalServerError(err.Error())
	}
	now := time.Now()
	userClubProfile := &models.UserClubProfile{
		OrganizationID: newClub.ID,
		UserID:         bson.ObjectIdHex(req.CreatorID),
		IsMaster:       true,
		IsCreator:      true,
		JoinTime:       now,
		State:          1,
	}

	if err := UserClubProfileModel.Create(userClubProfile); err != nil {
		return nil, c.Error.InternalServerError(err.Error())
	}

	return &newClub, nil
}

func (c *ClubService) FindClubListByPage(req *clubPB.GetClubListRequest, resp *clubPB.ClubListResponse) error {
	ClubModel := c.Connection.M("Organization")
	organizations := []models.Organization{}

	condition := bson.M{
		"kind": "club",
	}

	total, err := ClubModel.Find(condition).Count()
	if err != nil {
		return c.Error.InternalServerError(err.Error())
	}
	// list := []map
	// ClubModel.Find().Exec()
	if err := ClubModel.Find().Exec(&organizations); err != nil {
		return c.Error.InternalServerError(err.Error())
	}

	resp.Total = int64(total)
	if resp.Organizations == nil {
		resp.Organizations = make([]*orgPB.Organization, 0, 1)
	}
	for index, org := range organizations {
		if index == 0 {
			resp.FirstID = org.ID.Hex()
		}

		if index == len(organizations)-1 {
			resp.LastID = org.ID.Hex()
		}
		resp.Organizations = append(resp.Organizations, &orgPB.Organization{
			ID:          org.ID.Hex(),
			Name:        org.Name,
			Slug:        org.Slug,
			SchoolID:    org.SchoolID.Hex(),
			Kind:        org.Kind,
			Description: org.Description,
			ParentID:    org.ParentID.Hex(),
		})
	}

	return nil
}

func (c *ClubService) GetClubsByUserID(id string, pg *PagerBundle) (*ClubsResult, error) {
	// OrganizationModel := c.Model("Oragnization")
	UserClubModel := c.Model("UserClubProfile")
	userClubProfiles := []models.UserClubProfile{}
	res := ClubsResult{
		Total: 0,
	}
	organizations := make([]models.Organization, 0, 1)

	condition := bson.M{
		"user_id":           bson.ObjectIdHex(id),
		"organization.kind": "club",
	}

	total, err := UserClubModel.Find(condition).Populate("Organization").Count()
	if err != nil {
		return nil, c.Error.InternalServerError(err.Error())
	}

	if err := UserClubModel.Find(condition).Populate("Organization").Exec(&userClubProfiles); err != nil {
		return nil, c.Error.InternalServerError(err.Error())
	}

	for _, v := range userClubProfiles {
		if v.Organization != nil {
			organizations = append(organizations, *v.Organization)
		}
	}

	res.Organizations = organizations
	res.Total = total

	return &res, nil
}

// AcceptJoinOneClub 申请加入某个社团
func (c *ClubService) AcceptJoinOneClub(in *AcceptJoinClubBundle) error {
	UserClubProfileModel := c.Model("UserClubProfile")
	JoinOrganizationAccept := c.Model("JoinOrganizationAccept")
	condition := bson.M{
		"user_id":         bson.ObjectIdHex(in.UserID),
		"organization_id": bson.ObjectIdHex(in.ClubID),
	}
	total, err := UserClubProfileModel.Find(condition).Count()

	if err != nil {
		return c.Error.InternalServerError(err.Error())
	}

	userClubProfile := models.UserClubProfile{}

	if total > 0 {
		// 已经拥有该社团关系, 判断是否需要重新加入
		UserClubProfileModel.FindOne(condition).Exec(&userClubProfile)

		if userClubProfile.State == 1 {
			// 在职状态，不需要重新申请
			return c.Error.ActionError("HasJoinedError") // 已经加入过了
		}

		UserClubProfileModel.UpsertID(userClubProfile.ID, bson.M{"$set": bson.M{"state": 3}}) // 重新加入(申请)
		// return nil
	} else {
		newUserClubProfile := models.UserClubProfile{
			State:          0,
			UserID:         bson.ObjectIdHex(in.UserID),
			OrganizationID: bson.ObjectIdHex(in.ClubID),
			IsCreator:      false,
			IsMaster:       false,
		}

		if bson.IsObjectIdHex(in.JobID) {
			newUserClubProfile.JobID = bson.ObjectIdHex(in.JobID)
		}
		if bson.IsObjectIdHex(in.ClubID) {
			newUserClubProfile.OrganizationID = bson.ObjectIdHex(in.ClubID)
		}
		if err := UserClubProfileModel.Create(&newUserClubProfile); err != nil {
			return c.Error.InternalServerError(err.Error())
		}
	}

	JoinOrganizationAccept.Create(&models.JoinOrganizationAccept{
		UserID:         bson.ObjectIdHex(in.UserID),
		OrganizationID: bson.ObjectIdHex(in.ClubID),
		State:          0,
	})
	// TODO 检查jobID、OrganizationID、DepartmentID 的合法性
	return nil
}

// ExecuteJoinClubAccept 处理加入社团请求
func (c *ClubService) ExecuteJoinClubAccept(in *ExecuteJoinClubAccept) error {
	if !bson.IsObjectIdHex(in.AcceptID) {
		return c.Error.InternalServerError("acceptID is not objectID")
	}

	AcceptModel := c.Model("JoinOrganizationAccept")
	UserClubProfileModel := c.Model("UserClubProfile")
	accept := models.JoinOrganizationAccept{}

	AcceptModel.FindByID(bson.ObjectIdHex(in.AcceptID)).Exec(&accept)

	if accept.IsEmpty() {
		return c.Error.ActionError("NotFoundAccept")
	}
	updateFields := bson.M{
		"state": 0,
	}
	if in.IsPassed {
		updateFields["state"] = 2
		// 通过，将用户设置为在职
		if err := UserClubProfileModel.Update(bson.M{"_id": accept.UserID}, bson.M{"$set": bson.M{"state": 1}}); err != nil {
			return c.Error.InternalServerError(err.Error())
		}
	} else {
		updateFields["state"] = 1
	}
	_, err := AcceptModel.UpsertID(accept.ID, bson.M{"$set": updateFields})
	if err != nil {
		return c.Error.InternalServerError(err.Error())
	}

	return nil
}
