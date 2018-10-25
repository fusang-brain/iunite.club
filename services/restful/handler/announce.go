package handler

import (
	"context"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/micro/go-micro/client"
	announcePB "iunite.club/services/core/proto/announce"
	"iunite.club/services/restful/dto"
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
	params := struct {
		Name string `json:"name"`
		Body string `json:"body"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	createdResp, err := self.announceService.CreateInstructions(ctx, &announcePB.CreateInstructionsRequest{
		Name:   params.Name,
		Body:   params.Body,
		ClubID: self.GetCurrentClubIDFromRequest(req),
		UserID: self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	createdAnnounce := dto.PBToAnnounce(createdResp.Announce)

	SuccessResponse(rsp, D{
		"Announce": createdAnnounce,
	})
}

func (self *AnnounceHandler) CreateReminder(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Name         string   `json:"name,omitempty"`
		Body         string   `json:"body,omitempty"`
		ReminderTime int64    `json:"reminder_time,omitempty"`
		Users        []string `json:"users,omitempty"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reminderResp, err := self.announceService.CreateReminder(ctx, &announcePB.CreateReminderRequest{
		Name:         params.Name,
		Body:         params.Body,
		ReminderTime: hptypes.TimestampProto(time.Unix(params.ReminderTime/1e3, 0)),
		Users:        params.Users,
		ClubID:       self.GetCurrentClubIDFromRequest(req),
		UserID:       self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Announce": dto.PBToAnnounce(reminderResp.Announce),
	})
}

func (self *AnnounceHandler) CreateTask(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Name         string   `json:"name,omitempty"`
		Body         string   `json:"body,omitempty"`
		ReminderTime int64    `json:"reminder_time,omitempty"`
		Users        []string `json:"users,omitempty"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	reminderResp, err := self.announceService.CreateTask(ctx, &announcePB.CreateTaskRequest{
		Name: params.Name,
		Body: params.Body,
		// ReminderTime: hptypes.TimestampProto(time.Unix(params.ReminderTime/1e3, 0)),
		StartTime: hptypes.TimestampProto(time.Unix(params.ReminderTime/1e3, 0)),
		EndTime:   hptypes.TimestampProto(time.Unix(params.ReminderTime/1e3, 0).Add(time.Hour * 7 * 24)),
		Users:     params.Users,
		ClubID:    self.GetCurrentClubIDFromRequest(req),
		UserID:    self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Announce": dto.PBToAnnounce(reminderResp.Announce),
	})
}

func (self *AnnounceHandler) GetAnnounces(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		Page  int32  `json:"page,omitempty" query:"page"`
		Limit int32  `json:"limit,omitempty" query:"limit"`
		Kind  string `json:"kind,omitempty" query:"kind"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	listResp, err := self.announceService.GetAnnounces(ctx, &announcePB.GetAnnouncesRequest{
		Page:   params.Page,
		Limit:  params.Limit,
		Kind:   params.Kind,
		ClubID: self.GetCurrentClubIDFromRequest(req),
		UserID: self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	listCount := len(listResp.Announces)

	announces := make([]*dto.Announce, 0, listCount)
	for _, v := range listResp.Announces {
		announces = append(announces, dto.PBToAnnounce(v))
	}

	SuccessResponseWithPage(rsp, int64(params.Page), int64(params.Limit), int64(listResp.Total), announces)
}

func (self *AnnounceHandler) MarkedToRead(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		ID string `json:"id"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	_, err := self.announceService.MarkedOneToRead(ctx, &announcePB.MarkedOneToReadRequest{
		ID:     params.ID,
		UserID: self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}
