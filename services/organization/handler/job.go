package handler

import (
	"context"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/bundles"
	orgPB "iunite.club/services/organization/proto"
	jobPB "iunite.club/services/organization/proto/job"
	"iunite.club/services/organization/srv"
)

type JobHandler struct {
	ironic.BaseHandler
}

func (j *JobHandler) CreateJob(ctx context.Context, req *jobPB.CreateJobRequest, resp *jobPB.CreateJobResponse) error {
	jobService := srv.NewJobService(ctx)
	orgJob, err := jobService.CreateJob(&srv.CreateJobBundle{
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
	jobService := srv.NewJobService(ctx)
	err := jobService.UpdateJob(&srv.UpdateJobBundle{
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
	jobService := srv.NewJobService(ctx)
	if err := jobService.RemoveJob(req.ID); err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (j *JobHandler) GetJobListByParentID(ctx context.Context, req *jobPB.JobListRequest, resp *jobPB.JobListResponse) error {

	jobService := srv.NewJobService(ctx)
	result, err := jobService.GetJobListByParentID(&srv.JobListRequestBundle{
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
