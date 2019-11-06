// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getArea.proto

package go_micro_srv_getArea

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_420050f84ccd68e8, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 []*Area  `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_420050f84ccd68e8, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() []*Area {
	if m != nil {
		return m.Data
	}
	return nil
}

type Area struct {
	Aid                  int32    `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Aname                string   `protobuf:"bytes,2,opt,name=aname,proto3" json:"aname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_420050f84ccd68e8, []int{2}
}

func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (m *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(m, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetAid() int32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *Area) GetAname() string {
	if m != nil {
		return m.Aname
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.getArea.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.getArea.Response")
	proto.RegisterType((*Area)(nil), "go.micro.srv.getArea.area")
}

func init() { proto.RegisterFile("getArea.proto", fileDescriptor_420050f84ccd68e8) }

var fileDescriptor_420050f84ccd68e8 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4e, 0x84, 0x30,
	0x10, 0x86, 0x5d, 0x61, 0x77, 0x65, 0x8c, 0x89, 0x69, 0x88, 0x21, 0x24, 0x1a, 0xd2, 0x13, 0xa7,
	0x39, 0xe0, 0x13, 0x18, 0x0f, 0xdc, 0xfb, 0x06, 0xa3, 0x4c, 0x90, 0x04, 0x28, 0x4e, 0xab, 0xcf,
	0x6f, 0x68, 0x7b, 0xc4, 0x5b, 0xbf, 0xfe, 0x93, 0xf9, 0xf2, 0x0f, 0x3c, 0x8c, 0xec, 0xdf, 0x84,
	0x09, 0x37, 0xb1, 0xde, 0xaa, 0x72, 0xb4, 0xb8, 0x4c, 0x9f, 0x62, 0xd1, 0xc9, 0x2f, 0xa6, 0x4c,
	0x17, 0x70, 0x35, 0xfc, 0xfd, 0xc3, 0xce, 0xeb, 0x2f, 0xb8, 0x33, 0xec, 0x36, 0xbb, 0x3a, 0x56,
	0x25, 0x9c, 0x59, 0x64, 0xb5, 0xd5, 0xa9, 0x39, 0xb5, 0x85, 0x89, 0xa0, 0x9e, 0xe0, 0xc2, 0x22,
	0x8b, 0x1b, 0xab, 0xdb, 0xf0, 0x9d, 0x48, 0x21, 0xe4, 0x03, 0x79, 0xaa, 0xb2, 0x26, 0x6b, 0xef,
	0xbb, 0x1a, 0x8f, 0x4c, 0x48, 0xc2, 0x64, 0xc2, 0x9c, 0x46, 0xc8, 0x77, 0x52, 0x8f, 0x90, 0xd1,
	0x34, 0x04, 0xc7, 0xd9, 0xec, 0xcf, 0xdd, 0x4b, 0x2b, 0x2d, 0x9c, 0x04, 0x11, 0x3a, 0x03, 0xd7,
	0x3e, 0x6e, 0x51, 0x3d, 0xe4, 0xef, 0x34, 0xcf, 0xea, 0xf9, 0x58, 0x92, 0xba, 0xd4, 0x2f, 0xff,
	0xc5, 0xb1, 0x9f, 0xbe, 0xf9, 0xb8, 0x84, 0xab, 0xbc, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfc,
	0xdc, 0x8f, 0xea, 0x26, 0x01, 0x00, 0x00,
}