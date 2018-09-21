package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"iunite.club/srv/secruity/proto/auth"
)

type secruityAuthKey struct{}

/*
SecruityAuthServiceFromContext 从上下文中获取安全服务
*/
func SecruityAuthServiceFromContext(ctx context.Context) (kit_iron_srv_secruity.AuthService, bool) {
	c, ok := ctx.Value(secruityAuthKey{}).(kit_iron_srv_secruity.AuthService)

	return c, ok
}

/*
SecruityWrapper 安全服务注入中间件
*/
func SecruityWrapper(service micro.Service) server.HandlerWrapper {
	client := kit_iron_srv_secruity.NewAuthService("kit.iron.srv.secruity", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, resp interface{}) error {
			ctx = context.WithValue(ctx, secruityAuthKey{}, client)

			return fn(ctx, req, resp)
		}
	}
}
