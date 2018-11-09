// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/secruity/secruity.proto

/*
Package iunite_club_srv_user_secruity is a generated protocol buffer package.

It is generated from these files:
	proto/secruity/secruity.proto

It has these top-level messages:
	TokenRequest
	UserIDResponse
	SignupWithMobileRequest
	SignupResponse
	AuthRequest
	AuthResponse
*/
package iunite_club_srv_user_secruity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Client API for Secruity service

type SecruityService interface {
	Signin(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, opts ...client.CallOption) (*SignupResponse, error)
	GetUserIDFromToken(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*UserIDResponse, error)
}

type secruityService struct {
	c    client.Client
	name string
}

func NewSecruityService(name string, c client.Client) SecruityService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iunite.club.srv.user.secruity"
	}
	return &secruityService{
		c:    c,
		name: name,
	}
}

func (c *secruityService) Signin(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Secruity.Signin", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secruityService) SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, opts ...client.CallOption) (*SignupResponse, error) {
	req := c.c.NewRequest(c.name, "Secruity.SignupWithMobile", in)
	out := new(SignupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secruityService) GetUserIDFromToken(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*UserIDResponse, error) {
	req := c.c.NewRequest(c.name, "Secruity.GetUserIDFromToken", in)
	out := new(UserIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Secruity service

type SecruityHandler interface {
	Signin(context.Context, *AuthRequest, *AuthResponse) error
	SignupWithMobile(context.Context, *SignupWithMobileRequest, *SignupResponse) error
	GetUserIDFromToken(context.Context, *TokenRequest, *UserIDResponse) error
}

func RegisterSecruityHandler(s server.Server, hdlr SecruityHandler, opts ...server.HandlerOption) {
	type secruity interface {
		Signin(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, out *SignupResponse) error
		GetUserIDFromToken(ctx context.Context, in *TokenRequest, out *UserIDResponse) error
	}
	type Secruity struct {
		secruity
	}
	h := &secruityHandler{hdlr}
	s.Handle(s.NewHandler(&Secruity{h}, opts...))
}

type secruityHandler struct {
	SecruityHandler
}

func (h *secruityHandler) Signin(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.SecruityHandler.Signin(ctx, in, out)
}

func (h *secruityHandler) SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, out *SignupResponse) error {
	return h.SecruityHandler.SignupWithMobile(ctx, in, out)
}

func (h *secruityHandler) GetUserIDFromToken(ctx context.Context, in *TokenRequest, out *UserIDResponse) error {
	return h.SecruityHandler.GetUserIDFromToken(ctx, in, out)
}
