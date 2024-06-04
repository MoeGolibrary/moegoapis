// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: moego/business/customer/v1/customer_service.proto

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
	CustomerService_CreateCustomer_FullMethodName     = "/moego.business.customer.v1.CustomerService/CreateCustomer"
	CustomerService_GetCustomer_FullMethodName        = "/moego.business.customer.v1.CustomerService/GetCustomer"
	CustomerService_ListCustomers_FullMethodName      = "/moego.business.customer.v1.CustomerService/ListCustomers"
	CustomerService_GenCustomerCofLink_FullMethodName = "/moego.business.customer.v1.CustomerService/GenCustomerCofLink"
)

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// CustomerService openapi definitions for operate customer
type CustomerServiceClient interface {
	// Create a customer
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	// Get a customer
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	// List customers
	ListCustomers(ctx context.Context, in *ListCustomersRequest, opts ...grpc.CallOption) (*ListCustomersResponse, error)
	// Generate a link to add customer card on file
	GenCustomerCofLink(ctx context.Context, in *GenCustomerCofLinkRequest, opts ...grpc.CallOption) (*GenCustomerCofLinkResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Customer)
	err := c.cc.Invoke(ctx, CustomerService_CreateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*Customer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Customer)
	err := c.cc.Invoke(ctx, CustomerService_GetCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ListCustomers(ctx context.Context, in *ListCustomersRequest, opts ...grpc.CallOption) (*ListCustomersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCustomersResponse)
	err := c.cc.Invoke(ctx, CustomerService_ListCustomers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GenCustomerCofLink(ctx context.Context, in *GenCustomerCofLinkRequest, opts ...grpc.CallOption) (*GenCustomerCofLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenCustomerCofLinkResponse)
	err := c.cc.Invoke(ctx, CustomerService_GenCustomerCofLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
//
// CustomerService openapi definitions for operate customer
type CustomerServiceServer interface {
	// Create a customer
	CreateCustomer(context.Context, *CreateCustomerRequest) (*Customer, error)
	// Get a customer
	GetCustomer(context.Context, *GetCustomerRequest) (*Customer, error)
	// List customers
	ListCustomers(context.Context, *ListCustomersRequest) (*ListCustomersResponse, error)
	// Generate a link to add customer card on file
	GenCustomerCofLink(context.Context, *GenCustomerCofLinkRequest) (*GenCustomerCofLinkResponse, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomer(context.Context, *GetCustomerRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) ListCustomers(context.Context, *ListCustomersRequest) (*ListCustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomers not implemented")
}
func (UnimplementedCustomerServiceServer) GenCustomerCofLink(context.Context, *GenCustomerCofLinkRequest) (*GenCustomerCofLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenCustomerCofLink not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_CreateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ListCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ListCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_ListCustomers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ListCustomers(ctx, req.(*ListCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GenCustomerCofLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenCustomerCofLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GenCustomerCofLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GenCustomerCofLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GenCustomerCofLink(ctx, req.(*GenCustomerCofLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.customer.v1.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "ListCustomers",
			Handler:    _CustomerService_ListCustomers_Handler,
		},
		{
			MethodName: "GenCustomerCofLink",
			Handler:    _CustomerService_GenCustomerCofLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/customer/v1/customer_service.proto",
}
