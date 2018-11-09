// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

/*
Package iunite_club_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user.proto

It has these top-level messages:
	UpdateAvatarRequest
	ByOrganizationIDRequest
	FindUserClubProfileByIDRequest
	UserClubProfileResponse
	CreateMemberRequest
	FindUsersByClubIDRequest
	ResetPasswordResponse
	UserClubProfile
	Department
	User
	Profile
	ResetPasswordRequest
	SigninByMobileRequest
	RegisterUserRequest
	RegisterUserResponse
	Response
	UserListResponse
	PagerRequest
	UpdateUserRequest
	QueryUserRequest
	UserResponse
	QueryProfileRequest
	ProfileResponse
*/
package iunite_club_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for UserSrv service

type UserSrvService interface {
	FindUserByID(ctx context.Context, in *QueryUserRequest, opts ...client.CallOption) (*UserResponse, error)
	FindProfileByID(ctx context.Context, in *QueryProfileRequest, opts ...client.CallOption) (*ProfileResponse, error)
	CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*Response, error)
	FindUsers(ctx context.Context, in *PagerRequest, opts ...client.CallOption) (*UserListResponse, error)
	IsUserEnabled(ctx context.Context, in *QueryUserRequest, opts ...client.CallOption) (*Response, error)
	// RegisterUserByMobile 注册一个新的用户
	RegisterUserByMobile(ctx context.Context, in *RegisterUserRequest, opts ...client.CallOption) (*RegisterUserResponse, error)
	// ResetPasswordByMobile 重置密码
	ResetPasswordByMobile(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*ResetPasswordResponse, error)
	SigninByMobile(ctx context.Context, in *SigninByMobileRequest, opts ...client.CallOption) (*UserResponse, error)
	// rpc ResetPassword(ResetPasswordRequest) returns(ResetPasswordResponse) {}
	FindUsersByClubID(ctx context.Context, in *FindUsersByClubIDRequest, opts ...client.CallOption) (*UserListResponse, error)
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...client.CallOption) (*Response, error)
	FindUserClubProfileByID(ctx context.Context, in *FindUserClubProfileByIDRequest, opts ...client.CallOption) (*UserClubProfileResponse, error)
	FindUsersByOrganizationID(ctx context.Context, in *ByOrganizationIDRequest, opts ...client.CallOption) (*UserListResponse, error)
	UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, opts ...client.CallOption) (*Response, error)
}

type userSrvService struct {
	c    client.Client
	name string
}

func NewUserSrvService(name string, c client.Client) UserSrvService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iunite.club.srv.user"
	}
	return &userSrvService{
		c:    c,
		name: name,
	}
}

