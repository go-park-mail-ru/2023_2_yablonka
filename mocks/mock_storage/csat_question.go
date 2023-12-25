// Code generated by MockGen. DO NOT EDIT.
// Source: csat_question.go
//
// Generated by this command:
//
//	mockgen -source=csat_question.go -destination=../../mocks/mock_storage/csat_question.go -package=mock_storage
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

// MockICSATQuestionStorage is a mock of ICSATQuestionStorage interface.
type MockICSATQuestionStorage struct {
	ctrl     *gomock.Controller
	recorder *MockICSATQuestionStorageMockRecorder
}

// MockICSATQuestionStorageMockRecorder is the mock recorder for MockICSATQuestionStorage.
type MockICSATQuestionStorageMockRecorder struct {
	mock *MockICSATQuestionStorage
}

// NewMockICSATQuestionStorage creates a new mock instance.
func NewMockICSATQuestionStorage(ctrl *gomock.Controller) *MockICSATQuestionStorage {
	mock := &MockICSATQuestionStorage{ctrl: ctrl}
	mock.recorder = &MockICSATQuestionStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICSATQuestionStorage) EXPECT() *MockICSATQuestionStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICSATQuestionStorage) Create(arg0 context.Context, arg1 dto.NewCSATQuestion) (*dto.CSATQuestionFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dto.CSATQuestionFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICSATQuestionStorageMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICSATQuestionStorage)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockICSATQuestionStorage) Delete(arg0 context.Context, arg1 dto.CSATQuestionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICSATQuestionStorageMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICSATQuestionStorage)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockICSATQuestionStorage) GetAll(arg0 context.Context) (*[]dto.CSATQuestionFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(*[]dto.CSATQuestionFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockICSATQuestionStorageMockRecorder) GetAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockICSATQuestionStorage)(nil).GetAll), arg0)
}

// GetQuestionType mocks base method.
func (m *MockICSATQuestionStorage) GetQuestionType(arg0 context.Context, arg1 dto.CSATQuestionID) (*entities.QuestionType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionType", arg0, arg1)
	ret0, _ := ret[0].(*entities.QuestionType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionType indicates an expected call of GetQuestionType.
func (mr *MockICSATQuestionStorageMockRecorder) GetQuestionType(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionType", reflect.TypeOf((*MockICSATQuestionStorage)(nil).GetQuestionType), arg0, arg1)
}

// GetQuestionTypeWithName mocks base method.
func (m *MockICSATQuestionStorage) GetQuestionTypeWithName(arg0 context.Context, arg1 dto.CSATQuestionTypeName) (*entities.QuestionType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionTypeWithName", arg0, arg1)
	ret0, _ := ret[0].(*entities.QuestionType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionTypeWithName indicates an expected call of GetQuestionTypeWithName.
func (mr *MockICSATQuestionStorageMockRecorder) GetQuestionTypeWithName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionTypeWithName", reflect.TypeOf((*MockICSATQuestionStorage)(nil).GetQuestionTypeWithName), arg0, arg1)
}

// GetStats mocks base method.
func (m *MockICSATQuestionStorage) GetStats(arg0 context.Context) (*[]dto.QuestionWithStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", arg0)
	ret0, _ := ret[0].(*[]dto.QuestionWithStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockICSATQuestionStorageMockRecorder) GetStats(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockICSATQuestionStorage)(nil).GetStats), arg0)
}

// Update mocks base method.
func (m *MockICSATQuestionStorage) Update(arg0 context.Context, arg1 dto.UpdatedCSATQuestion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockICSATQuestionStorageMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICSATQuestionStorage)(nil).Update), arg0, arg1)
}
