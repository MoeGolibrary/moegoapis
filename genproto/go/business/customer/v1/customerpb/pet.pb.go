// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: moego/business/customer/v1/pet.proto

package customerpb

import (
	commonpb "github.com/MoeGolibrary/moegoapis/genproto/go/common/v1/commonpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
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

// PetType is the type of pet
type Pet_Type int32

const (
	Pet_TYPE_UNSPECIFIED Pet_Type = 0
	Pet_OTHER            Pet_Type = 1
	Pet_DOG              Pet_Type = 2
	Pet_CAT              Pet_Type = 3
	Pet_BIRD             Pet_Type = 4
	Pet_RABBIT           Pet_Type = 5
	Pet_GUINEA_PIG       Pet_Type = 6
	Pet_HORSE            Pet_Type = 7
	Pet_HAMSTER          Pet_Type = 8
	Pet_RAT              Pet_Type = 9
	Pet_MOUSE            Pet_Type = 10
	Pet_CHINCHILLA       Pet_Type = 11
)

// Enum value maps for Pet_Type.
var (
	Pet_Type_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "OTHER",
		2:  "DOG",
		3:  "CAT",
		4:  "BIRD",
		5:  "RABBIT",
		6:  "GUINEA_PIG",
		7:  "HORSE",
		8:  "HAMSTER",
		9:  "RAT",
		10: "MOUSE",
		11: "CHINCHILLA",
	}
	Pet_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"OTHER":            1,
		"DOG":              2,
		"CAT":              3,
		"BIRD":             4,
		"RABBIT":           5,
		"GUINEA_PIG":       6,
		"HORSE":            7,
		"HAMSTER":          8,
		"RAT":              9,
		"MOUSE":            10,
		"CHINCHILLA":       11,
	}
)

func (x Pet_Type) Enum() *Pet_Type {
	p := new(Pet_Type)
	*p = x
	return p
}

func (x Pet_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pet_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_customer_v1_pet_proto_enumTypes[0].Descriptor()
}

func (Pet_Type) Type() protoreflect.EnumType {
	return &file_moego_business_customer_v1_pet_proto_enumTypes[0]
}

func (x Pet_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pet_Type.Descriptor instead.
func (Pet_Type) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 0}
}

// Status is the status of a pet
type Pet_Status int32

const (
	Pet_STATUS_UNSPECIFIED Pet_Status = 0
	Pet_ALIVE              Pet_Status = 1
	Pet_PASSED_AWAY        Pet_Status = 2
)

// Enum value maps for Pet_Status.
var (
	Pet_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ALIVE",
		2: "PASSED_AWAY",
	}
	Pet_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ALIVE":              1,
		"PASSED_AWAY":        2,
	}
)

func (x Pet_Status) Enum() *Pet_Status {
	p := new(Pet_Status)
	*p = x
	return p
}

func (x Pet_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pet_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_customer_v1_pet_proto_enumTypes[1].Descriptor()
}

func (Pet_Status) Type() protoreflect.EnumType {
	return &file_moego_business_customer_v1_pet_proto_enumTypes[1]
}

func (x Pet_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pet_Status.Descriptor instead.
func (Pet_Status) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 1}
}

// Gender
type Pet_Gender int32

const (
	Pet_GENDER_UNSPECIFIED Pet_Gender = 0
	Pet_MALE               Pet_Gender = 1
	Pet_FEMALE             Pet_Gender = 2
	Pet_UNKNOWN            Pet_Gender = 3
)

// Enum value maps for Pet_Gender.
var (
	Pet_Gender_name = map[int32]string{
		0: "GENDER_UNSPECIFIED",
		1: "MALE",
		2: "FEMALE",
		3: "UNKNOWN",
	}
	Pet_Gender_value = map[string]int32{
		"GENDER_UNSPECIFIED": 0,
		"MALE":               1,
		"FEMALE":             2,
		"UNKNOWN":            3,
	}
)

