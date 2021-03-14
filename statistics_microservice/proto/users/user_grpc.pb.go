// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package users

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AddBook(ctx context.Context, in *UserAddBookRequest, opts ...grpc.CallOption) (*UserAddBookResponse, error)
	GetAllBooks(ctx context.Context, in *GetAllBooksRequest, opts ...grpc.CallOption) (*GetAllBooksResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddBook(ctx context.Context, in *UserAddBookRequest, opts ...grpc.CallOption) (*UserAddBookResponse, error) {
	out := new(UserAddBookResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllBooks(ctx context.Context, in *GetAllBooksRequest, opts ...grpc.CallOption) (*GetAllBooksResponse, error) {
	out := new(GetAllBooksResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetAllBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	AddBook(context.Context, *UserAddBookRequest) (*UserAddBookResponse, error)
	GetAllBooks(context.Context, *GetAllBooksRequest) (*GetAllBooksResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) AddBook(context.Context, *UserAddBookRequest) (*UserAddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedUserServiceServer) GetAllBooks(context.Context, *GetAllBooksRequest) (*GetAllBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddBook(ctx, req.(*UserAddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetAllBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllBooks(ctx, req.(*GetAllBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _UserService_AddBook_Handler,
		},
		{
			MethodName: "GetAllBooks",
			Handler:    _UserService_GetAllBooks_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users/user.proto",
}
