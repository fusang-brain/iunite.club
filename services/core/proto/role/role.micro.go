// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/role/role.proto

/*
Package iunite_club_srv_core_role is a generated protocol buffer package.

It is generated from these files:
	proto/role/role.proto

It has these top-level messages:
	RoleGroupPB
	RolePB
	UpdateRoleRequest
	UpdateRoleResponse
	ByIDRequest
	DeletedResponse
	CreateRoleRequest
	CreatedRoleResponse
	UsersAndRoleRequest
	CreatedResponse
	UpdateRoleGroupRequest
	UpdatedResponse
	ByGroupIDRequest
	DeletedRoleGroupResponse
	ByRoleRequest
	UsersResponse
	ByOrganizationRequest
	RolesResponse
	GroupsResponse
	CreateRoleGroupRequest
	CreatedGroupResponse
*/
package iunite_club_srv_core_role

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "iunite.club/services/user/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Role service

type RoleService interface {
	CreateRoleGroup(ctx context.Context, in *CreateRoleGroupRequest, opts ...client.CallOption) (*CreatedGroupResponse, error)
	FindAllRolesByOrganization(ctx context.Context, in *ByOrganizationRequest, opts ...client.CallOption) (*RolesResponse, error)
	FindAllGroupByOrganization(ctx context.Context, in *ByOrganizationRequest, opts ...client.CallOption) (*GroupsResponse, error)
	FindUsersByRoleID(ctx context.Context, in *ByRoleRequest, opts ...client.CallOption) (*UsersResponse, error)
	UpdateRoleGroup(ctx context.Context, in *UpdateRoleGroupRequest, opts ...client.CallOption) (*UpdatedResponse, error)
	DeleteRoleGroup(ctx context.Context, in *ByGroupIDRequest, opts ...client.CallOption) (*DeletedRoleGroupResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...client.CallOption) (*UpdateRoleResponse, error)
	DeleteRole(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*DeletedResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*CreatedRoleResponse, error)
	AddUsersToRole(ctx context.Context, in *UsersAndRoleRequest, opts ...client.CallOption) (*CreatedResponse, error)
	RemoveUsersToRole(ctx context.Context, in *UsersAndRoleRequest, opts ...client.CallOption) (*DeletedResponse, error)
}

type roleService struct {
	c    client.Client
	name string
}

func NewRoleService(name string, c client.Client) RoleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iunite.club.srv.core.role"
	}
	return &roleService{
		c:    c,
		name: name,
	}
}

