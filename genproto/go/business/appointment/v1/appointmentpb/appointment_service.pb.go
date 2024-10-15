// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: moego/business/appointment/v1/appointment_service.proto

package appointmentpb

import (
	commonpb "github.com/MoeGolibrary/moegoapis/genproto/go/common/v1/commonpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	interval "google.golang.org/genproto/googleapis/type/interval"
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

// GetAppointmentRequest get appointment request
type GetAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BusinessId string `protobuf:"bytes,2,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
}

func (x *GetAppointmentRequest) Reset() {
	*x = GetAppointmentRequest{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppointmentRequest) ProtoMessage() {}

func (x *GetAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppointmentRequest.ProtoReflect.Descriptor instead.
func (*GetAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAppointmentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAppointmentRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

type ListAppointmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination  *commonpb.Pagination            `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	CompanyId   string                          `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	BusinessIds []string                        `protobuf:"bytes,3,rep,name=business_ids,json=businessIds,proto3" json:"business_ids,omitempty"`
	Filter      *ListAppointmentsRequest_Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListAppointmentsRequest) Reset() {
	*x = ListAppointmentsRequest{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAppointmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppointmentsRequest) ProtoMessage() {}

func (x *ListAppointmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppointmentsRequest.ProtoReflect.Descriptor instead.
func (*ListAppointmentsRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListAppointmentsRequest) GetPagination() *commonpb.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListAppointmentsRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ListAppointmentsRequest) GetBusinessIds() []string {
	if x != nil {
		return x.BusinessIds
	}
	return nil
}

func (x *ListAppointmentsRequest) GetFilter() *ListAppointmentsRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListAppointmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The next page token
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The Appointments
	Appointments []*Appointment `protobuf:"bytes,2,rep,name=appointments,proto3" json:"appointments,omitempty"`
}

func (x *ListAppointmentsResponse) Reset() {
	*x = ListAppointmentsResponse{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAppointmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppointmentsResponse) ProtoMessage() {}

func (x *ListAppointmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppointmentsResponse.ProtoReflect.Descriptor instead.
func (*ListAppointmentsResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListAppointmentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListAppointmentsResponse) GetAppointments() []*Appointment {
	if x != nil {
		return x.Appointments
	}
	return nil
}

type CreateAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account Info
	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// PetService info
	PetServices []*CreateAppointmentRequest_PetService `protobuf:"bytes,3,rep,name=pet_services,json=petServices,proto3" json:"pet_services,omitempty"`
}

func (x *CreateAppointmentRequest) Reset() {
	*x = CreateAppointmentRequest{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppointmentRequest) ProtoMessage() {}

func (x *CreateAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppointmentRequest.ProtoReflect.Descriptor instead.
func (*CreateAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAppointmentRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *CreateAppointmentRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateAppointmentRequest) GetPetServices() []*CreateAppointmentRequest_PetService {
	if x != nil {
		return x.PetServices
	}
	return nil
}

type ListAppointmentsRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime       *interval.Interval   `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime         *interval.Interval   `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	LastUpdatedTime *interval.Interval   `protobuf:"bytes,3,opt,name=last_updated_time,json=lastUpdatedTime,proto3" json:"last_updated_time,omitempty"`
	Statuses        []Appointment_Status `protobuf:"varint,4,rep,packed,name=statuses,proto3,enum=moego.business.appointment.v1.Appointment_Status" json:"statuses,omitempty"`
}

func (x *ListAppointmentsRequest_Filter) Reset() {
	*x = ListAppointmentsRequest_Filter{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAppointmentsRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppointmentsRequest_Filter) ProtoMessage() {}

func (x *ListAppointmentsRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppointmentsRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListAppointmentsRequest_Filter) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListAppointmentsRequest_Filter) GetStartTime() *interval.Interval {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ListAppointmentsRequest_Filter) GetEndTime() *interval.Interval {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *ListAppointmentsRequest_Filter) GetLastUpdatedTime() *interval.Interval {
	if x != nil {
		return x.LastUpdatedTime
	}
	return nil
}

func (x *ListAppointmentsRequest_Filter) GetStatuses() []Appointment_Status {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type CreateAppointmentRequest_PetService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId    string                              `protobuf:"bytes,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Services []*CreateAppointmentRequest_Service `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *CreateAppointmentRequest_PetService) Reset() {
	*x = CreateAppointmentRequest_PetService{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAppointmentRequest_PetService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppointmentRequest_PetService) ProtoMessage() {}

func (x *CreateAppointmentRequest_PetService) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppointmentRequest_PetService.ProtoReflect.Descriptor instead.
func (*CreateAppointmentRequest_PetService) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CreateAppointmentRequest_PetService) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *CreateAppointmentRequest_PetService) GetServices() []*CreateAppointmentRequest_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type CreateAppointmentRequest_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Duration *interval.Interval `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`                 // the service start time && end time
	StaffIds []string           `protobuf:"bytes,3,rep,name=staff_ids,json=staffIds,proto3" json:"staff_ids,omitempty"` // only support first staff
}

