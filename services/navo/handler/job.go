package handler

import (
	"context"

	userPB "iunite.club/services/user/proto"

	"iunite.club/services/navo/dto"

	"iunite.club/services/navo/client"

	jobPB "iunite.club/services/organization/proto/job"

	go_api "github.com/micro/go-api/proto"
)

type JobHandler struct {
	BaseHandler
	jobService jobPB.JobService
}

func (j *JobHandler) getJobService(ctx context.Context) jobPB.JobService {

	if j.jobService == nil {
		jobSrv, ok := client.JobServiceFromContext(ctx)

		if !ok {
			panic("Not found job service")
		}

		j.jobService = jobSrv
	}

	return j.jobService
}

func (j *JobHandler) CreateJob(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	jobSrv := j.getJobService(ctx)
	params := struct {
		JobName string `json:"jobName"`
		ClubID  string `json:"org" validate:"nonzero,objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	createdRsp, err := jobSrv.CreateJob(ctx, &jobPB.CreateJobRequest{
		ClubID: params.ClubID,
		Name:   params.JobName,
	})
	if err != nil {
		return ErrorResponse(rsp, err)
	}

	if !createdRsp.OK {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest("CreatedError"))
	}

	return SuccessResponse(rsp, D{
		"Job": dto.PBToJob(createdRsp.Job),
	})

}

func (j *JobHandler) GetUsersWithJob(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Page  int64  `query:"page"`
		Limit int64  `query:"limit"`
		JobID string `json:"jobID" query:"jobID" validate:"nonzero,objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	jobSrv, ok := client.JobServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest("job service not found"))
	}

	usersResp, err := jobSrv.GetUsersByJobID(ctx, &jobPB.ListByJobIDRequest{Page: params.Page, Limit: params.Limit, JobID: params.JobID})

	if err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	// usersResp.Total
	users := make([]*dto.User, 0)

	for _, u := range usersResp.Users {
		users = append(users, dto.PBToUser(u))
	}

	return SuccessResponse(rsp, D{
		"CurrentPage": params.Page,
		"PageSize":    params.Limit,
		"PageTotal":   usersResp.Total,
		"Total":       usersResp.Total,
		"List":        users,
	})
}

func (j *JobHandler) AllCanSelectedUsers(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {

	params := struct {
		Page  int64
		Limit int64
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	currentUserID := j.GetUserIDFromRequest(req)

	userSrv, _ := client.UserServiceFromContext(ctx)

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: currentUserID})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	jobService := j.getJobService(ctx)

	foundUsersResp, err := jobService.GetAllCanSelectedUsers(ctx, &jobPB.ListByClubIDRequest{Page: params.Page, Limit: params.Limit, ClubID: foundUserResp.User.DefaultClubID})
	if err != nil {
		return ErrorResponse(rsp, err)
	}
	users := make([]*dto.User, 0)
	for _, u := range foundUsersResp.Users {
		users = append(users, dto.PBToUser(u))
	}

	return SuccessResponseWithPage(rsp, params.Page, params.Limit, foundUsersResp.Total, users)
}

func (j *JobHandler) AddUsersToJob(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Users []string `json:"users,omitempty"`
		JobID string   `json:"job_id,omitempty" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	userSrv, _ := client.UserServiceFromContext(ctx)
	currentUserID := j.GetUserIDFromRequest(req)

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: currentUserID})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	jobService := j.getJobService(ctx)

	resp, err := jobService.AddUsersToJob(
		ctx,
		&jobPB.UserFromJobRequest{
			Users:  params.Users,
			JobID:  params.JobID,
			ClubID: foundUserResp.User.DefaultClubID,
		})

	if err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if !resp.OK {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest("UpdateError"))
	}

	return SuccessResponse(rsp, D{})
}

func (j *JobHandler) RemoveUsersFromJob(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	params := struct {
		Users []string `json:"users,omitempty"`
		JobID string   `json:"job_id,omitempty" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	userSrv, _ := client.UserServiceFromContext(ctx)
	currentUserID := j.GetUserIDFromRequest(req)

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: currentUserID})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	jobService := j.getJobService(ctx)

	resp, err := jobService.RemoveUsersFromJob(
		ctx,
		&jobPB.UserFromJobRequest{
			Users:  params.Users,
			JobID:  params.JobID,
			ClubID: foundUserResp.User.DefaultClubID,
		})

	if err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if !resp.OK {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest("UpdateError"))
	}

	return SuccessResponse(rsp, D{})
}

func (j *JobHandler) GetAllJobs(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	params := struct {
		ID string `query:"id" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if params.ID == "" {
		params.ID = j.GetCurrentClubIDFromRequest(ctx, req)
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	jobService := j.getJobService(ctx)
	foundResp, err := jobService.GetJobListByParentID(ctx, &jobPB.JobListRequest{Page: 1, Limit: 999, OrganizationID: params.ID})
	if err != nil {
		return ErrorResponse(rsp, err)
	}
	jobs := make([]*dto.Job, 0)

	for _, job := range foundResp.Jobs {
		jobs = append(jobs, dto.PBToJob(job))
	}

	return SuccessResponse(rsp, D{
		"Jobs": jobs,
	})
}

func (j *JobHandler) Update(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	params := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	jobService := j.getJobService(ctx)

	currentClubID := j.GetCurrentClubIDFromRequest(ctx, req)
	_, err := jobService.UpdateJob(ctx, &jobPB.UpdateJobRequest{
		ID:     params.ID,
		ClubID: currentClubID,
		Name:   params.Name,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{})
}

func (j *JobHandler) Remove(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	params := struct {
		ID string `json:"id"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	if err := j.Validate(&params); err != nil {
		return ErrorResponse(rsp, j.Error(ctx).BadRequest(err.Error()))
	}

	jobService := j.getJobService(ctx)
	_, err := jobService.RemoveJob(ctx, &jobPB.RemoveJobRequest{ID: params.ID})
	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{})
}
