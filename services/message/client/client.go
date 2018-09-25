package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	core "iunite.club/services/core/proto/example"
)

type coreServiceKey struct{}

func TestServerFromContext(ctx context.Context) (core.ExampleService, bool) {
	v, ok := ctx.Value(coreServiceKey{}).(core.ExampleService)

	return v, ok
}

func TestServiceWrapper(service micro.Service) server.HandlerWrapper {
	name := "iunite.club.srv.core"

	coreSrv := core.NewExampleService(name, service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, coreServiceKey{}, coreSrv)

			return fn(ctx, req, rsp)
		}
	}
}
