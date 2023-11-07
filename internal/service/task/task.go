package task

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type TaskService struct {
	storage storage.ITaskStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewTaskService(storage storage.ITaskStorage) *TaskService {
	return &TaskService{
		storage: storage,
	}
}

// Create
// создает новое задание
// или возвращает ошибки ...
func (ts TaskService) Create(ctx context.Context, info dto.NewTaskInfo) (*entities.Task, error) {
	return nil, nil
}

// Update
// обновляет задание
// или возвращает ошибки ...
func (ts TaskService) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	return nil
}

// Delete
// удаляет задание
// или возвращает ошибки ...
func (ts TaskService) Delete(ctx context.Context, id dto.TaskID) error {
	return nil
}
