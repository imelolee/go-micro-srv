// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/userOrder.proto

package userOrder

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

// Api Endpoints for UserOrder service

func NewUserOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserOrder service

type UserOrderService interface {
	CreateOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetOrderInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type userOrderService struct {
	c    client.Client
	name string
}

func NewUserOrderService(name string, c client.Client) UserOrderService {
	return &userOrderService{
		c:    c,
		name: name,
	}
}

func (c *userOrderService) CreateOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserOrder.CreateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) GetOrderInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrder.GetOrderInfo", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) UpdateStatus(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrder.UpdateStatus", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserOrder service

type UserOrderHandler interface {
	CreateOrder(context.Context, *Request, *Response) error
	GetOrderInfo(context.Context, *GetRequest, *GetResponse) error
	UpdateStatus(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterUserOrderHandler(s server.Server, hdlr UserOrderHandler, opts ...server.HandlerOption) error {
	type userOrder interface {
		CreateOrder(ctx context.Context, in *Request, out *Response) error
		GetOrderInfo(ctx context.Context, in *GetRequest, out *GetResponse) error
		UpdateStatus(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
	}
	type UserOrder struct {
		userOrder
	}
	h := &userOrderHandler{hdlr}
	return s.Handle(s.NewHandler(&UserOrder{h}, opts...))
}

type userOrderHandler struct {
	UserOrderHandler
}

func (h *userOrderHandler) CreateOrder(ctx context.Context, in *Request, out *Response) error {
	return h.UserOrderHandler.CreateOrder(ctx, in, out)
}

func (h *userOrderHandler) GetOrderInfo(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.UserOrderHandler.GetOrderInfo(ctx, in, out)
}

func (h *userOrderHandler) UpdateStatus(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UserOrderHandler.UpdateStatus(ctx, in, out)
}