func (x *CreateAppointmentRequest_Service) Reset() {
	*x = CreateAppointmentRequest_Service{}
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAppointmentRequest_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppointmentRequest_Service) ProtoMessage() {}

func (x *CreateAppointmentRequest_Service) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppointmentRequest_Service.ProtoReflect.Descriptor instead.
func (*CreateAppointmentRequest_Service) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP(), []int{3, 1}
}

func (x *CreateAppointmentRequest_Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateAppointmentRequest_Service) GetDuration() *interval.Interval {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *CreateAppointmentRequest_Service) GetStaffIds() []string {
	if x != nil {
		return x.StaffIds
	}
	return nil
}

var File_moego_business_appointment_v1_appointment_service_proto protoreflect.FileDescriptor

var file_moego_business_appointment_v1_appointment_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0b,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x22, 0x97, 0x04, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x73, 0x12, 0x55, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6d,
	0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x96, 0x02, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x46, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x4e, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xca, 0x03, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x6a, 0x0a, 0x0c, 0x70, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x70, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x8a, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x60, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x69, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x73, 0x32, 0xe8,
	0x03, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xa5, 0x01, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x95, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0xa5, 0x01, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x17, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_appointment_v1_appointment_service_proto_rawDescOnce sync.Once
	file_moego_business_appointment_v1_appointment_service_proto_rawDescData = file_moego_business_appointment_v1_appointment_service_proto_rawDesc
)

func file_moego_business_appointment_v1_appointment_service_proto_rawDescGZIP() []byte {
	file_moego_business_appointment_v1_appointment_service_proto_rawDescOnce.Do(func() {
		file_moego_business_appointment_v1_appointment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_appointment_v1_appointment_service_proto_rawDescData)
	})
	return file_moego_business_appointment_v1_appointment_service_proto_rawDescData
}

