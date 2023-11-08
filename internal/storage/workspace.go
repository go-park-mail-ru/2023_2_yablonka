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
	GetUserWorkspaces(context.Context, dto.UserID) (*dto.AllWorkspaces, error)
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
	// UpdateThumbnail
	// обновляет ссылку на картину рабочего пространства в БД
	// или возвращает ошибки ...
	UpdateThumbnailUrl(context.Context, dto.ImageUrlInfo) error
	// Delete
	// удаляет данногt рабочее пространство в БД по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.WorkspaceID) error
}
