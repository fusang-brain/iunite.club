package services

import (
	"fmt"
	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	"iunite.club/srv/user-srv/proto/user"
	"iunite.club/srv/user-srv/utils"
)

const (
	MobileAuthType = "MobileAuthType"
)

// UserService 用户服务
type UserService struct {
	assistant.Service

	Error *assistant.ErrorManager
}

func (u *UserService) SigninUser(authType string, key string, password string) (*models.User, error) {
	UserModel := u.Connection.M("User")
	ProfileModel := u.Connection.M("Profile")
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
			return &user, u.Error.ActionError("NotFoundUser")
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
			return &user, u.Error.ActionError("UserHasBan")
		}
		fmt.Println(password)
		fmt.Println(secruityInfo.Secret)
		if err := utils.CheckPassword(password, secruityInfo.Secret); err != nil {
			return &user, u.Error.ActionError("ErrorPassword")
		}

		return &user, nil
	}

	return &user, u.Error.ActionError("ErrorAuthType")
}

func (u *UserService) RegisterUserByMobile(user *kit_iron_srv_user.RegisterUserRequest) (*models.User, error) {
	// fmt.Println("errmanager: ", u.Error)
	// TODO generate username
	password, err := utils.GeneratePassword(user.Password)

	if err != nil {
		return nil, u.Error.InternalServerError(err.Error())
	}

	if user.ConfirmPassword != user.Password {
		return nil, u.Error.ActionError("ConfirmPassword")
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
	}

	return newUser, u.CreateUser(newUser)
}

func (u *UserService) ResetPasswordByMobile(req *kit_iron_srv_user.ResetPasswordRequest) (bool, error) {
	// profile := u.
	if req.Password != req.ConfirmPassword {
		return false, u.Error.ActionError("ConfirmPassword")
	}

	// TODO validate mobile code

	profile, err := u.FindProfileByMobile(req.Mobile)

	if err != nil {
		log.Log(err.Error())
		return false, u.Error.ActionError("NotFoundUserProfile")
	}

	User := u.Connection.M("User")

	password, err := utils.GeneratePassword(req.Password)

	if err != nil {
		log.Log(err.Error())
		return false, u.Error.InternalServerError(err.Error())
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
		return false, u.Error.InternalServerError(err.Error())
	}

	return true, nil
	// foundUser, err
}

func (u *UserService) FindProfileByMobile(mobile string) (*models.Profile, error) {
	profile := models.Profile{}
	Profile := u.Connection.M("Profile")

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
func (u *UserService) UpdateUser(user models.User) error {
	UserModel := u.Model("User")

	// d, _ := assistant.MongerDocumentToMap(&user)
	// d, _ := json.Marshal(&user)

	return UserModel.Update(bson.M{"_id": user.ID}, bson.M{"$set": &user})
}

// GetProfileByID 通过ID获取简历
func (u *UserService) GetProfileByID(id string) *models.Profile {
	profile := &models.Profile{}

	UserModel := u.Model("Profile")
	UserModel.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}).Exec(profile)

	return profile
}

// IsUserEnabled 用户是否启用了
func (u *UserService) IsUserEnabled(id string) (bool, error) {
	userModel := u.Model("User")

	count, err := userModel.Where(bson.M{"_id": bson.ObjectIdHex(id), "enabled": true}).Count()

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
