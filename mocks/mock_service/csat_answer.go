// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\csat_answer.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\csat_answer.go --destination=./mocks/mock_service/csat_answer.go --package=mock_service
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"
	dto "server/internal/pkg/dto"

	gomock "go.uber.org/mock/gomock"
)

// MockICSATSAnswerService is a mock of ICSATSAnswerService interface.
type MockICSATSAnswerService struct {
	ctrl     *gomock.Controller
	recorder *MockICSATSAnswerServiceMockRecorder
}

// MockICSATSAnswerServiceMockRecorder is the mock recorder for MockICSATSAnswerService.
type MockICSATSAnswerServiceMockRecorder struct {
	mock *MockICSATSAnswerService
}

// NewMockICSATSAnswerService creates a new mock instance.
func NewMockICSATSAnswerService(ctrl *gomock.Controller) *MockICSATSAnswerService {
	mock := &MockICSATSAnswerService{ctrl: ctrl}
	mock.recorder = &MockICSATSAnswerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICSATSAnswerService) EXPECT() *MockICSATSAnswerServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICSATSAnswerService) Create(arg0 context.Context, arg1 dto.NewCSATAnswer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockICSATSAnswerServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICSATSAnswerService)(nil).Create), arg0, arg1)
}
