// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

package user

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*RegResponse, error)
	Register(ctx context.Context, in *RegRequest, opts ...client.CallOption) (*RegResponse, error)
	Login(ctx context.Context, in *RegRequest, opts ...client.CallOption) (*RegResponse, error)
	MicroGetUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateUserName(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	UploadAvatar(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error)
	AuthUpdate(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*RegResponse, error) {
	req := c.c.NewRequest(c.name, "User.SendSms", in)
	out := new(RegResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Register(ctx context.Context, in *RegRequest, opts ...client.CallOption) (*RegResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(RegResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *RegRequest, opts ...client.CallOption) (*RegResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(RegResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MicroGetUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.MicroGetUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserName(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserName", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UploadAvatar(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error) {
	req := c.c.NewRequest(c.name, "User.UploadAvatar", in)
	out := new(UploadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AuthUpdate(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "User.AuthUpdate", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	SendSms(context.Context, *SmsRequest, *RegResponse) error
	Register(context.Context, *RegRequest, *RegResponse) error
	Login(context.Context, *RegRequest, *RegResponse) error
	MicroGetUser(context.Context, *Request, *Response) error
	UpdateUserName(context.Context, *UpdateRequest, *UpdateResponse) error
	UploadAvatar(context.Context, *UploadRequest, *UploadResponse) error
	AuthUpdate(context.Context, *AuthRequest, *AuthResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		SendSms(ctx context.Context, in *SmsRequest, out *RegResponse) error
		Register(ctx context.Context, in *RegRequest, out *RegResponse) error
		Login(ctx context.Context, in *RegRequest, out *RegResponse) error
		MicroGetUser(ctx context.Context, in *Request, out *Response) error
		UpdateUserName(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		UploadAvatar(ctx context.Context, in *UploadRequest, out *UploadResponse) error
		AuthUpdate(ctx context.Context, in *AuthRequest, out *AuthResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) SendSms(ctx context.Context, in *SmsRequest, out *RegResponse) error {
	return h.UserHandler.SendSms(ctx, in, out)
}

func (h *userHandler) Register(ctx context.Context, in *RegRequest, out *RegResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *RegRequest, out *RegResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) MicroGetUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserHandler.MicroGetUser(ctx, in, out)
}

func (h *userHandler) UpdateUserName(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UserHandler.UpdateUserName(ctx, in, out)
}

func (h *userHandler) UploadAvatar(ctx context.Context, in *UploadRequest, out *UploadResponse) error {
	return h.UserHandler.UploadAvatar(ctx, in, out)
}

func (h *userHandler) AuthUpdate(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.UserHandler.AuthUpdate(ctx, in, out)
}
