package handler

import (
	"context"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	approvedPB "iunite.club/services/core/proto/approved"
)

type TodoHandler struct {
	BaseHandler

	approvedService approvedPB.ApprovedService
}

func NewTodoHandler(c client.Client) *TodoHandler {
	return &TodoHandler{
		approvedService: approvedPB.NewApprovedService(CoreService, c),
	}
}

func (self *TodoHandler) GetTaskCount(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	params := struct {
		UserID string `query:"userID"`
	}{}

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if params.UserID == "" {
		params.UserID = self.GetUserIDFromRequest(req)
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	countResp, err := self.approvedService.GetPendingApprovedCountByUserID(ctx, &approvedPB.GetPendingApprovedCountRequest{
		UserID: params.UserID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
	}

	counts := countResp.Counts
	result := make(map[string]int64)
	for _, v := range counts {
		result[v.ClubID] = v.Count
	}

	SuccessResponse(rsp, D{
		"Counts": result,
	})
}
