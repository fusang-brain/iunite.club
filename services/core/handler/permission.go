package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/core/proto/permission"
)

type Permission struct {
	ironic.BaseHandler
}

func (r *Permission) model(ctx context.Context, modelName string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		panic(err.Error())
	}

	return conn.M(modelName)
}

func (p *Permission) CreatePermission(ctx context.Context, request *pb.CreateRequest, resp *pb.CreatedPermissionResponse) error {
	PermissionModel := p.model(ctx, "Permission")
	newPermission := &models.Permission{
		PermissionName: request.PermissionName,
		PermissionKey: uuid.New().String(),
		PathRule: request.PathRule,
		Method: request.Method,
	}

	if err := PermissionModel.Create(newPermission); err != nil {
		return p.Error(ctx).InternalServerError(err.Error())
	}

	return nil
}

func (p *Permission) DeletePermission (ctx context.Context, req *pb.ByIDRequest, resp *pb.DeletedResponse) error {
	CasbinRuleModel := p.model(ctx, "CasbinRuleModel")

	if err := CasbinRuleModel.ForceDeleteAll(bson.M{
		"permission_id": bson.ObjectIdHex(req.ID),
	}); err != nil {
		return p.Error(ctx).InternalServerError(err.Error())
	}

	return nil
}