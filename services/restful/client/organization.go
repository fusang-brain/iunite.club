package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	club "iunite.club/services/organization/proto/club"
	dept "iunite.club/services/organization/proto/department"
	job "iunite.club/services/organization/proto/job"
	"iunite.club/services/organization/proto/school"
)

type schoolKey struct{}
type clubKey struct{}
type deptKey struct{}
type jobKey struct{}

func SchoolServiceFromContext(ctx context.Context) (iunite_club_srv_organization_school.SchoolSrvService, bool) {
	c, ok := ctx.Value(schoolKey{}).(iunite_club_srv_organization_school.SchoolSrvService)

	return c, ok
}

func ClubServiceFromContext(ctx context.Context) (club.ClubService, bool) {
	s, ok := ctx.Value(clubKey{}).(club.ClubService)

	return s, ok
}

func DepartmentServiceFromContext(ctx context.Context) (dept.DepartmentService, bool) {
	s, ok := ctx.Value(deptKey{}).(dept.DepartmentService)

	return s, ok
}

func JobServiceFromContext(ctx context.Context) (job.JobService, bool) {
	s, ok := ctx.Value(jobKey{}).(job.JobService)

	return s, ok
}

func OrganizationServiceWrapper(service micro.Service) server.HandlerWrapper {
	serviceName := "iunite.club.srv.organization"
	sclient := service.Client()
	client := iunite_club_srv_organization_school.NewSchoolSrvService(serviceName, service.Client())
	jobClient := job.NewJobService(serviceName, service.Client())
	departmentClient := dept.NewDepartmentService(serviceName, sclient)
	clubClient := club.NewClubService(serviceName, sclient)
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, schoolKey{}, client)
			ctx = context.WithValue(ctx, jobKey{}, jobClient)
			ctx = context.WithValue(ctx, deptKey{}, departmentClient)
			ctx = context.WithValue(ctx, clubKey{}, clubClient)
			return fn(ctx, req, rsp)
		}
	}
}
