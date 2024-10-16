// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: moego/business/agreement/v1/agreement_record.proto

package agreementrecordpb

import (
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

// signing status of agreement
type AgreementRecord_SignedStatus int32

const (
	// unspecified
	AgreementRecord_SIGNED_STATUS_UNSPECIFIED AgreementRecord_SignedStatus = 0
	// unsigned
	AgreementRecord_UNSIGNED AgreementRecord_SignedStatus = 1
	// signed
	AgreementRecord_SIGNED AgreementRecord_SignedStatus = 2
)

// Enum value maps for AgreementRecord_SignedStatus.
var (
	AgreementRecord_SignedStatus_name = map[int32]string{
		0: "SIGNED_STATUS_UNSPECIFIED",
		1: "UNSIGNED",
		2: "SIGNED",
	}
	AgreementRecord_SignedStatus_value = map[string]int32{
		"SIGNED_STATUS_UNSPECIFIED": 0,
		"UNSIGNED":                  1,
		"SIGNED":                    2,
	}
)

func (x AgreementRecord_SignedStatus) Enum() *AgreementRecord_SignedStatus {
	p := new(AgreementRecord_SignedStatus)
	*p = x
	return p
}

func (x AgreementRecord_SignedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgreementRecord_SignedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_agreement_v1_agreement_record_proto_enumTypes[0].Descriptor()
}

func (AgreementRecord_SignedStatus) Type() protoreflect.EnumType {
	return &file_moego_business_agreement_v1_agreement_record_proto_enumTypes[0]
}

func (x AgreementRecord_SignedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgreementRecord_SignedStatus.Descriptor instead.
func (AgreementRecord_SignedStatus) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_agreement_v1_agreement_record_proto_rawDescGZIP(), []int{0, 0}
}

// signed an agreement type
type AgreementRecord_SignedType int32

const (
	// unspecified
	AgreementRecord_SIGNED_TYPE_UNSPECIFIED AgreementRecord_SignedType = 0
	// customer signature
	AgreementRecord_CUSTOMER_SIGNED AgreementRecord_SignedType = 1
	// business upload file
	AgreementRecord_BY_BUSINESS_UPLOAD AgreementRecord_SignedType = 2
)

// Enum value maps for AgreementRecord_SignedType.
var (
	AgreementRecord_SignedType_name = map[int32]string{
		0: "SIGNED_TYPE_UNSPECIFIED",
		1: "CUSTOMER_SIGNED",
		2: "BY_BUSINESS_UPLOAD",
	}
	AgreementRecord_SignedType_value = map[string]int32{
		"SIGNED_TYPE_UNSPECIFIED": 0,
		"CUSTOMER_SIGNED":         1,
		"BY_BUSINESS_UPLOAD":      2,
	}
)

func (x AgreementRecord_SignedType) Enum() *AgreementRecord_SignedType {
	p := new(AgreementRecord_SignedType)
	*p = x
	return p
}

func (x AgreementRecord_SignedType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgreementRecord_SignedType) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_agreement_v1_agreement_record_proto_enumTypes[1].Descriptor()
}

func (AgreementRecord_SignedType) Type() protoreflect.EnumType {
	return &file_moego_business_agreement_v1_agreement_record_proto_enumTypes[1]
}

func (x AgreementRecord_SignedType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgreementRecord_SignedType.Descriptor instead.
func (AgreementRecord_SignedType) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_agreement_v1_agreement_record_proto_rawDescGZIP(), []int{0, 1}
}

// agreement source type
type AgreementRecord_SourceType int32

const (
	// unspecified
	AgreementRecord_SOURCE_TYPE_UNSPECIFIED AgreementRecord_SourceType = 0
	// from web url
	AgreementRecord_URL AgreementRecord_SourceType = 1
	// from mobile app
	AgreementRecord_MOBILE AgreementRecord_SourceType = 2
	// from online booking
	AgreementRecord_ONLINE_BOOKING AgreementRecord_SourceType = 3
	// from intake form
	AgreementRecord_INTAKE_FORM AgreementRecord_SourceType = 4
)

// Enum value maps for AgreementRecord_SourceType.
var (
	AgreementRecord_SourceType_name = map[int32]string{
		0: "SOURCE_TYPE_UNSPECIFIED",
		1: "URL",
		2: "MOBILE",
		3: "ONLINE_BOOKING",
		4: "INTAKE_FORM",
	}
	AgreementRecord_SourceType_value = map[string]int32{
		"SOURCE_TYPE_UNSPECIFIED": 0,
		"URL":                     1,
		"MOBILE":                  2,
		"ONLINE_BOOKING":          3,
		"INTAKE_FORM":             4,
	}
)

func (x AgreementRecord_SourceType) Enum() *AgreementRecord_SourceType {
	p := new(AgreementRecord_SourceType)
	*p = x
	return p
}

