package handler

import (
	"context"
	"fmt"
	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	smsPB "iunite.club/services/message/proto/sms"
	"iunite.club/services/user/client"
	user "iunite.club/services/user/proto"
	auth "iunite.club/services/user/proto/secruity"
	"iunite.club/services/user/utils"
)

type Secruity struct {
	ironic.BaseHandler
}

// SignupWithMobile is a function to signup a new account by mobile
func (a *Secruity) SignupWithMobile(ctx context.Context, req *auth.SignupWithMobileRequest, rsp *auth.SignupResponse) error {
	log.Log("Received Auth.Signup request")
	userService := newUserService(ctx)
	// validate code
	smsSrv, _ := client.SMSServerFromContext(ctx)
	if resp, err := smsSrv.ValidateMobileCode(ctx, &smsPB.ValidateMobileCodeRequest{
		Mobile: req.Mobile,
		Code:   req.ValidateCode,
	}); err != nil {
		return err
	} else if !resp.OK {
		return a.Error(ctx).BadRequest("Code is error")
	}

	user, e := userService.RegisterUserByMobile(&user.RegisterUserRequest{
		Mobile:          req.Mobile,
		SchoolID:        req.School,
		Password:        req.Password,
		Firstname:       req.FirstName,
		Lastname:        req.LastName,
		ConfirmPassword: req.ConfirmPassword,
		Code:            req.ValidateCode,
	})

	if e != nil {
		fmt.Println(e.Error())
		return e
	}

	tokenService := &utils.TokenService{}

	token, expiredAt, _ := tokenService.Encode(user.ToPB(), 7)
	log.Log("token: ", token)
	rsp.OK = true
	rsp.Token = token
	rsp.TokenExpiredAt = expiredAt
	return nil
}

// Signin 登入
func (a *Secruity) Signin(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	log.Log("Received Auth.Signin request")
	fmt.Println("Login Request")
	userService := newUserService(ctx)

	user, err := userService.SigninUser(MobileAuthType, req.Identify, req.Password)

	if err != nil {
		return err
	}

	tokenService := &utils.TokenService{}
	token, expiredAt, err := tokenService.Encode(user.ToPB(), 7)
	if err != nil {
		return a.Error(ctx).InternalServerError(err.Error())
	}

	rsp.Token = token
	rsp.ExpiredAt = expiredAt
	return nil
}

func (a *Secruity) GetUserIDFromToken(ctx context.Context, req *auth.TokenRequest, resp *auth.UserIDResponse) error {
	tokenService := &utils.TokenService{}

	claims, err := tokenService.Decode(req.Token)

	if err != nil {
		return a.Error(ctx).TemplateBadRequest(err.Error())
	}

	resp.UserID = claims.UserID

	return nil
}
