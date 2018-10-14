package handler

import (
	"context"

	userPB "iunite.club/services/user/proto"

	"gopkg.in/mgo.v2/bson"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/bundles"
	"github.com/iron-kit/monger"
	"iunite.club/models"
	orgPB "iunite.club/services/organization/proto"
	jobPB "iunite.club/services/organization/proto/job"
	"iunite.club/services/organization/srv"
)

type JobHandler struct {
	ironic.BaseHandler
}

func (j *JobHandler) Model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
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
	// fmt.Println(result.Jobs, "====")
	resp.Jobs = make([]*orgPB.Job, 0)
	// fmt.Println(resp.Jobs, "====")
	for _, v := range result.Jobs {
		resp.Jobs = append(resp.Jobs, v.ToPB())
	}
	resp.Total = int64(result.Total)
	return nil
}

func (j *JobHandler) GetUsersByJobID(ctx context.Context, req *jobPB.ListByJobIDRequest, rsp *jobPB.UserListResponse) error {
	UserClubProfile := j.Model(ctx, "UserClubProfile")

	ucps := make([]models.UserClubProfile, 0)
	users := make([]*userPB.User, 0)
	if !bson.IsObjectIdHex(req.JobID) {
		return j.Error(ctx).BadRequest("jobID must be objectid")
	}

	err := UserClubProfile.Where(bson.M{"job_id": bson.ObjectIdHex(req.JobID)}).
		Skip(int((req.Page-1)*req.Limit)).
		Limit(int(req.Limit)).
		Populate("User", "User.Profile").
		FindAll(&ucps)

	if err != nil {
		return j.Error(ctx).InternalServerError(err.Error())
	}

	for _, ucp := range ucps {
		user := ucp.User

		users = append(users, user.ToPB())
	}

	rsp.Users = users
	rsp.Total = int64(UserClubProfile.Count(bson.M{"job_id": bson.ObjectIdHex(req.JobID)}))

	return nil
}

func (j *JobHandler) GetAllCanSelectedUsers(ctx context.Context, req *jobPB.ListByClubIDRequest, rsp *jobPB.UserListResponse) error {
	UserClubProfile := j.Model(ctx, "UserClubProfile")
	ucps := make([]models.UserClubProfile, 0)
	users := make([]*userPB.User, 0)

	// if !bson.IsObjectIdHex
	condition := bson.M{
		"$or": []bson.M{
			{
				"job_id": bson.M{"$exists": false},
			},
			{
				"job_id": "-",
			},
		},
		"organization_id": bson.ObjectIdHex(req.ClubID),
		"state":           1,
	}
	rsp.Total = int64(UserClubProfile.Count(condition))
	err := UserClubProfile.
		Where(condition).
		Populate("User", "User.Profile").
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&ucps)

	if err != nil {
		return j.Error(ctx).InternalServerError(err.Error())
	}

	for _, ucp := range ucps {
		user := ucp.User

		users = append(users, user.ToPB())
	}

	rsp.Users = users

	return nil
}

func (j *JobHandler) AddUsersToJob(ctx context.Context, req *jobPB.UserFromJobRequest, rsp *jobPB.UpdateJobResponse) error {
	UserClubProfile := j.Model(ctx, "UserClubProfile")
	Job := j.Model(ctx, "OrganizationJob")
	foundJobCount := Job.Count(bson.M{"_id": bson.ObjectIdHex(req.JobID)})

	if foundJobCount == 0 {
		return j.Error(ctx).BadRequest("The job not exists")
	}

	userIDs := make([]bson.ObjectId, 0)

	for _, u := range req.Users {
		userIDs = append(userIDs, bson.ObjectIdHex(u))
	}

	_, err := UserClubProfile.Upsert(bson.M{
		"organization_id": bson.ObjectIdHex(req.ClubID),
		"user_id":         bson.M{"$in": userIDs},
	}, bson.M{"$set": bson.M{"job_id": bson.ObjectIdHex(req.JobID)}})

	if err != nil {
		return j.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true

	return nil
}

func (j *JobHandler) RemoveUsersFromJob(ctx context.Context, req *jobPB.UserFromJobRequest, rsp *jobPB.UpdateJobResponse) error {
	UserClubProfile := j.Model(ctx, "UserClubProfile")
	Job := j.Model(ctx, "OrganizationJob")

	foundJobCount := Job.Count(bson.M{"_id": bson.ObjectIdHex(req.JobID)})

	if foundJobCount == 0 {
		return j.Error(ctx).BadRequest("The job not exists")
	}
	userIDs := make([]bson.ObjectId, 0)

	for _, u := range req.Users {
		userIDs = append(userIDs, bson.ObjectIdHex(u))
	}

	_, err := UserClubProfile.Upsert(bson.M{
		"organization_id": bson.ObjectIdHex(req.ClubID),
		"user_id":         bson.M{"$in": userIDs},
	}, bson.M{"$set": bson.M{"job_id": "-"}})

	if err != nil {
		return j.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true

	return nil
}
