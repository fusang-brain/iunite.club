package services_test

import (
	"github.com/iron-kit/go-ironic/bundles"
	"iunite.club/srv/organization-srv/services"
	"iunite.club/srv/organization-srv/services/utils"
	"testing"
)

func Test_CreateJob(t *testing.T) {
	ass := utils.GetAssistant(&services.JobService{})

	jobService := ass.S(&services.JobService{}).(*services.JobService)

	org, err := jobService.CreateJob(&services.CreateJobBundle{
		Name:   "主席",
		ClubID: "5b90fa862af7f60001721e66",
	})

	if err != nil {
		t.Error(err.Error())
		return
	}

	if org.OrganizationID.Hex() == "5b90fa862af7f60001721e66" {
		t.Log("测试成功")
		return
	}

	t.Error("测试失败", org.OrganizationID)

}

func Test_ListJob(t *testing.T) {
	ass := utils.GetAssistant(&services.JobService{})

	jobService := ass.S(&services.JobService{}).(*services.JobService)

	list, err := jobService.GetJobListByParentID(&services.JobListRequestBundle{
		PaginationBundle: bundles.PaginationBundle{
			Page:  1,
			Limit: 10,
		},
		OrganizationID: "5b90fa862af7f60001721e66",
	})

	if err != nil {
		t.Error(err.Error())
		return
	}

	if list.Total > 0 && len(list.Jobs) > 0 {
		t.Log("Passed")
		return
	}

	t.Error("Not Passed")
}

// func Test_UpdateJob(t *testing.T) {
// 	ass := utils.GetAssistant(&services.JobService{})
// 	jobService := ass.S(&services.JobService{}).(*services.JobService)

// 	org, err := jobService.UpdateJob(&services.UpdateJobBundle{
// 		ID: ""
// 		Name: "主席1",

// 	})
// }
