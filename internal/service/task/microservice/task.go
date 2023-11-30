package task

import (
	"context"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

const nodeName = "service"

type TaskService struct {
	taskStorage storage.ITaskStorage
	userStorage storage.IUserStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewTaskService(ts storage.ITaskStorage, us storage.IUserStorage) *TaskService {
	return &TaskService{
		taskStorage: ts,
		userStorage: us,
	}
}

// Read
// возвращает задание
// или возвращает ошибки ...
func (ts TaskService) Read(ctx context.Context, id dto.TaskID) (*dto.SingleTaskInfo, error) {
	return ts.taskStorage.Read(ctx, id)
}

// Create
// создает новое задание
// или возвращает ошибки ...
func (ts TaskService) Create(ctx context.Context, info dto.NewTaskInfo) (*entities.Task, error) {
	return ts.taskStorage.Create(ctx, info)
}

// Update
// обновляет задание
// или возвращает ошибки ...
func (ts TaskService) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	return ts.taskStorage.Update(ctx, info)
}

// Delete
// удаляет задание
// или возвращает ошибки ...
func (ts TaskService) Delete(ctx context.Context, id dto.TaskID) error {
	return ts.taskStorage.Delete(ctx, id)
}

// AddUser
// добавляет пользователя в карточку
// или возвращает ошибки ...
func (ts TaskService) AddUser(ctx context.Context, info dto.AddTaskUserInfo) error {
	funcName := "TaskService.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	userAccess, err := ts.taskStorage.CheckAccess(ctx, (dto.CheckTaskAccessInfo)(info))
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.Debug("got user", funcName, nodeName)

	if userAccess {
		return apperrors.ErrUserAlreadyInBoard
	}
	logger.Debug("user not in task", funcName, nodeName)

	return ts.taskStorage.AddUser(ctx, info)
}

// RemoveUser
// удаляет пользователя из карточки
// или возвращает ошибки ...
func (ts TaskService) RemoveUser(ctx context.Context, info dto.RemoveTaskUserInfo) error {
	funcName := "TaskService.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	userAccess, err := ts.taskStorage.CheckAccess(ctx, (dto.CheckTaskAccessInfo)(info))
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.Debug("got user", funcName, nodeName)

	if !userAccess {
		return apperrors.ErrUserAlreadyInTask
	}
	logger.Debug("user not in task", funcName, nodeName)

	return ts.taskStorage.RemoveUser(ctx, info)
}
