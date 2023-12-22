package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для сервиса тэг
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ITagService interface {
	// Create
	// создает новый тэг
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTagInfo) (*entities.Tag, error)
	// Update
	// обновляет тэг
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTagInfo) error
	// Delete
	// удаляет тэг по id
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
