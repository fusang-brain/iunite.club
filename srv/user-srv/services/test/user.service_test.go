package services_test

import (
	"github.com/iron-kit/unite-services/user-srv/models"
	"github.com/iron-kit/unite-services/user-srv/services"
	"github.com/iron-kit/unite-services/user-srv/services/utils"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// func Test_CreateUser(t *testing.T) {
// 	assistant := getAssistant()
// 	srv := assistant.S(&UserService{}).(*UserService)
// 	user := models.User{
// 		Username: "100000001",
// 		Enabled:  false,
// 		SecruityInfos: []models.SecruityInfo{
// 			{
// 				AuthType:      "UniteApp",
// 				Key:           "-",
// 				Secret:        "1234567",
// 				PlainPassword: "1234567",
// 			},
// 		},
// 		Profile: &models.Profile{
// 			Avatar: "https://www.baidu.com",
// 			Mobile: "18627894265",
// 		},
// 	}

// 	err := srv.CreateUser(&user)
// 	if err != nil {
// 		t.Error("创建用户测试失败", err.Error())
// 	} else {
// 		t.Log("测试通过")
// 	}
// }

func Test_UpdateUser(t *testing.T) {
	assistant := test_utils.GetAssistant(&services.UserService{})
	srv := assistant.S(&services.UserService{}).(*services.UserService)
	user := models.User{
		Username: "7654321",
		Enabled:  true,
	}

	user.ID = bson.ObjectIdHex("5b89096716a44b6954c4aad4")
	err := srv.UpdateUser(user)

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("修改用户测试通过了")
	}
}

func Test_IsUserEnabled_1(t *testing.T) {
	assistant := test_utils.GetAssistant(&services.UserService{})

	srv := assistant.S(&services.UserService{}).(*services.UserService)

	isEnabled, err := srv.IsUserEnabled("5b89096716a44b6954c4aad4")

	if err != nil {
		t.Error("IsUserEnabled_1 测试失败：", err.Error())
	}

	if isEnabled {
		t.Log("IsUserEnabled_1 测试通过", isEnabled)
	} else {
		t.Error("IsUserEnabled_1 测试失败：", isEnabled)
	}
}

func Test_GetProfile(t *testing.T) {
	assistant := test_utils.GetAssistant(&services.UserService{})
	srv := assistant.S(&services.UserService{}).(*services.UserService)

	profile := srv.GetProfileByID("5b89098d16a44b698d5dd3b2")

	if profile.Mobile == "18627894264" {
		t.Log("测试通过")
	} else {
		t.Error("Test_GetProfile 测试失败", profile.ID)
	}
}
