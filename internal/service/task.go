package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для сервиса заданий
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ITaskService interface {
	// Create
	// создает новое задание
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTaskInfo) (*entities.Task, error)
	// Read
	// возвращает заание с привязанными пользователями
	// или возвращает ошибки ...
	Read(context.Context, dto.TaskID) (*dto.SingleTaskInfo, error)
	// Update
	// обновляет задание
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTaskInfo) error
	// Move
	// переносит задание в другой список
	// или возвращает ошибки ...
	Move(context.Context, dto.TaskMoveInfo) error
	// Delete
	// удаляет задание
	// или возвращает ошибки ...
	Delete(context.Context, dto.TaskID) error
	// AddUser
	// добавляет пользователя в карточку
	// или возвращает ошибки ...
	AddUser(context.Context, dto.AddTaskUserInfo) error
	// RemoveUser
	// удаляет пользователя из карточки
	// или возвращает ошибки ...
	RemoveUser(context.Context, dto.RemoveTaskUserInfo) error
	// Attach
	// добавляет файл в задание
	// или возвращает ошибки ...
	Attach(ctx context.Context, info dto.NewFileInfo) (*dto.AttachedFileInfo, error)
}
