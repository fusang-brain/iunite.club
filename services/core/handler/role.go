package handler

import (
	"context"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/core/proto/role"
	"iunite.club/services/user/proto"
)

type Role struct {
	ironic.BaseHandler
}

func (r *Role) model(ctx context.Context, modelName string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		panic(err.Error())
	}

	return conn.M(modelName)
}

func (r *Role) CreateRoleGroup(ctx context.Context, req *pb.CreateRoleGroupRequest, rsp *pb.CreatedGroupResponse) error {
	RoleGroupModel := r.model(ctx, "RoleGroup")
	newRoleGroup := &models.RoleGroup{
		GroupName: req.Name,
		GroupDescription: req.Description,
		Organization: bson.ObjectIdHex(req.Organization),
	}
	if err := RoleGroupModel.Create(newRoleGroup); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}
	// panic("not implemented")
	rsp.ID = newRoleGroup.ID.Hex()
	rsp.CreatedAt = hptypes.TimestampProto(newRoleGroup.CreatedAt)
	return nil
}

func (r *Role) FindAllRolesByOrganization(ctx context.Context, req *pb.ByOrganizationRequest, rsp *pb.RolesResponse) error {
	// panic("not implemented")
	RoleModel := r.model(ctx, "Role")
	roles := make([]models.Role, 0)

	if err := RoleModel.Where(bson.M{
		"organization": req.Organization,
	}).FindAll(&roles); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	total := RoleModel.Where(bson.M{
		"organization": req.Organization,
	}).Count()

	rsp.Total = int32(total)

	rolePBs := make([]*pb.RolePB, 0, total)
	for _, v := range roles {
		rolePBs = append(rolePBs, v.ToPB())
	}

	rsp.Roles = rolePBs

	return nil
}

func (r *Role) FindAllGroupByOrganization(ctx context.Context, req *pb.ByOrganizationRequest, rsp *pb.GroupsResponse) error {

	// panic("not implemented")
	RoleGroupModel := r.model(ctx, "RoleGroup")
	groups := make([]models.RoleGroup, 0)
	if err := RoleGroupModel.Where(bson.M{
		"organization": req.Organization,
	}).FindAll(&groups); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	total := RoleGroupModel.Count(bson.M{"organization": req.Organization})

	rsp.Total = int32(total)
	groupPBs := make([]*pb.RoleGroupPB, 0, total)
	for _, v := range groups {
		groupPBs = append(groupPBs, v.ToPB())
	}

	rsp.Groups = groupPBs

	return nil
}

func (r *Role) FindUsersByRoleID(ctx context.Context, req *pb.ByRoleRequest, rsp *pb.UsersResponse) error {
	RoleUserModel := r.model(ctx, "RoleUser")

	roleUsers := make([]models.RoleUser, 0)
	if err := RoleUserModel.FindAll(&roleUsers, bson.M{
		"role_id": req.Role,
	}); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	total := RoleUserModel.Count(bson.M{"role_id": req.Role})
	users := make([]*iunite_club_srv_user.User, 0, total)
	for _, v := range roleUsers {
		users = append(users, v.User.ToPB())
	}

	rsp.Total = int32(total)
	rsp.Users = users
	return nil
	// panic("not implemented")
}

func (r *Role) UpdateRoleGroup(ctx context.Context, req *pb.UpdateRoleGroupRequest, rsp *pb.UpdatedResponse) error {
	RoleGroupModel := r.model(ctx, "Role")
	foundRoleGroup := new(models.RoleGroup)
	RoleGroupModel.Where(bson.M{"_id": req.RoleGroup.ID}).FindOne(foundRoleGroup)

	if foundRoleGroup.IsEmpty() {
		return r.Error(ctx).NotFound("Not found the role group")
	}

	rg := req.RoleGroup

	if rg.Organization != "" {
		foundRoleGroup.Organization = bson.ObjectIdHex(rg.Organization)
	}
	if rg.GroupDescription != "" {
		foundRoleGroup.GroupDescription = rg.GroupDescription
	}
	if rg.GroupName != "" {
		foundRoleGroup.GroupName = rg.GroupName
	}

	if err := RoleGroupModel.Update(bson.M{"_id": req.RoleGroup.ID}, foundRoleGroup); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.UpdatedAt = hptypes.TimestampProto(foundRoleGroup.UpdatedAt)
	rsp.OK = true

	return nil
}

