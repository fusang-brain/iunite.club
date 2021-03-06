package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	smsPB "iunite.club/services/message/proto/sms"

	"iunite.club/services/user/client"
	user "iunite.club/services/user/proto"
	"iunite.club/services/user/utils"
)

type UserSrv struct {
	ironic.BaseHandler
}

func (u *UserSrv) model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (u *UserSrv) UpdateAvatar(ctx context.Context, req *user.UpdateAvatarRequest, rsp *user.Response) error {
	ProfileModel := u.model(ctx, "Profile")

	err := ProfileModel.Update(bson.M{
		"user_id": req.ID,
	}, bson.M{
		"$set": bson.M{
			"avatar": req.Avatar,
		},
	})

	if err != nil {
		return u.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (u *UserSrv) FindUserClubProfileByID(ctx context.Context, req *user.FindUserClubProfileByIDRequest, rsp *user.UserClubProfileResponse) error {
	userService := newUserService(ctx)
	clubProfile := userService.FindUserClubProfileByID(req.ID)

	if clubProfile.IsEmpty() {
		return u.Error(ctx).NotFound("NotFoundUserClubProfle By %s", req.ID)
	}
	rsp.ClubProfile = &user.UserClubProfile{
		ID:             clubProfile.ID.Hex(),
		UserID:         clubProfile.UserID.Hex(),
		OrganizationID: clubProfile.OrganizationID.Hex(),
		State:          int32(clubProfile.State),
		IsCreator:      clubProfile.IsCreator,
		IsMaster:       clubProfile.IsMaster,
		JoinTime:       hptypes.TimestampProto(clubProfile.JoinTime),
		LeaveTime:      hptypes.TimestampProto(clubProfile.LeaveTime),
		CreatedAt:      hptypes.TimestampProto(clubProfile.CreatedAt),
		UpdatedAt:      hptypes.TimestampProto(clubProfile.UpdatedAt),
		JobID:          clubProfile.JobID.Hex(),
		DepartmentID:   clubProfile.DepartmentID.Hex(),
	}
	return nil
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
	if req.Username == "" {
		req.Username = userService.GenerateUsername()
	}

	password, err := utils.GeneratePassword(req.Password)
	if err != nil {
		return u.Error(ctx).InternalServerError(err.Error())
	}

	newUser := models.User{
		Username: req.Username,
		Enabled:  req.Enabled,
		SecruityInfos: []models.SecruityInfo{
			{
				AuthType:      "UniteApp",
				Secret:        password,
				PlainPassword: req.Password,
			},
		},
		Profile: &models.Profile{
			Avatar:    req.Profile.Avatar,
			Lastname:  req.Profile.Lastname,
			Firstname: req.Profile.Firstname,
			Mobile:    req.Profile.Mobile,
			Birthday:  hptypes.Timestamp(req.Profile.Birthday),
			Nickname:  req.Profile.Nickname,
			Gender:    req.Profile.Gender,
			// RoomNumber: req.Profile.
		},
	}

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

		if profileFields["birthday"] != nil {
			// profileFields["birthday"] = time.Unix()
			profileFields["birthday"] = func(v string) time.Time {
				fmt.Println(v, "生日 =====")
				now := time.Now()
				// time.RFC1123
				// now.String(2006-01-02 15:04:05.999)
				// return time.Unix(v/1e3, 0)
				// now.String()
				// time.ParseInLocation
				if t, err := time.Parse(time.RFC1123Z, v); err == nil {
					fmt.Println(t.String(), "生日 11======")
					return t
				}
				return now
			}(profileFields["birthday"].(string))
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

func (u *UserSrv) CreateMember(ctx context.Context, req *user.CreateMemberRequest, rsp *user.Response) error {

	userService := newUserService(ctx)
	if req.User.Profile == nil {
		req.User.Profile = new(user.Profile)
	}

	profile := req.User.Profile
	password := "123456"
	pwd, _ := utils.GeneratePassword(password)

	if !bson.IsObjectIdHex(req.ClubID) || !bson.IsObjectIdHex(req.DepartmentID) || !bson.IsObjectIdHex(req.JobID) {
		return u.Error(ctx).BadRequest("ClubID/DepartmentID/JobID) must be objectid")
	}

	err := userService.CreateMember(&CreateMemberBundle{
		ClubID:       bson.ObjectIdHex(req.ClubID),
		DepartmentID: bson.ObjectIdHex(req.DepartmentID),
		JobID:        bson.ObjectIdHex(req.JobID),
		User: &models.User{
			Username: userService.GenerateUsername(),
			Enabled:  false,
			SecruityInfos: []models.SecruityInfo{
				{
					AuthType:      "UniteApp",
					Secret:        pwd,
					PlainPassword: password,
				},
			},
			Profile: &models.Profile{
				Avatar:           profile.Avatar,
				Mobile:           profile.Mobile,
				Email:            profile.Email,
				Firstname:        profile.Firstname,
				Lastname:         profile.Lastname,
				Gender:           profile.Gender,
				Birthday:         hptypes.Timestamp(profile.Birthday),
				Nickname:         profile.Nickname,
				SchoolClass:      profile.SchoolClass,
				SchoolDepartment: profile.SchoolDepartment,
				Major:            profile.Major,
				AdvisorMobile:    profile.AdvisorMobile,
				AdvisorName:      profile.AdvisorName,
				StudentID:        profile.StudentID,
				RoomNumber:       profile.RoomNumber,
			},
		},
	})

	if err != nil {
		return err
	}

	rsp.OK = true

	return nil
}

func (u *UserSrv) FindUsersByOrganizationID(ctx context.Context, req *user.ByOrganizationIDRequest, rsp *user.UserListResponse) error {
	UserClubProfile := u.model(ctx, "UserClubProfile")

	ucps := make([]models.UserClubProfile, 0)
	users := make([]*user.User, 0)
	if !bson.IsObjectIdHex(req.ID) {
		return u.Error(ctx).BadRequest("ID must be objectid")
	}

	key := "department_id"

	if req.Kind == "club" {
		key = "organization_id"
	}

	err := UserClubProfile.Where(bson.M{key: bson.ObjectIdHex(req.ID)}).
		Populate("User", "User.Profile").
		FindAll(&ucps)

	if err != nil {
		return u.Error(ctx).InternalServerError(err.Error())
	}

	for _, ucp := range ucps {
		user := ucp.User

		users = append(users, user.ToPB())
	}

	rsp.Users = users
	rsp.Count = int32(len(users))
	return nil

}

// func (u *UserSrv) RemoveUserFromClub(ctx context.Context, req *user.Remove, rsp *user.Response) error {
// 	userSrv := newUserService(ctx)

// 	if !bson.IsObjectIdHex(req.UserID) {
// 		return u.Error(ctx).BadRequest("user id must be objectid")
// 	}

// 	if !bson.IsObjectIdHex(req.ClubID) {
// 		return u.Error(ctx).BadRequest("club id must be objectid")
// 	}

// 	err := userSrv.RemoveFromClub(bson.ObjectIdHex(req.UserID), bson.ObjectIdHex(req.ClubID))

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
