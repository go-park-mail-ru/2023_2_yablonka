package workspace

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	embedded "server/internal/service/workspace/embedded"
	micro "server/internal/service/workspace/microservice"

	"google.golang.org/grpc"
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

func NewEmbeddedWorkspaceService(workspaceStorage storage.IWorkspaceStorage) *embedded.WorkspaceService {
	return embedded.NewWorkspaceService(workspaceStorage)
}

// TODO: User microservice
func NewMicroWorkspaceService(workspaceStorage storage.IWorkspaceStorage, connection *grpc.ClientConn) *micro.WorkspaceService {
	return micro.NewWorkspaceService(workspaceStorage)
}
