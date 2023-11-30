package service

import (
	"context"
	"server/internal/pkg/dto"
)

// Интерфейс для сервиса чеклистов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IChecklistService interface {
	// Create
	// создает новый чеклист в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewChecklistInfo) (*dto.ChecklistInfo, error)
	// Update
	// обновляет чеклист в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedChecklistInfo) error
	// Delete
	// удаляет чеклист в БД
	// или возвращает ошибки ...
	Delete(context.Context, dto.ChecklistID) error
}
