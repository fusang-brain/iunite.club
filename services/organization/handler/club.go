package handler

import (
	"context"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/iron-kit/go-ironic"
	orgPB "iunite.club/services/organization/proto"
	pb "iunite.club/services/organization/proto/club"
	"iunite.club/services/organization/srv"
)

type ClubHandler struct {
	ironic.BaseHandler
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

func (o ClubHandler) AcceptJoinOneClub(ctx context.Context, req *pb.AcceptJoinOneClubRequest, resp *pb.Response) error {
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

func (o ClubHandler) ExecuteJoinClubAccept(ctx context.Context, req *pb.ExecuteJoinClubAcceptRequest, resp *pb.Response) error {
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

func (o ClubHandler) SearchClubs(ctx context.Context, req *pb.SearchClubRequest, resp *pb.ClubListResponse) error {
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

func (o ClubHandler) FindRefusedAcceptByUserID(ctx context.Context, req *pb.FindRefusedAcceptRequest, rsp *pb.AcceptListResponse) error {
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

func (o ClubHandler) FindClubDetailsByID(ctx context.Context, req *pb.GetClubByIDRequest, rsp *pb.ClubDetailsResponse) error {
	clubService := srv.NewClubService(ctx)

	club, err := clubService.FindClubByID(req.ID)
	if err != nil {
		return err
	}

	rsp.Club = club.ToPB()

	return nil
}

func (o ClubHandler) UpdateClubInfo(ctx context.Context, req *pb.UpdateClubInfoRequest, rsp *pb.UpdatedResponse) error {
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

func (o ClubHandler) FindClubsBySchoolID(ctx context.Context, req *pb.GetClubsBySchoolIDRequest, rsp *pb.ClubListResponse) error {
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

func (o ClubHandler) GetUserClubProfilesByUserID(ctx context.Context, req *pb.GetUserClubProfilesByUserIDRequest, rsp *pb.ClubProfilesListResponse) {
	panic("not impl")
}
