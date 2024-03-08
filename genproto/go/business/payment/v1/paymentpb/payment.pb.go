// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: moego/business/payment/v1/payment.proto

package paymentpb

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

// Status
type Payment_Status int32

const (
	// PAYMENT_STATUS_UNSPECIFIED
	Payment_PAYMENT_STATUS_UNSPECIFIED Payment_Status = 0
	// UNPAID
	Payment_UNPAID Payment_Status = 1
	// PARTIAL_PAID
	Payment_PARTIAL_PAID Payment_Status = 2
	// FULL_PAID
	Payment_FULL_PAID Payment_Status = 3
)

// Enum value maps for Payment_Status.
var (
	Payment_Status_name = map[int32]string{
		0: "PAYMENT_STATUS_UNSPECIFIED",
		1: "UNPAID",
		2: "PARTIAL_PAID",
		3: "FULL_PAID",
	}
	Payment_Status_value = map[string]int32{
		"PAYMENT_STATUS_UNSPECIFIED": 0,
		"UNPAID":                     1,
		"PARTIAL_PAID":               2,
		"FULL_PAID":                  3,
	}
)

func (x Payment_Status) Enum() *Payment_Status {
	p := new(Payment_Status)
	*p = x
	return p
}

func (x Payment_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Payment_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_payment_v1_payment_proto_enumTypes[0].Descriptor()
}

func (Payment_Status) Type() protoreflect.EnumType {
	return &file_moego_business_payment_v1_payment_proto_enumTypes[0]
}

func (x Payment_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Payment_Status.Descriptor instead.
func (Payment_Status) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_payment_v1_payment_proto_rawDescGZIP(), []int{0, 0}
}

// Payment
type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_payment_v1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_payment_v1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_moego_business_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_moego_business_payment_v1_payment_proto protoreflect.FileDescriptor

var file_moego_business_payment_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x22, 0x70, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x55, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x50,
	0x41, 0x49, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c,
	0x5f, 0x50, 0x41, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4c, 0x4c, 0x5f,
	0x50, 0x41, 0x49, 0x44, 0x10, 0x03, 0x42, 0x8a, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_payment_v1_payment_proto_rawDescOnce sync.Once
	file_moego_business_payment_v1_payment_proto_rawDescData = file_moego_business_payment_v1_payment_proto_rawDesc
)

func file_moego_business_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_moego_business_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_moego_business_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_payment_v1_payment_proto_rawDescData)
	})
	return file_moego_business_payment_v1_payment_proto_rawDescData
}

var file_moego_business_payment_v1_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_moego_business_payment_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_moego_business_payment_v1_payment_proto_goTypes = []interface{}{
	(Payment_Status)(0), // 0: moego.business.payment.v1.Payment.Status
	(*Payment)(nil),     // 1: moego.business.payment.v1.Payment
}
var file_moego_business_payment_v1_payment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_moego_business_payment_v1_payment_proto_init() }
func file_moego_business_payment_v1_payment_proto_init() {
	if File_moego_business_payment_v1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_business_payment_v1_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
			RawDescriptor: file_moego_business_payment_v1_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_payment_v1_payment_proto_goTypes,
		DependencyIndexes: file_moego_business_payment_v1_payment_proto_depIdxs,
		EnumInfos:         file_moego_business_payment_v1_payment_proto_enumTypes,
		MessageInfos:      file_moego_business_payment_v1_payment_proto_msgTypes,
	}.Build()
	File_moego_business_payment_v1_payment_proto = out.File
	file_moego_business_payment_v1_payment_proto_rawDesc = nil
	file_moego_business_payment_v1_payment_proto_goTypes = nil
	file_moego_business_payment_v1_payment_proto_depIdxs = nil
}