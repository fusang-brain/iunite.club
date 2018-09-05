package handler

import (
	"context"
	"fmt"
	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"iunite.club/models"
	user "iunite.club/srv/user-srv/proto/user"
	"iunite.club/srv/user-srv/services"
)

type User struct {
	assistant.BaseHandler

	UserService *services.UserService
	Error       *assistant.ErrorManager
}

func (u *User) FindUserByID(ctx context.Context, req *user.QueryUserRequest, resp *user.UserResponse) error {
	log.Log("start load user info")
	// fmt.Println("Helo")
	// fmt.Println(req)
	userService := u.UserService

	userInfo := userService.GetUserInfoByID(req.Id)
	// ptypes.
	resp.User = userInfo.ToPB()
	return nil
}

func (u *User) FindProfileByID(ctx context.Context, in *user.QueryProfileRequest, out *user.ProfileResponse) error {
	log.Log("Find Profile")
	userService := u.UserService
	profile := userService.GetProfileByID(in.Id)
	out.Profile = profile.ToPB()
	return nil
}

func (u *User) CreateUser(ctx context.Context, req *user.User, resp *user.Response) error {
	userService := u.UserService
	newUser := models.User{
		Username: req.Username,
		Enabled:  req.Enabled,
		SecruityInfos: []models.SecruityInfo{
			models.SecruityInfo{
				AuthType:      "UniteApp",
				Secret:        req.Password,
				PlainPassword: req.Password,
			},
		},
		Profile: &models.Profile{
			Avatar: req.Profile.Avatar,
		},
	}

	// if err := assistant.TransformDTOToMongerSchema(req, &newUser); err != nil {
	// 	return err
	// }
	if err := userService.CreateUser(&newUser); err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (u *User) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, resp *user.Response) error {

	log.Log("receive UpdateUser request")
	// willUpdateUser := models.User{}

	// assistant.TransformDTOToMongerSchema(req.UserInfo, &willUpdateUser)

	// if err := u.UserService.UpdateUser(willUpdateUser); err != nil {
	// 	return err
	// }

	resp.OK = true
	return nil
}

// func (u *User) FindUsers(ctx context.Context, req *user.PagerRequest, resp *user.UserListResponse) error
// 	IsUserEnabled(context.Context, *OnlyIDRequest, *LogicResponse) error

func (u *User) FindUsers(ctx context.Context, req *user.PagerRequest, resp *user.UserListResponse) error {
	log.Log("to find users")

	return nil
}

func (u *User) IsUserEnabled(ctx context.Context, req *user.QueryUserRequest, resp *user.Response) error {
	log.Log("check is user enabled")
	// ok := false
	isSuccess, err := u.UserService.IsUserEnabled(req.Id)
	if err != nil {
		return err
	}

	resp.OK = isSuccess
	return nil
}

func (u *User) RegisterUserByMobile(ctx context.Context, req *user.RegisterUserRequest, resp *user.RegisterUserResponse) error {
	log.Log("receive register user request")
	fmt.Println(req)
	newUser, err := u.UserService.RegisterUserByMobile(req)
	if err != nil {
		return err
	}

	resp.OK = true
	resp.User = newUser.ToPB()
	return nil
}

func (u *User) ResetPasswordByMobile(ctx context.Context, req *user.ResetPasswordRequest, resp *user.Response) error {
	log.Log("receive reset password by mobile request")

	isSuccess, err := u.UserService.ResetPasswordByMobile(req)

	if err != nil {
		return err
	}

	resp.OK = isSuccess

	return nil
}

func (u *User) SigninByMobile(ctx context.Context, req *user.SigninByMobileRequest, resp *user.UserResponse) error {
	log.Log("receive signin user request")

	user, err := u.UserService.SigninUser(services.MobileAuthType, req.Mobile, req.Password)

	if err != nil {
		return err
	}

	resp.User = user.ToPB()
	return nil
}
