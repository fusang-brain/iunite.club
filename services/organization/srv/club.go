package srv

import (
	"errors"
	"time"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	orgPB "iunite.club/services/organization/proto"
	clubPB "iunite.club/services/organization/proto/club"
)

type ClubService struct {
	ironic.Service
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
	FirstID       string
	LastID        string
}

// type UserClubProfilesResult struct {
// 	UserClubProfiles
// }

func (c *ClubService) Model(name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(c.Ctx())

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (c *ClubService) CreateClub(req *clubPB.CreateClubRequest) (*models.Organization, error) {

	ClubModel := c.Model("Organization")
	UserClubProfileModel := c.Model("UserClubProfile")
	newClub := models.Organization{
		Name:     req.Name,
		SchoolID: bson.ObjectIdHex(req.SchoolID),
		Enabled:  false,
		ClubProfile: models.ClubProfile{
			Scale: req.Scale,
			// Paperworks
		},
		Kind: "club",
	}

	if err := ClubModel.Create(&newClub); err != nil {
		return nil, c.Error().InternalServerError(err.Error())
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
		return nil, c.Error().InternalServerError(err.Error())
	}

	return &newClub, nil
}

func (c *ClubService) FindClubListByPage(req *clubPB.GetClubListRequest, resp *clubPB.ClubListResponse) error {
	ClubModel := c.Model("Organization")
	organizations := []models.Organization{}

	condition := bson.M{
		"kind": "club",
	}

	total := ClubModel.Where(condition).Count()

	// list := []map
	// ClubModel.Find().Exec()
	if err := ClubModel.FindAll(&organizations); err != nil {
		return c.Error().InternalServerError(err.Error())
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
		"user_id":              bson.ObjectIdHex(id),
		"organization.kind":    "club",
		"organization.enabled": true,
	}

	total := UserClubModel.Where(condition).Populate("Organization").Count()
	// if err != nil {
	// 	return nil, c.Error().InternalServerError(err.Error())
	// }

	if err := UserClubModel.Where(condition).Populate("Organization").FindAll(&userClubProfiles); err != nil {
		return nil, c.Error().InternalServerError(err.Error())
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
	OrganizationAccept := c.Model("OrganizationAccept")
	condition := bson.M{
		"user_id":         bson.ObjectIdHex(in.UserID),
		"organization_id": bson.ObjectIdHex(in.ClubID),
	}
	total := UserClubProfileModel.Where(condition).Count()

	userClubProfile := models.UserClubProfile{}

	if total > 0 {
		// 已经拥有该社团关系, 判断是否需要重新加入
		UserClubProfileModel.Where(condition).FindOne(&userClubProfile)

		if userClubProfile.State == 1 {
			// 在职状态，不需要重新申请
			return c.Error().ActionError("HasJoinedError") // 已经加入过了
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
			return c.Error().InternalServerError(err.Error())
		}
	}

	OrganizationAccept.Create(&models.OrganizationAccept{
		UserID:         bson.ObjectIdHex(in.UserID),
		OrganizationID: bson.ObjectIdHex(in.ClubID),
		State:          0,
		Kind:           2,
	})
	// TODO 检查jobID、OrganizationID、DepartmentID 的合法性
	return nil
}

// ExecuteJoinClubAccept 处理加入社团请求
func (c *ClubService) ExecuteJoinClubAccept(in *ExecuteJoinClubAccept) error {
	if !bson.IsObjectIdHex(in.AcceptID) {
		return c.Error().InternalServerError("acceptID is not objectID")
	}

	AcceptModel := c.Model("OrganizationAccept")
	UserClubProfileModel := c.Model("UserClubProfile")
	accept := models.OrganizationAccept{}

	AcceptModel.FindByID(bson.ObjectIdHex(in.AcceptID), &accept)

	if accept.IsEmpty() {
		return c.Error().ActionError("NotFoundAccept")
	}
	updateFields := bson.M{
		"state": 0,
	}
	if in.IsPassed {
		updateFields["state"] = 2
		// 通过，将用户设置为在职
		if err := UserClubProfileModel.Update(bson.M{"user_id": accept.UserID, "organization_id": accept.OrganizationID}, bson.M{"$set": bson.M{"state": 1}}); err != nil {
			return c.Error().InternalServerError(err.Error())
		}
	} else {
		updateFields["state"] = 1
	}
	_, err := AcceptModel.UpsertID(accept.ID, bson.M{"$set": updateFields})
	if err != nil {
		return c.Error().InternalServerError(err.Error())
	}

	return nil
}

// SearchClubs is function to search clubs by name
func (c *ClubService) SearchClubs(search string, page, limit int64) ([]models.Organization, int, error) {
	OrganizationModel := c.Model("Organization")

	condition := bson.M{
		"kind": "club",
		"name": bson.RegEx{search, "i"},
	}

	clubs := make([]models.Organization, 0)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit >= 500 {
		limit = 500
	}
	total := OrganizationModel.Count(condition)

	err := OrganizationModel.Where(condition).Limit(int(limit)).Skip(int((page - 1) * limit)).FindAll(&clubs)

	if err != nil {
		return clubs, total, c.Error().InternalServerError(err.Error())
	}

	return clubs, total, nil
}

// FindRefusedAcceptByUserID is function to find refused accept that joinin or create clubs
func (c *ClubService) FindRefusedAcceptByUserID(id string, page, limit int64) ([]models.OrganizationAccept, int, error) {
	OrganizationAcceptModel := c.Model("OrganizationAccept")
	cond := bson.M{
		"state": 1,
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit >= 500 {
		limit = 500
	}

	joinAccepts := make([]models.OrganizationAccept, 0, 1)

	total := OrganizationAcceptModel.Count(cond)

	if err := OrganizationAcceptModel.
		Where(cond).Skip(int((page - 1) * limit)).
		Limit(int(limit)).
		Populate("Organization").
		FindAll(&joinAccepts); err != nil {
		return joinAccepts, total, c.Error().InternalServerError(err.Error())
	}

	return joinAccepts, total, nil
}

func (c *ClubService) FindClubByID(id string) (*models.Organization, error) {
	ClubModel := c.Model("Organization")
	club := new(models.Organization)
	if !bson.IsObjectIdHex(id) {
		return club, c.Error().InternalServerError("id must be a objectid hex")
	}
	err := ClubModel.FindByID(bson.ObjectIdHex(id), club)

	if err != nil {
		return club, c.Error().InternalServerError(err.Error())
	}

	return club, nil
}

func (c *ClubService) UpdateClub(id bson.ObjectId, set map[string]interface{}) (time.Time, error) {
	ClubModel := c.Model("Organization")
	now := time.Now()

	set["updatedAt"] = now
	changeInfo, err := ClubModel.UpsertID(id, bson.M{
		"$set": set,
	})

	if err != nil {
		return now, err
	}

	if changeInfo.Updated > 0 {
		return now, nil
	}

	return now, errors.New("Not found this club")
}

func (c *ClubService) GetClubsBySchoolID(id bson.ObjectId) (*ClubsResult, error) {
	// OrganizationModel := c.Model("Oragnization")
	ClubModel := c.Model("Organization")
	organizations := []models.Organization{}
	resp := new(ClubsResult)

	condition := bson.M{
		"kind":      "club",
		"school_id": id,
	}

	total := ClubModel.Where(condition).Count()
	// if err != nil {
	// 	return nil, c.Error().InternalServerError(err.Error())
	// }
	// list := []map
	// ClubModel.Find().Exec()
	if err := ClubModel.FindAll(&organizations); err != nil {
		return nil, c.Error().InternalServerError(err.Error())
	}
	resp.Total = total
	resp.Organizations = organizations
	for index, org := range organizations {
		if index == 0 {
			resp.FirstID = org.ID.Hex()
		}

		if index == len(organizations)-1 {
			resp.LastID = org.ID.Hex()
		}
	}

	return resp, nil
}

func (c *ClubService) GetUserClubProfilesByUserID(id bson.ObjectId, rsp *clubPB.UserClubProfilesListResponse) error {
	UserClubProfileModel := c.Model("UserClubProfile")
	userClubProfiles := make([]models.UserClubProfile, 0)

	err := UserClubProfileModel.
		Where(bson.M{"user_id": id}).
		Populate("User", "Organization", "Job", "Department").
		FindAll(&userClubProfiles)
	if err != nil {
		return c.Error().InternalServerError(err.Error())
	}
	rsp.UserClubProfiles = make([]*orgPB.UserClubProfile, 0)
	rsp.Total = int64(len(userClubProfiles))
	for i, v := range userClubProfiles {
		if i == 0 {
			rsp.FirstID = v.ID.Hex()
		}

		if i == len(userClubProfiles)-1 {
			rsp.LastID = v.ID.Hex()
		}

		rsp.UserClubProfiles = append(rsp.UserClubProfiles, v.ToPB())
	}

	return nil
}

func (c *ClubService) GetUserClubProfileDetailsByID(orgID, userID bson.ObjectId, rsp *clubPB.UserClubProfileResponse) error {
	UserClubProfileModel := c.Model("UserClubProfile")
	userClubProfile := new(models.UserClubProfile)

	err := UserClubProfileModel.
		Where(bson.M{"organization_id": orgID, "user_id": userID}).
		Populate("User", "User.Profile", "Organization", "Job", "Department").
		FindOne(userClubProfile)

	if err != nil {
		return c.Error().InternalServerError(err.Error())
	}

	rsp.UserClubProfile = userClubProfile.ToPB()

	return nil
}

func (c *ClubService) RemoveUserFromClub(userID, clubID bson.ObjectId) error {
	UserClubProfileModel := c.Model("UserClubProfile")

	return UserClubProfileModel.ForceDelete(bson.M{
		"user_id":         userID,
		"organization_id": clubID,
	})
}
