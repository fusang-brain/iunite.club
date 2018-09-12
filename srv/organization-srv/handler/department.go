package handler

import (
	"context"
	"github.com/iron-kit/go-ironic/bundles"
	"github.com/iron-kit/go-ironic/micro-assistant"
	orgPB "iunite.club/srv/organization-srv/proto"
	deptPB "iunite.club/srv/organization-srv/proto/department"
	"iunite.club/srv/organization-srv/services"
)

type DepartmentHandler struct {
	assistant.BaseHandler
	DepartmentService *services.DepartmentService
}

func (o *DepartmentHandler) CreateDepartment(ctx context.Context, req *deptPB.CreateDepartmentRequest, resp *deptPB.CreateDepartmentResponse) error {
	department, err := o.DepartmentService.CreateDepartment(&services.CreateDepartmentBundle{
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
	err := o.DepartmentService.UpdateDepartment(&services.UpdateDepartmentBundle{
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
	if err := o.DepartmentService.RemoveDepartment(req.ID); err != nil {
		return err
	}
	return nil
}

func (o *DepartmentHandler) GetDepartmentListByParentID(ctx context.Context, req *deptPB.DepartmentListByParentIDRequest, resp *deptPB.DepartmentListResponse) error {
	result, err := o.DepartmentService.GetDepartmentListByParentID(&services.GetDepartmentListBundle{
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