var file_moego_business_appointment_v1_appointment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_moego_business_appointment_v1_appointment_service_proto_goTypes = []any{
	(*GetAppointmentRequest)(nil),               // 0: moego.business.appointment.v1.GetAppointmentRequest
	(*ListAppointmentsRequest)(nil),             // 1: moego.business.appointment.v1.ListAppointmentsRequest
	(*ListAppointmentsResponse)(nil),            // 2: moego.business.appointment.v1.ListAppointmentsResponse
	(*CreateAppointmentRequest)(nil),            // 3: moego.business.appointment.v1.CreateAppointmentRequest
	(*ListAppointmentsRequest_Filter)(nil),      // 4: moego.business.appointment.v1.ListAppointmentsRequest.Filter
	(*CreateAppointmentRequest_PetService)(nil), // 5: moego.business.appointment.v1.CreateAppointmentRequest.PetService
	(*CreateAppointmentRequest_Service)(nil),    // 6: moego.business.appointment.v1.CreateAppointmentRequest.Service
	(*commonpb.Pagination)(nil),                 // 7: moego.common.v1.Pagination
	(*Appointment)(nil),                         // 8: moego.business.appointment.v1.Appointment
	(*interval.Interval)(nil),                   // 9: google.type.Interval
	(Appointment_Status)(0),                     // 10: moego.business.appointment.v1.Appointment.Status
}
var file_moego_business_appointment_v1_appointment_service_proto_depIdxs = []int32{
	7,  // 0: moego.business.appointment.v1.ListAppointmentsRequest.pagination:type_name -> moego.common.v1.Pagination
	4,  // 1: moego.business.appointment.v1.ListAppointmentsRequest.filter:type_name -> moego.business.appointment.v1.ListAppointmentsRequest.Filter
	8,  // 2: moego.business.appointment.v1.ListAppointmentsResponse.appointments:type_name -> moego.business.appointment.v1.Appointment
	5,  // 3: moego.business.appointment.v1.CreateAppointmentRequest.pet_services:type_name -> moego.business.appointment.v1.CreateAppointmentRequest.PetService
	9,  // 4: moego.business.appointment.v1.ListAppointmentsRequest.Filter.start_time:type_name -> google.type.Interval
	9,  // 5: moego.business.appointment.v1.ListAppointmentsRequest.Filter.end_time:type_name -> google.type.Interval
	9,  // 6: moego.business.appointment.v1.ListAppointmentsRequest.Filter.last_updated_time:type_name -> google.type.Interval
	10, // 7: moego.business.appointment.v1.ListAppointmentsRequest.Filter.statuses:type_name -> moego.business.appointment.v1.Appointment.Status
	6,  // 8: moego.business.appointment.v1.CreateAppointmentRequest.PetService.services:type_name -> moego.business.appointment.v1.CreateAppointmentRequest.Service
	9,  // 9: moego.business.appointment.v1.CreateAppointmentRequest.Service.duration:type_name -> google.type.Interval
	0,  // 10: moego.business.appointment.v1.AppointmentService.GetAppointment:input_type -> moego.business.appointment.v1.GetAppointmentRequest
	1,  // 11: moego.business.appointment.v1.AppointmentService.ListAppointments:input_type -> moego.business.appointment.v1.ListAppointmentsRequest
	3,  // 12: moego.business.appointment.v1.AppointmentService.CreateAppointment:input_type -> moego.business.appointment.v1.CreateAppointmentRequest
	8,  // 13: moego.business.appointment.v1.AppointmentService.GetAppointment:output_type -> moego.business.appointment.v1.Appointment
	2,  // 14: moego.business.appointment.v1.AppointmentService.ListAppointments:output_type -> moego.business.appointment.v1.ListAppointmentsResponse
	8,  // 15: moego.business.appointment.v1.AppointmentService.CreateAppointment:output_type -> moego.business.appointment.v1.Appointment
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_moego_business_appointment_v1_appointment_service_proto_init() }
func file_moego_business_appointment_v1_appointment_service_proto_init() {
	if File_moego_business_appointment_v1_appointment_service_proto != nil {
		return
	}
	file_moego_business_appointment_v1_appointment_proto_init()
	file_moego_business_appointment_v1_pet_service_detail_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moego_business_appointment_v1_appointment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moego_business_appointment_v1_appointment_service_proto_goTypes,
		DependencyIndexes: file_moego_business_appointment_v1_appointment_service_proto_depIdxs,
		MessageInfos:      file_moego_business_appointment_v1_appointment_service_proto_msgTypes,
	}.Build()
	File_moego_business_appointment_v1_appointment_service_proto = out.File
	file_moego_business_appointment_v1_appointment_service_proto_rawDesc = nil
	file_moego_business_appointment_v1_appointment_service_proto_goTypes = nil
	file_moego_business_appointment_v1_appointment_service_proto_depIdxs = nil
}
