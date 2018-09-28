package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	kit_iron_srv_user "iunite.club/services/user/proto"
	"iunite.club/services/user/utils"
)

const (
	MobileAuthType = "MobileAuthType"
)

// UserService 用户服务
type UserService struct {
	ironic.Service
}

func newUserService(ctx context.Context) *UserService {
	userService := &UserService{}

	if err := ironic.InitServiceFunc(userService, ctx); err != nil {
		panic(err.Error())
	}

	return userService
}

func (u *UserService) Model(name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(u.Ctx())

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (u *UserService) SigninUser(authType string, key string, password string) (*models.User, error) {

	UserModel := u.Model("User")
	ProfileModel := u.Model("Profile")
	user := models.User{}
	// signin by mobileasfasdfasdf
	if authType == MobileAuthType {
		profile := models.Profile{}

		ProfileModel.FindOne(bson.M{"mobile": key}).Exec(&profile)
		UserModel.FindOne(bson.M{
			"_id": profile.UserID,
			"secruity_infos": bson.M{
				"$elemMatch": bson.M{
					"auth_type": "UniteApp",
				},
			},
		}).Exec(&user)

		if user.IsEmpty() {
			return &user, u.Error().ActionError("NotFoundUser")
		}

		// disabled in this version
		// if !user.Enabled {
		// 	return &user, u.Error.ActionError("NotEnabledUser")
		// }

		secruityInfo := models.SecruityInfo{}

		for _, secruity := range user.SecruityInfos {
			if secruity.AuthType == "UniteApp" {
				secruityInfo = secruity
			}
		}

		if secruityInfo.Secret == "" {
			return &user, u.Error().ActionError("UserHasBan")
		}
		fmt.Println(password)
		fmt.Println(secruityInfo.Secret)
		if err := utils.CheckPassword(password, secruityInfo.Secret); err != nil {
			return &user, u.Error().ActionError("ErrorPassword")
		}

		return &user, nil
	}

	return &user, u.Error().ActionError("ErrorAuthType")
}

func (u *UserService) RegisterUserByMobile(user *kit_iron_srv_user.RegisterUserRequest) (*models.User, error) {
	// fmt.Println("errmanager: ", u.Error)
	// TODO generate username

	SchoolModel := u.Model("School")
	password, err := utils.GeneratePassword(user.Password)

	if err != nil {
		return nil, u.Error().InternalServerError(err.Error())
	}

	if user.ConfirmPassword != user.Password {
		return nil, u.Error().ActionError("ConfirmPassword")
	}

	if !bson.IsObjectIdHex(user.SchoolID) {
		return nil, u.Error().BadRequest("SchoolID must be a objectid")
	}

	foundSchoolCount := SchoolModel.Count(bson.M{"_id": bson.ObjectIdHex(user.SchoolID)})

	if foundSchoolCount <= 0 {
		return nil, u.Error().BadRequest("School is not exists")
	}

	if profile, err := u.FindProfileByMobile(user.Mobile); err != nil {
		fmt.Println(err.Error())
		return nil, u.Error().InternalServerError(err.Error())
		// return nil, u.Error().BadRequest("Account %s has be registered", user.Mobile)
	} else if !profile.IsEmpty() {
		fmt.Println(profile)

		return nil, u.Error().BadRequest("Account has be registered")
	}

	newUser := &models.User{
		Enabled: false,
		// SchoolID: req.SchoolID,
		SecruityInfos: []models.SecruityInfo{
			{
				AuthType:      "UniteApp",
				Secret:        password,
				PlainPassword: user.Password,
			},
		},
		Profile: &models.Profile{
			Mobile:    user.Mobile,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		},
		SchoolID: bson.ObjectIdHex(user.SchoolID),
	}

	return newUser, u.CreateUser(newUser)
}

func (u *UserService) ResetPasswordByMobile(req *kit_iron_srv_user.ResetPasswordRequest, rsp *kit_iron_srv_user.ResetPasswordResponse) (bool, error) {
	// profile := u.
	if req.Password != req.ConfirmPassword {
		return false, u.Error().ActionError("ConfirmPassword")
	}

	// validate mobile code
	profile, err := u.FindProfileByMobile(req.Mobile)

	if err != nil {
		log.Log(err.Error())
		return false, u.Error().ActionError("NotFoundUserProfile")
	}

	User := u.Model("User")

	password, err := utils.GeneratePassword(req.Password)

	if err != nil {
		log.Log(err.Error())
		return false, u.Error().InternalServerError(err.Error())
	}
	fmt.Println(profile)
	if err := User.Update(bson.M{
		"_id": profile.UserID,
		"secruity_infos": bson.M{
			"$elemMatch": bson.M{
				"auth_type": "UniteApp",
			},
		},
	}, bson.M{"$set": bson.M{"secruity_infos.$.secret": password, "secruity_infos.$.plain_password": req.Password}}); err != nil {
		return false, u.Error().InternalServerError(err.Error())
	}
	rsp.UpdatedAt = time.Now().String()
	rsp.ID = profile.UserID.Hex()
	return true, nil
	// foundUser, err
}

func (u *UserService) FindProfileByMobile(mobile string) (*models.Profile, error) {
	profile := models.Profile{}
	Profile := u.Model("Profile")

	if err := Profile.FindOne(bson.M{"mobile": mobile}).Exec(&profile); err != nil {
		return &profile, err
	}

	return &profile, nil
}

// GetUserInfoByID 通过ID获取用户信息
func (u *UserService) GetUserInfoByID(id string) *models.User {
	user := &models.User{}
	UserModel := u.Model("User")

	// UserModel.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}).One(user)
	UserModel.
		FindOne(bson.M{"_id": bson.ObjectIdHex(id)}).
		Populate("Profile").
		Exec(user)
	return user
}

// CreateUser 创建用户
func (u *UserService) CreateUser(user *models.User) error {
	UserModel := u.Model("User")
	ProfileModel := u.Model("Profile")
	profile := models.Profile{}
	if user.Profile != nil {
		// profile.Avatar =
		profile = *user.Profile
	}
	if err := UserModel.Create(user); err != nil {
		return err
	}
	profile.UserID = user.ID

	if err := ProfileModel.Create(&profile); err != nil {
		return err
	}

	return nil
}

// UpdateUser  更新用户
func (u *UserService) UpdateUser(id bson.ObjectId, user map[string]interface{}) error {
	UserModel := u.Model("User")
	for k, v := range user {
		if k == "defaultClubID" {
			user[k] = bson.ObjectIdHex(v.(string))
		}
		if k == "school_id" {
			user[k] = bson.ObjectIdHex(v.(string))
		}
	}
	_, err := UserModel.Upsert(bson.M{"_id": id}, bson.M{"$set": user})

	return err
}

// UpdateProfileByID  更新用户简历
func (u *UserService) UpdateProfileByID(id bson.ObjectId, profile interface{}) error {
	UserModel := u.Model("Profile")
	_, err := UserModel.Upsert(bson.M{"_id": id}, bson.M{"$set": profile})
	return err
}

// UpdateProfileByUserID  更新用户简历
func (u *UserService) UpdateProfileByUserID(id bson.ObjectId, profile interface{}) error {
	UserModel := u.Model("Profile")
	_, err := UserModel.Upsert(bson.M{"user_id": id}, bson.M{"$set": profile})
	return err
}

// func (u )

// GetProfileByID 通过ID获取简历
func (u *UserService) GetProfileByID(id string) *models.Profile {
	profile := &models.Profile{}

	UserModel := u.Model("Profile")
	UserModel.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}).Exec(profile)

	return profile
}

