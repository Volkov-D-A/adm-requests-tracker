// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: tsr.proto

package tsr_v1

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
	TSRService_CreateTSR_FullMethodName      = "/tsr.v1.TSRService/CreateTSR"
	TSRService_EmployeeTSR_FullMethodName    = "/tsr.v1.TSRService/EmployeeTSR"
	TSRService_ImportanceTSR_FullMethodName  = "/tsr.v1.TSRService/ImportanceTSR"
	TSRService_FinishTSR_FullMethodName      = "/tsr.v1.TSRService/FinishTSR"
	TSRService_ApplyTSR_FullMethodName       = "/tsr.v1.TSRService/ApplyTSR"
	TSRService_GetFullTsrInfo_FullMethodName = "/tsr.v1.TSRService/GetFullTsrInfo"
	TSRService_GetListTickets_FullMethodName = "/tsr.v1.TSRService/GetListTickets"
	TSRService_SetTsrComment_FullMethodName  = "/tsr.v1.TSRService/SetTsrComment"
	TSRService_GetTsrCommnts_FullMethodName  = "/tsr.v1.TSRService/GetTsrCommnts"
	TSRService_GetTsrStat_FullMethodName     = "/tsr.v1.TSRService/GetTsrStat"
)

// TSRServiceClient is the client API for TSRService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TSRServiceClient interface {
	CreateTSR(ctx context.Context, in *CreateTSRRequest, opts ...grpc.CallOption) (*CreateTSRResponse, error)
	EmployeeTSR(ctx context.Context, in *EmployeeTSRRequest, opts ...grpc.CallOption) (*EmployeeTSRResponse, error)
	ImportanceTSR(ctx context.Context, in *ImportanceTSRRequest, opts ...grpc.CallOption) (*ImportanceTSRResponse, error)
	FinishTSR(ctx context.Context, in *FinishTSRRequest, opts ...grpc.CallOption) (*FinishTSRResponse, error)
	ApplyTSR(ctx context.Context, in *ApplyTSRRequest, opts ...grpc.CallOption) (*ApplyTSRResponse, error)
	GetFullTsrInfo(ctx context.Context, in *GetFullTsrInfoRequest, opts ...grpc.CallOption) (*GetFullTsrInfoResponse, error)
	GetListTickets(ctx context.Context, in *GetListTicketRequest, opts ...grpc.CallOption) (*GetListTicketResponse, error)
	SetTsrComment(ctx context.Context, in *SetTsrCommentRequest, opts ...grpc.CallOption) (*SetTsrCommentResponse, error)
	GetTsrCommnts(ctx context.Context, in *GetTsrCommentsRequest, opts ...grpc.CallOption) (*GetTsrCommentsResponse, error)
	GetTsrStat(ctx context.Context, in *GetTsrStatRequest, opts ...grpc.CallOption) (*GetTsrStatResponse, error)
}

type tSRServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTSRServiceClient(cc grpc.ClientConnInterface) TSRServiceClient {
	return &tSRServiceClient{cc}
}

