// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: moego/business/business/v1/business_service.proto

package businesspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BusinessService_GetBusiness_FullMethodName  = "/moego.business.business.v1.BusinessService/GetBusiness"
	BusinessService_ListBusiness_FullMethodName = "/moego.business.business.v1.BusinessService/ListBusiness"
)

// BusinessServiceClient is the client API for BusinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// BusinessService openapi definitions for operate business
type BusinessServiceClient interface {
	// GetBusiness get a business
	GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*Business, error)
	// ListBusiness list businesses
	ListBusiness(ctx context.Context, in *ListBusinessRequest, opts ...grpc.CallOption) (*ListBusinessResponse, error)
}

type businessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessServiceClient(cc grpc.ClientConnInterface) BusinessServiceClient {
	return &businessServiceClient{cc}
}

func (c *businessServiceClient) GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*Business, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Business)
	err := c.cc.Invoke(ctx, BusinessService_GetBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) ListBusiness(ctx context.Context, in *ListBusinessRequest, opts ...grpc.CallOption) (*ListBusinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBusinessResponse)
	err := c.cc.Invoke(ctx, BusinessService_ListBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServiceServer is the server API for BusinessService service.
// All implementations must embed UnimplementedBusinessServiceServer
// for forward compatibility.
//
// BusinessService openapi definitions for operate business
type BusinessServiceServer interface {
	// GetBusiness get a business
	GetBusiness(context.Context, *GetBusinessRequest) (*Business, error)
	// ListBusiness list businesses
	ListBusiness(context.Context, *ListBusinessRequest) (*ListBusinessResponse, error)
	mustEmbedUnimplementedBusinessServiceServer()
}

// UnimplementedBusinessServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBusinessServiceServer struct{}

func (UnimplementedBusinessServiceServer) GetBusiness(context.Context, *GetBusinessRequest) (*Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) ListBusiness(context.Context, *ListBusinessRequest) (*ListBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) mustEmbedUnimplementedBusinessServiceServer() {}
func (UnimplementedBusinessServiceServer) testEmbeddedByValue()                         {}

// UnsafeBusinessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServiceServer will
// result in compilation errors.
type UnsafeBusinessServiceServer interface {
	mustEmbedUnimplementedBusinessServiceServer()
}

func RegisterBusinessServiceServer(s grpc.ServiceRegistrar, srv BusinessServiceServer) {
	// If the following call pancis, it indicates UnimplementedBusinessServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BusinessService_ServiceDesc, srv)
}

func _BusinessService_GetBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).GetBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_GetBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).GetBusiness(ctx, req.(*GetBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_ListBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).ListBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_ListBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).ListBusiness(ctx, req.(*ListBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BusinessService_ServiceDesc is the grpc.ServiceDesc for BusinessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusinessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.business.v1.BusinessService",
	HandlerType: (*BusinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusiness",
			Handler:    _BusinessService_GetBusiness_Handler,
		},
		{
			MethodName: "ListBusiness",
			Handler:    _BusinessService_ListBusiness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/business/v1/business_service.proto",
}
