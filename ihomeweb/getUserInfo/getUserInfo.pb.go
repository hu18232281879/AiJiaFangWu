// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/getUserInfo/getUserInfo.proto

package go_micro_srv_getUserInfo

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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3b2d42a07109986, []int{0}
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Errno                string    `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string    `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 *UserInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3b2d42a07109986, []int{1}
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

func (m *Response) GetData() *UserInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserInfo struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	RealName             string   `protobuf:"bytes,4,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	IdCard               string   `protobuf:"bytes,5,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3b2d42a07109986, []int{2}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UserInfo) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *UserInfo) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.getUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.getUserInfo.Response")
	proto.RegisterType((*UserInfo)(nil), "go.micro.srv.getUserInfo.UserInfo")
}

func init() {
	proto.RegisterFile("proto/getUserInfo/getUserInfo.proto", fileDescriptor_f3b2d42a07109986)
}

var fileDescriptor_f3b2d42a07109986 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x4a, 0xf3, 0x40,
	0x10, 0xff, 0xd2, 0x2f, 0x4d, 0x9b, 0xe9, 0x6d, 0x11, 0xbb, 0x28, 0x85, 0xba, 0x5e, 0x7a, 0x5a,
	0xa1, 0x82, 0x2f, 0xd0, 0x83, 0xf4, 0xa2, 0xb0, 0xd0, 0xab, 0x61, 0xdb, 0x1d, 0x43, 0x60, 0x93,
	0x8d, 0xb3, 0x49, 0x9f, 0xc8, 0x07, 0x95, 0x6c, 0x52, 0xab, 0x87, 0x7a, 0x9b, 0xdf, 0x3f, 0xf6,
	0x37, 0xb3, 0x70, 0x5f, 0x93, 0x6b, 0xdc, 0x43, 0x8e, 0xcd, 0xce, 0x23, 0x6d, 0xab, 0xf7, 0x5f,
	0xb3, 0x0c, 0x2a, 0xe3, 0xb9, 0x93, 0x65, 0x71, 0x20, 0x27, 0x3d, 0x1d, 0xe5, 0x0f, 0x5d, 0x2c,
	0x60, 0xa2, 0xf0, 0xa3, 0x45, 0xdf, 0x30, 0x06, 0x71, 0xa5, 0x4b, 0xe4, 0xd1, 0x32, 0x5a, 0xa5,
	0x2a, 0xcc, 0xa2, 0x86, 0xa9, 0x42, 0x5f, 0xbb, 0xca, 0x23, 0xbb, 0x82, 0x31, 0x12, 0x55, 0x6e,
	0x30, 0xf4, 0x80, 0x5d, 0x43, 0x82, 0x44, 0xa5, 0xcf, 0xf9, 0x28, 0xd0, 0x03, 0x62, 0x4f, 0x10,
	0x1b, 0xdd, 0x68, 0xfe, 0x7f, 0x19, 0xad, 0x66, 0x6b, 0x21, 0x2f, 0x35, 0x90, 0xa7, 0x41, 0x05,
	0xbf, 0xf8, 0x8c, 0x60, 0x7a, 0xa2, 0xd8, 0x1c, 0x26, 0xad, 0x47, 0xca, 0x0a, 0x13, 0x1e, 0x1d,
	0xab, 0xa4, 0x83, 0x5b, 0xf3, 0xdd, 0x75, 0x74, 0xee, 0xda, 0x35, 0x29, 0xdd, 0xbe, 0xb0, 0x18,
	0xde, 0x4c, 0xd5, 0x80, 0xd8, 0x2d, 0xa4, 0x84, 0xda, 0x66, 0x21, 0x10, 0x07, 0x69, 0xda, 0x11,
	0x2f, 0x5d, 0x68, 0x0e, 0x93, 0xc2, 0x64, 0x07, 0x4d, 0x86, 0x8f, 0xfb, 0x54, 0x61, 0x36, 0x9a,
	0x0c, 0x5b, 0x00, 0xe8, 0xa3, 0x6e, 0x34, 0x65, 0x2d, 0x59, 0x9e, 0x04, 0x2d, 0xed, 0x99, 0x1d,
	0xd9, 0xf5, 0x1b, 0xcc, 0x9e, 0xcf, 0x4b, 0xb0, 0x57, 0x88, 0x37, 0xda, 0x5a, 0x76, 0x77, 0x79,
	0xcf, 0xe1, 0xcc, 0x37, 0xe2, 0x2f, 0x4b, 0x7f, 0x6a, 0xf1, 0x6f, 0x9f, 0x84, 0x8f, 0x7b, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x42, 0x49, 0xa1, 0xdf, 0x01, 0x00, 0x00,
}
