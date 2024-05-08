// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: moego/business/order/v1/order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status
type Order_Status int32

const (
	// UNSPECIFIED
	Order_STATUS_UNSPECIFIED Order_Status = 0
	// CREATED
	Order_CREATED Order_Status = 1
	// PROCESSING
	Order_PROCESSING Order_Status = 2
	// COMPLETED
	Order_COMPLETED Order_Status = 3
	// REMOVED
	Order_REMOVED Order_Status = 4
)

// Enum value maps for Order_Status.
var (
	Order_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATED",
		2: "PROCESSING",
		3: "COMPLETED",
		4: "REMOVED",
	}
	Order_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"PROCESSING":         2,
		"COMPLETED":          3,
		"REMOVED":            4,
	}
)

func (x Order_Status) Enum() *Order_Status {
	p := new(Order_Status)
	*p = x
	return p
}

func (x Order_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_order_v1_order_proto_enumTypes[0].Descriptor()
}

func (Order_Status) Type() protoreflect.EnumType {
	return &file_moego_business_order_v1_order_proto_enumTypes[0]
}

func (x Order_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_Status.Descriptor instead.
func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_order_v1_order_proto_rawDescGZIP(), []int{0, 0}
}

// Agreement
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// business id
	BusinessId string `protobuf:"bytes,2,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	// customer id
	CustomerId string `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// order status
	Status Order_Status `protobuf:"varint,4,opt,name=status,proto3,enum=moego.business.order.v1.Order_Status" json:"status,omitempty"`
	// tips amount
	TipsAmount float64 `protobuf:"fixed64,5,opt,name=tips_amount,json=tipsAmount,proto3" json:"tips_amount,omitempty"`
	// tax amount
	TaxAmount float64 `protobuf:"fixed64,6,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	// discount amount
	DiscountAmount float64 `protobuf:"fixed64,7,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	// extra fee amount
	ExtraFeeAmount float64 `protobuf:"fixed64,8,opt,name=extra_fee_amount,json=extraFeeAmount,proto3" json:"extra_fee_amount,omitempty"`
	// subTotal amount
	SubTotalAmount float64 `protobuf:"fixed64,9,opt,name=sub_total_amount,json=subTotalAmount,proto3" json:"sub_total_amount,omitempty"`
	// amount of tips based on
	TipsBasedAmount float64 `protobuf:"fixed64,10,opt,name=tips_based_amount,json=tipsBasedAmount,proto3" json:"tips_based_amount,omitempty"`
	// total amount
	TotalAmount float64 `protobuf:"fixed64,11,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	// paid amount
	PaidAmount float64 `protobuf:"fixed64,12,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	// remain amount to pay
	RemainAmount float64 `protobuf:"fixed64,13,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`
	// refunded amount
	RefundedAmount float64 `protobuf:"fixed64,14,opt,name=refunded_amount,json=refundedAmount,proto3" json:"refunded_amount,omitempty"`
	// order title
	Title string `protobuf:"bytes,15,opt,name=title,proto3" json:"title,omitempty"`
	// staff id of creating this order
	CreateBy *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	// staff id of last updating this order
	UpdateBy *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	// create time
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// update time
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// description
	Description string `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_order_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_order_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_moego_business_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Order) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Order) GetStatus() Order_Status {
	if x != nil {
		return x.Status
	}
	return Order_STATUS_UNSPECIFIED
}

func (x *Order) GetTipsAmount() float64 {
	if x != nil {
		return x.TipsAmount
	}
	return 0
}

func (x *Order) GetTaxAmount() float64 {
	if x != nil {
		return x.TaxAmount
	}
	return 0
}

func (x *Order) GetDiscountAmount() float64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *Order) GetExtraFeeAmount() float64 {
	if x != nil {
		return x.ExtraFeeAmount
	}
	return 0
}

func (x *Order) GetSubTotalAmount() float64 {
	if x != nil {
		return x.SubTotalAmount
	}
	return 0
}

func (x *Order) GetTipsBasedAmount() float64 {
	if x != nil {
		return x.TipsBasedAmount
	}
	return 0
}

func (x *Order) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Order) GetPaidAmount() float64 {
	if x != nil {
		return x.PaidAmount
	}
	return 0
}

func (x *Order) GetRemainAmount() float64 {
	if x != nil {
		return x.RemainAmount
	}
	return 0
}

func (x *Order) GetRefundedAmount() float64 {
	if x != nil {
		return x.RefundedAmount
	}
	return 0
}

func (x *Order) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Order) GetCreateBy() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateBy
	}
	return nil
}

func (x *Order) GetUpdateBy() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateBy
	}
	return nil
}

func (x *Order) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Order) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_moego_business_order_v1_order_proto protoreflect.FileDescriptor

var file_moego_business_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x07, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69,
	0x70, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x74, 0x69, 0x70, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x74, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x65, 0x65,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x69, 0x70, 0x73, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0f, 0x74, 0x69, 0x70, 0x73, 0x42, 0x61, 0x73, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x69,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x44, 0x10, 0x04, 0x42, 0x80, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x3b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_order_v1_order_proto_rawDescOnce sync.Once
	file_moego_business_order_v1_order_proto_rawDescData = file_moego_business_order_v1_order_proto_rawDesc
)

func file_moego_business_order_v1_order_proto_rawDescGZIP() []byte {
	file_moego_business_order_v1_order_proto_rawDescOnce.Do(func() {
		file_moego_business_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_order_v1_order_proto_rawDescData)
	})
	return file_moego_business_order_v1_order_proto_rawDescData
}

var file_moego_business_order_v1_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_moego_business_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_moego_business_order_v1_order_proto_goTypes = []interface{}{
	(Order_Status)(0),             // 0: moego.business.order.v1.Order.Status
	(*Order)(nil),                 // 1: moego.business.order.v1.Order
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_moego_business_order_v1_order_proto_depIdxs = []int32{
	0, // 0: moego.business.order.v1.Order.status:type_name -> moego.business.order.v1.Order.Status
	2, // 1: moego.business.order.v1.Order.create_by:type_name -> google.protobuf.Timestamp
	2, // 2: moego.business.order.v1.Order.update_by:type_name -> google.protobuf.Timestamp
	2, // 3: moego.business.order.v1.Order.create_time:type_name -> google.protobuf.Timestamp
	2, // 4: moego.business.order.v1.Order.update_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_moego_business_order_v1_order_proto_init() }
func file_moego_business_order_v1_order_proto_init() {
	if File_moego_business_order_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_business_order_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_moego_business_order_v1_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_order_v1_order_proto_goTypes,
		DependencyIndexes: file_moego_business_order_v1_order_proto_depIdxs,
		EnumInfos:         file_moego_business_order_v1_order_proto_enumTypes,
		MessageInfos:      file_moego_business_order_v1_order_proto_msgTypes,
	}.Build()
	File_moego_business_order_v1_order_proto = out.File
	file_moego_business_order_v1_order_proto_rawDesc = nil
	file_moego_business_order_v1_order_proto_goTypes = nil
	file_moego_business_order_v1_order_proto_depIdxs = nil
}
