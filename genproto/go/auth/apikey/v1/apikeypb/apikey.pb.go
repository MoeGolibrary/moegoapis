// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: moego/auth/apikey/v1/apikey.proto

package apikeypb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Key is the representation of a key that can be used to access a resource.
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id in UUID4 format
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of API key
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The Secret Key string
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	// Key restrictions
	Restrictions *Restrictions `protobuf:"bytes,4,opt,name=restrictions,proto3" json:"restrictions,omitempty"`
	// Key created timestamp
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// Key updated timestamp
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// Key expired timestamp
	ExpiredTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_auth_apikey_v1_apikey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_moego_auth_apikey_v1_apikey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_moego_auth_apikey_v1_apikey_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Key) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Key) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Key) GetRestrictions() *Restrictions {
	if x != nil {
		return x.Restrictions
	}
	return nil
}

func (x *Key) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Key) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Key) GetExpiredTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredTime
	}
	return nil
}

// Restrictions are restrictions that are applied to a key.
type Restrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scopes that are allowed to be requested by the key.
	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// The client-side restrictions that are allowed to use the key.
	// You can specify only one type of client restrictions per key.
	//
	// Types that are assignable to ClientRestrictions:
	//
	//	*Restrictions_ServerRestrictions
	ClientRestrictions isRestrictions_ClientRestrictions `protobuf_oneof:"client_restrictions"`
}

func (x *Restrictions) Reset() {
	*x = Restrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_auth_apikey_v1_apikey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restrictions) ProtoMessage() {}

func (x *Restrictions) ProtoReflect() protoreflect.Message {
	mi := &file_moego_auth_apikey_v1_apikey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restrictions.ProtoReflect.Descriptor instead.
func (*Restrictions) Descriptor() ([]byte, []int) {
	return file_moego_auth_apikey_v1_apikey_proto_rawDescGZIP(), []int{1}
}

func (x *Restrictions) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (m *Restrictions) GetClientRestrictions() isRestrictions_ClientRestrictions {
	if m != nil {
		return m.ClientRestrictions
	}
	return nil
}

func (x *Restrictions) GetServerRestrictions() *ServerRestrictions {
	if x, ok := x.GetClientRestrictions().(*Restrictions_ServerRestrictions); ok {
		return x.ServerRestrictions
	}
	return nil
}

type isRestrictions_ClientRestrictions interface {
	isRestrictions_ClientRestrictions()
}

type Restrictions_ServerRestrictions struct {
	// The IP addresses of callers that are allowed to use the key.
	ServerRestrictions *ServerRestrictions `protobuf:"bytes,2,opt,name=server_restrictions,json=serverRestrictions,proto3,oneof"`
}

func (*Restrictions_ServerRestrictions) isRestrictions_ClientRestrictions() {}

// ServerRestrictions
type ServerRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	AllowedIps []string `protobuf:"bytes,1,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
}

func (x *ServerRestrictions) Reset() {
	*x = ServerRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moego_auth_apikey_v1_apikey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRestrictions) ProtoMessage() {}

func (x *ServerRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_moego_auth_apikey_v1_apikey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRestrictions.ProtoReflect.Descriptor instead.
func (*ServerRestrictions) Descriptor() ([]byte, []int) {
	return file_moego_auth_apikey_v1_apikey_proto_rawDescGZIP(), []int{2}
}

func (x *ServerRestrictions) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

var File_moego_auth_apikey_v1_apikey_proto protoreflect.FileDescriptor

var file_moego_auth_apikey_v1_apikey_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69,
	0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x02, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x46, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x9a, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48,
	0x00, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x43, 0x0a, 0x12,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0c, 0xba, 0x48, 0x09, 0x92, 0x01, 0x06, 0x22,
	0x04, 0x72, 0x02, 0x70, 0x01, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70,
	0x73, 0x42, 0x7d, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x65,
	0x47, 0x6f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x65, 0x67, 0x6f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moego_auth_apikey_v1_apikey_proto_rawDescOnce sync.Once
	file_moego_auth_apikey_v1_apikey_proto_rawDescData = file_moego_auth_apikey_v1_apikey_proto_rawDesc
)

func file_moego_auth_apikey_v1_apikey_proto_rawDescGZIP() []byte {
	file_moego_auth_apikey_v1_apikey_proto_rawDescOnce.Do(func() {
		file_moego_auth_apikey_v1_apikey_proto_rawDescData = protoimpl.X.CompressGZIP(file_moego_auth_apikey_v1_apikey_proto_rawDescData)
	})
	return file_moego_auth_apikey_v1_apikey_proto_rawDescData
}

var file_moego_auth_apikey_v1_apikey_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_moego_auth_apikey_v1_apikey_proto_goTypes = []interface{}{
	(*Key)(nil),                   // 0: moego.auth.apikey.v1.Key
	(*Restrictions)(nil),          // 1: moego.auth.apikey.v1.Restrictions
	(*ServerRestrictions)(nil),    // 2: moego.auth.apikey.v1.ServerRestrictions
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_moego_auth_apikey_v1_apikey_proto_depIdxs = []int32{
	1, // 0: moego.auth.apikey.v1.Key.restrictions:type_name -> moego.auth.apikey.v1.Restrictions
	3, // 1: moego.auth.apikey.v1.Key.created_time:type_name -> google.protobuf.Timestamp
	3, // 2: moego.auth.apikey.v1.Key.updated_time:type_name -> google.protobuf.Timestamp
	3, // 3: moego.auth.apikey.v1.Key.expired_time:type_name -> google.protobuf.Timestamp
	2, // 4: moego.auth.apikey.v1.Restrictions.server_restrictions:type_name -> moego.auth.apikey.v1.ServerRestrictions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_moego_auth_apikey_v1_apikey_proto_init() }
func file_moego_auth_apikey_v1_apikey_proto_init() {
	if File_moego_auth_apikey_v1_apikey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moego_auth_apikey_v1_apikey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_moego_auth_apikey_v1_apikey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restrictions); i {
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
		file_moego_auth_apikey_v1_apikey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerRestrictions); i {
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
	file_moego_auth_apikey_v1_apikey_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Restrictions_ServerRestrictions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moego_auth_apikey_v1_apikey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moego_auth_apikey_v1_apikey_proto_goTypes,
		DependencyIndexes: file_moego_auth_apikey_v1_apikey_proto_depIdxs,
		MessageInfos:      file_moego_auth_apikey_v1_apikey_proto_msgTypes,
	}.Build()
	File_moego_auth_apikey_v1_apikey_proto = out.File
	file_moego_auth_apikey_v1_apikey_proto_rawDesc = nil
	file_moego_auth_apikey_v1_apikey_proto_goTypes = nil
	file_moego_auth_apikey_v1_apikey_proto_depIdxs = nil
}
