// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go
//
// Generated by this command:
//
//	mockgen -source=comment.go -destination=../../mocks/mock_service/comment.go -package=mock_service
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

// MockICommentService is a mock of ICommentService interface.
type MockICommentService struct {
	ctrl     *gomock.Controller
	recorder *MockICommentServiceMockRecorder
}

// MockICommentServiceMockRecorder is the mock recorder for MockICommentService.
type MockICommentServiceMockRecorder struct {
	mock *MockICommentService
}

// NewMockICommentService creates a new mock instance.
func NewMockICommentService(ctrl *gomock.Controller) *MockICommentService {
	mock := &MockICommentService{ctrl: ctrl}
	mock.recorder = &MockICommentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommentService) EXPECT() *MockICommentServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICommentService) Create(arg0 context.Context, arg1 dto.NewCommentInfo) (*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICommentServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICommentService)(nil).Create), arg0, arg1)
}

// GetFromTask mocks base method.
func (m *MockICommentService) GetFromTask(arg0 context.Context, arg1 dto.TaskID) (*[]dto.CommentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromTask", arg0, arg1)
	ret0, _ := ret[0].(*[]dto.CommentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromTask indicates an expected call of GetFromTask.
func (mr *MockICommentServiceMockRecorder) GetFromTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromTask", reflect.TypeOf((*MockICommentService)(nil).GetFromTask), arg0, arg1)
}