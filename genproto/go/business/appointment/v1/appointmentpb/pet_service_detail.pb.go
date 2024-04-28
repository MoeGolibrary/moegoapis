// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: moego/business/appointment/v1/pet_service_detail.proto

package appointmentpb

import (
	customerpb "github.com/MoeGolibrary/moegoapis/genproto/go/business/customer/v1/customerpb"
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

// PetServiceDetail
type PetServiceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet            *customerpb.Pet  `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	ServiceDetails []*ServiceDetail `protobuf:"bytes,2,rep,name=service_details,json=serviceDetails,proto3" json:"service_details,omitempty"`
}

func (x *PetServiceDetail) Reset() {
	*x = PetServiceDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetServiceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetServiceDetail) ProtoMessage() {}

func (x *PetServiceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetServiceDetail.ProtoReflect.Descriptor instead.
func (*PetServiceDetail) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_pet_service_detail_proto_rawDescGZIP(), []int{0}
}

func (x *PetServiceDetail) GetPet() *customerpb.Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

func (x *PetServiceDetail) GetServiceDetails() []*ServiceDetail {
	if x != nil {
		return x.ServiceDetails
	}
	return nil
}

// ServiceDetail
type ServiceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price    *money.Money       `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Duration *interval.Interval `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	StaffIds []string           `protobuf:"bytes,5,rep,name=staff_ids,json=staffIds,proto3" json:"staff_ids,omitempty"`
}

func (x *ServiceDetail) Reset() {
	*x = ServiceDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDetail) ProtoMessage() {}

func (x *ServiceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDetail.ProtoReflect.Descriptor instead.
func (*ServiceDetail) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_pet_service_detail_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceDetail) GetPrice() *money.Money {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *ServiceDetail) GetDuration() *interval.Interval {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *ServiceDetail) GetStaffIds() []string {
	if x != nil {
		return x.StaffIds
	}
	return nil
}

var File_moego_business_appointment_v1_pet_service_detail_proto protoreflect.FileDescriptor

var file_moego_business_appointment_v1_pet_service_detail_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x12, 0x55, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64,
	0x73, 0x42, 0xa3, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x50, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_appointment_v1_pet_service_detail_proto_rawDescOnce sync.Once
	file_moego_business_appointment_v1_pet_service_detail_proto_rawDescData = file_moego_business_appointment_v1_pet_service_detail_proto_rawDesc
)

func file_moego_business_appointment_v1_pet_service_detail_proto_rawDescGZIP() []byte {
	file_moego_business_appointment_v1_pet_service_detail_proto_rawDescOnce.Do(func() {
		file_moego_business_appointment_v1_pet_service_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_appointment_v1_pet_service_detail_proto_rawDescData)
	})
	return file_moego_business_appointment_v1_pet_service_detail_proto_rawDescData
}

var file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_moego_business_appointment_v1_pet_service_detail_proto_goTypes = []interface{}{
	(*PetServiceDetail)(nil),  // 0: moego.business.appointment.v1.PetServiceDetail
	(*ServiceDetail)(nil),     // 1: moego.business.appointment.v1.ServiceDetail
	(*customerpb.Pet)(nil),    // 2: moego.business.customer.v1.Pet
	(*money.Money)(nil),       // 3: google.type.Money
	(*interval.Interval)(nil), // 4: google.type.Interval
}
var file_moego_business_appointment_v1_pet_service_detail_proto_depIdxs = []int32{
	2, // 0: moego.business.appointment.v1.PetServiceDetail.pet:type_name -> moego.business.customer.v1.Pet
	1, // 1: moego.business.appointment.v1.PetServiceDetail.service_details:type_name -> moego.business.appointment.v1.ServiceDetail
	3, // 2: moego.business.appointment.v1.ServiceDetail.price:type_name -> google.type.Money
	4, // 3: moego.business.appointment.v1.ServiceDetail.duration:type_name -> google.type.Interval
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_moego_business_appointment_v1_pet_service_detail_proto_init() }
func file_moego_business_appointment_v1_pet_service_detail_proto_init() {
	if File_moego_business_appointment_v1_pet_service_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetServiceDetail); i {
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
		file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDetail); i {
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
			RawDescriptor: file_moego_business_appointment_v1_pet_service_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_appointment_v1_pet_service_detail_proto_goTypes,
		DependencyIndexes: file_moego_business_appointment_v1_pet_service_detail_proto_depIdxs,
		MessageInfos:      file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes,
	}.Build()
	File_moego_business_appointment_v1_pet_service_detail_proto = out.File
	file_moego_business_appointment_v1_pet_service_detail_proto_rawDesc = nil
	file_moego_business_appointment_v1_pet_service_detail_proto_goTypes = nil
	file_moego_business_appointment_v1_pet_service_detail_proto_depIdxs = nil
}
