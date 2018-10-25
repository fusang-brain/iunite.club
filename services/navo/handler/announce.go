package handler

import (
	"context"

	"github.com/emicklei/go-restful"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/micro/go-micro/client"
	announcePB "iunite.club/services/core/proto/announce"
	"iunite.club/services/navo/dto/announce"
)

type AnnounceHandler struct {
	BaseHandler

	announceService announcePB.AnnounceService
}

func NewAnnounceHandler(c client.Client) *AnnounceHandler {
	return &AnnounceHandler{
		announceService: announcePB.NewAnnounceService(CoreService, c),
	}
}

func (self *AnnounceHandler) CreateInstructions(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := dto_announce.CreateInstructionsBundle{}
	if err := self.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	createdResp, err := self.announceService.CreateInstructions(ctx, &announcePB.CreateInstructionsRequest{
		Name:   params.Name,
		Body:   params.Body,
		ClubID: params.ClubID,
		UserID: self.GetUserIDFromRequest(req),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	announce := createdResp.Announce
	WriteJsonResponse(rsp, D{
		"CreatedAt": hptypes.Timestamp(announce.CreatedAt),
		"ID":        announce.ID,
	})
	return
}

func (self *AnnounceHandler) CreateReminder(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := dto_announce.CreateReminderBundle{}

	if err := self.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	reminderResp, err := self.announceService.CreateReminder(ctx, &announcePB.CreateReminderRequest{
		Name:         params.Name,
		Body:         params.Body,
		ReminderTime: hptypes.TimestampProto(params.ReminderTime),
		Users:        params.Users,
		ClubID:       params.ClubID,
		UserID:       self.GetUserIDFromRequest(req),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	announce := reminderResp.Announce
	WriteJsonResponse(rsp, D{
		"CreatedAt": hptypes.Timestamp(announce.CreatedAt),
		"ID":        announce.ID,
	})
	return
}

func (self *AnnounceHandler) CreateTask(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := dto_announce.CreateTaskBundle{}

	if err := self.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	taskResp, err := self.announceService.CreateTask(ctx, &announcePB.CreateTaskRequest{
		Name: params.Name,
		Body: params.Body,
		// ReminderTime: hptypes.TimestampProto(time.Unix(params.ReminderTime/1e3, 0)),
		StartTime: hptypes.TimestampProto(params.StartTime),
		EndTime:   hptypes.TimestampProto(params.EndTime),
		Users:     params.Users,
		ClubID:    params.ClubID,
		UserID:    self.GetUserIDFromRequest(req),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	announce := taskResp.Announce
	WriteJsonResponse(rsp, D{
		"CreatedAt": hptypes.Timestamp(announce.CreatedAt),
		"ID":        announce.ID,
	})
	return
}

func (self *AnnounceHandler) GetAnnounces(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		Page   int32  `json:"page,omitempty" query:"page"`
		Limit  int32  `json:"limit,omitempty" query:"limit"`
		Kind   string `json:"kind,omitempty" query:"kind"`
		ClubID string `json:"club_id,omitempty" query:"club_id"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	listResp, err := self.announceService.GetAnnounces(ctx, &announcePB.GetAnnouncesRequest{
		Page:   params.Page,
		Limit:  params.Limit,
		Kind:   params.Kind,
		ClubID: params.ClubID,
		UserID: self.GetUserIDFromRequest(req),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}
	
	WriteJsonResponse(rsp, listResp)
	return
}

func (self *AnnounceHandler) MarkedToRead(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	// params := struct {
	// 	ID string `json:"id"`
	// }{}

	// if err := self.BindAndValidate(req, &params); err != nil {
	// 	WriteError(rsp, err)
	// 	return
	// }

	ID := req.PathParameter("id")

	_, err := self.announceService.MarkedOneToRead(ctx, &announcePB.MarkedOneToReadRequest{
		ID:     ID,
		UserID: self.GetUserIDFromRequest(req),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, D{
		"ID": ID,
	})
	return
}
