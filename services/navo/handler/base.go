package handler

import (
	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/api"
	go_api "github.com/micro/go-api/proto"
	"iunite.club/services/user/utils"
)

type BaseHandler struct {
	ironic.BaseHandler
	api.Helper
}

func (h *BaseHandler) GetTokenCliamsFromRequest(req *go_api.Request) (*utils.CustomClaims, error) {
	token := h.GetTokenFromRequest(req)
	log.Logf("token is %v >> ", token)
	tokenSrv := utils.TokenService{}

	return tokenSrv.Decode(token)
}
