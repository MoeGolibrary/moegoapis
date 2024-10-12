// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
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

// ServiceItemType
type ServiceDetail_ServiceItemType int32

const (
	ServiceDetail_SERVICE_ITEM_TYPE_UNSPECIFIED ServiceDetail_ServiceItemType = 0
	ServiceDetail_GROOMING                      ServiceDetail_ServiceItemType = 1
	ServiceDetail_BOARDING                      ServiceDetail_ServiceItemType = 2
	ServiceDetail_DAYCARE                       ServiceDetail_ServiceItemType = 3
	ServiceDetail_EVALUATION                    ServiceDetail_ServiceItemType = 4
)

// Enum value maps for ServiceDetail_ServiceItemType.
var (
	ServiceDetail_ServiceItemType_name = map[int32]string{
		0: "SERVICE_ITEM_TYPE_UNSPECIFIED",
		1: "GROOMING",
		2: "BOARDING",
		3: "DAYCARE",
		4: "EVALUATION",
	}
	ServiceDetail_ServiceItemType_value = map[string]int32{
		"SERVICE_ITEM_TYPE_UNSPECIFIED": 0,
		"GROOMING":                      1,
		"BOARDING":                      2,
		"DAYCARE":                       3,
		"EVALUATION":                    4,
	}
)

func (x ServiceDetail_ServiceItemType) Enum() *ServiceDetail_ServiceItemType {
	p := new(ServiceDetail_ServiceItemType)
	*p = x
	return p
}

func (x ServiceDetail_ServiceItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceDetail_ServiceItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_appointment_v1_pet_service_detail_proto_enumTypes[0].Descriptor()
}

func (ServiceDetail_ServiceItemType) Type() protoreflect.EnumType {
	return &file_moego_business_appointment_v1_pet_service_detail_proto_enumTypes[0]
}

func (x ServiceDetail_ServiceItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceDetail_ServiceItemType.Descriptor instead.
func (ServiceDetail_ServiceItemType) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_pet_service_detail_proto_rawDescGZIP(), []int{1, 0}
}

// ServiceType
type ServiceDetail_ServiceType int32

const (
	ServiceDetail_SERVICE_TYPE_UNSPECIFIED ServiceDetail_ServiceType = 0
	ServiceDetail_SERVICE                  ServiceDetail_ServiceType = 1
	ServiceDetail_ADDON                    ServiceDetail_ServiceType = 2
)

// Enum value maps for ServiceDetail_ServiceType.
var (
	ServiceDetail_ServiceType_name = map[int32]string{
		0: "SERVICE_TYPE_UNSPECIFIED",
		1: "SERVICE",
		2: "ADDON",
	}
	ServiceDetail_ServiceType_value = map[string]int32{
		"SERVICE_TYPE_UNSPECIFIED": 0,
		"SERVICE":                  1,
		"ADDON":                    2,
	}
)

func (x ServiceDetail_ServiceType) Enum() *ServiceDetail_ServiceType {
	p := new(ServiceDetail_ServiceType)
	*p = x
	return p
}

func (x ServiceDetail_ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceDetail_ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_appointment_v1_pet_service_detail_proto_enumTypes[1].Descriptor()
}

func (ServiceDetail_ServiceType) Type() protoreflect.EnumType {
	return &file_moego_business_appointment_v1_pet_service_detail_proto_enumTypes[1]
}

func (x ServiceDetail_ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceDetail_ServiceType.Descriptor instead.
func (ServiceDetail_ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_pet_service_detail_proto_rawDescGZIP(), []int{1, 1}
}

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
	mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PetServiceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetServiceDetail) ProtoMessage() {}

