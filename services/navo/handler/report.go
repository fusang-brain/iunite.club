package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/jung-kurt/gofpdf"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/navo/dto/report"
	pb "iunite.club/services/report/proto"
)

// Reports 汇报控制器
type Reports struct {
	BaseHandler

	reportService pb.ReportService
}

// NewReportsHandler 创建汇报控制器
func NewReportsHandler(c client.Client) *Reports {
	return &Reports{
		reportService: pb.NewReportService(ReportService, c),
	}
}

// PostReport 发布汇报
func (r *Reports) PostReport(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := dto_report.SimpleReportBundle{}
	if err := r.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	reportResp, err := r.reportService.PostReport(ctx, &pb.PostReportBundle{
		Title:       params.Title,
		Kind:        pb.PostReportBundle_Default,
		Description: params.Description,
		Body:        params.Body,
		Receivers:   params.Receivers,
		ClubID:      params.ClubID,
	})
	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, D{
		"ID": reportResp.Report.ID,
	})
	return
}

// PostTemplateReport 发布模板汇报
func (r *Reports) PostTemplateReport(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	templateID := req.PathParameter("id")
	params := new(dto_report.ReportTemplateBundle)
	if err := r.BindAndValidate(req, params); err != nil {
		WriteError(rsp, err)
		return
	}

	reportResp, err := r.reportService.PostReport(ctx, &pb.PostReportBundle{
		Kind:       pb.PostReportBundle_Template,
		TemplateID: templateID,
		Results:    hptypes.EncodeToStruct(params.Results),
		ClubID:     params.ClubID,
		UserID:     r.GetUserIDFromRequest(req),
	})

	fmt.Println(r.GetUserIDFromRequest(req))

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, D{
		"ID": reportResp.Report.ID,
	})
	return
}

// Reports 获取汇报列表
func (r *Reports) Reports(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		ClubID string `json:"club_id,omitempty" query:"club_id"`
		Page   int32  `json:"page,omitempty" query:"page"`
		Limit  int32  `json:"limit,omitempty" query:"limit"`
	}{}

	if err := r.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	reportsResp, err := r.reportService.FindReports(ctx, &pb.FindReportsRequest{
		Page:   params.Page,
		Limit:  params.Limit,
		ClubID: params.ClubID,
		UserID: r.GetUserIDFromRequest(req),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reportsResp)
	return
}

// PendingReportTemplates 获取待汇报的模板列表
func (r *Reports) PendingReportTemplates(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Page   int32  `json:"page,omitempty" query:"page"`
		Limit  int32  `json:"limit,omitempty" query:"limit"`
		ClubID string `json:"club_id,omitempty" query:"club_id" validate:"nonzero,objectid"`
	}{}

	if err := r.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	templatesResp, err := r.reportService.FindPendingTemplates(ctx, &pb.FindTemplatesRequest{
		UserID: r.GetUserIDFromRequest(req),
		Page:   params.Page,
		Limit:  params.Limit,
		ClubID: params.ClubID,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, templatesResp)
	return
}

// GetReportTemplates 获取汇报模板列表
func (r *Reports) GetReportTemplates(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Page   int32  `json:"page,omitempty" query:"page"`
		Limit  int32  `json:"limit,omitempty" query:"limit"`
		ClubID string `json:"club_id,omitempty" query:"club_id" validate:"nonzero,objectid"`
	}{}

	if err := r.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	templatesResp, err := r.reportService.FindTemplates(ctx, &pb.FindTemplatesRequest{
		UserID: r.GetUserIDFromRequest(req),
		Page:   params.Page,
		Limit:  params.Limit,
		ClubID: params.ClubID,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, templatesResp)
	return
}

// DownloadReport 下载一个汇报
func (r *Reports) DownloadReport(req *restful.Request, rsp *restful.Response) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Write(8, "This line doesn't belong to any layer.\n")

	// Define layers
	l1 := pdf.AddLayer("Layer 1", true)
	l2 := pdf.AddLayer("Layer 2", true)

	// Open layer pane in PDF viewer
	pdf.OpenLayerPane()

	// First layer
	pdf.BeginLayer(l1)
	pdf.Write(8, "This line belongs to layer 1.\n")
	pdf.EndLayer()

	// Second layer
	pdf.BeginLayer(l2)
	pdf.Write(8, "This line belongs to layer 2.\n")
	pdf.EndLayer()

	// First layer again
	pdf.BeginLayer(l1)
	pdf.Write(8, "This line belongs to layer 1 again.\n")
	pdf.EndLayer()

	// rsp.AddHeader()
	pdf.Output(rsp.ResponseWriter)
}

