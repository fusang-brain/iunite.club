package routers

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/navo/dto/announce"
	"iunite.club/services/navo/handler"
	"iunite.club/services/navo/router"
	"time"
)

func getAuthHeaderParam(r *router.Router) *restful.Parameter {
	return r.GetWS().HeaderParameter("Authorization", "密钥")
}

func AnnounceRoute(r *router.Router) {

	announceHandler := handler.NewAnnounceHandler(client.DefaultClient)
	authHeaderParam := getAuthHeaderParam(r)
	announceRoute := r.Group(
		"Announce",
		"/announces",
		router.Description("通告模块"),
		router.Produces(restful.MIME_JSON),
	)

	announceRoute.GET(
		"/",
		announceHandler.GetAnnounces,
		router.RouteDoc("获取通告列表"),
		router.RouteParam(
			announceRoute.WS().HeaderParameter("Authorization", "密钥"),
		),
		router.RouteParams(
			announceRoute.WS().QueryParameter("page", "页数").DataType("int32"),
			announceRoute.WS().QueryParameter("limit", "限定").DataType("int32"),
			announceRoute.WS().QueryParameter("club_id", "社团ID").DataType("string"),
			announceRoute.WS().
				QueryParameter("kind", "类型(instructions:社长令,task:任务,reminder:提醒)").
				DataType("string"),
		),
		router.RouteReturns(200, "ok", AnnounceListResponse{}),
	)
	announceRoute.POST(
		"/instructions",
		announceHandler.CreateInstructions,
		router.RouteDoc("发布社长令"),
		router.RouteParam(authHeaderParam),
		router.RouteReads(dto_announce.CreateInstructionsBundle{}),
		router.RouteReturns(200, "ok", struct {
			CreatedAt time.Time
			ID string
		}{}),
	)
	announceRoute.POST(
		"/task",
		announceHandler.CreateTask,
		router.RouteDoc("发布任务"),
		router.RouteParam(authHeaderParam),
		router.RouteReads(dto_announce.CreateTaskBundle{}),
		router.RouteReturns(200, "ok", struct {
			CreatedAt time.Time
			ID string
		}{}),
	)
	announceRoute.POST(
		"/reminder",
		announceHandler.CreateReminder,
		router.RouteDoc("发起提醒"),
		router.RouteParam(authHeaderParam),
		router.RouteReads(dto_announce.CreateReminderBundle{}),
		router.RouteReturns(200, "ok", struct {
			CreatedAt time.Time
			ID string
		}{}),
	)
	announceRoute.PUT(
		"/{id}/marked_to_read",
		announceHandler.MarkedToRead,
		router.RouteDoc("标记为已读"),
		router.RouteParam(authHeaderParam),
		router.RouteParam(announceRoute.WS().PathParameter("id", "通告ID").DataType("string")),
		router.RouteReturns(200, "ok", struct {
			ID string
		}{}),
	)
}
