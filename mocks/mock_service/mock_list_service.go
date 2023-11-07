// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\list.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\list.go --destination=./mocks/mock_service/mock_list_service.go --package=mock_service
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

// MockIListService is a mock of IListService interface.
type MockIListService struct {
	ctrl     *gomock.Controller
	recorder *MockIListServiceMockRecorder
}

// MockIListServiceMockRecorder is the mock recorder for MockIListService.
type MockIListServiceMockRecorder struct {
	mock *MockIListService
}

// NewMockIListService creates a new mock instance.
func NewMockIListService(ctrl *gomock.Controller) *MockIListService {
	mock := &MockIListService{ctrl: ctrl}
	mock.recorder = &MockIListServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIListService) EXPECT() *MockIListServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIListService) Create(arg0 context.Context, arg1 dto.NewListInfo) (*entities.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIListServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIListService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIListService) Delete(arg0 context.Context, arg1 dto.ListID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIListServiceMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIListService)(nil).Delete), arg0, arg1)
}

// ReadListsInBoard mocks base method.
func (m *MockIListService) ReadListsInBoard(arg0 context.Context, arg1 dto.BoardID) (*[]entities.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadListsInBoard", arg0, arg1)
	ret0, _ := ret[0].(*[]entities.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadListsInBoard indicates an expected call of ReadListsInBoard.
func (mr *MockIListServiceMockRecorder) ReadListsInBoard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadListsInBoard", reflect.TypeOf((*MockIListService)(nil).ReadListsInBoard), arg0, arg1)
}

// Update mocks base method.
func (m *MockIListService) Update(arg0 context.Context, arg1 dto.UpdatedListInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIListServiceMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIListService)(nil).Update), arg0, arg1)
}
