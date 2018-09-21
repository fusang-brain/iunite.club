package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"iunite.club/services/user/proto/secruity"
)

type secruityAuthKey struct{}

/*
SecruityAuthServiceFromContext 从上下文中获取安全服务
*/
func SecruityAuthServiceFromContext(ctx context.Context) (iunite_club_srv_user_secruity.SecruityService, bool) {
	c, ok := ctx.Value(secruityAuthKey{}).(iunite_club_srv_user_secruity.SecruityService)

	return c, ok
}

/*
SecruityWrapper 安全服务注入中间件
*/
func SecruityWrapper(service micro.Service) server.HandlerWrapper {
	client := iunite_club_srv_user_secruity.NewSecruityService("iunite.club.srv.user", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, resp interface{}) error {
			ctx = context.WithValue(ctx, secruityAuthKey{}, client)

			return fn(ctx, req, resp)
		}
	}
}
