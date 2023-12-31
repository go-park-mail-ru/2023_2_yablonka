package microservice

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	logger "server/internal/logging"

	"github.com/google/uuid"
)

type WorkspaceService struct {
	storage storage.IWorkspaceStorage
}

const nodeName = "service"

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
func (ws WorkspaceService) GetUserWorkspaces(ctx context.Context, userID dto.UserID) (*dto.AllWorkspaces, error) {
	funcName := "WorkspaceService.GetUserWorkspaces"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	ownedWorkspaces, err := ws.storage.GetUserOwnedWorkspaces(ctx, userID)
	if err != nil {
		logger.DebugFmt("Failed to get user owned workspaces", requestID.String(), funcName, nodeName)
		return nil, err
	}

	guestWorkspaces, err := ws.storage.GetUserGuestWorkspaces(ctx, userID)
	if err != nil {
		logger.DebugFmt("Failed to get user guest workspaces", requestID.String(), funcName, nodeName)
		return nil, err
	}

	return &dto.AllWorkspaces{
		OwnedWorkspaces: *ownedWorkspaces,
		GuestWorkspaces: *guestWorkspaces,
	}, nil
}

/*

// GetWorkspace
// находит рабочее пространство в БД по его id
// или возвращает ошибки ...
func (ws WorkspaceService) GetWorkspace(ctx context.Context, id dto.WorkspaceID) (*entities.Workspace, error) {
	return ws.storage.GetWorkspace(ctx, id)
}

*/

// Create
// создает новоt рабочее пространство по данным
// или возвращает ошибки ...
func (ws WorkspaceService) Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error) {
	return ws.storage.Create(ctx, info)
}

// UpdateData
// обновляет рабочее пространство
// или возвращает ошибки .....
func (ws WorkspaceService) UpdateData(ctx context.Context, info dto.UpdatedWorkspaceInfo) error {
	return ws.storage.UpdateData(ctx, info)
}

// Delete
// удаляет рабочее пространство в БД по id
// или возвращает ошибки ...
func (ws WorkspaceService) Delete(ctx context.Context, id dto.WorkspaceID) error {
	return ws.storage.Delete(ctx, id)
}
