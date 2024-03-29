// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/postLogin/postLogin.proto

package go_micro_srv_PostLogin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PostLogin service

type PostLoginService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type postLoginService struct {
	c    client.Client
	name string
}

func NewPostLoginService(name string, c client.Client) PostLoginService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.PostLogin"
	}
	return &postLoginService{
		c:    c,
		name: name,
	}
}

func (c *postLoginService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PostLogin.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostLogin service

type PostLoginHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterPostLoginHandler(s server.Server, hdlr PostLoginHandler, opts ...server.HandlerOption) error {
	type postLogin interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type PostLogin struct {
		postLogin
	}
	h := &postLoginHandler{hdlr}
	return s.Handle(s.NewHandler(&PostLogin{h}, opts...))
}

type postLoginHandler struct {
	PostLoginHandler
}

func (h *postLoginHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.PostLoginHandler.Call(ctx, in, out)
}
