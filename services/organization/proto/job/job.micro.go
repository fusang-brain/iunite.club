// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/job/job.proto

/*
Package iunite_club_srv_organization_job is a generated protocol buffer package.

It is generated from these files:
	proto/job/job.proto

It has these top-level messages:
	CreateJobRequest
	CreateJobResponse
	UpdateJobRequest
	UpdateJobResponse
	RemoveJobRequest
	RemoveJobResponse
	JobListRequest
	JobListResponse
*/
package iunite_club_srv_organization_job

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "iunite.club/services/organization/proto"

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

// Client API for Job service

type JobService interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...client.CallOption) (*CreateJobResponse, error)
	UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...client.CallOption) (*UpdateJobResponse, error)
	RemoveJob(ctx context.Context, in *RemoveJobRequest, opts ...client.CallOption) (*RemoveJobResponse, error)
	GetJobListByParentID(ctx context.Context, in *JobListRequest, opts ...client.CallOption) (*JobListResponse, error)
}

type jobService struct {
	c    client.Client
	name string
}

func NewJobService(name string, c client.Client) JobService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iunite.club.srv.organization.job"
	}
	return &jobService{
		c:    c,
		name: name,
	}
}

func (c *jobService) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...client.CallOption) (*CreateJobResponse, error) {
	req := c.c.NewRequest(c.name, "Job.CreateJob", in)
	out := new(CreateJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...client.CallOption) (*UpdateJobResponse, error) {
	req := c.c.NewRequest(c.name, "Job.UpdateJob", in)
	out := new(UpdateJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) RemoveJob(ctx context.Context, in *RemoveJobRequest, opts ...client.CallOption) (*RemoveJobResponse, error) {
	req := c.c.NewRequest(c.name, "Job.RemoveJob", in)
	out := new(RemoveJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobService) GetJobListByParentID(ctx context.Context, in *JobListRequest, opts ...client.CallOption) (*JobListResponse, error) {
	req := c.c.NewRequest(c.name, "Job.GetJobListByParentID", in)
	out := new(JobListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobHandler interface {
	CreateJob(context.Context, *CreateJobRequest, *CreateJobResponse) error
	UpdateJob(context.Context, *UpdateJobRequest, *UpdateJobResponse) error
	RemoveJob(context.Context, *RemoveJobRequest, *RemoveJobResponse) error
	GetJobListByParentID(context.Context, *JobListRequest, *JobListResponse) error
}

func RegisterJobHandler(s server.Server, hdlr JobHandler, opts ...server.HandlerOption) {
	type job interface {
		CreateJob(ctx context.Context, in *CreateJobRequest, out *CreateJobResponse) error
		UpdateJob(ctx context.Context, in *UpdateJobRequest, out *UpdateJobResponse) error
		RemoveJob(ctx context.Context, in *RemoveJobRequest, out *RemoveJobResponse) error
		GetJobListByParentID(ctx context.Context, in *JobListRequest, out *JobListResponse) error
	}
	type Job struct {
		job
	}
	h := &jobHandler{hdlr}
	s.Handle(s.NewHandler(&Job{h}, opts...))
}

type jobHandler struct {
	JobHandler
}

func (h *jobHandler) CreateJob(ctx context.Context, in *CreateJobRequest, out *CreateJobResponse) error {
	return h.JobHandler.CreateJob(ctx, in, out)
}

func (h *jobHandler) UpdateJob(ctx context.Context, in *UpdateJobRequest, out *UpdateJobResponse) error {
	return h.JobHandler.UpdateJob(ctx, in, out)
}

func (h *jobHandler) RemoveJob(ctx context.Context, in *RemoveJobRequest, out *RemoveJobResponse) error {
	return h.JobHandler.RemoveJob(ctx, in, out)
}

func (h *jobHandler) GetJobListByParentID(ctx context.Context, in *JobListRequest, out *JobListResponse) error {
	return h.JobHandler.GetJobListByParentID(ctx, in, out)
}
