package handler

import (
	"context"
	"github.com/iron-kit/go-ironic/micro-assistant"
	orgPB "iunite.club/srv/organization-srv/proto"
	pb "iunite.club/srv/organization-srv/proto/club"
	"iunite.club/srv/organization-srv/services"
)

type ClubHandler struct {
	assistant.BaseHandler
	ClubService *services.ClubService
}

func (o *ClubHandler) CreateClub(ctx context.Context, req *pb.CreateClubRequest, resp *pb.CreateClubResponse) error {
	_, err := o.ClubService.CreateClub(req)
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
	err := o.ClubService.FindClubListByPage(req, resp)
	if err != nil {
		return err
	}

	return nil
}

func (o *ClubHandler) GetClubsByUserID(ctx context.Context, req *pb.GetClubsByUserIDRequest, resp *pb.ClubListResponse) error {
	// result := services.ClubsResult{}
	result, err := o.ClubService.GetClubsByUserID(req.UserID, &services.PagerBundle{
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
	if err := o.ClubService.AcceptJoinOneClub(&services.AcceptJoinClubBundle{
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
	if err := o.ClubService.ExecuteJoinClubAccept(&services.ExecuteJoinClubAccept{
		IsPassed: req.IsPassed,
		AcceptID: req.AcceptID,
	}); err != nil {
		return err
	}

	resp.OK = true

	return nil
}
