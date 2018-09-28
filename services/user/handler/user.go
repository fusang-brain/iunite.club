package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	smsPB "iunite.club/services/message/proto/sms"
	"iunite.club/services/user/client"
	user "iunite.club/services/user/proto"
)

type UserSrv struct {
	ironic.BaseHandler
}

func (u *UserSrv) FindUserByID(ctx context.Context, req *user.QueryUserRequest, resp *user.UserResponse) error {
	// log.Log("start load user info")
	// fmt.Println("Helo")
	// fmt.Println(req)
	userService := newUserService(ctx)

	userInfo := userService.GetUserInfoByID(req.Id)

	if userInfo.IsEmpty() {
		return u.Error(ctx).NotFound("NotFoundUser")
	}
	// ptypes.
	resp.User = userInfo.ToPB()
	return nil
}

func (u *UserSrv) FindProfileByID(ctx context.Context, in *user.QueryProfileRequest, out *user.ProfileResponse) error {

	userService := newUserService(ctx)
	profile := userService.GetProfileByID(in.Id)
	out.Profile = profile.ToPB()
	return nil
}

func (u *UserSrv) CreateUser(ctx context.Context, req *user.User, resp *user.Response) error {
	userService := newUserService(ctx)
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

func (u *UserSrv) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, resp *user.Response) error {

	//
	if !bson.IsObjectIdHex(req.ID) {
		return u.Error(ctx).BadRequest("ID must be a objectid")
	}

	userFields := make(map[string]interface{})
	log.Logf("will updated user fields %v \r\n", string(req.User))
	profileFields := make(map[string]interface{})
	log.Logf("will updated profile fields %v \r\n", string(req.Profile))
	userService := newUserService(ctx)
	if len(req.User) > 0 {
		if err := json.Unmarshal(req.User, &userFields); err != nil {
			return u.Error(ctx).InternalServerError(err.Error())
		}
		if err := userService.UpdateUser(bson.ObjectIdHex(req.ID), userFields); err != nil {
			return u.Error(ctx).InternalServerError(err.Error())
		}
	}
	if len(req.Profile) > 0 {
		if err := json.Unmarshal(req.Profile, &profileFields); err != nil {
			return u.Error(ctx).InternalServerError(err.Error())
		}
		if err := userService.UpdateProfileByUserID(bson.ObjectIdHex(req.ID), profileFields); err != nil {
			return u.Error(ctx).InternalServerError(err.Error())
		}
	}

	resp.OK = true
	return nil
}

// func (u *User) FindUsers(ctx context.Context, req *user.PagerRequest, resp *user.UserListResponse) error
// 	IsUserEnabled(context.Context, *OnlyIDRequest, *LogicResponse) error

func (u *UserSrv) FindUsers(ctx context.Context, req *user.PagerRequest, resp *user.UserListResponse) error {
	log.Log("to find users")

	return nil
}

func (u *UserSrv) IsUserEnabled(ctx context.Context, req *user.QueryUserRequest, resp *user.Response) error {
	fmt.Println("check is user enabled", req)
	log.Log("check is user enabled")
	userService := newUserService(ctx)
	// ok := false
	isSuccess := userService.IsUserEnabled(req.Id)
	fmt.Println("checked user :", isSuccess)
	// if err != nil {
	// 	return err
	// }

	resp.OK = isSuccess
	return nil
}

func (u *UserSrv) RegisterUserByMobile(ctx context.Context, req *user.RegisterUserRequest, resp *user.RegisterUserResponse) error {
	log.Log("receive register user request")
	userService := newUserService(ctx)
	if smsService, ok := client.SMSServerFromContext(ctx); ok {
		resp, err := smsService.ValidateMobileCode(ctx, &smsPB.ValidateMobileCodeRequest{
			Mobile: req.Mobile,
			Code:   req.Code,
		})

		if err != nil {
			return err
		}

		if !resp.OK {
			return u.Error(ctx).BadRequest("Code is error")
		}
	} else {
		return u.Error(ctx).BadRequest("SMS Service is not enable")
	}

	newUser, err := userService.RegisterUserByMobile(req)
	if err != nil {
		return err
	}

	resp.OK = true
	resp.User = newUser.ToPB()
	return nil
}

func (u *UserSrv) ResetPasswordByMobile(ctx context.Context, req *user.ResetPasswordRequest, resp *user.ResetPasswordResponse) error {
	log.Log("receive reset password by mobile request")
	userService := newUserService(ctx)
	_, err := userService.ResetPasswordByMobile(req, resp)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserSrv) SigninByMobile(ctx context.Context, req *user.SigninByMobileRequest, resp *user.UserResponse) error {
	log.Log("receive signin user request")
	fmt.Println("receive signin user request")
	userService := newUserService(ctx)

	user, err := userService.SigninUser(MobileAuthType, req.Mobile, req.Password)

	if err != nil {
		return err
	}

	resp.User = user.ToPB()
	return nil
}

func (u *UserSrv) FindUsersByClubID(ctx context.Context, req *user.FindUsersByClubIDRequest, rsp *user.UserListResponse) error {
	userService := newUserService(ctx)

	err := userService.FindUsersByClubID(req, rsp)
	if err != nil {
		return err
	}

	return nil
}
