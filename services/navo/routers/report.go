package routers

import (
	"iunite.club/services/navo/dto/report"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
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
	)

	reports.GET(
		"/templates/pending",
		reportsHandler.PendingReportTemplates,
		router.RouteDoc("获取待我汇报的模板列表"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParams(paganizationParams...),
	)

	reports.POST(
		"/",
		reportsHandler.PostReport,
		router.RouteDoc("发布汇报"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteReads(dto_report.SimpleReportBundle{}),
	)

	reports.POST(
		"/template/{id}",
		reportsHandler.PostReport,
		router.RouteDoc("发布模板汇报"),
		router.RouteReads(dto_report.ReportTemplateBundle{}),
	)

	reports.GET(
		"/{id}/download",
		reportsHandler.DownloadReport,
		router.RouteDoc("下载汇报"),
	)

	reports.GET(
		"/batch_download/{start_time}/{end_time}",
		reportsHandler.BatchDownload,
		router.RouteDoc("批量下载"),
		router.RouteParam(reports.WS().QueryParameter("ids", "要下载的报告ID")),
	)

	reports.GET(
		"/{id}",
		reportsHandler.Details,
		router.RouteDoc("获取汇报详情"),
	)

	reports.GET(
		"/templates/{id}",
		reportsHandler.GetReportTemplate,
		router.RouteDoc("获取汇报模板详情"),
	)

	reports.POST(
		"/templates",
		reportsHandler.PostTemplate,
		router.RouteDoc("发布一个新的汇报模板"),
		router.RouteReads(dto_report.TemplateBundle{}),
	)


}
