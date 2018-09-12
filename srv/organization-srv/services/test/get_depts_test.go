package services_test

import (
	"iunite.club/srv/organization-srv/services"
	"iunite.club/srv/organization-srv/services/utils"
	"testing"
)

func Test_GetDepartments(t *testing.T) {
	ass := utils.GetAssistant(&services.DepartmentService{})

	departmentService := ass.S(&services.DepartmentService{}).(*services.DepartmentService)

	res, err := departmentService.GetDepartmentListByParentID(&services.GetDepartmentListBundle{
		ParentID: "5b90fa862af7f60001721e66",
	})

	if err != nil {
		t.Error("测试失败", err.Error())
		return
	}

	if res.Total > 0 && len(res.Departments) > 0 {
		t.Log("测试通过")
		return
	}

	t.Error("测试失败")
}
