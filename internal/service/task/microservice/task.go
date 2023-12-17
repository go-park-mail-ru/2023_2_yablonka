package task

import (
	"context"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	"github.com/google/uuid"
)

const nodeName = "service"

type TaskService struct {
	storage     storage.ITaskStorage
	userStorage storage.IUserStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewTaskService(ts storage.ITaskStorage, us storage.IUserStorage) *TaskService {
	return &TaskService{
		storage:     ts,
		userStorage: us,
	}
}

// Read
// возвращает задание
// или возвращает ошибки ...
func (ts TaskService) Read(ctx context.Context, id dto.TaskID) (*dto.SingleTaskInfo, error) {
	return ts.storage.Read(ctx, id)
}

// Create
// создает новое задание
// или возвращает ошибки ...
func (ts TaskService) Create(ctx context.Context, info dto.NewTaskInfo) (*entities.Task, error) {
	return ts.storage.Create(ctx, info)
}

// Update
// обновляет задание
// или возвращает ошибки ...
func (ts TaskService) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	return ts.storage.Update(ctx, info)
}

// Delete
// удаляет задание
// или возвращает ошибки ...
func (ts TaskService) Delete(ctx context.Context, id dto.TaskID) error {
	return ts.storage.Delete(ctx, id)
}

// AddUser
// добавляет пользователя в карточку
// или возвращает ошибки ...
func (ts TaskService) AddUser(ctx context.Context, info dto.AddTaskUserInfo) error {
	funcName := "TaskService.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	userAccess, err := ts.storage.CheckAccess(ctx, (dto.CheckTaskAccessInfo)(info))
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if userAccess {
		return apperrors.ErrUserAlreadyInTask
	}
	logger.DebugFmt("user not in task", requestID.String(), funcName, nodeName)

	return ts.storage.AddUser(ctx, info)
}

// RemoveUser
// удаляет пользователя из карточки
// или возвращает ошибки ...
func (ts TaskService) RemoveUser(ctx context.Context, info dto.RemoveTaskUserInfo) error {
	funcName := "TaskService.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	userAccess, err := ts.storage.CheckAccess(ctx, (dto.CheckTaskAccessInfo)(info))
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if !userAccess {
		return apperrors.ErrUserNotInTask
	}
	logger.DebugFmt("user not in task", requestID.String(), funcName, nodeName)

	return ts.storage.RemoveUser(ctx, info)
}

// Move
// переносит задание в другой список
// или возвращает ошибки ...
func (ts TaskService) Move(ctx context.Context, taskMoveInfo dto.TaskMoveInfo) error {
	return ts.storage.Move(ctx, taskMoveInfo)
}
