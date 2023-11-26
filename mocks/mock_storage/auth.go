// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=../../mocks/mock_storage/auth.go -package=mock_storage
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

// MockIAuthStorage is a mock of IAuthStorage interface.
type MockIAuthStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthStorageMockRecorder
}

// MockIAuthStorageMockRecorder is the mock recorder for MockIAuthStorage.
type MockIAuthStorageMockRecorder struct {
	mock *MockIAuthStorage
}

// NewMockIAuthStorage creates a new mock instance.
func NewMockIAuthStorage(ctrl *gomock.Controller) *MockIAuthStorage {
	mock := &MockIAuthStorage{ctrl: ctrl}
	mock.recorder = &MockIAuthStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthStorage) EXPECT() *MockIAuthStorageMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockIAuthStorage) CreateSession(arg0 context.Context, arg1 *entities.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockIAuthStorageMockRecorder) CreateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockIAuthStorage)(nil).CreateSession), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockIAuthStorage) DeleteSession(arg0 context.Context, arg1 dto.SessionToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockIAuthStorageMockRecorder) DeleteSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockIAuthStorage)(nil).DeleteSession), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockIAuthStorage) GetSession(arg0 context.Context, arg1 dto.SessionToken) (*entities.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(*entities.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockIAuthStorageMockRecorder) GetSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockIAuthStorage)(nil).GetSession), arg0, arg1)
}