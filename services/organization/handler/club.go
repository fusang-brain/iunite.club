package handler

import (
	"context"
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
