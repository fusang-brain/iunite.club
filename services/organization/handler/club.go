package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"iunite.club/models"
	orgPB "iunite.club/services/organization/proto"
	pb "iunite.club/services/organization/proto/club"
	"iunite.club/services/organization/srv"
)

type ClubHandler struct {
	ironic.BaseHandler
}

func (u *ClubHandler) model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (o *ClubHandler) CreateClub(ctx context.Context, req *pb.CreateClubRequest, resp *pb.CreateClubResponse) error {
	clubService := srv.NewClubService(ctx)

	_, err := clubService.CreateClub(req)
	if err != nil {
		return err
	}

	resp.OK = true
	return nil
}

// func (o *ClubHandler) FindClubList(ctx context.Context, req *pb.GetClubListRequest, resp *pb.ClubListResponse) error {
// 	// OrganizationModel :=
// 	err := o.ClubService.FindClubListByPage(req, resp)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (o *ClubHandler) FindClubListByPage(ctx context.Context, req *pb.GetClubListRequest, resp *pb.ClubListResponse) error {
	clubService := srv.NewClubService(ctx)
	err := clubService.FindClubListByPage(req, resp)
	if err != nil {
		return err
	}

	return nil
}

func (o *ClubHandler) GetClubsByUserID(ctx context.Context, req *pb.GetClubsByUserIDRequest, resp *pb.ClubListResponse) error {
	// result := services.ClubsResult{}
	clubService := srv.NewClubService(ctx)
	result, err := clubService.GetClubsByUserID(req.UserID, &srv.PagerBundle{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	})

	if err != nil {
		return err
	}

	resp.Organizations = make([]*orgPB.Organization, 0, 1)
	for _, v := range result.Organizations {
		orgRaw := v.ToPB()
		resp.Organizations = append(resp.Organizations, orgRaw)
	}

	resp.Total = int64(result.Total)
	return nil
}

func (o *ClubHandler) AcceptJoinOneClub(ctx context.Context, req *pb.AcceptJoinOneClubRequest, resp *pb.Response) error {
	clubService := srv.NewClubService(ctx)
	if err := clubService.AcceptJoinOneClub(&srv.AcceptJoinClubBundle{
		UserID:       req.UserID,
		ClubID:       req.ClubID,
		JobID:        req.JobID,
		DepartmentID: req.DepartmentID,
	}); err != nil {
		return err
	}

	resp.OK = true

	return nil
}

func (self *ClubHandler) FindAcceptByUserClubProfileID(ctx context.Context, req *pb.ByUserClubProfileIDRequest, rsp *pb.AcceptResponse) error {
	OrganizationClub := self.model(ctx, "OrganizationAccept")
	UserClubProfile := self.model(ctx, "UserClubProfile")

	organizationAccept := new(models.OrganizationAccept)
	ucp := new(models.UserClubProfile)

	if err := UserClubProfile.Where(bson.M{"_id": req.ID}).FindOne(ucp); err != nil {
		return self.Error(ctx).InternalServerError(err.Error())
	}

	err := OrganizationClub.Where(bson.M{
		"user_id":         ucp.UserID,
		"organization_id": ucp.OrganizationID,
		"state":           0,
	}).FindOne(organizationAccept)

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.ID = organizationAccept.ID.Hex()
	rsp.UserID = organizationAccept.UserID.Hex()
	rsp.OrganizationID = organizationAccept.OrganizationID.Hex()
	rsp.State = int32(organizationAccept.State)
	rsp.Kind = int32(organizationAccept.Kind)

	return nil
}

// func (self *ClubHandler) FindAcceptBy(ctx context.Context, req *pb.UserIDWithClubID, rsp *pb.AcceptResponse) error {
// 	OrganizationClub := self.model(ctx, "OrganizationAccept")

// 	organizationAccept := new(models.OrganizationAccept)

// 	err := OrganizationClub.Where(bson.M{
// 		"user_id":         req.UserID,
// 		"organization_id": req.ClubID,
// 		"state":           0,
// 	}).FindOne(organizationAccept)

// 	if err != nil {
// 		return self.Error(ctx).BadRequest(err.Error())
// 	}

