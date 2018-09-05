package services_test

import (
	"github.com/iron-kit/unite-services/user-srv/services"
	"github.com/iron-kit/unite-services/user-srv/services/utils"
	"testing"
)

func Test_SigninUserByMobile(t *testing.T) {
	assistant := test_utils.GetAssistant(&services.UserService{})

	userService := assistant.S(&services.UserService{}).(*services.UserService)

	user, err := userService.SigninUser(services.MobileAuthType, "18627894265", "12345679")

	if err != nil {
		t.Error("测试失败: ", err.Error())
		return
	}

	t.Log("测试通过: ", user)
}
