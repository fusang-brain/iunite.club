// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/job/job.proto

package iunite_club_srv_organization_job

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "iunite.club/services/organization/proto"
import proto1 "iunite.club/services/user/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserFromJobRequest struct {
	Users                []string `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	JobID                string   `protobuf:"bytes,2,opt,name=JobID,proto3" json:"JobID,omitempty"`
	ClubID               string   `protobuf:"bytes,3,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFromJobRequest) Reset()         { *m = UserFromJobRequest{} }
func (m *UserFromJobRequest) String() string { return proto.CompactTextString(m) }
func (*UserFromJobRequest) ProtoMessage()    {}
func (*UserFromJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{0}
}
func (m *UserFromJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFromJobRequest.Unmarshal(m, b)
}
func (m *UserFromJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFromJobRequest.Marshal(b, m, deterministic)
}
func (dst *UserFromJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFromJobRequest.Merge(dst, src)
}
func (m *UserFromJobRequest) XXX_Size() int {
	return xxx_messageInfo_UserFromJobRequest.Size(m)
}
func (m *UserFromJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFromJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserFromJobRequest proto.InternalMessageInfo

func (m *UserFromJobRequest) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserFromJobRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *UserFromJobRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type ListByClubIDRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	ClubID               string   `protobuf:"bytes,3,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListByClubIDRequest) Reset()         { *m = ListByClubIDRequest{} }
func (m *ListByClubIDRequest) String() string { return proto.CompactTextString(m) }
func (*ListByClubIDRequest) ProtoMessage()    {}
func (*ListByClubIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{1}
}
func (m *ListByClubIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByClubIDRequest.Unmarshal(m, b)
}
func (m *ListByClubIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByClubIDRequest.Marshal(b, m, deterministic)
}
func (dst *ListByClubIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByClubIDRequest.Merge(dst, src)
}
func (m *ListByClubIDRequest) XXX_Size() int {
	return xxx_messageInfo_ListByClubIDRequest.Size(m)
}
func (m *ListByClubIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByClubIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListByClubIDRequest proto.InternalMessageInfo

func (m *ListByClubIDRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListByClubIDRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListByClubIDRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type ListByJobIDRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	JobID                string   `protobuf:"bytes,3,opt,name=JobID,proto3" json:"JobID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListByJobIDRequest) Reset()         { *m = ListByJobIDRequest{} }
func (m *ListByJobIDRequest) String() string { return proto.CompactTextString(m) }
func (*ListByJobIDRequest) ProtoMessage()    {}
func (*ListByJobIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{2}
}
func (m *ListByJobIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByJobIDRequest.Unmarshal(m, b)
}
func (m *ListByJobIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByJobIDRequest.Marshal(b, m, deterministic)
}
func (dst *ListByJobIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByJobIDRequest.Merge(dst, src)
}
func (m *ListByJobIDRequest) XXX_Size() int {
	return xxx_messageInfo_ListByJobIDRequest.Size(m)
}
func (m *ListByJobIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByJobIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListByJobIDRequest proto.InternalMessageInfo

func (m *ListByJobIDRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListByJobIDRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListByJobIDRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

type UserListResponse struct {
	Users                []*proto1.User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	Total                int64          `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{3}
}
func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (dst *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(dst, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUsers() []*proto1.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserListResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type CreateJobRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ClubID               string   `protobuf:"bytes,2,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{4}
}
func (m *CreateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobRequest.Unmarshal(m, b)
}
func (m *CreateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobRequest.Marshal(b, m, deterministic)
}
func (dst *CreateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobRequest.Merge(dst, src)
}
func (m *CreateJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateJobRequest.Size(m)
}
func (m *CreateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobRequest proto.InternalMessageInfo

func (m *CreateJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateJobRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type CreateJobResponse struct {
	OK                   bool        `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	Job                  *proto2.Job `protobuf:"bytes,2,opt,name=Job,proto3" json:"Job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateJobResponse) Reset()         { *m = CreateJobResponse{} }
func (m *CreateJobResponse) String() string { return proto.CompactTextString(m) }
func (*CreateJobResponse) ProtoMessage()    {}
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{5}
}
func (m *CreateJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobResponse.Unmarshal(m, b)
}
func (m *CreateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobResponse.Marshal(b, m, deterministic)
}
func (dst *CreateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobResponse.Merge(dst, src)
}
func (m *CreateJobResponse) XXX_Size() int {
	return xxx_messageInfo_CreateJobResponse.Size(m)
}
func (m *CreateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobResponse proto.InternalMessageInfo

func (m *CreateJobResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *CreateJobResponse) GetJob() *proto2.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type UpdateJobRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ClubID               string   `protobuf:"bytes,3,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateJobRequest) Reset()         { *m = UpdateJobRequest{} }
func (m *UpdateJobRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateJobRequest) ProtoMessage()    {}
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{6}
}
func (m *UpdateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJobRequest.Unmarshal(m, b)
}
func (m *UpdateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJobRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJobRequest.Merge(dst, src)
}
func (m *UpdateJobRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateJobRequest.Size(m)
}
func (m *UpdateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJobRequest proto.InternalMessageInfo

func (m *UpdateJobRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UpdateJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateJobRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type UpdateJobResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateJobResponse) Reset()         { *m = UpdateJobResponse{} }
func (m *UpdateJobResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateJobResponse) ProtoMessage()    {}
func (*UpdateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{7}
}
func (m *UpdateJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJobResponse.Unmarshal(m, b)
}
func (m *UpdateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJobResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJobResponse.Merge(dst, src)
}
func (m *UpdateJobResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateJobResponse.Size(m)
}
func (m *UpdateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJobResponse proto.InternalMessageInfo

func (m *UpdateJobResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type RemoveJobRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveJobRequest) Reset()         { *m = RemoveJobRequest{} }
func (m *RemoveJobRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveJobRequest) ProtoMessage()    {}
func (*RemoveJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{8}
}
func (m *RemoveJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveJobRequest.Unmarshal(m, b)
}
func (m *RemoveJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveJobRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveJobRequest.Merge(dst, src)
}
func (m *RemoveJobRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveJobRequest.Size(m)
}
func (m *RemoveJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveJobRequest proto.InternalMessageInfo

func (m *RemoveJobRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RemoveJobResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveJobResponse) Reset()         { *m = RemoveJobResponse{} }
func (m *RemoveJobResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveJobResponse) ProtoMessage()    {}
func (*RemoveJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{9}
}
func (m *RemoveJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveJobResponse.Unmarshal(m, b)
}
func (m *RemoveJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveJobResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveJobResponse.Merge(dst, src)
}
func (m *RemoveJobResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveJobResponse.Size(m)
}
func (m *RemoveJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveJobResponse proto.InternalMessageInfo

func (m *RemoveJobResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type JobListRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	MaxFlag              string   `protobuf:"bytes,3,opt,name=MaxFlag,proto3" json:"MaxFlag,omitempty"`
	MinFlag              string   `protobuf:"bytes,4,opt,name=MinFlag,proto3" json:"MinFlag,omitempty"`
	OrganizationID       string   `protobuf:"bytes,5,opt,name=OrganizationID,proto3" json:"OrganizationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobListRequest) Reset()         { *m = JobListRequest{} }
func (m *JobListRequest) String() string { return proto.CompactTextString(m) }
func (*JobListRequest) ProtoMessage()    {}
func (*JobListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{10}
}
func (m *JobListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobListRequest.Unmarshal(m, b)
}
func (m *JobListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobListRequest.Marshal(b, m, deterministic)
}
func (dst *JobListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobListRequest.Merge(dst, src)
}
func (m *JobListRequest) XXX_Size() int {
	return xxx_messageInfo_JobListRequest.Size(m)
}
func (m *JobListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobListRequest proto.InternalMessageInfo

func (m *JobListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *JobListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *JobListRequest) GetMaxFlag() string {
	if m != nil {
		return m.MaxFlag
	}
	return ""
}

func (m *JobListRequest) GetMinFlag() string {
	if m != nil {
		return m.MinFlag
	}
	return ""
}

func (m *JobListRequest) GetOrganizationID() string {
	if m != nil {
		return m.OrganizationID
	}
	return ""
}

type JobListResponse struct {
	Jobs                 []*proto2.Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Total                int64         `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JobListResponse) Reset()         { *m = JobListResponse{} }
func (m *JobListResponse) String() string { return proto.CompactTextString(m) }
func (*JobListResponse) ProtoMessage()    {}
func (*JobListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_a74bcd39a0cacd4c, []int{11}
}
func (m *JobListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobListResponse.Unmarshal(m, b)
}
func (m *JobListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobListResponse.Marshal(b, m, deterministic)
}
func (dst *JobListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobListResponse.Merge(dst, src)
}
func (m *JobListResponse) XXX_Size() int {
	return xxx_messageInfo_JobListResponse.Size(m)
}
func (m *JobListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobListResponse proto.InternalMessageInfo

func (m *JobListResponse) GetJobs() []*proto2.Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *JobListResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*UserFromJobRequest)(nil), "iunite.club.srv.organization.job.UserFromJobRequest")
	proto.RegisterType((*ListByClubIDRequest)(nil), "iunite.club.srv.organization.job.ListByClubIDRequest")
	proto.RegisterType((*ListByJobIDRequest)(nil), "iunite.club.srv.organization.job.ListByJobIDRequest")
	proto.RegisterType((*UserListResponse)(nil), "iunite.club.srv.organization.job.UserListResponse")
	proto.RegisterType((*CreateJobRequest)(nil), "iunite.club.srv.organization.job.CreateJobRequest")
	proto.RegisterType((*CreateJobResponse)(nil), "iunite.club.srv.organization.job.CreateJobResponse")
	proto.RegisterType((*UpdateJobRequest)(nil), "iunite.club.srv.organization.job.UpdateJobRequest")
	proto.RegisterType((*UpdateJobResponse)(nil), "iunite.club.srv.organization.job.UpdateJobResponse")
	proto.RegisterType((*RemoveJobRequest)(nil), "iunite.club.srv.organization.job.RemoveJobRequest")
	proto.RegisterType((*RemoveJobResponse)(nil), "iunite.club.srv.organization.job.RemoveJobResponse")
	proto.RegisterType((*JobListRequest)(nil), "iunite.club.srv.organization.job.JobListRequest")
	proto.RegisterType((*JobListResponse)(nil), "iunite.club.srv.organization.job.JobListResponse")
}

func init() { proto.RegisterFile("proto/job/job.proto", fileDescriptor_job_a74bcd39a0cacd4c) }

var fileDescriptor_job_a74bcd39a0cacd4c = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0xc5, 0x36, 0x24, 0x65, 0xa2, 0x12, 0xb2, 0x89, 0x22, 0xe4, 0x13, 0x75, 0xa5, 0x2a, 0xea,
	0xc1, 0xa4, 0x90, 0x5c, 0x7a, 0xa8, 0x94, 0x80, 0x82, 0x20, 0x69, 0x88, 0x5c, 0xa2, 0x46, 0x3d,
	0x54, 0xb2, 0x61, 0x85, 0x8c, 0x8c, 0x97, 0xda, 0x0b, 0x6a, 0xaa, 0xf6, 0xd0, 0x43, 0x7f, 0xa1,
	0x7f, 0xda, 0x7b, 0xb5, 0xbb, 0xc6, 0x5a, 0x1b, 0x12, 0x4c, 0x0e, 0x39, 0x58, 0xf2, 0xcc, 0xce,
	0xbc, 0x79, 0xb3, 0xf3, 0x34, 0x0b, 0xfb, 0xd3, 0x80, 0x50, 0x52, 0x1b, 0x13, 0x87, 0x7d, 0x26,
	0xb7, 0x50, 0xd5, 0x9d, 0xf9, 0x2e, 0xc5, 0xe6, 0xc0, 0x9b, 0x39, 0x66, 0x18, 0xcc, 0x4d, 0x12,
	0x8c, 0x6c, 0xdf, 0xfd, 0x61, 0x53, 0x97, 0xf8, 0xe6, 0x98, 0x38, 0xfa, 0x7b, 0x29, 0xa2, 0x16,
	0xe2, 0x60, 0xee, 0x0e, 0x70, 0x58, 0x93, 0xc3, 0x6a, 0x02, 0x38, 0x91, 0xc9, 0x5d, 0xfa, 0xdb,
	0x95, 0xb9, 0xb3, 0x10, 0x07, 0x51, 0x0e, 0xfb, 0x15, 0xb1, 0xc6, 0x1d, 0xa0, 0xdb, 0x10, 0x07,
	0x17, 0x01, 0x99, 0x74, 0x89, 0x63, 0xe1, 0x6f, 0x33, 0x1c, 0x52, 0x74, 0x00, 0x05, 0xe6, 0x0d,
	0x2b, 0x4a, 0x55, 0x3b, 0x2a, 0x5a, 0xc2, 0x60, 0xde, 0x2e, 0x71, 0x3a, 0xad, 0x8a, 0x5a, 0x55,
	0x98, 0x97, 0x1b, 0xe8, 0x10, 0xb6, 0x9a, 0xde, 0x8c, 0xb9, 0x35, 0xee, 0x8e, 0x2c, 0xe3, 0x33,
	0xec, 0x5f, 0xb9, 0x21, 0x3d, 0xbf, 0x17, 0xf6, 0x02, 0x1a, 0x41, 0xfe, 0xc6, 0x1e, 0xe1, 0x8a,
	0x52, 0x55, 0x8e, 0x34, 0x8b, 0xff, 0x33, 0xe0, 0x2b, 0x77, 0xe2, 0x52, 0x0e, 0xac, 0x59, 0xc2,
	0x78, 0x10, 0xb8, 0x0f, 0x48, 0x00, 0xf3, 0xfa, 0x9b, 0xe3, 0xc6, 0x6d, 0x68, 0x52, 0x1b, 0xc6,
	0x17, 0x28, 0xb3, 0x2e, 0x19, 0xb2, 0x85, 0xc3, 0x29, 0xf1, 0x43, 0x8c, 0x8e, 0xe5, 0x6b, 0xd8,
	0xa9, 0xeb, 0x66, 0x7a, 0x6c, 0xfc, 0x22, 0x59, 0x88, 0x74, 0x45, 0x7d, 0x42, 0x6d, 0x6f, 0x51,
	0x91, 0x1b, 0xc6, 0x07, 0x28, 0x37, 0x03, 0x6c, 0x53, 0x2c, 0x5d, 0x31, 0x82, 0xfc, 0xb5, 0x3d,
	0x11, 0x7c, 0x8b, 0x16, 0xff, 0x97, 0x3a, 0x56, 0x13, 0x1d, 0xdf, 0xc1, 0x9e, 0x94, 0x1f, 0x91,
	0x2b, 0x81, 0xda, 0xbb, 0xe4, 0xe9, 0x2f, 0x2c, 0xb5, 0x77, 0x89, 0x1a, 0xa0, 0x75, 0x89, 0xc3,
	0x33, 0x77, 0xea, 0xaf, 0xcc, 0x47, 0x15, 0xc6, 0x70, 0x58, 0xb4, 0x71, 0x0d, 0xe5, 0xdb, 0xe9,
	0x30, 0xc9, 0xac, 0x04, 0x6a, 0xa7, 0x15, 0xf1, 0x52, 0x3b, 0xad, 0x98, 0xa9, 0xba, 0x92, 0x69,
	0x72, 0x36, 0xaf, 0x61, 0x4f, 0xc2, 0x5b, 0xcd, 0xd4, 0x30, 0xa0, 0x6c, 0xe1, 0x09, 0x99, 0x3f,
	0x52, 0x94, 0x01, 0x49, 0x31, 0x0f, 0x00, 0xfd, 0x55, 0xa0, 0xd4, 0x25, 0x8e, 0x98, 0xd9, 0xb2,
	0x0c, 0x0a, 0xab, 0x64, 0x50, 0x58, 0xc8, 0xa0, 0x02, 0xdb, 0x1f, 0xed, 0xef, 0x17, 0x9e, 0x3d,
	0x8a, 0x7a, 0x58, 0x98, 0xfc, 0xc4, 0xf5, 0xf9, 0x49, 0x3e, 0x3a, 0x11, 0x26, 0x7a, 0x03, 0xa5,
	0x9e, 0x74, 0x8f, 0x9d, 0x56, 0xa5, 0xc0, 0x03, 0x52, 0x5e, 0xe3, 0x2b, 0xec, 0xc6, 0xbc, 0x22,
	0xee, 0xa7, 0x90, 0x1f, 0x13, 0x67, 0x21, 0xa5, 0x0c, 0xf3, 0xe1, 0xe1, 0xab, 0x05, 0x55, 0xff,
	0xb7, 0xcd, 0x87, 0x8d, 0xe6, 0x50, 0x8c, 0x85, 0x81, 0xea, 0xe6, 0xba, 0xad, 0x62, 0xa6, 0x55,
	0xa8, 0x37, 0x36, 0xca, 0x11, 0xad, 0x18, 0x39, 0x56, 0x37, 0x1e, 0x73, 0x96, 0xba, 0x69, 0x8d,
	0x65, 0xa9, 0xbb, 0xa4, 0x23, 0x51, 0x37, 0x56, 0x45, 0x96, 0xba, 0x69, 0x99, 0x65, 0xa9, 0xbb,
	0x24, 0x3b, 0x23, 0x87, 0x7e, 0xc1, 0x41, 0x1b, 0xd3, 0x68, 0xa4, 0xe7, 0xf7, 0x37, 0x76, 0x80,
	0x7d, 0xda, 0x69, 0xa1, 0xe3, 0xf5, 0x70, 0x49, 0x7d, 0xea, 0xef, 0x36, 0xc8, 0x90, 0xca, 0xef,
	0xb6, 0x31, 0xe5, 0x1b, 0x26, 0xda, 0x7a, 0xe8, 0x64, 0x3d, 0xce, 0xf2, 0x92, 0xd4, 0xb3, 0x8c,
	0x2a, 0xb5, 0x04, 0x8d, 0x1c, 0xfa, 0xa3, 0xc0, 0x61, 0x1b, 0xd3, 0x33, 0xcf, 0x6b, 0xda, 0xfe,
	0x27, 0xec, 0xe1, 0x01, 0xc5, 0x43, 0xb1, 0xef, 0x4e, 0xb3, 0xd2, 0x48, 0x3c, 0x02, 0x4f, 0xe4,
	0xf1, 0x13, 0x5e, 0x9e, 0x0d, 0x45, 0xe1, 0x3e, 0x61, 0x0a, 0x38, 0xc9, 0x06, 0x93, 0x7c, 0xdc,
	0x9e, 0xaa, 0xbd, 0xdf, 0x0a, 0x20, 0xa1, 0x0d, 0xce, 0x20, 0x02, 0x7d, 0x56, 0x0e, 0xce, 0x16,
	0x7f, 0xb4, 0x1b, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x8b, 0x1d, 0xb8, 0x55, 0x08, 0x00,
	0x00,
}
