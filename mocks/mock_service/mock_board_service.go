// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/board.go
//
// Generated by this command:
//
//	mockgen.exe --source=internal/service/board.go --destination=mocks/mock_board_service.go --package=mock_service
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"
	dto "server/internal/pkg/dto"

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

// GetUserGuestBoards mocks base method.
func (m *MockIBoardService) GetUserGuestBoards(arg0 context.Context, arg1 dto.VerifiedAuthInfo) ([]dto.UserGuestBoardInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGuestBoards", arg0, arg1)
	ret0, _ := ret[0].([]dto.UserGuestBoardInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGuestBoards indicates an expected call of GetUserGuestBoards.
func (mr *MockIBoardServiceMockRecorder) GetUserGuestBoards(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGuestBoards", reflect.TypeOf((*MockIBoardService)(nil).GetUserGuestBoards), arg0, arg1)
}

// GetUserOwnedBoards mocks base method.
func (m *MockIBoardService) GetUserOwnedBoards(arg0 context.Context, arg1 dto.VerifiedAuthInfo) ([]dto.UserOwnedBoardInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserOwnedBoards", arg0, arg1)
	ret0, _ := ret[0].([]dto.UserOwnedBoardInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserOwnedBoards indicates an expected call of GetUserOwnedBoards.
func (mr *MockIBoardServiceMockRecorder) GetUserOwnedBoards(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserOwnedBoards", reflect.TypeOf((*MockIBoardService)(nil).GetUserOwnedBoards), arg0, arg1)
}
