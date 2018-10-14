package handler

import (
	"context"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"

	"iunite.club/services/restful/dto"

	schoolPB "iunite.club/services/organization/proto/school"
)

type SchoolSrv struct {
	BaseHandler
	schoolService schoolPB.SchoolSrvService
}

func (s *SchoolSrv) SchoolList(req *go_api.Request, rsp *go_api.Response) {
	schoolSrv := s.schoolService
	ctx := context.Background()
	params := struct {
		Page  int `json:"page,omitempty" query:"page"`
		Limit int `json:"limit,omitempty" query:"page"`
	}{}

	if err := s.Bind(req, &params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	if err := s.Validate(&params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	listResp, err := schoolSrv.GetSchoolList(ctx, &schoolPB.ListRequest{
		Page:  int32(params.Page),
		Limit: int32(params.Limit),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	schools := make([]*dto.School, 0)

	for _, v := range listResp.Schools {
		schools = append(schools, dto.PBToSchool(v))
	}
	SuccessResponse(rsp, D{
		"CurrentPage": params.Page,
		"PageSize":    params.Limit,
		"Total":       listResp.Total,
		"List":        schools,
	})
}

func (s *SchoolSrv) SearchSchools(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		SearchName string
	}{}

	if err := s.Bind(req, &params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	if err := s.Validate(&params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	schoolSrv := s.schoolService

	resp, err := schoolSrv.SearchSchools(ctx, &schoolPB.SearchSchoolsRequest{
		Search: params.SearchName,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	schools := make([]*dto.School, 0)

	for _, v := range resp.Schools {
		schools = append(schools, dto.PBToSchool(v))
	}

	SuccessResponse(rsp, D{
		"Schools": schools,
	})
}

func (s *SchoolSrv) Create(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	// panic("not implemented")
	params := struct {
		Name        string `json:"name,omitempty" validate:"nonzero"`
		SchoolCode  string `json:"schoolCode,omitempty"`
		Description string `json:"description,omitempty"`
	}{}

	if err := s.Bind(req, &params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	if err := s.Validate(&params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	schoolSrv := s.schoolService

	createResp, err := schoolSrv.CreateSchool(ctx, &schoolPB.CreateSchoolRequest{
		Name:        params.Name,
		Description: params.Description,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !createResp.OK {
		ErrorResponse(rsp, s.Error().BadRequest("创建失败"))
		return
	}
	SuccessResponse(rsp, params)
}

func NewSchoolHandler(c client.Client) *SchoolSrv {
	return &SchoolSrv{
		schoolService: schoolPB.NewSchoolSrvService(OrganizationService, c),
	}
}
