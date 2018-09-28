package handler

import (
	"context"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/bundles"
	orgPB "iunite.club/services/organization/proto"
	deptPB "iunite.club/services/organization/proto/department"
	"iunite.club/services/organization/srv"
)

type DepartmentHandler struct {
	ironic.BaseHandler
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
		Search: req.Search,
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
