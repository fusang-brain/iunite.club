package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	go_api "github.com/micro/go-api/proto"
	smsPB "iunite.club/services/message/proto/sms"
	"iunite.club/services/navo/client"
	"iunite.club/services/navo/dto"
	clubPB "iunite.club/services/organization/proto/club"
	userPB "iunite.club/services/user/proto"
)

type UserHandler struct {
	BaseHandler
}

func (u *UserHandler) Info(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID string `query:"id" validate:"objectid"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}

	if err := u.Validate(&params); err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}

	if params.ID == "" {
		if token, err := u.GetTokenCliamsFromRequest(req); err == nil {
			params.ID = token.UserID
		} else {
			return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
		}
	}

	userSrv, ok := client.UserServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError("Not found user serivce"))
	}

	userResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{
		Id: params.ID,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}
	isMaster := false

	// TODO find organization user info
	return SuccessResponse(rsp, D{
		"info":     dto.PBToUser(userResp.User),
		"IsMaster": isMaster,
	})
	// panic("not implemented")
}

func (u *UserHandler) UpdateCurrentOrg(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		UserID         string `json:"user_id,omitempty" validate:"objectid"`
		OrganizationID string `json:"organization_id,omitempty" validate:"nonzero,objectid"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}

	if err := u.Validate(&params); err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}

	userSrv, ok := client.UserServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError("Not found user serivce"))
	}

	updateResp, err := userSrv.UpdateUser(ctx, &userPB.UpdateUserRequest{
		ID: params.UserID,
		// Profile: profile,
		User: []byte(`{"defaultClubID": "` + params.OrganizationID + `"}`),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if !updateResp.OK {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest("更新失败"))
	}

	return SuccessResponse(rsp, D{})
}

