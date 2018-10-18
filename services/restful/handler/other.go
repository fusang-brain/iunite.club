package handler

import restful "github.com/emicklei/go-restful"

type OtherHandler struct {
	BaseHandler
}

func (self *OtherHandler) GetUnreadCount(req *restful.Request, rsp *restful.Response) {
	// TODO GetUnreadCount
	SuccessResponse(rsp, D{
		"NotifyCount": 0,
	})
}

func (self *OtherHandler) GetAllCanUseRoles(req *restful.Request, rsp *restful.Response) {
	// TODO GetAllCanUseRoles
	SuccessResponse(rsp, D{
		"Roles": []string{},
	})
}
