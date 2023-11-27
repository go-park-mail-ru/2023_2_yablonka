// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source=user.go -destination=../../mocks/mock_service/user.go -package=mock_service
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"
	dto "server/internal/pkg/dto"
	entities "server/internal/pkg/entities"

	gomock "go.uber.org/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// CheckPassword mocks base method.
func (m *MockIUserService) CheckPassword(arg0 context.Context, arg1 dto.AuthInfo) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockIUserServiceMockRecorder) CheckPassword(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockIUserService)(nil).CheckPassword), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockIUserService) DeleteUser(arg0 context.Context, arg1 dto.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserServiceMockRecorder) DeleteUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserService)(nil).DeleteUser), arg0, arg1)
}

// GetWithID mocks base method.
func (m *MockIUserService) GetWithID(arg0 context.Context, arg1 dto.UserID) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithID", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithID indicates an expected call of GetWithID.
func (mr *MockIUserServiceMockRecorder) GetWithID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithID", reflect.TypeOf((*MockIUserService)(nil).GetWithID), arg0, arg1)
}

// RegisterUser mocks base method.
func (m *MockIUserService) RegisterUser(arg0 context.Context, arg1 dto.AuthInfo) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIUserServiceMockRecorder) RegisterUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserService)(nil).RegisterUser), arg0, arg1)
}

// UpdateAvatar mocks base method.
func (m *MockIUserService) UpdateAvatar(arg0 context.Context, arg1 dto.AvatarChangeInfo) (*dto.UrlObj, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1)
	ret0, _ := ret[0].(*dto.UrlObj)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockIUserServiceMockRecorder) UpdateAvatar(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockIUserService)(nil).UpdateAvatar), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockIUserService) UpdatePassword(arg0 context.Context, arg1 dto.PasswordChangeInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockIUserServiceMockRecorder) UpdatePassword(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockIUserService)(nil).UpdatePassword), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockIUserService) UpdateProfile(arg0 context.Context, arg1 dto.UserProfileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockIUserServiceMockRecorder) UpdateProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockIUserService)(nil).UpdateProfile), arg0, arg1)
}
