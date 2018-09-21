// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/secruity/secruity.proto

package iunite_club_srv_user_secruity

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

type SignupWithMobileRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
	School               string   `protobuf:"bytes,2,opt,name=School,proto3" json:"School,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	MobileAreaCode       int32    `protobuf:"varint,4,opt,name=MobileAreaCode,proto3" json:"MobileAreaCode,omitempty"`
	ValidateCode         string   `protobuf:"bytes,5,opt,name=ValidateCode,proto3" json:"ValidateCode,omitempty"`
	FirstName            string   `protobuf:"bytes,6,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,7,opt,name=LastName,proto3" json:"LastName,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,8,opt,name=ConfirmPassword,proto3" json:"ConfirmPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupWithMobileRequest) Reset()         { *m = SignupWithMobileRequest{} }
func (m *SignupWithMobileRequest) String() string { return proto.CompactTextString(m) }
func (*SignupWithMobileRequest) ProtoMessage()    {}
func (*SignupWithMobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_secruity_c368f40297f61236, []int{0}
}
func (m *SignupWithMobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupWithMobileRequest.Unmarshal(m, b)
}
func (m *SignupWithMobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupWithMobileRequest.Marshal(b, m, deterministic)
}
func (dst *SignupWithMobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupWithMobileRequest.Merge(dst, src)
}
func (m *SignupWithMobileRequest) XXX_Size() int {
	return xxx_messageInfo_SignupWithMobileRequest.Size(m)
}
func (m *SignupWithMobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupWithMobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupWithMobileRequest proto.InternalMessageInfo

func (m *SignupWithMobileRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignupWithMobileRequest) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *SignupWithMobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SignupWithMobileRequest) GetMobileAreaCode() int32 {
	if m != nil {
		return m.MobileAreaCode
	}
	return 0
}

func (m *SignupWithMobileRequest) GetValidateCode() string {
	if m != nil {
		return m.ValidateCode
	}
	return ""
}

