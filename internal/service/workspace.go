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
	GetUserWorkspaces(context.Context, dto.UserID) (*entities.Workspace, error)
	// GetByID
	// находит рабочее пространство в БД по его id
	// или возвращает ошибки ...
	GetWorkspace(context.Context, dto.WorkspaceID) (*entities.Workspace, error)
	// Create
	// создает новоt рабочее пространство в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewWorkspaceInfo) (*entities.Workspace, error)
	// Update
	// обновляет рабочее пространство в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedWorkspaceInfo) error
	// Delete
	// удаляет рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.WorkspaceID) error
}
