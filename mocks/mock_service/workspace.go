// Code generated by MockGen. DO NOT EDIT.
// Source: workspace.go
//
// Generated by this command:
//
//	mockgen -source=workspace.go -destination=../../mocks/mock_service/workspace.go -package=mock_service
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

// MockIWorkspaceService is a mock of IWorkspaceService interface.
type MockIWorkspaceService struct {
	ctrl     *gomock.Controller
	recorder *MockIWorkspaceServiceMockRecorder
}

// MockIWorkspaceServiceMockRecorder is the mock recorder for MockIWorkspaceService.
type MockIWorkspaceServiceMockRecorder struct {
	mock *MockIWorkspaceService
}

// NewMockIWorkspaceService creates a new mock instance.
func NewMockIWorkspaceService(ctrl *gomock.Controller) *MockIWorkspaceService {
	mock := &MockIWorkspaceService{ctrl: ctrl}
	mock.recorder = &MockIWorkspaceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWorkspaceService) EXPECT() *MockIWorkspaceServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIWorkspaceService) Create(arg0 context.Context, arg1 dto.NewWorkspaceInfo) (*entities.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entities.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIWorkspaceServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIWorkspaceService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIWorkspaceService) Delete(arg0 context.Context, arg1 dto.WorkspaceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIWorkspaceServiceMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIWorkspaceService)(nil).Delete), arg0, arg1)
}

// GetUserWorkspaces mocks base method.
func (m *MockIWorkspaceService) GetUserWorkspaces(arg0 context.Context, arg1 dto.UserID) (*dto.AllWorkspaces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWorkspaces", arg0, arg1)
	ret0, _ := ret[0].(*dto.AllWorkspaces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWorkspaces indicates an expected call of GetUserWorkspaces.
func (mr *MockIWorkspaceServiceMockRecorder) GetUserWorkspaces(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWorkspaces", reflect.TypeOf((*MockIWorkspaceService)(nil).GetUserWorkspaces), arg0, arg1)
}

// GetWorkspace mocks base method.
func (m *MockIWorkspaceService) GetWorkspace(arg0 context.Context, arg1 dto.WorkspaceID) (*entities.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*entities.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockIWorkspaceServiceMockRecorder) GetWorkspace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockIWorkspaceService)(nil).GetWorkspace), arg0, arg1)
}

// UpdateData mocks base method.
func (m *MockIWorkspaceService) UpdateData(arg0 context.Context, arg1 dto.UpdatedWorkspaceInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateData indicates an expected call of UpdateData.
func (mr *MockIWorkspaceServiceMockRecorder) UpdateData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateData", reflect.TypeOf((*MockIWorkspaceService)(nil).UpdateData), arg0, arg1)
}

// UpdateUsers mocks base method.
func (m *MockIWorkspaceService) UpdateUsers(arg0 context.Context, arg1 dto.ChangeWorkspaceGuestsInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUsers indicates an expected call of UpdateUsers.
func (mr *MockIWorkspaceServiceMockRecorder) UpdateUsers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsers", reflect.TypeOf((*MockIWorkspaceService)(nil).UpdateUsers), arg0, arg1)
}
