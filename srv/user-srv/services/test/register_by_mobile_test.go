package services_test

import (
	"github.com/iron-kit/unite-services/user-srv/proto/user"
	"github.com/iron-kit/unite-services/user-srv/services"
	"github.com/iron-kit/unite-services/user-srv/services/utils"
	"testing"
)

func Test_RegisterByMobile(t *testing.T) {
	assistant := test_utils.GetAssistant(&services.UserService{})
	userService := assistant.S(&services.UserService{}).(*services.UserService)

	user := &kit_iron_srv_user.RegisterUserRequest{
		Mobile:          "18627894265",
		Code:            "12345",
		Firstname:       "周",
		Lastname:        "金顺",
		Password:        "12345678",
		ConfirmPassword: "12345678",
	}
	newUser, err := userService.RegisterUserByMobile(user)
	if err != nil {
		t.Error("测试失败", err.Error())
	}

	if newUser.Profile.Mobile == user.Mobile {
		t.Log("测试通过", newUser)
	}
}
