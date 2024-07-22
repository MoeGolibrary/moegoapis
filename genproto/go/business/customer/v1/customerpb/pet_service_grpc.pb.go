// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: moego/business/customer/v1/pet_service.proto

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
	PetService_CreatePet_FullMethodName     = "/moego.business.customer.v1.PetService/CreatePet"
	PetService_CreatePetCode_FullMethodName = "/moego.business.customer.v1.PetService/CreatePetCode"
	PetService_ListPetCodes_FullMethodName  = "/moego.business.customer.v1.PetService/ListPetCodes"
	PetService_CreatePetNote_FullMethodName = "/moego.business.customer.v1.PetService/CreatePetNote"
	PetService_ListPetNotes_FullMethodName  = "/moego.business.customer.v1.PetService/ListPetNotes"
)

// PetServiceClient is the client API for PetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetServiceClient interface {
	// Create a pet
	CreatePet(ctx context.Context, in *CreatePetRequest, opts ...grpc.CallOption) (*CreatePetResponse, error)
	// Create a pet code
	CreatePetCode(ctx context.Context, in *CreatePetCodeRequest, opts ...grpc.CallOption) (*CreatePetCodeResponse, error)
	// List Pet Codes
	ListPetCodes(ctx context.Context, in *ListPetCodeRequest, opts ...grpc.CallOption) (*ListPetCodeResponse, error)
	// Create a pet note
	CreatePetNote(ctx context.Context, in *CreatePetNoteRequest, opts ...grpc.CallOption) (*CreatePetNoteResponse, error)
	// List Pet Notes
	ListPetNotes(ctx context.Context, in *ListPetNoteRequest, opts ...grpc.CallOption) (*ListPetNoteResponse, error)
}

type petServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetServiceClient(cc grpc.ClientConnInterface) PetServiceClient {
	return &petServiceClient{cc}
}

func (c *petServiceClient) CreatePet(ctx context.Context, in *CreatePetRequest, opts ...grpc.CallOption) (*CreatePetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePetResponse)
	err := c.cc.Invoke(ctx, PetService_CreatePet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) CreatePetCode(ctx context.Context, in *CreatePetCodeRequest, opts ...grpc.CallOption) (*CreatePetCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePetCodeResponse)
	err := c.cc.Invoke(ctx, PetService_CreatePetCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) ListPetCodes(ctx context.Context, in *ListPetCodeRequest, opts ...grpc.CallOption) (*ListPetCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPetCodeResponse)
	err := c.cc.Invoke(ctx, PetService_ListPetCodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) CreatePetNote(ctx context.Context, in *CreatePetNoteRequest, opts ...grpc.CallOption) (*CreatePetNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePetNoteResponse)
	err := c.cc.Invoke(ctx, PetService_CreatePetNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) ListPetNotes(ctx context.Context, in *ListPetNoteRequest, opts ...grpc.CallOption) (*ListPetNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPetNoteResponse)
	err := c.cc.Invoke(ctx, PetService_ListPetNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetServiceServer is the server API for PetService service.
// All implementations must embed UnimplementedPetServiceServer
// for forward compatibility
type PetServiceServer interface {
	// Create a pet
	CreatePet(context.Context, *CreatePetRequest) (*CreatePetResponse, error)
	// Create a pet code
	CreatePetCode(context.Context, *CreatePetCodeRequest) (*CreatePetCodeResponse, error)
	// List Pet Codes
	ListPetCodes(context.Context, *ListPetCodeRequest) (*ListPetCodeResponse, error)
	// Create a pet note
	CreatePetNote(context.Context, *CreatePetNoteRequest) (*CreatePetNoteResponse, error)
	// List Pet Notes
	ListPetNotes(context.Context, *ListPetNoteRequest) (*ListPetNoteResponse, error)
	mustEmbedUnimplementedPetServiceServer()
}

// UnimplementedPetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPetServiceServer struct {
}

func (UnimplementedPetServiceServer) CreatePet(context.Context, *CreatePetRequest) (*CreatePetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePet not implemented")
}
func (UnimplementedPetServiceServer) CreatePetCode(context.Context, *CreatePetCodeRequest) (*CreatePetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetCode not implemented")
}
func (UnimplementedPetServiceServer) ListPetCodes(context.Context, *ListPetCodeRequest) (*ListPetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPetCodes not implemented")
}
func (UnimplementedPetServiceServer) CreatePetNote(context.Context, *CreatePetNoteRequest) (*CreatePetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetNote not implemented")
}
func (UnimplementedPetServiceServer) ListPetNotes(context.Context, *ListPetNoteRequest) (*ListPetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPetNotes not implemented")
}
func (UnimplementedPetServiceServer) mustEmbedUnimplementedPetServiceServer() {}

// UnsafePetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetServiceServer will
// result in compilation errors.
type UnsafePetServiceServer interface {
	mustEmbedUnimplementedPetServiceServer()
}

func RegisterPetServiceServer(s grpc.ServiceRegistrar, srv PetServiceServer) {
	s.RegisterService(&PetService_ServiceDesc, srv)
}

func _PetService_CreatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).CreatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_CreatePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).CreatePet(ctx, req.(*CreatePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_CreatePetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).CreatePetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_CreatePetCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).CreatePetCode(ctx, req.(*CreatePetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_ListPetCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).ListPetCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_ListPetCodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).ListPetCodes(ctx, req.(*ListPetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_CreatePetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).CreatePetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_CreatePetNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).CreatePetNote(ctx, req.(*CreatePetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_ListPetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).ListPetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_ListPetNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).ListPetNotes(ctx, req.(*ListPetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetService_ServiceDesc is the grpc.ServiceDesc for PetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.customer.v1.PetService",
	HandlerType: (*PetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePet",
			Handler:    _PetService_CreatePet_Handler,
		},
		{
			MethodName: "CreatePetCode",
			Handler:    _PetService_CreatePetCode_Handler,
		},
		{
			MethodName: "ListPetCodes",
			Handler:    _PetService_ListPetCodes_Handler,
		},
		{
			MethodName: "CreatePetNote",
			Handler:    _PetService_CreatePetNote_Handler,
		},
		{
			MethodName: "ListPetNotes",
			Handler:    _PetService_ListPetNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/customer/v1/pet_service.proto",
}

const (
	PetCodeService_CreatePetCode_FullMethodName = "/moego.business.customer.v1.PetCodeService/CreatePetCode"
	PetCodeService_ListPetCodes_FullMethodName  = "/moego.business.customer.v1.PetCodeService/ListPetCodes"
)

// PetCodeServiceClient is the client API for PetCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetCodeServiceClient interface {
	// Create a petCode
	CreatePetCode(ctx context.Context, in *CreatePetCodeRequest, opts ...grpc.CallOption) (*CreatePetCodeResponse, error)
	// ListPetCodes
	ListPetCodes(ctx context.Context, in *ListPetCodeRequest, opts ...grpc.CallOption) (*ListPetCodeResponse, error)
}

type petCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetCodeServiceClient(cc grpc.ClientConnInterface) PetCodeServiceClient {
	return &petCodeServiceClient{cc}
}

func (c *petCodeServiceClient) CreatePetCode(ctx context.Context, in *CreatePetCodeRequest, opts ...grpc.CallOption) (*CreatePetCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePetCodeResponse)
	err := c.cc.Invoke(ctx, PetCodeService_CreatePetCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petCodeServiceClient) ListPetCodes(ctx context.Context, in *ListPetCodeRequest, opts ...grpc.CallOption) (*ListPetCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPetCodeResponse)
	err := c.cc.Invoke(ctx, PetCodeService_ListPetCodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetCodeServiceServer is the server API for PetCodeService service.
// All implementations must embed UnimplementedPetCodeServiceServer
// for forward compatibility
type PetCodeServiceServer interface {
	// Create a petCode
	CreatePetCode(context.Context, *CreatePetCodeRequest) (*CreatePetCodeResponse, error)
	// ListPetCodes
	ListPetCodes(context.Context, *ListPetCodeRequest) (*ListPetCodeResponse, error)
	mustEmbedUnimplementedPetCodeServiceServer()
}

// UnimplementedPetCodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPetCodeServiceServer struct {
}

func (UnimplementedPetCodeServiceServer) CreatePetCode(context.Context, *CreatePetCodeRequest) (*CreatePetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetCode not implemented")
}
func (UnimplementedPetCodeServiceServer) ListPetCodes(context.Context, *ListPetCodeRequest) (*ListPetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPetCodes not implemented")
}
func (UnimplementedPetCodeServiceServer) mustEmbedUnimplementedPetCodeServiceServer() {}

// UnsafePetCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetCodeServiceServer will
// result in compilation errors.
type UnsafePetCodeServiceServer interface {
	mustEmbedUnimplementedPetCodeServiceServer()
}

func RegisterPetCodeServiceServer(s grpc.ServiceRegistrar, srv PetCodeServiceServer) {
	s.RegisterService(&PetCodeService_ServiceDesc, srv)
}

func _PetCodeService_CreatePetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCodeServiceServer).CreatePetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetCodeService_CreatePetCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCodeServiceServer).CreatePetCode(ctx, req.(*CreatePetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetCodeService_ListPetCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCodeServiceServer).ListPetCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetCodeService_ListPetCodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCodeServiceServer).ListPetCodes(ctx, req.(*ListPetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetCodeService_ServiceDesc is the grpc.ServiceDesc for PetCodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetCodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.customer.v1.PetCodeService",
	HandlerType: (*PetCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePetCode",
			Handler:    _PetCodeService_CreatePetCode_Handler,
		},
		{
			MethodName: "ListPetCodes",
			Handler:    _PetCodeService_ListPetCodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/customer/v1/pet_service.proto",
}

const (
	PetNoteService_CreatePetNote_FullMethodName = "/moego.business.customer.v1.PetNoteService/CreatePetNote"
	PetNoteService_ListPetNotes_FullMethodName  = "/moego.business.customer.v1.PetNoteService/ListPetNotes"
)

// PetNoteServiceClient is the client API for PetNoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetNoteServiceClient interface {
	CreatePetNote(ctx context.Context, in *CreatePetNoteRequest, opts ...grpc.CallOption) (*CreatePetNoteResponse, error)
	ListPetNotes(ctx context.Context, in *ListPetNoteRequest, opts ...grpc.CallOption) (*ListPetNoteResponse, error)
}

type petNoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetNoteServiceClient(cc grpc.ClientConnInterface) PetNoteServiceClient {
	return &petNoteServiceClient{cc}
}

func (c *petNoteServiceClient) CreatePetNote(ctx context.Context, in *CreatePetNoteRequest, opts ...grpc.CallOption) (*CreatePetNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePetNoteResponse)
	err := c.cc.Invoke(ctx, PetNoteService_CreatePetNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petNoteServiceClient) ListPetNotes(ctx context.Context, in *ListPetNoteRequest, opts ...grpc.CallOption) (*ListPetNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPetNoteResponse)
	err := c.cc.Invoke(ctx, PetNoteService_ListPetNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetNoteServiceServer is the server API for PetNoteService service.
// All implementations must embed UnimplementedPetNoteServiceServer
// for forward compatibility
type PetNoteServiceServer interface {
	CreatePetNote(context.Context, *CreatePetNoteRequest) (*CreatePetNoteResponse, error)
	ListPetNotes(context.Context, *ListPetNoteRequest) (*ListPetNoteResponse, error)
	mustEmbedUnimplementedPetNoteServiceServer()
}

// UnimplementedPetNoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPetNoteServiceServer struct {
}

func (UnimplementedPetNoteServiceServer) CreatePetNote(context.Context, *CreatePetNoteRequest) (*CreatePetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetNote not implemented")
}
func (UnimplementedPetNoteServiceServer) ListPetNotes(context.Context, *ListPetNoteRequest) (*ListPetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPetNotes not implemented")
}
func (UnimplementedPetNoteServiceServer) mustEmbedUnimplementedPetNoteServiceServer() {}

// UnsafePetNoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetNoteServiceServer will
// result in compilation errors.
type UnsafePetNoteServiceServer interface {
	mustEmbedUnimplementedPetNoteServiceServer()
}

func RegisterPetNoteServiceServer(s grpc.ServiceRegistrar, srv PetNoteServiceServer) {
	s.RegisterService(&PetNoteService_ServiceDesc, srv)
}

func _PetNoteService_CreatePetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetNoteServiceServer).CreatePetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetNoteService_CreatePetNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetNoteServiceServer).CreatePetNote(ctx, req.(*CreatePetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetNoteService_ListPetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetNoteServiceServer).ListPetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetNoteService_ListPetNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetNoteServiceServer).ListPetNotes(ctx, req.(*ListPetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetNoteService_ServiceDesc is the grpc.ServiceDesc for PetNoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetNoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moego.business.customer.v1.PetNoteService",
	HandlerType: (*PetNoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePetNote",
			Handler:    _PetNoteService_CreatePetNote_Handler,
		},
		{
			MethodName: "ListPetNotes",
			Handler:    _PetNoteService_ListPetNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moego/business/customer/v1/pet_service.proto",
}
