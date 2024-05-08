// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: moego/business/event/v1/health_check.proto

package eventpb

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

// HealthCheck is the health check event payload
type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validation is a simple string that the webhook will echo back to the client
	Validation string `protobuf:"bytes,1,opt,name=validation,proto3" json:"validation,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_event_v1_health_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_event_v1_health_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_moego_business_event_v1_health_check_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheck) GetValidation() string {
	if x != nil {
		return x.Validation
	}
	return ""
}

var File_moego_business_event_v1_health_check_proto protoreflect.FileDescriptor

var file_moego_business_event_v1_health_check_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x2d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x86, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_event_v1_health_check_proto_rawDescOnce sync.Once
	file_moego_business_event_v1_health_check_proto_rawDescData = file_moego_business_event_v1_health_check_proto_rawDesc
)

func file_moego_business_event_v1_health_check_proto_rawDescGZIP() []byte {
	file_moego_business_event_v1_health_check_proto_rawDescOnce.Do(func() {
		file_moego_business_event_v1_health_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_event_v1_health_check_proto_rawDescData)
	})
	return file_moego_business_event_v1_health_check_proto_rawDescData
}

var file_moego_business_event_v1_health_check_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_moego_business_event_v1_health_check_proto_goTypes = []interface{}{
	(*HealthCheck)(nil), // 0: moego.business.event.v1.HealthCheck
}
var file_moego_business_event_v1_health_check_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_moego_business_event_v1_health_check_proto_init() }
func file_moego_business_event_v1_health_check_proto_init() {
	if File_moego_business_event_v1_health_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_business_event_v1_health_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
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
			RawDescriptor: file_moego_business_event_v1_health_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_event_v1_health_check_proto_goTypes,
		DependencyIndexes: file_moego_business_event_v1_health_check_proto_depIdxs,
		MessageInfos:      file_moego_business_event_v1_health_check_proto_msgTypes,
	}.Build()
	File_moego_business_event_v1_health_check_proto = out.File
	file_moego_business_event_v1_health_check_proto_rawDesc = nil
	file_moego_business_event_v1_health_check_proto_goTypes = nil
	file_moego_business_event_v1_health_check_proto_depIdxs = nil
}
