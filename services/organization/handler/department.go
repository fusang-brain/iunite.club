package handler

import (
	"context"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/bundles"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	orgPB "iunite.club/services/organization/proto"
	deptPB "iunite.club/services/organization/proto/department"
	"iunite.club/services/organization/srv"
	userPB "iunite.club/services/user/proto"
)

type DepartmentHandler struct {
	ironic.BaseHandler
}

func (j *DepartmentHandler) Model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (o *DepartmentHandler) GetDepartmentDetails(ctx context.Context, req *deptPB.GetDepartmentWithIDRequest, rsp *deptPB.DepartmentResponse) error {
	departmentService := srv.NewDepartmentService(ctx)

	department, err := departmentService.GetDepartmentDetailsByID(req.ID)

	if err != nil {
		return err
	}

	rsp.Department = department.ToPB()

	return nil
}

func (o *DepartmentHandler) CreateDepartment(ctx context.Context, req *deptPB.CreateDepartmentRequest, resp *deptPB.CreateDepartmentResponse) error {
	departmentService := srv.NewDepartmentService(ctx)
	department, err := departmentService.CreateDepartment(&srv.CreateDepartmentBundle{
		Name:        req.Name,
		ParentID:    req.ParentID,
		Description: req.Description,
	})

	if err != nil {
		return err
	}

	pb := department.ToPB()
	resp.Department = pb
	resp.OK = true

	return nil
}

func (o *DepartmentHandler) UpdateDepartment(ctx context.Context, req *deptPB.UpdateDepartmentRequest, resp *deptPB.UpdateDepartmentResponse) error {
	departmentService := srv.NewDepartmentService(ctx)
	err := departmentService.UpdateDepartment(&srv.UpdateDepartmentBundle{
		Name:        req.Name,
		ID:          req.ID,
		ParentID:    req.ParentID,
		Description: req.Description,
	})

	if err != nil {
		return err
	}

	resp.OK = true
	return nil
}

func (o *DepartmentHandler) RemoveDepartment(ctx context.Context, req *deptPB.RemoveDepartmentRequest, resp *deptPB.RemoveDepartmentResponse) error {
	departmentService := srv.NewDepartmentService(ctx)
	if err := departmentService.RemoveDepartment(req.ID); err != nil {
		return err
	}
	return nil
}

func (o *DepartmentHandler) GetDepartmentListByParentID(ctx context.Context, req *deptPB.DepartmentListByParentIDRequest, resp *deptPB.DepartmentListResponse) error {
	departmentService := srv.NewDepartmentService(ctx)
	result, err := departmentService.GetDepartmentListByParentID(&srv.GetDepartmentListBundle{
		ParentID: req.ParentID,
		Spread:   req.Spread,
		Search:   req.Search,
		PaginationBundle: bundles.PaginationBundle{
			Page:  int64(req.Page),
			Limit: int64(req.Limit),
		},
	})

	if err != nil {
		return err
	}

	resp.Departments = make([]*orgPB.Organization, 0, 1)

	for _, v := range result.Departments {
		pb := v.ToPB()
		resp.Departments = append(resp.Departments, pb)
	}

	resp.Total = int64(result.Total)

	return nil
}

func (o *DepartmentHandler) AddUsersToDepartment(ctx context.Context, req *deptPB.UserFromDepartmentRequest, rsp *deptPB.UpdateDepartmentResponse) error {
	UserClubProfile := o.Model(ctx, "UserClubProfile")
	Department := o.Model(ctx, "Organization")
	foundDeptCount := Department.Count(bson.M{"_id": bson.ObjectIdHex(req.DepartmentID)})

	if foundDeptCount == 0 {
		return o.Error(ctx).BadRequest("The department not exists")
	}

	userIDs := make([]bson.ObjectId, 0)

	for _, u := range req.Users {
		userIDs = append(userIDs, bson.ObjectIdHex(u))
	}

	_, err := UserClubProfile.Upsert(bson.M{
		"organization_id": bson.ObjectIdHex(req.ClubID),
		"user_id":         bson.M{"$in": userIDs},
	}, bson.M{"$set": bson.M{"department_id": bson.ObjectIdHex(req.DepartmentID)}})

	if err != nil {
		return o.Error(ctx).InternalServerError(err.Error())
	}

	return nil
}

func (o *DepartmentHandler) RemoveUsersFromDepartment(ctx context.Context, req *deptPB.UserFromDepartmentRequest, rsp *deptPB.UpdateDepartmentResponse) error {
	UserClubProfile := o.Model(ctx, "UserClubProfile")
	Department := o.Model(ctx, "Organization")
	foundDeptCount := Department.Count(bson.M{"_id": bson.ObjectIdHex(req.DepartmentID)})

	if foundDeptCount == 0 {
		return o.Error(ctx).BadRequest("The department not exists")
	}

	userIDs := make([]bson.ObjectId, 0)

	for _, u := range req.Users {
		userIDs = append(userIDs, bson.ObjectIdHex(u))
	}

	_, err := UserClubProfile.Upsert(bson.M{
		"organization_id": bson.ObjectIdHex(req.ClubID),
		"user_id":         bson.M{"$in": userIDs},
	}, bson.M{"$set": bson.M{"department_id": "-"}})

	if err != nil {
		return o.Error(ctx).InternalServerError(err.Error())
	}

	return nil
}

func (j *DepartmentHandler) GetUsersByDepartmentID(ctx context.Context, req *deptPB.ListByDepartmentIDRequest, rsp *deptPB.UserListResponse) error {
	UserClubProfile := j.Model(ctx, "UserClubProfile")

	ucps := make([]models.UserClubProfile, 0)
	users := make([]*userPB.User, 0)
	if !bson.IsObjectIdHex(req.DepartmentID) {
		return j.Error(ctx).BadRequest("departmentID must be objectid")
	}

	err := UserClubProfile.Where(bson.M{"department_id": bson.ObjectIdHex(req.DepartmentID)}).
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
	rsp.Total = int64(UserClubProfile.Count(bson.M{"department_id": bson.ObjectIdHex(req.DepartmentID)}))

	return nil
}

func (j *DepartmentHandler) GetAllCanSelectUsers(ctx context.Context, req *deptPB.ListByClubIDRequest, rsp *deptPB.UserListResponse) error {
	UserClubProfile := j.Model(ctx, "UserClubProfile")
	ucps := make([]models.UserClubProfile, 0)
	users := make([]*userPB.User, 0)

	// if !bson.IsObjectIdHex
	condition := bson.M{
		"$or": []bson.M{
			{
				"department_id": bson.M{"$exists": false},
			},
			{
				"department_id": "-",
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

func (o *DepartmentHandler) SearchDepartment(ctx context.Context, req *deptPB.SearchDepartmentRequest, rsp *deptPB.DepartmentListResponse) error {
	// panic("The function has not Impl")
	Department := o.Model(ctx, "Organization")

	departments := make([]models.Organization, 0)

	query := Department.Where(bson.M{
		"name": bson.RegEx{
			Pattern: req.Name,
			Options: "i",
		},
	})

	err := query.FindAll(&departments)
	if err != nil {
		return o.Error(ctx).InternalServerError(err.Error())
	}
	rsp.Total = int64(query.Count())
	depts := make([]*orgPB.Organization, 0)
	for _, dept := range departments {
		depts = append(depts, dept.ToPB())
	}
	rsp.Departments = depts

	return nil
}
