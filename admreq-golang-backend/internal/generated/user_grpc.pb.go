// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: user.proto

package tsr_v1

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
	UserService_RegisterUser_FullMethodName       = "/tsr.v1.UserService/RegisterUser"
	UserService_UserAuth_FullMethodName           = "/tsr.v1.UserService/UserAuth"
	UserService_DisableUser_FullMethodName        = "/tsr.v1.UserService/DisableUser"
	UserService_GetUsers_FullMethodName           = "/tsr.v1.UserService/GetUsers"
	UserService_AddDepartment_FullMethodName      = "/tsr.v1.UserService/AddDepartment"
	UserService_GetDepartments_FullMethodName     = "/tsr.v1.UserService/GetDepartments"
	UserService_ChangeUserPassword_FullMethodName = "/tsr.v1.UserService/ChangeUserPassword"
	UserService_UpdateUserRight_FullMethodName    = "/tsr.v1.UserService/UpdateUserRight"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	UserAuth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserAuthResponse, error)
	DisableUser(ctx context.Context, in *DisableUserRequest, opts ...grpc.CallOption) (*DisableUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	AddDepartment(ctx context.Context, in *AddDepartmentRequest, opts ...grpc.CallOption) (*AddDepartmentResponse, error)
	GetDepartments(ctx context.Context, in *GetDepartmentsRequest, opts ...grpc.CallOption) (*GetDepartmentsResponse, error)
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error)
	UpdateUserRight(ctx context.Context, in *UpdateUserRightRequest, opts ...grpc.CallOption) (*UpdateUserRightResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, UserService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserAuth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserAuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserAuthResponse)
	err := c.cc.Invoke(ctx, UserService_UserAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DisableUser(ctx context.Context, in *DisableUserRequest, opts ...grpc.CallOption) (*DisableUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisableUserResponse)
	err := c.cc.Invoke(ctx, UserService_DisableUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddDepartment(ctx context.Context, in *AddDepartmentRequest, opts ...grpc.CallOption) (*AddDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDepartmentResponse)
	err := c.cc.Invoke(ctx, UserService_AddDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetDepartments(ctx context.Context, in *GetDepartmentsRequest, opts ...grpc.CallOption) (*GetDepartmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDepartmentsResponse)
	err := c.cc.Invoke(ctx, UserService_GetDepartments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeUserPasswordResponse)
	err := c.cc.Invoke(ctx, UserService_ChangeUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserRight(ctx context.Context, in *UpdateUserRightRequest, opts ...grpc.CallOption) (*UpdateUserRightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserRightResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserRight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	UserAuth(context.Context, *UserAuthRequest) (*UserAuthResponse, error)
	DisableUser(context.Context, *DisableUserRequest) (*DisableUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	AddDepartment(context.Context, *AddDepartmentRequest) (*AddDepartmentResponse, error)
	GetDepartments(context.Context, *GetDepartmentsRequest) (*GetDepartmentsResponse, error)
	ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error)
	UpdateUserRight(context.Context, *UpdateUserRightRequest) (*UpdateUserRightResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) UserAuth(context.Context, *UserAuthRequest) (*UserAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuth not implemented")
}
func (UnimplementedUserServiceServer) DisableUser(context.Context, *DisableUserRequest) (*DisableUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableUser not implemented")
}
func (UnimplementedUserServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServiceServer) AddDepartment(context.Context, *AddDepartmentRequest) (*AddDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDepartment not implemented")
}
func (UnimplementedUserServiceServer) GetDepartments(context.Context, *GetDepartmentsRequest) (*GetDepartmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartments not implemented")
}
func (UnimplementedUserServiceServer) ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserRight(context.Context, *UpdateUserRightRequest) (*UpdateUserRightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRight not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserAuth(ctx, req.(*UserAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DisableUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DisableUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DisableUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DisableUser(ctx, req.(*DisableUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddDepartment(ctx, req.(*AddDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetDepartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetDepartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetDepartments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetDepartments(ctx, req.(*GetDepartmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangeUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserRight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserRight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserRight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserRight(ctx, req.(*UpdateUserRightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tsr.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "UserAuth",
			Handler:    _UserService_UserAuth_Handler,
		},
		{
			MethodName: "DisableUser",
			Handler:    _UserService_DisableUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "AddDepartment",
			Handler:    _UserService_AddDepartment_Handler,
		},
		{
			MethodName: "GetDepartments",
			Handler:    _UserService_GetDepartments_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _UserService_ChangeUserPassword_Handler,
		},
		{
			MethodName: "UpdateUserRight",
			Handler:    _UserService_UpdateUserRight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
