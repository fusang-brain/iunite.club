// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization.proto

package iunite_club_srv_organization

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Organization struct {
	ID                   string       `protobuf:"bytes,9,opt,name=ID,proto3" json:"ID,omitempty"`
	Kind                 string       `protobuf:"bytes,1,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string       `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	SchoolID             string       `protobuf:"bytes,4,opt,name=SchoolID,proto3" json:"SchoolID,omitempty"`
	Description          string       `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	ParentID             string       `protobuf:"bytes,6,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	ClubProfile          *ClubProfile `protobuf:"bytes,7,opt,name=ClubProfile,proto3" json:"ClubProfile,omitempty"`
	Jobs                 []*Job       `protobuf:"bytes,8,rep,name=Jobs,proto3" json:"Jobs,omitempty"`
	CreatedAt            string       `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	ClubID               string       `protobuf:"bytes,12,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Organization) Reset()         { *m = Organization{} }
func (m *Organization) String() string { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()    {}
func (*Organization) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_a77b5e7552a276ba, []int{0}
}
func (m *Organization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organization.Unmarshal(m, b)
}
func (m *Organization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organization.Marshal(b, m, deterministic)
}
func (dst *Organization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organization.Merge(dst, src)
}
func (m *Organization) XXX_Size() int {
	return xxx_messageInfo_Organization.Size(m)
}
func (m *Organization) XXX_DiscardUnknown() {
	xxx_messageInfo_Organization.DiscardUnknown(m)
}

var xxx_messageInfo_Organization proto.InternalMessageInfo

func (m *Organization) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Organization) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organization) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Organization) GetSchoolID() string {
	if m != nil {
		return m.SchoolID
	}
	return ""
}

func (m *Organization) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Organization) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *Organization) GetClubProfile() *ClubProfile {
	if m != nil {
		return m.ClubProfile
	}
	return nil
}

func (m *Organization) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *Organization) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Organization) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Organization) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type ClubAccept struct {
	ID                   string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string        `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OrganizationID       string        `protobuf:"bytes,3,opt,name=OrganizationID,proto3" json:"OrganizationID,omitempty"`
	Organization         *Organization `protobuf:"bytes,4,opt,name=Organization,proto3" json:"Organization,omitempty"`
	State                int64         `protobuf:"varint,5,opt,name=State,proto3" json:"State,omitempty"`
	Kind                 int64         `protobuf:"varint,6,opt,name=Kind,proto3" json:"Kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClubAccept) Reset()         { *m = ClubAccept{} }
func (m *ClubAccept) String() string { return proto.CompactTextString(m) }
func (*ClubAccept) ProtoMessage()    {}
func (*ClubAccept) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_a77b5e7552a276ba, []int{1}
}
func (m *ClubAccept) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubAccept.Unmarshal(m, b)
}
func (m *ClubAccept) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubAccept.Marshal(b, m, deterministic)
}
func (dst *ClubAccept) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubAccept.Merge(dst, src)
}
func (m *ClubAccept) XXX_Size() int {
	return xxx_messageInfo_ClubAccept.Size(m)
}
func (m *ClubAccept) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubAccept.DiscardUnknown(m)
}

var xxx_messageInfo_ClubAccept proto.InternalMessageInfo

func (m *ClubAccept) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ClubAccept) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ClubAccept) GetOrganizationID() string {
	if m != nil {
		return m.OrganizationID
	}
	return ""
}

func (m *ClubAccept) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

