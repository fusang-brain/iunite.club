package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"iunite.club/services/restful/file"
	storagePB "iunite.club/services/storage/proto"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"

	"github.com/iron-kit/go-ironic/utils"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"

	smsPB "iunite.club/services/message/proto/sms"
	orgPB "iunite.club/services/organization/proto"
	clubPB "iunite.club/services/organization/proto/club"
	schoolPB "iunite.club/services/organization/proto/school"
	"iunite.club/services/restful/dto"
	userPB "iunite.club/services/user/proto"
)

type UserHandler struct {
	BaseHandler

	userService   userPB.UserSrvService
	clubService   clubPB.ClubService
	smsService    smsPB.SMSService
	schoolService schoolPB.SchoolSrvService
	fileService   storagePB.StorageService
}

func NewUserHandler(c client.Client) *UserHandler {
	return &UserHandler{
		userService:   userPB.NewUserSrvService(UserSerivce, c),
		clubService:   clubPB.NewClubService(OrganizationService, c),
		smsService:    smsPB.NewSMSService(SMSService, c),
		schoolService: schoolPB.NewSchoolSrvService(OrganizationService, c),
		fileService:   storagePB.NewStorageService(StorageService, c),
	}
}

func (u *UserHandler) Info(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `query:"id" validate:"objectid"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if params.ID == "" {
		if token, err := u.GetTokenCliamsFromRequest(req); err == nil {
			params.ID = token.UserID
		} else {
			ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
			return
		}
	}

	userSrv := u.userService

	userResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{
		Id: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	// find organization user info
	clubProfileResp, err := u.clubService.GetUserClubProfilesByUserID(ctx, &clubPB.GetUserClubProfilesByUserIDRequest{UserID: userResp.User.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	currentUserClubProfile := new(orgPB.UserClubProfile)

	userClubProfiles := clubProfileResp.UserClubProfiles
	for _, v := range userClubProfiles {
		if v.OrganizationID == userResp.User.DefaultClubID {
			currentUserClubProfile = v
		}
	}

	isMaster := currentUserClubProfile.IsMaster
	uR := userResp
	// find school
	schoolService := u.schoolService
	var schoolResp *schoolPB.SchoolResponse
	if sR, e := schoolService.GetSchoolByID(ctx, &schoolPB.GetSchoolRequest{ID: uR.User.SchoolID}); e == nil {
		schoolResp = sR
	} else {
		sR := new(schoolPB.SchoolResponse)
		schoolResp = sR
	}

	// find organization
	clubsResp, err := u.clubService.GetClubsByUserID(ctx, &clubPB.GetClubsByUserIDRequest{UserID: uR.User.ID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	org := new(dto.Organization)
	orgs := make([]*dto.Organization, 0)
	for _, v := range clubsResp.Organizations {
		if v.ID == uR.User.DefaultClubID {
			org = dto.PBToOrganization(v)
		}
		orgs = append(orgs, dto.PBToOrganization(v))
	}

	SuccessResponse(rsp, D{
		"OrganizationUserInfo": dto.PBToOrganizationUser(currentUserClubProfile),
		"info": dto.User{
			ID:            uR.User.ID,
			Username:      uR.User.Username,
			CreatedAt:     utils.ISOTime2MicroUnix(uR.User.CreatedAt),
			UpdatedAt:     utils.ISOTime2MicroUnix(uR.User.UpdatedAt),
			IsTeacher:     false,
			IsAdmin:       false,
			Mobile:        uR.User.Profile.Mobile,
			AreaCode:      "+86",
			Organizations: orgs,

			CurrentOrganization: org,
			Profile: &dto.Profile{
				ID:        uR.User.Profile.ID,
				CreatedAt: utils.Time2MicroUnix(hptypes.Timestamp(uR.User.Profile.CreatedAt)),
				UpdatedAt: utils.Time2MicroUnix(hptypes.Timestamp(uR.User.Profile.UpdatedAt)),
				// CreatedAt: utils.ISOTime2MicroUnix(uR.User.Profile.CreatedAt),
				// UpdatedAt: utils.ISOTime2MicroUnix(uR.User.Profile.UpdatedAt),
				UserNO:    "-",
				Avatar:    uR.User.Profile.Avatar,
				FirstName: uR.User.Profile.Firstname,
				LastName:  uR.User.Profile.Lastname,
				Gender:    uR.User.Profile.Gender,
			},
			School: &dto.School{
				Name:       schoolResp.School.Name,
				SlugName:   schoolResp.School.SlugName,
				SchoolCode: schoolResp.School.SchoolCode,
				ID:         schoolResp.School.ID,
			},
		},
		"IsMaster": isMaster,
	})
	// panic("not implemented")
}

func (u *UserHandler) UpdateCurrentOrg(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		UserID         string `json:"user_id,omitempty" validate:"objectid"`
		OrganizationID string `json:"organization_id,omitempty" validate:"nonzero,objectid"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if params.UserID == "" {
		params.UserID = u.GetUserIDFromRequest(req)
	}

	userSrv := u.userService

	updateResp, err := userSrv.UpdateUser(ctx, &userPB.UpdateUserRequest{
		ID: params.UserID,
		// Profile: profile,
		User: []byte(`{"defaultClubID": "` + params.OrganizationID + `"}`),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !updateResp.OK {
		ErrorResponse(rsp, u.Error().BadRequest("更新失败"))
		return
	}

	SuccessResponse(rsp, D{})
}

func (u *UserHandler) ForgetPassword(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Mobile          string `json:"mobile"`
		Code            int    `json:"code"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		u.Error().BadRequest(err.Error())
		return
	}

	if err := u.Validate(&params); err != nil {
		u.Error().BadRequest(err.Error())
		return
	}

	// platform := u.GetPlatformFromRequest(req)
	smsSrv := u.smsService

	validateResp, err := smsSrv.ValidateMobileCode(ctx, &smsPB.ValidateMobileCodeRequest{
		Mobile: params.Mobile,
		Code:   strconv.Itoa(params.Code),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !validateResp.OK {
		ErrorResponse(rsp, u.Error().TemplateBadRequest("ErrorVerifyCode"))
		return
	}

	userSrv := u.userService

	_, err = userSrv.ResetPasswordByMobile(ctx, &userPB.ResetPasswordRequest{
		Password:        params.Password,
		ConfirmPassword: params.ConfirmPassword,
		Mobile:          params.Mobile,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

func (u *UserHandler) AllUser(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Page   int    `json:"page" form:"page" query:"page"`
		Limit  int    `json:"limit" form:"limit" query:"limit"`
		ClubID string `json:"club_id" query:"club_id"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		u.Error().BadRequest(err.Error())
		return
	}

	if err := u.Validate(&params); err != nil {
		u.Error().BadRequest(err.Error())
		return
	}
	defaultClubID := params.ClubID
	// fmt.Println(params.ClubID)
	userService := u.userService

	if defaultClubID == "" {
		token, _ := u.GetTokenCliamsFromRequest(req)
		foundUserResp, err := userService.FindUserByID(ctx, &userPB.QueryUserRequest{Id: token.UserID})
		if err != nil {
			ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
			return
		}
		defaultClubID = foundUserResp.User.DefaultClubID
		if defaultClubID == "" {
			ErrorResponse(rsp, u.Error().BadRequest("ClubID can't be empty"))
			return
		}
	}

	userListResp, err := userService.FindUsersByClubID(ctx, &userPB.FindUsersByClubIDRequest{Page: int64(params.Page), Limit: int64(params.Limit), ClubID: defaultClubID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	list := make([]*dto.User, 0)

	for _, v := range userListResp.Users {
		list = append(list, dto.PBToUser(v))
	}
	SuccessResponse(rsp, D{
		"CurrentPage": params.Page,
		"PageSize":    params.Limit,
		"PageTotal":   userListResp.Count,
		"Total":       userListResp.Count,
		"List":        list,
	})
}

func (u *UserHandler) GetCurrentOrganization(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	clubID := ""

	userID := u.GetUserIDFromRequest(req)

	userSrv := u.userService

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: userID})
	if err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}
	clubID = foundUserResp.User.DefaultClubID

	if clubID == "" {
		SuccessResponse(rsp, D{
			"CurrentOrganization": &dto.Organization{},
		})
		return
	}

	clubService := u.clubService

	clubDetailsResp, err := clubService.FindClubDetailsByID(ctx, &clubPB.GetClubByIDRequest{ID: clubID})
	if err != nil {
		ErrorResponse(rsp, err)
	}

	res := dto.PBToOrganization(clubDetailsResp.Club)

	SuccessResponse(rsp, D{
		"CurrentOrganization": res,
	})
}

