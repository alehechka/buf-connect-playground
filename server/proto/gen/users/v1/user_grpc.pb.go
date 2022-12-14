// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: users/v1/user.proto

package usersv1

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

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UsersService_ListUsersClient, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GenerateUsers(ctx context.Context, in *GenerateUsersRequest, opts ...grpc.CallOption) (*GenerateUsersResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	DeleteAllUsers(ctx context.Context, in *DeleteAllUsersRequest, opts ...grpc.CallOption) (*DeleteAllUsersResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UsersService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UsersService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UsersService_ServiceDesc.Streams[0], "/users.v1.UsersService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UsersService_ListUsersClient interface {
	Recv() (*ListUsersResponse, error)
	grpc.ClientStream
}

type usersServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *usersServiceListUsersClient) Recv() (*ListUsersResponse, error) {
	m := new(ListUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UsersService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GenerateUsers(ctx context.Context, in *GenerateUsersRequest, opts ...grpc.CallOption) (*GenerateUsersResponse, error) {
	out := new(GenerateUsersResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UsersService/GenerateUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UsersService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UsersService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) DeleteAllUsers(ctx context.Context, in *DeleteAllUsersRequest, opts ...grpc.CallOption) (*DeleteAllUsersResponse, error) {
	out := new(DeleteAllUsersResponse)
	err := c.cc.Invoke(ctx, "/users.v1.UsersService/DeleteAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUsers(*ListUsersRequest, UsersService_ListUsersServer) error
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GenerateUsers(context.Context, *GenerateUsersRequest) (*GenerateUsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	DeleteAllUsers(context.Context, *DeleteAllUsersRequest) (*DeleteAllUsersResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsersServiceServer) ListUsers(*ListUsersRequest, UsersService_ListUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUsersServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServiceServer) GenerateUsers(context.Context, *GenerateUsersRequest) (*GenerateUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateUsers not implemented")
}
func (UnimplementedUsersServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersServiceServer) DeleteAllUsers(context.Context, *DeleteAllUsersRequest) (*DeleteAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllUsers not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UsersService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServiceServer).ListUsers(m, &usersServiceListUsersServer{stream})
}

type UsersService_ListUsersServer interface {
	Send(*ListUsersResponse) error
	grpc.ServerStream
}

type usersServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *usersServiceListUsersServer) Send(m *ListUsersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UsersService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UsersService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GenerateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GenerateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UsersService/GenerateUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GenerateUsers(ctx, req.(*GenerateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UsersService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UsersService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_DeleteAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.v1.UsersService/DeleteAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteAllUsers(ctx, req.(*DeleteAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.v1.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UsersService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UsersService_CreateUser_Handler,
		},
		{
			MethodName: "GenerateUsers",
			Handler:    _UsersService_GenerateUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UsersService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UsersService_DeleteUser_Handler,
		},
		{
			MethodName: "DeleteAllUsers",
			Handler:    _UsersService_DeleteAllUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _UsersService_ListUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/v1/user.proto",
}
