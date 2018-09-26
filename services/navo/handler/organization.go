package handler

import (
	"context"
	"fmt"

	"github.com/go-log/log"

	"iunite.club/services/navo/dto"

	go_api "github.com/micro/go-api/proto"
	"iunite.club/services/navo/client"
	clubPB "iunite.club/services/organization/proto/club"
	userPB "iunite.club/services/user/proto"
)

type OrganizationHandler struct {
	BaseHandler
}

func (o *OrganizationHandler) CreateOrganization(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	clubService, ok := client.ClubServiceFromContext(ctx)
	if !ok {
		return o.Error(ctx).InternalServerError("ClubService not found")
	}

	params := struct {
		Name      string   `json:"name,omitempty" validate:"nonzero"`
		Scale     int32    `json:"scale,omitempty" validate:"nonzero"`
		Paperwork []string `json:"paperwork,omitempty" validate:"nonzero"`
		Logo      string   `json:"logo,omitempty" validate:"nonzero"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	// 字段验证
	if err := o.Validate(&params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	// get current user school id
	cliams, err := o.GetTokenCliamsFromRequest(req)

	if err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	userService, ok := client.UserServiceFromContext(ctx)

	if !ok {
		return o.Error(ctx).InternalServerError("ClubService not found")
	}

	userResp, err := userService.FindUserByID(ctx, &userPB.QueryUserRequest{
		Id: cliams.UserID,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	createResp, err := clubService.CreateClub(ctx, &clubPB.CreateClubRequest{
		Name:       params.Name,
		Paperworks: params.Paperwork,
		Scale:      params.Scale,
		SchoolID:   userResp.User.SchoolID,
		CreatorID:  userResp.User.ID,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if createResp.OK {
		return SuccessResponse(rsp, D{})
	}

	return ErrorResponse(rsp, o.Error(ctx).TemplateBadRequest("CreateError"))
}

func (o *OrganizationHandler) GetAllOrgByUserID(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID string `json:"id,omitempty" query:"id"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if params.ID == "" {
		// tokens := o.GetTokenFromRequest(req)
		if token, err := o.GetTokenCliamsFromRequest(req); err == nil {
			log.Logf("UserID is %v", token.UserID)
			params.ID = token.UserID
		}
	}

	orgService, ok := client.ClubServiceFromContext(ctx)
	if !ok {
		return ErrorResponse(rsp, o.Error(ctx).InternalServerError("ClubService not found"))
	}

	clubListResp, err := orgService.GetClubsByUserID(ctx, &clubPB.GetClubsByUserIDRequest{
		UserID: params.ID,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	resultClubs := make([]dto.Organization, 0, 1)

	for _, val := range clubListResp.Organizations {
		resultClubs = append(resultClubs, dto.Organization{
			Name:        val.Name,
			SlugName:    val.Slug,
			ID:          val.ID,
			Description: val.Description,
			SchoolRefer: val.SchoolID,
			Logo:        "",
		})
	}

	return SuccessResponse(rsp, D{
		"Organizations": resultClubs,
	})
}

func (o *OrganizationHandler) GetAllOrgUsersByUserID(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) SearchHostOrganization(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Search  string `json:"search,omitempty" query:"search"`
		MaxSize int    `json:"maxSize,omitempty" query:"maxSize"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).TemplateBadRequest(err.Error()))
	}

	if err := o.Validate(&params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	fmt.Println(params, ">>> params")

	clubService, ok := client.ClubServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, o.Error(ctx).InternalServerError("ClubService not found"))
	}

	clubListResponse, err := clubService.SearchClubs(ctx, &clubPB.SearchClubRequest{
		Search: params.Search,
		Page:   1,
		Limit:  int64(params.MaxSize),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	resultClubs := make([]dto.Organization, 0, 1)

	for _, val := range clubListResponse.Organizations {
		resultClubs = append(resultClubs, dto.Organization{
			Name:        val.Name,
			SlugName:    val.Slug,
			ID:          val.ID,
			Description: val.Description,
			SchoolRefer: val.SchoolID,
			Logo:        "",
		})
	}

	return SuccessResponse(rsp, D{
		"Organizations": resultClubs,
	})
}

func (o *OrganizationHandler) AcceptJoin(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		SchoolID     string `json:"schoolID,omitempty"`
		Organization string `json:"organizationID,omitempty"`
		Department   string `json:"departmentID,omitempty"`
		Job          string `json:"jobID,omitempty"`
		UserID       string `json:"userID,omitempty"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if err := o.Validate(&params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if params.UserID == "" {
		if token, err := o.GetTokenCliamsFromRequest(req); err == nil {
			params.UserID = token.UserID
		}
	}

	orgService, ok := client.ClubServiceFromContext(ctx)
	if !ok {
		return ErrorResponse(rsp, o.Error(ctx).InternalServerError("Club service not found"))
	}

	resp, err := orgService.AcceptJoinOneClub(ctx, &clubPB.AcceptJoinOneClubRequest{
		UserID:       params.UserID,
		ClubID:       params.Organization,
		JobID:        params.Job,
		DepartmentID: params.Department,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if !resp.OK {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest("Action failed"))
	}

	return SuccessResponse(rsp, D{})
}

func (o *OrganizationHandler) AgreeJoin(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID string `json:"id"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if err := o.Validate(&params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	orgService, ok := client.ClubServiceFromContext(ctx)
	if !ok {
		return ErrorResponse(rsp, o.Error(ctx).InternalServerError("Club service not found"))
	}

	resp, err := orgService.ExecuteJoinClubAccept(ctx, &clubPB.ExecuteJoinClubAcceptRequest{
		IsPassed: true,
		AcceptID: params.ID,
	})

	if err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if !resp.OK {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	return SuccessResponse(rsp, D{})
}

func (o *OrganizationHandler) RefuseJoin(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID string `json:"id"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if err := o.Validate(&params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	orgService, ok := client.ClubServiceFromContext(ctx)
	if !ok {
		return ErrorResponse(rsp, o.Error(ctx).InternalServerError("Club service not found"))
	}

	resp, err := orgService.ExecuteJoinClubAccept(ctx, &clubPB.ExecuteJoinClubAcceptRequest{
		IsPassed: false,
		AcceptID: params.ID,
	})

	if err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if !resp.OK {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	return SuccessResponse(rsp, D{})
}

func (o *OrganizationHandler) FindRefusedAccept(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {

	params := struct {
		UserID string `json:"id" validate:"nonzero"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	if err := o.Validate(&params); err != nil {
		return ErrorResponse(rsp, o.Error(ctx).BadRequest(err.Error()))
	}

	orgService, ok := client.ClubServiceFromContext(ctx)
	if !ok {
		return ErrorResponse(rsp, o.Error(ctx).InternalServerError("Club service not found"))
	}

	acceptResp, err := orgService.FindRefusedAcceptByUserID(ctx, &clubPB.FindRefusedAcceptRequest{
		UserID: params.UserID,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	organizationUsers := make([]dto.OrganizationUser, 0, 1)

	for _, v := range acceptResp.Accepts {
		organizationUsers = append(organizationUsers, dto.OrganizationUser{
			ID:          v.ID,
			Kind:        0,
			AcceptState: 0,
			State:       0,
			OrganizationInfo: &dto.Organization{
				ID:          v.Organization.ID,
				Name:        v.Organization.Name,
				SlugName:    v.Organization.Slug,
				Logo:        v.Organization.ClubProfile.Logo,
				Scale:       int(v.Organization.ClubProfile.Scale),
				SchoolRefer: v.Organization.SchoolID,
				Description: v.Organization.Description,
			},
			IsCreator: v.Kind == 1,
		})
	}

	return SuccessResponse(rsp, D{
		"OrganizationAccepts": organizationUsers,
	})
}

func (o *OrganizationHandler) GetDepartmentDetails(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) Info(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) UploadLogo(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) UpdateOrganizationDescription(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) GetAllOrganizationBySchool(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) GetOrganizationDetails(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) GetOrganizationUserInfoDetails(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (o *OrganizationHandler) SelectOrganizations(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}
