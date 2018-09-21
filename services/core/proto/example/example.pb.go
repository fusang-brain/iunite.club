// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/core/proto/example/example.proto

package iunite_club_srv_core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{2}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{3}
}
func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (dst *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(dst, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{4}
}
func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (dst *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(dst, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{5}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_1ff061818afdd91a, []int{6}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (dst *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(dst, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "iunite.club.srv.core.Message")
	proto.RegisterType((*Request)(nil), "iunite.club.srv.core.Request")
	proto.RegisterType((*Response)(nil), "iunite.club.srv.core.Response")
	proto.RegisterType((*StreamingRequest)(nil), "iunite.club.srv.core.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "iunite.club.srv.core.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "iunite.club.srv.core.Ping")
	proto.RegisterType((*Pong)(nil), "iunite.club.srv.core.Pong")
}

func init() {
	proto.RegisterFile("services/core/proto/example/example.proto", fileDescriptor_example_1ff061818afdd91a)
}

var fileDescriptor_example_1ff061818afdd91a = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0xb4, 0xff, 0xb6, 0xff, 0x39, 0xd5, 0xa5, 0x88, 0x44, 0x5b, 0x64, 0x0f, 0x9a,
	0x5e, 0x36, 0x45, 0x3f, 0x82, 0x88, 0x5e, 0x04, 0x89, 0x67, 0x0f, 0xdb, 0x30, 0x2c, 0xc1, 0x64,
	0x37, 0xee, 0x6c, 0x8a, 0x7e, 0x76, 0x2f, 0x92, 0x4d, 0x02, 0x22, 0x09, 0x9e, 0x32, 0x93, 0xdf,
	0x7b, 0xc3, 0xcc, 0x63, 0x61, 0x47, 0xe8, 0x8e, 0x79, 0x86, 0x94, 0x64, 0xd6, 0x61, 0x52, 0x39,
	0xeb, 0x6d, 0x82, 0x1f, 0xaa, 0xac, 0x0a, 0xec, 0xbf, 0x32, 0xfc, 0xe5, 0xeb, 0xbc, 0x36, 0xb9,
	0x47, 0x99, 0x15, 0xf5, 0x41, 0x92, 0x3b, 0xca, 0xc6, 0x21, 0xce, 0x61, 0xf1, 0x84, 0x44, 0x4a,
	0x23, 0x5f, 0xc1, 0x94, 0xd4, 0xe7, 0x19, 0xbb, 0x64, 0xf1, 0xff, 0xb4, 0x29, 0xc5, 0x06, 0x16,
	0x29, 0xbe, 0xd7, 0x48, 0x9e, 0x73, 0x98, 0x19, 0x55, 0x62, 0x47, 0x43, 0x2d, 0x2e, 0x60, 0x99,
	0x22, 0x55, 0xd6, 0x50, 0x30, 0x97, 0xa4, 0x7b, 0x73, 0x49, 0x5a, 0xc4, 0xb0, 0x7a, 0xf1, 0x0e,
	0x55, 0x99, 0x1b, 0xdd, 0x4f, 0x59, 0xc3, 0xbf, 0xcc, 0xd6, 0xc6, 0x07, 0xdd, 0x34, 0x6d, 0x1b,
	0xb1, 0x83, 0x93, 0x1f, 0xca, 0x6e, 0xe0, 0xb0, 0x74, 0x0b, 0xb3, 0xe7, 0xdc, 0x68, 0x7e, 0x0a,
	0x73, 0xf2, 0xce, 0xbe, 0x61, 0x87, 0xbb, 0x2e, 0x70, 0x3b, 0xce, 0x6f, 0xbe, 0x18, 0x2c, 0xee,
	0xdb, 0x58, 0xf8, 0x03, 0xcc, 0xee, 0x54, 0x51, 0xf0, 0x8d, 0x1c, 0x4a, 0x46, 0x76, 0x3b, 0x47,
	0xdb, 0x31, 0xdc, 0x2e, 0x2a, 0x26, 0xfc, 0x15, 0xe6, 0xed, 0xfe, 0xfc, 0x6a, 0x58, 0xfb, 0x3b,
	0x87, 0xe8, 0xfa, 0x4f, 0x5d, 0x3f, 0x7c, 0xcf, 0xf8, 0x23, 0x2c, 0x9b, 0x9b, 0xc3, 0x5d, 0xd1,
	0xb0, 0xb1, 0xe1, 0xd1, 0x18, 0xb3, 0x46, 0x8b, 0x49, 0xcc, 0xf6, 0xec, 0x30, 0x0f, 0x2f, 0xe1,
	0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xd6, 0xaa, 0xcc, 0x36, 0x02, 0x00, 0x00,
}