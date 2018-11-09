package routers

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/navo/dto/report"
	"iunite.club/services/navo/handler"
	"iunite.club/services/navo/router"
)

func getPaganizationParams(route *router.Route) []*restful.Parameter {
	return []*restful.Parameter{
		route.WS().QueryParameter("page", "页数").DataType("int32"),
		route.WS().QueryParameter("limit", "限定").DataType("int32"),
	}
}

func ReportRoute(r *router.Router) {
	reportsHandler := handler.NewReportsHandler(client.DefaultClient)

	reports := r.Group(
		"Report",
		"/reports",
		router.Description("汇报"),
		router.Produces(restful.MIME_JSON),
	)
	paganizationParams := getPaganizationParams(reports)

	reports.GET(
		"/",
		reportsHandler.Reports,
		router.RouteDoc("获取汇报列表"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParams(paganizationParams...),
		router.RouteParam(reports.WS().QueryParameter("club_id", "社团ID")),
	)

	reports.GET(
		"/templates/pending",
		reportsHandler.PendingReportTemplates,
		router.RouteDoc("获取待我汇报的模板列表"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParams(paganizationParams...),
		router.RouteParam(reports.WS().QueryParameter("club_id", "社团ID")),
	)

	reports.POST(
		"/",
		reportsHandler.PostReport,
		router.RouteDoc("发布汇报"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteReads(dto_report.SimpleReportBundle{}),
	)

	reports.POST(
		"/templates/{id}",
		reportsHandler.PostTemplateReport,
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteDoc("发布模板汇报"),
		router.RouteParams(
			reports.WS().PathParameter("id", "模板ID"),
		),
		router.RouteReads(dto_report.ReportTemplateBundle{}),
	)

	reports.GET(
		"/{id}/download",
		reportsHandler.DownloadReport,
		// router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParams(
			reports.WS().PathParameter("id", "汇报ID"),
		),
		router.RouteDoc("下载汇报"),
	)

	reports.GET(
		"/batch_download/{start_time}/{end_time}",
		reportsHandler.BatchDownload,
		router.RouteDoc("批量下载"),
		// router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParams(
			reports.WS().PathParameter("start_time", "开始时间"),
			reports.WS().PathParameter("end_time", "结束时间"),
		),
		router.RouteParam(reports.WS().QueryParameter("ids", "要下载的报告ID")),
	)

	reports.GET(
		"/{id}",
		reportsHandler.Details,
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParam(reports.WS().PathParameter("id", "汇报ID")),
		router.RouteDoc("获取汇报详情"),
	)

	reports.GET(
		"/templates/{id}",
		reportsHandler.GetReportTemplate,
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParam(reports.WS().PathParameter("id", "汇报模板ID")),
		router.RouteDoc("获取汇报模板详情"),
	)

	reports.POST(
		"/templates",
		reportsHandler.PostTemplate,
		router.RouteDoc("发布一个新的汇报模板"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteReads(dto_report.TemplateBundle{}),
	)

	reports.PUT(
		"/templates/{id}",
		reportsHandler.UpdateTemplate,
		router.RouteDoc("修改一个模板"),
		router.RouteParams(getAuthHeaderParam(r)),
		router.RouteReads(dto_report.TemplateBundle{}),
		router.RouteParam(reports.WS().PathParameter("id", "汇报模板ID")),
	)

	reports.PUT(
		"/templates/{id}/disable",
		reportsHandler.DisableTemplate,
		router.RouteDoc("禁用一个模板"),
		router.RouteParams(getAuthHeaderParam(r)),
		router.RouteParam(reports.WS().PathParameter("id", "汇报模板ID")),
	)

	reports.PUT(
		"/templates/{id}/enable",
		reportsHandler.EnableTemplate,
		router.RouteDoc("启用一个模板"),
		router.RouteParams(getAuthHeaderParam(r)),
		router.RouteParam(reports.WS().PathParameter("id", "汇报模板ID")),
	)

	reports.DELETE(
		"/templates/{id}",
		reportsHandler.DeleteTemplate,
		router.RouteDoc("删除一个模板"),
		router.RouteParams(getAuthHeaderParam(r)),
		router.RouteParam(reports.WS().PathParameter("id", "汇报模板ID")),
	)
}
