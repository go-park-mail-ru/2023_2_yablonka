package service

import (
	"context"
	"server/internal/pkg/dto"
)

// Интерфейс для сервиса элемента чеклиста
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IChecklistItemService interface {
	// Create
	// создает новый элемент элемент чеклистаа в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewChecklistItemInfo) (*dto.ChecklistItemInfo, error)
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
