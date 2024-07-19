// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: moego/business/customer/v1/pet_service.proto

package customerpb

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

// Pet
type CreatePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Pet        *Pet   `protobuf:"bytes,2,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *CreatePetRequest) Reset() {
	*x = CreatePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetRequest) ProtoMessage() {}

func (x *CreatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetRequest.ProtoReflect.Descriptor instead.
func (*CreatePetRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePetRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreatePetRequest) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type CreatePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *CreatePetResponse) Reset() {
	*x = CreatePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetResponse) ProtoMessage() {}

func (x *CreatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetResponse.ProtoReflect.Descriptor instead.
func (*CreatePetResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

// Pet Code
type CreatePetCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId int64 `protobuf:"varint,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	PetCode   *Code `protobuf:"bytes,3,opt,name=pet_code,json=petCode,proto3" json:"pet_code,omitempty"`
}

func (x *CreatePetCodeRequest) Reset() {
	*x = CreatePetCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetCodeRequest) ProtoMessage() {}

func (x *CreatePetCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetCodeRequest.ProtoReflect.Descriptor instead.
func (*CreatePetCodeRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePetCodeRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CreatePetCodeRequest) GetPetCode() *Code {
	if x != nil {
		return x.PetCode
	}
	return nil
}

type CreatePetCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetCode *Code `protobuf:"bytes,1,opt,name=pet_code,json=petCode,proto3" json:"pet_code,omitempty"`
}

func (x *CreatePetCodeResponse) Reset() {
	*x = CreatePetCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetCodeResponse) ProtoMessage() {}

func (x *CreatePetCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetCodeResponse.ProtoReflect.Descriptor instead.
func (*CreatePetCodeResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePetCodeResponse) GetPetCode() *Code {
	if x != nil {
		return x.PetCode
	}
	return nil
}

type ListPetCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId int64 `protobuf:"varint,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *ListPetCodeRequest) Reset() {
	*x = ListPetCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetCodeRequest) ProtoMessage() {}

func (x *ListPetCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[4]
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
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListPetCodeRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type ListPetCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pet_codes
	PetCodes []*Code `protobuf:"bytes,2,rep,name=pet_codes,json=petCodes,proto3" json:"pet_codes,omitempty"`
}

func (x *ListPetCodeResponse) Reset() {
	*x = ListPetCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetCodeResponse) ProtoMessage() {}

func (x *ListPetCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[5]
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
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListPetCodeResponse) GetPetCodes() []*Code {
	if x != nil {
		return x.PetCodes
	}
	return nil
}

// Pet Note
type CreatePetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pet id
	PetId int64 `protobuf:"varint,2,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	// pet note
	Note *Note `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CreatePetNoteRequest) Reset() {
	*x = CreatePetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetNoteRequest) ProtoMessage() {}

func (x *CreatePetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetNoteRequest.ProtoReflect.Descriptor instead.
func (*CreatePetNoteRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePetNoteRequest) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

func (x *CreatePetNoteRequest) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type CreatePetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pet note
	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CreatePetNoteResponse) Reset() {
	*x = CreatePetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetNoteResponse) ProtoMessage() {}

func (x *CreatePetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetNoteResponse.ProtoReflect.Descriptor instead.
func (*CreatePetNoteResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePetNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type ListPetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *commonpb.PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	PetId      int64                       `protobuf:"varint,3,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
}

func (x *ListPetNoteRequest) Reset() {
	*x = ListPetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetNoteRequest) ProtoMessage() {}

func (x *ListPetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetNoteRequest.ProtoReflect.Descriptor instead.
func (*ListPetNoteRequest) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListPetNoteRequest) GetPagination() *commonpb.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPetNoteRequest) GetPetId() int64 {
	if x != nil {
		return x.PetId
	}
	return 0
}

type ListPetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *commonpb.PaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Note
	Notes []*Note `protobuf:"bytes,2,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *ListPetNoteResponse) Reset() {
	*x = ListPetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetNoteResponse) ProtoMessage() {}

func (x *ListPetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetNoteResponse.ProtoReflect.Descriptor instead.
func (*ListPetNoteResponse) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListPetNoteResponse) GetPagination() *commonpb.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPetNoteResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

var File_moego_business_customer_v1_pet_service_proto protoreflect.FileDescriptor

var file_moego_business_customer_v1_pet_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x70,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03,
	0x70, 0x65, 0x74, 0x22, 0x7c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12,
	0x40, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x59, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x65,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d,
	0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x07, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x09, 0x70, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x08, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x22, 0x74, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x06, 0x70, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x32, 0xa4, 0x01, 0x0a, 0x0a,
	0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a,
	0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x65,
	0x74, 0x73, 0x32, 0xe1, 0x02, 0x0a, 0x0e, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x6f, 0x65, 0x67,
	0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0xa5,
	0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x2e, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xd9, 0x02, 0x0a, 0x0e, 0x50, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa2, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0xa1,
	0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x2e, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x3a, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x96, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f,
	0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x70, 0x62,
	0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_moego_business_customer_v1_pet_service_proto_rawDescOnce sync.Once
	file_moego_business_customer_v1_pet_service_proto_rawDescData = file_moego_business_customer_v1_pet_service_proto_rawDesc
)

func file_moego_business_customer_v1_pet_service_proto_rawDescGZIP() []byte {
	file_moego_business_customer_v1_pet_service_proto_rawDescOnce.Do(func() {
		file_moego_business_customer_v1_pet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_customer_v1_pet_service_proto_rawDescData)
	})
	return file_moego_business_customer_v1_pet_service_proto_rawDescData
}

var file_moego_business_customer_v1_pet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_moego_business_customer_v1_pet_service_proto_goTypes = []any{
	(*CreatePetRequest)(nil),            // 0: moego.business.customer.v1.CreatePetRequest
	(*CreatePetResponse)(nil),           // 1: moego.business.customer.v1.CreatePetResponse
	(*CreatePetCodeRequest)(nil),        // 2: moego.business.customer.v1.CreatePetCodeRequest
	(*CreatePetCodeResponse)(nil),       // 3: moego.business.customer.v1.CreatePetCodeResponse
	(*ListPetCodeRequest)(nil),          // 4: moego.business.customer.v1.ListPetCodeRequest
	(*ListPetCodeResponse)(nil),         // 5: moego.business.customer.v1.ListPetCodeResponse
	(*CreatePetNoteRequest)(nil),        // 6: moego.business.customer.v1.CreatePetNoteRequest
	(*CreatePetNoteResponse)(nil),       // 7: moego.business.customer.v1.CreatePetNoteResponse
	(*ListPetNoteRequest)(nil),          // 8: moego.business.customer.v1.ListPetNoteRequest
	(*ListPetNoteResponse)(nil),         // 9: moego.business.customer.v1.ListPetNoteResponse
	(*Pet)(nil),                         // 10: moego.business.customer.v1.Pet
	(*Code)(nil),                        // 11: moego.business.customer.v1.Code
	(*Note)(nil),                        // 12: moego.business.customer.v1.Note
	(*commonpb.PaginationRequest)(nil),  // 13: moego.common.v1.PaginationRequest
	(*commonpb.PaginationResponse)(nil), // 14: moego.common.v1.PaginationResponse
}
var file_moego_business_customer_v1_pet_service_proto_depIdxs = []int32{
	10, // 0: moego.business.customer.v1.CreatePetRequest.pet:type_name -> moego.business.customer.v1.Pet
	10, // 1: moego.business.customer.v1.CreatePetResponse.pet:type_name -> moego.business.customer.v1.Pet
	11, // 2: moego.business.customer.v1.CreatePetCodeRequest.pet_code:type_name -> moego.business.customer.v1.Code
	11, // 3: moego.business.customer.v1.CreatePetCodeResponse.pet_code:type_name -> moego.business.customer.v1.Code
	11, // 4: moego.business.customer.v1.ListPetCodeResponse.pet_codes:type_name -> moego.business.customer.v1.Code
	12, // 5: moego.business.customer.v1.CreatePetNoteRequest.note:type_name -> moego.business.customer.v1.Note
	12, // 6: moego.business.customer.v1.CreatePetNoteResponse.note:type_name -> moego.business.customer.v1.Note
	13, // 7: moego.business.customer.v1.ListPetNoteRequest.pagination:type_name -> moego.common.v1.PaginationRequest
	14, // 8: moego.business.customer.v1.ListPetNoteResponse.pagination:type_name -> moego.common.v1.PaginationResponse
	12, // 9: moego.business.customer.v1.ListPetNoteResponse.notes:type_name -> moego.business.customer.v1.Note
	0,  // 10: moego.business.customer.v1.PetService.CreatePet:input_type -> moego.business.customer.v1.CreatePetRequest
	2,  // 11: moego.business.customer.v1.PetCodeService.CreatePetCode:input_type -> moego.business.customer.v1.CreatePetCodeRequest
	4,  // 12: moego.business.customer.v1.PetCodeService.ListPetCodes:input_type -> moego.business.customer.v1.ListPetCodeRequest
	6,  // 13: moego.business.customer.v1.PetNoteService.CreatePetNote:input_type -> moego.business.customer.v1.CreatePetNoteRequest
	8,  // 14: moego.business.customer.v1.PetNoteService.ListPetNotes:input_type -> moego.business.customer.v1.ListPetNoteRequest
	1,  // 15: moego.business.customer.v1.PetService.CreatePet:output_type -> moego.business.customer.v1.CreatePetResponse
	3,  // 16: moego.business.customer.v1.PetCodeService.CreatePetCode:output_type -> moego.business.customer.v1.CreatePetCodeResponse
	5,  // 17: moego.business.customer.v1.PetCodeService.ListPetCodes:output_type -> moego.business.customer.v1.ListPetCodeResponse
	7,  // 18: moego.business.customer.v1.PetNoteService.CreatePetNote:output_type -> moego.business.customer.v1.CreatePetNoteResponse
	9,  // 19: moego.business.customer.v1.PetNoteService.ListPetNotes:output_type -> moego.business.customer.v1.ListPetNoteResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_moego_business_customer_v1_pet_service_proto_init() }
func file_moego_business_customer_v1_pet_service_proto_init() {
	if File_moego_business_customer_v1_pet_service_proto != nil {
		return
	}
	file_moego_business_customer_v1_customer_proto_init()
	file_moego_business_customer_v1_code_proto_init()
	file_moego_business_customer_v1_note_proto_init()
	file_moego_business_customer_v1_pet_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_moego_business_customer_v1_pet_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePetRequest); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePetResponse); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePetCodeRequest); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePetCodeResponse); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePetNoteRequest); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePetNoteResponse); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListPetNoteRequest); i {
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
		file_moego_business_customer_v1_pet_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListPetNoteResponse); i {
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
			RawDescriptor: file_moego_business_customer_v1_pet_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_moego_business_customer_v1_pet_service_proto_goTypes,
		DependencyIndexes: file_moego_business_customer_v1_pet_service_proto_depIdxs,
		MessageInfos:      file_moego_business_customer_v1_pet_service_proto_msgTypes,
	}.Build()
	File_moego_business_customer_v1_pet_service_proto = out.File
	file_moego_business_customer_v1_pet_service_proto_rawDesc = nil
	file_moego_business_customer_v1_pet_service_proto_goTypes = nil
	file_moego_business_customer_v1_pet_service_proto_depIdxs = nil
}
