package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type ITaskStorage interface {
	// Create
	// создает новую задачу по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTaskInfo) (*entities.Task, error)
	// Read
	// возвращает заание с привязанными пользователями
	// или возвращает ошибки ...
	Read(context.Context, dto.TaskID) (*entities.Task, error)
	// Update
	// обновляет задачу
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTaskInfo) error
	// Delete
	// удаляет задачу
	// или возвращает ошибки ...
	Delete(context.Context, dto.TaskID) error
}
