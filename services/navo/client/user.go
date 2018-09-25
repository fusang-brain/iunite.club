package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"iunite.club/services/user/proto"
	"iunite.club/services/user/proto/secruity"
)

type userKey struct{}

type secruityAuthKey struct{}

/*
SecruityAuthServiceFromContext 从上下文中获取安全服务
*/
func SecruityAuthServiceFromContext(ctx context.Context) (iunite_club_srv_user_secruity.SecruityService, bool) {
	c, ok := ctx.Value(secruityAuthKey{}).(iunite_club_srv_user_secruity.SecruityService)

	return c, ok
}

func UserServiceFromContext(ctx context.Context) (iunite_club_srv_user.UserSrvService, bool) {
	c, ok := ctx.Value(userKey{}).(iunite_club_srv_user.UserSrvService)

	return c, ok
}

func UserServiceWrapper(service micro.Service) server.HandlerWrapper {
	client := iunite_club_srv_user.NewUserSrvService("iunite.club.srv.user", service.Client())
	secruityClient := iunite_club_srv_user_secruity.NewSecruityService("iunite.club.srv.user", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, userKey{}, client)
			ctx = context.WithValue(ctx, secruityAuthKey{}, secruityClient)
			return fn(ctx, req, rsp)
		}
	}
}
