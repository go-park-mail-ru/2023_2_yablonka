// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\storage\user.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\storage\user.go --destination=./mocks/mock_storage/user.go --package=mock_storage
//
// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	dto "server/internal/pkg/dto"
	entities "server/internal/pkg/entities"

	gomock "go.uber.org/mock/gomock"
)

// MockIUserStorage is a mock of IUserStorage interface.
type MockIUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIUserStorageMockRecorder
}

// MockIUserStorageMockRecorder is the mock recorder for MockIUserStorage.
type MockIUserStorageMockRecorder struct {
	mock *MockIUserStorage
}

// NewMockIUserStorage creates a new mock instance.
func NewMockIUserStorage(ctrl *gomock.Controller) *MockIUserStorage {
	mock := &MockIUserStorage{ctrl: ctrl}
	mock.recorder = &MockIUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserStorage) EXPECT() *MockIUserStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserStorage) Create(arg0 context.Context, arg1 dto.SignupInfo) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserStorageMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserStorage)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIUserStorage) Delete(arg0 context.Context, arg1 dto.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserStorageMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserStorage)(nil).Delete), arg0, arg1)
}

// GetLoginInfoWithID mocks base method.
func (m *MockIUserStorage) GetLoginInfoWithID(arg0 context.Context, arg1 dto.UserID) (*dto.LoginInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoginInfoWithID", arg0, arg1)
	ret0, _ := ret[0].(*dto.LoginInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoginInfoWithID indicates an expected call of GetLoginInfoWithID.
func (mr *MockIUserStorageMockRecorder) GetLoginInfoWithID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginInfoWithID", reflect.TypeOf((*MockIUserStorage)(nil).GetLoginInfoWithID), arg0, arg1)
}

// GetWithID mocks base method.
func (m *MockIUserStorage) GetWithID(arg0 context.Context, arg1 dto.UserID) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithID", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithID indicates an expected call of GetWithID.
func (mr *MockIUserStorageMockRecorder) GetWithID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithID", reflect.TypeOf((*MockIUserStorage)(nil).GetWithID), arg0, arg1)
}

// GetWithLogin mocks base method.
func (m *MockIUserStorage) GetWithLogin(arg0 context.Context, arg1 dto.UserLogin) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithLogin", arg0, arg1)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithLogin indicates an expected call of GetWithLogin.
func (mr *MockIUserStorageMockRecorder) GetWithLogin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithLogin", reflect.TypeOf((*MockIUserStorage)(nil).GetWithLogin), arg0, arg1)
}

// UpdateAvatarUrl mocks base method.
func (m *MockIUserStorage) UpdateAvatarUrl(arg0 context.Context, arg1 dto.UserImageUrlInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatarUrl", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatarUrl indicates an expected call of UpdateAvatarUrl.
func (mr *MockIUserStorageMockRecorder) UpdateAvatarUrl(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatarUrl", reflect.TypeOf((*MockIUserStorage)(nil).UpdateAvatarUrl), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockIUserStorage) UpdatePassword(arg0 context.Context, arg1 dto.PasswordHashesInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockIUserStorageMockRecorder) UpdatePassword(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockIUserStorage)(nil).UpdatePassword), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockIUserStorage) UpdateProfile(arg0 context.Context, arg1 dto.UserProfileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockIUserStorageMockRecorder) UpdateProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockIUserStorage)(nil).UpdateProfile), arg0, arg1)
}
