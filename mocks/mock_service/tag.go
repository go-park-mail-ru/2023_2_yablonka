// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\tag.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\tag.go --destination=./mocks/mock_service/tag.go --package=mock_service
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

// MockITagService is a mock of ITagService interface.
type MockITagService struct {
	ctrl     *gomock.Controller
	recorder *MockITagServiceMockRecorder
}

// MockITagServiceMockRecorder is the mock recorder for MockITagService.
type MockITagServiceMockRecorder struct {
	mock *MockITagService
}

// NewMockITagService creates a new mock instance.
func NewMockITagService(ctrl *gomock.Controller) *MockITagService {
	mock := &MockITagService{ctrl: ctrl}
	mock.recorder = &MockITagServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITagService) EXPECT() *MockITagServiceMockRecorder {
	return m.recorder
}

// AddToTask mocks base method.
func (m *MockITagService) AddToTask(arg0 context.Context, arg1 dto.TagAndTaskIDs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToTask indicates an expected call of AddToTask.
func (mr *MockITagServiceMockRecorder) AddToTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToTask", reflect.TypeOf((*MockITagService)(nil).AddToTask), arg0, arg1)
}

// Create mocks base method.
func (m *MockITagService) Create(arg0 context.Context, arg1 dto.NewTagInfo) (*entities.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockITagServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITagService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockITagService) Delete(arg0 context.Context, arg1 dto.TagID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITagServiceMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITagService)(nil).Delete), arg0, arg1)
}

// RemoveFromTask mocks base method.
func (m *MockITagService) RemoveFromTask(arg0 context.Context, arg1 dto.TagAndTaskIDs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromTask indicates an expected call of RemoveFromTask.
func (mr *MockITagServiceMockRecorder) RemoveFromTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromTask", reflect.TypeOf((*MockITagService)(nil).RemoveFromTask), arg0, arg1)
}

// Update mocks base method.
func (m *MockITagService) Update(arg0 context.Context, arg1 dto.UpdatedTagInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockITagServiceMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITagService)(nil).Update), arg0, arg1)
}
