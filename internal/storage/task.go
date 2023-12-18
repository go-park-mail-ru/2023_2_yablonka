package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища заданий
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type ITaskStorage interface {
	// Create
	// создает новую задачу по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTaskInfo) (*entities.Task, error)
	// Read
	// возвращает заание с привязанными пользователями
	// или возвращает ошибки ...
	Read(context.Context, dto.TaskID) (*dto.SingleTaskInfo, error)
	// ReadMany
	// возвращает задания
	// или возвращает ошибки ...
	ReadMany(context.Context, dto.TaskIDs) (*[]dto.SingleTaskInfo, error)
	// CheckAccess
	// находит пользователя в задании
	// или возвращает ошибки ...
	CheckAccess(context.Context, dto.CheckTaskAccessInfo) (bool, error)
	// Update
	// обновляет задачу
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTaskInfo) error
	// Delete
	// удаляет задачу
	// или возвращает ошибки ...
	Delete(context.Context, dto.TaskID) error
	// Move
	// переносит задание в другой список
	// или возвращает ошибки ...
	Move(context.Context, dto.TaskMoveInfo) error
	// AddUser
	// добавляет пользователя в карточку
	// или возвращает ошибки ...
	AddUser(context.Context, dto.AddTaskUserInfo) error
	// RemoveUser
	// удаляет пользователя из карточки
	// или возвращает ошибки ...
	RemoveUser(context.Context, dto.RemoveTaskUserInfo) error
	// AttachFile
	// добавляет файл в задание
	// или возвращает ошибки ...
	AttachFile(context.Context, dto.AttachedFileInfo) error
}
