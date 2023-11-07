// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/workspace.go
//
// Generated by this command:
//
//	mockgen.exe --source=internal/service/workspace.go --destination=mocks/mock_workspace_service.go --package=mock_service
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
func (m *MockIWorkspaceService) Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, info)
	ret0, _ := ret[0].(*entities.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIWorkspaceServiceMockRecorder) Create(ctx, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIWorkspaceService)(nil).Create), ctx, info)
}

// Delete mocks base method.
func (m *MockIWorkspaceService) Delete(ctx context.Context, id dto.WorkspaceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIWorkspaceServiceMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIWorkspaceService)(nil).Delete), ctx, id)
}

// GetUserWorkspaces mocks base method.
func (m *MockIWorkspaceService) GetUserWorkspaces(ctx context.Context, id dto.UserID) (*entities.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWorkspaces", ctx, id)
	ret0, _ := ret[0].(*entities.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWorkspaces indicates an expected call of GetUserWorkspaces.
func (mr *MockIWorkspaceServiceMockRecorder) GetUserWorkspaces(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWorkspaces", reflect.TypeOf((*MockIWorkspaceService)(nil).GetUserWorkspaces), ctx, id)
}

// GetWorkspace mocks base method.
func (m *MockIWorkspaceService) GetWorkspace(ctx context.Context, id dto.WorkspaceID) (*entities.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", ctx, id)
	ret0, _ := ret[0].(*entities.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockIWorkspaceServiceMockRecorder) GetWorkspace(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockIWorkspaceService)(nil).GetWorkspace), ctx, id)
}

// Update mocks base method.
func (m *MockIWorkspaceService) Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIWorkspaceServiceMockRecorder) Update(ctx, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIWorkspaceService)(nil).Update), ctx, info)
}