func (m *SignupWithMobileRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SignupWithMobileRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *SignupWithMobileRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type SignupResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	TokenExpiredAt       int64    `protobuf:"varint,3,opt,name=TokenExpiredAt,proto3" json:"TokenExpiredAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupResponse) Reset()         { *m = SignupResponse{} }
func (m *SignupResponse) String() string { return proto.CompactTextString(m) }
func (*SignupResponse) ProtoMessage()    {}
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_secruity_c368f40297f61236, []int{1}
}
func (m *SignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupResponse.Unmarshal(m, b)
}
func (m *SignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupResponse.Marshal(b, m, deterministic)
}
func (dst *SignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupResponse.Merge(dst, src)
}
func (m *SignupResponse) XXX_Size() int {
	return xxx_messageInfo_SignupResponse.Size(m)
}
func (m *SignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignupResponse proto.InternalMessageInfo

func (m *SignupResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *SignupResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SignupResponse) GetTokenExpiredAt() int64 {
	if m != nil {
		return m.TokenExpiredAt
	}
	return 0
}

type AuthRequest struct {
	Identify             string   `protobuf:"bytes,1,opt,name=Identify,proto3" json:"Identify,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_secruity_c368f40297f61236, []int{2}
}
func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (dst *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(dst, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetIdentify() string {
	if m != nil {
		return m.Identify
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	ExpiredAt            int64    `protobuf:"varint,2,opt,name=ExpiredAt,proto3" json:"ExpiredAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_secruity_c368f40297f61236, []int{3}
}
func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (dst *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(dst, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetExpiredAt() int64 {
	if m != nil {
		return m.ExpiredAt
	}
	return 0
}

func init() {
	proto.RegisterType((*SignupWithMobileRequest)(nil), "iunite.club.srv.user.secruity.SignupWithMobileRequest")
	proto.RegisterType((*SignupResponse)(nil), "iunite.club.srv.user.secruity.SignupResponse")
	proto.RegisterType((*AuthRequest)(nil), "iunite.club.srv.user.secruity.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "iunite.club.srv.user.secruity.AuthResponse")
}

func init() {
	proto.RegisterFile("proto/secruity/secruity.proto", fileDescriptor_secruity_c368f40297f61236)
}

var fileDescriptor_secruity_c368f40297f61236 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x4b, 0xe3, 0x40,
	0x14, 0xdd, 0xa4, 0xdb, 0x6c, 0x7a, 0xb7, 0x74, 0x97, 0x61, 0xd9, 0x0d, 0xa5, 0x85, 0x92, 0x87,
	0xa5, 0x28, 0x46, 0x50, 0xf0, 0xbd, 0x96, 0x0a, 0x52, 0xb5, 0x92, 0x8a, 0xbe, 0x09, 0x69, 0x32,
	0xb5, 0x83, 0x69, 0x26, 0xce, 0x4c, 0xd4, 0xe2, 0xef, 0xf0, 0x5f, 0xfa, 0x23, 0x24, 0x33, 0xf9,
	0x68, 0x03, 0x6a, 0xdf, 0xee, 0x39, 0x73, 0x72, 0x73, 0xee, 0x99, 0x3b, 0xd0, 0x8d, 0x19, 0x15,
	0x74, 0x9f, 0x63, 0x9f, 0x25, 0x44, 0xac, 0x8a, 0xc2, 0x91, 0x3c, 0xea, 0x92, 0x24, 0x22, 0x02,
	0x3b, 0x7e, 0x98, 0xcc, 0x1c, 0xce, 0x1e, 0x9d, 0x84, 0x63, 0xe6, 0xe4, 0x22, 0xfb, 0x55, 0x87,
	0x7f, 0x53, 0x72, 0x17, 0x25, 0xf1, 0x0d, 0x11, 0x8b, 0x73, 0x3a, 0x23, 0x21, 0x76, 0xf1, 0x43,
	0x82, 0xb9, 0x40, 0x6d, 0x30, 0x2f, 0x3d, 0xce, 0x9f, 0x28, 0x0b, 0x2c, 0xad, 0xa7, 0xf5, 0x1b,
	0x6e, 0x81, 0xd1, 0x5f, 0x30, 0xa6, 0xfe, 0x82, 0xd2, 0xd0, 0xd2, 0xe5, 0x49, 0x86, 0x52, 0x5e,
	0x35, 0xb1, 0x6a, 0x8a, 0x57, 0x08, 0xfd, 0x87, 0x96, 0xaa, 0x06, 0x0c, 0x7b, 0x43, 0x1a, 0x60,
	0xeb, 0x7b, 0x4f, 0xeb, 0xd7, 0xdd, 0x0a, 0x8b, 0x6c, 0x68, 0x5e, 0x7b, 0x21, 0x09, 0x3c, 0x81,
	0xa5, 0xaa, 0x2e, 0xbb, 0x6c, 0x70, 0xa8, 0x03, 0x8d, 0x13, 0xc2, 0xb8, 0xb8, 0xf0, 0x96, 0xd8,
	0x32, 0xa4, 0xa0, 0x24, 0x52, 0xd7, 0x67, 0x5e, 0x76, 0xf8, 0x43, 0xb9, 0xce, 0x31, 0xea, 0xc3,
	0xaf, 0x21, 0x8d, 0xe6, 0x84, 0x2d, 0x8b, 0xc1, 0x4c, 0x29, 0xa9, 0xd2, 0xf6, 0x2d, 0xb4, 0x54,
	0x2c, 0x2e, 0xe6, 0x31, 0x8d, 0x38, 0x46, 0x2d, 0xd0, 0x27, 0x63, 0x99, 0x83, 0xe9, 0xea, 0x93,
	0x31, 0xfa, 0x03, 0xf5, 0x2b, 0x7a, 0x8f, 0xa3, 0x2c, 0x00, 0x05, 0xd2, 0x39, 0x65, 0x31, 0x7a,
	0x8e, 0x09, 0xc3, 0xc1, 0x40, 0xc8, 0x1c, 0x6a, 0x6e, 0x85, 0xb5, 0x47, 0xf0, 0x73, 0x90, 0x88,
	0xc5, 0x5a, 0xd4, 0xa7, 0x01, 0x8e, 0x04, 0x99, 0xaf, 0xf2, 0xa8, 0x73, 0xbc, 0x71, 0x0d, 0xfa,
	0xe6, 0x35, 0xd8, 0xc7, 0xd0, 0x54, 0x6d, 0x32, 0x93, 0x85, 0x29, 0x6d, 0xdd, 0x54, 0x07, 0x1a,
	0xa5, 0x1f, 0x5d, 0xfa, 0x29, 0x89, 0x83, 0x37, 0x0d, 0xcc, 0x69, 0xb6, 0x0f, 0xc8, 0x07, 0x23,
	0x9d, 0x9b, 0x44, 0x68, 0xc7, 0xf9, 0x74, 0x73, 0x9c, 0x35, 0xfb, 0xed, 0xdd, 0xad, 0xb4, 0xca,
	0xa3, 0xfd, 0x0d, 0xbd, 0xc0, 0xef, 0xea, 0xce, 0xa1, 0xa3, 0x2f, 0x5a, 0x7c, 0xb0, 0xa4, 0xed,
	0xbd, 0xad, 0xbe, 0x2b, 0x7f, 0x3e, 0x33, 0xe4, 0xbb, 0x38, 0x7c, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0x01, 0xb2, 0xdc, 0x38, 0x03, 0x00, 0x00,
}