func (c *roleService) CreateRoleGroup(ctx context.Context, in *CreateRoleGroupRequest, opts ...client.CallOption) (*CreatedGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Role.CreateRoleGroup", in)
	out := new(CreatedGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) FindAllRolesByOrganization(ctx context.Context, in *ByOrganizationRequest, opts ...client.CallOption) (*RolesResponse, error) {
	req := c.c.NewRequest(c.name, "Role.FindAllRolesByOrganization", in)
	out := new(RolesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) FindAllGroupByOrganization(ctx context.Context, in *ByOrganizationRequest, opts ...client.CallOption) (*GroupsResponse, error) {
	req := c.c.NewRequest(c.name, "Role.FindAllGroupByOrganization", in)
	out := new(GroupsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) FindUsersByRoleID(ctx context.Context, in *ByRoleRequest, opts ...client.CallOption) (*UsersResponse, error) {
	req := c.c.NewRequest(c.name, "Role.FindUsersByRoleID", in)
	out := new(UsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) UpdateRoleGroup(ctx context.Context, in *UpdateRoleGroupRequest, opts ...client.CallOption) (*UpdatedResponse, error) {
	req := c.c.NewRequest(c.name, "Role.UpdateRoleGroup", in)
	out := new(UpdatedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) DeleteRoleGroup(ctx context.Context, in *ByGroupIDRequest, opts ...client.CallOption) (*DeletedRoleGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Role.DeleteRoleGroup", in)
	out := new(DeletedRoleGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...client.CallOption) (*UpdateRoleResponse, error) {
	req := c.c.NewRequest(c.name, "Role.UpdateRole", in)
	out := new(UpdateRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) DeleteRole(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*DeletedResponse, error) {
	req := c.c.NewRequest(c.name, "Role.DeleteRole", in)
	out := new(DeletedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*CreatedRoleResponse, error) {
	req := c.c.NewRequest(c.name, "Role.CreateRole", in)
	out := new(CreatedRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) AddUsersToRole(ctx context.Context, in *UsersAndRoleRequest, opts ...client.CallOption) (*CreatedResponse, error) {
	req := c.c.NewRequest(c.name, "Role.AddUsersToRole", in)
	out := new(CreatedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) RemoveUsersToRole(ctx context.Context, in *UsersAndRoleRequest, opts ...client.CallOption) (*DeletedResponse, error) {
	req := c.c.NewRequest(c.name, "Role.RemoveUsersToRole", in)
	out := new(DeletedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Role service

type RoleHandler interface {
	CreateRoleGroup(context.Context, *CreateRoleGroupRequest, *CreatedGroupResponse) error
	FindAllRolesByOrganization(context.Context, *ByOrganizationRequest, *RolesResponse) error
	FindAllGroupByOrganization(context.Context, *ByOrganizationRequest, *GroupsResponse) error
	FindUsersByRoleID(context.Context, *ByRoleRequest, *UsersResponse) error
	UpdateRoleGroup(context.Context, *UpdateRoleGroupRequest, *UpdatedResponse) error
	DeleteRoleGroup(context.Context, *ByGroupIDRequest, *DeletedRoleGroupResponse) error
	UpdateRole(context.Context, *UpdateRoleRequest, *UpdateRoleResponse) error
	DeleteRole(context.Context, *ByIDRequest, *DeletedResponse) error
	CreateRole(context.Context, *CreateRoleRequest, *CreatedRoleResponse) error
	AddUsersToRole(context.Context, *UsersAndRoleRequest, *CreatedResponse) error
	RemoveUsersToRole(context.Context, *UsersAndRoleRequest, *DeletedResponse) error
}

func RegisterRoleHandler(s server.Server, hdlr RoleHandler, opts ...server.HandlerOption) {
	type role interface {
		CreateRoleGroup(ctx context.Context, in *CreateRoleGroupRequest, out *CreatedGroupResponse) error
		FindAllRolesByOrganization(ctx context.Context, in *ByOrganizationRequest, out *RolesResponse) error
		FindAllGroupByOrganization(ctx context.Context, in *ByOrganizationRequest, out *GroupsResponse) error
		FindUsersByRoleID(ctx context.Context, in *ByRoleRequest, out *UsersResponse) error
		UpdateRoleGroup(ctx context.Context, in *UpdateRoleGroupRequest, out *UpdatedResponse) error
		DeleteRoleGroup(ctx context.Context, in *ByGroupIDRequest, out *DeletedRoleGroupResponse) error
		UpdateRole(ctx context.Context, in *UpdateRoleRequest, out *UpdateRoleResponse) error
		DeleteRole(ctx context.Context, in *ByIDRequest, out *DeletedResponse) error
		CreateRole(ctx context.Context, in *CreateRoleRequest, out *CreatedRoleResponse) error
		AddUsersToRole(ctx context.Context, in *UsersAndRoleRequest, out *CreatedResponse) error
		RemoveUsersToRole(ctx context.Context, in *UsersAndRoleRequest, out *DeletedResponse) error
	}
	type Role struct {
		role
	}
	h := &roleHandler{hdlr}
	s.Handle(s.NewHandler(&Role{h}, opts...))
}

type roleHandler struct {
	RoleHandler
}

func (h *roleHandler) CreateRoleGroup(ctx context.Context, in *CreateRoleGroupRequest, out *CreatedGroupResponse) error {
	return h.RoleHandler.CreateRoleGroup(ctx, in, out)
}

func (h *roleHandler) FindAllRolesByOrganization(ctx context.Context, in *ByOrganizationRequest, out *RolesResponse) error {
	return h.RoleHandler.FindAllRolesByOrganization(ctx, in, out)
}

func (h *roleHandler) FindAllGroupByOrganization(ctx context.Context, in *ByOrganizationRequest, out *GroupsResponse) error {
	return h.RoleHandler.FindAllGroupByOrganization(ctx, in, out)
}

func (h *roleHandler) FindUsersByRoleID(ctx context.Context, in *ByRoleRequest, out *UsersResponse) error {
	return h.RoleHandler.FindUsersByRoleID(ctx, in, out)
}

func (h *roleHandler) UpdateRoleGroup(ctx context.Context, in *UpdateRoleGroupRequest, out *UpdatedResponse) error {
	return h.RoleHandler.UpdateRoleGroup(ctx, in, out)
}

func (h *roleHandler) DeleteRoleGroup(ctx context.Context, in *ByGroupIDRequest, out *DeletedRoleGroupResponse) error {
	return h.RoleHandler.DeleteRoleGroup(ctx, in, out)
}

func (h *roleHandler) UpdateRole(ctx context.Context, in *UpdateRoleRequest, out *UpdateRoleResponse) error {
	return h.RoleHandler.UpdateRole(ctx, in, out)
}

func (h *roleHandler) DeleteRole(ctx context.Context, in *ByIDRequest, out *DeletedResponse) error {
	return h.RoleHandler.DeleteRole(ctx, in, out)
}

func (h *roleHandler) CreateRole(ctx context.Context, in *CreateRoleRequest, out *CreatedRoleResponse) error {
	return h.RoleHandler.CreateRole(ctx, in, out)
}

func (h *roleHandler) AddUsersToRole(ctx context.Context, in *UsersAndRoleRequest, out *CreatedResponse) error {
	return h.RoleHandler.AddUsersToRole(ctx, in, out)
}

func (h *roleHandler) RemoveUsersToRole(ctx context.Context, in *UsersAndRoleRequest, out *DeletedResponse) error {
	return h.RoleHandler.RemoveUsersToRole(ctx, in, out)
}