func (x AgreementRecord_SourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgreementRecord_SourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_agreement_v1_agreement_record_proto_enumTypes[2].Descriptor()
}

func (AgreementRecord_SourceType) Type() protoreflect.EnumType {
	return &file_moego_business_agreement_v1_agreement_record_proto_enumTypes[2]
}

func (x AgreementRecord_SourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgreementRecord_SourceType.Descriptor instead.
func (AgreementRecord_SourceType) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_agreement_v1_agreement_record_proto_rawDescGZIP(), []int{0, 2}
}

// Status
type AgreementRecord_Status int32

const (
	// unspecified
	AgreementRecord_STATUS_UNSPECIFIED AgreementRecord_Status = 0
	// normal
	AgreementRecord_NORMAL AgreementRecord_Status = 1
	// deleted
	AgreementRecord_DELETED AgreementRecord_Status = 2
)

// Enum value maps for AgreementRecord_Status.
var (
	AgreementRecord_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "NORMAL",
		2: "DELETED",
	}
	AgreementRecord_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"NORMAL":             1,
		"DELETED":            2,
	}
)

func (x AgreementRecord_Status) Enum() *AgreementRecord_Status {
	p := new(AgreementRecord_Status)
	*p = x
	return p
}

func (x AgreementRecord_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgreementRecord_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_agreement_v1_agreement_record_proto_enumTypes[3].Descriptor()
}

func (AgreementRecord_Status) Type() protoreflect.EnumType {
	return &file_moego_business_agreement_v1_agreement_record_proto_enumTypes[3]
}

func (x AgreementRecord_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgreementRecord_Status.Descriptor instead.
func (AgreementRecord_Status) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_agreement_v1_agreement_record_proto_rawDescGZIP(), []int{0, 3}
}

type AgreementRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// agreement record id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// agreement record uuid
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// agreement id
	AgreementId string `protobuf:"bytes,3,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	// business id
	BusinessId string `protobuf:"bytes,4,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	// company id
	CompanyId string `protobuf:"bytes,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// customer id
	CustomerId string `protobuf:"bytes,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// associated target id
	TargetId string `protobuf:"bytes,7,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// status: normal, deleted
	Status AgreementRecord_Status `protobuf:"varint,8,opt,name=status,proto3,enum=moego.business.agreement.v1.AgreementRecord_Status" json:"status,omitempty"`
	// signed status: unsigned, signed
	SignedStatus AgreementRecord_SignedStatus `protobuf:"varint,9,opt,name=signed_status,json=signedStatus,proto3,enum=moego.business.agreement.v1.AgreementRecord_SignedStatus" json:"signed_status,omitempty"`
	// signed type, see definition in SignedType
	SignedType AgreementRecord_SignedType `protobuf:"varint,10,opt,name=signed_type,json=signedType,proto3,enum=moego.business.agreement.v1.AgreementRecord_SignedType" json:"signed_type,omitempty"`
	// source type, see definition in SourceType
	SourceType AgreementRecord_SourceType `protobuf:"varint,11,opt,name=source_type,json=sourceType,proto3,enum=moego.business.agreement.v1.AgreementRecord_SourceType" json:"source_type,omitempty"`
	// agreement record link
	Link string `protobuf:"bytes,12,opt,name=link,proto3" json:"link,omitempty"`
	// agreement title
	Title string `protobuf:"bytes,13,opt,name=title,proto3" json:"title,omitempty"`
	// agreement content
	Content string `protobuf:"bytes,14,opt,name=content,proto3" json:"content,omitempty"`
	// customer signature
	Signature string `protobuf:"bytes,15,opt,name=signature,proto3" json:"signature,omitempty"`
	// signed time: milliseconds
	SignedTime *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=signed_time,json=signedTime,proto3" json:"signed_time,omitempty"`
	// create time: milliseconds
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// update time: milliseconds
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
}

func (x *AgreementRecord) Reset() {
	*x = AgreementRecord{}
	mi := &file_moego_business_agreement_v1_agreement_record_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgreementRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreementRecord) ProtoMessage() {}

func (x *AgreementRecord) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_agreement_v1_agreement_record_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreementRecord.ProtoReflect.Descriptor instead.
func (*AgreementRecord) Descriptor() ([]byte, []int) {
	return file_moego_business_agreement_v1_agreement_record_proto_rawDescGZIP(), []int{0}
}

func (x *AgreementRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgreementRecord) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AgreementRecord) GetAgreementId() string {
	if x != nil {
		return x.AgreementId
	}
	return ""
}

func (x *AgreementRecord) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *AgreementRecord) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *AgreementRecord) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *AgreementRecord) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *AgreementRecord) GetStatus() AgreementRecord_Status {
	if x != nil {
		return x.Status
	}
	return AgreementRecord_STATUS_UNSPECIFIED
}

func (x *AgreementRecord) GetSignedStatus() AgreementRecord_SignedStatus {
	if x != nil {
		return x.SignedStatus
	}
	return AgreementRecord_SIGNED_STATUS_UNSPECIFIED
}

