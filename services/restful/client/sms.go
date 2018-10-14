package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	sms "iunite.club/services/message/proto/sms"
)

type smsServiceKey struct{}

func SMSServiceFromContext(ctx context.Context) (sms.SMSService, bool) {
	v, ok := ctx.Value(smsServiceKey{}).(sms.SMSService)
	return v, ok
}

func MessageServiceWrapper(service micro.Service) server.HandlerWrapper {
	name := "iunite.club.srv.message"
	smsSrv := sms.NewSMSService(name, service.Client())
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, smsServiceKey{}, smsSrv)
			return fn(ctx, req, rsp)
		}
	}
}
