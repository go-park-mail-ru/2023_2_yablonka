package microservice

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type ListService struct {
	storage storage.IListStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewListService(storage storage.IListStorage) *ListService {
	return &ListService{
		storage: storage,
	}
}

// Create
// создает новый список
// или возвращает ошибки ...
func (ls ListService) Create(ctx context.Context, info dto.NewListInfo) (*entities.List, error) {
	return ls.storage.Create(ctx, info)
}

// Update
// обновляет список
// или возвращает ошибки ...
func (ls ListService) Update(ctx context.Context, info dto.UpdatedListInfo) error {
	return ls.storage.Update(ctx, info)
}

// Delete
// удаляет список по id
// или возвращает ошибки ...
func (ls ListService) Delete(ctx context.Context, id dto.ListID) error {
	return ls.storage.Delete(ctx, id)
}
