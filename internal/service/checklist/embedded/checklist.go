package embedded

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/storage"
)

type ChecklistService struct {
	storage storage.IChecklistStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewChecklistService(storage storage.IChecklistStorage) *ChecklistService {
	return &ChecklistService{
		storage: storage,
	}
}

// Create
// создает новый чеклист
// или возвращает ошибки ...
func (cls ChecklistService) Create(ctx context.Context, info dto.NewChecklistInfo) (*dto.ChecklistInfo, error) {
	return cls.storage.Create(ctx, info)
}

// Update
// обновляет чеклист
// или возвращает ошибки ...
func (cls ChecklistService) Update(ctx context.Context, info dto.UpdatedChecklistInfo) error {
	return cls.storage.Update(ctx, info)
}

// Delete
// удаляет чеклист по id
// или возвращает ошибки ...
func (cls ChecklistService) Delete(ctx context.Context, id dto.ChecklistID) error {
	return cls.storage.Delete(ctx, id)
}
