// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/approved/approved.proto

package iunite_club_srv_core_approved

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Title                string          `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Kind                 string          `protobuf:"bytes,2,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Summary              string          `protobuf:"bytes,3,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Status               string          `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Description          string          `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Content              *_struct.Struct `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	ApprovedUsers        []string        `protobuf:"bytes,7,rep,name=ApprovedUsers,proto3" json:"ApprovedUsers,omitempty"`
	CopyUsers            []string        `protobuf:"bytes,8,rep,name=CopyUsers,proto3" json:"CopyUsers,omitempty"`
	ClubID               string          `protobuf:"bytes,9,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	CreatorID            string          `protobuf:"bytes,10,opt,name=CreatorID,proto3" json:"CreatorID,omitempty"`
	DepartmentID         string          `protobuf:"bytes,11,opt,name=DepartmentID,proto3" json:"DepartmentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CreateRequest) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *CreateRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRequest) GetContent() *_struct.Struct {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *CreateRequest) GetApprovedUsers() []string {
	if m != nil {
		return m.ApprovedUsers
	}
	return nil
}

func (m *CreateRequest) GetCopyUsers() []string {
	if m != nil {
		return m.CopyUsers
	}
	return nil
}

func (m *CreateRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

func (m *CreateRequest) GetCreatorID() string {
	if m != nil {
		return m.CreatorID
	}
	return ""
}

func (m *CreateRequest) GetDepartmentID() string {
	if m != nil {
		return m.DepartmentID
	}
	return ""
}

type ApprovedResponse struct {
	Approved             *ApprovedPB `protobuf:"bytes,1,opt,name=Approved,proto3" json:"Approved,omitempty"`
	OK                   bool        `protobuf:"varint,2,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ApprovedResponse) Reset()         { *m = ApprovedResponse{} }
func (m *ApprovedResponse) String() string { return proto.CompactTextString(m) }
func (*ApprovedResponse) ProtoMessage()    {}
func (*ApprovedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{1}
}
func (m *ApprovedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApprovedResponse.Unmarshal(m, b)
}
func (m *ApprovedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApprovedResponse.Marshal(b, m, deterministic)
}
func (dst *ApprovedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedResponse.Merge(dst, src)
}
func (m *ApprovedResponse) XXX_Size() int {
	return xxx_messageInfo_ApprovedResponse.Size(m)
}
func (m *ApprovedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedResponse proto.InternalMessageInfo

func (m *ApprovedResponse) GetApproved() *ApprovedPB {
	if m != nil {
		return m.Approved
	}
	return nil
}

func (m *ApprovedResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type Response struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type ApprovedPB struct {
	ID                   string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Kind                 string               `protobuf:"bytes,3,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Summary              string               `protobuf:"bytes,4,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Status               string               `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	Content              *_struct.Struct      `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`
	Flows                []*ApprovedFlowPB    `protobuf:"bytes,8,rep,name=Flows,proto3" json:"Flows,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	ClubID               string               `protobuf:"bytes,11,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApprovedPB) Reset()         { *m = ApprovedPB{} }
func (m *ApprovedPB) String() string { return proto.CompactTextString(m) }
func (*ApprovedPB) ProtoMessage()    {}
func (*ApprovedPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{3}
}
func (m *ApprovedPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApprovedPB.Unmarshal(m, b)
}
func (m *ApprovedPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApprovedPB.Marshal(b, m, deterministic)
}
func (dst *ApprovedPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedPB.Merge(dst, src)
}
func (m *ApprovedPB) XXX_Size() int {
	return xxx_messageInfo_ApprovedPB.Size(m)
}
func (m *ApprovedPB) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedPB.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedPB proto.InternalMessageInfo

func (m *ApprovedPB) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ApprovedPB) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ApprovedPB) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ApprovedPB) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ApprovedPB) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ApprovedPB) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ApprovedPB) GetContent() *_struct.Struct {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ApprovedPB) GetFlows() []*ApprovedFlowPB {
	if m != nil {
		return m.Flows
	}
	return nil
}

