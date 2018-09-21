package handler

import (
	"context"
	"encoding/json"
	"github.com/iron-kit/go-ironic"
	go_api "github.com/micro/go-api/proto"
	"iunite.club/services/navo/client"
	authPB "iunite.club/services/user/proto/secruity"
	// authPB "iunite.club/srv/secruity/proto/auth"
)

type AuthHandler struct {
	ironic.BaseHandler
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
		return a.Error(ctx).InternalServerError("Params parsed error : " + err.Error())
	}

	loginResp, err := authService.Signin(ctx, &authPB.AuthRequest{
		Identify: params.Mobile,
		Password: params.Password,
	})

	if err != nil {
		return err
	}

	resp.Body = APISuccess(map[string]interface{}{
		"Token":     loginResp.Token,
		"TokenTime": loginResp.ExpiredAt,
	}).String()

	resp.StatusCode = 200

	return nil
}

func (a *AuthHandler) Register(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {
	authService, ok := client.SecruityAuthServiceFromContext(ctx)

	if !ok {
		return a.Error(ctx).InternalServerError("Not found AuthService")
	}

	params := struct {
		Mobile          string `json:"mobile,omitempty"`
		Code            int64  `json:"code,omitempty"`
		Password        string `json:"password,omitempty"`
		ConfirmPassword string `json:"confirmPassword,omitempty"`
		FirstName       string `json:"firstName,omitempty"`
		LastName        string `json:"lastName,omitempty"`
		IsTeacher       bool   `json:"isTeacher,omitempty"`
		School          string `json:"school,omitempty"`
	}{}

	err := json.Unmarshal([]byte(req.Body), &params)

	if err != nil {
		return a.Error(ctx).InternalServerError(err.Error())
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
		return a.Error(ctx).InternalServerError(err.Error())
	}

	if !registerResp.OK {
		resp.Body = APIError("注册失败").String()
	}

	return nil
}
