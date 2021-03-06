// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/storage.proto

package iunite_club_srv_storage

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

type File struct {
	Filename             string   `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{0}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type FilePB struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FileKey              string   `protobuf:"bytes,2,opt,name=FileKey,proto3" json:"FileKey,omitempty"`
	Storage              string   `protobuf:"bytes,3,opt,name=Storage,proto3" json:"Storage,omitempty"`
	Ext                  string   `protobuf:"bytes,4,opt,name=Ext,proto3" json:"Ext,omitempty"`
	Host                 string   `protobuf:"bytes,5,opt,name=Host,proto3" json:"Host,omitempty"`
	Bulket               string   `protobuf:"bytes,6,opt,name=Bulket,proto3" json:"Bulket,omitempty"`
	OriginalFilename     string   `protobuf:"bytes,7,opt,name=OriginalFilename,proto3" json:"OriginalFilename,omitempty"`
	Size                 int64    `protobuf:"varint,8,opt,name=Size,proto3" json:"Size,omitempty"`
	Path                 string   `protobuf:"bytes,9,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePB) Reset()         { *m = FilePB{} }
func (m *FilePB) String() string { return proto.CompactTextString(m) }
func (*FilePB) ProtoMessage()    {}
func (*FilePB) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{1}
}
func (m *FilePB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePB.Unmarshal(m, b)
}
func (m *FilePB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePB.Marshal(b, m, deterministic)
}
func (dst *FilePB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePB.Merge(dst, src)
}
func (m *FilePB) XXX_Size() int {
	return xxx_messageInfo_FilePB.Size(m)
}
func (m *FilePB) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePB.DiscardUnknown(m)
}

var xxx_messageInfo_FilePB proto.InternalMessageInfo

func (m *FilePB) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *FilePB) GetFileKey() string {
	if m != nil {
		return m.FileKey
	}
	return ""
}

func (m *FilePB) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *FilePB) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

func (m *FilePB) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *FilePB) GetBulket() string {
	if m != nil {
		return m.Bulket
	}
	return ""
}

func (m *FilePB) GetOriginalFilename() string {
	if m != nil {
		return m.OriginalFilename
	}
	return ""
}

