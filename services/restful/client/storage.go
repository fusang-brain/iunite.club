package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	storagePB "iunite.club/services/storage/proto"
)

type storageKey struct{}

func StroageServiceFromContext(ctx context.Context) (storagePB.StorageService, bool) {
	c, ok := ctx.Value(storageKey{}).(storagePB.StorageService)

	return c, ok
}

func StorageServiceWrapper(service micro.Service) server.HandlerWrapper {
	serviceName := "iunite.club.srv.storage"

	sclient := service.Client()
	// approvedService := approvedPB.NewApprovedService(serviceName, sclient)
	storageSrv := storagePB.NewStorageService(serviceName, sclient)
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, storageKey{}, storageSrv)

			return fn(ctx, req, rsp)
		}
	}
}
