// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\storage\checklist_item.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\storage\checklist_item.go --destination=./mocks/mock_storage/checklist_item.go --package=mock_storage
//
// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	dto "server/internal/pkg/dto"

	gomock "go.uber.org/mock/gomock"
)

// MockIChecklistItemStorage is a mock of IChecklistItemStorage interface.
type MockIChecklistItemStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIChecklistItemStorageMockRecorder
}

// MockIChecklistItemStorageMockRecorder is the mock recorder for MockIChecklistItemStorage.
type MockIChecklistItemStorageMockRecorder struct {
	mock *MockIChecklistItemStorage
}

// NewMockIChecklistItemStorage creates a new mock instance.
func NewMockIChecklistItemStorage(ctrl *gomock.Controller) *MockIChecklistItemStorage {
	mock := &MockIChecklistItemStorage{ctrl: ctrl}
	mock.recorder = &MockIChecklistItemStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIChecklistItemStorage) EXPECT() *MockIChecklistItemStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIChecklistItemStorage) Create(arg0 context.Context, arg1 dto.NewChecklistItemInfo) (*dto.ChecklistItemInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dto.ChecklistItemInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIChecklistItemStorageMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIChecklistItemStorage)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIChecklistItemStorage) Delete(arg0 context.Context, arg1 dto.ChecklistItemID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIChecklistItemStorageMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIChecklistItemStorage)(nil).Delete), arg0, arg1)
}

// ReadMany mocks base method.
func (m *MockIChecklistItemStorage) ReadMany(arg0 context.Context, arg1 dto.ChecklistItemStringIDs) (*[]dto.ChecklistItemInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMany", arg0, arg1)
	ret0, _ := ret[0].(*[]dto.ChecklistItemInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMany indicates an expected call of ReadMany.
func (mr *MockIChecklistItemStorageMockRecorder) ReadMany(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMany", reflect.TypeOf((*MockIChecklistItemStorage)(nil).ReadMany), arg0, arg1)
}

// Update mocks base method.
func (m *MockIChecklistItemStorage) Update(arg0 context.Context, arg1 dto.UpdatedChecklistItemInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIChecklistItemStorageMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIChecklistItemStorage)(nil).Update), arg0, arg1)
}

// UpdateOrder mocks base method.
func (m *MockIChecklistItemStorage) UpdateOrder(arg0 context.Context, arg1 dto.ChecklistItemIDs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockIChecklistItemStorageMockRecorder) UpdateOrder(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockIChecklistItemStorage)(nil).UpdateOrder), arg0, arg1)
}
