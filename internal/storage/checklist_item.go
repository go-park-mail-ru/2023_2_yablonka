package storage

import (
	"context"
	"server/internal/pkg/dto"
)

// Интерфейс для хранилища элемент чеклистаов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type IChecklistItemStorage interface {
	// Create
	// создает новый элемент чеклиста в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewChecklistItemInfo) (*dto.ChecklistItemInfo, error)
	// ReadMany
	// возвращает элементы чеклиста
	// или возвращает ошибки ...
	ReadMany(context.Context, dto.ChecklistItemIDs) (*[]dto.ChecklistItemInfo, error)
	// Update
	// обновляет элемент чеклиста в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedChecklistItemInfo) error
	// Delete
	// удаляет элемент чеклиста в БД
	// или возвращает ошибки ...
	Delete(context.Context, dto.ChecklistItemID) error
	// UpdateOrder
	// изменяет порядок вещей в чеклистах по id
	// или возвращает ошибки ...
	UpdateOrder(context.Context, dto.ChecklistItemIDs) error
}
