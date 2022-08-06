// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: proto/order.proto

package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Id     string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *UpdateRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateResponse) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *UpdateResponse) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role     string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *GetData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *GetResponse) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *GetResponse) GetData() *GetData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrdersData `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetData) Reset() {
	*x = GetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetData) ProtoMessage() {}

func (x *GetData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetData.ProtoReflect.Descriptor instead.
func (*GetData) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetData) GetOrders() []*OrdersData {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrdersData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Comment   string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime     string `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Days      int32  `protobuf:"varint,4,opt,name=days,proto3" json:"days,omitempty"`
	EndDate   string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	ImgUrl    string `protobuf:"bytes,6,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	OrderId   int32  `protobuf:"varint,7,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	StartDate string `protobuf:"bytes,8,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Status    string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Title     string `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *OrdersData) Reset() {
	*x = OrdersData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersData) ProtoMessage() {}

func (x *OrdersData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersData.ProtoReflect.Descriptor instead.
func (*OrdersData) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrdersData) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrdersData) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *OrdersData) GetCtime() string {
	if x != nil {
		return x.Ctime
	}
	return ""
}

func (x *OrdersData) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *OrdersData) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *OrdersData) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *OrdersData) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrdersData) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *OrdersData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrdersData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseId   string `protobuf:"bytes,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	StartDate string `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	UserName  string `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{6}
}

func (x *Request) GetHouseId() string {
	if x != nil {
		return x.HouseId
	}
	return ""
}

func (x *Request) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Request) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Request) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string     `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string     `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *OrderData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Response) GetData() *OrderData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderData) Reset() {
	*x = OrderData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderData) ProtoMessage() {}

func (x *OrderData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderData.ProtoReflect.Descriptor instead.
func (*OrderData) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{8}
}

func (x *OrderData) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_proto_order_proto protoreflect.FileDescriptor

var file_proto_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x3c, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55,
	0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x7a, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67,
	0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32,
	0x97, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData = file_proto_order_proto_rawDesc
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_proto_rawDescData)
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_order_proto_goTypes = []interface{}{
	(*UpdateRequest)(nil),  // 0: go.micro.srv.userOrder.UpdateRequest
	(*UpdateResponse)(nil), // 1: go.micro.srv.userOrder.UpdateResponse
	(*GetRequest)(nil),     // 2: go.micro.srv.userOrder.GetRequest
	(*GetResponse)(nil),    // 3: go.micro.srv.userOrder.GetResponse
	(*GetData)(nil),        // 4: go.micro.srv.userOrder.GetData
	(*OrdersData)(nil),     // 5: go.micro.srv.userOrder.OrdersData
	(*Request)(nil),        // 6: go.micro.srv.userOrder.Request
	(*Response)(nil),       // 7: go.micro.srv.userOrder.Response
	(*OrderData)(nil),      // 8: go.micro.srv.userOrder.OrderData
}
var file_proto_order_proto_depIdxs = []int32{
	4, // 0: go.micro.srv.userOrder.GetResponse.data:type_name -> go.micro.srv.userOrder.GetData
	5, // 1: go.micro.srv.userOrder.GetData.orders:type_name -> go.micro.srv.userOrder.OrdersData
	8, // 2: go.micro.srv.userOrder.Response.data:type_name -> go.micro.srv.userOrder.OrderData
	6, // 3: go.micro.srv.userOrder.Order.CreateOrder:input_type -> go.micro.srv.userOrder.Request
	2, // 4: go.micro.srv.userOrder.Order.GetOrderInfo:input_type -> go.micro.srv.userOrder.GetRequest
	0, // 5: go.micro.srv.userOrder.Order.UpdateStatus:input_type -> go.micro.srv.userOrder.UpdateRequest
	7, // 6: go.micro.srv.userOrder.Order.CreateOrder:output_type -> go.micro.srv.userOrder.Response
	3, // 7: go.micro.srv.userOrder.Order.GetOrderInfo:output_type -> go.micro.srv.userOrder.GetResponse
	1, // 8: go.micro.srv.userOrder.Order.UpdateStatus:output_type -> go.micro.srv.userOrder.UpdateResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_rawDesc = nil
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
