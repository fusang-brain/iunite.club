package handler

import (
	"context"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"
	"iunite.club/services/restful/dto"
	userPB "iunite.club/services/user/proto"

	jobPB "iunite.club/services/organization/proto/job"
)

type JobHandler struct {
	BaseHandler
	jobService  jobPB.JobService
	userService userPB.UserSrvService
}

func NewJobService(c client.Client) *JobHandler {
	return &JobHandler{
		jobService:  jobPB.NewJobService(OrganizationService, c),
		userService: userPB.NewUserSrvService(UserSerivce, c),
	}
}

func (j *JobHandler) getJobService() jobPB.JobService {
	return j.jobService
}

func (j *JobHandler) CreateJob(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	ctx := context.Background()
	jobSrv := j.getJobService()
	params := struct {
		JobName string `json:"jobName"`
		ClubID  string `json:"org" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if params.ClubID == "" {
		params.ClubID = j.GetCurrentClubIDFromRequest(req)
	}

	createdRsp, err := jobSrv.CreateJob(ctx, &jobPB.CreateJobRequest{
		ClubID: params.ClubID,
		Name:   params.JobName,
	})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	if !createdRsp.OK {
		ErrorResponse(rsp, j.Error().BadRequest("CreatedError"))
		return
	}

	SuccessResponse(rsp, D{
		"Job": dto.PBToJob(createdRsp.Job),
	})

}

func (j *JobHandler) GetUsersWithJob(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Page  int64  `query:"page"`
		Limit int64  `query:"limit"`
		JobID string `json:"jobID" query:"jobID" validate:"nonzero,objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	jobSrv := j.jobService

	usersResp, err := jobSrv.GetUsersByJobID(ctx, &jobPB.ListByJobIDRequest{Page: params.Page, Limit: params.Limit, JobID: params.JobID})

	if err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	// usersResp.Total
	users := make([]*dto.User, 0)

	for _, u := range usersResp.Users {
		users = append(users, dto.PBToUser(u))
	}

	SuccessResponse(rsp, D{
		"CurrentPage": params.Page,
		"PageSize":    params.Limit,
		"PageTotal":   usersResp.Total,
		"Total":       usersResp.Total,
		"List":        users,
	})
}

func (j *JobHandler) AllCanSelectedUsers(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Page  int64
		Limit int64
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	currentUserID := j.GetUserIDFromRequest(req)

	userSrv := j.userService

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: currentUserID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	jobService := j.getJobService()

	foundUsersResp, err := jobService.GetAllCanSelectedUsers(ctx, &jobPB.ListByClubIDRequest{Page: params.Page, Limit: params.Limit, ClubID: foundUserResp.User.DefaultClubID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	users := make([]*dto.User, 0)
	for _, u := range foundUsersResp.Users {
		users = append(users, dto.PBToUser(u))
	}

	SuccessResponseWithPage(rsp, params.Page, params.Limit, foundUsersResp.Total, users)
}

func (j *JobHandler) AddUsersToJob(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Users []string `json:"users,omitempty"`
		JobID string   `json:"jobID,omitempty" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	userSrv := j.userService
	currentUserID := j.GetUserIDFromRequest(req)

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: currentUserID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	jobService := j.getJobService()

	resp, err := jobService.AddUsersToJob(
		ctx,
		&jobPB.UserFromJobRequest{
			Users:  params.Users,
			JobID:  params.JobID,
			ClubID: foundUserResp.User.DefaultClubID,
		})

	if err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if !resp.OK {
		ErrorResponse(rsp, j.Error().BadRequest("UpdateError"))
		return
	}

	SuccessResponse(rsp, D{})
}

func (j *JobHandler) RemoveUsersFromJob(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	params := struct {
		Users []string `json:"users,omitempty"`
		JobID string   `json:"job_id,omitempty" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	userSrv := j.userService
	currentUserID := j.GetUserIDFromRequest(req)

	foundUserResp, err := userSrv.FindUserByID(ctx, &userPB.QueryUserRequest{Id: currentUserID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	jobService := j.getJobService()

	resp, err := jobService.RemoveUsersFromJob(
		ctx,
		&jobPB.UserFromJobRequest{
			Users:  params.Users,
			JobID:  params.JobID,
			ClubID: foundUserResp.User.DefaultClubID,
		})

	if err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if !resp.OK {
		ErrorResponse(rsp, j.Error().BadRequest("UpdateError"))
		return
	}

	SuccessResponse(rsp, D{})
}

func (j *JobHandler) GetAllJobs(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	ctx := context.Background()
	params := struct {
		ID string `query:"id" validate:"objectid"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if params.ID == "" {
		params.ID = j.GetCurrentClubIDFromRequest(req)
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	jobService := j.getJobService()
	foundResp, err := jobService.GetJobListByParentID(ctx, &jobPB.JobListRequest{Page: 1, Limit: 999, OrganizationID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}
	jobs := make([]*dto.Job, 0)

	for _, job := range foundResp.Jobs {
		jobs = append(jobs, dto.PBToJob(job))
	}

	SuccessResponse(rsp, D{
		"Jobs": jobs,
	})
}

func (j *JobHandler) Update(req *go_api.Request, rsp *go_api.Response) {
	// panic("not implemented")
	ctx := context.Background()
	params := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	jobService := j.getJobService()

	currentClubID := j.GetCurrentClubIDFromRequest(req)
	_, err := jobService.UpdateJob(ctx, &jobPB.UpdateJobRequest{
		ID:     params.ID,
		ClubID: currentClubID,
		Name:   params.Name,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

func (j *JobHandler) Remove(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	// panic("not implemented")
	params := struct {
		ID string `json:"id"`
	}{}

	if err := j.Bind(req, &params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	if err := j.Validate(&params); err != nil {
		ErrorResponse(rsp, j.Error().BadRequest(err.Error()))
		return
	}

	jobService := j.getJobService()
	_, err := jobService.RemoveJob(ctx, &jobPB.RemoveJobRequest{ID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}
