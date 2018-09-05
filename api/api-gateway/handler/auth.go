package handler

import (
	"context"
	"fmt"
	ironic_api "github.com/iron-kit/go-ironic/api"
	"github.com/iron-kit/go-ironic/micro-assistant"
	go_api "github.com/micro/go-api/proto"
)

type Auth struct {
	ironic_api.Helper
	Error *assistant.ErrorManager
}

func (a *Auth) Login(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {
	token := a.GetTokenFromRequest(req)
	fmt.Println(token)
	resp.StatusCode = 200
	resp.Body = "{}"

	return nil
}

func (a *Auth) Register(ctx context.Context, req *go_api.Request, resp *go_api.Response) error {
	header := req.GetHeader()

	fmt.Println(header)

	resp.StatusCode = 200
	resp.Body = "{}"

	return nil
}
