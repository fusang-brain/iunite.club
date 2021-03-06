// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/department/department.proto

package iunite_club_srv_organization_department

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

type SearchDepartmentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchDepartmentRequest) Reset()         { *m = SearchDepartmentRequest{} }
func (m *SearchDepartmentRequest) String() string { return proto.CompactTextString(m) }
func (*SearchDepartmentRequest) ProtoMessage()    {}
func (*SearchDepartmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{0}
}
func (m *SearchDepartmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchDepartmentRequest.Unmarshal(m, b)
}
func (m *SearchDepartmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchDepartmentRequest.Marshal(b, m, deterministic)
}
func (dst *SearchDepartmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchDepartmentRequest.Merge(dst, src)
}
func (m *SearchDepartmentRequest) XXX_Size() int {
	return xxx_messageInfo_SearchDepartmentRequest.Size(m)
}
func (m *SearchDepartmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchDepartmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchDepartmentRequest proto.InternalMessageInfo

func (m *SearchDepartmentRequest) GetName() string {
	if m != nil {
		return m.Name
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
	return fileDescriptor_department_9b874b28ea07144d, []int{1}
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

type UserFromDepartmentRequest struct {
	Users                []string `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	DepartmentID         string   `protobuf:"bytes,2,opt,name=DepartmentID,proto3" json:"DepartmentID,omitempty"`
	ClubID               string   `protobuf:"bytes,3,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFromDepartmentRequest) Reset()         { *m = UserFromDepartmentRequest{} }
func (m *UserFromDepartmentRequest) String() string { return proto.CompactTextString(m) }
func (*UserFromDepartmentRequest) ProtoMessage()    {}
func (*UserFromDepartmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{2}
}
func (m *UserFromDepartmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFromDepartmentRequest.Unmarshal(m, b)
}
func (m *UserFromDepartmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFromDepartmentRequest.Marshal(b, m, deterministic)
}
func (dst *UserFromDepartmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFromDepartmentRequest.Merge(dst, src)
}
func (m *UserFromDepartmentRequest) XXX_Size() int {
	return xxx_messageInfo_UserFromDepartmentRequest.Size(m)
}
func (m *UserFromDepartmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFromDepartmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserFromDepartmentRequest proto.InternalMessageInfo

func (m *UserFromDepartmentRequest) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserFromDepartmentRequest) GetDepartmentID() string {
	if m != nil {
		return m.DepartmentID
	}
	return ""
}

func (m *UserFromDepartmentRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type ListByDepartmentIDRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	DepartmentID         string   `protobuf:"bytes,3,opt,name=DepartmentID,proto3" json:"DepartmentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListByDepartmentIDRequest) Reset()         { *m = ListByDepartmentIDRequest{} }
func (m *ListByDepartmentIDRequest) String() string { return proto.CompactTextString(m) }
func (*ListByDepartmentIDRequest) ProtoMessage()    {}
func (*ListByDepartmentIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{3}
}
func (m *ListByDepartmentIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByDepartmentIDRequest.Unmarshal(m, b)
}
func (m *ListByDepartmentIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByDepartmentIDRequest.Marshal(b, m, deterministic)
}
func (dst *ListByDepartmentIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByDepartmentIDRequest.Merge(dst, src)
}
func (m *ListByDepartmentIDRequest) XXX_Size() int {
	return xxx_messageInfo_ListByDepartmentIDRequest.Size(m)
}
func (m *ListByDepartmentIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByDepartmentIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListByDepartmentIDRequest proto.InternalMessageInfo

func (m *ListByDepartmentIDRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListByDepartmentIDRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListByDepartmentIDRequest) GetDepartmentID() string {
	if m != nil {
		return m.DepartmentID
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
	return fileDescriptor_department_9b874b28ea07144d, []int{4}
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

type GetDepartmentWithIDRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDepartmentWithIDRequest) Reset()         { *m = GetDepartmentWithIDRequest{} }
func (m *GetDepartmentWithIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetDepartmentWithIDRequest) ProtoMessage()    {}
func (*GetDepartmentWithIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{5}
}
func (m *GetDepartmentWithIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDepartmentWithIDRequest.Unmarshal(m, b)
}
func (m *GetDepartmentWithIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDepartmentWithIDRequest.Marshal(b, m, deterministic)
}
func (dst *GetDepartmentWithIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDepartmentWithIDRequest.Merge(dst, src)
}
func (m *GetDepartmentWithIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetDepartmentWithIDRequest.Size(m)
}
func (m *GetDepartmentWithIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDepartmentWithIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDepartmentWithIDRequest proto.InternalMessageInfo

func (m *GetDepartmentWithIDRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type DepartmentResponse struct {
	Department           *proto2.Organization `protobuf:"bytes,1,opt,name=Department,proto3" json:"Department,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DepartmentResponse) Reset()         { *m = DepartmentResponse{} }
func (m *DepartmentResponse) String() string { return proto.CompactTextString(m) }
func (*DepartmentResponse) ProtoMessage()    {}
func (*DepartmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{6}
}
func (m *DepartmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepartmentResponse.Unmarshal(m, b)
}
func (m *DepartmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepartmentResponse.Marshal(b, m, deterministic)
}
func (dst *DepartmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepartmentResponse.Merge(dst, src)
}
func (m *DepartmentResponse) XXX_Size() int {
	return xxx_messageInfo_DepartmentResponse.Size(m)
}
func (m *DepartmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DepartmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DepartmentResponse proto.InternalMessageInfo

func (m *DepartmentResponse) GetDepartment() *proto2.Organization {
	if m != nil {
		return m.Department
	}
	return nil
}

type CreateDepartmentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ParentID             string   `protobuf:"bytes,2,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDepartmentRequest) Reset()         { *m = CreateDepartmentRequest{} }
func (m *CreateDepartmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDepartmentRequest) ProtoMessage()    {}
func (*CreateDepartmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{7}
}
func (m *CreateDepartmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDepartmentRequest.Unmarshal(m, b)
}
func (m *CreateDepartmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDepartmentRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDepartmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDepartmentRequest.Merge(dst, src)
}
func (m *CreateDepartmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDepartmentRequest.Size(m)
}
func (m *CreateDepartmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDepartmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDepartmentRequest proto.InternalMessageInfo

func (m *CreateDepartmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDepartmentRequest) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *CreateDepartmentRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateDepartmentResponse struct {
	Department           *proto2.Organization `protobuf:"bytes,1,opt,name=Department,proto3" json:"Department,omitempty"`
	OK                   bool                 `protobuf:"varint,2,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateDepartmentResponse) Reset()         { *m = CreateDepartmentResponse{} }
func (m *CreateDepartmentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDepartmentResponse) ProtoMessage()    {}
func (*CreateDepartmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{8}
}
func (m *CreateDepartmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDepartmentResponse.Unmarshal(m, b)
}
func (m *CreateDepartmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDepartmentResponse.Marshal(b, m, deterministic)
}
func (dst *CreateDepartmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDepartmentResponse.Merge(dst, src)
}
func (m *CreateDepartmentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDepartmentResponse.Size(m)
}
func (m *CreateDepartmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDepartmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDepartmentResponse proto.InternalMessageInfo

func (m *CreateDepartmentResponse) GetDepartment() *proto2.Organization {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *CreateDepartmentResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type UpdateDepartmentRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ParentID             string   `protobuf:"bytes,3,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDepartmentRequest) Reset()         { *m = UpdateDepartmentRequest{} }
func (m *UpdateDepartmentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDepartmentRequest) ProtoMessage()    {}
func (*UpdateDepartmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{9}
}
func (m *UpdateDepartmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDepartmentRequest.Unmarshal(m, b)
}
func (m *UpdateDepartmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDepartmentRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateDepartmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDepartmentRequest.Merge(dst, src)
}
func (m *UpdateDepartmentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDepartmentRequest.Size(m)
}
func (m *UpdateDepartmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDepartmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDepartmentRequest proto.InternalMessageInfo

func (m *UpdateDepartmentRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UpdateDepartmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateDepartmentRequest) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *UpdateDepartmentRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateDepartmentResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDepartmentResponse) Reset()         { *m = UpdateDepartmentResponse{} }
func (m *UpdateDepartmentResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateDepartmentResponse) ProtoMessage()    {}
func (*UpdateDepartmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{10}
}
func (m *UpdateDepartmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDepartmentResponse.Unmarshal(m, b)
}
func (m *UpdateDepartmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDepartmentResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateDepartmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDepartmentResponse.Merge(dst, src)
}
func (m *UpdateDepartmentResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateDepartmentResponse.Size(m)
}
func (m *UpdateDepartmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDepartmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDepartmentResponse proto.InternalMessageInfo

func (m *UpdateDepartmentResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type RemoveDepartmentRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDepartmentRequest) Reset()         { *m = RemoveDepartmentRequest{} }
func (m *RemoveDepartmentRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveDepartmentRequest) ProtoMessage()    {}
func (*RemoveDepartmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{11}
}
func (m *RemoveDepartmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDepartmentRequest.Unmarshal(m, b)
}
func (m *RemoveDepartmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDepartmentRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveDepartmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDepartmentRequest.Merge(dst, src)
}
func (m *RemoveDepartmentRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveDepartmentRequest.Size(m)
}
func (m *RemoveDepartmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDepartmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDepartmentRequest proto.InternalMessageInfo

func (m *RemoveDepartmentRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RemoveDepartmentResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDepartmentResponse) Reset()         { *m = RemoveDepartmentResponse{} }
func (m *RemoveDepartmentResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveDepartmentResponse) ProtoMessage()    {}
func (*RemoveDepartmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{12}
}
func (m *RemoveDepartmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDepartmentResponse.Unmarshal(m, b)
}
func (m *RemoveDepartmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDepartmentResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveDepartmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDepartmentResponse.Merge(dst, src)
}
func (m *RemoveDepartmentResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveDepartmentResponse.Size(m)
}
func (m *RemoveDepartmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDepartmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDepartmentResponse proto.InternalMessageInfo

func (m *RemoveDepartmentResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type DepartmentListByParentIDRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	MaxFlag              string   `protobuf:"bytes,3,opt,name=MaxFlag,proto3" json:"MaxFlag,omitempty"`
	MinFlag              string   `protobuf:"bytes,4,opt,name=MinFlag,proto3" json:"MinFlag,omitempty"`
	ParentID             string   `protobuf:"bytes,5,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Search               string   `protobuf:"bytes,6,opt,name=Search,proto3" json:"Search,omitempty"`
	Spread               bool     `protobuf:"varint,7,opt,name=Spread,proto3" json:"Spread,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DepartmentListByParentIDRequest) Reset()         { *m = DepartmentListByParentIDRequest{} }
func (m *DepartmentListByParentIDRequest) String() string { return proto.CompactTextString(m) }
func (*DepartmentListByParentIDRequest) ProtoMessage()    {}
func (*DepartmentListByParentIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{13}
}
func (m *DepartmentListByParentIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepartmentListByParentIDRequest.Unmarshal(m, b)
}
func (m *DepartmentListByParentIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepartmentListByParentIDRequest.Marshal(b, m, deterministic)
}
func (dst *DepartmentListByParentIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepartmentListByParentIDRequest.Merge(dst, src)
}
func (m *DepartmentListByParentIDRequest) XXX_Size() int {
	return xxx_messageInfo_DepartmentListByParentIDRequest.Size(m)
}
func (m *DepartmentListByParentIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DepartmentListByParentIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DepartmentListByParentIDRequest proto.InternalMessageInfo

func (m *DepartmentListByParentIDRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *DepartmentListByParentIDRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DepartmentListByParentIDRequest) GetMaxFlag() string {
	if m != nil {
		return m.MaxFlag
	}
	return ""
}

func (m *DepartmentListByParentIDRequest) GetMinFlag() string {
	if m != nil {
		return m.MinFlag
	}
	return ""
}

func (m *DepartmentListByParentIDRequest) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *DepartmentListByParentIDRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *DepartmentListByParentIDRequest) GetSpread() bool {
	if m != nil {
		return m.Spread
	}
	return false
}

type DepartmentListResponse struct {
	Departments          []*proto2.Organization `protobuf:"bytes,1,rep,name=Departments,proto3" json:"Departments,omitempty"`
	Total                int64                  `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DepartmentListResponse) Reset()         { *m = DepartmentListResponse{} }
func (m *DepartmentListResponse) String() string { return proto.CompactTextString(m) }
func (*DepartmentListResponse) ProtoMessage()    {}
func (*DepartmentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_department_9b874b28ea07144d, []int{14}
}
func (m *DepartmentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepartmentListResponse.Unmarshal(m, b)
}
func (m *DepartmentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepartmentListResponse.Marshal(b, m, deterministic)
}
func (dst *DepartmentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepartmentListResponse.Merge(dst, src)
}
func (m *DepartmentListResponse) XXX_Size() int {
	return xxx_messageInfo_DepartmentListResponse.Size(m)
}
func (m *DepartmentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DepartmentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DepartmentListResponse proto.InternalMessageInfo

func (m *DepartmentListResponse) GetDepartments() []*proto2.Organization {
	if m != nil {
		return m.Departments
	}
	return nil
}

func (m *DepartmentListResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchDepartmentRequest)(nil), "iunite.club.srv.organization.department.SearchDepartmentRequest")
	proto.RegisterType((*UserListResponse)(nil), "iunite.club.srv.organization.department.UserListResponse")
	proto.RegisterType((*UserFromDepartmentRequest)(nil), "iunite.club.srv.organization.department.UserFromDepartmentRequest")
	proto.RegisterType((*ListByDepartmentIDRequest)(nil), "iunite.club.srv.organization.department.ListByDepartmentIDRequest")
	proto.RegisterType((*ListByClubIDRequest)(nil), "iunite.club.srv.organization.department.ListByClubIDRequest")
	proto.RegisterType((*GetDepartmentWithIDRequest)(nil), "iunite.club.srv.organization.department.GetDepartmentWithIDRequest")
	proto.RegisterType((*DepartmentResponse)(nil), "iunite.club.srv.organization.department.DepartmentResponse")
	proto.RegisterType((*CreateDepartmentRequest)(nil), "iunite.club.srv.organization.department.CreateDepartmentRequest")
	proto.RegisterType((*CreateDepartmentResponse)(nil), "iunite.club.srv.organization.department.CreateDepartmentResponse")
	proto.RegisterType((*UpdateDepartmentRequest)(nil), "iunite.club.srv.organization.department.UpdateDepartmentRequest")
	proto.RegisterType((*UpdateDepartmentResponse)(nil), "iunite.club.srv.organization.department.UpdateDepartmentResponse")
	proto.RegisterType((*RemoveDepartmentRequest)(nil), "iunite.club.srv.organization.department.RemoveDepartmentRequest")
	proto.RegisterType((*RemoveDepartmentResponse)(nil), "iunite.club.srv.organization.department.RemoveDepartmentResponse")
	proto.RegisterType((*DepartmentListByParentIDRequest)(nil), "iunite.club.srv.organization.department.DepartmentListByParentIDRequest")
	proto.RegisterType((*DepartmentListResponse)(nil), "iunite.club.srv.organization.department.DepartmentListResponse")
}

func init() {
	proto.RegisterFile("proto/department/department.proto", fileDescriptor_department_9b874b28ea07144d)
}

var fileDescriptor_department_9b874b28ea07144d = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x93, 0x26, 0x6d, 0xa7, 0x15, 0xaa, 0x96, 0x28, 0x71, 0xcd, 0x81, 0xe0, 0x0b, 0xa5,
	0x02, 0x07, 0x95, 0x13, 0x3f, 0x12, 0xb4, 0xb1, 0x5a, 0x42, 0x0b, 0xad, 0xdc, 0x56, 0x95, 0x38,
	0xb1, 0x4d, 0x56, 0xad, 0x85, 0x63, 0x87, 0xf5, 0x26, 0xe2, 0xef, 0x25, 0x38, 0xa1, 0x72, 0xe1,
	0xc2, 0x89, 0x77, 0xe0, 0x09, 0x78, 0x29, 0xe4, 0x5d, 0x3b, 0x59, 0xdb, 0x71, 0x1b, 0xb7, 0x39,
	0x70, 0xf3, 0xec, 0xce, 0xce, 0x7c, 0xdf, 0xec, 0x37, 0xeb, 0x81, 0x3b, 0x3d, 0xea, 0x31, 0xaf,
	0xd1, 0x21, 0x3d, 0x4c, 0x59, 0x97, 0xb8, 0x4c, 0xfa, 0x34, 0xf8, 0x1e, 0xba, 0x6b, 0xf7, 0x5d,
	0x9b, 0x11, 0xa3, 0xed, 0xf4, 0x4f, 0x0c, 0x9f, 0x0e, 0x0c, 0x8f, 0x9e, 0x62, 0xd7, 0xfe, 0x8c,
	0x99, 0xed, 0xb9, 0xc6, 0xc8, 0x5d, 0x7b, 0x22, 0x39, 0x36, 0x7c, 0x42, 0x07, 0x76, 0x9b, 0xf8,
	0x0d, 0xd9, 0xbb, 0x21, 0xb2, 0xc5, 0x02, 0xf0, 0x25, 0x6d, 0x6d, 0xec, 0xd9, 0xbe, 0x4f, 0x68,
	0x78, 0x26, 0xf8, 0x14, 0xbe, 0xfa, 0x03, 0xa8, 0x1d, 0x10, 0x4c, 0xdb, 0x67, 0xe6, 0x30, 0xb7,
	0x45, 0x3e, 0xf4, 0x89, 0xcf, 0x10, 0x82, 0xd9, 0x37, 0xb8, 0x4b, 0x54, 0xa5, 0xae, 0xac, 0x2e,
	0x58, 0xfc, 0x5b, 0x7f, 0x0b, 0xcb, 0x47, 0x3e, 0xa1, 0xbb, 0xb6, 0xcf, 0x2c, 0xe2, 0xf7, 0x3c,
	0xd7, 0x27, 0xe8, 0x21, 0x94, 0x82, 0x35, 0x5f, 0x55, 0xea, 0xc5, 0xd5, 0xc5, 0x75, 0xcd, 0x48,
	0x72, 0xe4, 0xe9, 0x02, 0x17, 0x4b, 0x38, 0xa2, 0x0a, 0x94, 0x0e, 0x3d, 0x86, 0x1d, 0xb5, 0x50,
	0x57, 0x56, 0x8b, 0x96, 0x30, 0xf4, 0x2e, 0xac, 0x04, 0xdb, 0x5b, 0xd4, 0xeb, 0xa6, 0xc1, 0x54,
	0xe4, 0x24, 0x0b, 0x51, 0x20, 0x1d, 0x96, 0x46, 0xae, 0x2d, 0x93, 0xc7, 0x5b, 0xb0, 0x62, 0x6b,
	0xa8, 0x0a, 0xe5, 0xa6, 0xd3, 0x3f, 0x69, 0x99, 0x6a, 0x91, 0xef, 0x86, 0x96, 0x6e, 0xc3, 0x4a,
	0x40, 0x63, 0xf3, 0x93, 0xec, 0x2d, 0x71, 0xdf, 0xc7, 0xa7, 0x82, 0x7b, 0xd1, 0xe2, 0xdf, 0x01,
	0x84, 0x5d, 0xbb, 0x6b, 0xb3, 0x08, 0x35, 0x37, 0x52, 0x10, 0x8a, 0x69, 0x08, 0xfa, 0x31, 0xdc,
	0x14, 0xa9, 0x44, 0xea, 0xfc, 0x49, 0xb2, 0x38, 0xdc, 0x07, 0x6d, 0x9b, 0xb0, 0x51, 0xae, 0x63,
	0x9b, 0x9d, 0x8d, 0xe2, 0xdf, 0x80, 0x42, 0xcb, 0x0c, 0xaf, 0xaf, 0xd0, 0x32, 0xf5, 0x77, 0x80,
	0xe4, 0xc2, 0x86, 0xd7, 0xf7, 0x0a, 0x60, 0xb4, 0xca, 0xbd, 0x17, 0xd7, 0xd7, 0x8c, 0x0b, 0x75,
	0xba, 0x27, 0x19, 0x96, 0x74, 0x5a, 0x7f, 0x0f, 0xb5, 0x26, 0x25, 0x98, 0x91, 0x89, 0xd4, 0x84,
	0x34, 0x98, 0xdf, 0xc7, 0x54, 0xbe, 0xba, 0xa1, 0x8d, 0xea, 0xb0, 0x68, 0x12, 0xbf, 0x4d, 0xed,
	0x5e, 0x90, 0x25, 0xe4, 0x2d, 0x2f, 0xe9, 0x03, 0x50, 0xd3, 0xc9, 0xa6, 0x4f, 0x2a, 0x28, 0xe3,
	0xde, 0x0e, 0xc7, 0x37, 0x6f, 0x15, 0xf6, 0x76, 0xf4, 0x2f, 0x50, 0x3b, 0xea, 0x75, 0xc6, 0x92,
	0x4c, 0x54, 0x7c, 0x48, 0xba, 0x90, 0x41, 0xba, 0x78, 0x31, 0xe9, 0xd9, 0x34, 0xe9, 0x35, 0x50,
	0xd3, 0xc9, 0x43, 0xd2, 0x02, 0xa8, 0x32, 0x04, 0x7a, 0x0f, 0x6a, 0x16, 0xe9, 0x7a, 0x83, 0xcb,
	0x81, 0x06, 0x61, 0xd3, 0xae, 0x19, 0x61, 0xff, 0x2a, 0x70, 0x7b, 0xe4, 0x26, 0x84, 0x1d, 0x31,
	0x18, 0x27, 0xed, 0xd2, 0x38, 0x69, 0x97, 0x22, 0x69, 0xab, 0x30, 0xf7, 0x1a, 0x7f, 0xdc, 0x72,
	0xf0, 0x69, 0x58, 0x8d, 0xc8, 0xe4, 0x3b, 0xb6, 0xcb, 0x77, 0x66, 0xc3, 0x1d, 0x61, 0xc6, 0x4a,
	0x58, 0x4a, 0x94, 0xb0, 0x0a, 0x65, 0xf1, 0xa0, 0xa9, 0x65, 0xd1, 0x2a, 0xc2, 0xe2, 0xeb, 0x3d,
	0x4a, 0x70, 0x47, 0x9d, 0xe3, 0x4c, 0x42, 0x4b, 0xff, 0x0a, 0xd5, 0x38, 0x99, 0x21, 0xef, 0xdd,
	0xe0, 0x32, 0xa2, 0x9d, 0xe8, 0x75, 0xcb, 0x23, 0x22, 0xf9, 0xf8, 0xf8, 0x37, 0x6f, 0xfd, 0xcf,
	0x92, 0x2c, 0x54, 0x74, 0xae, 0xc0, 0x72, 0x52, 0xd3, 0xe8, 0x85, 0x31, 0xe1, 0x4f, 0xc3, 0xc8,
	0xe8, 0x3d, 0x6d, 0xe3, 0x1a, 0x11, 0x44, 0x31, 0xf4, 0x19, 0x8e, 0x2d, 0x29, 0xbd, 0x1c, 0xd8,
	0x32, 0x5a, 0x26, 0x07, 0xb6, 0x2c, 0xdd, 0x87, 0xd8, 0x92, 0xfa, 0xcd, 0x81, 0x2d, 0xa3, 0x4b,
	0x72, 0x60, 0xcb, 0x6a, 0x1e, 0x7d, 0x06, 0xfd, 0x56, 0xe0, 0x56, 0xec, 0x91, 0x8e, 0x77, 0x0c,
	0x7a, 0x39, 0x71, 0x92, 0x4b, 0x9a, 0x4e, 0x7b, 0x7e, 0xc5, 0x48, 0x12, 0xd8, 0x1f, 0x0a, 0x54,
	0x62, 0x60, 0x4d, 0xc2, 0xb0, 0xed, 0xf8, 0xa8, 0x39, 0x71, 0xec, 0xec, 0x1f, 0x92, 0xf6, 0xf4,
	0x0a, 0x00, 0x25, 0x70, 0x3f, 0x15, 0xa8, 0x6c, 0x74, 0x3a, 0xfc, 0xd7, 0x7f, 0xe8, 0x49, 0x37,
	0xbd, 0x39, 0xb9, 0x86, 0xb2, 0x06, 0x8c, 0xe9, 0xe8, 0xf0, 0x97, 0x02, 0x2b, 0x42, 0x0a, 0x1c,
	0x64, 0x3c, 0xd3, 0xff, 0x03, 0xf3, 0x5c, 0x81, 0xea, 0x36, 0x61, 0x1c, 0x63, 0x7c, 0xfe, 0xc9,
	0x81, 0x31, 0x73, 0x78, 0xd2, 0x1e, 0xe7, 0xe2, 0x99, 0x50, 0xe0, 0x37, 0xa1, 0xc0, 0x0d, 0xc7,
	0x69, 0x62, 0xf7, 0x80, 0x38, 0xa4, 0x2d, 0x70, 0xa2, 0x67, 0x39, 0x91, 0xc5, 0x66, 0xad, 0xeb,
	0x61, 0xfa, 0xae, 0xc0, 0x72, 0x72, 0x4a, 0xce, 0xf1, 0xbc, 0x64, 0x0c, 0xd8, 0x53, 0xe8, 0xd7,
	0x93, 0x32, 0x9f, 0xe2, 0x1f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x34, 0xc7, 0xc0, 0xc9, 0x7b,
	0x0c, 0x00, 0x00,
}
