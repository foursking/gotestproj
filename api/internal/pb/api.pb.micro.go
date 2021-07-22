// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package pb

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

// Client API for UserService service

type UserService interface {
	//增
	InsertUser(ctx context.Context, in *InsertUserReq, opts ...client.CallOption) (*InsertUserRep, error)
	//删
	DeletetUser(ctx context.Context, in *DeletetUserReq, opts ...client.CallOption) (*DeletetUserRep, error)
	//查
	SelectUser(ctx context.Context, in *SelectUserReq, opts ...client.CallOption) (*SelectUserRep, error)
	//改
	ModifyUser(ctx context.Context, in *ModifyUserReq, opts ...client.CallOption) (*ModifyUserRep, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pb"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) InsertUser(ctx context.Context, in *InsertUserReq, opts ...client.CallOption) (*InsertUserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.InsertUser", in)
	out := new(InsertUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeletetUser(ctx context.Context, in *DeletetUserReq, opts ...client.CallOption) (*DeletetUserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.DeletetUser", in)
	out := new(DeletetUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SelectUser(ctx context.Context, in *SelectUserReq, opts ...client.CallOption) (*SelectUserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.SelectUser", in)
	out := new(SelectUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ModifyUser(ctx context.Context, in *ModifyUserReq, opts ...client.CallOption) (*ModifyUserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.ModifyUser", in)
	out := new(ModifyUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//增
	InsertUser(context.Context, *InsertUserReq, *InsertUserRep) error
	//删
	DeletetUser(context.Context, *DeletetUserReq, *DeletetUserRep) error
	//查
	SelectUser(context.Context, *SelectUserReq, *SelectUserRep) error
	//改
	ModifyUser(context.Context, *ModifyUserReq, *ModifyUserRep) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		InsertUser(ctx context.Context, in *InsertUserReq, out *InsertUserRep) error
		DeletetUser(ctx context.Context, in *DeletetUserReq, out *DeletetUserRep) error
		SelectUser(ctx context.Context, in *SelectUserReq, out *SelectUserRep) error
		ModifyUser(ctx context.Context, in *ModifyUserReq, out *ModifyUserRep) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) InsertUser(ctx context.Context, in *InsertUserReq, out *InsertUserRep) error {
	return h.UserServiceHandler.InsertUser(ctx, in, out)
}

func (h *userServiceHandler) DeletetUser(ctx context.Context, in *DeletetUserReq, out *DeletetUserRep) error {
	return h.UserServiceHandler.DeletetUser(ctx, in, out)
}

func (h *userServiceHandler) SelectUser(ctx context.Context, in *SelectUserReq, out *SelectUserRep) error {
	return h.UserServiceHandler.SelectUser(ctx, in, out)
}

func (h *userServiceHandler) ModifyUser(ctx context.Context, in *ModifyUserReq, out *ModifyUserRep) error {
	return h.UserServiceHandler.ModifyUser(ctx, in, out)
}