func (u *UserHandler) ForgetPassword(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {

	params := struct {
		Mobile          string `json:"mobile"`
		Code            int    `json:"code"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		return u.Error(ctx).BadRequest(err.Error())
	}

	if err := u.Validate(&params); err != nil {
		return u.Error(ctx).BadRequest(err.Error())
	}

	// platform := u.GetPlatformFromRequest(req)
	smsSrv, _ := client.SMSServiceFromContext(ctx)

	validateResp, err := smsSrv.ValidateMobileCode(ctx, &smsPB.ValidateMobileCodeRequest{
		Mobile: params.Mobile,
		Code:   strconv.Itoa(params.Code),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if !validateResp.OK {
		return ErrorResponse(rsp, u.Error(ctx).TemplateBadRequest("ErrorVerifyCode"))
	}

	userSrv, found := client.UserServiceFromContext(ctx)
	if !found {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest("Not found user service"))
	}

	_, err = userSrv.ResetPasswordByMobile(ctx, &userPB.ResetPasswordRequest{
		Password:        params.Password,
		ConfirmPassword: params.ConfirmPassword,
		Mobile:          params.Mobile,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{})
}

func (u *UserHandler) AllUser(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Page   int    `json:"page" form:"page" query:"page"`
		Limit  int    `json:"limit" form:"limit" query:"limit"`
		ClubID string `json:"club_id" query:"club_id"`
	}{}

	defaultClubID := params.ClubID

	if err := u.Bind(req, &params); err != nil {
		return u.Error(ctx).BadRequest(err.Error())
	}

	if err := u.Validate(&params); err != nil {
		return u.Error(ctx).BadRequest(err.Error())
	}

	userService, found := client.UserServiceFromContext(ctx)

	if !found {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError("Not found user service"))
	}
	if defaultClubID == "" {
		token, _ := u.GetTokenCliamsFromRequest(req)
		foundUserResp, err := userService.FindUserByID(ctx, &userPB.QueryUserRequest{Id: token.UserID})
		if err != nil {
			return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
		}
		defaultClubID = foundUserResp.User.DefaultClubID
		if defaultClubID == "" {
			return ErrorResponse(rsp, u.Error(ctx).BadRequest("ClubID can't be empty"))
		}
	}

	userListResp, err := userService.FindUsersByClubID(ctx, &userPB.FindUsersByClubIDRequest{Page: int64(params.Page), Limit: int64(params.Limit), ClubID: defaultClubID})

	if err != nil {
		return ErrorResponse(rsp, err)
	}
	list := make([]*dto.User, 0)

	for _, v := range userListResp.Users {
		list = append(list, dto.PBToUser(v))
	}
	return SuccessResponse(rsp, D{
		"CurrentPage": params.Page,
		"PageSize":    params.Limit,
		"PageTotal":   userListResp.Count,
		"Total":       userListResp.Count,
		"List":        list,
	})
}

func (u *UserHandler) GetCurrentOrganization(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	clubID := ""

	userID := u.GetUserIDFromRequest(req)

	userSrv, found := client.UserServiceFromContext(ctx)
	if !found {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError("Not found user service"))
	}

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: userID})
	if err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}
	clubID = foundUserResp.User.DefaultClubID

	if clubID == "" {
		return SuccessResponse(rsp, D{
			"CurrentOrganization": &dto.Organization{},
		})
	}

	clubService, found := client.ClubServiceFromContext(ctx)

	if !found {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError("Not found club service"))
	}

	clubDetailsResp, err := clubService.FindClubDetailsByID(ctx, &clubPB.GetClubByIDRequest{ID: clubID})
	if err != nil {
		return ErrorResponse(rsp, err)
	}

	res := dto.PBToOrganization(clubDetailsResp.Club)

	return SuccessResponse(rsp, D{
		"CurrentOrganization": res,
	})
}

func (u *UserHandler) GetAllMembers(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")

}

func (u *UserHandler) CreateMember(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) RemvoeMemberFromOrg(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UpdateMember(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) GetMemberDetails(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) RemoveOrg(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UpdateUserInfo(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID               *string `json:"id,omitempty"`
		Nickname         *string `json:"nickname,omitempty"`
		Name             *string `json:"name,omitempty"`
		Birthday         *int64  `json:"birthday,omitempty"`
		Email            *string `json:"email,omitempty"`
		Gender           *string `json:"gender,omitempty"`
		SchoolDepartment *string `json:"SchoolDepartment,omitempty"`
		SchoolClass      *string `json:"SchoolClass,omitempty"`
		Major            *string `json:"Major,omitempty"`
		AdvisorMobile    *string `json:"AdvisorMobile,omitempty"` // 辅导员手机
		AdvisorName      *string `json:"AdvisorName,omitempty"`   // 辅导员姓名
		StudentID        *string `json:"StudentID,omitempty"`     // 学号
		RoomNumber       *string `json:"RoomNumber,omitempty"`    // 寝室号
	}{}
	// var map[string]interface{}
	if err := u.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}

	if err := u.Validate(&params); err != nil {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest(err.Error()))
	}
	fieldsMap := make(map[string]interface{})

	if params.Name != nil {
		rsName := []rune(*params.Name)
		firstname := string(rsName[0:1])
		lastname := string(rsName[1:])
		fieldsMap["first_name"] = firstname
		fieldsMap["last_name"] = lastname
	}

	if params.Nickname != nil {
		fieldsMap["nickname"] = *params.Nickname
	}
	if params.Birthday != nil {
		fieldsMap["birthday"] = time.Unix(*params.Birthday/1e3, 0)
	}
	if params.Email != nil {
		fieldsMap["email"] = *params.Email
	}
	if params.Gender != nil {
		fieldsMap["gender"] = *params.Gender
	}
	if params.SchoolDepartment != nil {
		fieldsMap["SchoolDepartment"] = *params.SchoolDepartment
	}
	if params.SchoolClass != nil {
		fieldsMap["SchoolClass"] = *params.SchoolClass
	}
	if params.Major != nil {
		fieldsMap["Major"] = *params.Major
	}
	if params.RoomNumber != nil {
		fieldsMap["RoomNumber"] = *params.RoomNumber
	}
	if params.AdvisorMobile != nil {
		fieldsMap["AdvisorMobile"] = *params.AdvisorMobile
	}
	if params.AdvisorName != nil {
		fieldsMap["AdvisorName"] = *params.AdvisorName
	}
	if params.StudentID != nil {
		fieldsMap["StudentID"] = *params.StudentID
	}

	profile, err := json.Marshal(fieldsMap)

	if err != nil {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError(err.Error()))
	}

	userSrv, ok := client.UserServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError("Not found user serivce"))
	}

	updateResp, err := userSrv.UpdateUser(ctx, &userPB.UpdateUserRequest{
		ID:      *params.ID,
		Profile: profile,
	})

	if err != nil {
		return ErrorResponse(rsp, u.Error(ctx).InternalServerError(err.Error()))
	}

	if !updateResp.OK {
		return ErrorResponse(rsp, u.Error(ctx).BadRequest("更新失败"))
	}

	return SuccessResponse(rsp, D{})
}

func (u *UserHandler) FlagMemberState(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) GetHotUsers(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UploadAvatar(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) ExportList(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) DownloadExportTemplate(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UploadUserList(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}
