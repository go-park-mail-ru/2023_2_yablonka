package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища тэгов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type ITagStorage interface {
	// Create
	// создает новый тэг в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTagInfo) (*entities.Tag, error)
	// Update
	// обновляет списсок задач в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTagInfo) error
	// Delete
	// удаляет списсок задач в БД
	// или возвращает ошибки ...
	Delete(context.Context, dto.TagID) error
	// AddToTask
	// добавляет тэг к заданию
	// или возвращает ошибки ...
	AddToTask(context.Context, dto.TagAndTaskIDs) error
	// RemoveFromTask
	// удаляет тэг из заданию
	// или возвращает ошибки ...
	RemoveFromTask(context.Context, dto.TagAndTaskIDs) error
}
