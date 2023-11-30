package microservice

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/storage"
)

type ChecklistItemService struct {
	storage storage.IChecklistItemStorage
}

// NewChecklistItemService
// возвращает ChecklistItemService с инициализированным хранилищем
func NewChecklistItemService(storage storage.IChecklistItemStorage) *ChecklistItemService {
	return &ChecklistItemService{
		storage: storage,
	}
}

// Create
// создает новый элемент чеклиста
// или возвращает ошибки ...
func (cls ChecklistItemService) Create(ctx context.Context, info dto.NewChecklistItemInfo) (*dto.ChecklistItemInfo, error) {
	return cls.storage.Create(ctx, info)
}

// Update
// обновляет элемент чеклиста
// или возвращает ошибки ...
func (cls ChecklistItemService) Update(ctx context.Context, info dto.UpdatedChecklistItemInfo) error {
	return cls.storage.Update(ctx, info)
}

// Delete
// удаляет элемент чеклиста по id
// или возвращает ошибки ...
func (cls ChecklistItemService) Delete(ctx context.Context, id dto.ChecklistItemID) error {
	return cls.storage.Delete(ctx, id)
}
