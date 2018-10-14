package handler

import (
	"context"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"

	"iunite.club/services/navo/dto"

	go_api "github.com/micro/go-api/proto"
	approvedPB "iunite.club/services/core/proto/approved"
	"iunite.club/services/navo/client"
)

type ApprovedHandler struct {
	BaseHandler
	approvedService approvedPB.ApprovedService
}

func (a *ApprovedHandler) getApprovedService(ctx context.Context) approvedPB.ApprovedService {
	if a.approvedService == nil {
		srv, ok := client.ApprovedServiceFromContext(ctx)
		if !ok {
			panic("not found approved service")
		}
		a.approvedService = srv
	}

	return a.approvedService
}

func (a *ApprovedHandler) List(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {

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
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	if err := a.Validate(&params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	approvedService := a.getApprovedService(ctx)

	handlerID := a.GetUserIDFromRequest(req)
	currentClubID := a.GetCurrentClubIDFromRequest(ctx, req)

	listResp, err := approvedService.List(ctx, &approvedPB.ListRequest{
		ClubID:    currentClubID,
		Kind:      params.Kind,
		Status:    params.Status,
		Search:    params.Search,
		HandlerID: handlerID,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}
	approveds := make([]*dto.ApprovedTask, 0, len(listResp.Approveds))
	if len(listResp.Approveds) > 0 {

		for _, v := range listResp.Approveds {
			approveds = append(approveds, dto.PBToApprovedTask(v))
		}

	}
	return SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, approveds)
}

func (a *ApprovedHandler) Details(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		ID string `query:"id"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	if err := a.Validate(&params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	approvedService := a.getApprovedService(ctx)

	detailsResp, err := approvedService.Details(ctx, &approvedPB.DetailsRequest{ID: params.ID})
	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{
		"approved": dto.PBToApprovedTask(detailsResp.Approved),
		"details":  hptypes.DecodeToMap(detailsResp.Approved.Content),
	})
}

func (a *ApprovedHandler) ExecuteOne(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	params := struct {
		ID      string `json:"id"`
		Result  bool   `json:"result"`
		Options string `json:"content"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	if err := a.Validate(&params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	approvedService := a.getApprovedService(ctx)

	_, err := approvedService.Execute(ctx, &approvedPB.ExecuteRequest{
		ID:      params.ID,
		Result:  params.Result,
		Options: params.Options,
		ClubID:  a.GetCurrentClubIDFromRequest(ctx, req),
		UserID:  a.GetUserIDFromRequest(req),
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{})
}

func (a *ApprovedHandler) Archive(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (a *ApprovedHandler) BatchArchive(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (a *ApprovedHandler) WaitingTaskList(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (a *ApprovedHandler) ListV2(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
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
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	if err := a.Validate(&params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	approvedService := a.getApprovedService(ctx)
	listResp, err := approvedService.ListV2(ctx, &approvedPB.ListV2Request{
		Page:       params.Page,
		Limit:      params.Limit,
		ClubID:     params.ClubID,
		Search:     params.Search,
		ReadState:  params.ReadState,
		FlowStatus: params.Status,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}
	approveds := make([]*dto.ApprovedTask, 0, len(listResp.Approveds))
	if len(listResp.Approveds) > 0 {

		for _, v := range listResp.Approveds {
			approveds = append(approveds, dto.PBToApprovedTask(v))
		}

	}
	return SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, approveds)
}

func (a *ApprovedHandler) ListByPusher(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	params := struct {
		Page   int64  `json:"page,omitempty" query:"page"`
		Limit  int64  `json:"limit,omitempty" query:"limit"`
		UserID string `query:"userID"`
		Search string `query:"search"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	if err := a.Validate(&params); err != nil {
		return ErrorResponse(rsp, a.Error(ctx).BadRequest(err.Error()))
	}

	approvedService := a.getApprovedService(ctx)
	listResp, err := approvedService.ListByPusher(ctx, &approvedPB.ListByPusherRequest{
		UserID: a.GetUserIDFromRequest(req),
		ClubID: a.GetCurrentClubIDFromRequest(ctx, req),
		Page:   params.Page,
		Limit:  params.Limit,
		Search: params.Search,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	approveds := make([]*dto.ApprovedTask, 0, len(listResp.Approveds))
	if len(listResp.Approveds) > 0 {

		for _, v := range listResp.Approveds {
			approveds = append(approveds, dto.PBToApprovedTask(v))
		}

	}
	return SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, approveds)
}