// 	rsp.ID = organizationAccept.ID.Hex()
// 	rsp.UserID = organizationAccept.UserID.Hex()
// 	rsp.OrganizationID = organizationAccept.OrganizationID.Hex()
// 	rsp.State = int32(organizationAccept.State)
// 	rsp.Kind = int32(organizationAccept.Kind)

// 	return nil
// }

func (o *ClubHandler) ExecuteJoinClubAccept(ctx context.Context, req *pb.ExecuteJoinClubAcceptRequest, resp *pb.Response) error {
	clubService := srv.NewClubService(ctx)
	if err := clubService.ExecuteJoinClubAccept(&srv.ExecuteJoinClubAccept{
		IsPassed: req.IsPassed,
		AcceptID: req.AcceptID,
	}); err != nil {
		return err
	}

	resp.OK = true

	return nil
}

func (o *ClubHandler) SearchClubs(ctx context.Context, req *pb.SearchClubRequest, resp *pb.ClubListResponse) error {
	clubService := srv.NewClubService(ctx)

	clubs, total, err := clubService.SearchClubs(req.Search, req.Page, req.Limit)

	if err != nil {
		return err
	}

	resp.Organizations = make([]*orgPB.Organization, 0, 1)

	for index, raw := range clubs {
		if index == 0 {
			resp.FirstID = raw.ID.Hex()
		}

		if index == len(clubs)-1 {
			resp.LastID = raw.ID.Hex()
		}

		resp.Organizations = append(resp.Organizations, raw.ToPB())
	}

	resp.Total = int64(total)

	return nil
}

func (o *ClubHandler) FindRefusedAcceptByUserID(ctx context.Context, req *pb.FindRefusedAcceptRequest, rsp *pb.AcceptListResponse) error {
	clubService := srv.NewClubService(ctx)

	clubAccepts, total, err := clubService.FindRefusedAcceptByUserID(req.UserID, req.Page, req.Limit)

	if err != nil {
		return err
	}

	rsp.Total = int64(total)
	rsp.Accepts = make([]*orgPB.ClubAccept, 0, 1)
	for index, raw := range clubAccepts {
		if index == 0 {
			rsp.FirstID = raw.ID.Hex()
		}

		if index == len(clubAccepts)-1 {
			rsp.LastID = raw.ID.Hex()
		}

		rsp.Accepts = append(rsp.Accepts, &orgPB.ClubAccept{
			ID:             raw.ID.Hex(),
			UserID:         raw.UserID.Hex(),
			OrganizationID: raw.OrganizationID.Hex(),
			State:          int64(raw.State),
		})
	}

	return nil
}

func (o *ClubHandler) FindClubDetailsByID(ctx context.Context, req *pb.GetClubByIDRequest, rsp *pb.ClubDetailsResponse) error {
	clubService := srv.NewClubService(ctx)

	club, err := clubService.FindClubByID(req.ID)
	if err != nil {
		return err
	}

	rsp.Club = club.ToPB()

	return nil
}

func (o *ClubHandler) UpdateClubInfo(ctx context.Context, req *pb.UpdateClubInfoRequest, rsp *pb.UpdatedResponse) error {
	clubService := srv.NewClubService(ctx)
	// req.String()
	if !bson.IsObjectIdHex(req.ID) {
		return o.Error(ctx).InternalServerError("ID must be a objectid")
	}

	toSet := make(map[string]interface{})

	if err := json.Unmarshal(req.ToSet, &toSet); err != nil {
		return o.Error(ctx).InternalServerError(err.Error())
	}

	updatedAt, err := clubService.UpdateClub(bson.ObjectIdHex(req.ID), toSet)

	if err != nil {
		return o.Error(ctx).InternalServerError(err.Error())
	}

	rsp.UpdateAt = updatedAt.String()
	rsp.OK = true
	return nil
}

