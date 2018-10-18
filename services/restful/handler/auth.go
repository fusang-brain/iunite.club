package handler

import (
	"context"
	"strconv"

	orgPB "iunite.club/services/organization/proto"

	clubPB "iunite.club/services/organization/proto/club"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"

	"github.com/iron-kit/go-ironic/utils"
	schoolPB "iunite.club/services/organization/proto/school"
	"iunite.club/services/restful/dto"

	userPB "iunite.club/services/user/proto"
	authPB "iunite.club/services/user/proto/secruity"
	// authPB "iunite.club/srv/secruity/proto/auth"
)

type AuthHandler struct {
	BaseHandler

	authService   authPB.SecruityService
	userService   userPB.UserSrvService
	schoolService schoolPB.SchoolSrvService
	clubService   clubPB.ClubService
}

func NewAuthHandler(c client.Client) *AuthHandler {
	return &AuthHandler{
		authService:   authPB.NewSecruityService(UserSerivce, c),
		userService:   userPB.NewUserSrvService(UserSerivce, c),
		schoolService: schoolPB.NewSchoolSrvService(OrganizationService, c),
		clubService:   clubPB.NewClubService(OrganizationService, c),
	}
}

func (a *AuthHandler) Login(req *go_api.Request, resp *go_api.Response) {
	ctx := context.Background()
	authService := a.authService

	params := struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(resp, a.Error().BadRequest(err.Error()))
		return
	}

	loginResp, err := authService.Signin(ctx, &authPB.AuthRequest{
		Identify: params.Mobile,
		Password: params.Password,
	})

	if err != nil {
		ErrorResponse(resp, err)
		return
	}

	userIDResp, err := authService.GetUserIDFromToken(ctx, &authPB.TokenRequest{Token: loginResp.Token})

	if err != nil {
		ErrorResponse(resp, err)
		return
	}

	// fmt.Println(userIDResp.UserID)

	uR, err := a.userService.FindUserByID(ctx, &userPB.QueryUserRequest{Id: userIDResp.UserID})

	if err != nil {
		ErrorResponse(resp, err)
		return
	}

	schoolService := a.schoolService
	var schoolResp *schoolPB.SchoolResponse
	if sR, e := schoolService.GetSchoolByID(ctx, &schoolPB.GetSchoolRequest{ID: uR.User.SchoolID}); e == nil {
		schoolResp = sR
	} else {
		sR := new(schoolPB.SchoolResponse)
		schoolResp = sR
	}

	clubProfileResp, err := a.clubService.GetUserClubProfilesByUserID(ctx, &clubPB.GetUserClubProfilesByUserIDRequest{UserID: uR.User.ID})
	if err != nil {
		ErrorResponse(resp, err)
		return
	}
	currentUserClubProfile := new(orgPB.UserClubProfile)

	userClubProfiles := clubProfileResp.UserClubProfiles
	for _, v := range userClubProfiles {
		if v.OrganizationID == uR.User.DefaultClubID {
			currentUserClubProfile = v
		}
	}

	clubsResp, err := a.clubService.GetClubsByUserID(ctx, &clubPB.GetClubsByUserIDRequest{UserID: uR.User.ID})

	if err != nil {
		ErrorResponse(resp, err)
		return
	}
	org := new(dto.Organization)
	orgs := make([]*dto.Organization, 0)
	for _, v := range clubsResp.Organizations {
		if v.ID == uR.User.DefaultClubID {
			org = dto.PBToOrganization(v)
		}
		orgs = append(orgs, dto.PBToOrganization(v))
	}

	// pCreatedAt := hptypes.Timestamp(uR.User.Profile.CreatedAt)
	// pUpdatedAt := hptypes.Timestamp(uR.User.Profile.UpdatedAt)
	SuccessResponse(resp, map[string]interface{}{
		"Token":                loginResp.Token,
		"TokenTime":            loginResp.ExpiredAt,
		"IsMaster":             currentUserClubProfile.IsMaster,
		"OrganizationUserInfo": dto.PBToOrganizationUser(currentUserClubProfile),
		"User": dto.User{
			ID:                  uR.User.ID,
			Username:            uR.User.Username,
			CreatedAt:           utils.ISOTime2MicroUnix(uR.User.CreatedAt),
			UpdatedAt:           utils.ISOTime2MicroUnix(uR.User.UpdatedAt),
			IsTeacher:           false,
			IsAdmin:             false,
			Mobile:              uR.User.Profile.Mobile,
			AreaCode:            "+86",
			Organizations:       orgs,
			CurrentOrganization: org,
			Profile:             dto.PBToProfile(uR.User.Profile),
			// Profile: &dto.Profile{
			// 	ID:        uR.User.Profile.ID,
			// 	CreatedAt: utils.Time2MicroUnix(pCreatedAt),
			// 	UpdatedAt: utils.Time2MicroUnix(pUpdatedAt),
			// 	// CreatedAt: utils.ISOTime2MicroUnix(uR.User.Profile.CreatedAt),
			// 	// UpdatedAt: utils.ISOTime2MicroUnix(uR.User.Profile.UpdatedAt),
			// 	UserNO:    "-",
			// 	Avatar:    uR.User.Profile.Avatar,
			// 	FirstName: uR.User.Profile.Firstname,
			// 	LastName:  uR.User.Profile.Lastname,
			// 	Gender:    uR.User.Profile.Gender,

			// },
			School: &dto.School{
				Name:       schoolResp.School.Name,
				SlugName:   schoolResp.School.SlugName,
				SchoolCode: schoolResp.School.SchoolCode,
				ID:         schoolResp.School.ID,
				// CreatedAt: utils.ISOTime2MicroUnix(schoolResp.School.)
			},
		},
	})
	return
}

func (a *AuthHandler) Register(req *go_api.Request, resp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Mobile          string `json:"mobile,omitempty" validate:"nonzero"`
		Code            int    `json:"code,omitempty" validate:"nonzero"`
		Password        string `json:"password,omitempty" validate:"nonzero"`
		ConfirmPassword string `json:"confirmPassword,omitempty" validate:"nonzero"`
		FirstName       string `json:"firstName,omitempty" validate:"nonzero"`
		LastName        string `json:"lastName,omitempty" validate:"nonzero"`
		IsTeacher       bool   `json:"isTeacher,omitempty"`
		School          string `json:"school,omitempty" validate:"objectid,nonzero"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(resp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(resp, a.Error().BadRequest(err.Error()))
		return
	}

	registerResp, err := a.authService.SignupWithMobile(ctx, &authPB.SignupWithMobileRequest{
		Mobile:          params.Mobile,
		Password:        params.Password,
		ConfirmPassword: params.ConfirmPassword,
		FirstName:       params.FirstName,
		LastName:        params.LastName,
		School:          params.School,
		ValidateCode:    strconv.Itoa(params.Code),
	})

	if err != nil {
		ErrorResponse(resp, err)
		return
	}

	if !registerResp.OK {
		ErrorResponse(resp, "注册失败")
		return
	}

	SuccessResponse(resp, D{})
}
