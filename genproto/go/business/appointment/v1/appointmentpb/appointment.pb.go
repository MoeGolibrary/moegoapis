// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: moego/business/appointment/v1/appointment.proto

package appointmentpb

import (
	paymentpb "github.com/MoeGolibrary/moegoapis/genproto/go/business/payment/v1/paymentpb"
	commonpb "github.com/MoeGolibrary/moegoapis/genproto/go/common/v1/commonpb"
	interval "google.golang.org/genproto/googleapis/type/interval"
	money "google.golang.org/genproto/googleapis/type/money"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
type Appointment_Status int32

const (
	Appointment_STATUS_UNSPECIFIED Appointment_Status = 0
	Appointment_UNCONFIRMED        Appointment_Status = 1
	Appointment_CONFIRMED          Appointment_Status = 2
	Appointment_CHECKED_IN         Appointment_Status = 3
	Appointment_READY              Appointment_Status = 4
	Appointment_FINISHED           Appointment_Status = 5
	Appointment_CANCELED           Appointment_Status = 6
)

// Enum value maps for Appointment_Status.
var (
	Appointment_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "UNCONFIRMED",
		2: "CONFIRMED",
		3: "CHECKED_IN",
		4: "READY",
		5: "FINISHED",
		6: "CANCELED",
	}
	Appointment_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"UNCONFIRMED":        1,
		"CONFIRMED":          2,
		"CHECKED_IN":         3,
		"READY":              4,
		"FINISHED":           5,
		"CANCELED":           6,
	}
)

func (x Appointment_Status) Enum() *Appointment_Status {
	p := new(Appointment_Status)
	*p = x
	return p
}

func (x Appointment_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Appointment_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_appointment_v1_appointment_proto_enumTypes[0].Descriptor()
}

func (Appointment_Status) Type() protoreflect.EnumType {
	return &file_moego_business_appointment_v1_appointment_proto_enumTypes[0]
}

func (x Appointment_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Appointment_Status.Descriptor instead.
func (Appointment_Status) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_proto_rawDescGZIP(), []int{0, 0}
}

// Appointment
type Appointment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the appointment
	Id                string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BusinessId        string                   `protobuf:"bytes,2,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	CustomerId        string                   `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Address           *commonpb.Address        `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Duration          *interval.Interval       `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	PetServiceDetails []*PetServiceDetail      `protobuf:"bytes,6,rep,name=pet_service_details,json=petServiceDetails,proto3" json:"pet_service_details,omitempty"`
	Status            Appointment_Status       `protobuf:"varint,7,opt,name=status,proto3,enum=moego.business.appointment.v1.Appointment_Status" json:"status,omitempty"`
	TicketComment     string                   `protobuf:"bytes,8,opt,name=ticket_comment,json=ticketComment,proto3" json:"ticket_comment,omitempty"`
	ColorCode         string                   `protobuf:"bytes,9,opt,name=color_code,json=colorCode,proto3" json:"color_code,omitempty"`
	InvoiceId         string                   `protobuf:"bytes,10,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	TotalAmount       *money.Money             `protobuf:"bytes,11,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	PaidAmount        *money.Money             `protobuf:"bytes,12,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	RefundAmount      *money.Money             `protobuf:"bytes,13,opt,name=refund_amount,json=refundAmount,proto3" json:"refund_amount,omitempty"`
	PaymentStatus     paymentpb.Payment_Status `protobuf:"varint,14,opt,name=payment_status,json=paymentStatus,proto3,enum=moego.business.payment.v1.Payment_Status" json:"payment_status,omitempty"`
	CreatedBy         string                   `protobuf:"bytes,15,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedTime       *timestamppb.Timestamp   `protobuf:"bytes,16,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	LastUpdatedBy     string                   `protobuf:"bytes,17,opt,name=last_updated_by,json=lastUpdatedBy,proto3" json:"last_updated_by,omitempty"`
	LastUpdatedTime   *timestamppb.Timestamp   `protobuf:"bytes,18,opt,name=last_updated_time,json=lastUpdatedTime,proto3" json:"last_updated_time,omitempty"`
}

func (x *Appointment) Reset() {
	*x = Appointment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_appointment_v1_appointment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appointment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appointment) ProtoMessage() {}

func (x *Appointment) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_appointment_v1_appointment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appointment.ProtoReflect.Descriptor instead.
func (*Appointment) Descriptor() ([]byte, []int) {
	return file_moego_business_appointment_v1_appointment_proto_rawDescGZIP(), []int{0}
}

func (x *Appointment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Appointment) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Appointment) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Appointment) GetAddress() *commonpb.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Appointment) GetDuration() *interval.Interval {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Appointment) GetPetServiceDetails() []*PetServiceDetail {
	if x != nil {
		return x.PetServiceDetails
	}
	return nil
}

func (x *Appointment) GetStatus() Appointment_Status {
	if x != nil {
		return x.Status
	}
	return Appointment_STATUS_UNSPECIFIED
}

func (x *Appointment) GetTicketComment() string {
	if x != nil {
		return x.TicketComment
	}
	return ""
}

