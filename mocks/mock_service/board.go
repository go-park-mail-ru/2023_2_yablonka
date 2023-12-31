// Code generated by MockGen. DO NOT EDIT.
// Source: P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\board.go
//
// Generated by this command:
//
//	mockgen.exe --source=P:\VK Образование\Web\Sem_2\Project\2023_2_yablonka\internal\service\board.go --destination=./mocks/mock_service/board.go --package=mock_service
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

// MockIBoardService is a mock of IBoardService interface.
type MockIBoardService struct {
	ctrl     *gomock.Controller
	recorder *MockIBoardServiceMockRecorder
}

// MockIBoardServiceMockRecorder is the mock recorder for MockIBoardService.
type MockIBoardServiceMockRecorder struct {
	mock *MockIBoardService
}

// NewMockIBoardService creates a new mock instance.
func NewMockIBoardService(ctrl *gomock.Controller) *MockIBoardService {
	mock := &MockIBoardService{ctrl: ctrl}
	mock.recorder = &MockIBoardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBoardService) EXPECT() *MockIBoardServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockIBoardService) AddUser(arg0 context.Context, arg1 dto.AddBoardUserRequest) (dto.UserPublicInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(dto.UserPublicInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIBoardServiceMockRecorder) AddUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIBoardService)(nil).AddUser), arg0, arg1)
}

// Create mocks base method.
func (m *MockIBoardService) Create(arg0 context.Context, arg1 dto.NewBoardInfo) (*entities.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIBoardServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBoardService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIBoardService) Delete(arg0 context.Context, arg1 dto.BoardDeleteRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIBoardServiceMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIBoardService)(nil).Delete), arg0, arg1)
}

// GetFullBoard mocks base method.
func (m *MockIBoardService) GetFullBoard(arg0 context.Context, arg1 dto.IndividualBoardRequest) (*dto.FullBoardResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullBoard", arg0, arg1)
	ret0, _ := ret[0].(*dto.FullBoardResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullBoard indicates an expected call of GetFullBoard.
func (mr *MockIBoardServiceMockRecorder) GetFullBoard(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullBoard", reflect.TypeOf((*MockIBoardService)(nil).GetFullBoard), arg0, arg1)
}

// GetHistory mocks base method.
func (m *MockIBoardService) GetHistory(arg0 context.Context, arg1 dto.BoardID) (*[]dto.BoardHistoryEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", arg0, arg1)
	ret0, _ := ret[0].(*[]dto.BoardHistoryEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockIBoardServiceMockRecorder) GetHistory(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockIBoardService)(nil).GetHistory), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockIBoardService) RemoveUser(arg0 context.Context, arg1 dto.RemoveBoardUserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockIBoardServiceMockRecorder) RemoveUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockIBoardService)(nil).RemoveUser), arg0, arg1)
}

// SubmitEdit mocks base method.
func (m *MockIBoardService) SubmitEdit(arg0 context.Context, arg1 dto.NewHistoryEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitEdit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitEdit indicates an expected call of SubmitEdit.
func (mr *MockIBoardServiceMockRecorder) SubmitEdit(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitEdit", reflect.TypeOf((*MockIBoardService)(nil).SubmitEdit), arg0, arg1)
}

// UpdateData mocks base method.
func (m *MockIBoardService) UpdateData(arg0 context.Context, arg1 dto.UpdatedBoardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateData indicates an expected call of UpdateData.
func (mr *MockIBoardServiceMockRecorder) UpdateData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateData", reflect.TypeOf((*MockIBoardService)(nil).UpdateData), arg0, arg1)
}

// UpdateThumbnail mocks base method.
func (m *MockIBoardService) UpdateThumbnail(arg0 context.Context, arg1 dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThumbnail", arg0, arg1)
	ret0, _ := ret[0].(*dto.UrlObj)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateThumbnail indicates an expected call of UpdateThumbnail.
func (mr *MockIBoardServiceMockRecorder) UpdateThumbnail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThumbnail", reflect.TypeOf((*MockIBoardService)(nil).UpdateThumbnail), arg0, arg1)
}