func (r *Role) DeleteRoleGroup(ctx context.Context, req *pb.ByGroupIDRequest, rsp *pb.DeletedRoleGroupResponse) error {

	RoleGroupModel := r.model(ctx, "RoleGroup")
	if err := RoleGroupModel.Where(bson.M{"_id": req.GroupID}).Delete(); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	rsp.DeletedAt = hptypes.TimestampNow()
	return nil
}

func (r *Role) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest, rsp *pb.UpdateRoleResponse) error {
	// panic("not implemented")
	RoleModel := r.model(ctx, "Role")

	foundRole := new(models.Role)
	if err := RoleModel.Where(bson.M{"_id": req.ID}).FindOne(foundRole); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	if foundRole.IsEmpty() {
		return r.Error(ctx).NotFound("Not found this role")
	}

	if req.GroupID != "" && bson.IsObjectIdHex(req.GroupID) {
		foundRole.GroupID =  bson.ObjectIdHex(req.GroupID)
	}

	if req.Organization != "" && bson.IsObjectIdHex(req.Organization) {
		foundRole.Organization = bson.ObjectIdHex(req.Organization)
	}

	if req.Level != "" {
		foundRole.Level = req.Level
	}

	if req.Name != "" {
		foundRole.Name = req.Name
	}

	if err := RoleModel.Update(bson.M{"_id": req.ID}, foundRole); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.ID = foundRole.ID.Hex()
	rsp.UpdatedAt = hptypes.TimestampProto(foundRole.UpdatedAt)

	return nil
}

func (r *Role) DeleteRole(ctx context.Context, req *pb.ByIDRequest, rsp *pb.DeletedResponse) error {
	RoleModel := r.model(ctx, "Role")
	// foundRole := new(models.Role)
	if err := RoleModel.Delete(bson.M{"_id": req.ID}); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	rsp.DeletedAt = hptypes.TimestampNow()
	return nil
}

func (r *Role) CreateRole(ctx context.Context, req *pb.CreateRoleRequest, rsp *pb.CreatedRoleResponse) error {
	// panic("not implemented")
	RoleModel := r.model(ctx, "Role")
	newRole := &models.Role{
		Name: req.Name,
		Level: req.Level,
		GroupID: bson.ObjectIdHex(req.GroupID),
		Organization: bson.ObjectIdHex(req.Organization),
	}
	if err := RoleModel.Create(newRole); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.ID = newRole.ID.Hex()
	rsp.CreatedAt = hptypes.TimestampProto(newRole.CreatedAt)
	return nil
}

func (r *Role) AddUsersToRole(ctx context.Context, req *pb.UsersAndRoleRequest, rsp *pb.CreatedResponse) error {
	RoleUserModel := r.model(ctx, "RoleUser")

	roleUsers := make([]models.RoleUser, 0, len(req.Users))

	for _, value := range req.Users {
		if bson.IsObjectIdHex(value) && bson.IsObjectIdHex(req.Role) {
			roleUsers = append(roleUsers, models.RoleUser{
				UserID: bson.ObjectIdHex(value),
				RoleID: bson.ObjectIdHex(req.Role),
			})
		}
	}
	if len(roleUsers) > 0 {
		if err := RoleUserModel.Collection().Insert(roleUsers); err != nil {
			return r.Error(ctx).InternalServerError(err.Error())
		}
	}

	rsp.OK = true
	rsp.CreatedAt = hptypes.TimestampNow()
	return nil
}

func (r *Role) RemoveUsersToRole(ctx context.Context, req *pb.UsersAndRoleRequest, rsp *pb.DeletedResponse) error {
	RoleUserModel := r.model(ctx, "RoleUser")

	users := make([]bson.ObjectId, 0)
	for _, v := range req.Users {
		users = append(users, bson.ObjectIdHex(v))
	}

	if err := RoleUserModel.ForceDeleteAll(bson.M{
		"role_id": bson.ObjectIdHex(req.Role),
		"user_id": bson.M{"$in": users},
	}); err != nil {
		return r.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	rsp.DeletedAt = hptypes.TimestampNow()
	return nil
}



