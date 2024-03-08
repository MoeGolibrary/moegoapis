// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: moego/business/discount/v1/discount_service.proto

package discountpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiscountService_CreateDiscount_FullMethodName = "/moego.business.discount.v1.DiscountService/CreateDiscount"
	DiscountService_GetDiscount_FullMethodName    = "/moego.business.discount.v1.DiscountService/GetDiscount"
	DiscountService_ListDiscounts_FullMethodName  = "/moego.business.discount.v1.DiscountService/ListDiscounts"
)

// DiscountServiceClient is the client API for DiscountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountServiceClient interface {
	// CreateDiscount
	CreateDiscount(ctx context.Context, in *CreateDiscountRequest, opts ...grpc.CallOption) (*Discount, error)
	// GetDiscount
	GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*Discount, error)
	// ListDiscounts
	ListDiscounts(ctx context.Context, in *ListDiscountsRequest, opts ...grpc.CallOption) (*ListDiscountsResponse, error)
}

type discountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountServiceClient(cc grpc.ClientConnInterface) DiscountServiceClient {
	return &discountServiceClient{cc}
}

func (c *discountServiceClient) CreateDiscount(ctx context.Context, in *CreateDiscountRequest, opts ...grpc.CallOption) (*Discount, error) {
	out := new(Discount)
	err := c.cc.Invoke(ctx, DiscountService_CreateDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountServiceClient) GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*Discount, error) {
	out := new(Discount)
	err := c.cc.Invoke(ctx, DiscountService_GetDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discountServiceClient) ListDiscounts(ctx context.Context, in *ListDiscountsRequest, opts ...grpc.CallOption) (*ListDiscountsResponse, error) {
	out := new(ListDiscountsResponse)
	err := c.cc.Invoke(ctx, DiscountService_ListDiscounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountServiceServer is the server API for DiscountService service.
// All implementations must embed UnimplementedDiscountServiceServer
// for forward compatibility
type DiscountServiceServer interface {
	// CreateDiscount
	CreateDiscount(context.Context, *CreateDiscountRequest) (*Discount, error)
	// GetDiscount
	GetDiscount(context.Context, *GetDiscountRequest) (*Discount, error)
	// ListDiscounts
	ListDiscounts(context.Context, *ListDiscountsRequest) (*ListDiscountsResponse, error)
	mustEmbedUnimplementedDiscountServiceServer()
}

// UnimplementedDiscountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountServiceServer struct {
}

func (UnimplementedDiscountServiceServer) CreateDiscount(context.Context, *CreateDiscountRequest) (*Discount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscount not implemented")
}
func (UnimplementedDiscountServiceServer) GetDiscount(context.Context, *GetDiscountRequest) (*Discount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscount not implemented")
}
func (UnimplementedDiscountServiceServer) ListDiscounts(context.Context, *ListDiscountsRequest) (*ListDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiscounts not implemented")
}
func (UnimplementedDiscountServiceServer) mustEmbedUnimplementedDiscountServiceServer() {}

// UnsafeDiscountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountServiceServer will
// result in compilation errors.
type UnsafeDiscountServiceServer interface {
	mustEmbedUnimplementedDiscountServiceServer()
}

func RegisterDiscountServiceServer(s grpc.ServiceRegistrar, srv DiscountServiceServer) {
	s.RegisterService(&DiscountService_ServiceDesc, srv)
}

func _DiscountService_CreateDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServiceServer).CreateDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountService_CreateDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServiceServer).CreateDiscount(ctx, req.(*CreateDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountService_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServiceServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountService_GetDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServiceServer).GetDiscount(ctx, req.(*GetDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscountService_ListDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServiceServer).ListDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscountService_ListDiscounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServiceServer).ListDiscounts(ctx, req.(*ListDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscountService_ServiceDesc is the grpc.ServiceDesc for DiscountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.discount.v1.DiscountService",
	HandlerType: (*DiscountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDiscount",
			Handler:    _DiscountService_CreateDiscount_Handler,
		},
		{
			MethodName: "GetDiscount",
			Handler:    _DiscountService_GetDiscount_Handler,
		},
		{
			MethodName: "ListDiscounts",
			Handler:    _DiscountService_ListDiscounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/discount/v1/discount_service.proto",
}
