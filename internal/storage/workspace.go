package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IWorkspaceStorage interface {
	// GetUserWorkspaces
	// находит пользователя в БД по почте
	// или возвращает ошибки ...
	GetUserWorkspaces(ctx context.Context, userID uint64) (*entities.Workspace, error)
	// GetByID
	// находит рабочее пространство в БД по его id
	// или возвращает ошибки ...
	GetByID(ctx context.Context, id uint64) (*entities.Workspace, error)
	// Create
	// создает новоt рабочее пространство в БД по данным
	// или возвращает ошибки ...
	Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error)
	// Update
	// обновляет рабочее пространство в БД
	// или возвращает ошибки ...
	Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) (*entities.Workspace, error)
	// Delete
	// удаляет данногt рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(ctx context.Context, id uint64) error
}
