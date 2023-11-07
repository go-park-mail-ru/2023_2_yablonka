package list

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
// создает новоt рабочее пространство в БД по данным
// или возвращает ошибки ...
func (ls ListService) Create(ctx context.Context, info dto.NewListInfo) (*entities.List, error) {
	return nil, nil
}

// Update
// обновляет рабочее пространство в БД
// или возвращает ошибки ...
func (ls ListService) Update(ctx context.Context, info dto.UpdatedListInfo) error {
	return nil
}

// Delete
// удаляет рабочее пространство в БД по id
// или возвращает ошибки ...
func (ls ListService) Delete(ctx context.Context, id dto.ListID) error {
	return nil
}