func (o *ClubHandler) FindClubsBySchoolID(ctx context.Context, req *pb.GetClubsBySchoolIDRequest, rsp *pb.ClubListResponse) error {
	clubService := srv.NewClubService(ctx)
	if !bson.IsObjectIdHex(req.SchoolID) {
		return o.Error(ctx).BadRequest("SchoolID must be a objectid")
	}

	clubListResp, err := clubService.GetClubsBySchoolID(bson.ObjectIdHex(req.SchoolID))

	if err != nil {
		return err
	}

	rsp.FirstID = clubListResp.FirstID
	rsp.LastID = clubListResp.LastID

	if rsp.Organizations == nil {
		rsp.Organizations = make([]*orgPB.Organization, 0)
	}

	for _, val := range clubListResp.Organizations {
		rsp.Organizations = append(rsp.Organizations, val.ToPB())
	}

	rsp.Total = int64(clubListResp.Total)

	return nil
}

func (o *ClubHandler) GetUserClubProfilesByUserID(ctx context.Context, req *pb.GetUserClubProfilesByUserIDRequest, rsp *pb.UserClubProfilesListResponse) error {
	// panic("not impl")
	clubSrv := srv.NewClubService(ctx)

	if !bson.IsObjectIdHex(req.UserID) {
		return o.Error(ctx).BadRequest("UserID must be a objectid")
	}

	if err := clubSrv.GetUserClubProfilesByUserID(bson.ObjectIdHex(req.UserID), rsp); err != nil {
		return err
	}

	return nil
}

func (o *ClubHandler) GetUserClubProfileDetailsByID(ctx context.Context, req *pb.GetUserClubProfileDetailsByIDRequest, rsp *pb.UserClubProfileResponse) error {
	clubSrv := srv.NewClubService(ctx)

	if !bson.IsObjectIdHex(req.OrganizationID) {
		return o.Error(ctx).BadRequest("OrganizationID must be a objectid")
	}

	if !bson.IsObjectIdHex(req.UserID) {
		return o.Error(ctx).BadRequest("UserID must be a objectid")
	}

	if err := clubSrv.GetUserClubProfileDetailsByID(bson.ObjectIdHex(req.OrganizationID), bson.ObjectIdHex(req.UserID), rsp); err != nil {
		return err
	}

	return nil
}

func (o *ClubHandler) RemoveUserFromClub(ctx context.Context, req *pb.RemoveUserFromClubRequest, rsp *pb.Response) error {
	clubSrv := srv.NewClubService(ctx)

	if !bson.IsObjectIdHex(req.UserID) {
		return o.Error(ctx).BadRequest("userID must be objectid")
	}

	if !bson.IsObjectIdHex(req.ClubID) {
		return o.Error(ctx).BadRequest("clubID must be objectid")
	}

	if err := clubSrv.RemoveUserFromClub(bson.ObjectIdHex(req.UserID), bson.ObjectIdHex(req.ClubID)); err != nil {
		return o.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	return nil
}

func (u *ClubHandler) FindUserClubProfiles(ctx context.Context, req *pb.FindUserClubProfilesRequest, rsp *pb.UserClubProfilesResponse) error {
	UserClubProfile := u.model(ctx, "UserClubProfile")

	ucps := make([]models.UserClubProfile, 0)
	ucppbs := make([]*orgPB.UserClubProfile, 0)

	condition := bson.M{
		"organization_id": req.ClubID,
		// "department_id":   req.DepartmentID,
		// "job_id":          req.JobID,
	}

	if req.DepartmentID != "" {
		condition["department_id"] = req.DepartmentID
	}

	if req.JobID != "" {
		condition["job_id"] = req.JobID
	}

	if req.Category == "accept" {
		condition["state"] = 0
	} else {
		condition["state"] = bson.M{
			"$eq": 1,
		}
	}

	fmt.Println(condition, "cond")

	query := UserClubProfile.Where(condition).
		Populate("User", "User.Profile", "Job", "Department").
		Query()
	// FindAll(&ucps)
	err := query.Query().
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&ucps)

	fmt.Println(ucps, "ucps")
	if err != nil {
		return u.Error(ctx).InternalServerError(err.Error())
	}
	for _, v := range ucps {
		ucppbs = append(ucppbs, v.ToPB())
	}

	rsp.UserClubProfiles = ucppbs
	rsp.Total = int32(query.Query().Count())

	return nil
}
