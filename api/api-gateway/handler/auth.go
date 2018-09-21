package handler

import (
	"context"
	"encoding/json"
	"fmt"
	ironic_api "github.com/iron-kit/go-ironic/api"
	"github.com/iron-kit/go-ironic/micro-assistant"
	authPB "iunite.club/srv/secruity/proto/auth"

	go_api "github.com/micro/go-api/proto"
	"iunite.club/api/api-gateway/client"
)

type Auth struct {
	// assistant.BaseHandler
	ironic_api.Helper
	Error *assistant.ErrorManager
}

func (a *Auth) Login(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {
	authService, ok := client.SecruityAuthServiceFromContext(ctx)
	if !ok {
		return a.Error.InternalServerError("not found AuthService")
	}

	authReq := authPB.AuthRequest{}
	err := json.Unmarshal([]byte(req.Body), &authReq)

	if err != nil {
		return a.Error.InternalServerError("Params parse error" + err.Error())
	}

	authResp, err := authService.Signin(ctx, &authReq)

	if err != nil {
		return err
	}

	b, err := json.Marshal(authResp)
	if err != nil {

		return a.Error.InternalServerError(err.Error())
	}

	resp.StatusCode = 200
	resp.Body = string(b)
	// resp.Body = "{}"
	return nil
}

func (a *Auth) Register(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {
	header := req.GetHeader()

	fmt.Println(header)

	resp.StatusCode = 200
	resp.Body = "{}"

	return nil
}
