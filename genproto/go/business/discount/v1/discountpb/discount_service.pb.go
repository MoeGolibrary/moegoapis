// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: moego/business/discount/v1/discount_service.proto

package discountpb

import (
	commonpb "github.com/MoeGolibrary/moegoapis/genproto/go/common/v1/commonpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CreateDiscountRequest
type CreateDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// The discount to be created
	Discount *Discount `protobuf:"bytes,2,opt,name=discount,proto3" json:"discount,omitempty"`
}

func (x *CreateDiscountRequest) Reset() {
	*x = CreateDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDiscountRequest) ProtoMessage() {}

func (x *CreateDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDiscountRequest.ProtoReflect.Descriptor instead.
func (*CreateDiscountRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDiscountRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreateDiscountRequest) GetDiscount() *Discount {
	if x != nil {
		return x.Discount
	}
	return nil
}

// GetDiscountRequest
type GetDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Code      string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetDiscountRequest) Reset() {
	*x = GetDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiscountRequest) ProtoMessage() {}

func (x *GetDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiscountRequest.ProtoReflect.Descriptor instead.
func (*GetDiscountRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDiscountRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *GetDiscountRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// ListDiscountsRequest is the request message for DiscountAPI.ListDiscounts
type ListDiscountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pagination information
	Pagination *commonpb.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	CompanyId  string               `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *ListDiscountsRequest) Reset() {
	*x = ListDiscountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDiscountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDiscountsRequest) ProtoMessage() {}

func (x *ListDiscountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDiscountsRequest.ProtoReflect.Descriptor instead.
func (*ListDiscountsRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListDiscountsRequest) GetPagination() *commonpb.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListDiscountsRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

// ListDiscountsResponse is the response message for DiscountAPI.ListDiscounts
type ListDiscountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The next page token
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The Discounts
	Discounts []*Discount `protobuf:"bytes,2,rep,name=discounts,proto3" json:"discounts,omitempty"`
}

func (x *ListDiscountsResponse) Reset() {
	*x = ListDiscountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDiscountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDiscountsResponse) ProtoMessage() {}

func (x *ListDiscountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_discount_v1_discount_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDiscountsResponse.ProtoReflect.Descriptor instead.
func (*ListDiscountsResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_discount_v1_discount_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListDiscountsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListDiscountsResponse) GetDiscounts() []*Discount {
	if x != nil {
		return x.Discounts
	}
	return nil
}

var File_moego_business_discount_v1_discount_service_proto protoreflect.FileDescriptor

var file_moego_business_discount_v1_discount_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x29, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x51, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x7c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x42, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x32, 0xb1, 0x03, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a,
	0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x81, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x64, 0x65, 0x7d, 0x12, 0x93, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x30, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x96, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x14, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_discount_v1_discount_service_proto_rawDescOnce sync.Once
	file_moego_business_discount_v1_discount_service_proto_rawDescData = file_moego_business_discount_v1_discount_service_proto_rawDesc
)

func file_moego_business_discount_v1_discount_service_proto_rawDescGZIP() []byte {
	file_moego_business_discount_v1_discount_service_proto_rawDescOnce.Do(func() {
		file_moego_business_discount_v1_discount_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_discount_v1_discount_service_proto_rawDescData)
	})
	return file_moego_business_discount_v1_discount_service_proto_rawDescData
}

var file_moego_business_discount_v1_discount_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_moego_business_discount_v1_discount_service_proto_goTypes = []interface{}{
	(*CreateDiscountRequest)(nil), // 0: moego.business.discount.v1.CreateDiscountRequest
	(*GetDiscountRequest)(nil),    // 1: moego.business.discount.v1.GetDiscountRequest
	(*ListDiscountsRequest)(nil),  // 2: moego.business.discount.v1.ListDiscountsRequest
	(*ListDiscountsResponse)(nil), // 3: moego.business.discount.v1.ListDiscountsResponse
	(*Discount)(nil),              // 4: moego.business.discount.v1.Discount
	(*commonpb.Pagination)(nil),   // 5: moego.common.v1.Pagination
}
var file_moego_business_discount_v1_discount_service_proto_depIdxs = []int32{
	4, // 0: moego.business.discount.v1.CreateDiscountRequest.discount:type_name -> moego.business.discount.v1.Discount
	5, // 1: moego.business.discount.v1.ListDiscountsRequest.pagination:type_name -> moego.common.v1.Pagination
	4, // 2: moego.business.discount.v1.ListDiscountsResponse.discounts:type_name -> moego.business.discount.v1.Discount
	0, // 3: moego.business.discount.v1.DiscountService.CreateDiscount:input_type -> moego.business.discount.v1.CreateDiscountRequest
	1, // 4: moego.business.discount.v1.DiscountService.GetDiscount:input_type -> moego.business.discount.v1.GetDiscountRequest
	2, // 5: moego.business.discount.v1.DiscountService.ListDiscounts:input_type -> moego.business.discount.v1.ListDiscountsRequest
	4, // 6: moego.business.discount.v1.DiscountService.CreateDiscount:output_type -> moego.business.discount.v1.Discount
	4, // 7: moego.business.discount.v1.DiscountService.GetDiscount:output_type -> moego.business.discount.v1.Discount
	3, // 8: moego.business.discount.v1.DiscountService.ListDiscounts:output_type -> moego.business.discount.v1.ListDiscountsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_moego_business_discount_v1_discount_service_proto_init() }
func file_moego_business_discount_v1_discount_service_proto_init() {
	if File_moego_business_discount_v1_discount_service_proto != nil {
		return
	}
	file_moego_business_discount_v1_discount_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_moego_business_discount_v1_discount_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDiscountRequest); i {
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
		file_moego_business_discount_v1_discount_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiscountRequest); i {
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
		file_moego_business_discount_v1_discount_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDiscountsRequest); i {
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
		file_moego_business_discount_v1_discount_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDiscountsResponse); i {
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
			RawDescriptor: file_moego_business_discount_v1_discount_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moego_business_discount_v1_discount_service_proto_goTypes,
		DependencyIndexes: file_moego_business_discount_v1_discount_service_proto_depIdxs,
		MessageInfos:      file_moego_business_discount_v1_discount_service_proto_msgTypes,
	}.Build()
	File_moego_business_discount_v1_discount_service_proto = out.File
	file_moego_business_discount_v1_discount_service_proto_rawDesc = nil
	file_moego_business_discount_v1_discount_service_proto_goTypes = nil
	file_moego_business_discount_v1_discount_service_proto_depIdxs = nil
}
