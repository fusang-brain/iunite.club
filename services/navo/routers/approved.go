package routers

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/navo/dto/approved"
	"iunite.club/services/navo/handler"
	"iunite.club/services/navo/router"
)

func ApprovedRoute(r *router.Router) {
	approvedHandler := handler.NewApprovedHandler(client.DefaultClient)

	approved := r.Group(
		"Approved",
		"/approved",
		router.Description("审批模块"),
		router.Produces(restful.MIME_JSON),
	)

	authHeaderParam := getAuthHeaderParam(r)
	paginationParams := getPaganizationParams(approved)

	approved.GET(
		"/templates",
		approvedHandler.FindTemplates,
		router.RouteDoc("获取审批模板列表"),
		router.RouteParam(authHeaderParam),
		router.RouteParams(paginationParams...),
		router.RouteParam(approved.WS().QueryParameter("club_id", "社团ID")),
		router.RouteParam(approved.WS().QueryParameter("populates", "需要展开的字段，多个字段以英文逗号分隔")),
	)

	approved.POST(
		"/templates",
		approvedHandler.PostTemplate,
		router.RouteDoc("发布一个模板"),
		router.RouteParam(authHeaderParam),
		router.RouteReads(dto_approved.TemplateBundle{}),
	)

	approved.PUT(
		"/templates/{id}",
		approvedHandler.UpdateTemplate,
		router.RouteDoc("修改一个模板"),
		router.RouteParam(authHeaderParam),
		router.RouteParam(approved.WS().PathParameter("id", "模板ID")),
		router.RouteReads(dto_approved.UpdateTemplateBundle{}),
	)

	approved.PUT(
		"/templates/{id}/enable",
		approvedHandler.EnableTemplate,
		router.RouteDoc("启用一个模板"),
		router.RouteParam(authHeaderParam),
		router.RouteParam(approved.WS().PathParameter("id", "模板ID")),
	)

	approved.PUT(
		"/templates/{id}/disable",
		approvedHandler.DisableTemplate,
		router.RouteDoc("停用一个模板"),
		router.RouteParam(authHeaderParam),
		router.RouteParam(approved.WS().PathParameter("id", "模板ID")),
	)

	approved.DELETE(
		"/templates/{id}",
		approvedHandler.DeleteTemplates,
		router.RouteDoc("删除一个模板"),
		router.RouteParam(authHeaderParam),
		router.RouteParam(approved.WS().PathParameter("id", "模板ID")),
	)
}
