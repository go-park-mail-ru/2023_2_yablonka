package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type WorkspaceService struct {
	storage storage.IWorkspaceStorage
}

// NewWorkspaceService
// возвращает UserService с инициализированным хранилищем пользователей
func NewWorkspaceService(storage storage.IWorkspaceStorage) *WorkspaceService {
	return &WorkspaceService{
		storage: storage,
	}
}

// GetUserWorkspaces
// находит пользователя в БД по почте
// или возвращает ошибки ...
func (us WorkspaceService) GetUserWorkspaces(ctx context.Context, id dto.UserID) (*entities.Workspace, error) {
	return nil, nil
}

// GetByID
// находит рабочее пространство в БД по его id
// или возвращает ошибки ...
func (us WorkspaceService) GetWorkspace(ctx context.Context, id dto.WorkspaceID) (*entities.Workspace, error) {
	return nil, nil
}

// Create
// создает новоt рабочее пространство в БД по данным
// или возвращает ошибки ...
func (us WorkspaceService) Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error) {
	return nil, nil
}

// Update
// обновляет рабочее пространство в БД
// или возвращает ошибки ...
func (us WorkspaceService) Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) error {
	return nil
}

// Delete
// удаляет рабочее пространство в БД по id
// или возвращает ошибки ...
func (us WorkspaceService) Delete(ctx context.Context, id dto.WorkspaceID) error {
	return nil
}
