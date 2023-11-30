// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth_microservice

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	AuthUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SessionToken, error)
	VerifyAuth(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*UserID, error)
	LogOut(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetLifetime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*durationpb.Duration, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AuthUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SessionToken, error) {
	out := new(SessionToken)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AuthUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyAuth(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/auth.AuthService/VerifyAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LogOut(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetLifetime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*durationpb.Duration, error) {
	out := new(durationpb.Duration)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetLifetime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	AuthUser(context.Context, *UserID) (*SessionToken, error)
	VerifyAuth(context.Context, *SessionToken) (*UserID, error)
	LogOut(context.Context, *SessionToken) (*emptypb.Empty, error)
	GetLifetime(context.Context, *emptypb.Empty) (*durationpb.Duration, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AuthUser(context.Context, *UserID) (*SessionToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
}
func (UnimplementedAuthServiceServer) VerifyAuth(context.Context, *SessionToken) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAuth not implemented")
}
func (UnimplementedAuthServiceServer) LogOut(context.Context, *SessionToken) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedAuthServiceServer) GetLifetime(context.Context, *emptypb.Empty) (*durationpb.Duration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLifetime not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AuthUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/VerifyAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyAuth(ctx, req.(*SessionToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LogOut(ctx, req.(*SessionToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetLifetime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetLifetime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetLifetime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetLifetime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthUser",
			Handler:    _AuthService_AuthUser_Handler,
		},
		{
			MethodName: "VerifyAuth",
			Handler:    _AuthService_VerifyAuth_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _AuthService_LogOut_Handler,
		},
		{
			MethodName: "GetLifetime",
			Handler:    _AuthService_GetLifetime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/api/auth.proto",
}
