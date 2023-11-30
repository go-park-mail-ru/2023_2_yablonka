package storage

import (
	"context"
	"server/internal/pkg/dto"
)

// Интерфейс для хранилища чеклистов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type IChecklistStorage interface {
	// Create
	// создает новый чеклист в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewChecklistInfo) (*dto.ChecklistInfo, error)
	// ReadMany
	// возвращает чеклисты
	// или возвращает ошибки ...
	ReadMany(context.Context, dto.ChecklistIDs) (*[]dto.ChecklistInfo, error)
	// Update
	// обновляет чеклист в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedChecklistInfo) error
	// Delete
	// удаляет чеклист в БД
	// или возвращает ошибки ...
	Delete(context.Context, dto.ChecklistID) error
}
