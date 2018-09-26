package handler

import (
	"context"
	"encoding/json"
	"github.com/iron-kit/go-ironic/utils"
	"iunite.club/services/navo/dto"
	schoolPB "iunite.club/services/organization/proto/school"

	go_api "github.com/micro/go-api/proto"
	"iunite.club/services/navo/client"
	userPB "iunite.club/services/user/proto"
	authPB "iunite.club/services/user/proto/secruity"
	// authPB "iunite.club/srv/secruity/proto/auth"
)

type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Login(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {

	authService, ok := client.SecruityAuthServiceFromContext(ctx)

	if !ok {
		return a.Error(ctx).InternalServerError("Not found AuthService")
	}

	params := struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}{}

	err := json.Unmarshal([]byte(req.Body), &params)

	if err != nil {
		return ErrorResponse(resp, a.Error(ctx).BadRequest(err.Error()))
	}

	loginResp, err := authService.Signin(ctx, &authPB.AuthRequest{
		Identify: params.Mobile,
		Password: params.Password,
	})

	if err != nil {
		return ErrorResponse(resp, err)
	}

	userIDResp, err := authService.GetUserIDFromToken(ctx, &authPB.TokenRequest{Token: loginResp.Token})

	if err != nil {
		return ErrorResponse(resp, err)
	}

	// fmt.Println(userIDResp.UserID)

	userService, ok := client.UserServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(resp, a.Error(ctx).TemplateBadRequest("Not found user server"))
	}

	uR, err := userService.FindUserByID(ctx, &userPB.QueryUserRequest{Id: userIDResp.UserID})

	if err != nil {
		return err
	}

	schoolService, _ := client.SchoolServiceFromContext(ctx)

	var schoolResp *schoolPB.SchoolResponse

	if sR, e := schoolService.GetSchoolByID(ctx, &schoolPB.GetSchoolRequest{ID: uR.User.SchoolID}); e == nil {
		schoolResp = sR
	} else {
		sR := new(schoolPB.SchoolResponse)
		schoolResp = sR
	}

	return SuccessResponse(resp, map[string]interface{}{
		"Token":                loginResp.Token,
		"TokenTime":            loginResp.ExpiredAt,
		"IsMaster":             false,
		"OrganizationUserInfo": dto.OrganizationUser{},
		"User": dto.User{
			Username:  uR.User.Username,
			CreatedAt: utils.ISOTime2MicroUnix(uR.User.CreatedAt),
			UpdatedAt: utils.ISOTime2MicroUnix(uR.User.UpdatedAt),
			IsTeacher: false,
			IsAdmin:   false,
			Mobile:    uR.User.Profile.Mobile,
			AreaCode:  "+86",
			Profile: &dto.Profile{
				ID:        uR.User.Profile.ID,
				CreatedAt: utils.ISOTime2MicroUnix(uR.User.Profile.CreatedAt),
				UpdatedAt: utils.ISOTime2MicroUnix(uR.User.Profile.UpdatedAt),
				UserNO:    "-",
				Avatar:    uR.User.Profile.Avatar,
				FirstName: uR.User.Profile.Firstname,
				LastName:  uR.User.Profile.Lastname,
				Gender:    uR.User.Profile.Gender,
			},
		},
		"School": dto.School{
			Name:       schoolResp.School.Name,
			SlugName:   schoolResp.School.SlugName,
			SchoolCode: schoolResp.School.SchoolCode,
			ID:         schoolResp.School.ID,
			// CreatedAt: utils.ISOTime2MicroUnix(schoolResp.School.)
		},
	})
}

func (a *AuthHandler) Register(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {
	authService, ok := client.SecruityAuthServiceFromContext(ctx)

	if !ok {
		return a.Error(ctx).InternalServerError("Not found AuthService")
	}

	params := struct {
		Mobile          string `json:"mobile,omitempty" validate:"nonzero"`
		Code            int64  `json:"code,omitempty" validate:"nonzero"`
		Password        string `json:"password,omitempty" validate:"nonzero"`
		ConfirmPassword string `json:"confirmPassword,omitempty" validate:"nonzero"`
		FirstName       string `json:"firstName,omitempty" validate:"nonzero"`
		LastName        string `json:"lastName,omitempty" validate:"nonzero"`
		IsTeacher       bool   `json:"isTeacher,omitempty"`
		School          string `json:"school,omitempty" validate:"objectid,nonzero"`
	}{}

	err := json.Unmarshal([]byte(req.Body), &params)

	if err != nil {
		return ErrorResponse(resp, a.Error(ctx).BadRequest(err.Error()))
	}

	if err := a.Validate(&params); err != nil {
		return ErrorResponse(resp, a.Error(ctx).BadRequest(err.Error()))
	}

	registerResp, err := authService.SignupWithMobile(ctx, &authPB.SignupWithMobileRequest{
		Mobile:          params.Mobile,
		Password:        params.Password,
		ConfirmPassword: params.ConfirmPassword,
		FirstName:       params.FirstName,
		LastName:        params.LastName,
		School:          params.School,
	})

	if err != nil {
		return ErrorResponse(resp, err)
	}

	if !registerResp.OK {
		return ErrorResponse(resp, "注册失败")
	}

	return SuccessResponse(resp, D{})
}
