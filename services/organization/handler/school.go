package handler

import (
	"context"
	"fmt"

	"github.com/iron-kit/go-ironic"

	pb "iunite.club/services/organization/proto/school"
	"iunite.club/services/organization/srv"
)

type SchoolHandler struct {
	ironic.BaseHandler
}

func (s *SchoolHandler) GetSchoolByID(ctx context.Context, req *pb.GetSchoolRequest, rsp *pb.SchoolResponse) error {
	fmt.Println("get school by id")
	schoolService := srv.NewSchoolService(ctx)

	foundSchool, err := schoolService.GetSchoolByID(req.ID)

	if err != nil {
		return err
	}

	rsp.School = &pb.School{
		SchoolCode:  foundSchool.SchoolCode,
		Name:        foundSchool.Name,
		SlugName:    foundSchool.SlugName,
		Description: foundSchool.Description,
		ID:          foundSchool.ID.Hex(),
	}

	return nil
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

func (s *SchoolHandler) SearchSchools(ctx context.Context, req *pb.SearchSchoolsRequest, rsp *pb.ListResponse) error {
	schoolService := srv.NewSchoolService(ctx)

	err := schoolService.SearchSchools(req, rsp)

	if err != nil {
		return err
	}

	return nil
}
