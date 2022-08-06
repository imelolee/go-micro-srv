// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order.proto

package order

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

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	CreateOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetOrderInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrderInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Order.GetOrderInfo", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateStatus(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateStatus", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	CreateOrder(context.Context, *Request, *Response) error
	GetOrderInfo(context.Context, *GetRequest, *GetResponse) error
	UpdateStatus(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		CreateOrder(ctx context.Context, in *Request, out *Response) error
		GetOrderInfo(ctx context.Context, in *GetRequest, out *GetResponse) error
		UpdateStatus(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *Request, out *Response) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}

func (h *orderHandler) GetOrderInfo(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.OrderHandler.GetOrderInfo(ctx, in, out)
}

func (h *orderHandler) UpdateStatus(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.OrderHandler.UpdateStatus(ctx, in, out)
}