func (x *AgreementRecord) GetSignedType() AgreementRecord_SignedType {
	if x != nil {
		return x.SignedType
	}
	return AgreementRecord_SIGNED_TYPE_UNSPECIFIED
}

func (x *AgreementRecord) GetSourceType() AgreementRecord_SourceType {
	if x != nil {
		return x.SourceType
	}
	return AgreementRecord_SOURCE_TYPE_UNSPECIFIED
}

func (x *AgreementRecord) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *AgreementRecord) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AgreementRecord) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AgreementRecord) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AgreementRecord) GetSignedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SignedTime
	}
	return nil
}

func (x *AgreementRecord) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *AgreementRecord) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

var File_moego_business_agreement_v1_agreement_record_proto protoreflect.FileDescriptor

var file_moego_business_agreement_v1_agreement_record_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67,
	0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x09, 0x0a, 0x0f, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67,
	0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x67, 0x72,
	0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5e, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x39, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x58, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x0c, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x49, 0x47,
	0x4e, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x53, 0x49,
	0x47, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44,
	0x10, 0x02, 0x22, 0x56, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x17, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x59, 0x5f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53,
	0x53, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x02, 0x22, 0x63, 0x0a, 0x0a, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x4e,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x49, 0x4e, 0x54, 0x41, 0x4b, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x04, 0x22,
	0x39, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0xa6, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x14, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x70,
	0x62, 0x3b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_agreement_v1_agreement_record_proto_rawDescOnce sync.Once
	file_moego_business_agreement_v1_agreement_record_proto_rawDescData = file_moego_business_agreement_v1_agreement_record_proto_rawDesc
)

func file_moego_business_agreement_v1_agreement_record_proto_rawDescGZIP() []byte {
	file_moego_business_agreement_v1_agreement_record_proto_rawDescOnce.Do(func() {
		file_moego_business_agreement_v1_agreement_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_agreement_v1_agreement_record_proto_rawDescData)
	})
	return file_moego_business_agreement_v1_agreement_record_proto_rawDescData
}

var file_moego_business_agreement_v1_agreement_record_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_moego_business_agreement_v1_agreement_record_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_moego_business_agreement_v1_agreement_record_proto_goTypes = []any{
	(AgreementRecord_SignedStatus)(0), // 0: moego.business.agreement.v1.AgreementRecord.SignedStatus
	(AgreementRecord_SignedType)(0),   // 1: moego.business.agreement.v1.AgreementRecord.SignedType
	(AgreementRecord_SourceType)(0),   // 2: moego.business.agreement.v1.AgreementRecord.SourceType
	(AgreementRecord_Status)(0),       // 3: moego.business.agreement.v1.AgreementRecord.Status
	(*AgreementRecord)(nil),           // 4: moego.business.agreement.v1.AgreementRecord
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
}
var file_moego_business_agreement_v1_agreement_record_proto_depIdxs = []int32{
	3, // 0: moego.business.agreement.v1.AgreementRecord.status:type_name -> moego.business.agreement.v1.AgreementRecord.Status
	0, // 1: moego.business.agreement.v1.AgreementRecord.signed_status:type_name -> moego.business.agreement.v1.AgreementRecord.SignedStatus
	1, // 2: moego.business.agreement.v1.AgreementRecord.signed_type:type_name -> moego.business.agreement.v1.AgreementRecord.SignedType
	2, // 3: moego.business.agreement.v1.AgreementRecord.source_type:type_name -> moego.business.agreement.v1.AgreementRecord.SourceType
	5, // 4: moego.business.agreement.v1.AgreementRecord.signed_time:type_name -> google.protobuf.Timestamp
	5, // 5: moego.business.agreement.v1.AgreementRecord.created_time:type_name -> google.protobuf.Timestamp
	5, // 6: moego.business.agreement.v1.AgreementRecord.updated_time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_moego_business_agreement_v1_agreement_record_proto_init() }
func file_moego_business_agreement_v1_agreement_record_proto_init() {
	if File_moego_business_agreement_v1_agreement_record_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moego_business_agreement_v1_agreement_record_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_agreement_v1_agreement_record_proto_goTypes,
		DependencyIndexes: file_moego_business_agreement_v1_agreement_record_proto_depIdxs,
		EnumInfos:         file_moego_business_agreement_v1_agreement_record_proto_enumTypes,
		MessageInfos:      file_moego_business_agreement_v1_agreement_record_proto_msgTypes,
	}.Build()
	File_moego_business_agreement_v1_agreement_record_proto = out.File
	file_moego_business_agreement_v1_agreement_record_proto_rawDesc = nil
	file_moego_business_agreement_v1_agreement_record_proto_goTypes = nil
	file_moego_business_agreement_v1_agreement_record_proto_depIdxs = nil
}
