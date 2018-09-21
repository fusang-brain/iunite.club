package handler

import (
	"context"
	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
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

	user, err := userService.RegisterUserByMobile(&user.RegisterUserRequest{
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
	tokenService := &utils.TokenService{}

	token, expiredAt, err := tokenService.Encode(user.ToPB(), 7)
	log.Log("token: ", token)
	rsp.OK = true
	rsp.Token = token
	rsp.TokenExpiredAt = expiredAt
	return nil
}

// Signin 登入
func (a *Secruity) Signin(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	log.Log("Received Auth.Signin request")

	userService := newUserService(ctx)

	user, err := userService.SigninUser(MobileAuthType, req.Identify, req.Password)

	// userInfo, err := userService.SigninByMobile(ctx, &user.SigninByMobileRequest{
	// 	Mobile:   req.Identify,
	// 	Password: req.Password,
	// })

	if err != nil {
		return err
	}

	tokenService := &utils.TokenService{}
	token, expiredAt, err := tokenService.Encode(user.ToPB(), 7)
	// fmt.Println(userInfo.User.ID, "TON")
	if err != nil {
		return a.Error(ctx).InternalServerError(err.Error())
	}

	rsp.Token = token
	rsp.ExpiredAt = expiredAt
	return nil
}
