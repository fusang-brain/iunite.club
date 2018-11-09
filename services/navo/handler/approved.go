package handler

import (
	"context"
	"encoding/json"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	pb "iunite.club/services/approved/proto"
	"iunite.club/services/navo/dto/approved"
)

type ApprovedHandler struct {
	BaseHandler

	approvedService pb.ApprovedService
}

func NewApprovedHandler(c client.Client) *ApprovedHandler {
	return &ApprovedHandler{
		approvedService: pb.NewApprovedService(ApprovedService, c),
	}
}

func (a *ApprovedHandler) PostTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := new(dto_approved.TemplateBundle)
	if err := a.BindAndValidate(req, params); err != nil {
		WriteError(rsp, err)
		return
	}
	temp := new(pb.ApprovedTemplatePB)

	if b, err := json.Marshal(params); err == nil {
		json.Unmarshal(b, temp)
	}

	temp.UserID = a.GetUserIDFromRequest(req)

	postedResp, err := a.approvedService.PostTemplate(ctx, &pb.PostTemplateRequest{
		Template: temp,
	})
	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, D{
		"ID":        postedResp.Template.ID,
		"CreatedAt": hptypes.Timestamp(postedResp.Template.CreatedAt),
	})
}

func (a *ApprovedHandler) FindTemplates(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Page      int32  `json:"page,omitempty" query:"page"`
		Limit     int32  `json:"limit,omitempty" query:"limit"`
		ClubID    string `json:"club_id,omitempty" query:"club_id"`
		Populates string `json:"populates,omitempty" query:"populates"`
	}{}

	if err := a.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	populates := strings.Split(params.Populates, ",")
	templatesResp, err := a.approvedService.FindTemplates(ctx, &pb.FindTemplatesRequest{
		Page:     params.Page,
		Limit:    params.Limit,
		ClubID:   params.ClubID,
		Populate: populates,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, templatesResp)
}

func (a *ApprovedHandler) DeleteTemplates(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	id := req.PathParameter("id")

	reply, err := a.approvedService.DeleteTemplate(ctx, &pb.DeleteTemplateRequest{
		ID: id,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
}

func (a *ApprovedHandler) UpdateTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	id := req.PathParameter("id")

	updateFields := make(map[string]interface{})

	if err := json.NewDecoder(req.Request.Body).Decode(&updateFields); err != nil {
		WriteError(rsp, a.Error().InternalServerError(err.Error()))
		return
	}

	// params := new(dto_approved.TemplateBundle)
	// if err := a.BindAndValidate(req, params); err != nil {
	// 	WriteError(rsp, err)
	// 	return
	// }

	reply, err := a.approvedService.UpdateTemplate(ctx, &pb.UpdateTemplateRequest{
		ID:     id,
		Fields: hptypes.EncodeToStruct(updateFields),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
}

func (a *ApprovedHandler) EnableTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	id := req.PathParameter("id")
	reply, err := a.approvedService.ToggleTemplateEnableState(ctx, &pb.ToggleEnableStateReq{
		ID:        id,
		IsEnabled: true,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
}

func (a *ApprovedHandler) DisableTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	id := req.PathParameter("id")
	reply, err := a.approvedService.ToggleTemplateEnableState(ctx, &pb.ToggleEnableStateReq{
		ID:        id,
		IsEnabled: false,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
}
