package handler

import (
	"context"
	"github.com/iron-kit/go-ironic/bundles"
	"github.com/iron-kit/go-ironic/micro-assistant"
	orgPB "iunite.club/srv/organization-srv/proto"
	jobPB "iunite.club/srv/organization-srv/proto/job"
	"iunite.club/srv/organization-srv/services"
)

type JobHandler struct {
	assistant.BaseHandler
	JobService *services.JobService
}

func (j *JobHandler) CreateJob(ctx context.Context, req *jobPB.CreateJobRequest, resp *jobPB.CreateJobResponse) error {
	orgJob, err := j.JobService.CreateJob(&services.CreateJobBundle{
		Name:   req.Name,
		ClubID: req.ClubID,
	})

	if err != nil {
		return err
	}

	resp.OK = true
	resp.Job = orgJob.ToPB()

	return nil
}

func (j *JobHandler) UpdateJob(ctx context.Context, req *jobPB.UpdateJobRequest, resp *jobPB.UpdateJobResponse) error {

	err := j.JobService.UpdateJob(&services.UpdateJobBundle{
		ID:     req.ID,
		Name:   req.Name,
		ClubID: req.ClubID,
	})

	if err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (j *JobHandler) RemoveJob(ctx context.Context, req *jobPB.RemoveJobRequest, resp *jobPB.RemoveJobResponse) error {
	if err := j.JobService.RemoveJob(req.ID); err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (j *JobHandler) GetJobListByParentID(ctx context.Context, req *jobPB.JobListRequest, resp *jobPB.JobListResponse) error {
	result, err := j.JobService.GetJobListByParentID(&services.JobListRequestBundle{
		OrganizationID: req.OrganizationID,
		PaginationBundle: bundles.PaginationBundle{
			Page:  int64(req.Page),
			Limit: int64(req.Limit),
		},
	})

	if err != nil {
		return err
	}

	resp.Jobs = make([]*orgPB.Job, 0, 1)

	for _, v := range result.Jobs {
		resp.Jobs = append(resp.Jobs, v.ToPB())
	}
	resp.Total = int64(result.Total)
	return nil
}