func (m *ClubAccept) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *ClubAccept) GetKind() int64 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type Job struct {
	ID                   string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Slug                 string               `protobuf:"bytes,3,opt,name=Slug,proto3" json:"Slug,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	ClubID               string               `protobuf:"bytes,6,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_a77b5e7552a276ba, []int{2}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (dst *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(dst, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Job) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Job) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Job) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type Paperwork struct {
	FileID               string   `protobuf:"bytes,1,opt,name=FileID,proto3" json:"FileID,omitempty"`
	UploadAt             string   `protobuf:"bytes,2,opt,name=UploadAt,proto3" json:"UploadAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paperwork) Reset()         { *m = Paperwork{} }
func (m *Paperwork) String() string { return proto.CompactTextString(m) }
func (*Paperwork) ProtoMessage()    {}
func (*Paperwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_a77b5e7552a276ba, []int{3}
}
func (m *Paperwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paperwork.Unmarshal(m, b)
}
func (m *Paperwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paperwork.Marshal(b, m, deterministic)
}
func (dst *Paperwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paperwork.Merge(dst, src)
}
func (m *Paperwork) XXX_Size() int {
	return xxx_messageInfo_Paperwork.Size(m)
}
func (m *Paperwork) XXX_DiscardUnknown() {
	xxx_messageInfo_Paperwork.DiscardUnknown(m)
}

var xxx_messageInfo_Paperwork proto.InternalMessageInfo

func (m *Paperwork) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *Paperwork) GetUploadAt() string {
	if m != nil {
		return m.UploadAt
	}
	return ""
}

type ClubProfile struct {
	Logo                 string       `protobuf:"bytes,1,opt,name=Logo,proto3" json:"Logo,omitempty"`
	Scale                int32        `protobuf:"varint,2,opt,name=Scale,proto3" json:"Scale,omitempty"`
	Paperworks           []*Paperwork `protobuf:"bytes,3,rep,name=Paperworks,proto3" json:"Paperworks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClubProfile) Reset()         { *m = ClubProfile{} }
func (m *ClubProfile) String() string { return proto.CompactTextString(m) }
func (*ClubProfile) ProtoMessage()    {}
func (*ClubProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_a77b5e7552a276ba, []int{4}
}
func (m *ClubProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubProfile.Unmarshal(m, b)
}
func (m *ClubProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubProfile.Marshal(b, m, deterministic)
}
func (dst *ClubProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubProfile.Merge(dst, src)
}
func (m *ClubProfile) XXX_Size() int {
	return xxx_messageInfo_ClubProfile.Size(m)
}
func (m *ClubProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ClubProfile proto.InternalMessageInfo

func (m *ClubProfile) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *ClubProfile) GetScale() int32 {
	if m != nil {
		return m.Scale
	}
	return 0
}

func (m *ClubProfile) GetPaperworks() []*Paperwork {
	if m != nil {
		return m.Paperworks
	}
	return nil
}

type UserClubProfile struct {
	ID                   string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	State                int64         `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	UserID               string        `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	User                 *proto1.User  `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
	IsCreator            bool          `protobuf:"varint,5,opt,name=IsCreator,proto3" json:"IsCreator,omitempty"`
	IsMaster             bool          `protobuf:"varint,6,opt,name=IsMaster,proto3" json:"IsMaster,omitempty"`
	JoinTime             string        `protobuf:"bytes,7,opt,name=JoinTime,proto3" json:"JoinTime,omitempty"`
	LeaveTime            string        `protobuf:"bytes,8,opt,name=LeaveTime,proto3" json:"LeaveTime,omitempty"`
	JobID                string        `protobuf:"bytes,9,opt,name=JobID,proto3" json:"JobID,omitempty"`
	Job                  *Job          `protobuf:"bytes,10,opt,name=Job,proto3" json:"Job,omitempty"`
	DepartmentID         string        `protobuf:"bytes,11,opt,name=DepartmentID,proto3" json:"DepartmentID,omitempty"`
	Department           *Organization `protobuf:"bytes,12,opt,name=Department,proto3" json:"Department,omitempty"`
	CreatedAt            string        `protobuf:"bytes,13,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,14,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	OrganizationID       string        `protobuf:"bytes,15,opt,name=OrganizationID,proto3" json:"OrganizationID,omitempty"`
	Organization         *Organization `protobuf:"bytes,16,opt,name=Organization,proto3" json:"Organization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserClubProfile) Reset()         { *m = UserClubProfile{} }
func (m *UserClubProfile) String() string { return proto.CompactTextString(m) }
func (*UserClubProfile) ProtoMessage()    {}
func (*UserClubProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_a77b5e7552a276ba, []int{5}
}
func (m *UserClubProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserClubProfile.Unmarshal(m, b)
}
func (m *UserClubProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserClubProfile.Marshal(b, m, deterministic)
}
func (dst *UserClubProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserClubProfile.Merge(dst, src)
}
func (m *UserClubProfile) XXX_Size() int {
	return xxx_messageInfo_UserClubProfile.Size(m)
}
func (m *UserClubProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserClubProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserClubProfile proto.InternalMessageInfo

func (m *UserClubProfile) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UserClubProfile) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *UserClubProfile) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserClubProfile) GetUser() *proto1.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserClubProfile) GetIsCreator() bool {
	if m != nil {
		return m.IsCreator
	}
	return false
}

func (m *UserClubProfile) GetIsMaster() bool {
	if m != nil {
		return m.IsMaster
	}
	return false
}

func (m *UserClubProfile) GetJoinTime() string {
	if m != nil {
		return m.JoinTime
	}
	return ""
}

func (m *UserClubProfile) GetLeaveTime() string {
	if m != nil {
		return m.LeaveTime
	}
	return ""
}

func (m *UserClubProfile) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *UserClubProfile) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *UserClubProfile) GetDepartmentID() string {
	if m != nil {
		return m.DepartmentID
	}
	return ""
}

func (m *UserClubProfile) GetDepartment() *Organization {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *UserClubProfile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserClubProfile) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UserClubProfile) GetOrganizationID() string {
	if m != nil {
		return m.OrganizationID
	}
	return ""
}

func (m *UserClubProfile) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

func init() {
	proto.RegisterType((*Organization)(nil), "iunite.club.srv.organization.Organization")
	proto.RegisterType((*ClubAccept)(nil), "iunite.club.srv.organization.ClubAccept")
	proto.RegisterType((*Job)(nil), "iunite.club.srv.organization.Job")
	proto.RegisterType((*Paperwork)(nil), "iunite.club.srv.organization.Paperwork")
	proto.RegisterType((*ClubProfile)(nil), "iunite.club.srv.organization.ClubProfile")
	proto.RegisterType((*UserClubProfile)(nil), "iunite.club.srv.organization.UserClubProfile")
}

func init() {
	proto.RegisterFile("proto/organization.proto", fileDescriptor_organization_a77b5e7552a276ba)
}

var fileDescriptor_organization_a77b5e7552a276ba = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0x24, 0x24, 0x93, 0xd2, 0xa2, 0x15, 0xaa, 0x56, 0x51, 0x25, 0x42, 0x3e, 0xa0,
	0xf4, 0xc3, 0x91, 0x5a, 0x21, 0xf5, 0x0f, 0x55, 0xb5, 0x40, 0x4e, 0x4b, 0xa9, 0x5c, 0x7a, 0x80,
	0xb5, 0xbb, 0x0d, 0x16, 0x4e, 0xd6, 0x5a, 0xaf, 0x8b, 0xc4, 0x17, 0x07, 0xe0, 0x62, 0x5c, 0x80,
	0x03, 0x70, 0x12, 0xb4, 0xb3, 0xce, 0x66, 0x9d, 0xa2, 0x50, 0xc4, 0xdf, 0xcc, 0x9b, 0x19, 0x7b,
	0x67, 0xde, 0xbc, 0x01, 0x5a, 0x48, 0xa1, 0xc4, 0x44, 0xc8, 0x19, 0x5b, 0x64, 0x5f, 0x99, 0xca,
	0xc4, 0x22, 0x40, 0x88, 0xec, 0x65, 0xd5, 0x22, 0x53, 0x3c, 0x48, 0xf3, 0x2a, 0x09, 0x4a, 0x79,
	0x17, 0xb8, 0x39, 0xc3, 0x67, 0x33, 0x21, 0x66, 0x39, 0x9f, 0x60, 0x6e, 0x52, 0xdd, 0x4e, 0x54,
	0x36, 0xe7, 0xa5, 0x62, 0xf3, 0xc2, 0x94, 0x0f, 0x0f, 0x9c, 0xf2, 0x49, 0xc9, 0xe5, 0x5d, 0x96,
	0xf2, 0x72, 0x52, 0x95, 0x5c, 0x9a, 0x1a, 0x34, 0x4d, 0xee, 0xf8, 0xbb, 0x0f, 0x5b, 0x1f, 0x9c,
	0xaf, 0x93, 0x6d, 0x68, 0x45, 0x21, 0xed, 0x8f, 0xbc, 0xfd, 0x7e, 0xdc, 0x8a, 0x42, 0x42, 0xa0,
	0x7d, 0x96, 0x2d, 0x6e, 0xa8, 0x87, 0x08, 0xda, 0x1a, 0xbb, 0x60, 0x73, 0x4e, 0x5b, 0x06, 0xd3,
	0xb6, 0xc6, 0xae, 0xf2, 0x6a, 0x46, 0x7d, 0x83, 0x69, 0x9b, 0x0c, 0xa1, 0x77, 0x95, 0x7e, 0x12,
	0x22, 0x8f, 0x42, 0xda, 0x46, 0xdc, 0xfa, 0x64, 0x04, 0x83, 0x90, 0x97, 0xa9, 0xcc, 0x0a, 0xfd,
	0x5b, 0xda, 0xc1, 0xb0, 0x0b, 0xe9, 0xea, 0x4b, 0x26, 0xf9, 0x42, 0x45, 0x21, 0xed, 0x9a, 0xea,
	0xa5, 0x4f, 0xce, 0x60, 0x70, 0x9a, 0x57, 0xc9, 0xa5, 0x14, 0xb7, 0x59, 0xce, 0xe9, 0xa3, 0x91,
	0xb7, 0x3f, 0x38, 0x7c, 0x15, 0x6c, 0x9a, 0x5b, 0xe0, 0x14, 0xc4, 0x6e, 0x35, 0x79, 0x0d, 0xed,
	0xa9, 0x48, 0x4a, 0xda, 0x1b, 0xf9, 0xfb, 0x83, 0xc3, 0xe7, 0x9b, 0xbf, 0x32, 0x15, 0x49, 0x8c,
	0xe9, 0x64, 0x0f, 0xfa, 0xa7, 0x92, 0x33, 0xc5, 0x6f, 0x4e, 0x14, 0x05, 0x7c, 0xe0, 0x0a, 0xd0,
	0xd1, 0xeb, 0xe2, 0xa6, 0x8e, 0x0e, 0x4c, 0xd4, 0x02, 0x64, 0x17, 0xba, 0xfa, 0x05, 0x51, 0x48,
	0xb7, 0x30, 0x54, 0x7b, 0xe3, 0x9f, 0x1e, 0x80, 0x36, 0x4f, 0xd2, 0x94, 0x17, 0xaa, 0x26, 0xc3,
	0xb3, 0x64, 0xec, 0x42, 0xf7, 0xba, 0xe4, 0x32, 0x0a, 0xeb, 0xd1, 0xd7, 0x1e, 0x79, 0x01, 0xdb,
	0x2e, 0x89, 0x51, 0x58, 0xd3, 0xb0, 0x86, 0x92, 0x8b, 0x26, 0xd9, 0x48, 0xca, 0xe0, 0xf0, 0x60,
	0x73, 0xc7, 0x6e, 0x45, 0xdc, 0x5c, 0x96, 0xa7, 0xd0, 0xb9, 0x52, 0x4c, 0x71, 0xa4, 0xcf, 0x8f,
	0x8d, 0x63, 0x57, 0xa6, 0x8b, 0x20, 0xda, 0xe3, 0x1f, 0x1e, 0xf8, 0x53, 0x91, 0xdc, 0xeb, 0xe8,
	0xa1, 0xab, 0x74, 0xec, 0x0e, 0xdb, 0x3c, 0x7b, 0x18, 0x18, 0x21, 0x04, 0x4b, 0x21, 0x04, 0x1f,
	0x97, 0x42, 0x70, 0x89, 0x38, 0x76, 0x89, 0xe8, 0xfc, 0xbd, 0xf2, 0x4f, 0x24, 0x75, 0x1b, 0x24,
	0xbd, 0x81, 0xfe, 0x25, 0x2b, 0xb8, 0xfc, 0x22, 0xe4, 0x67, 0x9d, 0xf4, 0x36, 0xcb, 0xb9, 0x6d,
	0xaa, 0xf6, 0xf4, 0xf6, 0x5e, 0x17, 0xb9, 0x60, 0xfa, 0xaf, 0xa6, 0x39, 0xeb, 0x8f, 0xbf, 0x79,
	0x8d, 0xf5, 0xd5, 0x0d, 0x9f, 0x8b, 0x99, 0x58, 0x6a, 0x4c, 0xdb, 0x38, 0xda, 0x94, 0xe5, 0x66,
	0x32, 0x9d, 0xd8, 0x38, 0xe4, 0x1d, 0x80, 0xfd, 0x75, 0x49, 0x7d, 0x5c, 0xd8, 0x97, 0x9b, 0xe9,
	0xb3, 0xf9, 0xb1, 0x53, 0x3a, 0xfe, 0xd5, 0x86, 0x1d, 0xbd, 0x3c, 0xee, 0x33, 0xd6, 0xb9, 0xb1,
	0xec, 0xb6, 0x5c, 0x76, 0x57, 0x3b, 0xe8, 0x37, 0x76, 0x30, 0x80, 0xb6, 0xb6, 0x2c, 0x39, 0xeb,
	0x8f, 0xc2, 0xa3, 0xa3, 0x33, 0x62, 0xcc, 0xd3, 0x02, 0x89, 0x4a, 0xa4, 0x49, 0x48, 0xe4, 0xa5,
	0x17, 0xaf, 0x00, 0x3d, 0xbe, 0xa8, 0x7c, 0xcf, 0x4a, 0xc5, 0x25, 0x4e, 0xbf, 0x17, 0x5b, 0x5f,
	0xc7, 0xa6, 0x22, 0x5b, 0x68, 0xce, 0x50, 0xf9, 0xfd, 0xd8, 0xfa, 0xfa, 0xab, 0xe7, 0x9c, 0xdd,
	0x71, 0x0c, 0xf6, 0x8c, 0xec, 0x2c, 0xa0, 0x3b, 0x9a, 0x8a, 0xc4, 0xde, 0x37, 0xe3, 0x90, 0x23,
	0x5c, 0x4d, 0x94, 0xf0, 0x83, 0xe4, 0x8f, 0x8b, 0x3c, 0x86, 0xad, 0x90, 0x17, 0x4c, 0xaa, 0xb9,
	0xb9, 0x50, 0x46, 0xe2, 0x0d, 0x8c, 0x4c, 0x01, 0x56, 0x3e, 0x2a, 0xfd, 0xdf, 0xc4, 0xe6, 0x54,
	0x37, 0xaf, 0xcd, 0xe3, 0x8d, 0xd7, 0x66, 0x7b, 0xfd, 0xda, 0xdc, 0x3f, 0x0f, 0x3b, 0x0f, 0x3a,
	0x0f, 0x4f, 0xfe, 0xef, 0x3c, 0x24, 0x5d, 0xd4, 0xd7, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0x75, 0xf0, 0x86, 0xea, 0x06, 0x00, 0x00,
}
