package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища списков
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type IListStorage interface {
	// Create
	// создает новый список задач в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewListInfo) (*entities.List, error)
	// Update
	// обновляет списсок задач в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedListInfo) error
	// Delete
	// удаляет списсок задач в БД
	// или возвращает ошибки ...
	Delete(context.Context, dto.ListID) error
}
