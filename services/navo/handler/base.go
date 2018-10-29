package handler

import (
	"context"
	"strings"

	"github.com/micro/go-micro/client"

	restful "github.com/emicklei/go-restful"

	userPB "iunite.club/services/user/proto"

	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/api"
	"iunite.club/services/user/utils"
)

const (
	PlatformWEB     = "WEB"
	PlatformAndroid = "ANDROID"
	PlatformIOS     = "IOS"
)

const (
	UserSerivce         = "iunite.club.srv.user"
	OrganizationService = "iunite.club.srv.organization"
	SMSService          = "iunite.club.srv.message"
	StorageService      = "iunite.club.srv.storage"
	CoreService         = "iunite.club.srv.core"
	ReportService 			= "iunite.club.srv.report"
	ApprovedService     = "iunite.club.srv.approved"
)

type BaseHandler struct {
	ironic.BaseWebHandler
	api.Helper
	userClient userPB.UserSrvService
}

func (h *BaseHandler) getTokenFromRequest(req *restful.Request) string {

	// header := r.GetHeader()

	// fmt.Println(header)
	tokenString := req.HeaderParameter("U-Access-Token")

	if tokenString == "" {
		t := req.HeaderParameter("Authorization")
		tokenString = strings.TrimSpace(strings.Replace(t, "Bearer", "", 1))
	}

	return tokenString
}

func (h *BaseHandler) GetTokenCliamsFromRequest(req *restful.Request) (*utils.CustomClaims, error) {
	token := h.getTokenFromRequest(req)

	log.Logf("token is %v >> ", token)
	tokenSrv := utils.TokenService{}

	return tokenSrv.Decode(token)
}

func (h *BaseHandler) GetUserIDFromRequest(req *restful.Request) string {
	token, err := h.GetTokenCliamsFromRequest(req)

	if err != nil {
		return ""
	}

	return token.UserID
}

func (h *BaseHandler) GetCurrentClubIDFromRequest(req *restful.Request) string {
	currentUserID := h.GetUserIDFromRequest(req)
	// if h.userClient
	// userSrv, _ := client.UserServiceFromContext(ctx)
	if h.userClient == nil {
		h.userClient = userPB.NewUserSrvService("iunite.club.srv.user", client.DefaultClient)
	}

	resp, err := h.userClient.FindUserByID(context.Background(), &userPB.QueryUserRequest{Id: currentUserID})
	if err != nil {
		panic(err)
	}

	return resp.User.DefaultClubID
}

func (h *BaseHandler) GetPlatformFromRequest(req *restful.Request) string {

	platformFlag := req.HeaderParameter("U-Platform")

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

func (self *BaseHandler) BindAndValidate(req *restful.Request, in interface{}) error {
	if err := self.Bind(req, in); err != nil {
		return self.Error().BadRequest(err.Error())
	}

	if err := self.Validate(in); err != nil {
		return self.Error().BadRequest(err.Error())
	}

	return nil
}