func (x *PetServiceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[0]
	if x != nil {
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

	Id              string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServiceItemType ServiceDetail_ServiceItemType `protobuf:"varint,3,opt,name=service_item_type,json=serviceItemType,proto3,enum=moego.business.appointment.v1.ServiceDetail_ServiceItemType" json:"service_item_type,omitempty"`
	Category        string                        `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Price           *money.Money                  `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Duration        *interval.Interval            `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"` // the service start time && end time
	StaffIds        []string                      `protobuf:"bytes,7,rep,name=staff_ids,json=staffIds,proto3" json:"staff_ids,omitempty"`
	ServiceType     ServiceDetail_ServiceType     `protobuf:"varint,8,opt,name=service_type,json=serviceType,proto3,enum=moego.business.appointment.v1.ServiceDetail_ServiceType" json:"service_type,omitempty"`
	ServiceTime     int32                         `protobuf:"varint,9,opt,name=service_time,json=serviceTime,proto3" json:"service_time,omitempty"` // the service provision time
}

func (x *ServiceDetail) Reset() {
	*x = ServiceDetail{}
	mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDetail) ProtoMessage() {}

func (x *ServiceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes[1]
	if x != nil {
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

func (x *ServiceDetail) GetServiceItemType() ServiceDetail_ServiceItemType {
	if x != nil {
		return x.ServiceItemType
	}
	return ServiceDetail_SERVICE_ITEM_TYPE_UNSPECIFIED
}

func (x *ServiceDetail) GetCategory() string {
	if x != nil {
		return x.Category
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

func (x *ServiceDetail) GetServiceType() ServiceDetail_ServiceType {
	if x != nil {
		return x.ServiceType
	}
	return ServiceDetail_SERVICE_TYPE_UNSPECIFIED
}

func (x *ServiceDetail) GetServiceTime() int32 {
	if x != nil {
		return x.ServiceTime
	}
	return 0
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
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
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
	0x73, 0x22, 0xe7, 0x04, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x73, 0x12, 0x5b, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x47, 0x52, 0x4f, 0x4f, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42,
	0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x59,
	0x43, 0x41, 0x52, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x22, 0x43, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x44, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0xa3, 0x01, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_moego_business_appointment_v1_pet_service_detail_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_moego_business_appointment_v1_pet_service_detail_proto_goTypes = []any{
	(ServiceDetail_ServiceItemType)(0), // 0: moego.business.appointment.v1.ServiceDetail.ServiceItemType
	(ServiceDetail_ServiceType)(0),     // 1: moego.business.appointment.v1.ServiceDetail.ServiceType
	(*PetServiceDetail)(nil),           // 2: moego.business.appointment.v1.PetServiceDetail
	(*ServiceDetail)(nil),              // 3: moego.business.appointment.v1.ServiceDetail
	(*customerpb.Pet)(nil),             // 4: moego.business.customer.v1.Pet
	(*money.Money)(nil),                // 5: google.type.Money
	(*interval.Interval)(nil),          // 6: google.type.Interval
}
var file_moego_business_appointment_v1_pet_service_detail_proto_depIdxs = []int32{
	4, // 0: moego.business.appointment.v1.PetServiceDetail.pet:type_name -> moego.business.customer.v1.Pet
	3, // 1: moego.business.appointment.v1.PetServiceDetail.service_details:type_name -> moego.business.appointment.v1.ServiceDetail
	0, // 2: moego.business.appointment.v1.ServiceDetail.service_item_type:type_name -> moego.business.appointment.v1.ServiceDetail.ServiceItemType
	5, // 3: moego.business.appointment.v1.ServiceDetail.price:type_name -> google.type.Money
	6, // 4: moego.business.appointment.v1.ServiceDetail.duration:type_name -> google.type.Interval
	1, // 5: moego.business.appointment.v1.ServiceDetail.service_type:type_name -> moego.business.appointment.v1.ServiceDetail.ServiceType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_moego_business_appointment_v1_pet_service_detail_proto_init() }
func file_moego_business_appointment_v1_pet_service_detail_proto_init() {
	if File_moego_business_appointment_v1_pet_service_detail_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moego_business_appointment_v1_pet_service_detail_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_appointment_v1_pet_service_detail_proto_goTypes,
		DependencyIndexes: file_moego_business_appointment_v1_pet_service_detail_proto_depIdxs,
		EnumInfos:         file_moego_business_appointment_v1_pet_service_detail_proto_enumTypes,
		MessageInfos:      file_moego_business_appointment_v1_pet_service_detail_proto_msgTypes,
	}.Build()
	File_moego_business_appointment_v1_pet_service_detail_proto = out.File
	file_moego_business_appointment_v1_pet_service_detail_proto_rawDesc = nil
	file_moego_business_appointment_v1_pet_service_detail_proto_goTypes = nil
	file_moego_business_appointment_v1_pet_service_detail_proto_depIdxs = nil
}
