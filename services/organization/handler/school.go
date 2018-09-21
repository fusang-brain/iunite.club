package handler

import (
	"context"
	"github.com/iron-kit/go-ironic"

	pb "iunite.club/services/organization/proto/school"
	"iunite.club/services/organization/srv"
)

type SchoolHandler struct {
	ironic.BaseHandler
}

func (s *SchoolHandler) CreateSchool(ctx context.Context, req *pb.CreateSchoolRequest, resp *pb.CreateSchoolResponse) error {
	schoolService := srv.NewSchoolService(ctx)
	_, err := schoolService.CreateSchool(req)

	if err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (s *SchoolHandler) GetSchoolList(ctx context.Context, req *pb.ListRequest, resp *pb.ListResponse) error {
	schoolService := srv.NewSchoolService(ctx)
	err := schoolService.GetSchoolList(req, resp)

	if err != nil {
		return err
	}
	return nil
}
