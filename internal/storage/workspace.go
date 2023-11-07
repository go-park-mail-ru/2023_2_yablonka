package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IWorkspaceStorage interface {
	// GetUserOwnedWorkspaces
	// находит все рабочие пространства, связанные с пользователем, группированные по ролям
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserWorkspaces(context.Context, dto.UserID) (*[]entities.Workspace, error)
	// GetByID
	// находит рабочее пространство в БД по его id
	// или возвращает ошибки ...
	GetWithBoards(context.Context, dto.WorkspaceID) (*entities.Workspace, error)
	// Create
	// создает новоt рабочее пространство в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewWorkspaceInfo) (*entities.Workspace, error)
	// Update
	// обновляет рабочее пространство в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedWorkspaceInfo) error
	// Delete
	// удаляет данногt рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.WorkspaceID) error
	GetUsersInWorkspace(context.Context, dto.WorkspaceID) (*dto.UsersAndRoles, error)
}
