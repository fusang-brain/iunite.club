package srv

import (
	"context"
	"github.com/iron-kit/go-ironic"
)

func NewJobService(ctx context.Context) *JobService {
	jobService := &JobService{}

	if err := ironic.InitServiceFunc(jobService, ctx); err != nil {
		panic(err.Error())
	}

	return jobService
}

func NewDepartmentService(ctx context.Context) *DepartmentService {
	departmentService := &DepartmentService{}
	if err := ironic.InitServiceFunc(departmentService, ctx); err != nil {
		panic(err.Error())
	}

	return departmentService
}

func NewSchoolService(ctx context.Context) *SchoolService {
	schoolService := &SchoolService{}

	if err := ironic.InitServiceFunc(schoolService, ctx); err != nil {
		panic(err.Error())
	}

	return schoolService
}

func NewClubService(ctx context.Context) *ClubService {
	clubService := &ClubService{}

	if err := ironic.InitServiceFunc(clubService, ctx); err != nil {
		panic(err.Error())
	}

	return clubService
}