func (m *FilePB) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FilePB) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type FileInfoRequest struct {
	File                 *FilePB  `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfoRequest) Reset()         { *m = FileInfoRequest{} }
func (m *FileInfoRequest) String() string { return proto.CompactTextString(m) }
func (*FileInfoRequest) ProtoMessage()    {}
func (*FileInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{2}
}
func (m *FileInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfoRequest.Unmarshal(m, b)
}
func (m *FileInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfoRequest.Marshal(b, m, deterministic)
}
func (dst *FileInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfoRequest.Merge(dst, src)
}
func (m *FileInfoRequest) XXX_Size() int {
	return xxx_messageInfo_FileInfoRequest.Size(m)
}
func (m *FileInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfoRequest proto.InternalMessageInfo

func (m *FileInfoRequest) GetFile() *FilePB {
	if m != nil {
		return m.File
	}
	return nil
}

type FileRequest struct {
	File                 *File    `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{3}
}
func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (dst *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(dst, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

type FilesRequest struct {
	// File File = 1;
	Files                []*File  `protobuf:"bytes,2,rep,name=Files,proto3" json:"Files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilesRequest) Reset()         { *m = FilesRequest{} }
func (m *FilesRequest) String() string { return proto.CompactTextString(m) }
func (*FilesRequest) ProtoMessage()    {}
func (*FilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{4}
}
func (m *FilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilesRequest.Unmarshal(m, b)
}
func (m *FilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilesRequest.Marshal(b, m, deterministic)
}
func (dst *FilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilesRequest.Merge(dst, src)
}
func (m *FilesRequest) XXX_Size() int {
	return xxx_messageInfo_FilesRequest.Size(m)
}
func (m *FilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilesRequest proto.InternalMessageInfo

func (m *FilesRequest) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type FileResponse struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	File                 *FilePB  `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponse) Reset()         { *m = FileResponse{} }
func (m *FileResponse) String() string { return proto.CompactTextString(m) }
func (*FileResponse) ProtoMessage()    {}
func (*FileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{5}
}
func (m *FileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponse.Unmarshal(m, b)
}
func (m *FileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponse.Marshal(b, m, deterministic)
}
func (dst *FileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponse.Merge(dst, src)
}
func (m *FileResponse) XXX_Size() int {
	return xxx_messageInfo_FileResponse.Size(m)
}
func (m *FileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponse proto.InternalMessageInfo

func (m *FileResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *FileResponse) GetFile() *FilePB {
	if m != nil {
		return m.File
	}
	return nil
}

type RepeatedFileResponse struct {
	OK                   bool      `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	Files                []*FilePB `protobuf:"bytes,2,rep,name=Files,proto3" json:"Files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RepeatedFileResponse) Reset()         { *m = RepeatedFileResponse{} }
func (m *RepeatedFileResponse) String() string { return proto.CompactTextString(m) }
func (*RepeatedFileResponse) ProtoMessage()    {}
func (*RepeatedFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_c90ae75d3ee1d6fc, []int{6}
}
func (m *RepeatedFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedFileResponse.Unmarshal(m, b)
}
func (m *RepeatedFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedFileResponse.Marshal(b, m, deterministic)
}
func (dst *RepeatedFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedFileResponse.Merge(dst, src)
}
func (m *RepeatedFileResponse) XXX_Size() int {
	return xxx_messageInfo_RepeatedFileResponse.Size(m)
}
func (m *RepeatedFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedFileResponse proto.InternalMessageInfo

func (m *RepeatedFileResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *RepeatedFileResponse) GetFiles() []*FilePB {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "iunite.club.srv.storage.File")
	proto.RegisterType((*FilePB)(nil), "iunite.club.srv.storage.FilePB")
	proto.RegisterType((*FileInfoRequest)(nil), "iunite.club.srv.storage.FileInfoRequest")
	proto.RegisterType((*FileRequest)(nil), "iunite.club.srv.storage.FileRequest")
	proto.RegisterType((*FilesRequest)(nil), "iunite.club.srv.storage.FilesRequest")
	proto.RegisterType((*FileResponse)(nil), "iunite.club.srv.storage.FileResponse")
	proto.RegisterType((*RepeatedFileResponse)(nil), "iunite.club.srv.storage.RepeatedFileResponse")
}

func init() { proto.RegisterFile("proto/storage.proto", fileDescriptor_storage_c90ae75d3ee1d6fc) }

var fileDescriptor_storage_c90ae75d3ee1d6fc = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x3f, 0xea, 0xda, 0xd3, 0x08, 0xaa, 0x01, 0xc1, 0x2a, 0x12, 0xc2, 0xb2, 0x40, 0xb2,
	0x90, 0x30, 0xa2, 0x15, 0x77, 0xd4, 0x96, 0x8a, 0x28, 0x87, 0x44, 0xeb, 0x13, 0x07, 0x0e, 0x4e,
	0x18, 0x82, 0x85, 0xb1, 0x83, 0x77, 0x1d, 0x01, 0xbf, 0x89, 0xdf, 0xc4, 0x6f, 0x41, 0xbb, 0x6b,
	0x9b, 0x44, 0x10, 0x87, 0x9e, 0x76, 0xde, 0xec, 0xcc, 0xdb, 0xf7, 0xc6, 0x23, 0xc3, 0xbd, 0x75,
	0x5d, 0xc9, 0xea, 0x85, 0x90, 0x55, 0x9d, 0xad, 0x28, 0xd1, 0x08, 0x1f, 0xe6, 0x4d, 0x99, 0x4b,
	0x4a, 0x96, 0x45, 0xb3, 0x48, 0x44, 0xbd, 0x49, 0xda, 0xeb, 0x68, 0x0e, 0xee, 0x4d, 0x5e, 0x10,
	0x8e, 0xc1, 0x57, 0x67, 0x99, 0x7d, 0x21, 0x66, 0x85, 0x56, 0x1c, 0xf0, 0x1e, 0x23, 0x82, 0x9b,
	0xe6, 0x3f, 0x88, 0xd9, 0xa1, 0x15, 0x3b, 0x5c, 0xc7, 0xc8, 0xe0, 0xe4, 0xaa, 0x2a, 0x25, 0x95,
	0x92, 0x39, 0xa1, 0x15, 0x8f, 0x78, 0x07, 0xa3, 0x5f, 0x16, 0x78, 0xaa, 0x75, 0x7e, 0x89, 0x77,
	0xc0, 0x9e, 0x5c, 0xb7, 0x74, 0xf6, 0xe4, 0x5a, 0x35, 0xa9, 0x9b, 0x29, 0x7d, 0xd7, 0x5c, 0x01,
	0xef, 0xa0, 0xba, 0x49, 0x8d, 0x22, 0x4d, 0x17, 0xf0, 0x0e, 0xe2, 0x19, 0x38, 0x6f, 0xbe, 0x49,
	0xe6, 0xea, 0xac, 0x0a, 0x95, 0x9c, 0xb7, 0x95, 0x90, 0xec, 0x58, 0xa7, 0x74, 0x8c, 0x0f, 0xc0,
	0xbb, 0x6c, 0x8a, 0xcf, 0x24, 0x99, 0xa7, 0xb3, 0x2d, 0xc2, 0x67, 0x70, 0x36, 0xab, 0xf3, 0x55,
	0x5e, 0x66, 0x45, 0x6f, 0xef, 0x44, 0x57, 0xfc, 0x95, 0xef, 0x6d, 0xfa, 0x5b, 0x36, 0x11, 0xdc,
	0x79, 0x26, 0x3f, 0xb1, 0xc0, 0xbc, 0xa5, 0xe2, 0xe8, 0x06, 0xee, 0xaa, 0x9e, 0x49, 0xf9, 0xb1,
	0xe2, 0xf4, 0xb5, 0x21, 0x21, 0xf1, 0xc2, 0x4c, 0x51, 0x5b, 0x3d, 0x3d, 0x7f, 0x9c, 0xec, 0x99,
	0x76, 0x62, 0xe6, 0xc2, 0x75, 0x71, 0xf4, 0x1a, 0x4e, 0xd5, 0xd9, 0x71, 0xbc, 0xdc, 0xe1, 0x78,
	0x34, 0xc8, 0xd1, 0x32, 0x5c, 0xc1, 0x48, 0x9d, 0xe2, 0x8f, 0x8c, 0x63, 0x8d, 0x99, 0x1d, 0x3a,
	0x87, 0x39, 0x4c, 0x6d, 0x94, 0x1a, 0x12, 0x4e, 0x62, 0x5d, 0x95, 0x82, 0xd4, 0x47, 0x9b, 0x4d,
	0xb5, 0x0a, 0x9f, 0xdb, 0xb3, 0x69, 0xef, 0xcd, 0xbe, 0x8d, 0xb7, 0xf7, 0x70, 0x9f, 0xd3, 0x9a,
	0x32, 0x49, 0x1f, 0x06, 0xc9, 0x5f, 0xed, 0x2a, 0x3e, 0xc8, 0x6e, 0xaa, 0xcf, 0x7f, 0xda, 0xfd,
	0xbe, 0xe0, 0x3b, 0xf0, 0xd3, 0x6c, 0x43, 0x7a, 0x8b, 0x9f, 0x0c, 0x3b, 0x36, 0x63, 0x1a, 0x3f,
	0x3d, 0x50, 0x65, 0xb4, 0x46, 0x47, 0xb8, 0x84, 0xa0, 0xa3, 0x16, 0x38, 0xdc, 0xd5, 0x7d, 0x83,
	0xf1, 0xf3, 0xbd, 0x65, 0xff, 0x1a, 0x48, 0x74, 0x84, 0x19, 0x8c, 0xba, 0x47, 0xd4, 0x4a, 0x61,
	0x3c, 0xf8, 0xce, 0xd6, 0xd6, 0xfd, 0xb7, 0x8f, 0x85, 0xa7, 0x7f, 0x02, 0x17, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0x58, 0x7d, 0xe8, 0x1b, 0x04, 0x00, 0x00,
}
