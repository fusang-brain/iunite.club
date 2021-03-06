// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sms/sms.proto

package iunite_club_srv_message_sms

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

type ValidateMobileCodeRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateMobileCodeRequest) Reset()         { *m = ValidateMobileCodeRequest{} }
func (m *ValidateMobileCodeRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateMobileCodeRequest) ProtoMessage()    {}
func (*ValidateMobileCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sms_9fa7ebaaa25889e2, []int{0}
}
func (m *ValidateMobileCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateMobileCodeRequest.Unmarshal(m, b)
}
func (m *ValidateMobileCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateMobileCodeRequest.Marshal(b, m, deterministic)
}
func (dst *ValidateMobileCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateMobileCodeRequest.Merge(dst, src)
}
func (m *ValidateMobileCodeRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateMobileCodeRequest.Size(m)
}
func (m *ValidateMobileCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateMobileCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateMobileCodeRequest proto.InternalMessageInfo

func (m *ValidateMobileCodeRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ValidateMobileCodeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ValidateResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sms_9fa7ebaaa25889e2, []int{1}
}
func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (dst *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(dst, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type SendVerifyCodeRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerifyCodeRequest) Reset()         { *m = SendVerifyCodeRequest{} }
func (m *SendVerifyCodeRequest) String() string { return proto.CompactTextString(m) }
func (*SendVerifyCodeRequest) ProtoMessage()    {}
func (*SendVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sms_9fa7ebaaa25889e2, []int{2}
}
func (m *SendVerifyCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerifyCodeRequest.Unmarshal(m, b)
}
func (m *SendVerifyCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerifyCodeRequest.Marshal(b, m, deterministic)
}
func (dst *SendVerifyCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerifyCodeRequest.Merge(dst, src)
}
func (m *SendVerifyCodeRequest) XXX_Size() int {
	return xxx_messageInfo_SendVerifyCodeRequest.Size(m)
}
func (m *SendVerifyCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerifyCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerifyCodeRequest proto.InternalMessageInfo

func (m *SendVerifyCodeRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SendVerifyCodeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type SendResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sms_9fa7ebaaa25889e2, []int{3}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *SendResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SendResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidateMobileCodeRequest)(nil), "iunite.club.srv.message.sms.ValidateMobileCodeRequest")
	proto.RegisterType((*ValidateResponse)(nil), "iunite.club.srv.message.sms.ValidateResponse")
	proto.RegisterType((*SendVerifyCodeRequest)(nil), "iunite.club.srv.message.sms.SendVerifyCodeRequest")
	proto.RegisterType((*SendResponse)(nil), "iunite.club.srv.message.sms.SendResponse")
}

func init() { proto.RegisterFile("proto/sms/sms.proto", fileDescriptor_sms_9fa7ebaaa25889e2) }

var fileDescriptor_sms_9fa7ebaaa25889e2 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0x2d, 0x06, 0x61, 0x3d, 0x30, 0x4f, 0x48, 0x3a, 0xb3, 0x34, 0x2f, 0xb3,
	0x24, 0x55, 0x2f, 0x39, 0xa7, 0x34, 0x49, 0xaf, 0xb8, 0xa8, 0x4c, 0x2f, 0x37, 0xb5, 0xb8, 0x38,
	0x31, 0x3d, 0x55, 0xaf, 0x38, 0xb7, 0x58, 0xc9, 0x9d, 0x4b, 0x32, 0x2c, 0x31, 0x27, 0x33, 0x25,
	0xb1, 0x24, 0xd5, 0x37, 0x3f, 0x29, 0x33, 0x27, 0xd5, 0x39, 0x3f, 0x25, 0x35, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x0d, 0x22, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0xe5, 0x09, 0x09, 0x71, 0xb1, 0x80, 0x94, 0x49, 0x30, 0x81, 0x45, 0xc1, 0x6c, 0x25, 0x25,
	0x2e, 0x01, 0x98, 0x41, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c,
	0xfe, 0xde, 0x60, 0xbd, 0x1c, 0x41, 0x4c, 0xfe, 0xde, 0x4a, 0xce, 0x5c, 0xa2, 0xc1, 0xa9, 0x79,
	0x29, 0x61, 0xa9, 0x45, 0x99, 0x69, 0x95, 0x44, 0x5a, 0x14, 0x52, 0x59, 0x00, 0xb7, 0x08, 0xc4,
	0x56, 0xf2, 0xe1, 0xe2, 0x01, 0x19, 0x82, 0xcb, 0x12, 0x6c, 0x8e, 0x13, 0x92, 0xe0, 0x62, 0xf7,
	0x85, 0x78, 0x5a, 0x82, 0x19, 0x2c, 0x0c, 0xe3, 0x1a, 0x7d, 0x67, 0xe4, 0x62, 0x0e, 0xf6, 0x0d,
	0x16, 0x2a, 0xe4, 0xe2, 0x43, 0x75, 0x9a, 0x90, 0x91, 0x1e, 0x9e, 0x70, 0xd3, 0xc3, 0xea, 0x0f,
	0x29, 0x4d, 0x82, 0x7a, 0x60, 0xce, 0x56, 0x62, 0x10, 0xaa, 0xe5, 0x12, 0xc2, 0x0c, 0x7a, 0x21,
	0x33, 0xbc, 0x46, 0xe0, 0x8c, 0x2b, 0x29, 0x5d, 0xa2, 0xf4, 0x21, 0xac, 0x4f, 0x62, 0x03, 0xa7,
	0x0e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x36, 0xfd, 0xed, 0x34, 0x02, 0x00, 0x00,
}
