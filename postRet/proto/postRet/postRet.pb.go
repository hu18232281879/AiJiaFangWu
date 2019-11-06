// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/postRet/postRet.proto

package go_micro_srv_postRet

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
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	SmsCode              string   `protobuf:"bytes,3,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_553841f8054e9341, []int{0}
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

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Request) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_553841f8054e9341, []int{1}
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

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.postRet.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.postRet.Response")
}

func init() { proto.RegisterFile("proto/postRet/postRet.proto", fileDescriptor_553841f8054e9341) }

var fileDescriptor_553841f8054e9341 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbd, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x8d, 0x1f, 0x49, 0x9c, 0x72, 0x09, 0x12, 0x23, 0x8a, 0xa4, 0xb2, 0x5a, 0x41, 0x1b,
	0xfb, 0x14, 0xb6, 0xb2, 0x95, 0x9d, 0xe4, 0x63, 0x08, 0x81, 0x6c, 0x66, 0x9d, 0xd9, 0xbb, 0xfb,
	0xf7, 0x8f, 0xdb, 0xec, 0x5d, 0x75, 0x57, 0x2d, 0xbf, 0xfd, 0xc1, 0x7b, 0x8f, 0x81, 0x27, 0xc7,
	0xe4, 0xe9, 0xdd, 0x91, 0x78, 0x83, 0xfe, 0xf8, 0xea, 0xf0, 0xab, 0x8a, 0x91, 0xb4, 0x9d, 0x7a,
	0x26, 0x2d, 0xbc, 0xd5, 0xd1, 0xd5, 0xbf, 0x90, 0x19, 0xfc, 0xdf, 0xa0, 0x78, 0xf5, 0x00, 0xa9,
	0xa5, 0x6e, 0x9a, 0xb1, 0x4c, 0x5e, 0x93, 0xb7, 0x7b, 0x13, 0x49, 0x55, 0x90, 0xbb, 0x56, 0x64,
	0x47, 0x3c, 0x94, 0xd7, 0xc1, 0x9c, 0x58, 0x3d, 0x42, 0x2e, 0x56, 0xfe, 0x7a, 0x1a, 0xb0, 0xbc,
	0x09, 0x2e, 0x13, 0x2b, 0x0d, 0x0d, 0x58, 0x7f, 0x41, 0x6e, 0x50, 0x1c, 0x2d, 0x82, 0xaa, 0x80,
	0x3b, 0x64, 0x5e, 0x28, 0x26, 0xaf, 0x70, 0x28, 0x44, 0x66, 0x2b, 0x63, 0x8c, 0x8d, 0xf4, 0x61,
	0x20, 0xfb, 0x59, 0xe7, 0xa9, 0x6f, 0xb8, 0x6d, 0xda, 0x79, 0x56, 0xcf, 0xfa, 0xdc, 0x7a, 0x1d,
	0xa7, 0x57, 0x2f, 0x97, 0xf4, 0xda, 0x5f, 0x5f, 0x75, 0x69, 0x38, 0xc2, 0xe7, 0x3e, 0x00, 0x00,
	0xff, 0xff, 0x67, 0x27, 0x1f, 0xa7, 0x23, 0x01, 0x00, 0x00,
}
