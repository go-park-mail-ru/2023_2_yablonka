// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package csrf_microservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CSRFServiceClient is the client API for CSRFService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CSRFServiceClient interface {
	SetupCSRF(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*CSRFData, error)
	VerifyCSRF(ctx context.Context, in *CSRFToken, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCSRF(ctx context.Context, in *CSRFToken, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetLifetime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*durationpb.Duration, error)
}

type cSRFServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCSRFServiceClient(cc grpc.ClientConnInterface) CSRFServiceClient {
	return &cSRFServiceClient{cc}
}

func (c *cSRFServiceClient) SetupCSRF(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*CSRFData, error) {
	out := new(CSRFData)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/SetupCSRF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSRFServiceClient) VerifyCSRF(ctx context.Context, in *CSRFToken, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/VerifyCSRF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSRFServiceClient) DeleteCSRF(ctx context.Context, in *CSRFToken, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/DeleteCSRF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSRFServiceClient) GetLifetime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*durationpb.Duration, error) {
	out := new(durationpb.Duration)
	err := c.cc.Invoke(ctx, "/csrf.CSRFService/GetLifetime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CSRFServiceServer is the server API for CSRFService service.
// All implementations must embed UnimplementedCSRFServiceServer
// for forward compatibility
type CSRFServiceServer interface {
	SetupCSRF(context.Context, *UserID) (*CSRFData, error)
	VerifyCSRF(context.Context, *CSRFToken) (*emptypb.Empty, error)
	DeleteCSRF(context.Context, *CSRFToken) (*emptypb.Empty, error)
	GetLifetime(context.Context, *emptypb.Empty) (*durationpb.Duration, error)
	mustEmbedUnimplementedCSRFServiceServer()
}

// UnimplementedCSRFServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCSRFServiceServer struct {
}

func (UnimplementedCSRFServiceServer) SetupCSRF(context.Context, *UserID) (*CSRFData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupCSRF not implemented")
}
func (UnimplementedCSRFServiceServer) VerifyCSRF(context.Context, *CSRFToken) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCSRF not implemented")
}
func (UnimplementedCSRFServiceServer) DeleteCSRF(context.Context, *CSRFToken) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCSRF not implemented")
}
func (UnimplementedCSRFServiceServer) GetLifetime(context.Context, *emptypb.Empty) (*durationpb.Duration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLifetime not implemented")
}
func (UnimplementedCSRFServiceServer) mustEmbedUnimplementedCSRFServiceServer() {}

// UnsafeCSRFServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CSRFServiceServer will
// result in compilation errors.
type UnsafeCSRFServiceServer interface {
	mustEmbedUnimplementedCSRFServiceServer()
}

func RegisterCSRFServiceServer(s grpc.ServiceRegistrar, srv CSRFServiceServer) {
	s.RegisterService(&CSRFService_ServiceDesc, srv)
}

func _CSRFService_SetupCSRF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).SetupCSRF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/SetupCSRF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).SetupCSRF(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSRFService_VerifyCSRF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRFToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).VerifyCSRF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/VerifyCSRF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).VerifyCSRF(ctx, req.(*CSRFToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSRFService_DeleteCSRF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRFToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).DeleteCSRF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/DeleteCSRF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).DeleteCSRF(ctx, req.(*CSRFToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSRFService_GetLifetime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRFServiceServer).GetLifetime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CSRFService/GetLifetime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRFServiceServer).GetLifetime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CSRFService_ServiceDesc is the grpc.ServiceDesc for CSRFService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CSRFService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "csrf.CSRFService",
	HandlerType: (*CSRFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetupCSRF",
			Handler:    _CSRFService_SetupCSRF_Handler,
		},
		{
			MethodName: "VerifyCSRF",
			Handler:    _CSRFService_VerifyCSRF_Handler,
		},
		{
			MethodName: "DeleteCSRF",
			Handler:    _CSRFService_DeleteCSRF_Handler,
		},
		{
			MethodName: "GetLifetime",
			Handler:    _CSRFService_GetLifetime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "csrf/api/csrf.proto",
}
