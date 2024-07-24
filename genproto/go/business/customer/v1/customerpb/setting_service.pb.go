// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: moego/business/setting/setting_service.proto

package customerpb

import (
	_ "github.com/MoeGolibrary/moegoapis/genproto/go/common/v1/commonpb"
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

type ListPetCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *ListPetCodeRequest) Reset() {
	*x = ListPetCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_setting_setting_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetCodeRequest) ProtoMessage() {}

func (x *ListPetCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_setting_setting_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetCodeRequest.ProtoReflect.Descriptor instead.
func (*ListPetCodeRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_setting_setting_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListPetCodeRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

type ListPetCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []*Pet_Code `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *ListPetCodeResponse) Reset() {
	*x = ListPetCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_setting_setting_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetCodeResponse) ProtoMessage() {}

func (x *ListPetCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_setting_setting_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetCodeResponse.ProtoReflect.Descriptor instead.
func (*ListPetCodeResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_setting_setting_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPetCodeResponse) GetCodes() []*Pet_Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_moego_business_setting_setting_service_proto protoreflect.FileDescriptor

var file_moego_business_setting_setting_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x38, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0xbf,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xac, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x2e, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x3a, 0x01, 0x2a, 0x22, 0x30,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x96, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47,
	0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_moego_business_setting_setting_service_proto_rawDescOnce sync.Once
	file_moego_business_setting_setting_service_proto_rawDescData = file_moego_business_setting_setting_service_proto_rawDesc
)

func file_moego_business_setting_setting_service_proto_rawDescGZIP() []byte {
	file_moego_business_setting_setting_service_proto_rawDescOnce.Do(func() {
		file_moego_business_setting_setting_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_setting_setting_service_proto_rawDescData)
	})
	return file_moego_business_setting_setting_service_proto_rawDescData
}

var file_moego_business_setting_setting_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_moego_business_setting_setting_service_proto_goTypes = []any{
	(*ListPetCodeRequest)(nil),  // 0: moego.business.customer.v1.ListPetCodeRequest
	(*ListPetCodeResponse)(nil), // 1: moego.business.customer.v1.ListPetCodeResponse
	(*Pet_Code)(nil),            // 2: moego.business.customer.v1.Pet.Code
}
var file_moego_business_setting_setting_service_proto_depIdxs = []int32{
	2, // 0: moego.business.customer.v1.ListPetCodeResponse.codes:type_name -> moego.business.customer.v1.Pet.Code
	0, // 1: moego.business.customer.v1.SettingService.ListPetCodes:input_type -> moego.business.customer.v1.ListPetCodeRequest
	1, // 2: moego.business.customer.v1.SettingService.ListPetCodes:output_type -> moego.business.customer.v1.ListPetCodeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_moego_business_setting_setting_service_proto_init() }
func file_moego_business_setting_setting_service_proto_init() {
	if File_moego_business_setting_setting_service_proto != nil {
		return
	}
	file_moego_business_customer_v1_pet_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_moego_business_setting_setting_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListPetCodeRequest); i {
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
		file_moego_business_setting_setting_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListPetCodeResponse); i {
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
			RawDescriptor: file_moego_business_setting_setting_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moego_business_setting_setting_service_proto_goTypes,
		DependencyIndexes: file_moego_business_setting_setting_service_proto_depIdxs,
		MessageInfos:      file_moego_business_setting_setting_service_proto_msgTypes,
	}.Build()
	File_moego_business_setting_setting_service_proto = out.File
	file_moego_business_setting_setting_service_proto_rawDesc = nil
	file_moego_business_setting_setting_service_proto_goTypes = nil
	file_moego_business_setting_setting_service_proto_depIdxs = nil
}
