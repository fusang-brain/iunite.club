package wrappers

import (
	"context"
	"fmt"
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"os"
)

type AuthWhiteListItemer interface {
	Service() string
	Method() string
}

type authWhiteListItem struct {
	service string
	method  string
}

func (s *authWhiteListItem) Service() string {
	return s.service
}

func (s *authWhiteListItem) Method() string {
	return s.method
}

func NewWhiteItem(service, method string) AuthWhiteListItemer {
	return &authWhiteListItem{service, method}
}

func GenerateAuthWrapper(whiteList ...AuthWhiteListItemer) server.HandlerWrapper {
	errManager := assistant.NewErrorManager("kit.iron.srv.secruity")
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, resp interface{}) error {
			if os.Getenv("DISABLE_AUTH") == "true" {
				return fn(ctx, req, resp)
			}
			// exec white list
			for _, item := range whiteList {
				// if not method, the all serivce all needn't auth toke
				fmt.Println(item.Service(), item.Method())
				fmt.Println(req.Service(), req.Method())
				if item.Method() == "" {
					if item.Service() == req.Service() {
						return fn(ctx, req, resp)
					}
				} else {
					if item.Service() == req.Service() && item.Method() == req.Method() {
						return fn(ctx, req, resp)
					}
				}
				// if item.Service() == req.Service()
			}

			meta, ok := metadata.FromContext(ctx)
			if !ok {
				return errManager.Unauthorized("no auth meta-data found in request")
			}

			// Note this is now uppercase (not entirely sure why this is...)
			token := meta["Token"]

			if token == "" {
				return errManager.Unauthorized("token couldn't be null")
			}

			// Auth here
			// authClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
			// authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			// 	Token: token,
			// })
			// log.Println("Auth Resp:", authResp)

			// if err != nil {
			// 	return err
			// }
			err := fn(ctx, req, resp)
			return err
		}
	}
}