func (x *Appointment) GetColorCode() string {
	if x != nil {
		return x.ColorCode
	}
	return ""
}

func (x *Appointment) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *Appointment) GetTotalAmount() *money.Money {
	if x != nil {
		return x.TotalAmount
	}
	return nil
}

func (x *Appointment) GetPaidAmount() *money.Money {
	if x != nil {
		return x.PaidAmount
	}
	return nil
}

func (x *Appointment) GetRefundAmount() *money.Money {
	if x != nil {
		return x.RefundAmount
	}
	return nil
}

func (x *Appointment) GetPaymentStatus() paymentpb.Payment_Status {
	if x != nil {
		return x.PaymentStatus
	}
	return paymentpb.Payment_Status(0)
}

func (x *Appointment) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Appointment) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Appointment) GetLastUpdatedBy() string {
	if x != nil {
		return x.LastUpdatedBy
	}
	return ""
}

func (x *Appointment) GetLastUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedTime
	}
	return nil
}

var File_moego_business_appointment_v1_appointment_proto protoreflect.FileDescriptor

var file_moego_business_appointment_v1_appointment_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x36, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x08, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x5f, 0x0a, 0x13, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d,
	0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x11, 0x70,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x61, 0x69, 0x64, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x50, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x46, 0x0a,
	0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x45, 0x43, 0x4b,
	0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x42, 0x9e,
	0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x3b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_appointment_v1_appointment_proto_rawDescOnce sync.Once
	file_moego_business_appointment_v1_appointment_proto_rawDescData = file_moego_business_appointment_v1_appointment_proto_rawDesc
)

func file_moego_business_appointment_v1_appointment_proto_rawDescGZIP() []byte {
	file_moego_business_appointment_v1_appointment_proto_rawDescOnce.Do(func() {
		file_moego_business_appointment_v1_appointment_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_appointment_v1_appointment_proto_rawDescData)
	})
	return file_moego_business_appointment_v1_appointment_proto_rawDescData
}

var file_moego_business_appointment_v1_appointment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_moego_business_appointment_v1_appointment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_moego_business_appointment_v1_appointment_proto_goTypes = []interface{}{
	(Appointment_Status)(0),       // 0: moego.business.appointment.v1.Appointment.Status
	(*Appointment)(nil),           // 1: moego.business.appointment.v1.Appointment
	(*commonpb.Address)(nil),      // 2: moego.common.v1.Address
	(*interval.Interval)(nil),     // 3: google.type.Interval
	(*PetServiceDetail)(nil),      // 4: moego.business.appointment.v1.PetServiceDetail
	(*money.Money)(nil),           // 5: google.type.Money
	(paymentpb.Payment_Status)(0), // 6: moego.business.payment.v1.Payment.Status
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_moego_business_appointment_v1_appointment_proto_depIdxs = []int32{
	2,  // 0: moego.business.appointment.v1.Appointment.address:type_name -> moego.common.v1.Address
	3,  // 1: moego.business.appointment.v1.Appointment.duration:type_name -> google.type.Interval
	4,  // 2: moego.business.appointment.v1.Appointment.pet_service_details:type_name -> moego.business.appointment.v1.PetServiceDetail
	0,  // 3: moego.business.appointment.v1.Appointment.status:type_name -> moego.business.appointment.v1.Appointment.Status
	5,  // 4: moego.business.appointment.v1.Appointment.total_amount:type_name -> google.type.Money
	5,  // 5: moego.business.appointment.v1.Appointment.paid_amount:type_name -> google.type.Money
	5,  // 6: moego.business.appointment.v1.Appointment.refund_amount:type_name -> google.type.Money
	6,  // 7: moego.business.appointment.v1.Appointment.payment_status:type_name -> moego.business.payment.v1.Payment.Status
	7,  // 8: moego.business.appointment.v1.Appointment.created_time:type_name -> google.protobuf.Timestamp
	7,  // 9: moego.business.appointment.v1.Appointment.last_updated_time:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_moego_business_appointment_v1_appointment_proto_init() }
func file_moego_business_appointment_v1_appointment_proto_init() {
	if File_moego_business_appointment_v1_appointment_proto != nil {
		return
	}
	file_moego_business_appointment_v1_pet_service_detail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_moego_business_appointment_v1_appointment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appointment); i {
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
			RawDescriptor: file_moego_business_appointment_v1_appointment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_appointment_v1_appointment_proto_goTypes,
		DependencyIndexes: file_moego_business_appointment_v1_appointment_proto_depIdxs,
		EnumInfos:         file_moego_business_appointment_v1_appointment_proto_enumTypes,
		MessageInfos:      file_moego_business_appointment_v1_appointment_proto_msgTypes,
	}.Build()
	File_moego_business_appointment_v1_appointment_proto = out.File
	file_moego_business_appointment_v1_appointment_proto_rawDesc = nil
	file_moego_business_appointment_v1_appointment_proto_goTypes = nil
	file_moego_business_appointment_v1_appointment_proto_depIdxs = nil
}
