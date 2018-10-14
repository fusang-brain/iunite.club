package handler

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"

	"github.com/go-log/log"

	"iunite.club/services/restful/dto"

	clubPB "iunite.club/services/organization/proto/club"
	deptPB "iunite.club/services/organization/proto/department"
	userPB "iunite.club/services/user/proto"
)

type OrganizationHandler struct {
	BaseHandler
	clubService       clubPB.ClubService
	userService       userPB.UserSrvService
	departmentService deptPB.DepartmentService
}

func NewOrganizationHandler(c client.Client) *OrganizationHandler {
	return &OrganizationHandler{
		clubService:       clubPB.NewClubService(OrganizationService, c),
		userService:       userPB.NewUserSrvService(UserSerivce, c),
		departmentService: deptPB.NewDepartmentService(OrganizationService, c),
	}
}

func (o *OrganizationHandler) CreateOrganization(req *go_api.Request, rsp *go_api.Response) {
	clubService := o.clubService
	ctx := context.Background()
	params := struct {
		Name      string   `json:"name,omitempty" validate:"nonzero"`
		Scale     int32    `json:"scale,omitempty" validate:"nonzero"`
		Paperwork []string `json:"paperwork,omitempty" validate:"nonzero"`
		Logo      string   `json:"logo,omitempty" validate:"nonzero"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	// 字段验证
	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	// get current user school id
	cliams, err := o.GetTokenCliamsFromRequest(req)

	if err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	userService := o.userService

	userResp, err := userService.FindUserByID(ctx, &userPB.QueryUserRequest{
		Id: cliams.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	createResp, err := clubService.CreateClub(ctx, &clubPB.CreateClubRequest{
		Name:       params.Name,
		Paperworks: params.Paperwork,
		Scale:      params.Scale,
		SchoolID:   userResp.User.SchoolID,
		CreatorID:  userResp.User.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if createResp.OK {
		SuccessResponse(rsp, D{})
		return
	}

	ErrorResponse(rsp, o.Error().TemplateBadRequest("CreateError"))
	return
}

func (o *OrganizationHandler) GetAllOrgByUserID(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id,omitempty" query:"id"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if params.ID == "" {
		// tokens := o.GetTokenFromRequest(req)
		if token, err := o.GetTokenCliamsFromRequest(req); err == nil {
			log.Logf("UserID is %v", token.UserID)
			params.ID = token.UserID
		}
	}

	orgService := o.clubService

	clubListResp, err := orgService.GetClubsByUserID(ctx, &clubPB.GetClubsByUserIDRequest{
		UserID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
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

	SuccessResponse(rsp, D{
		"Organizations": resultClubs,
	})
	return
}

func (o *OrganizationHandler) GetAllOrgUsersByUserID(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `query:"id" validate:"objectid"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, err.Error())
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, err.Error())
		return
	}

	if params.ID == "" {
		if token, err := o.GetTokenCliamsFromRequest(req); err != nil {
			params.ID = token.UserID
		}
	}

	clubSrv := o.clubService

	userClubProfileResp, err := clubSrv.GetUserClubProfilesByUserID(ctx, &clubPB.GetUserClubProfilesByUserIDRequest{UserID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	userOrgs := make([]*dto.OrganizationUser, 0)

	for _, v := range userClubProfileResp.UserClubProfiles {
		userOrgs = append(userOrgs, dto.PBToOrganizationUser(v))
	}
	SuccessResponse(rsp, D{
		"Organizations": userOrgs,
	})
}

func (o *OrganizationHandler) SearchHotOrganization(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Search  string `json:"search,omitempty" query:"search"`
		MaxSize int    `json:"maxSize,omitempty" query:"maxSize"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().TemplateBadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	fmt.Println(params, ">>> params")

	clubService := o.clubService

	clubListResponse, err := clubService.SearchClubs(ctx, &clubPB.SearchClubRequest{
		Search: params.Search,
		Page:   1,
		Limit:  int64(params.MaxSize),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
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

	SuccessResponse(rsp, D{
		"Organizations": resultClubs,
	})
}

func (o *OrganizationHandler) AcceptJoin(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		SchoolID     string `json:"schoolID,omitempty"`
		Organization string `json:"organizationID,omitempty"`
		Department   string `json:"departmentID,omitempty"`
		Job          string `json:"jobID,omitempty"`
		UserID       string `json:"userID,omitempty"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if params.UserID == "" {
		if token, err := o.GetTokenCliamsFromRequest(req); err == nil {
			params.UserID = token.UserID
		}
	}

	orgService := o.clubService

	resp, err := orgService.AcceptJoinOneClub(ctx, &clubPB.AcceptJoinOneClubRequest{
		UserID:       params.UserID,
		ClubID:       params.Organization,
		JobID:        params.Job,
		DepartmentID: params.Department,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !resp.OK {
		ErrorResponse(rsp, o.Error().BadRequest("Action failed"))
		return
	}

	SuccessResponse(rsp, D{})
}

func (o *OrganizationHandler) AgreeJoin(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	orgService := o.clubService

	resp, err := orgService.ExecuteJoinClubAccept(ctx, &clubPB.ExecuteJoinClubAcceptRequest{
		IsPassed: true,
		AcceptID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !resp.OK {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	SuccessResponse(rsp, D{})
}

func (o *OrganizationHandler) RefuseJoin(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	orgService := o.clubService

	resp, err := orgService.ExecuteJoinClubAccept(ctx, &clubPB.ExecuteJoinClubAcceptRequest{
		IsPassed: false,
		AcceptID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if !resp.OK {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (o *OrganizationHandler) FindRefusedAccept(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		UserID string `json:"id" validate:"nonzero"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	orgService := o.clubService

	acceptResp, err := orgService.FindRefusedAcceptByUserID(ctx, &clubPB.FindRefusedAcceptRequest{
		UserID: params.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
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

	SuccessResponse(rsp, D{
		"OrganizationAccepts": organizationUsers,
	})
}

func (o *OrganizationHandler) GetDepartmentDetails(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	ctx := context.Background()
	departmentSrv := o.departmentService

	params := struct {
		ID string `json:"departmentID,omitempty" query:"departmentID" validate:"nonzero,objectid"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	fmt.Println(params.ID)

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	deptResp, err := departmentSrv.GetDepartmentDetails(ctx, &deptPB.GetDepartmentWithIDRequest{ID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Details": dto.PBToDepartment(deptResp.Department),
	})

}

func (o *OrganizationHandler) Info(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	ctx := context.Background()
	params := struct {
		ID string `query:"id" validate:"nonzero,objectid"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	clubSrv := o.clubService

	clubResp, err := clubSrv.FindClubDetailsByID(ctx, &clubPB.GetClubByIDRequest{ID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, dto.PBToOrganization(clubResp.Club))
}

func (o *OrganizationHandler) UploadLogo(*go_api.Request, *go_api.Response) {
	// TODO 文件服务实现之后实现该接口
	panic("not implemented")
}

func (o *OrganizationHandler) UpdateOrganizationDescription(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	ctx := context.Background()
	params := struct {
		ID          string `json:"id,omitempty" validate:"nonzero,objectid"`
		Description string `json:"description,omitempty"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	clubSrv := o.clubService

	updateResp, err := clubSrv.UpdateClubInfo(ctx, &clubPB.UpdateClubInfoRequest{ID: params.ID, ToSet: []byte(`{"description": "` + params.Description + `"}`)})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, updateResp)
	return
}

func (o *OrganizationHandler) GetAllOrganizationBySchool(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		SchoolID string `query:"schoolID" validate:"nonzero,objectid"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	clubSrv := o.clubService

	clubListResp, err := clubSrv.FindClubsBySchoolID(ctx, &clubPB.GetClubsBySchoolIDRequest{SchoolID: params.SchoolID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	orgs := make([]*dto.Organization, 0)

	for _, o := range clubListResp.Organizations {
		orgs = append(orgs, dto.PBToOrganization(o))
	}

	SuccessResponse(rsp, D{
		"Organizations": orgs,
	})
}

func (o *OrganizationHandler) GetOrganizationDetails(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `query:"id" validate:"nonzero,objectid"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}
	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	clubSrv := o.clubService

	clubResp, err := clubSrv.FindClubDetailsByID(ctx, &clubPB.GetClubByIDRequest{ID: params.ID})
	if err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	SuccessResponse(rsp, dto.PBToOrganization(clubResp.Club))
}

func (o *OrganizationHandler) GetOrganizationUserInfoDetails(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID     string `query:"orgID" validate:"nonzero,objectid"`
		UserID string `query:"userID" validate:"objectid"`
	}{}

	if err := o.Bind(req, &params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if err := o.Validate(&params); err != nil {
		ErrorResponse(rsp, o.Error().BadRequest(err.Error()))
		return
	}

	if params.UserID == "" {
		if token, err := o.GetTokenCliamsFromRequest(req); err == nil {
			params.UserID = token.UserID
		} else {
			ErrorResponse(rsp, o.Error().Unauthorized("非法访问"))
			return
		}
	}

	clubSrv := o.clubService

	resp, err := clubSrv.GetUserClubProfileDetailsByID(ctx, &clubPB.GetUserClubProfileDetailsByIDRequest{
		OrganizationID: params.ID,
		UserID:         params.UserID,
	})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	details := dto.PBToOrganizationUser(resp.UserClubProfile)

	SuccessResponse(rsp, D{
		"Details": details,
	})
}

func (o *OrganizationHandler) SelectOrganizations(*go_api.Request, *go_api.Response) {

	// TODO 最后实现老师端的接口
	panic("not implemented")
}
