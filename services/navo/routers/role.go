package routers

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/navo/handler"
	"iunite.club/services/navo/router"
)

func RoleRoute(r *router.Router) {
	roleHandler := handler.NewRoleHandler(client.DefaultClient)

	roles := r.Group(
		"Role",
		"/roles",
		router.Description("角色"),
		router.Produces(restful.MIME_JSON),
	)

	roles.GET(
		"/CheckTest",
		roleHandler.CheckTest,
		router.RouteDoc("测试权限"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteReads(struct {
			Role string `json:"role"`
			Path string `json:"path"`
			Method string `json:"method"`
		}{}),
	)

	roles.POST(
		"/",
		roleHandler.CreateRole,
		router.RouteDoc("创建角色"),
		router.RouteParam(getAuthHeaderParam(r)),
		router.RouteParams(
			roles.WS().QueryParameter("role", "角色").DataType("string"),
			roles.WS().QueryParameter("path", "路径").DataType("string"),
			roles.WS().QueryParameter("method", "请求方式").DataType("string"),
		),
	)
}