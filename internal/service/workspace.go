package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для сервиса рабочих пространств
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IWorkspaceService interface {
	// GetUserWorkspaces
	// находит пользователя по почте
	// или возвращает ошибки ...
	GetUserWorkspaces(context.Context, dto.UserID) (*dto.AllWorkspaces, error)
	// GetByID
	// находит рабочее пространство по его id
	// или возвращает ошибки ...
	GetWorkspace(context.Context, dto.WorkspaceID) (*entities.Workspace, error)
	// Create
	// создает новоt рабочее пространство по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewWorkspaceInfo) (*entities.Workspace, error)
	// UpdateData
	// обновляет рабочее пространство
	// или возвращает ошибки ...
	UpdateData(context.Context, dto.UpdatedWorkspaceInfo) error
	// UpdateUsers
	// обновляет список пользователей рабочего пространства
	// или возвращает ошибки ...
	UpdateUsers(context.Context, dto.ChangeWorkspaceGuestsInfo) error
	// Delete
	// удаляет рабочее пространство по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.WorkspaceID) error
}