func (x Pet_Gender) Enum() *Pet_Gender {
	p := new(Pet_Gender)
	*p = x
	return p
}

func (x Pet_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pet_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_moego_business_customer_v1_pet_proto_enumTypes[2].Descriptor()
}

func (Pet_Gender) Type() protoreflect.EnumType {
	return &file_moego_business_customer_v1_pet_proto_enumTypes[2]
}

func (x Pet_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pet_Gender.Descriptor instead.
func (Pet_Gender) EnumDescriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 2}
}

// Pet is a pet
type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of a pet
	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Birthday *date.Date       `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Status   Pet_Status       `protobuf:"varint,4,opt,name=status,proto3,enum=moego.business.customer.v1.Pet_Status" json:"status,omitempty"`
	Type     Pet_Type         `protobuf:"varint,5,opt,name=type,proto3,enum=moego.business.customer.v1.Pet_Type" json:"type,omitempty"`
	Breed    string           `protobuf:"bytes,6,opt,name=breed,proto3" json:"breed,omitempty"`
	Gender   Pet_Gender       `protobuf:"varint,7,opt,name=gender,proto3,enum=moego.business.customer.v1.Pet_Gender" json:"gender,omitempty"`
	Weight   *commonpb.Weight `protobuf:"bytes,8,opt,name=weight,proto3" json:"weight,omitempty"`
	Fixed    string           `protobuf:"bytes,9,opt,name=fixed,proto3" json:"fixed,omitempty"`
	Coat     string           `protobuf:"bytes,10,opt,name=coat,proto3" json:"coat,omitempty"`
	Behavior string           `protobuf:"bytes,11,opt,name=behavior,proto3" json:"behavior,omitempty"`
	PetCodes []*Pet_Code      `protobuf:"bytes,12,rep,name=pet_codes,json=petCodes,proto3" json:"pet_codes,omitempty"`
	Notes    []*Pet_Note      `protobuf:"bytes,13,rep,name=notes,proto3" json:"notes,omitempty"`
	// This field is not supported when creating a pet.
	Vaccinations []*Pet_Vaccination `protobuf:"bytes,14,rep,name=vaccinations,proto3" json:"vaccinations,omitempty"`
	CustomerId   string             `protobuf:"bytes,15,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Vet          *Pet_Vet           `protobuf:"bytes,16,opt,name=vet,proto3" json:"vet,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetBirthday() *date.Date {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *Pet) GetStatus() Pet_Status {
	if x != nil {
		return x.Status
	}
	return Pet_STATUS_UNSPECIFIED
}

func (x *Pet) GetType() Pet_Type {
	if x != nil {
		return x.Type
	}
	return Pet_TYPE_UNSPECIFIED
}

func (x *Pet) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

func (x *Pet) GetGender() Pet_Gender {
	if x != nil {
		return x.Gender
	}
	return Pet_GENDER_UNSPECIFIED
}

func (x *Pet) GetWeight() *commonpb.Weight {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *Pet) GetFixed() string {
	if x != nil {
		return x.Fixed
	}
	return ""
}

func (x *Pet) GetCoat() string {
	if x != nil {
		return x.Coat
	}
	return ""
}

func (x *Pet) GetBehavior() string {
	if x != nil {
		return x.Behavior
	}
	return ""
}

func (x *Pet) GetPetCodes() []*Pet_Code {
	if x != nil {
		return x.PetCodes
	}
	return nil
}

func (x *Pet) GetNotes() []*Pet_Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *Pet) GetVaccinations() []*Pet_Vaccination {
	if x != nil {
		return x.Vaccinations
	}
	return nil
}

func (x *Pet) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Pet) GetVet() *Pet_Vet {
	if x != nil {
		return x.Vet
	}
	return nil
}

