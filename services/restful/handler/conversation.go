package handler

import (
	"context"
	"net/url"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	pb "iunite.club/services/core/proto/conversation"
	"iunite.club/services/restful/dto"
)

type Conversation struct {
	BaseHandler

	conversationService pb.ConversationService
}

func NewConversationHandler(c client.Client) *Conversation {
	return &Conversation{
		conversationService: pb.NewConversationService(CoreService, c),
	}
}

// GetConversationsByMemberID 通过成员ID获取会话列表
func (conv *Conversation) GetConversationsByMemberID(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		UserID string `query:"userID" validate:"objectid"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := conv.conversationService.GetConversationsByMemberID(ctx, &pb.ByUserID{
		ID: params.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	pbConvs := make([]*dto.Conversation, 0, len(reply.Conversations))

	for _, val := range reply.Conversations {
		pbConvs = append(pbConvs, dto.PBToConversation(val))
	}
	// reply.

	SuccessResponse(rsp, D{
		"Conversations": pbConvs,
	})
	return
}

// CreateConversation 创建一个会话
func (conv *Conversation) CreateConversation(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		// 会话ID
		ConversationID string `json:"conversationID"`

		// 会话类型 simple, group, temporary
		Kind string `json:"kind"`

		// 会话名称，群会话时需要设置
		Name string `json:"name"`

		// 会话头像，群会话时需要设置
		Avatar string `json:"avatar"`

		// 会话成员 [用户ID]
		Members []string `json:"members"`

		// 会话管理员ID, 设置为会话的创建者的ID
		Master string `json:"master"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := conv.conversationService.CreateConversation(ctx, &pb.WithConversationBundle{
		Kind:            params.Kind,
		Name:            params.Name,
		Avatar:          params.Avatar,
		Members:         params.Members,
		IsStartValidate: false,
		IsTop:           false,
		ConversationID:  params.ConversationID,
		Master:          params.Master,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	detailsReply, err := conv.conversationService.FindConversationDetails(ctx, &pb.ByID{
		ID: reply.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		// "Conversation":,
		"Conversation": dto.PBToConversation(detailsReply.Conversation),
		"Metadata":     dto.PBToConversatoinMetaData(reply.MetaData),
		"IsExists":     reply.IsExists,

	})
}

func (conv *Conversation) GetConversationDetails(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		ID string `query:"id"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}
	pbReq := &pb.ByID{
		ID: params.ID,
	}
	pbReq.Put = true
	reply, err := conv.conversationService.FindConversationDetails(ctx, pbReq)

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	urlQuery := url.Values{}
	urlQuery.Set("pg", "conversation.group.details")
	urlQuery.Add("do", "showGroupDetails")
	urlQuery.Add("id", reply.Conversation.ID)

	// UpdateConversationMetaDataByID 更新leancloud远端元数据

	SuccessResponse(rsp, D{
		"Details":    dto.PBToConversation(reply.Conversation),
		"QRCodeBody": "unite://client//action?" + urlQuery.Encode(),
	})
}

func (conv *Conversation) ExitGroup(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		ID string `query:"id"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := conv.conversationService.ExitGroup(ctx, &pb.ByIDWithUserID{
		ID:     params.ID,
		UserID: conv.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) DismissGroup(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		ID string `query:"id"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := conv.conversationService.DismissGroup(ctx, &pb.ByIDWithUserID{
		ID:     params.ID,
		UserID: conv.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) UpdateGroupConversation(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		ID              string `json:"id"`
		Nickname        string `json:"nickname"`
		GroupName       string `json:"groupName"`
		IsTop           string `json:"isTop"`
		IsStartValidate string `json:"isStartValidate"`
		Master          string `json:"master"`
		Avatar          string `json:"avatar"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := conv.conversationService.UpdateGroupConversation(ctx, &pb.WithUpdateBundle{
		ID:              params.ID,
		Nickname:        params.Nickname,
		GroupName:       params.GroupName,
		IsTop:           params.IsTop,
		IsStartValidate: params.IsStartValidate,
		Master:          params.Master,
		Avatar:          params.Avatar,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	// UpdateConversationMetaDataByID 更新leancloud远端元数据

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) CreateGroupNotice(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		Title          string `json:"title,omitempty"`
		Body           string `json:"body,omitempty"`
		ConversationID string `json:"conversationID,omitempty"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := conv.conversationService.CreateNotice(ctx, &pb.WithNoticeBundle{
		ConversationID: params.ConversationID,
		Title:          params.Title,
		Body:           params.Body,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) MarkNoticeToHasRead(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		NoticeID string `json:"noticeID,omitempty"`
		UserID   string `json:"userID,omitempty"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := conv.conversationService.MarkedNoticeToHasRead(ctx, &pb.WithMarkedBundle{
		NoticeID: params.NoticeID,
		UserID:   params.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

// GetNoticeList 获取群公告列表
// Path:
func (conv *Conversation) GetNoticeList(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Page           int32  `json:"page,omitempty" query:"page"`
		Limit          int32  `json:"limit,omitempty" query:"limit"`
		ConversationID string `json:"conversation_id,omitempty" query:"conversation_id"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reply, err := conv.conversationService.GetNoticeList(ctx, &pb.ByIDWithPager{
		Page:  int64(params.Page),
		Limit: int64(params.Limit),
		ID:    params.ConversationID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	pbNoticeList := make([]*dto.ConversationNotice, 0, len(reply.Notices))

	for _, val := range reply.Notices {
		pbNoticeList = append(pbNoticeList, dto.PBToConversationNotice(val))
	}

	SuccessResponseWithPage(rsp, int64(params.Page), int64(params.Limit), reply.Total, pbNoticeList)
	return
}

func (conv *Conversation) RemoveConversationNotice(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID string
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}
	ctx := context.Background()

	_, err := conv.conversationService.RemoveConversationNotice(ctx, &pb.ByNoticeID{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) RemoveGroupMember(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ConversationID string
		Members        []string
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}
	ctx := context.Background()

	_, err := conv.conversationService.RemoveGroupMember(ctx, &pb.WithIDAndMembers{
		ID:      params.ConversationID,
		Members: params.Members,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) AddGroupMember(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ConversationID string
		Members        []string
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	ctx := context.Background()

	_, err := conv.conversationService.AddGroupMember(ctx, &pb.WithIDAndMembers{
		ID:      params.ConversationID,
		Members: params.Members,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (conv *Conversation) GetAllMembersOfConversation(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID string `query:"id"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	ctx := context.Background()

	reply, err := conv.conversationService.GetAllMembersOfConversation(ctx, &pb.ByID{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	members := make([]*dto.User, 0, len(reply.Members))

	for _, val := range reply.Members {
		members = append(members, dto.PBToUser(val))
	}

	SuccessResponse(rsp, D{
		"Members": members,
	})
	return
}

// JoinGroup 加入群组
// Path:
func (conv *Conversation) JoinGroup(req *restful.Request, rsp *restful.Response) {
	params := struct {
		GroupID string `json:"groupID"`
		UserID  string `json:"userID"`
	}{}

	if err := conv.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	ctx := context.Background()

	_, err := conv.conversationService.JoinInGroup(ctx, &pb.WithIDAndUserID{
		ID:     params.GroupID,
		UserID: params.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}