func (c *tSRServiceClient) CreateTSR(ctx context.Context, in *CreateTSRRequest, opts ...grpc.CallOption) (*CreateTSRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTSRResponse)
	err := c.cc.Invoke(ctx, TSRService_CreateTSR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) EmployeeTSR(ctx context.Context, in *EmployeeTSRRequest, opts ...grpc.CallOption) (*EmployeeTSRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployeeTSRResponse)
	err := c.cc.Invoke(ctx, TSRService_EmployeeTSR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) ImportanceTSR(ctx context.Context, in *ImportanceTSRRequest, opts ...grpc.CallOption) (*ImportanceTSRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportanceTSRResponse)
	err := c.cc.Invoke(ctx, TSRService_ImportanceTSR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) FinishTSR(ctx context.Context, in *FinishTSRRequest, opts ...grpc.CallOption) (*FinishTSRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FinishTSRResponse)
	err := c.cc.Invoke(ctx, TSRService_FinishTSR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) ApplyTSR(ctx context.Context, in *ApplyTSRRequest, opts ...grpc.CallOption) (*ApplyTSRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApplyTSRResponse)
	err := c.cc.Invoke(ctx, TSRService_ApplyTSR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) GetFullTsrInfo(ctx context.Context, in *GetFullTsrInfoRequest, opts ...grpc.CallOption) (*GetFullTsrInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFullTsrInfoResponse)
	err := c.cc.Invoke(ctx, TSRService_GetFullTsrInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) GetListTickets(ctx context.Context, in *GetListTicketRequest, opts ...grpc.CallOption) (*GetListTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListTicketResponse)
	err := c.cc.Invoke(ctx, TSRService_GetListTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) SetTsrComment(ctx context.Context, in *SetTsrCommentRequest, opts ...grpc.CallOption) (*SetTsrCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetTsrCommentResponse)
	err := c.cc.Invoke(ctx, TSRService_SetTsrComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) GetTsrCommnts(ctx context.Context, in *GetTsrCommentsRequest, opts ...grpc.CallOption) (*GetTsrCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTsrCommentsResponse)
	err := c.cc.Invoke(ctx, TSRService_GetTsrCommnts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tSRServiceClient) GetTsrStat(ctx context.Context, in *GetTsrStatRequest, opts ...grpc.CallOption) (*GetTsrStatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTsrStatResponse)
	err := c.cc.Invoke(ctx, TSRService_GetTsrStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TSRServiceServer is the server API for TSRService service.
// All implementations must embed UnimplementedTSRServiceServer
// for forward compatibility
type TSRServiceServer interface {
	CreateTSR(context.Context, *CreateTSRRequest) (*CreateTSRResponse, error)
	EmployeeTSR(context.Context, *EmployeeTSRRequest) (*EmployeeTSRResponse, error)
	ImportanceTSR(context.Context, *ImportanceTSRRequest) (*ImportanceTSRResponse, error)
	FinishTSR(context.Context, *FinishTSRRequest) (*FinishTSRResponse, error)
	ApplyTSR(context.Context, *ApplyTSRRequest) (*ApplyTSRResponse, error)
	GetFullTsrInfo(context.Context, *GetFullTsrInfoRequest) (*GetFullTsrInfoResponse, error)
	GetListTickets(context.Context, *GetListTicketRequest) (*GetListTicketResponse, error)
	SetTsrComment(context.Context, *SetTsrCommentRequest) (*SetTsrCommentResponse, error)
	GetTsrCommnts(context.Context, *GetTsrCommentsRequest) (*GetTsrCommentsResponse, error)
	GetTsrStat(context.Context, *GetTsrStatRequest) (*GetTsrStatResponse, error)
	mustEmbedUnimplementedTSRServiceServer()
}

// UnimplementedTSRServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTSRServiceServer struct {
}

func (UnimplementedTSRServiceServer) CreateTSR(context.Context, *CreateTSRRequest) (*CreateTSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTSR not implemented")
}
func (UnimplementedTSRServiceServer) EmployeeTSR(context.Context, *EmployeeTSRRequest) (*EmployeeTSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmployeeTSR not implemented")
}
func (UnimplementedTSRServiceServer) ImportanceTSR(context.Context, *ImportanceTSRRequest) (*ImportanceTSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportanceTSR not implemented")
}
func (UnimplementedTSRServiceServer) FinishTSR(context.Context, *FinishTSRRequest) (*FinishTSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTSR not implemented")
}
func (UnimplementedTSRServiceServer) ApplyTSR(context.Context, *ApplyTSRRequest) (*ApplyTSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyTSR not implemented")
}
func (UnimplementedTSRServiceServer) GetFullTsrInfo(context.Context, *GetFullTsrInfoRequest) (*GetFullTsrInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullTsrInfo not implemented")
}
func (UnimplementedTSRServiceServer) GetListTickets(context.Context, *GetListTicketRequest) (*GetListTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTickets not implemented")
}
func (UnimplementedTSRServiceServer) SetTsrComment(context.Context, *SetTsrCommentRequest) (*SetTsrCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTsrComment not implemented")
}
func (UnimplementedTSRServiceServer) GetTsrCommnts(context.Context, *GetTsrCommentsRequest) (*GetTsrCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTsrCommnts not implemented")
}
func (UnimplementedTSRServiceServer) GetTsrStat(context.Context, *GetTsrStatRequest) (*GetTsrStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTsrStat not implemented")
}
func (UnimplementedTSRServiceServer) mustEmbedUnimplementedTSRServiceServer() {}

// UnsafeTSRServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TSRServiceServer will
// result in compilation errors.
type UnsafeTSRServiceServer interface {
	mustEmbedUnimplementedTSRServiceServer()
}

func RegisterTSRServiceServer(s grpc.ServiceRegistrar, srv TSRServiceServer) {
	s.RegisterService(&TSRService_ServiceDesc, srv)
}

func _TSRService_CreateTSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).CreateTSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_CreateTSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).CreateTSR(ctx, req.(*CreateTSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_EmployeeTSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeTSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).EmployeeTSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_EmployeeTSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).EmployeeTSR(ctx, req.(*EmployeeTSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_ImportanceTSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportanceTSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).ImportanceTSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_ImportanceTSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).ImportanceTSR(ctx, req.(*ImportanceTSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_FinishTSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishTSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).FinishTSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_FinishTSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).FinishTSR(ctx, req.(*FinishTSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_ApplyTSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyTSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).ApplyTSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_ApplyTSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).ApplyTSR(ctx, req.(*ApplyTSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_GetFullTsrInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullTsrInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).GetFullTsrInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_GetFullTsrInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).GetFullTsrInfo(ctx, req.(*GetFullTsrInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_GetListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).GetListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_GetListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).GetListTickets(ctx, req.(*GetListTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_SetTsrComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTsrCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).SetTsrComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_SetTsrComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).SetTsrComment(ctx, req.(*SetTsrCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_GetTsrCommnts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTsrCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).GetTsrCommnts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_GetTsrCommnts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).GetTsrCommnts(ctx, req.(*GetTsrCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TSRService_GetTsrStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTsrStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TSRServiceServer).GetTsrStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TSRService_GetTsrStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TSRServiceServer).GetTsrStat(ctx, req.(*GetTsrStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TSRService_ServiceDesc is the grpc.ServiceDesc for TSRService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TSRService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tsr.v1.TSRService",
	HandlerType: (*TSRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTSR",
			Handler:    _TSRService_CreateTSR_Handler,
		},
		{
			MethodName: "EmployeeTSR",
			Handler:    _TSRService_EmployeeTSR_Handler,
		},
		{
			MethodName: "ImportanceTSR",
			Handler:    _TSRService_ImportanceTSR_Handler,
		},
		{
			MethodName: "FinishTSR",
			Handler:    _TSRService_FinishTSR_Handler,
		},
		{
			MethodName: "ApplyTSR",
			Handler:    _TSRService_ApplyTSR_Handler,
		},
		{
			MethodName: "GetFullTsrInfo",
			Handler:    _TSRService_GetFullTsrInfo_Handler,
		},
		{
			MethodName: "GetListTickets",
			Handler:    _TSRService_GetListTickets_Handler,
		},
		{
			MethodName: "SetTsrComment",
			Handler:    _TSRService_SetTsrComment_Handler,
		},
		{
			MethodName: "GetTsrCommnts",
			Handler:    _TSRService_GetTsrCommnts_Handler,
		},
		{
			MethodName: "GetTsrStat",
			Handler:    _TSRService_GetTsrStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tsr.proto",
}
