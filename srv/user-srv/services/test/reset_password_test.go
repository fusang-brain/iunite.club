package services_test

import (
	"github.com/iron-kit/unite-services/user-srv/proto/user"
	"github.com/iron-kit/unite-services/user-srv/services"
	"github.com/iron-kit/unite-services/user-srv/services/utils"
	"testing"
)

func Test_ResetPassword(t *testing.T) {
	assistant := test_utils.GetAssistant(&services.UserService{})

	userService := assistant.S(&services.UserService{}).(*services.UserService)

	isSuccess, err := userService.ResetPasswordByMobile(&kit_iron_srv_user.ResetPasswordRequest{
		Mobile:          "18627894265",
		MobileAreaCode:  "86",
		Password:        "12345679",
		ConfirmPassword: "12345679",
		Code:            "123456",
	})

	if err != nil {
		t.Error("测试失败，出现错误！>", err.Error())
		return
	}

	if isSuccess {
		t.Log("测试成功")
		return
	}

	t.Error("测试失败")

}
