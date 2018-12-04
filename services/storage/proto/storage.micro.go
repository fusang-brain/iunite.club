// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/storage.proto

/*
Package iunite_club_srv_storage is a generated protocol buffer package.

It is generated from these files:
	proto/storage.proto

It has these top-level messages:
	File
	FilePB
	FileInfoRequest
	FileRequest
	FilesRequest
	FileResponse
	RepeatedFileResponse
*/
package iunite_club_srv_storage

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

// Client API for Storage service

type StorageService interface {
	SaveFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (*FileResponse, error)
	SaveFiles(ctx context.Context, in *FilesRequest, opts ...client.CallOption) (*RepeatedFileResponse, error)
	SaveFileInfo(ctx context.Context, in *FileInfoRequest, opts ...client.CallOption) (*FileResponse, error)
}

type storageService struct {
	c    client.Client
	name string
}

func NewStorageService(name string, c client.Client) StorageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iunite.club.srv.storage"
	}
	return &storageService{
		c:    c,
		name: name,
	}
}

func (c *storageService) SaveFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (*FileResponse, error) {
	req := c.c.NewRequest(c.name, "Storage.SaveFile", in)
	out := new(FileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) SaveFiles(ctx context.Context, in *FilesRequest, opts ...client.CallOption) (*RepeatedFileResponse, error) {
	req := c.c.NewRequest(c.name, "Storage.SaveFiles", in)
	out := new(RepeatedFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) SaveFileInfo(ctx context.Context, in *FileInfoRequest, opts ...client.CallOption) (*FileResponse, error) {
	req := c.c.NewRequest(c.name, "Storage.SaveFileInfo", in)
	out := new(FileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Storage service

type StorageHandler interface {
	SaveFile(context.Context, *FileRequest, *FileResponse) error
	SaveFiles(context.Context, *FilesRequest, *RepeatedFileResponse) error
	SaveFileInfo(context.Context, *FileInfoRequest, *FileResponse) error
}

func RegisterStorageHandler(s server.Server, hdlr StorageHandler, opts ...server.HandlerOption) {
	type storage interface {
		SaveFile(ctx context.Context, in *FileRequest, out *FileResponse) error
		SaveFiles(ctx context.Context, in *FilesRequest, out *RepeatedFileResponse) error
		SaveFileInfo(ctx context.Context, in *FileInfoRequest, out *FileResponse) error
	}
	type Storage struct {
		storage
	}
	h := &storageHandler{hdlr}
	s.Handle(s.NewHandler(&Storage{h}, opts...))
}

type storageHandler struct {
	StorageHandler
}

func (h *storageHandler) SaveFile(ctx context.Context, in *FileRequest, out *FileResponse) error {
	return h.StorageHandler.SaveFile(ctx, in, out)
}

func (h *storageHandler) SaveFiles(ctx context.Context, in *FilesRequest, out *RepeatedFileResponse) error {
	return h.StorageHandler.SaveFiles(ctx, in, out)
}

func (h *storageHandler) SaveFileInfo(ctx context.Context, in *FileInfoRequest, out *FileResponse) error {
	return h.StorageHandler.SaveFileInfo(ctx, in, out)
}
