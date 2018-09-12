package services_test

import (
	"iunite.club/srv/organization-srv/services"
	"iunite.club/srv/organization-srv/services/utils"
	"testing"
)

func Test_UpdateDepartment(t *testing.T) {
	ass := utils.GetAssistant(&services.DepartmentService{})

	departmentService := ass.S(&services.DepartmentService{}).(*services.DepartmentService)

	err := departmentService.UpdateDepartment(&services.UpdateDepartmentBundle{
		ID:   "5b92325416a44b8d1f7b935a",
		Name: "主席团1",
	})

	if err != nil {
		t.Error("测试失败", err.Error())
		return
	}

	t.Logf("测试通过")
}