// Code
type Pet_Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Abbreviation string `protobuf:"bytes,2,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Color        string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Pet_Code) Reset() {
	*x = Pet_Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet_Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet_Code) ProtoMessage() {}

func (x *Pet_Code) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet_Code.ProtoReflect.Descriptor instead.
func (*Pet_Code) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Pet_Code) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet_Code) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *Pet_Code) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Pet_Code) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

// Note
type Pet_Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content         string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	LastUpdatedBy   string                 `protobuf:"bytes,3,opt,name=last_updated_by,json=lastUpdatedBy,proto3" json:"last_updated_by,omitempty"`
	LastUpdatedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_updated_time,json=lastUpdatedTime,proto3" json:"last_updated_time,omitempty"`
}

func (x *Pet_Note) Reset() {
	*x = Pet_Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet_Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet_Note) ProtoMessage() {}

func (x *Pet_Note) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet_Note.ProtoReflect.Descriptor instead.
func (*Pet_Note) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Pet_Note) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet_Note) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Pet_Note) GetLastUpdatedBy() string {
	if x != nil {
		return x.LastUpdatedBy
	}
	return ""
}

func (x *Pet_Note) GetLastUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedTime
	}
	return nil
}

// Vaccination
type Pet_Vaccination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *Pet_Vaccination) Reset() {
	*x = Pet_Vaccination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet_Vaccination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet_Vaccination) ProtoMessage() {}

func (x *Pet_Vaccination) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet_Vaccination.ProtoReflect.Descriptor instead.
func (*Pet_Vaccination) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Pet_Vaccination) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet_Vaccination) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

// Vet
type Pet_Vet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Pet_Vet) Reset() {
	*x = Pet_Vet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_business_customer_v1_pet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet_Vet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet_Vet) ProtoMessage() {}

func (x *Pet_Vet) ProtoReflect() protoreflect.Message {
	mi := &file_moego_business_customer_v1_pet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet_Vet.ProtoReflect.Descriptor instead.
func (*Pet_Vet) Descriptor() ([]byte, []int) {
	return file_moego_business_customer_v1_pet_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Pet_Vet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet_Vet) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Pet_Vet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_moego_business_customer_v1_pet_proto protoreflect.FileDescriptor

var file_moego_business_customer_v1_pet_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x6f,
	0x65, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x0b, 0x0a, 0x03, 0x50,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x65,
	0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x62, 0x72, 0x65,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x62,
	0x72, 0x65, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x61, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x09, 0x70,
	0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x2e, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x76, 0x61,
	0x63, 0x63, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x74, 0x2e, 0x56, 0x61, 0x63, 0x63, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x76,
	0x61, 0x63, 0x63, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x03, 0x76, 0x65, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x2e,
	0x56, 0x65, 0x74, 0x52, 0x03, 0x76, 0x65, 0x74, 0x1a, 0x72, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x1a, 0xa0, 0x01, 0x0a,
	0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x46, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a,
	0x5c, 0x0a, 0x0b, 0x56, 0x61, 0x63, 0x63, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x56, 0x0a,
	0x03, 0x56, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x4f, 0x47, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x41, 0x54, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x49, 0x52, 0x44, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x41, 0x42, 0x42, 0x49, 0x54, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x55, 0x49, 0x4e, 0x45,
	0x41, 0x5f, 0x50, 0x49, 0x47, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x4f, 0x52, 0x53, 0x45,
	0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x41, 0x4d, 0x53, 0x54, 0x45, 0x52, 0x10, 0x08, 0x12,
	0x07, 0x0a, 0x03, 0x52, 0x41, 0x54, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x55, 0x53,
	0x45, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x49, 0x4e, 0x43, 0x48, 0x49, 0x4c, 0x4c,
	0x41, 0x10, 0x0b, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x5f, 0x41, 0x57, 0x41, 0x59, 0x10,
	0x02, 0x22, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x42, 0x8a, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x50,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_business_customer_v1_pet_proto_rawDescOnce sync.Once
	file_moego_business_customer_v1_pet_proto_rawDescData = file_moego_business_customer_v1_pet_proto_rawDesc
)

