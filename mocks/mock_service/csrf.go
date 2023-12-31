// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\csrf.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\csrf.go --destination=./mocks/mock_service/csrf.go --package=mock_service
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"
	dto "server/internal/pkg/dto"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockICSRFService is a mock of ICSRFService interface.
type MockICSRFService struct {
	ctrl     *gomock.Controller
	recorder *MockICSRFServiceMockRecorder
}

// MockICSRFServiceMockRecorder is the mock recorder for MockICSRFService.
type MockICSRFServiceMockRecorder struct {
	mock *MockICSRFService
}

// NewMockICSRFService creates a new mock instance.
func NewMockICSRFService(ctrl *gomock.Controller) *MockICSRFService {
	mock := &MockICSRFService{ctrl: ctrl}
	mock.recorder = &MockICSRFServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICSRFService) EXPECT() *MockICSRFServiceMockRecorder {
	return m.recorder
}

// DeleteCSRF mocks base method.
func (m *MockICSRFService) DeleteCSRF(arg0 context.Context, arg1 dto.CSRFToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCSRF", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCSRF indicates an expected call of DeleteCSRF.
func (mr *MockICSRFServiceMockRecorder) DeleteCSRF(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCSRF", reflect.TypeOf((*MockICSRFService)(nil).DeleteCSRF), arg0, arg1)
}

// GetLifetime mocks base method.
func (m *MockICSRFService) GetLifetime(arg0 context.Context) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifetime", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetLifetime indicates an expected call of GetLifetime.
func (mr *MockICSRFServiceMockRecorder) GetLifetime(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifetime", reflect.TypeOf((*MockICSRFService)(nil).GetLifetime), arg0)
}

// SetupCSRF mocks base method.
func (m *MockICSRFService) SetupCSRF(arg0 context.Context, arg1 dto.UserID) (dto.CSRFData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupCSRF", arg0, arg1)
	ret0, _ := ret[0].(dto.CSRFData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupCSRF indicates an expected call of SetupCSRF.
func (mr *MockICSRFServiceMockRecorder) SetupCSRF(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupCSRF", reflect.TypeOf((*MockICSRFService)(nil).SetupCSRF), arg0, arg1)
}

// VerifyCSRF mocks base method.
func (m *MockICSRFService) VerifyCSRF(arg0 context.Context, arg1 dto.CSRFToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCSRF", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyCSRF indicates an expected call of VerifyCSRF.
func (mr *MockICSRFServiceMockRecorder) VerifyCSRF(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCSRF", reflect.TypeOf((*MockICSRFService)(nil).VerifyCSRF), arg0, arg1)
}