func (u *UserHandler) GetAllMembers(req *go_api.Request, rsp *go_api.Response) {
	params := struct {
		Page           int64  `query:"page"`
		Limit          int64  `query:"limit"`
		Department     string `json:"department" query:"department"`
		Job            string `json:"job" query:"job"`
		OrganizationID string `json:"organization" query:"organization"`
		Category       string `json:"category" query:"category"`
		Search         string `json:"search" query:"search"`
	}{}

	ctx := context.Background()

	if err := u.Bind(req, &params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if params.OrganizationID == "" {
		params.OrganizationID = u.GetCurrentClubIDFromRequest(req)
	}

	reply, err := u.clubService.FindUserClubProfiles(ctx, &clubPB.FindUserClubProfilesRequest{
		ClubID:       params.OrganizationID,
		Page:         params.Page,
		Limit:        params.Limit,
		DepartmentID: params.Department,
		JobID:        params.Job,
		Category:     params.Category,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	// for reply.UserClubProfiles.
	ucps := make([]*dto.OrganizationUser, 0)
	for _, v := range reply.UserClubProfiles {
		ucps = append(ucps, dto.PBToOrganizationUser(v))
	}
	SuccessResponseWithPage(rsp, params.Page, params.Limit, int64(reply.Total), ucps)
}

func (u *UserHandler) CreateMember(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Name       string   `json:"name"`
		Mobile     string   `json:"mobile"`
		Email      string   `json:"email"`
		Job        string   `json:"job" validate:"nonzero,objectid"`
		Department string   `json:"department" validate:"nonzero,objectid"`
		Roles      []string `json:"roles"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		ErrorResponse(rsp, u.Error().InternalServerError(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, err.Error())
		return
	}

	userSrv := u.userService

	userID := u.GetUserIDFromRequest(req)

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: userID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	name := []rune(params.Name)
	firstname := string(name[0:1])
	lastname := string(name[1:])

	createdResp, err := userSrv.CreateMember(ctx, &userPB.CreateMemberRequest{
		User: &userPB.User{
			Profile: &userPB.Profile{
				Firstname: firstname,
				Lastname:  lastname,
				Email:     params.Email,
				Mobile:    params.Mobile,
			},
		},
		JobID:        params.Job,
		DepartmentID: params.Department,
		ClubID:       foundUserResp.User.DefaultClubID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !createdResp.OK {
		ErrorResponse(rsp, u.Error().BadRequest("CreatedError"))
		return
	}

	SuccessResponse(rsp, D{})
	// panic("not implemented")
}

func (u *UserHandler) RemoveMemberFromOrg(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id" validate:"nonzero,objectid"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	// userSrv, found := .UserServiceFromContext(ctx)
	clubSrv := u.clubService

	userID := u.GetUserIDFromRequest(req)

	removeResp, err := clubSrv.RemoveUserFromClub(ctx, &clubPB.RemoveUserFromClubRequest{UserID: userID, ClubID: params.ID})
	// removeResp, err := userSrv.RemoveUserFromClub(ctx, &userPB.RemoveFromClubRequest{UserID: userID, ClubID: params.ID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !removeResp.OK {
		ErrorResponse(rsp, "Action Error")
		return
	}

	SuccessResponse(rsp, D{})
}

func (u *UserHandler) UpdateMember(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}

func (u *UserHandler) GetMemberDetails(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		ID     string `json:"id,omitempty"`
		ClubID string `json:"org,omitempty"`
	}{}

	if err := u.Bind(req, &params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	clubSrv := u.clubService

	if params.ID == "" {
		params.ID = u.GetUserIDFromRequest(req)
	}

	if params.ClubID == "" {
		params.ClubID = u.GetCurrentClubIDFromRequest(req)
	}

	ucpResp, err := clubSrv.GetUserClubProfileDetailsByID(ctx, &clubPB.GetUserClubProfileDetailsByIDRequest{
		OrganizationID: params.ClubID,
		UserID:         params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	SuccessResponse(rsp, D{
		"Details": dto.PBToOrganizationUser(ucpResp.UserClubProfile),
	})
}

func (u *UserHandler) RemoveOrg(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}

func (u *UserHandler) UpdateUserInfo(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
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
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	if err := u.Validate(&params); err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
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
		bir := time.Unix(*params.Birthday/1e3, 0)
		fieldsMap["birthday"] = bir.Local().Format(time.RFC1123Z)
	}
	if params.Email != nil {
		fieldsMap["email"] = *params.Email
	}
	if params.Gender != nil {
		fieldsMap["gender"] = *params.Gender
	}
	if params.SchoolDepartment != nil {
		fieldsMap["school_department"] = *params.SchoolDepartment
	}
	if params.SchoolClass != nil {
		fieldsMap["school_class"] = *params.SchoolClass
	}
	if params.Major != nil {
		fieldsMap["major"] = *params.Major
	}
	if params.RoomNumber != nil {
		fieldsMap["room_number"] = *params.RoomNumber
	}
	if params.AdvisorMobile != nil {
		fieldsMap["advisor_mobile"] = *params.AdvisorMobile
	}
	if params.AdvisorName != nil {
		fieldsMap["advisor_name"] = *params.AdvisorName
	}
	if params.StudentID != nil {
		fieldsMap["student_id"] = *params.StudentID
	}

	profile, err := json.Marshal(fieldsMap)

	if err != nil {
		ErrorResponse(rsp, u.Error().InternalServerError(err.Error()))
		return
	}

	userSrv := u.userService

	updateResp, err := userSrv.UpdateUser(ctx, &userPB.UpdateUserRequest{
		ID:      *params.ID,
		Profile: profile,
	})

	if err != nil {
		ErrorResponse(rsp, u.Error().InternalServerError(err.Error()))
		return
	}

	if !updateResp.OK {
		ErrorResponse(rsp, u.Error().BadRequest("更新失败"))
		return
	}

	SuccessResponse(rsp, D{})
}

func (u *UserHandler) FlagMemberState(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}

func (u *UserHandler) GetHotUsers(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}

func (u *UserHandler) UploadAvatar(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()

	request := req.Request

	_, fh, err := request.FormFile("file")
	if err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	// type savedResp struct {
	// 	filePB *iunite_club_srv_storage.FilePB
	// 	err    error
	// }

	// saveCh := make(chan savedResp, 1)

	filePB, err := file.SaveFile(fh)

	if err != nil {
		ErrorResponse(rsp, u.Error().BadRequest(err.Error()))
		return
	}

	fileResp, err := u.fileService.SaveFileInfo(ctx, &storagePB.FileInfoRequest{
		File: filePB,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	// Update Avatar
	// fileResp.File.ID
	_, err = u.userService.UpdateAvatar(ctx, &userPB.UpdateAvatarRequest{
		ID:     u.GetUserIDFromRequest(req),
		Avatar: fileResp.File.ID,
	})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

func (u *UserHandler) ExportList(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}

func (u *UserHandler) DownloadExportTemplate(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}

func (u *UserHandler) UploadUserList(*go_api.Request, *go_api.Response) {
	panic("not implemented")
}