// IsUserEnabled 用户是否启用了
func (u *UserService) IsUserEnabled(id string) bool {
	userModel := u.Model("User")
	count := userModel.Count(bson.M{"_id": bson.ObjectIdHex(id), "enabled": true})
	log.Logf("IsUserEnabled 找到的用户数: %d", count)

	if count > 0 {
		return true
	}

	return false
}

func (u *UserService) FindUsersByClubID(req *kit_iron_srv_user.FindUsersByClubIDRequest, rsp *kit_iron_srv_user.UserListResponse) error {
	if !bson.IsObjectIdHex(req.ClubID) {
		return u.Error().BadRequest("ID must be a objectid")
	}
	UserModel := u.Model("User")
	total := 0
	users := make([]models.User, 0)
	condition := bson.M{
		"user_club_profiles": bson.M{
			"$elemMatch": bson.M{
				"organization_id": bson.ObjectIdHex(req.ClubID),
			},
		},
	}

	// fmt.Println(condition)
	query := UserModel.Find(condition).Populate("UserClubProfiles", "Profile")

	total, _ = query.Count()

	err := query.Skip(int((req.Page - 1) * req.Limit)).Limit(int(req.Limit)).Exec(&users)
	// b, _ := json.Marshal(users)

	// fmt.Println(string(b))
	if err != nil {
		return u.Error().InternalServerError(err.Error())
	}

	rsp.Count = int32(total)
	rsp.Page = int32(req.Page)
	rsp.Limit = int32(req.Limit)
	rsp.Users = make([]*kit_iron_srv_user.User, 0)

	for _, v := range users {
		rsp.Users = append(rsp.Users, v.ToPB())
	}

	return nil
}
