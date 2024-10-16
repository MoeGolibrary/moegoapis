// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: moego/business/order/v1/order.proto

package orderpb

import (
	money "google.golang.org/genproto/googleapis/type/money"
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

// Order
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
	// order title
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	// description
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// tips amount
	TipsAmount *money.Money `protobuf:"bytes,7,opt,name=tips_amount,json=tipsAmount,proto3" json:"tips_amount,omitempty"`
	// tax amount
	TaxAmount *money.Money `protobuf:"bytes,8,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	// discount amount
	DiscountAmount *money.Money `protobuf:"bytes,9,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	// extra fee amount
	ExtraFeeAmount *money.Money `protobuf:"bytes,10,opt,name=extra_fee_amount,json=extraFeeAmount,proto3" json:"extra_fee_amount,omitempty"`
	// subTotal amount
	SubTotalAmount *money.Money `protobuf:"bytes,11,opt,name=sub_total_amount,json=subTotalAmount,proto3" json:"sub_total_amount,omitempty"`
	// amount of tips based on
	TipsBasedAmount *money.Money `protobuf:"bytes,12,opt,name=tips_based_amount,json=tipsBasedAmount,proto3" json:"tips_based_amount,omitempty"`
	// total amount
	TotalAmount *money.Money `protobuf:"bytes,13,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	// paid amount
	PaidAmount *money.Money `protobuf:"bytes,14,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	// remain amount to pay
	RemainAmount *money.Money `protobuf:"bytes,15,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`
	// refunded amount
	RefundedAmount *money.Money `protobuf:"bytes,16,opt,name=refunded_amount,json=refundedAmount,proto3" json:"refunded_amount,omitempty"`
	// staff id of creating this order
	CreatedBy string `protobuf:"bytes,17,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// create time
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// staff id of last updating this order
	LastUpdatedBy string `protobuf:"bytes,19,opt,name=last_updated_by,json=lastUpdatedBy,proto3" json:"last_updated_by,omitempty"`
	// update time
	LastUpdatedTime *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=last_updated_time,json=lastUpdatedTime,proto3" json:"last_updated_time,omitempty"`
	// sales datetime
	SalesDatetime *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=sales_datetime,json=salesDatetime,proto3" json:"sales_datetime,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_moego_business_order_v1_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_order_v1_order_proto_msgTypes[0]
	if x != nil {
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

func (x *Order) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetTipsAmount() *money.Money {
	if x != nil {
		return x.TipsAmount
	}
	return nil
}

func (x *Order) GetTaxAmount() *money.Money {
	if x != nil {
		return x.TaxAmount
	}
	return nil
}

func (x *Order) GetDiscountAmount() *money.Money {
	if x != nil {
		return x.DiscountAmount
	}
	return nil
}

func (x *Order) GetExtraFeeAmount() *money.Money {
	if x != nil {
		return x.ExtraFeeAmount
	}
	return nil
}

func (x *Order) GetSubTotalAmount() *money.Money {
	if x != nil {
		return x.SubTotalAmount
	}
	return nil
}

func (x *Order) GetTipsBasedAmount() *money.Money {
	if x != nil {
		return x.TipsBasedAmount
	}
	return nil
}

func (x *Order) GetTotalAmount() *money.Money {
	if x != nil {
		return x.TotalAmount
	}
	return nil
}

func (x *Order) GetPaidAmount() *money.Money {
	if x != nil {
		return x.PaidAmount
	}
	return nil
}

func (x *Order) GetRemainAmount() *money.Money {
	if x != nil {
		return x.RemainAmount
	}
	return nil
}

func (x *Order) GetRefundedAmount() *money.Money {
	if x != nil {
		return x.RefundedAmount
	}
	return nil
}

func (x *Order) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Order) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Order) GetLastUpdatedBy() string {
	if x != nil {
		return x.LastUpdatedBy
	}
	return ""
}

func (x *Order) GetLastUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedTime
	}
	return nil
}

func (x *Order) GetSalesDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.SalesDatetime
	}
	return nil
}

var File_moego_business_order_v1_order_proto protoreflect.FileDescriptor

var file_moego_business_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x08, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x74,
	0x69, 0x70, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0a, 0x74, 0x69, 0x70, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x31, 0x0a, 0x0a, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x09, 0x74, 0x61, 0x78, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3c, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0e,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c,
	0x0a, 0x10, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0e, 0x73, 0x75,
	0x62, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x11,
	0x74, 0x69, 0x70, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0f, 0x74, 0x69, 0x70,
	0x73, 0x42, 0x61, 0x73, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0a, 0x70, 0x61,
	0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3b, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0e,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3d, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x46, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0e,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x59, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x42, 0x80, 0x01, 0x0a, 0x1f, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_moego_business_order_v1_order_proto_goTypes = []any{
	(Order_Status)(0),             // 0: moego.business.order.v1.Order.Status
	(*Order)(nil),                 // 1: moego.business.order.v1.Order
	(*money.Money)(nil),           // 2: google.type.Money
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_moego_business_order_v1_order_proto_depIdxs = []int32{
	0,  // 0: moego.business.order.v1.Order.status:type_name -> moego.business.order.v1.Order.Status
	2,  // 1: moego.business.order.v1.Order.tips_amount:type_name -> google.type.Money
	2,  // 2: moego.business.order.v1.Order.tax_amount:type_name -> google.type.Money
	2,  // 3: moego.business.order.v1.Order.discount_amount:type_name -> google.type.Money
	2,  // 4: moego.business.order.v1.Order.extra_fee_amount:type_name -> google.type.Money
	2,  // 5: moego.business.order.v1.Order.sub_total_amount:type_name -> google.type.Money
	2,  // 6: moego.business.order.v1.Order.tips_based_amount:type_name -> google.type.Money
	2,  // 7: moego.business.order.v1.Order.total_amount:type_name -> google.type.Money
	2,  // 8: moego.business.order.v1.Order.paid_amount:type_name -> google.type.Money
	2,  // 9: moego.business.order.v1.Order.remain_amount:type_name -> google.type.Money
	2,  // 10: moego.business.order.v1.Order.refunded_amount:type_name -> google.type.Money
	3,  // 11: moego.business.order.v1.Order.created_time:type_name -> google.protobuf.Timestamp
	3,  // 12: moego.business.order.v1.Order.last_updated_time:type_name -> google.protobuf.Timestamp
	3,  // 13: moego.business.order.v1.Order.sales_datetime:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_moego_business_order_v1_order_proto_init() }
func file_moego_business_order_v1_order_proto_init() {
	if File_moego_business_order_v1_order_proto != nil {
		return
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
