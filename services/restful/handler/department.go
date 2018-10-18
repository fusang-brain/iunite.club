package handler

import (
	"context"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"

	"iunite.club/services/restful/dto"

	deptPB "iunite.club/services/organization/proto/department"
)

type DepartmentHandler struct {
	BaseHandler
	departmentService deptPB.DepartmentService
}

func NewDepartmentHandler(c client.Client) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: deptPB.NewDepartmentService(OrganizationService, c),
	}
}

func (d *DepartmentHandler) getDepartmentService() deptPB.DepartmentService {
	return d.departmentService
}

func (d *DepartmentHandler) AddDept(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Organization string `json:"organization" validate:"objectid"`
		Name         string `json:"name"`
		Parent       string `json:"parent" validate:"objectid"`
		Description  string `json:"description"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	parentID := params.Parent

	if params.Organization == "" {
		params.Organization = d.GetCurrentClubIDFromRequest(req)
	}

	if parentID == "" {
		parentID = params.Organization
	}

	createResp, err := deptSrv.CreateDepartment(ctx, &deptPB.CreateDepartmentRequest{
		Name:        params.Name,
		ParentID:    parentID,
		Description: params.Description,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !createResp.OK {
		ErrorResponse(rsp, d.Error().BadRequest("CreateError"))
		return
	}

	SuccessResponse(rsp, D{
		"Department": dto.PBToDepartment(createResp.Department),
	})
}

func (d *DepartmentHandler) GetDepartmentByOrg(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ClubID string `query:"org" validate:"nonzero,objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	deptsResp, err := deptSrv.GetDepartmentListByParentID(ctx, &deptPB.DepartmentListByParentIDRequest{
		Page:     1,
		Limit:    500,
		ParentID: params.ClubID,
		Spread:   false,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	departments := make([]*dto.Department, 0)

	for _, dept := range deptsResp.Departments {
		departments = append(departments, dto.PBToDepartment(dept))
	}

	SuccessResponse(rsp, D{
		"Departments": departments,
	})
}

func (d *DepartmentHandler) GetAllDepartmentByOrg(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ClubID string `query:"org" validate:"nonzero,objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if params.ClubID == "" {
		params.ClubID = d.GetCurrentClubIDFromRequest(req)
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	deptsResp, err := deptSrv.GetDepartmentListByParentID(ctx, &deptPB.DepartmentListByParentIDRequest{
		Page:     1,
		Limit:    500,
		ParentID: params.ClubID,
		Spread:   true,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	departments := make([]*dto.Department, 0)

	for _, dept := range deptsResp.Departments {
		departments = append(departments, dto.PBToDepartment(dept))
	}

	SuccessResponse(rsp, D{
		"Departments": departments,
	})
}

func (d *DepartmentHandler) SearchDepartment(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ClubID string `query:"org" validate:"objectid"`
		Name   string `query:"name"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if params.ClubID == "" {
		params.ClubID = d.GetCurrentClubIDFromRequest(req)
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	deptsResp, err := deptSrv.GetDepartmentListByParentID(ctx, &deptPB.DepartmentListByParentIDRequest{
		Page:     1,
		Limit:    500,
		ParentID: params.ClubID,
		Spread:   true,
		Search:   params.Name,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	departments := make([]*dto.Department, 0)

	for _, dept := range deptsResp.Departments {
		departments = append(departments, dto.PBToDepartment(dept))
	}

	SuccessResponse(rsp, D{
		"Departments": departments,
	})
}

func (d *DepartmentHandler) AddUserToDepartment(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Users        []string `json:"users"`
		DepartmentID string   `json:"dept" validate:"objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	departmentSrv := d.getDepartmentService()

	_, err := departmentSrv.AddUsersToDepartment(ctx, &deptPB.UserFromDepartmentRequest{
		Users:        params.Users,
		DepartmentID: params.DepartmentID,
		ClubID:       d.GetCurrentClubIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	// if !updateResp.OK {
	// 	return ErrorResponse(rsp, d.Error(ctx).BadRequest("UpdateError"))
	// }

	SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) RemoveUserFromDepartment(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Users        []string `json:"users"`
		DepartmentID string   `json:"dept" validate:"objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	departmentSrv := d.getDepartmentService()

	_, err := departmentSrv.RemoveUsersFromDepartment(ctx, &deptPB.UserFromDepartmentRequest{
		Users:        params.Users,
		DepartmentID: params.DepartmentID,
		ClubID:       d.GetCurrentClubIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) AllCanSelectedUsers(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Page  int64 `json:"page,omitempty"`
		Limit int64 `json:"limit,omitempty"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	resp, err := deptSrv.GetAllCanSelectUsers(ctx, &deptPB.ListByClubIDRequest{
		ClubID: d.GetCurrentClubIDFromRequest(req),
		Page:   params.Page,
		Limit:  params.Limit,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	users := make([]*dto.User, 0)

	for _, u := range resp.Users {
		users = append(users, dto.PBToUser(u))
	}

	SuccessResponseWithPage(rsp, params.Page, params.Limit, resp.Total, users)
}

func (d *DepartmentHandler) GetAllUsersWithDepartment(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Page         int64  `json:"page,omitempty" query:"page"`
		Limit        int64  `json:"limit,omitempty" query:"limit"`
		DepartmentID string `json:"dept,omitempty" query:"dept"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	resp, err := deptSrv.GetUsersByDepartmentID(ctx, &deptPB.ListByDepartmentIDRequest{
		DepartmentID: params.DepartmentID,
		Page:         params.Page,
		Limit:        params.Limit,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	users := make([]*dto.User, 0)

	for _, u := range resp.Users {
		users = append(users, dto.PBToUser(u))
	}

	SuccessResponseWithPage(rsp, params.Page, params.Limit, resp.Total, users)
}

func (d *DepartmentHandler) Update(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID           string `json:"id" validate:"nonzero,objectid"`
		Organization string `json:"organization" validate:"objectid"`
		Name         string `json:"name"`
		Parent       string `json:"parent" validate:"objectid"`
		Description  string `json:"description"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if params.Parent == "" {
		params.Parent = params.Organization
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	_, err := deptSrv.UpdateDepartment(ctx, &deptPB.UpdateDepartmentRequest{
		ParentID:    params.Parent,
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

func (d *DepartmentHandler) Remove(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id" validate:"nonzero,objectid"`
	}{}

	if err := d.Bind(req, &params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	if err := d.Validate(&params); err != nil {
		ErrorResponse(rsp, d.Error().BadRequest(err.Error()))
		return
	}

	deptSrv := d.getDepartmentService()

	_, err := deptSrv.RemoveDepartment(ctx, &deptPB.RemoveDepartmentRequest{ID: params.ID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}
