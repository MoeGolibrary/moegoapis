// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: moego/business/business/v1/business_service.proto

package businesspb

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

// GetBusinessRequest request for GetBusiness
type GetBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of business
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBusinessRequest) Reset() {
	*x = GetBusinessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_business_v1_business_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessRequest) ProtoMessage() {}

func (x *GetBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_business_v1_business_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessRequest.ProtoReflect.Descriptor instead.
func (*GetBusinessRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_business_v1_business_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBusinessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ListBusinessRequest request for ListBusiness
type ListBusinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pagination information
	Pagination *commonpb.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// The company id
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *ListBusinessRequest) Reset() {
	*x = ListBusinessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_business_v1_business_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBusinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBusinessRequest) ProtoMessage() {}

func (x *ListBusinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_business_v1_business_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBusinessRequest.ProtoReflect.Descriptor instead.
func (*ListBusinessRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_business_v1_business_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListBusinessRequest) GetPagination() *commonpb.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListBusinessRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

// ListBusinessResponse response for ListBusiness
type ListBusinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The next page token
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The businesses
	Businesses []*Business `protobuf:"bytes,2,rep,name=businesses,proto3" json:"businesses,omitempty"`
}

func (x *ListBusinessResponse) Reset() {
	*x = ListBusinessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_business_v1_business_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBusinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBusinessResponse) ProtoMessage() {}

func (x *ListBusinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_business_v1_business_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBusinessResponse.ProtoReflect.Descriptor instead.
func (*ListBusinessResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_business_v1_business_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBusinessResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListBusinessResponse) GetBusinesses() []*Business {
	if x != nil {
		return x.Businesses
	}
	return nil
}

var File_moego_business_business_v1_business_service_proto protoreflect.FileDescriptor

var file_moego_business_business_v1_business_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x32, 0xa8, 0x02, 0x0a, 0x0f, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80,
	0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x2e,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x91, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x2f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a,
	0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x3a, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x96, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x70, 0x62, 0x3b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_business_v1_business_service_proto_rawDescOnce sync.Once
	file_moego_business_business_v1_business_service_proto_rawDescData = file_moego_business_business_v1_business_service_proto_rawDesc
)

func file_moego_business_business_v1_business_service_proto_rawDescGZIP() []byte {
	file_moego_business_business_v1_business_service_proto_rawDescOnce.Do(func() {
		file_moego_business_business_v1_business_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_business_v1_business_service_proto_rawDescData)
	})
	return file_moego_business_business_v1_business_service_proto_rawDescData
}

var file_moego_business_business_v1_business_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_moego_business_business_v1_business_service_proto_goTypes = []any{
	(*GetBusinessRequest)(nil),   // 0: moego.business.business.v1.GetBusinessRequest
	(*ListBusinessRequest)(nil),  // 1: moego.business.business.v1.ListBusinessRequest
	(*ListBusinessResponse)(nil), // 2: moego.business.business.v1.ListBusinessResponse
	(*commonpb.Pagination)(nil),  // 3: moego.common.v1.Pagination
	(*Business)(nil),             // 4: moego.business.business.v1.Business
}
var file_moego_business_business_v1_business_service_proto_depIdxs = []int32{
	3, // 0: moego.business.business.v1.ListBusinessRequest.pagination:type_name -> moego.common.v1.Pagination
	4, // 1: moego.business.business.v1.ListBusinessResponse.businesses:type_name -> moego.business.business.v1.Business
	0, // 2: moego.business.business.v1.BusinessService.GetBusiness:input_type -> moego.business.business.v1.GetBusinessRequest
	1, // 3: moego.business.business.v1.BusinessService.ListBusiness:input_type -> moego.business.business.v1.ListBusinessRequest
	4, // 4: moego.business.business.v1.BusinessService.GetBusiness:output_type -> moego.business.business.v1.Business
	2, // 5: moego.business.business.v1.BusinessService.ListBusiness:output_type -> moego.business.business.v1.ListBusinessResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_moego_business_business_v1_business_service_proto_init() }
func file_moego_business_business_v1_business_service_proto_init() {
	if File_moego_business_business_v1_business_service_proto != nil {
		return
	}
	file_moego_business_business_v1_business_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_moego_business_business_v1_business_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetBusinessRequest); i {
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
		file_moego_business_business_v1_business_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListBusinessRequest); i {
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
		file_moego_business_business_v1_business_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListBusinessResponse); i {
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
			RawDescriptor: file_moego_business_business_v1_business_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moego_business_business_v1_business_service_proto_goTypes,
		DependencyIndexes: file_moego_business_business_v1_business_service_proto_depIdxs,
		MessageInfos:      file_moego_business_business_v1_business_service_proto_msgTypes,
	}.Build()
	File_moego_business_business_v1_business_service_proto = out.File
	file_moego_business_business_v1_business_service_proto_rawDesc = nil
	file_moego_business_business_v1_business_service_proto_goTypes = nil
	file_moego_business_business_v1_business_service_proto_depIdxs = nil
}
