// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: auth/v1/client_api.proto

package authv1

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
	ClientAPI_CreateClient_FullMethodName     = "/auth.v1.ClientAPI/CreateClient"
	ClientAPI_GetClientByID_FullMethodName    = "/auth.v1.ClientAPI/GetClientByID"
	ClientAPI_UpdateClientByID_FullMethodName = "/auth.v1.ClientAPI/UpdateClientByID"
	ClientAPI_DeleteClientByID_FullMethodName = "/auth.v1.ClientAPI/DeleteClientByID"
	ClientAPI_ListClients_FullMethodName      = "/auth.v1.ClientAPI/ListClients"
)

// ClientAPIClient is the client API for ClientAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ClientAPI service definition
type ClientAPIClient interface {
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
	GetClientByID(ctx context.Context, in *GetClientByIDRequest, opts ...grpc.CallOption) (*GetClientByIDResponse, error)
	UpdateClientByID(ctx context.Context, in *UpdateClientByIDRequest, opts ...grpc.CallOption) (*UpdateClientByIDResponse, error)
	DeleteClientByID(ctx context.Context, in *DeleteClientByIDRequest, opts ...grpc.CallOption) (*DeleteClientByIDResponse, error)
	ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*ListClientsResponse, error)
}

type clientAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewClientAPIClient(cc grpc.ClientConnInterface) ClientAPIClient {
	return &clientAPIClient{cc}
}

func (c *clientAPIClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, ClientAPI_CreateClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAPIClient) GetClientByID(ctx context.Context, in *GetClientByIDRequest, opts ...grpc.CallOption) (*GetClientByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClientByIDResponse)
	err := c.cc.Invoke(ctx, ClientAPI_GetClientByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAPIClient) UpdateClientByID(ctx context.Context, in *UpdateClientByIDRequest, opts ...grpc.CallOption) (*UpdateClientByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateClientByIDResponse)
	err := c.cc.Invoke(ctx, ClientAPI_UpdateClientByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAPIClient) DeleteClientByID(ctx context.Context, in *DeleteClientByIDRequest, opts ...grpc.CallOption) (*DeleteClientByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteClientByIDResponse)
	err := c.cc.Invoke(ctx, ClientAPI_DeleteClientByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAPIClient) ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*ListClientsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListClientsResponse)
	err := c.cc.Invoke(ctx, ClientAPI_ListClients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAPIServer is the server API for ClientAPI service.
// All implementations should embed UnimplementedClientAPIServer
// for forward compatibility.
//
// ClientAPI service definition
type ClientAPIServer interface {
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
	GetClientByID(context.Context, *GetClientByIDRequest) (*GetClientByIDResponse, error)
	UpdateClientByID(context.Context, *UpdateClientByIDRequest) (*UpdateClientByIDResponse, error)
	DeleteClientByID(context.Context, *DeleteClientByIDRequest) (*DeleteClientByIDResponse, error)
	ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error)
}

// UnimplementedClientAPIServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientAPIServer struct{}

func (UnimplementedClientAPIServer) CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedClientAPIServer) GetClientByID(context.Context, *GetClientByIDRequest) (*GetClientByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientByID not implemented")
}
func (UnimplementedClientAPIServer) UpdateClientByID(context.Context, *UpdateClientByIDRequest) (*UpdateClientByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientByID not implemented")
}
func (UnimplementedClientAPIServer) DeleteClientByID(context.Context, *DeleteClientByIDRequest) (*DeleteClientByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientByID not implemented")
}
func (UnimplementedClientAPIServer) ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClients not implemented")
}
func (UnimplementedClientAPIServer) testEmbeddedByValue() {}

// UnsafeClientAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientAPIServer will
// result in compilation errors.
type UnsafeClientAPIServer interface {
	mustEmbedUnimplementedClientAPIServer()
}

func RegisterClientAPIServer(s grpc.ServiceRegistrar, srv ClientAPIServer) {
	// If the following call pancis, it indicates UnimplementedClientAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientAPI_ServiceDesc, srv)
}

func _ClientAPI_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAPI_CreateClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAPI_GetClientByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServer).GetClientByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAPI_GetClientByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServer).GetClientByID(ctx, req.(*GetClientByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAPI_UpdateClientByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServer).UpdateClientByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAPI_UpdateClientByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServer).UpdateClientByID(ctx, req.(*UpdateClientByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAPI_DeleteClientByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServer).DeleteClientByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAPI_DeleteClientByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServer).DeleteClientByID(ctx, req.(*DeleteClientByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAPI_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAPIServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAPI_ListClients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAPIServer).ListClients(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientAPI_ServiceDesc is the grpc.ServiceDesc for ClientAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.ClientAPI",
	HandlerType: (*ClientAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientAPI_CreateClient_Handler,
		},
		{
			MethodName: "GetClientByID",
			Handler:    _ClientAPI_GetClientByID_Handler,
		},
		{
			MethodName: "UpdateClientByID",
			Handler:    _ClientAPI_UpdateClientByID_Handler,
		},
		{
			MethodName: "DeleteClientByID",
			Handler:    _ClientAPI_DeleteClientByID_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _ClientAPI_ListClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/client_api.proto",
}
