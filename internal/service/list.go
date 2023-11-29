package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для сервиса списков заданий
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IListService interface {
	// Create
	// создает новый список
	// или возвращает ошибки ...
	Create(context.Context, dto.NewListInfo) (*entities.List, error)
	// Update
	// обновляет список
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedListInfo) error
	// Delete
	// удаляет список по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.ListID) error
}
