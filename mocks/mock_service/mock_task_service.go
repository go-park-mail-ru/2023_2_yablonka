// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\task.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\task.go --destination=./mocks/mock_service/mock_task_service.go --package=mock_service
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

// MockITaskService is a mock of ITaskService interface.
type MockITaskService struct {
	ctrl     *gomock.Controller
	recorder *MockITaskServiceMockRecorder
}

// MockITaskServiceMockRecorder is the mock recorder for MockITaskService.
type MockITaskServiceMockRecorder struct {
	mock *MockITaskService
}

// NewMockITaskService creates a new mock instance.
func NewMockITaskService(ctrl *gomock.Controller) *MockITaskService {
	mock := &MockITaskService{ctrl: ctrl}
	mock.recorder = &MockITaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskService) EXPECT() *MockITaskServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITaskService) Create(arg0 context.Context, arg1 dto.NewTaskInfo) (*entities.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockITaskServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITaskService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockITaskService) Delete(arg0 context.Context, arg1 dto.TaskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITaskServiceMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITaskService)(nil).Delete), arg0, arg1)
}

// Update mocks base method.
func (m *MockITaskService) Update(arg0 context.Context, arg1 dto.UpdatedTaskInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockITaskServiceMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITaskService)(nil).Update), arg0, arg1)
}