package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	user "iunite.club/srv/user-srv/proto/user"
)

/*
UserServiceFromContext is to get kit.iron.srv.user.UserService
*/
func UserServiceFromContext(ctx context.Context) (user.UserSrvService, bool) {
	c, ok := ctx.Value(userKey).(user.UserSrvService)

	return c, ok
}

func UserWrapper(service micro.Service) server.HandlerWrapper {
	client := user.NewUserSrvService("iron.kit.srv.user", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, userKey, client)

			return fn(ctx, req, rsp)
		}
	}
}