func (c *userSrvService) FindUserByID(ctx context.Context, in *QueryUserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.FindUserByID", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) FindProfileByID(ctx context.Context, in *QueryProfileRequest, opts ...client.CallOption) (*ProfileResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.FindProfileByID", in)
	out := new(ProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserSrv.CreateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserSrv.UpdateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) FindUsers(ctx context.Context, in *PagerRequest, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.FindUsers", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) IsUserEnabled(ctx context.Context, in *QueryUserRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserSrv.IsUserEnabled", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) RegisterUserByMobile(ctx context.Context, in *RegisterUserRequest, opts ...client.CallOption) (*RegisterUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.RegisterUserByMobile", in)
	out := new(RegisterUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) ResetPasswordByMobile(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*ResetPasswordResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.ResetPasswordByMobile", in)
	out := new(ResetPasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) SigninByMobile(ctx context.Context, in *SigninByMobileRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.SigninByMobile", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) FindUsersByClubID(ctx context.Context, in *FindUsersByClubIDRequest, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.FindUsersByClubID", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserSrv.CreateMember", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) FindUserClubProfileByID(ctx context.Context, in *FindUserClubProfileByIDRequest, opts ...client.CallOption) (*UserClubProfileResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.FindUserClubProfileByID", in)
	out := new(UserClubProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) FindUsersByOrganizationID(ctx context.Context, in *ByOrganizationIDRequest, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.FindUsersByOrganizationID", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserSrv.UpdateAvatar", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserSrv service

type UserSrvHandler interface {
	FindUserByID(context.Context, *QueryUserRequest, *UserResponse) error
	FindProfileByID(context.Context, *QueryProfileRequest, *ProfileResponse) error
	CreateUser(context.Context, *User, *Response) error
	UpdateUser(context.Context, *UpdateUserRequest, *Response) error
	FindUsers(context.Context, *PagerRequest, *UserListResponse) error
	IsUserEnabled(context.Context, *QueryUserRequest, *Response) error
	// RegisterUserByMobile 注册一个新的用户
	RegisterUserByMobile(context.Context, *RegisterUserRequest, *RegisterUserResponse) error
	// ResetPasswordByMobile 重置密码
	ResetPasswordByMobile(context.Context, *ResetPasswordRequest, *ResetPasswordResponse) error
	SigninByMobile(context.Context, *SigninByMobileRequest, *UserResponse) error
	// rpc ResetPassword(ResetPasswordRequest) returns(ResetPasswordResponse) {}
	FindUsersByClubID(context.Context, *FindUsersByClubIDRequest, *UserListResponse) error
	CreateMember(context.Context, *CreateMemberRequest, *Response) error
	FindUserClubProfileByID(context.Context, *FindUserClubProfileByIDRequest, *UserClubProfileResponse) error
	FindUsersByOrganizationID(context.Context, *ByOrganizationIDRequest, *UserListResponse) error
	UpdateAvatar(context.Context, *UpdateAvatarRequest, *Response) error
}

func RegisterUserSrvHandler(s server.Server, hdlr UserSrvHandler, opts ...server.HandlerOption) {
	type userSrv interface {
		FindUserByID(ctx context.Context, in *QueryUserRequest, out *UserResponse) error
		FindProfileByID(ctx context.Context, in *QueryProfileRequest, out *ProfileResponse) error
		CreateUser(ctx context.Context, in *User, out *Response) error
		UpdateUser(ctx context.Context, in *UpdateUserRequest, out *Response) error
		FindUsers(ctx context.Context, in *PagerRequest, out *UserListResponse) error
		IsUserEnabled(ctx context.Context, in *QueryUserRequest, out *Response) error
		RegisterUserByMobile(ctx context.Context, in *RegisterUserRequest, out *RegisterUserResponse) error
		ResetPasswordByMobile(ctx context.Context, in *ResetPasswordRequest, out *ResetPasswordResponse) error
		SigninByMobile(ctx context.Context, in *SigninByMobileRequest, out *UserResponse) error
		FindUsersByClubID(ctx context.Context, in *FindUsersByClubIDRequest, out *UserListResponse) error
		CreateMember(ctx context.Context, in *CreateMemberRequest, out *Response) error
		FindUserClubProfileByID(ctx context.Context, in *FindUserClubProfileByIDRequest, out *UserClubProfileResponse) error
		FindUsersByOrganizationID(ctx context.Context, in *ByOrganizationIDRequest, out *UserListResponse) error
		UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, out *Response) error
	}
	type UserSrv struct {
		userSrv
	}
	h := &userSrvHandler{hdlr}
	s.Handle(s.NewHandler(&UserSrv{h}, opts...))
}

type userSrvHandler struct {
	UserSrvHandler
}

func (h *userSrvHandler) FindUserByID(ctx context.Context, in *QueryUserRequest, out *UserResponse) error {
	return h.UserSrvHandler.FindUserByID(ctx, in, out)
}

func (h *userSrvHandler) FindProfileByID(ctx context.Context, in *QueryProfileRequest, out *ProfileResponse) error {
	return h.UserSrvHandler.FindProfileByID(ctx, in, out)
}

func (h *userSrvHandler) CreateUser(ctx context.Context, in *User, out *Response) error {
	return h.UserSrvHandler.CreateUser(ctx, in, out)
}

func (h *userSrvHandler) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *Response) error {
	return h.UserSrvHandler.UpdateUser(ctx, in, out)
}

func (h *userSrvHandler) FindUsers(ctx context.Context, in *PagerRequest, out *UserListResponse) error {
	return h.UserSrvHandler.FindUsers(ctx, in, out)
}

func (h *userSrvHandler) IsUserEnabled(ctx context.Context, in *QueryUserRequest, out *Response) error {
	return h.UserSrvHandler.IsUserEnabled(ctx, in, out)
}

func (h *userSrvHandler) RegisterUserByMobile(ctx context.Context, in *RegisterUserRequest, out *RegisterUserResponse) error {
	return h.UserSrvHandler.RegisterUserByMobile(ctx, in, out)
}

func (h *userSrvHandler) ResetPasswordByMobile(ctx context.Context, in *ResetPasswordRequest, out *ResetPasswordResponse) error {
	return h.UserSrvHandler.ResetPasswordByMobile(ctx, in, out)
}

func (h *userSrvHandler) SigninByMobile(ctx context.Context, in *SigninByMobileRequest, out *UserResponse) error {
	return h.UserSrvHandler.SigninByMobile(ctx, in, out)
}

func (h *userSrvHandler) FindUsersByClubID(ctx context.Context, in *FindUsersByClubIDRequest, out *UserListResponse) error {
	return h.UserSrvHandler.FindUsersByClubID(ctx, in, out)
}

func (h *userSrvHandler) CreateMember(ctx context.Context, in *CreateMemberRequest, out *Response) error {
	return h.UserSrvHandler.CreateMember(ctx, in, out)
}

func (h *userSrvHandler) FindUserClubProfileByID(ctx context.Context, in *FindUserClubProfileByIDRequest, out *UserClubProfileResponse) error {
	return h.UserSrvHandler.FindUserClubProfileByID(ctx, in, out)
}

func (h *userSrvHandler) FindUsersByOrganizationID(ctx context.Context, in *ByOrganizationIDRequest, out *UserListResponse) error {
	return h.UserSrvHandler.FindUsersByOrganizationID(ctx, in, out)
}

func (h *userSrvHandler) UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, out *Response) error {
	return h.UserSrvHandler.UpdateAvatar(ctx, in, out)
}
