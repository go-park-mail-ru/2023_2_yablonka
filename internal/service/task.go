package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type ITaskService interface {
	// Create
	// создает новое задание
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTaskInfo) (*entities.Task, error)
	// Update
	// обновляет задание
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTaskInfo) error
	// Delete
	// удаляет задание
	// или возвращает ошибки ...
	Delete(context.Context, dto.TaskID) error
}
