// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: moego/business/appointment/v1/appointment_service.proto

package appointmentpb

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
	AppointmentService_GetAppointment_FullMethodName    = "/moego.business.appointment.v1.AppointmentService/GetAppointment"
	AppointmentService_ListAppointments_FullMethodName  = "/moego.business.appointment.v1.AppointmentService/ListAppointments"
	AppointmentService_CreateAppointment_FullMethodName = "/moego.business.appointment.v1.AppointmentService/CreateAppointment"
)

// AppointmentServiceClient is the client API for AppointmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AppointmentService openapi definitions for operate appointments
type AppointmentServiceClient interface {
	// GetAppointment
	GetAppointment(ctx context.Context, in *GetAppointmentRequest, opts ...grpc.CallOption) (*Appointment, error)
	// ListStaffs
	ListAppointments(ctx context.Context, in *ListAppointmentsRequest, opts ...grpc.CallOption) (*ListAppointmentsResponse, error)
	// CreateAppointment
	CreateAppointment(ctx context.Context, in *CreateAppointmentRequest, opts ...grpc.CallOption) (*Appointment, error)
}

type appointmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentServiceClient(cc grpc.ClientConnInterface) AppointmentServiceClient {
	return &appointmentServiceClient{cc}
}

func (c *appointmentServiceClient) GetAppointment(ctx context.Context, in *GetAppointmentRequest, opts ...grpc.CallOption) (*Appointment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Appointment)
	err := c.cc.Invoke(ctx, AppointmentService_GetAppointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) ListAppointments(ctx context.Context, in *ListAppointmentsRequest, opts ...grpc.CallOption) (*ListAppointmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAppointmentsResponse)
	err := c.cc.Invoke(ctx, AppointmentService_ListAppointments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) CreateAppointment(ctx context.Context, in *CreateAppointmentRequest, opts ...grpc.CallOption) (*Appointment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Appointment)
	err := c.cc.Invoke(ctx, AppointmentService_CreateAppointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServiceServer is the server API for AppointmentService service.
// All implementations must embed UnimplementedAppointmentServiceServer
// for forward compatibility.
//
// AppointmentService openapi definitions for operate appointments
type AppointmentServiceServer interface {
	// GetAppointment
	GetAppointment(context.Context, *GetAppointmentRequest) (*Appointment, error)
	// ListStaffs
	ListAppointments(context.Context, *ListAppointmentsRequest) (*ListAppointmentsResponse, error)
	// CreateAppointment
	CreateAppointment(context.Context, *CreateAppointmentRequest) (*Appointment, error)
	mustEmbedUnimplementedAppointmentServiceServer()
}

// UnimplementedAppointmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppointmentServiceServer struct{}

func (UnimplementedAppointmentServiceServer) GetAppointment(context.Context, *GetAppointmentRequest) (*Appointment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppointment not implemented")
}
func (UnimplementedAppointmentServiceServer) ListAppointments(context.Context, *ListAppointmentsRequest) (*ListAppointmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppointments not implemented")
}
func (UnimplementedAppointmentServiceServer) CreateAppointment(context.Context, *CreateAppointmentRequest) (*Appointment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppointment not implemented")
}
func (UnimplementedAppointmentServiceServer) mustEmbedUnimplementedAppointmentServiceServer() {}
func (UnimplementedAppointmentServiceServer) testEmbeddedByValue()                            {}

// UnsafeAppointmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServiceServer will
// result in compilation errors.
type UnsafeAppointmentServiceServer interface {
	mustEmbedUnimplementedAppointmentServiceServer()
}

func RegisterAppointmentServiceServer(s grpc.ServiceRegistrar, srv AppointmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedAppointmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppointmentService_ServiceDesc, srv)
}

func _AppointmentService_GetAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetAppointment(ctx, req.(*GetAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_ListAppointments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppointmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).ListAppointments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_ListAppointments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).ListAppointments(ctx, req.(*ListAppointmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_CreateAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).CreateAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_CreateAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).CreateAppointment(ctx, req.(*CreateAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppointmentService_ServiceDesc is the grpc.ServiceDesc for AppointmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppointmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.appointment.v1.AppointmentService",
	HandlerType: (*AppointmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppointment",
			Handler:    _AppointmentService_GetAppointment_Handler,
		},
		{
			MethodName: "ListAppointments",
			Handler:    _AppointmentService_ListAppointments_Handler,
		},
		{
			MethodName: "CreateAppointment",
			Handler:    _AppointmentService_CreateAppointment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/appointment/v1/appointment_service.proto",
}
