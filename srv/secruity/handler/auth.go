package handler

import (
	"context"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/errors"
	"iunite.club/srv/secruity/client"
	auth "iunite.club/srv/secruity/proto/auth"
	"iunite.club/srv/secruity/services"
	user "iunite.club/srv/user-srv/proto/user"
)

// Auth 认证
type Auth struct {
	Error        *assistant.ErrorManager
	TokenService *services.TokenService
}

// SignupWithMobile is a function to signup a new account by mobile
func (a *Auth) SignupWithMobile(ctx context.Context, req *auth.SignupWithMobileRequest, rsp *auth.SignupResponse) error {
	log.Log("Received Auth.Signup request")

	userService, ok := client.UserServiceFromContext(ctx)
	if !ok {
		return errors.InternalServerError("iron.kit.srv.secruity", "user client not found")
	}

	var err error
	user, err := userService.RegisterUserByMobile(ctx, &user.RegisterUserRequest{
		Mobile:          req.Mobile,
		SchoolID:        req.School,
		Password:        req.Password,
		Firstname:       req.FirstName,
		Lastname:        req.LastName,
		ConfirmPassword: req.ConfirmPassword,
		Code:            req.ValidateCode,
	})

	if err != nil {
		return err
	}

	token, err := a.TokenService.Encode(user.User, 7)
	log.Log("token: ", token)
	rsp.Ok = true

	return nil
}

// Signin 登入
func (a *Auth) Signin(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	log.Log("Received Auth.Signin request")

	userService, ok := client.UserServiceFromContext(ctx)

	if !ok {
		return errors.InternalServerError("iron.kit.srv.secruity", "user client not found")
	}
	var err error
	// userInfo := user.UserResponse{}
	userInfo, err := userService.SigninByMobile(ctx, &user.SigninByMobileRequest{
		Mobile:   req.Identify,
		Password: req.Password,
	})

	rsp.Token, err = a.TokenService.Encode(userInfo.User, 7)
	if err != nil {
		return err
	}
	return nil
}
