// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: moego/business/setting/setting_service.proto

package customerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SettingService_AppendPetCode_FullMethodName = "/moego.business.customer.v1.SettingService/AppendPetCode"
)

// SettingServiceClient is the client API for SettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingServiceClient interface {
	// Create a pet code
	AppendPetCode(ctx context.Context, in *AppendPetCodeRequest, opts ...grpc.CallOption) (*AppendPetCodeResponse, error)
}

type settingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingServiceClient(cc grpc.ClientConnInterface) SettingServiceClient {
	return &settingServiceClient{cc}
}

func (c *settingServiceClient) AppendPetCode(ctx context.Context, in *AppendPetCodeRequest, opts ...grpc.CallOption) (*AppendPetCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppendPetCodeResponse)
	err := c.cc.Invoke(ctx, SettingService_AppendPetCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingServiceServer is the server API for SettingService service.
// All implementations must embed UnimplementedSettingServiceServer
// for forward compatibility
type SettingServiceServer interface {
	// Create a pet code
	AppendPetCode(context.Context, *AppendPetCodeRequest) (*AppendPetCodeResponse, error)
	mustEmbedUnimplementedSettingServiceServer()
}

// UnimplementedSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSettingServiceServer struct {
}

func (UnimplementedSettingServiceServer) AppendPetCode(context.Context, *AppendPetCodeRequest) (*AppendPetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendPetCode not implemented")
}
func (UnimplementedSettingServiceServer) mustEmbedUnimplementedSettingServiceServer() {}

// UnsafeSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingServiceServer will
// result in compilation errors.
type UnsafeSettingServiceServer interface {
	mustEmbedUnimplementedSettingServiceServer()
}

func RegisterSettingServiceServer(s grpc.ServiceRegistrar, srv SettingServiceServer) {
	s.RegisterService(&SettingService_ServiceDesc, srv)
}

func _SettingService_AppendPetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendPetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).AppendPetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_AppendPetCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).AppendPetCode(ctx, req.(*AppendPetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingService_ServiceDesc is the grpc.ServiceDesc for SettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.customer.v1.SettingService",
	HandlerType: (*SettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendPetCode",
			Handler:    _SettingService_AppendPetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/setting/setting_service.proto",
}