// BatchDownload 批量下载汇报列表
func (r *Reports) BatchDownload(req *restful.Request, rsp *restful.Response) {

}

// Details 获取一个汇报详情
func (r *Reports) Details(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	ID := req.PathParameter("id")
	// params := struct {
	// 	ID string `json:"id,omitempty" query:"id" validate:"nonzero,objectid"`
	// }{}

	// if err := r.BindAndValidate(req, &params); err != nil {
	// 	WriteError(rsp, err)
	// 	return
	// }

	reportResp, err := r.reportService.FindOneReport(ctx, &pb.ByIDRequest{
		ID: ID,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reportResp)
	return
}

// GetReportTemplate 获取模板汇报详情
func (r *Reports) GetReportTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	ID := req.PathParameter("id")
	// params := struct {
	// 	ID string `json:"id,omitempty" query:"id" validate:"nonzero,objectid"`
	// }{}

	// if err := r.BindAndValidate(req, &params); err != nil {
	// 	WriteError(rsp, err)
	// 	return
	// }

	reportResp, err := r.reportService.FindOneTemplate(ctx, &pb.ByTemplateIDRequest{
		ID: ID,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reportResp)
	return
}

// PostTemplate 发布模版
func (r *Reports) PostTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := new(dto_report.TemplateBundle)

	if err := r.BindAndValidate(req, params); err != nil {
		WriteError(rsp, err)
		return
	}

	pbReq := &pb.PostTemplateBundle{
		Title:       params.Title,
		Description: params.Description,
		Receivers:   params.Receivers,
		ClubID:      params.ClubID,
		UserID:      r.GetUserIDFromRequest(req),
	}

	customFields := make([]*pb.TemplateCustomFieldPB, 0, len(params.CustomFields))
	for _, v := range params.CustomFields {
		customFields = append(customFields, &pb.TemplateCustomFieldPB{
			Key:     v.Key,
			Kind:    v.Kind,
			Label:   v.Label,
			Options: hptypes.EncodeToStruct(v.Options),
			Sort:    v.Sort,
		})
	}
	pbReq.CustomFields = customFields
	c := params.Config
	pbReq.Config = &pb.TemplateConfigPB{
		Kind:      c.Kind,
		Weeks:     c.Weeks,
		StartTime: hptypes.TimestampProto(c.StartTime),
		EndTime:   hptypes.TimestampProto(c.EndTime),
	}
	templateResp, err := r.reportService.PostTemplate(ctx, pbReq)

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, D{
		"ID": templateResp.ReportTemplate.ID,
	})
	return
}

// UpdateTemplate 修改模板
func (r *Reports) UpdateTemplate(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()

	// params := new(dto_report.TemplateBundle)
	// if err := r.BindAndValidate(req, params); err != nil {
	// 	WriteError(rsp, err)
	// 	return
	// }

	fields := make(map[string]interface{})

	if err := json.NewDecoder(req.Request.Body).Decode(&fields); err != nil {
		WriteError(rsp, r.Error().InternalServerError(err.Error()))
		return
	}
	// if b, err := json.Marshal(params); err == nil {
	// 	json.Unmarshal(b, fields)
	// }
	id := req.PathParameter("id")

	reply, err := r.reportService.UpdateTemplate(ctx, &pb.UpdateTemplateBundle{
		ID:     id,
		Fields: hptypes.EncodeToStruct(fields),
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
	return
}

// DisableTemplate 禁用一个模板
func (r *Reports) DisableTemplate(req *restful.Request, rsp *restful.Response) {
	id := req.PathParameter("id")
	ctx := context.Background()
	reply, err := r.reportService.ToggleTemplateEnableState(ctx, &pb.ToggleTemplateEnableReq{
		ID:        id,
		IsEnabled: false,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
	return
}

// EnableTemplate 启用一个模板
func (r *Reports) EnableTemplate(req *restful.Request, rsp *restful.Response) {
	id := req.PathParameter("id")
	ctx := context.Background()
	reply, err := r.reportService.ToggleTemplateEnableState(ctx, &pb.ToggleTemplateEnableReq{
		ID:        id,
		IsEnabled: true,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
	return
}

// DeleteTemplate 删除模板
func (r *Reports) DeleteTemplate(req *restful.Request, rsp *restful.Response) {
	id := req.PathParameter("id")
	ctx := context.Background()
	reply, err := r.reportService.DeleteTemplate(ctx, &pb.DeleteTemplateRequest{
		ID: id,
	})

	if err != nil {
		WriteError(rsp, err)
		return
	}

	WriteJsonResponse(rsp, reply)
	return
}

// func (r *Reports) Download(req *restful.)
