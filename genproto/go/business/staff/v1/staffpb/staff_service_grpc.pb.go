// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: moego/business/staff/v1/staff_service.proto

package staffpb

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
	StaffService_GetStaff_FullMethodName   = "/moego.business.staff.v1.StaffService/GetStaff"
	StaffService_ListStaffs_FullMethodName = "/moego.business.staff.v1.StaffService/ListStaffs"
)

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	// GetStaff
	GetStaff(ctx context.Context, in *GetStaffRequest, opts ...grpc.CallOption) (*Staff, error)
	// ListStaffs
	ListStaffs(ctx context.Context, in *ListStaffsRequest, opts ...grpc.CallOption) (*ListStaffsResponse, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) GetStaff(ctx context.Context, in *GetStaffRequest, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, StaffService_GetStaff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) ListStaffs(ctx context.Context, in *ListStaffsRequest, opts ...grpc.CallOption) (*ListStaffsResponse, error) {
	out := new(ListStaffsResponse)
	err := c.cc.Invoke(ctx, StaffService_ListStaffs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	// GetStaff
	GetStaff(context.Context, *GetStaffRequest) (*Staff, error)
	// ListStaffs
	ListStaffs(context.Context, *ListStaffsRequest) (*ListStaffsResponse, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) GetStaff(context.Context, *GetStaffRequest) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaff not implemented")
}
func (UnimplementedStaffServiceServer) ListStaffs(context.Context, *ListStaffsRequest) (*ListStaffsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStaffs not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_GetStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_GetStaff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetStaff(ctx, req.(*GetStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_ListStaffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).ListStaffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_ListStaffs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).ListStaffs(ctx, req.(*ListStaffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.staff.v1.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStaff",
			Handler:    _StaffService_GetStaff_Handler,
		},
		{
			MethodName: "ListStaffs",
			Handler:    _StaffService_ListStaffs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/staff/v1/staff_service.proto",
}
