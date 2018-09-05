// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: secruity/proto/auth/auth.proto

/*
Package kit_iron_srv_secruity is a generated protocol buffer package.

It is generated from these files:
	secruity/proto/auth/auth.proto

It has these top-level messages:
	SignupWithMobileRequest
	SignupResponse
	AuthRequest
	AuthResponse
*/
package kit_iron_srv_secruity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Client API for Auth service

type AuthService interface {
	Signin(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, opts ...client.CallOption) (*SignupResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "kit.iron.srv.secruity"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Signin(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Signin", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, opts ...client.CallOption) (*SignupResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.SignupWithMobile", in)
	out := new(SignupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Signin(context.Context, *AuthRequest, *AuthResponse) error
	SignupWithMobile(context.Context, *SignupWithMobileRequest, *SignupResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Signin(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, out *SignupResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Signin(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.AuthHandler.Signin(ctx, in, out)
}

func (h *authHandler) SignupWithMobile(ctx context.Context, in *SignupWithMobileRequest, out *SignupResponse) error {
	return h.AuthHandler.SignupWithMobile(ctx, in, out)
}
