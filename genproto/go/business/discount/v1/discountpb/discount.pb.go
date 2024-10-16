// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: moego/business/discount/v1/discount.proto

package discountpb

import (
	interval "google.golang.org/genproto/googleapis/type/interval"
	money "google.golang.org/genproto/googleapis/type/money"
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

// Discount
type Discount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to Value:
	//
	//	*Discount_Amount
	//	*Discount_Percentage
	Value       isDiscount_Value    `protobuf_oneof:"value"`
	ValidPeriod *interval.Interval  `protobuf:"bytes,5,opt,name=valid_period,json=validPeriod,proto3" json:"valid_period,omitempty"`
	Limitation  *DiscountLimitation `protobuf:"bytes,6,opt,name=limitation,proto3" json:"limitation,omitempty"`
	Settings    *DiscountSettings   `protobuf:"bytes,7,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *Discount) Reset() {
	*x = Discount{}
	mi := &file_moego_business_discount_v1_discount_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discount) ProtoMessage() {}

func (x *Discount) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discount.ProtoReflect.Descriptor instead.
func (*Discount) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_proto_rawDescGZIP(), []int{0}
}

func (x *Discount) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Discount) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *Discount) GetValue() isDiscount_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Discount) GetAmount() *money.Money {
	if x, ok := x.GetValue().(*Discount_Amount); ok {
		return x.Amount
	}
	return nil
}

func (x *Discount) GetPercentage() uint32 {
	if x, ok := x.GetValue().(*Discount_Percentage); ok {
		return x.Percentage
	}
	return 0
}

func (x *Discount) GetValidPeriod() *interval.Interval {
	if x != nil {
		return x.ValidPeriod
	}
	return nil
}

func (x *Discount) GetLimitation() *DiscountLimitation {
	if x != nil {
		return x.Limitation
	}
	return nil
}

func (x *Discount) GetSettings() *DiscountSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type isDiscount_Value interface {
	isDiscount_Value()
}

type Discount_Amount struct {
	Amount *money.Money `protobuf:"bytes,3,opt,name=amount,proto3,oneof"`
}

type Discount_Percentage struct {
	Percentage uint32 `protobuf:"varint,4,opt,name=percentage,proto3,oneof"`
}

func (*Discount_Amount) isDiscount_Value() {}

func (*Discount_Percentage) isDiscount_Value() {}

// DiscountLimitation
type DiscountLimitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxRedeemTimes        uint32   `protobuf:"varint,1,opt,name=max_redeem_times,json=maxRedeemTimes,proto3" json:"max_redeem_times,omitempty"`
	BusinessIds           []string `protobuf:"bytes,2,rep,name=business_ids,json=businessIds,proto3" json:"business_ids,omitempty"`
	RedeemOncePerCustomer bool     `protobuf:"varint,3,opt,name=redeem_once_per_customer,json=redeemOncePerCustomer,proto3" json:"redeem_once_per_customer,omitempty"`
	CustomerIds           []string `protobuf:"bytes,4,rep,name=customer_ids,json=customerIds,proto3" json:"customer_ids,omitempty"`
}

func (x *DiscountLimitation) Reset() {
	*x = DiscountLimitation{}
	mi := &file_moego_business_discount_v1_discount_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscountLimitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountLimitation) ProtoMessage() {}

func (x *DiscountLimitation) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountLimitation.ProtoReflect.Descriptor instead.
func (*DiscountLimitation) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_proto_rawDescGZIP(), []int{1}
}

func (x *DiscountLimitation) GetMaxRedeemTimes() uint32 {
	if x != nil {
		return x.MaxRedeemTimes
	}
	return 0
}

func (x *DiscountLimitation) GetBusinessIds() []string {
	if x != nil {
		return x.BusinessIds
	}
	return nil
}

func (x *DiscountLimitation) GetRedeemOncePerCustomer() bool {
	if x != nil {
		return x.RedeemOncePerCustomer
	}
	return false
}

func (x *DiscountLimitation) GetCustomerIds() []string {
	if x != nil {
		return x.CustomerIds
	}
	return nil
}

// DiscountSettings
type DiscountSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoApplyOnEligibleAppointment bool `protobuf:"varint,1,opt,name=auto_apply_on_eligible_appointment,json=autoApplyOnEligibleAppointment,proto3" json:"auto_apply_on_eligible_appointment,omitempty"`
	AllowForOnlineBooking          bool `protobuf:"varint,2,opt,name=allow_for_online_booking,json=allowForOnlineBooking,proto3" json:"allow_for_online_booking,omitempty"`
}

func (x *DiscountSettings) Reset() {
	*x = DiscountSettings{}
	mi := &file_moego_business_discount_v1_discount_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscountSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountSettings) ProtoMessage() {}

func (x *DiscountSettings) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountSettings.ProtoReflect.Descriptor instead.
func (*DiscountSettings) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_proto_rawDescGZIP(), []int{2}
}

func (x *DiscountSettings) GetAutoApplyOnEligibleAppointment() bool {
	if x != nil {
		return x.AutoApplyOnEligibleAppointment
	}
	return false
}

func (x *DiscountSettings) GetAllowForOnlineBooking() bool {
	if x != nil {
		return x.AllowForOnlineBooking
	}
	return false
}

var File_moego_business_discount_v1_discount_proto protoreflect.FileDescriptor

var file_moego_business_discount_v1_discount_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x02, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x38, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x0b, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x4e, 0x0a, 0x0a, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xbd, 0x01, 0x0a,
	0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x64, 0x65, 0x65,
	0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d,
	0x61, 0x78, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x73,
	0x12, 0x37, 0x0a, 0x18, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x5f, 0x6f, 0x6e, 0x63, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x4f, 0x6e, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x97, 0x01, 0x0a,
	0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x4a, 0x0a, 0x22, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x6f, 0x6e, 0x5f, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x61,
	0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62,
	0x6c, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a,
	0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x6f, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x8f, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_discount_v1_discount_proto_rawDescOnce sync.Once
	file_moego_business_discount_v1_discount_proto_rawDescData = file_moego_business_discount_v1_discount_proto_rawDesc
)

func file_moego_business_discount_v1_discount_proto_rawDescGZIP() []byte {
	file_moego_business_discount_v1_discount_proto_rawDescOnce.Do(func() {
		file_moego_business_discount_v1_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_discount_v1_discount_proto_rawDescData)
	})
	return file_moego_business_discount_v1_discount_proto_rawDescData
}

var file_moego_business_discount_v1_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_moego_business_discount_v1_discount_proto_goTypes = []any{
	(*Discount)(nil),           // 0: moego.business.discount.v1.Discount
	(*DiscountLimitation)(nil), // 1: moego.business.discount.v1.DiscountLimitation
	(*DiscountSettings)(nil),   // 2: moego.business.discount.v1.DiscountSettings
	(*money.Money)(nil),        // 3: google.type.Money
	(*interval.Interval)(nil),  // 4: google.type.Interval
}
var file_moego_business_discount_v1_discount_proto_depIdxs = []int32{
	3, // 0: moego.business.discount.v1.Discount.amount:type_name -> google.type.Money
	4, // 1: moego.business.discount.v1.Discount.valid_period:type_name -> google.type.Interval
	1, // 2: moego.business.discount.v1.Discount.limitation:type_name -> moego.business.discount.v1.DiscountLimitation
	2, // 3: moego.business.discount.v1.Discount.settings:type_name -> moego.business.discount.v1.DiscountSettings
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_moego_business_discount_v1_discount_proto_init() }
func file_moego_business_discount_v1_discount_proto_init() {
	if File_moego_business_discount_v1_discount_proto != nil {
		return
	}
	file_moego_business_discount_v1_discount_proto_msgTypes[0].OneofWrappers = []any{
		(*Discount_Amount)(nil),
		(*Discount_Percentage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moego_business_discount_v1_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_discount_v1_discount_proto_goTypes,
		DependencyIndexes: file_moego_business_discount_v1_discount_proto_depIdxs,
		MessageInfos:      file_moego_business_discount_v1_discount_proto_msgTypes,
	}.Build()
	File_moego_business_discount_v1_discount_proto = out.File
	file_moego_business_discount_v1_discount_proto_rawDesc = nil
	file_moego_business_discount_v1_discount_proto_goTypes = nil
	file_moego_business_discount_v1_discount_proto_depIdxs = nil
}
