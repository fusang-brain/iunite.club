// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/approved/approved.proto

/*
Package iunite_club_srv_core_approved is a generated protocol buffer package.

It is generated from these files:
	proto/approved/approved.proto

It has these top-level messages:
	CreateRequest
	ApprovedResponse
	Response
	ApprovedPB
	ApprovedFlowPB
	ListResponse
	ListRequest
	ListV2Request
	ListByPusherRequest
	ListByCountRequest
	DetailsRequest
	ExecuteRequest
*/
package iunite_club_srv_core_approved

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/struct"

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

// Client API for Approved service

type ApprovedService interface {
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	ListV2(ctx context.Context, in *ListV2Request, opts ...client.CallOption) (*ListResponse, error)
	ListByPusher(ctx context.Context, in *ListByPusherRequest, opts ...client.CallOption) (*ListResponse, error)
	WaitingExecuteList(ctx context.Context, in *ListByCountRequest, opts ...client.CallOption) (*ListResponse, error)
	Details(ctx context.Context, in *DetailsRequest, opts ...client.CallOption) (*ApprovedResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*ApprovedResponse, error)
}

type approvedService struct {
	c    client.Client
	name string
}

func NewApprovedService(name string, c client.Client) ApprovedService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iunite.club.srv.core.approved"
	}
	return &approvedService{
		c:    c,
		name: name,
	}
}

func (c *approvedService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Approved.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvedService) ListV2(ctx context.Context, in *ListV2Request, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Approved.ListV2", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvedService) ListByPusher(ctx context.Context, in *ListByPusherRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Approved.ListByPusher", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvedService) WaitingExecuteList(ctx context.Context, in *ListByCountRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Approved.WaitingExecuteList", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvedService) Details(ctx context.Context, in *DetailsRequest, opts ...client.CallOption) (*ApprovedResponse, error) {
	req := c.c.NewRequest(c.name, "Approved.Details", in)
	out := new(ApprovedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvedService) Execute(ctx context.Context, in *ExecuteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Approved.Execute", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvedService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*ApprovedResponse, error) {
	req := c.c.NewRequest(c.name, "Approved.Create", in)
	out := new(ApprovedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Approved service

type ApprovedHandler interface {
	List(context.Context, *ListRequest, *ListResponse) error
	ListV2(context.Context, *ListV2Request, *ListResponse) error
	ListByPusher(context.Context, *ListByPusherRequest, *ListResponse) error
	WaitingExecuteList(context.Context, *ListByCountRequest, *ListResponse) error
	Details(context.Context, *DetailsRequest, *ApprovedResponse) error
	Execute(context.Context, *ExecuteRequest, *Response) error
	Create(context.Context, *CreateRequest, *ApprovedResponse) error
}

func RegisterApprovedHandler(s server.Server, hdlr ApprovedHandler, opts ...server.HandlerOption) {
	type approved interface {
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		ListV2(ctx context.Context, in *ListV2Request, out *ListResponse) error
		ListByPusher(ctx context.Context, in *ListByPusherRequest, out *ListResponse) error
		WaitingExecuteList(ctx context.Context, in *ListByCountRequest, out *ListResponse) error
		Details(ctx context.Context, in *DetailsRequest, out *ApprovedResponse) error
		Execute(ctx context.Context, in *ExecuteRequest, out *Response) error
		Create(ctx context.Context, in *CreateRequest, out *ApprovedResponse) error
	}
	type Approved struct {
		approved
	}
	h := &approvedHandler{hdlr}
	s.Handle(s.NewHandler(&Approved{h}, opts...))
}

type approvedHandler struct {
	ApprovedHandler
}

func (h *approvedHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ApprovedHandler.List(ctx, in, out)
}

func (h *approvedHandler) ListV2(ctx context.Context, in *ListV2Request, out *ListResponse) error {
	return h.ApprovedHandler.ListV2(ctx, in, out)
}

func (h *approvedHandler) ListByPusher(ctx context.Context, in *ListByPusherRequest, out *ListResponse) error {
	return h.ApprovedHandler.ListByPusher(ctx, in, out)
}

func (h *approvedHandler) WaitingExecuteList(ctx context.Context, in *ListByCountRequest, out *ListResponse) error {
	return h.ApprovedHandler.WaitingExecuteList(ctx, in, out)
}

func (h *approvedHandler) Details(ctx context.Context, in *DetailsRequest, out *ApprovedResponse) error {
	return h.ApprovedHandler.Details(ctx, in, out)
}

func (h *approvedHandler) Execute(ctx context.Context, in *ExecuteRequest, out *Response) error {
	return h.ApprovedHandler.Execute(ctx, in, out)
}

func (h *approvedHandler) Create(ctx context.Context, in *CreateRequest, out *ApprovedResponse) error {
	return h.ApprovedHandler.Create(ctx, in, out)
}