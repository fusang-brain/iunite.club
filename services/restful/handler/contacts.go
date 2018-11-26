package handler

import (
	"context"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	pb "iunite.club/services/core/proto/contacts"
	"iunite.club/services/restful/dto"
)

type ContactsHandler struct {
	BaseHandler

	contactsService pb.ContactsService
}

func NewContactsHandler(c client.Client) *ContactsHandler {
	return &ContactsHandler{
		contactsService: pb.NewContactsService(CoreService, c),
	}
}

func (hdl *ContactsHandler) GetFriendList(req *restful.Request, rsp *restful.Response) {
	params := struct {
		UserID string `query:"userID"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}
	ctx := context.Background()

	reply, err := hdl.contactsService.FindFriendList(ctx, &pb.FindFriendListRequest{
		UserID: params.UserID,
	})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoUsers := make([]*dto.User, 0)
	for _, user := range reply.Users {
		dtoUsers = append(dtoUsers, dto.PBToUser(user))
	}

	SuccessResponse(rsp, D{
		"Users": dtoUsers,
	})
	return
}

func (hdl *ContactsHandler) GetAllGroup(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		UserID string `query:"userID"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}
	withUsers := false
	platform := hdl.GetPlatformFromRequest(req)

	if platform == PlatformWEB {
		withUsers = true
	}

	replay, err := hdl.contactsService.FindAllGroup(ctx, &pb.FindAllGroupRequest{
		UserID:    params.UserID,
		WithUsers: withUsers,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoConversations := make([]*dto.Conversation, 0)

	for _, conv := range replay.Conversations {
		dtoConversations = append(dtoConversations, dto.PBToConversation(conv))
	}

	SuccessResponse(rsp, D{
		"Conversations": dtoConversations,
	})

	return
}

func (hdl *ContactsHandler) GetContactsList(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		Page       int32  `query:"page"`
		Limit      int32  `query:"limit"`
		Search     string `query:"search"`
		Department string `query:"department"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.FindContactsList(ctx, &pb.FindContactsRequest{
		Search:     params.Search,
		Limit:      params.Limit,
		Page:       params.Page,
		Department: params.Department,
		ClubID:     hdl.GetCurrentClubIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoContacts := make([]*dto.OrganizationUser, 0)
	for _, c := range reply.Contacts {
		dtoContacts = append(dtoContacts, dto.PBToOrganizationUser(c))
	}

	SuccessResponseWithPage(rsp, int64(params.Page), int64(params.Limit), int64(reply.Total), dtoContacts)
	return
}

func (hdl *ContactsHandler) GetDepartmentGroupByUserID(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		UserID string `query:"userID"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.FindDepartmentGroupByUserID(ctx, &pb.UserIDRequest{
		UserID: params.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoOrganizatioins := make([]*dto.Organization, 0)
	dtoDepartments := make([]*dto.Department, 0)

	for _, org := range reply.Organizations {
		if org.Kind == "club" {
			dtoOrganizatioins = append(dtoOrganizatioins, dto.PBToOrganization(org))
			continue
		}

		if org.Kind == "department" {
			dtoDepartments = append(dtoDepartments, dto.PBToDepartment(org))
			continue
		}
	}

	SuccessResponse(rsp, D{
		"Organizations": dtoOrganizatioins,
		"Departments":   dtoDepartments,
	})

	return
}

func (hdl *ContactsHandler) GetUsersByDepartmentID(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		DepartmentID string `query:"department"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.FindUsersByDepartment(ctx, &pb.DepartmentRequest{
		DepartmentID: params.DepartmentID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	type userScan struct {
		dto.User
		Job dto.Job
	}

	rsUsers := make([]*userScan, 0)
	for _, val := range reply.UserJobs {
		dtoUser := dto.PBToUser(val.User)
		dtoJob := dto.PBToJob(val.Job)
		rsUsers = append(rsUsers, &userScan{
			User: *dtoUser,
			Job:  *dtoJob,
		})
	}

	SuccessResponse(rsp, D{
		"Users": rsUsers,
	})

	return
}

func (hdl *ContactsHandler) GetDepartmentsByOrganizationID(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		OrgID string `query:"orgID"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.FindDepartmentsByOrganization(ctx, &pb.OrgRequest{
		OrgID: params.OrgID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoDepartments := make([]*dto.Department, 0)

	for _, dept := range reply.Organizations {
		dtoDepartments = append(dtoDepartments, dto.PBToDepartment(dept))
	}

	SuccessResponse(rsp, D{
		"Departments": dtoDepartments,
		"Counts":      reply.Counts,
	})

	return
}

func (hdl *ContactsHandler) GetUserCardDetails(req *restful.Request, rsp *restful.Response) {
	// TODO change user card details
	params := struct {
		UserID string `query:"userID"`
	}{}
	ctx := context.Background()

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.GetUserCardDetails(ctx, &pb.UserIDRequest{
		UserID:        params.UserID,
		CurrentUserID: hdl.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoOrgUserInfos := make([]*dto.OrganizationUser, 0)
	for _, oui := range reply.OrganizationUserInfo {
		dtoOrgUserInfos = append(dtoOrgUserInfos, dto.PBToOrganizationUser(oui))
	}

	SuccessResponse(rsp, D{
		"BaseInfo":             dto.PBToUser(reply.BaseInfo),
		"OrganizationUserInfo": dtoOrgUserInfos,
		"IsFriend":             reply.IsFriend,
		"InSameOrganization":   reply.InSameOrganization,
	})

	return
}

func (hdl *ContactsHandler) SearchUser(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Mobile string `json:"mobile,omitempty" query:"mobile"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.SearchUsers(ctx, &pb.SearchUserRequest{
		Mobile: params.Mobile,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	dtoUsers := make([]*dto.User, 0)

	for _, user := range reply.Users {
		dtoUsers = append(dtoUsers, dto.PBToUser(user))
	}

	SuccessResponse(rsp, D{
		"Users": dtoUsers,
	})
	return
}

func (hdl *ContactsHandler) AddFriend(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		UserID   string `json:"user_id,omitempty"`
		FriendID string `json:"friend_id,omitempty"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.AddFriend(ctx, &pb.AddFriendRequest{
		UserID:   params.UserID,
		FriendID: params.FriendID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !reply.OK {
		ErrorResponse(rsp, hdl.Error().BadRequest("Error to add friend"))
		return
	}

	SuccessResponse(rsp, nil)
}

func (hdl *ContactsHandler) RemoveFriend(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		UserID   string `json:"user_id,omitempty"`
		FriendID string `json:"friend_id,omitempty"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := hdl.contactsService.RemoveFriend(ctx, &pb.RemoveFriendRequest{
		UserID:   params.UserID,
		FriendID: params.FriendID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, nil)
}

func (hdl *ContactsHandler) GetFriendAcceptList(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id,omitempty" query:"id"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.FindFriendAcceptList(ctx, &pb.UserIDRequest{
		UserID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	dtoList := make([]*dto.FriendAccept, 0)
	for _, item := range reply.AcceptList {
		// dtoList = append(dtoList, dto.)
		dtoList = append(dtoList, dto.PBToFriendAccept(item))
	}

	SuccessResponse(rsp, D{
		"AcceptList": dtoList,
	})
}

func (hdl *ContactsHandler) GetFrientAcceptCount(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id,omitempty" query:"id"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := hdl.contactsService.FindFrientAcceptCount(ctx, &pb.UserIDRequest{
		UserID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Count": reply.Count,
	})
	return
}

func (hdl *ContactsHandler) AgreeFriendAccept(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id,omitempty" query:"id"`
	}{}

	if err := hdl.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := hdl.contactsService.AgreeFriendAccept(ctx, &pb.AgreeFriendAcceptRequest{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, nil)
	return
}
