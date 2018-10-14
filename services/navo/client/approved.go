package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	approvedPB "iunite.club/services/core/proto/approved"
)

type approvedKey struct{}

func ApprovedServiceFromContext(ctx context.Context) (approvedPB.ApprovedService, bool) {
	c, ok := ctx.Value(approvedKey{}).(approvedPB.ApprovedService)

	return c, ok
}

func CoreServiceWrapper(service micro.Service) server.HandlerWrapper {
	serviceName := "iunite.club.srv.core"

	sclient := service.Client()
	approvedService := approvedPB.NewApprovedService(serviceName, sclient)
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, approvedKey{}, approvedService)

			return fn(ctx, req, rsp)
		}
	}
}
