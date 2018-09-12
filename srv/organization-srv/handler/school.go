package handler

import (
	"context"
	"github.com/iron-kit/go-ironic/micro-assistant"
	pb "iunite.club/srv/organization-srv/proto/school"
	"iunite.club/srv/organization-srv/services"
)

type SchoolHandler struct {
	assistant.BaseHandler

	SchoolService *services.SchoolService
}

func (s *SchoolHandler) CreateSchool(ctx context.Context, req *pb.CreateSchoolRequest, resp *pb.CreateSchoolResponse) error {
	_, err := s.SchoolService.CreateSchool(req)

	if err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (s *SchoolHandler) GetSchoolList(ctx context.Context, req *pb.ListRequest, resp *pb.ListResponse) error {
	err := s.SchoolService.GetSchoolList(req, resp)

	if err != nil {
		return err
	}
	return nil
}
