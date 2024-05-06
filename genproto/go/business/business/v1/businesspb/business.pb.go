// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: moego/business/business/v1/business.proto

package businesspb

import (
	commonpb "github.com/MoeGolibrary/moegoapis/genproto/go/common/v1/commonpb"
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

// Business information
type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar        string            `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Phone         string            `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	BusinessPhone string            `protobuf:"bytes,5,opt,name=business_phone,json=businessPhone,proto3" json:"business_phone,omitempty"`
	Email         string            `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Address       *commonpb.Address `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Website       string            `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Facebook      string            `protobuf:"bytes,9,opt,name=facebook,proto3" json:"facebook,omitempty"`
	Instagram     string            `protobuf:"bytes,10,opt,name=instagram,proto3" json:"instagram,omitempty"`
	Yelp          string            `protobuf:"bytes,11,opt,name=yelp,proto3" json:"yelp,omitempty"`
	Google        string            `protobuf:"bytes,12,opt,name=google,proto3" json:"google,omitempty"`
	CompanyId     string            `protobuf:"bytes,13,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_business_v1_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_business_v1_business_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_moego_business_business_v1_business_proto_rawDescGZIP(), []int{0}
}

func (x *Business) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Business) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Business) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Business) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Business) GetBusinessPhone() string {
	if x != nil {
		return x.BusinessPhone
	}
	return ""
}

func (x *Business) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Business) GetAddress() *commonpb.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Business) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Business) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *Business) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *Business) GetYelp() string {
	if x != nil {
		return x.Yelp
	}
	return ""
}

func (x *Business) GetGoogle() string {
	if x != nil {
		return x.Google
	}
	return ""
}

func (x *Business) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

var File_moego_business_business_v1_business_proto protoreflect.FileDescriptor

var file_moego_business_business_v1_business_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x6c, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x6c, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x42, 0x8f, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x70, 0x62, 0x3b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_business_v1_business_proto_rawDescOnce sync.Once
	file_moego_business_business_v1_business_proto_rawDescData = file_moego_business_business_v1_business_proto_rawDesc
)

func file_moego_business_business_v1_business_proto_rawDescGZIP() []byte {
	file_moego_business_business_v1_business_proto_rawDescOnce.Do(func() {
		file_moego_business_business_v1_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_business_v1_business_proto_rawDescData)
	})
	return file_moego_business_business_v1_business_proto_rawDescData
}

var file_moego_business_business_v1_business_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_moego_business_business_v1_business_proto_goTypes = []interface{}{
	(*Business)(nil),         // 0: moego.business.business.v1.Business
	(*commonpb.Address)(nil), // 1: moego.common.v1.Address
}
var file_moego_business_business_v1_business_proto_depIdxs = []int32{
	1, // 0: moego.business.business.v1.Business.address:type_name -> moego.common.v1.Address
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_moego_business_business_v1_business_proto_init() }
func file_moego_business_business_v1_business_proto_init() {
	if File_moego_business_business_v1_business_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_business_business_v1_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
			RawDescriptor: file_moego_business_business_v1_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_business_v1_business_proto_goTypes,
		DependencyIndexes: file_moego_business_business_v1_business_proto_depIdxs,
		MessageInfos:      file_moego_business_business_v1_business_proto_msgTypes,
	}.Build()
	File_moego_business_business_v1_business_proto = out.File
	file_moego_business_business_v1_business_proto_rawDesc = nil
	file_moego_business_business_v1_business_proto_goTypes = nil
	file_moego_business_business_v1_business_proto_depIdxs = nil
}
