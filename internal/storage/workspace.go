package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища рабочих пространств
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type IWorkspaceStorage interface {
	// GetUserOwnedWorkspaces
	// находит все рабочие пространства, созданные пользователем
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserOwnedWorkspaces(context.Context, dto.UserID) (*[]dto.UserOwnedWorkspaceInfo, error)
	// GetUserGuestWorkspaces
	// находит все рабочие пространства, где пользователь гость
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestWorkspaces(context.Context, dto.UserID) (*[]dto.UserGuestWorkspaceInfo, error)
	// Create
	// создает новоt рабочее пространство в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewWorkspaceInfo) (*entities.Workspace, error)
	// Update
	// обновляет рабочее пространство в БД
	// или возвращает ошибки ...
	UpdateData(context.Context, dto.UpdatedWorkspaceInfo) error
	// Delete
	// удаляет данногt рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.WorkspaceID) error
}
