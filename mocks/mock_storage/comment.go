// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go
//
// Generated by this command:
//
//	mockgen -source=comment.go -destination=../../mocks/mock_storage/comment.go -package=mock_storage
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

// MockICommentStorage is a mock of ICommentStorage interface.
type MockICommentStorage struct {
	ctrl     *gomock.Controller
	recorder *MockICommentStorageMockRecorder
}

// MockICommentStorageMockRecorder is the mock recorder for MockICommentStorage.
type MockICommentStorageMockRecorder struct {
	mock *MockICommentStorage
}

// NewMockICommentStorage creates a new mock instance.
func NewMockICommentStorage(ctrl *gomock.Controller) *MockICommentStorage {
	mock := &MockICommentStorage{ctrl: ctrl}
	mock.recorder = &MockICommentStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommentStorage) EXPECT() *MockICommentStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICommentStorage) Create(arg0 context.Context, arg1 dto.NewCommentInfo) (*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICommentStorageMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICommentStorage)(nil).Create), arg0, arg1)
}

// GetFromTask mocks base method.
func (m *MockICommentStorage) GetFromTask(arg0 context.Context, arg1 dto.TaskID) (*[]dto.CommentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromTask", arg0, arg1)
	ret0, _ := ret[0].(*[]dto.CommentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromTask indicates an expected call of GetFromTask.
func (mr *MockICommentStorageMockRecorder) GetFromTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromTask", reflect.TypeOf((*MockICommentStorage)(nil).GetFromTask), arg0, arg1)
}
