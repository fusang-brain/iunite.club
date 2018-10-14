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

type DepartmentService struct {
	ironic.Service
}

type CreateDepartmentBundle struct {
	Name        string `validate:"max=25",nonzero`
	ParentID    string `validate:"objectid,nonzero"`
	Description string
}

type UpdateDepartmentBundle struct {
	ID          string `validate:"objectid,nonzero"`
	Name        string `validate:"max=25"`
	ParentID    string `validate:"objectid"`
	Description string
}

type GetDepartmentListBundle struct {
	bundles.PaginationBundle

	ParentID string `validate:"objectid,nonzero"`
	Search   string
	Spread   bool
}

type DepartmentListResponseBundle struct {
	Departments []models.Organization
	Total       int
}

func (d *DepartmentService) Model(name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(d.Ctx())

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

// CreateDepartment 创建部门
func (d *DepartmentService) CreateDepartment(in *CreateDepartmentBundle) (*models.Organization, error) {

	if err := validator.Validate(in); err != nil {
		return nil, d.Error().TemplateBadRequest(err.Error())
	}

	OrgModel := d.Model("Organization")
	foundParentOrg := models.Organization{}

	OrgModel.FindByID(bson.ObjectIdHex(in.ParentID), &foundParentOrg)

	if foundParentOrg.IsEmpty() {
		return nil, d.Error().InternalServerError("NotFoundClub")
	}

	clubID := foundParentOrg.ClubID
	if len(clubID) == 0 {
		clubID = foundParentOrg.ID
	}

	newDepartment := models.Organization{
		Kind:        "department",
		Name:        in.Name,
		Description: in.Description,
		Slug:        utils.Hans2Pinyin(in.Name, "_"),
		ParentID:    foundParentOrg.ID,
		SchoolID:    foundParentOrg.SchoolID,
		ClubID:      clubID,
		// PathIndexs: append(foundParentOrg.PathIndexs, models.PathIndex{
		// 	ID:   foundParentOrg.ID,
		// 	Name: foundParentOrg.Name,
		// 	Slug: foundParentOrg.Slug,
		// 	Sort: len(foundParentOrg.PathIndexs),
		// }),
	}

	OrgModel.Create(&newDepartment)

	return &newDepartment, nil
}

// UpdateDepartment 更新部门
func (d *DepartmentService) UpdateDepartment(in *UpdateDepartmentBundle) error {
	if err := validator.Validate(in); err != nil {
		return d.Error().TemplateBadRequest(err.Error())
	}
	OrgModel := d.Model("Organization")
	foundOrg := models.Organization{}

	OrgModel.Where(bson.M{"_id": bson.ObjectIdHex(in.ID)}).FindOne(&foundOrg)

	if foundOrg.IsEmpty() {
		return d.Error().ActionError("NotFoundOrganization")
	}

	if in.Name != "" {
		foundOrg.Name = in.Name
	}

	if in.Description != "" {
		foundOrg.Description = in.Description
	}

	if in.ParentID != "" {
		foundParentOrg := models.Organization{}
		OrgModel.Where(bson.M{"_id": bson.ObjectIdHex(in.ParentID)}).FindOne(&foundParentOrg)
		if foundParentOrg.IsEmpty() {
			return d.Error().ActionError("NotFoundOrganization")
		}

		foundOrg.ParentID = foundParentOrg.ID
		// foundOrg.PathIndexs = append(foundParentOrg.PathIndexs, models.PathIndex{
		// 	ID:   foundParentOrg.ID,
		// 	Name: foundParentOrg.Name,
		// 	Slug: foundParentOrg.Slug,
		// 	Sort: len(foundParentOrg.PathIndexs),
		// })
	}

	if err := OrgModel.Update(bson.M{"_id": foundOrg.ID}, &foundOrg); err != nil {
		return d.Error().InternalServerError(err.Error())
	}

	return nil
}

// RemoveDepartment 移出部门
func (d *DepartmentService) RemoveDepartment(id string) error {

	OrgModel := d.Model("Organization")

	if _, err := OrgModel.UpsertID(bson.ObjectIdHex(id), bson.M{"$set": bson.M{
		"deleted": true,
	}}); err != nil {
		return d.Error().InternalServerError(err.Error())
	}

	return nil
}

// GetDepartmentListByParentID  获取部门列表
func (d *DepartmentService) GetDepartmentListByParentID(in *GetDepartmentListBundle) (*DepartmentListResponseBundle, error) {
	if err := validator.Validate(in); err != nil {
		return nil, d.Error().TemplateBadRequest(err.Error())
	}

	OrgModel := d.Model("Organization")

	departments := []models.Organization{}

	condition := bson.M{}

	if in.Spread {
		condition["$or"] = []bson.M{
			{
				"parent_id": bson.ObjectIdHex(in.ParentID),
			},
			{
				"pathindexs._id": bson.ObjectIdHex(in.ParentID),
			},
		}
	} else {
		condition["parent_id"] = bson.ObjectIdHex(in.ParentID)
	}

	// total := OrgModel.Where(condition).Count()

	// err := OrgModel.
	// 	Where(condition).
	// 	Skip(int((in.Page - 1) * in.Limit)).
	// 	Limit(int(in.Limit)).
	// 	FindAll(&departments)

	err := OrgModel.Aggregate([]bson.M{
		{
			"$graphLookup": bson.M{
				"from":             "organization",
				"startWith":        "$parent_id",
				"connectFromField": "parent_id",
				"connectToField":   "_id",
				"as":               "pathindexs",
			},
		},
		{
			"$match": condition,
		},
		{
			"$skip": int((in.Page - 1) * in.Limit),
		},
		{
			"$limit": int(in.Limit),
		},
	}).Pipe().All(&departments)
	pipes := []bson.M{
		{
			"$graphLookup": bson.M{
				"from":             "organization",
				"startWith":        "$parent_id",
				"connectFromField": "parent_id",
				"connectToField":   "_id",
				"as":               "pathindexs",
			},
		},
		{
			"$match": condition,
		},
		{
			"$group": bson.M{
				"_id": "null",
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		{
			"$project": bson.M{"_id": 0},
		},
	}

	Cs := struct {
		Count int64 `bson:"count"`
	}{}
	OrgModel.Aggregate(pipes).Pipe().One(&Cs)

	if err != nil {
		return nil, d.Error().InternalServerError(err.Error())
	}

	return &DepartmentListResponseBundle{
		Departments: departments,
		Total:       int(Cs.Count),
	}, err
}

func (d *DepartmentService) GetDepartmentDetailsByID(id string) (*models.Organization, error) {
	DepartmentModel := d.Model("Organization")
	dept := new(models.Organization)
	if !bson.IsObjectIdHex(id) {
		return dept, d.Error().BadRequest("ID must be a objectid hex string")
	}
	err := DepartmentModel.FindByID(bson.ObjectIdHex(id), dept)

	if err != nil {
		return dept, d.Error().InternalServerError(err.Error())
	}

	return dept, nil
}

// func (d *DepartmentService) JoinPeopleToDepartment()
