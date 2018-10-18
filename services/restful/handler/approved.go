package handler

import (
	"context"
	"fmt"

	go_api "github.com/emicklei/go-restful"
	c "github.com/micro/go-micro/client"

	"iunite.club/services/restful/dto"

	approvedPB "iunite.club/services/core/proto/approved"
	// "iunite.club/services/navo/client"
)

type ApprovedHandler struct {
	BaseHandler
	approvedService approvedPB.ApprovedService
}

func NewApprovedHandler(client c.Client) *ApprovedHandler {
	return &ApprovedHandler{
		approvedService: approvedPB.NewApprovedService(CoreService, client),
	}
}

// func (a *ApprovedHandler) getApprovedService(ctx context.Context) approvedPB.ApprovedService {
// 	if a.approvedService == nil {
// 		srv, ok := client.ApprovedServiceFromContext(ctx)
// 		if !ok {
// 			panic("not found approved service")
// 		}
// 		a.approvedService = srv
// 	}

// 	return a.approvedService
// }

func (a *ApprovedHandler) List(req *go_api.Request, rsp *go_api.Response) {

	params := struct {
		Page      int64  `json:"page,omitempty" query:"page"`
		Limit     int64  `json:"limit,omitempty" query:"limit"`
		ClubID    string `json:"org,omitempty" query:"org" validate:"objectid"`
		Kind      string `json:"kind,omitempty" query:"kind"`
		Status    string `json:"status,omitempty" query:"status"`
		ReadState string `json:"readState,omitempty" query:"readState"`
		Search    string `json:"search,omitempty" query:"search"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	approvedService := a.approvedService

	handlerID := a.GetUserIDFromRequest(req)
	currentClubID := a.GetCurrentClubIDFromRequest(req)
	ctx := context.Background()

	listResp, err := approvedService.List(ctx, &approvedPB.ListRequest{
		ClubID:    currentClubID,
		Kind:      params.Kind,
		Status:    params.Status,
		Search:    params.Search,
		HandlerID: handlerID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	approveds := make([]*dto.ApprovedTask, 0, len(listResp.Approveds))
	if len(listResp.Approveds) > 0 {

		for _, v := range listResp.Approveds {
			approveds = append(approveds, dto.PBToApprovedTask(v))
		}

	}
	SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, approveds)
	return
}

func (a *ApprovedHandler) Details(req *go_api.Request, rsp *go_api.Response) {
	params := struct {
		ID string `query:"id"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	approvedService := a.approvedService

	detailsResp, err := approvedService.Details(context.Background(), &approvedPB.DetailsRequest{ID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	// pusher := detailsResp.Approved.Pu
	SuccessResponse(rsp, D{
		"approved": dto.PBToApprovedTask(detailsResp.Approved),
		// "details":  hptypes.DecodeToMap(detailsResp.Approved.Content),
		"details": dto.GetApprovedContent(detailsResp.Approved),
	})
	return
}

func (a *ApprovedHandler) ExecuteOne(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	params := struct {
		ID      string `json:"id"`
		Result  bool   `json:"result"`
		Options string `json:"content"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	approvedService := a.approvedService
	ctx := context.Background()
	_, err := approvedService.Execute(ctx, &approvedPB.ExecuteRequest{
		ID:      params.ID,
		Result:  params.Result,
		Options: params.Options,
		ClubID:  a.GetCurrentClubIDFromRequest(req),
		UserID:  a.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

// func (a *ApprovedHandler) Archive(context.Context, *go_api.Request, *go_api.Response) error {
// 	panic("not implemented")
// }

// func (a *ApprovedHandler) BatchArchive(context.Context, *go_api.Request, *go_api.Response) error {
// 	panic("not implemented")
// }

// func (a *ApprovedHandler) WaitingTaskList(context.Context, *go_api.Request, *go_api.Response) error {
// 	panic("not implemented")
// }

func (a *ApprovedHandler) ListV2(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	fmt.Println("list v2")
	params := struct {
		Page      int64  `json:"page,omitempty" query:"page"`
		Limit     int64  `json:"limit,omitempty" query:"limit"`
		ClubID    string `json:"org,omitempty" query:"org" validate:"objectid"`
		Kind      string `json:"kind,omitempty" query:"kind"`
		Status    string `json:"status,omitempty" query:"status"`
		ReadState string `json:"readState,omitempty" query:"readState"`
		Search    string `json:"search,omitempty" query:"search"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if params.ClubID == "" {
		params.ClubID = a.GetCurrentClubIDFromRequest(req)
	}

	approvedService := a.approvedService
	listResp, err := approvedService.ListV2(ctx, &approvedPB.ListV2Request{
		HandlerID:  a.GetUserIDFromRequest(req),
		Page:       params.Page,
		Limit:      params.Limit,
		ClubID:     params.ClubID,
		Search:     params.Search,
		ReadState:  params.ReadState,
		FlowStatus: params.Status,
	})

	if err != nil {
		fmt.Println(err.Error(), "error")
		ErrorResponse(rsp, err)
		return
	}
	approveds := make([]*dto.ApprovedTask, 0, len(listResp.Approveds))
	if len(listResp.Approveds) > 0 {

		for _, v := range listResp.Approveds {
			approveds = append(approveds, dto.PBToApprovedTask(v))
		}

	}
	SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, approveds)
	return
}

func (a *ApprovedHandler) ListByPusher(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Page   int64  `json:"page,omitempty" query:"page"`
		Limit  int64  `json:"limit,omitempty" query:"limit"`
		UserID string `query:"userID"`
		Search string `query:"search"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}
	userID := params.UserID

	if userID == "" {
		userID = a.GetUserIDFromRequest(req)
	}

	approvedService := a.approvedService
	listResp, err := approvedService.ListByPusher(ctx, &approvedPB.ListByPusherRequest{
		UserID: userID,
		ClubID: a.GetCurrentClubIDFromRequest(req),
		Page:   params.Page,
		Limit:  params.Limit,
		Search: params.Search,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	approveds := make([]*dto.ApprovedTask, 0, len(listResp.Approveds))
	if len(listResp.Approveds) > 0 {

		for _, v := range listResp.Approveds {
			approveds = append(approveds, dto.PBToApprovedTask(v))
		}

	}
	SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, approveds)
	return
}
