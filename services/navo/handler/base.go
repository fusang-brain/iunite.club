package handler

import (
	"strings"

	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/api"
	go_api "github.com/micro/go-api/proto"
	"iunite.club/services/user/utils"
)

const (
	PlatformWEB     = "WEB"
	PlatformAndroid = "ANDROID"
	PlatformIOS     = "IOS"
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

func (h *BaseHandler) GetUserIDFromRequest(req *go_api.Request) string {
	token, err := h.GetTokenCliamsFromRequest(req)

	if err != nil {
		return ""
	}

	return token.UserID
}

func (h *BaseHandler) GetPlatformFromRequest(req *go_api.Request) string {

	platformFlag := h.GetHeaderFieldFromRequest(req, "U-Platform")

	if platformFlag == "" || strings.ToUpper(platformFlag) == PlatformWEB {
		return PlatformWEB
	}

	if strings.ToUpper(platformFlag) == PlatformAndroid {
		return PlatformAndroid
	}

	if strings.ToUpper(platformFlag) == PlatformIOS {
		return PlatformIOS
	}

	return PlatformWEB
}