func file_moego_business_customer_v1_pet_proto_rawDescGZIP() []byte {
	file_moego_business_customer_v1_pet_proto_rawDescOnce.Do(func() {
		file_moego_business_customer_v1_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_business_customer_v1_pet_proto_rawDescData)
	})
	return file_moego_business_customer_v1_pet_proto_rawDescData
}

var file_moego_business_customer_v1_pet_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_moego_business_customer_v1_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_moego_business_customer_v1_pet_proto_goTypes = []any{
	(Pet_Type)(0),                 // 0: moego.business.customer.v1.Pet.Type
	(Pet_Status)(0),               // 1: moego.business.customer.v1.Pet.Status
	(Pet_Gender)(0),               // 2: moego.business.customer.v1.Pet.Gender
	(*Pet)(nil),                   // 3: moego.business.customer.v1.Pet
	(*Pet_Code)(nil),              // 4: moego.business.customer.v1.Pet.Code
	(*Pet_Note)(nil),              // 5: moego.business.customer.v1.Pet.Note
	(*Pet_Vaccination)(nil),       // 6: moego.business.customer.v1.Pet.Vaccination
	(*Pet_Vet)(nil),               // 7: moego.business.customer.v1.Pet.Vet
	(*date.Date)(nil),             // 8: google.type.Date
	(*commonpb.Weight)(nil),       // 9: moego.common.v1.Weight
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_moego_business_customer_v1_pet_proto_depIdxs = []int32{
	8,  // 0: moego.business.customer.v1.Pet.birthday:type_name -> google.type.Date
	1,  // 1: moego.business.customer.v1.Pet.status:type_name -> moego.business.customer.v1.Pet.Status
	0,  // 2: moego.business.customer.v1.Pet.type:type_name -> moego.business.customer.v1.Pet.Type
	2,  // 3: moego.business.customer.v1.Pet.gender:type_name -> moego.business.customer.v1.Pet.Gender
	9,  // 4: moego.business.customer.v1.Pet.weight:type_name -> moego.common.v1.Weight
	4,  // 5: moego.business.customer.v1.Pet.pet_codes:type_name -> moego.business.customer.v1.Pet.Code
	5,  // 6: moego.business.customer.v1.Pet.notes:type_name -> moego.business.customer.v1.Pet.Note
	6,  // 7: moego.business.customer.v1.Pet.vaccinations:type_name -> moego.business.customer.v1.Pet.Vaccination
	7,  // 8: moego.business.customer.v1.Pet.vet:type_name -> moego.business.customer.v1.Pet.Vet
	10, // 9: moego.business.customer.v1.Pet.Note.last_updated_time:type_name -> google.protobuf.Timestamp
	10, // 10: moego.business.customer.v1.Pet.Vaccination.expired_at:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_moego_business_customer_v1_pet_proto_init() }
func file_moego_business_customer_v1_pet_proto_init() {
	if File_moego_business_customer_v1_pet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_business_customer_v1_pet_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Pet); i {
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
		file_moego_business_customer_v1_pet_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Pet_Code); i {
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
		file_moego_business_customer_v1_pet_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Pet_Note); i {
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
		file_moego_business_customer_v1_pet_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Pet_Vaccination); i {
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
		file_moego_business_customer_v1_pet_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Pet_Vet); i {
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
			RawDescriptor: file_moego_business_customer_v1_pet_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_business_customer_v1_pet_proto_goTypes,
		DependencyIndexes: file_moego_business_customer_v1_pet_proto_depIdxs,
		EnumInfos:         file_moego_business_customer_v1_pet_proto_enumTypes,
		MessageInfos:      file_moego_business_customer_v1_pet_proto_msgTypes,
	}.Build()
	File_moego_business_customer_v1_pet_proto = out.File
	file_moego_business_customer_v1_pet_proto_rawDesc = nil
	file_moego_business_customer_v1_pet_proto_goTypes = nil
	file_moego_business_customer_v1_pet_proto_depIdxs = nil
}
