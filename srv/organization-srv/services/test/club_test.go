package services_test

import (
	"fmt"
	"iunite.club/srv/organization-srv/services"
	"iunite.club/srv/organization-srv/services/utils"
	"testing"
)

func Test_GetClubsByUserID(t *testing.T) {
	ass := utils.GetAssistant(&services.ClubService{})

	clubService := ass.S(&services.ClubService{}).(*services.ClubService)

	res, err := clubService.GetClubsByUserID("5b90f99e2af7f60001b23cea", &services.PagerBundle{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		t.Error("测试失败 >>>", err.Error())
		return
	}

	if len(res.Organizations) > 0 && res.Total == 1 {
		fmt.Println(res)
		t.Log("测试成功", res.Total)
		return
	}

	t.Error("测试失败 >>>", res.Total)
}

func Test_JoinClub(t *testing.T) {
	ass := utils.GetAssistant(&services.ClubService{})

	clubService := ass.S(&services.ClubService{}).(*services.ClubService)

	err := clubService.AcceptJoinOneClub(&services.AcceptJoinClubBundle{
		UserID: "5b90f99e2af7f60001b23cea",
		ClubID: "5b90fa862af7f60001721e66",
	})

	if err != nil {
		// t.Error(err.Error())
		t.Log("测试通过")
		return
	}

	t.Error("测试失败")

}

func Test_JoinClub1(t *testing.T) {
	ass := utils.GetAssistant(&services.ClubService{})

	clubService := ass.S(&services.ClubService{}).(*services.ClubService)

	err := clubService.AcceptJoinOneClub(&services.AcceptJoinClubBundle{
		UserID: "5b90f99e2af7f60001b23cea",
		ClubID: "5b90f99e2af7f60001b23cea",
	})

	if err != nil {
		t.Error(err.Error())
		// t.Log("测试通过")
		return
	}

	t.Log("测试通过")

}
