// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/house.proto

package house

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

// Api Endpoints for House service

func NewHouseEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for House service

type HouseService interface {
	PubHouse(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UploadHouseImg(ctx context.Context, in *ImgRequest, opts ...client.CallOption) (*ImgResponse, error)
	GetHouseInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	GetHouseDetail(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error)
	GetIndexHouse(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*GetResponse, error)
	SearchHouse(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*GetResponse, error)
}

type houseService struct {
	c    client.Client
	name string
}

func NewHouseService(name string, c client.Client) HouseService {
	return &houseService{
		c:    c,
		name: name,
	}
}

func (c *houseService) PubHouse(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "House.PubHouse", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseService) UploadHouseImg(ctx context.Context, in *ImgRequest, opts ...client.CallOption) (*ImgResponse, error) {
	req := c.c.NewRequest(c.name, "House.UploadHouseImg", in)
	out := new(ImgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseService) GetHouseInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "House.GetHouseInfo", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseService) GetHouseDetail(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error) {
	req := c.c.NewRequest(c.name, "House.GetHouseDetail", in)
	out := new(DetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseService) GetIndexHouse(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "House.GetIndexHouse", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseService) SearchHouse(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "House.SearchHouse", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for House service

type HouseHandler interface {
	PubHouse(context.Context, *Request, *Response) error
	UploadHouseImg(context.Context, *ImgRequest, *ImgResponse) error
	GetHouseInfo(context.Context, *GetRequest, *GetResponse) error
	GetHouseDetail(context.Context, *DetailRequest, *DetailResponse) error
	GetIndexHouse(context.Context, *IndexRequest, *GetResponse) error
	SearchHouse(context.Context, *SearchRequest, *GetResponse) error
}

func RegisterHouseHandler(s server.Server, hdlr HouseHandler, opts ...server.HandlerOption) error {
	type house interface {
		PubHouse(ctx context.Context, in *Request, out *Response) error
		UploadHouseImg(ctx context.Context, in *ImgRequest, out *ImgResponse) error
		GetHouseInfo(ctx context.Context, in *GetRequest, out *GetResponse) error
		GetHouseDetail(ctx context.Context, in *DetailRequest, out *DetailResponse) error
		GetIndexHouse(ctx context.Context, in *IndexRequest, out *GetResponse) error
		SearchHouse(ctx context.Context, in *SearchRequest, out *GetResponse) error
	}
	type House struct {
		house
	}
	h := &houseHandler{hdlr}
	return s.Handle(s.NewHandler(&House{h}, opts...))
}

type houseHandler struct {
	HouseHandler
}

func (h *houseHandler) PubHouse(ctx context.Context, in *Request, out *Response) error {
	return h.HouseHandler.PubHouse(ctx, in, out)
}

func (h *houseHandler) UploadHouseImg(ctx context.Context, in *ImgRequest, out *ImgResponse) error {
	return h.HouseHandler.UploadHouseImg(ctx, in, out)
}

func (h *houseHandler) GetHouseInfo(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.HouseHandler.GetHouseInfo(ctx, in, out)
}

func (h *houseHandler) GetHouseDetail(ctx context.Context, in *DetailRequest, out *DetailResponse) error {
	return h.HouseHandler.GetHouseDetail(ctx, in, out)
}

func (h *houseHandler) GetIndexHouse(ctx context.Context, in *IndexRequest, out *GetResponse) error {
	return h.HouseHandler.GetIndexHouse(ctx, in, out)
}

func (h *houseHandler) SearchHouse(ctx context.Context, in *SearchRequest, out *GetResponse) error {
	return h.HouseHandler.SearchHouse(ctx, in, out)
}