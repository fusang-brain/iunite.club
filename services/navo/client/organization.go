package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"iunite.club/services/organization/proto/school"
)

type schoolKey struct{}

func SchoolServiceFromContext(ctx context.Context) (iunite_club_srv_organization_school.SchoolSrvService, bool) {
	c, ok := ctx.Value(schoolKey{}).(iunite_club_srv_organization_school.SchoolSrvService)

	return c, ok
}

func OrganizationServiceWrapper(service micro.Service) server.HandlerWrapper {
	serviceName := "iunite.club.srv.organization"
	client := iunite_club_srv_organization_school.NewSchoolSrvService(serviceName, service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, schoolKey{}, client)

			return fn(ctx, req, rsp)
		}
	}
}
