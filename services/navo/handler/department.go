package handler

import (
	"context"

	"iunite.club/services/navo/dto"

	"iunite.club/services/navo/client"

	go_api "github.com/micro/go-api/proto"
	deptPB "iunite.club/services/organization/proto/department"
)

type DepartmentHandler struct {
	BaseHandler
	departmentService deptPB.DepartmentService
}

func (d *DepartmentHandler) getDepartmentService(ctx context.Context) deptPB.DepartmentService {
	if d.departmentService == nil {
		deptSrv, ok := client.DepartmentServiceFromContext(ctx)

		if !ok {
			panic("Not found department service")
		}

		d.departmentService = deptSrv
	}

	return d.departmentService
}

func (d *DepartmentHandler) AddDept(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {

	params := struct {
		Organization string `json:"organization" validate:"nonzero,objectid"`
		Name         string `json:"name"`
		Parent       string `json:"parent" validate:"objectid"`
		Description  string `json:"description"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	parentID := params.Parent

	if parentID == "" {
		parentID = params.Organization
	}

	createResp, err := deptSrv.CreateDepartment(ctx, &deptPB.CreateDepartmentRequest{
		Name:        params.Name,
		ParentID:    parentID,
		Description: params.Description,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if !createResp.OK {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest("CreateError"))
	}

	return SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) GetDepartmentByOrg(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ClubID string `query:"org" validate:"nonzero,objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	deptsResp, err := deptSrv.GetDepartmentListByParentID(ctx, &deptPB.DepartmentListByParentIDRequest{
		Page:     1,
		Limit:    500,
		ParentID: params.ClubID,
		Spread:   false,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	departments := make([]*dto.Department, 0)

	for _, dept := range deptsResp.Departments {
		departments = append(departments, dto.PBToDepartment(dept))
	}

	return SuccessResponse(rsp, D{
		"Departments": departments,
	})
}

func (d *DepartmentHandler) GetAllDepartmentByOrg(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ClubID string `query:"org" validate:"nonzero,objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if params.ClubID == "" {
		params.ClubID = d.GetCurrentClubIDFromRequest(ctx, req)
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	deptsResp, err := deptSrv.GetDepartmentListByParentID(ctx, &deptPB.DepartmentListByParentIDRequest{
		Page:     1,
		Limit:    500,
		ParentID: params.ClubID,
		Spread:   true,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	departments := make([]*dto.Department, 0)

	for _, dept := range deptsResp.Departments {
		departments = append(departments, dto.PBToDepartment(dept))
	}

	return SuccessResponse(rsp, D{
		"Departments": departments,
	})
}

func (d *DepartmentHandler) SearchDepartment(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ClubID string `query:"org" validate:"objectid"`
		Name   string `query:"name"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if params.ClubID == "" {
		params.ClubID = d.GetCurrentClubIDFromRequest(ctx, req)
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	deptsResp, err := deptSrv.GetDepartmentListByParentID(ctx, &deptPB.DepartmentListByParentIDRequest{
		Page:     1,
		Limit:    500,
		ParentID: params.ClubID,
		Spread:   true,
		Search:   params.Name,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	departments := make([]*dto.Department, 0)

	for _, dept := range deptsResp.Departments {
		departments = append(departments, dto.PBToDepartment(dept))
	}

	return SuccessResponse(rsp, D{
		"Departments": departments,
	})
}

func (d *DepartmentHandler) AddUserToDepartment(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Users        []string `json:"users"`
		DepartmentID string   `json:"dept" validate:"objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	departmentSrv := d.getDepartmentService(ctx)

	_, err := departmentSrv.AddUsersToDepartment(ctx, &deptPB.UserFromDepartmentRequest{
		Users:        params.Users,
		DepartmentID: params.DepartmentID,
		ClubID:       d.GetCurrentClubIDFromRequest(ctx, req),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	// if !updateResp.OK {
	// 	return ErrorResponse(rsp, d.Error(ctx).BadRequest("UpdateError"))
	// }

	return SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) RemoveUserFromDepartment(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Users        []string `json:"users"`
		DepartmentID string   `json:"dept" validate:"objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	departmentSrv := d.getDepartmentService(ctx)

	_, err := departmentSrv.RemoveUsersFromDepartment(ctx, &deptPB.UserFromDepartmentRequest{
		Users:        params.Users,
		DepartmentID: params.DepartmentID,
		ClubID:       d.GetCurrentClubIDFromRequest(ctx, req),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	// if !updateResp.OK {
	// 	return ErrorResponse(rsp, d.Error(ctx).BadRequest("UpdateError"))
	// }

	return SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) AllCanSelectedUsers(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Page  int64 `json:"page,omitempty"`
		Limit int64 `json:"limit,omitempty"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	resp, err := deptSrv.GetAllCanSelectUsers(ctx, &deptPB.ListByClubIDRequest{
		ClubID: d.GetCurrentClubIDFromRequest(ctx, req),
		Page:   params.Page,
		Limit:  params.Limit,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	users := make([]*dto.User, 0)

	for _, u := range resp.Users {
		users = append(users, dto.PBToUser(u))
	}

	return SuccessResponseWithPage(rsp, params.Page, params.Limit, resp.Total, users)
}

func (d *DepartmentHandler) GetAllUsersWithDepartment(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Page         int64  `json:"page,omitempty" query:"page"`
		Limit        int64  `json:"limit,omitempty" query:"limit"`
		DepartmentID string `json:"dept,omitempty" query:"dept"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	resp, err := deptSrv.GetUsersByDepartmentID(ctx, &deptPB.ListByDepartmentIDRequest{
		DepartmentID: params.DepartmentID,
		Page:         params.Page,
		Limit:        params.Limit,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	users := make([]*dto.User, 0)

	for _, u := range resp.Users {
		users = append(users, dto.PBToUser(u))
	}

	return SuccessResponseWithPage(rsp, params.Page, params.Limit, resp.Total, users)
}

func (d *DepartmentHandler) Update(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID           string `json:"id" validate:"nonzero,objectid"`
		Organization string `json:"organization" validate:"objectid"`
		Name         string `json:"name"`
		Parent       string `json:"parent" validate:"objectid"`
		Description  string `json:"description"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if params.Parent == "" {
		params.Parent = params.Organization
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	_, err := deptSrv.UpdateDepartment(ctx, &deptPB.UpdateDepartmentRequest{
		ParentID:    params.Parent,
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) Remove(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID string `json:"id" validate:"nonzero,objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	if err := d.Validate(&params); err != nil {
		return ErrorResponse(rsp, d.Error(ctx).BadRequest(err.Error()))
	}

	deptSrv := d.getDepartmentService(ctx)

	_, err := deptSrv.RemoveDepartment(ctx, &deptPB.RemoveDepartmentRequest{ID: params.ID})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{})
}
