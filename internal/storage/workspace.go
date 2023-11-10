package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IWorkspaceStorage interface {
	// GetUserOwnedWorkspaces
	// находит все рабочие пространства, созданные пользователем
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserOwnedWorkspaces(context.Context, dto.UserID) (*[]dto.UserOwnedWorkspaceInfo, error)
	// GetUserGuestWorkspaces
	// находит все рабочие пространства, где пользователь гость
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestWorkspaces(context.Context, dto.UserID) (*[]dto.UserGuestWorkspaceInfo, error)
	// GetWorkspace
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
	UpdateData(context.Context, dto.UpdatedWorkspaceInfo) error
	// UpdateUsers
	// обновляет людей с доступом в рабочее пространство в БД
	// или возвращает ошибки ...
	UpdateUsers(context.Context, dto.ChangeWorkspaceGuestsInfo) error
	// Delete
	// удаляет данногt рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.WorkspaceID) error
}
