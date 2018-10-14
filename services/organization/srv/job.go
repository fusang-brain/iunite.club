package srv

import (
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/bundles"
	"github.com/iron-kit/go-ironic/utils"
	"github.com/iron-kit/go-ironic/validator"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
)

type JobService struct {
	ironic.Service
}

type CreateJobBundle struct {
	Name   string `validate:"max=25,nonzero"`
	ClubID string `validate:"objectid,nonzero"`
}

type UpdateJobBundle struct {
	ID     string `validate:"objectid,nonzero"`
	Name   string `validate:"max=25"`
	ClubID string `validate:"objectid"`
}

type JobListRequestBundle struct {
	bundles.PaginationBundle
	OrganizationID string `validate:"objectid,nonzero"`
}

type JobListResponseBundle struct {
	Jobs  []models.OrganizationJob
	Total int
}

func (d *JobService) Model(name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(d.Ctx())

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

// CreateJob 新增职位
func (j *JobService) CreateJob(in *CreateJobBundle) (*models.OrganizationJob, error) {
	JobModel := j.Model("OrganizationJob")
	OrgModel := j.Model("Organization")
	if err := validator.Validate(in); err != nil {
		return nil, j.Error().TemplateBadRequest(err.Error())
	}
	foundClub := models.Organization{}

	OrgModel.FindByID(bson.ObjectIdHex(in.ClubID), &foundClub)

	if foundClub.IsEmpty() {
		return nil, j.Error().TemplateBadRequest("NotFoundClub")
	}

	newJob := models.OrganizationJob{
		Name:           in.Name,
		OrganizationID: foundClub.ID,
		Slug:           utils.Hans2Pinyin(in.Name, "_"),
	}

	if err := JobModel.Create(&newJob); err != nil {
		return nil, j.Error().InternalServerError(err.Error())
	}
	// if err := JobModel.Create(in)

	return &newJob, nil
}

// UpdateJob 更新职位
func (j *JobService) UpdateJob(in *UpdateJobBundle) error {
	JobModel := j.Model("OrganizationJob")

	if err := validator.Validate(in); err != nil {
		return j.Error().TemplateBadRequest(err.Error())
	}

	willUpdateJob := models.OrganizationJob{
		Name: in.Name,
		Slug: utils.Hans2Pinyin(in.Name, "_"),
	}

	if in.ClubID != "" {
		OrganizationModel := j.Model("Organization")
		foundClub := models.Organization{}
		// fmt.Println(in.ClubID)
		OrganizationModel.FindByID(bson.ObjectIdHex(in.ClubID), &foundClub)
		if foundClub.IsEmpty() {
			return j.Error().TemplateBadRequest("NotFoundCLub")
		}
		willUpdateJob.OrganizationID = foundClub.ID
		willUpdateJob.ID = bson.ObjectIdHex(in.ID)
	}
	// fmt.Println(in.ID)
	// return j.Error.TemplateBadRequest()
	if err := JobModel.Update(bson.M{"_id": bson.ObjectIdHex(in.ID)}, &willUpdateJob); err != nil {
		return j.Error().InternalServerError(err.Error())
	}

	return nil
}

func (j *JobService) RemoveJob(id string) error {
	if !bson.IsObjectIdHex(id) {
		return j.Error().TemplateBadRequest("ParamsError")
	}

	JobModel := j.Model("OrganizationJob")

	_, err := JobModel.UpsertID(bson.ObjectIdHex(id), bson.M{
		"$set": bson.M{
			"deleted": true,
		},
		// "deleted": true,
	})

	if err != nil {
		return j.Error().InternalServerError(err.Error())
	}

	return nil
}

// GetJobListByParentID 获取职位列表
func (j *JobService) GetJobListByParentID(in *JobListRequestBundle) (*JobListResponseBundle, error) {
	if err := validator.Validate(in); err != nil {
		return nil, j.Error().BadRequest(err.Error())
	}

	jobs := []models.OrganizationJob{}

	JobModel := j.Model("OrganizationJob")

	condition := bson.M{
		"organization_id": bson.ObjectIdHex(in.OrganizationID),
	}

	total := JobModel.Where(condition).Count()
	err := JobModel.
		Where(condition).
		Skip(int((in.Page - 1) * in.Limit)).
		Limit(int(in.Limit)).
		FindAll(&jobs)

	if err != nil {
		return nil, j.Error().InternalServerError(err.Error())
	}

	return &JobListResponseBundle{
		Jobs:  jobs,
		Total: total,
	}, nil
}
