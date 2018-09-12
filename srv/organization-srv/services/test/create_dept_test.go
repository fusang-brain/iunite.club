package services_test

import (
	"iunite.club/srv/organization-srv/services"
	"iunite.club/srv/organization-srv/services/utils"
	"testing"
)

func Test_CreateDepartment(t *testing.T) {
	ass := utils.GetAssistant(&services.DepartmentService{})

	departmentService := ass.S(&services.DepartmentService{}).(*services.DepartmentService)

	department, err := departmentService.CreateDepartment(&services.CreateDepartmentBundle{
		Name:        "主席团",
		Description: "主席团描述",
		ParentID:    "5b90fa862af7f60001721e66",
	})

	if err != nil {
		t.Error("测试失败", err.Error())
		return
	}

	t.Logf("测试通过 %s", department.Name)
}
