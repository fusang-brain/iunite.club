package handler

import (
	"context"

	"iunite.club/services/navo/dto"

	"github.com/iron-kit/go-ironic"
	go_api "github.com/micro/go-api/proto"
	"iunite.club/services/navo/client"
	schoolPB "iunite.club/services/organization/proto/school"
)

type SchoolSrv struct {
	ironic.BaseHandler
}

func (s *SchoolSrv) SchoolList(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	schoolSrv, ok := client.SchoolServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, s.Error(ctx).InternalServerError("Not found school service"))
	}

	params := struct {
		Page  int `json:"page,omitempty" query:"page"`
		Limit int `json:"limit,omitempty" query:"page"`
	}{}

	if err := s.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	if err := s.Validate(&params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	listResp, err := schoolSrv.GetSchoolList(ctx, &schoolPB.ListRequest{
		Page:  int32(params.Page),
		Limit: int32(params.Limit),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}
	schools := make([]*dto.School, 0)

	for _, v := range listResp.Schools {
		schools = append(schools, dto.PBToSchool(v))
	}
	return SuccessResponse(rsp, D{
		"CurrentPage": params.Page,
		"PageSize":    params.Limit,
		"Total":       listResp.Total,
		"List":        schools,
	})
}

func (s *SchoolSrv) SearchSchools(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		SearchName string
	}{}

	if err := s.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	if err := s.Validate(&params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	schoolSrv, ok := client.SchoolServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, s.Error(ctx).InternalServerError("Not found school service"))
	}

	resp, err := schoolSrv.SearchSchools(ctx, &schoolPB.SearchSchoolsRequest{
		Search: params.SearchName,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	schools := make([]*dto.School, 0)

	for _, v := range resp.Schools {
		schools = append(schools, dto.PBToSchool(v))
	}

	return SuccessResponse(rsp, D{
		"Schools": schools,
	})
}

func (s *SchoolSrv) Create(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	params := struct {
		Name        string `json:"name,omitempty" validate:"nonzero"`
		SchoolCode  string `json:"schoolCode,omitempty"`
		Description string `json:"description,omitempty"`
	}{}

	if err := s.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	if err := s.Validate(&params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	schoolSrv, ok := client.SchoolServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, s.Error(ctx).InternalServerError("Not found school service"))
	}

	createResp, err := schoolSrv.CreateSchool(ctx, &schoolPB.CreateSchoolRequest{
		Name:        params.Name,
		Description: params.Description,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if !createResp.OK {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest("创建失败"))
	}
	return SuccessResponse(rsp, params)
}
