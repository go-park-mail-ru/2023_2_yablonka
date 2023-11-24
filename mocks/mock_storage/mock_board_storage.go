// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\storage\board.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\storage\board.go --destination=./mocks/mock_storage/mock_board_storage.go --package=mock_storage
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

// MockIBoardStorage is a mock of IBoardStorage interface.
type MockIBoardStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIBoardStorageMockRecorder
}

// MockIBoardStorageMockRecorder is the mock recorder for MockIBoardStorage.
type MockIBoardStorageMockRecorder struct {
	mock *MockIBoardStorage
}

// NewMockIBoardStorage creates a new mock instance.
func NewMockIBoardStorage(ctrl *gomock.Controller) *MockIBoardStorage {
	mock := &MockIBoardStorage{ctrl: ctrl}
	mock.recorder = &MockIBoardStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBoardStorage) EXPECT() *MockIBoardStorageMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockIBoardStorage) AddUser(arg0 context.Context, arg1 dto.UserBoardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIBoardStorageMockRecorder) AddUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIBoardStorage)(nil).AddUser), arg0, arg1)
}

// Create mocks base method.
func (m *MockIBoardStorage) Create(arg0 context.Context, arg1 dto.NewBoardInfo) (*entities.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIBoardStorageMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBoardStorage)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIBoardStorage) Delete(arg0 context.Context, arg1 dto.BoardID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIBoardStorageMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIBoardStorage)(nil).Delete), arg0, arg1)
}

// GetById mocks base method.
func (m *MockIBoardStorage) GetById(arg0 context.Context, arg1 dto.BoardID) (*dto.FullBoardResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*dto.FullBoardResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIBoardStorageMockRecorder) GetById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIBoardStorage)(nil).GetById), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockIBoardStorage) GetUsers(arg0 context.Context, arg1 dto.BoardID) (*[]dto.UserPublicInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].(*[]dto.UserPublicInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockIBoardStorageMockRecorder) GetUsers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIBoardStorage)(nil).GetUsers), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockIBoardStorage) RemoveUser(arg0 context.Context, arg1 dto.UserBoardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockIBoardStorageMockRecorder) RemoveUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockIBoardStorage)(nil).RemoveUser), arg0, arg1)
}

// UpdateData mocks base method.
func (m *MockIBoardStorage) UpdateData(arg0 context.Context, arg1 dto.UpdatedBoardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateData indicates an expected call of UpdateData.
func (mr *MockIBoardStorageMockRecorder) UpdateData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateData", reflect.TypeOf((*MockIBoardStorage)(nil).UpdateData), arg0, arg1)
}

// UpdateThumbnailUrl mocks base method.
func (m *MockIBoardStorage) UpdateThumbnailUrl(arg0 context.Context, arg1 dto.ImageUrlInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThumbnailUrl", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateThumbnailUrl indicates an expected call of UpdateThumbnailUrl.
func (mr *MockIBoardStorageMockRecorder) UpdateThumbnailUrl(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThumbnailUrl", reflect.TypeOf((*MockIBoardStorage)(nil).UpdateThumbnailUrl), arg0, arg1)
}
