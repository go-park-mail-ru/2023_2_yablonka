// Code generated by MockGen. DO NOT EDIT.
// Source: auth_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=auth_grpc.pb.go -destination=../../../mocks/mock_grcp/auth_grpc.pb.go -package=mock_grcp
//
// Package mock_grcp is a generated GoMock package.
package mock_grcp

import (
	context "context"
	reflect "reflect"
	auth_microservice "server/microservices/auth/auth"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockAuthServiceClient is a mock of AuthServiceClient interface.
type MockAuthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceClientMockRecorder
}

// MockAuthServiceClientMockRecorder is the mock recorder for MockAuthServiceClient.
type MockAuthServiceClientMockRecorder struct {
	mock *MockAuthServiceClient
}

// NewMockAuthServiceClient creates a new mock instance.
func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockAuthServiceClient) AuthUser(ctx context.Context, in *auth_microservice.AuthUserRequest, opts ...grpc.CallOption) (*auth_microservice.AuthUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthUser", varargs...)
	ret0, _ := ret[0].(*auth_microservice.AuthUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockAuthServiceClientMockRecorder) AuthUser(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockAuthServiceClient)(nil).AuthUser), varargs...)
}

// GetLifetime mocks base method.
func (m *MockAuthServiceClient) GetLifetime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*auth_microservice.GetLifetimeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLifetime", varargs...)
	ret0, _ := ret[0].(*auth_microservice.GetLifetimeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifetime indicates an expected call of GetLifetime.
func (mr *MockAuthServiceClientMockRecorder) GetLifetime(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifetime", reflect.TypeOf((*MockAuthServiceClient)(nil).GetLifetime), varargs...)
}

// LogOut mocks base method.
func (m *MockAuthServiceClient) LogOut(ctx context.Context, in *auth_microservice.LogOutRequest, opts ...grpc.CallOption) (*auth_microservice.LogOutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LogOut", varargs...)
	ret0, _ := ret[0].(*auth_microservice.LogOutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogOut indicates an expected call of LogOut.
func (mr *MockAuthServiceClientMockRecorder) LogOut(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOut", reflect.TypeOf((*MockAuthServiceClient)(nil).LogOut), varargs...)
}

// VerifyAuth mocks base method.
func (m *MockAuthServiceClient) VerifyAuth(ctx context.Context, in *auth_microservice.VerifyAuthRequest, opts ...grpc.CallOption) (*auth_microservice.VerifyAuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyAuth", varargs...)
	ret0, _ := ret[0].(*auth_microservice.VerifyAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAuth indicates an expected call of VerifyAuth.
func (mr *MockAuthServiceClientMockRecorder) VerifyAuth(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAuth", reflect.TypeOf((*MockAuthServiceClient)(nil).VerifyAuth), varargs...)
}

// MockAuthServiceServer is a mock of AuthServiceServer interface.
type MockAuthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceServerMockRecorder
}

// MockAuthServiceServerMockRecorder is the mock recorder for MockAuthServiceServer.
type MockAuthServiceServerMockRecorder struct {
	mock *MockAuthServiceServer
}

// NewMockAuthServiceServer creates a new mock instance.
func NewMockAuthServiceServer(ctrl *gomock.Controller) *MockAuthServiceServer {
	mock := &MockAuthServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceServer) EXPECT() *MockAuthServiceServerMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockAuthServiceServer) AuthUser(arg0 context.Context, arg1 *auth_microservice.AuthUserRequest) (*auth_microservice.AuthUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", arg0, arg1)
	ret0, _ := ret[0].(*auth_microservice.AuthUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockAuthServiceServerMockRecorder) AuthUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockAuthServiceServer)(nil).AuthUser), arg0, arg1)
}

// GetLifetime mocks base method.
func (m *MockAuthServiceServer) GetLifetime(arg0 context.Context, arg1 *emptypb.Empty) (*auth_microservice.GetLifetimeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifetime", arg0, arg1)
	ret0, _ := ret[0].(*auth_microservice.GetLifetimeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifetime indicates an expected call of GetLifetime.
func (mr *MockAuthServiceServerMockRecorder) GetLifetime(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifetime", reflect.TypeOf((*MockAuthServiceServer)(nil).GetLifetime), arg0, arg1)
}

// LogOut mocks base method.
func (m *MockAuthServiceServer) LogOut(arg0 context.Context, arg1 *auth_microservice.LogOutRequest) (*auth_microservice.LogOutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogOut", arg0, arg1)
	ret0, _ := ret[0].(*auth_microservice.LogOutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogOut indicates an expected call of LogOut.
func (mr *MockAuthServiceServerMockRecorder) LogOut(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOut", reflect.TypeOf((*MockAuthServiceServer)(nil).LogOut), arg0, arg1)
}

// VerifyAuth mocks base method.
func (m *MockAuthServiceServer) VerifyAuth(arg0 context.Context, arg1 *auth_microservice.VerifyAuthRequest) (*auth_microservice.VerifyAuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAuth", arg0, arg1)
	ret0, _ := ret[0].(*auth_microservice.VerifyAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAuth indicates an expected call of VerifyAuth.
func (mr *MockAuthServiceServerMockRecorder) VerifyAuth(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAuth", reflect.TypeOf((*MockAuthServiceServer)(nil).VerifyAuth), arg0, arg1)
}

// mustEmbedUnimplementedAuthServiceServer mocks base method.
func (m *MockAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServiceServer")
}

// mustEmbedUnimplementedAuthServiceServer indicates an expected call of mustEmbedUnimplementedAuthServiceServer.
func (mr *MockAuthServiceServerMockRecorder) mustEmbedUnimplementedAuthServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServiceServer", reflect.TypeOf((*MockAuthServiceServer)(nil).mustEmbedUnimplementedAuthServiceServer))
}

// MockUnsafeAuthServiceServer is a mock of UnsafeAuthServiceServer interface.
type MockUnsafeAuthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAuthServiceServerMockRecorder
}

// MockUnsafeAuthServiceServerMockRecorder is the mock recorder for MockUnsafeAuthServiceServer.
type MockUnsafeAuthServiceServerMockRecorder struct {
	mock *MockUnsafeAuthServiceServer
}

// NewMockUnsafeAuthServiceServer creates a new mock instance.
func NewMockUnsafeAuthServiceServer(ctrl *gomock.Controller) *MockUnsafeAuthServiceServer {
	mock := &MockUnsafeAuthServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAuthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuthServiceServer) EXPECT() *MockUnsafeAuthServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAuthServiceServer mocks base method.
func (m *MockUnsafeAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServiceServer")
}

// mustEmbedUnimplementedAuthServiceServer indicates an expected call of mustEmbedUnimplementedAuthServiceServer.
func (mr *MockUnsafeAuthServiceServerMockRecorder) mustEmbedUnimplementedAuthServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServiceServer", reflect.TypeOf((*MockUnsafeAuthServiceServer)(nil).mustEmbedUnimplementedAuthServiceServer))
}
