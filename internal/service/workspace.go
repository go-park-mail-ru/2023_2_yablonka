package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IWorkspaceService interface {
	// GetUserWorkspaces
	// находит пользователя в БД по почте
	// или возвращает ошибки ...
	GetUserWorkspaces(ctx context.Context, id dto.UserID) (*entities.Workspace, error)
	// GetByID
	// находит рабочее пространство в БД по его id
	// или возвращает ошибки ...
	GetWorkspace(ctx context.Context, id dto.WorkspaceID) (*entities.Workspace, error)
	// Create
	// создает новоt рабочее пространство в БД по данным
	// или возвращает ошибки ...
	Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error)
	// Update
	// обновляет рабочее пространство в БД
	// или возвращает ошибки ...
	Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) error
	// Delete
	// удаляет рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(ctx context.Context, id dto.WorkspaceID) error
}