func (m *ApprovedPB) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ApprovedPB) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *ApprovedPB) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type ApprovedFlowPB struct {
	ID                   string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Kind                 string               `protobuf:"bytes,2,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Options              string               `protobuf:"bytes,3,opt,name=Options,proto3" json:"Options,omitempty"`
	HandlerID            string               `protobuf:"bytes,4,opt,name=HandlerID,proto3" json:"HandlerID,omitempty"`
	Status               int32                `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Sort                 int32                `protobuf:"varint,6,opt,name=Sort,proto3" json:"Sort,omitempty"`
	ApprovedID           string               `protobuf:"bytes,7,opt,name=ApprovedID,proto3" json:"ApprovedID,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApprovedFlowPB) Reset()         { *m = ApprovedFlowPB{} }
func (m *ApprovedFlowPB) String() string { return proto.CompactTextString(m) }
func (*ApprovedFlowPB) ProtoMessage()    {}
func (*ApprovedFlowPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{4}
}
func (m *ApprovedFlowPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApprovedFlowPB.Unmarshal(m, b)
}
func (m *ApprovedFlowPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApprovedFlowPB.Marshal(b, m, deterministic)
}
func (dst *ApprovedFlowPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedFlowPB.Merge(dst, src)
}
func (m *ApprovedFlowPB) XXX_Size() int {
	return xxx_messageInfo_ApprovedFlowPB.Size(m)
}
func (m *ApprovedFlowPB) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedFlowPB.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedFlowPB proto.InternalMessageInfo

func (m *ApprovedFlowPB) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ApprovedFlowPB) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ApprovedFlowPB) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *ApprovedFlowPB) GetHandlerID() string {
	if m != nil {
		return m.HandlerID
	}
	return ""
}

func (m *ApprovedFlowPB) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ApprovedFlowPB) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *ApprovedFlowPB) GetApprovedID() string {
	if m != nil {
		return m.ApprovedID
	}
	return ""
}

func (m *ApprovedFlowPB) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ApprovedFlowPB) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ListResponse struct {
	Approveds            []*ApprovedPB `protobuf:"bytes,1,rep,name=Approveds,proto3" json:"Approveds,omitempty"`
	Total                int64         `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{5}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetApproveds() []*ApprovedPB {
	if m != nil {
		return m.Approveds
	}
	return nil
}

func (m *ListResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ListRequest struct {
	ClubID               string   `protobuf:"bytes,1,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Search               string   `protobuf:"bytes,4,opt,name=Search,proto3" json:"Search,omitempty"`
	Page                 int64    `protobuf:"varint,5,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,6,opt,name=Limit,proto3" json:"Limit,omitempty"`
	HandlerID            string   `protobuf:"bytes,7,opt,name=HandlerID,proto3" json:"HandlerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{6}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

func (m *ListRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ListRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ListRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *ListRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetHandlerID() string {
	if m != nil {
		return m.HandlerID
	}
	return ""
}

type ListV2Request struct {
	ClubID               string   `protobuf:"bytes,1,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	FlowStatus           string   `protobuf:"bytes,3,opt,name=FlowStatus,proto3" json:"FlowStatus,omitempty"`
	ReadState            string   `protobuf:"bytes,4,opt,name=ReadState,proto3" json:"ReadState,omitempty"`
	Search               string   `protobuf:"bytes,5,opt,name=Search,proto3" json:"Search,omitempty"`
	Page                 int64    `protobuf:"varint,6,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,7,opt,name=Limit,proto3" json:"Limit,omitempty"`
	HandlerID            string   `protobuf:"bytes,8,opt,name=HandlerID,proto3" json:"HandlerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListV2Request) Reset()         { *m = ListV2Request{} }
func (m *ListV2Request) String() string { return proto.CompactTextString(m) }
func (*ListV2Request) ProtoMessage()    {}
func (*ListV2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{7}
}
func (m *ListV2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListV2Request.Unmarshal(m, b)
}
func (m *ListV2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListV2Request.Marshal(b, m, deterministic)
}
func (dst *ListV2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListV2Request.Merge(dst, src)
}
func (m *ListV2Request) XXX_Size() int {
	return xxx_messageInfo_ListV2Request.Size(m)
}
func (m *ListV2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ListV2Request.DiscardUnknown(m)
}

var xxx_messageInfo_ListV2Request proto.InternalMessageInfo

func (m *ListV2Request) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

func (m *ListV2Request) GetFlowStatus() string {
	if m != nil {
		return m.FlowStatus
	}
	return ""
}

func (m *ListV2Request) GetReadState() string {
	if m != nil {
		return m.ReadState
	}
	return ""
}

func (m *ListV2Request) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *ListV2Request) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListV2Request) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListV2Request) GetHandlerID() string {
	if m != nil {
		return m.HandlerID
	}
	return ""
}

type ListByPusherRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	UserID               string   `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Search               string   `protobuf:"bytes,4,opt,name=Search,proto3" json:"Search,omitempty"`
	ClubID               string   `protobuf:"bytes,5,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListByPusherRequest) Reset()         { *m = ListByPusherRequest{} }
func (m *ListByPusherRequest) String() string { return proto.CompactTextString(m) }
func (*ListByPusherRequest) ProtoMessage()    {}
func (*ListByPusherRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{8}
}
func (m *ListByPusherRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByPusherRequest.Unmarshal(m, b)
}
func (m *ListByPusherRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByPusherRequest.Marshal(b, m, deterministic)
}
func (dst *ListByPusherRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByPusherRequest.Merge(dst, src)
}
func (m *ListByPusherRequest) XXX_Size() int {
	return xxx_messageInfo_ListByPusherRequest.Size(m)
}
func (m *ListByPusherRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByPusherRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListByPusherRequest proto.InternalMessageInfo

func (m *ListByPusherRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListByPusherRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListByPusherRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ListByPusherRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *ListByPusherRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

type ListByCountRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListByCountRequest) Reset()         { *m = ListByCountRequest{} }
func (m *ListByCountRequest) String() string { return proto.CompactTextString(m) }
func (*ListByCountRequest) ProtoMessage()    {}
func (*ListByCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{9}
}
func (m *ListByCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByCountRequest.Unmarshal(m, b)
}
func (m *ListByCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByCountRequest.Marshal(b, m, deterministic)
}
func (dst *ListByCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByCountRequest.Merge(dst, src)
}
func (m *ListByCountRequest) XXX_Size() int {
	return xxx_messageInfo_ListByCountRequest.Size(m)
}
func (m *ListByCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListByCountRequest proto.InternalMessageInfo

func (m *ListByCountRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type DetailsRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailsRequest) Reset()         { *m = DetailsRequest{} }
func (m *DetailsRequest) String() string { return proto.CompactTextString(m) }
func (*DetailsRequest) ProtoMessage()    {}
func (*DetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{10}
}
func (m *DetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailsRequest.Unmarshal(m, b)
}
func (m *DetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailsRequest.Marshal(b, m, deterministic)
}
func (dst *DetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailsRequest.Merge(dst, src)
}
func (m *DetailsRequest) XXX_Size() int {
	return xxx_messageInfo_DetailsRequest.Size(m)
}
func (m *DetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailsRequest proto.InternalMessageInfo

func (m *DetailsRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ExecuteRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Options              string   `protobuf:"bytes,3,opt,name=Options,proto3" json:"Options,omitempty"`
	UserID               string   `protobuf:"bytes,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ClubID               string   `protobuf:"bytes,5,opt,name=ClubID,proto3" json:"ClubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_approved_28292d24413e18ef, []int{11}
}
func (m *ExecuteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteRequest.Unmarshal(m, b)
}
func (m *ExecuteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteRequest.Marshal(b, m, deterministic)
}
func (dst *ExecuteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteRequest.Merge(dst, src)
}
func (m *ExecuteRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteRequest.Size(m)
}
func (m *ExecuteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteRequest proto.InternalMessageInfo

func (m *ExecuteRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ExecuteRequest) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *ExecuteRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *ExecuteRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ExecuteRequest) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "iunite.club.srv.core.approved.CreateRequest")
	proto.RegisterType((*ApprovedResponse)(nil), "iunite.club.srv.core.approved.ApprovedResponse")
	proto.RegisterType((*Response)(nil), "iunite.club.srv.core.approved.Response")
	proto.RegisterType((*ApprovedPB)(nil), "iunite.club.srv.core.approved.ApprovedPB")
	proto.RegisterType((*ApprovedFlowPB)(nil), "iunite.club.srv.core.approved.ApprovedFlowPB")
	proto.RegisterType((*ListResponse)(nil), "iunite.club.srv.core.approved.ListResponse")
	proto.RegisterType((*ListRequest)(nil), "iunite.club.srv.core.approved.ListRequest")
	proto.RegisterType((*ListV2Request)(nil), "iunite.club.srv.core.approved.ListV2Request")
	proto.RegisterType((*ListByPusherRequest)(nil), "iunite.club.srv.core.approved.ListByPusherRequest")
	proto.RegisterType((*ListByCountRequest)(nil), "iunite.club.srv.core.approved.ListByCountRequest")
	proto.RegisterType((*DetailsRequest)(nil), "iunite.club.srv.core.approved.DetailsRequest")
	proto.RegisterType((*ExecuteRequest)(nil), "iunite.club.srv.core.approved.ExecuteRequest")
}

func init() {
	proto.RegisterFile("proto/approved/approved.proto", fileDescriptor_approved_28292d24413e18ef)
}

var fileDescriptor_approved_28292d24413e18ef = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xc6, 0x76, 0x6c, 0x27, 0xe5, 0x9d, 0x11, 0x6a, 0xd0, 0x60, 0x45, 0xbb, 0x4b, 0x64, 0x21,
	0x31, 0x2c, 0xe0, 0xd1, 0x86, 0x0b, 0xd7, 0x9d, 0x78, 0x81, 0x68, 0x56, 0x9a, 0xa8, 0xb3, 0x0b,
	0x67, 0x27, 0x69, 0xb2, 0x16, 0x8e, 0xed, 0x75, 0xb7, 0x07, 0xe6, 0xca, 0x85, 0xb7, 0xe0, 0x09,
	0xb8, 0xf0, 0x10, 0xbc, 0x06, 0x4f, 0xc2, 0x01, 0x75, 0xb7, 0xdb, 0x7f, 0x9b, 0xdf, 0xb9, 0xb9,
	0xaa, 0xab, 0xba, 0x7e, 0xbe, 0xaf, 0xaa, 0x0d, 0x4f, 0xb2, 0x3c, 0x65, 0xe9, 0x55, 0x98, 0x65,
	0x79, 0x7a, 0x47, 0x56, 0xd5, 0x87, 0x2f, 0xf4, 0xe8, 0x49, 0x54, 0x24, 0x11, 0x23, 0xfe, 0x32,
	0x2e, 0x16, 0x3e, 0xcd, 0xef, 0xfc, 0x65, 0x9a, 0x13, 0x5f, 0x19, 0x0d, 0x3f, 0x5d, 0xa7, 0xe9,
	0x3a, 0x26, 0x57, 0xc2, 0x78, 0x51, 0xfc, 0x7c, 0xc5, 0xa2, 0x0d, 0xa1, 0x2c, 0xdc, 0x64, 0xd2,
	0x7f, 0xf8, 0xb8, 0x6b, 0x40, 0x59, 0x5e, 0x2c, 0x99, 0x3c, 0xf5, 0xfe, 0xd5, 0xe1, 0x6c, 0x92,
	0x93, 0x90, 0x11, 0x4c, 0xde, 0x15, 0x84, 0x32, 0xf4, 0x31, 0x98, 0xaf, 0x23, 0x16, 0x13, 0x57,
	0x1b, 0x69, 0x97, 0x03, 0x2c, 0x05, 0x84, 0xa0, 0x77, 0x13, 0x25, 0x2b, 0x57, 0x17, 0x4a, 0xf1,
	0x8d, 0x5c, 0xb0, 0xe7, 0xc5, 0x66, 0x13, 0xe6, 0xf7, 0xae, 0x21, 0xd4, 0x4a, 0x44, 0x17, 0x60,
	0xcd, 0x59, 0xc8, 0x0a, 0xea, 0xf6, 0xc4, 0x41, 0x29, 0xa1, 0x11, 0x38, 0x01, 0xa1, 0xcb, 0x3c,
	0xca, 0x58, 0x94, 0x26, 0xae, 0x29, 0x0e, 0x9b, 0x2a, 0xf4, 0x1c, 0xec, 0x49, 0x9a, 0x30, 0x92,
	0x30, 0xd7, 0x1a, 0x69, 0x97, 0xce, 0xf8, 0x13, 0x5f, 0xe6, 0xef, 0xab, 0xfc, 0xfd, 0xb9, 0xc8,
	0x1f, 0x2b, 0x3b, 0xf4, 0x19, 0x9c, 0xbd, 0x28, 0xbb, 0xf1, 0x86, 0x92, 0x9c, 0xba, 0xf6, 0xc8,
	0xb8, 0x1c, 0xe0, 0xb6, 0x12, 0x3d, 0x86, 0xc1, 0x24, 0xcd, 0xee, 0xa5, 0x45, 0x5f, 0x58, 0xd4,
	0x0a, 0x9e, 0xf0, 0x24, 0x2e, 0x16, 0xd3, 0xc0, 0x1d, 0xc8, 0x84, 0xa5, 0x24, 0xbc, 0x78, 0x77,
	0xd2, 0x7c, 0x1a, 0xb8, 0x20, 0x8e, 0x6a, 0x05, 0xf2, 0xe0, 0x51, 0x40, 0xb2, 0x30, 0x67, 0x1b,
	0x92, 0xb0, 0x69, 0xe0, 0x3a, 0xc2, 0xa0, 0xa5, 0xf3, 0x22, 0xf8, 0x50, 0x25, 0x82, 0x09, 0xcd,
	0xd2, 0x84, 0x12, 0xf4, 0x12, 0xfa, 0x4a, 0x27, 0xba, 0xec, 0x8c, 0xbf, 0xf0, 0xf7, 0xa2, 0xec,
	0x2b, 0xf3, 0xd9, 0x35, 0xae, 0x5c, 0xd1, 0x39, 0xe8, 0xb7, 0x37, 0x02, 0x91, 0x3e, 0xd6, 0x6f,
	0x6f, 0xbc, 0x21, 0xf4, 0xab, 0x10, 0xf2, 0x4c, 0xab, 0xce, 0xfe, 0x34, 0x00, 0xea, 0x4b, 0xf8,
	0xf1, 0x34, 0x28, 0x11, 0xd6, 0xa7, 0x41, 0x0d, 0xba, 0xbe, 0x0d, 0x74, 0x63, 0x3b, 0xe8, 0xbd,
	0x5d, 0xa0, 0x9b, 0xfb, 0x40, 0xb7, 0xf6, 0x82, 0x6e, 0x1f, 0x09, 0xfa, 0x04, 0xcc, 0xef, 0xe2,
	0xf4, 0x57, 0x09, 0xa5, 0x33, 0xfe, 0xfa, 0xc8, 0xfe, 0x71, 0x9f, 0xd9, 0x35, 0x96, 0xbe, 0xe8,
	0xdb, 0x12, 0x5d, 0xb2, 0x7a, 0xc1, 0x04, 0xf0, 0xce, 0x78, 0xf8, 0x5e, 0xe4, 0xd7, 0x6a, 0x9e,
	0x70, 0x6d, 0xcc, 0x3d, 0xdf, 0x64, 0xab, 0xd2, 0x13, 0x0e, 0x7b, 0x56, 0xc6, 0x0d, 0xa6, 0x39,
	0x4d, 0xa6, 0x79, 0x7f, 0xe9, 0x70, 0xde, 0xce, 0xf2, 0x3d, 0x90, 0x76, 0xcc, 0xe0, 0xad, 0x68,
	0x22, 0x55, 0x33, 0x58, 0x8a, 0x9c, 0xba, 0x3f, 0x84, 0xc9, 0x2a, 0x26, 0x9c, 0xba, 0x12, 0xaa,
	0x5a, 0xd1, 0x01, 0xcb, 0xac, 0xc0, 0x42, 0xd0, 0x9b, 0xa7, 0xb9, 0x1c, 0x3e, 0x13, 0x8b, 0x6f,
	0xf4, 0xb4, 0xa6, 0xce, 0x34, 0x10, 0x08, 0x0d, 0x70, 0x43, 0xd3, 0x6e, 0x63, 0xff, 0xc1, 0x6d,
	0x1c, 0x9c, 0xd0, 0x46, 0x6f, 0x03, 0x8f, 0x5e, 0x45, 0x94, 0x55, 0x7c, 0xff, 0x1e, 0x06, 0x2a,
	0x23, 0xea, 0x6a, 0x82, 0x13, 0x27, 0xcc, 0x54, 0xed, 0x2b, 0x26, 0x21, 0x65, 0x61, 0x2c, 0xba,
	0x6c, 0x60, 0x29, 0x78, 0x7f, 0x6b, 0xe0, 0xc8, 0x78, 0x72, 0x49, 0xd6, 0x28, 0x6a, 0xad, 0x7d,
	0xb1, 0x0d, 0xa2, 0xba, 0xd5, 0x46, 0x6b, 0x2e, 0xb8, 0x9e, 0x84, 0xf9, 0xf2, 0x6d, 0xb5, 0x24,
	0x85, 0xc4, 0xef, 0x98, 0x85, 0x6b, 0x22, 0x80, 0x31, 0xb0, 0xf8, 0xe6, 0x59, 0xbd, 0x8a, 0x36,
	0x91, 0xc4, 0xc5, 0xc0, 0x52, 0x68, 0x43, 0x6c, 0x77, 0x20, 0xf6, 0xfe, 0xd1, 0xe0, 0x8c, 0xe7,
	0xfc, 0xe3, 0xf8, 0x50, 0xd6, 0x4f, 0x01, 0x38, 0xe5, 0x5a, 0x59, 0x36, 0x34, 0x3c, 0x0e, 0x26,
	0xe1, 0x8a, 0x4b, 0x44, 0x51, 0xa9, 0x52, 0x34, 0xea, 0x30, 0xb7, 0xd6, 0x61, 0x6d, 0xab, 0xc3,
	0xde, 0x59, 0x47, 0xbf, 0x5b, 0xc7, 0x1f, 0x1a, 0x7c, 0xc4, 0xeb, 0xb8, 0xbe, 0x9f, 0x15, 0xf4,
	0x2d, 0xc9, 0x55, 0x35, 0xea, 0x7e, 0x6d, 0xdb, 0xfd, 0x7a, 0xf3, 0xfe, 0x0b, 0xb0, 0xf8, 0x9a,
	0x9f, 0x06, 0x0a, 0x01, 0x29, 0xed, 0x44, 0xa0, 0xee, 0x93, 0xd9, 0x9a, 0xd1, 0x67, 0x80, 0x64,
	0x22, 0x93, 0xb4, 0x48, 0x58, 0xe3, 0xc1, 0x14, 0x72, 0x99, 0x88, 0x14, 0xbc, 0x11, 0x9c, 0x07,
	0x84, 0x85, 0x51, 0x4c, 0x95, 0x5d, 0x67, 0x9c, 0xbd, 0xdf, 0x35, 0x38, 0x7f, 0xf9, 0x1b, 0x59,
	0x16, 0xf5, 0xdb, 0xdb, 0x9d, 0xf8, 0x0b, 0xb0, 0x30, 0xa1, 0x45, 0xcc, 0xca, 0x2d, 0x5f, 0x4a,
	0x7b, 0xa6, 0xbe, 0x2e, 0xb5, 0xd7, 0x2d, 0x75, 0x5b, 0x49, 0xe3, 0xff, 0xcc, 0xfa, 0x2d, 0x42,
	0x21, 0xf4, 0x78, 0x7d, 0xe8, 0xd9, 0x81, 0xc9, 0x69, 0x4c, 0xc2, 0xf0, 0xcb, 0xa3, 0x6c, 0xe5,
	0x94, 0x7a, 0x1f, 0x20, 0x02, 0x96, 0xe4, 0x24, 0xfa, 0xea, 0x08, 0xc7, 0x8a, 0xba, 0xa7, 0x86,
	0x79, 0x27, 0xd7, 0x83, 0xa2, 0x0c, 0x1a, 0x1f, 0xe1, 0xde, 0xe1, 0xd7, 0xa9, 0x21, 0xef, 0x00,
	0xfd, 0x14, 0x46, 0x2c, 0x4a, 0xd6, 0x25, 0xa8, 0xa2, 0x95, 0xcf, 0x8f, 0x0a, 0xdc, 0xe4, 0xd3,
	0xa9, 0x71, 0x7f, 0x01, 0xbb, 0x24, 0x1a, 0x3a, 0xf4, 0x0a, 0xb6, 0x09, 0x39, 0xbc, 0x3a, 0x72,
	0x41, 0x36, 0x82, 0x2d, 0xc1, 0x2e, 0xab, 0x3b, 0x18, 0xac, 0x4d, 0xed, 0xe1, 0xe7, 0x07, 0xcc,
	0x1b, 0x41, 0x22, 0xb0, 0xe4, 0x13, 0x71, 0x90, 0x23, 0xad, 0x3f, 0xd7, 0x07, 0xd4, 0xb3, 0xb0,
	0xc4, 0x2b, 0xf3, 0xcd, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0xf7, 0xbd, 0xdb, 0x84, 0x0b,
	0x00, 0x00,
}