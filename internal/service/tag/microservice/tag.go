package tag

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type TagService struct {
	storage storage.ITagStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewTagService(ts storage.ITagStorage) *TagService {
	return &TagService{
		storage: ts,
	}
}

// Create
// создает новое задание
// или возвращает ошибки ...
func (ts TagService) Create(ctx context.Context, info dto.NewTagInfo) (*entities.Tag, error) {
	return ts.storage.Create(ctx, info)
}

// Update
// обновляет задание
// или возвращает ошибки ...
func (ts TagService) Update(ctx context.Context, info dto.UpdatedTagInfo) error {
	return ts.storage.Update(ctx, info)
}

// Delete
// удаляет задание
// или возвращает ошибки ...
func (ts TagService) Delete(ctx context.Context, id dto.TagID) error {
	return ts.storage.Delete(ctx, id)
}

// AddToTask
// добавляет тэг к заданию
// или возвращает ошибки ...
func (ts TagService) AddToTask(ctx context.Context, ids dto.TagAndTaskIDs) error {
	return ts.storage.AddToTask(ctx, ids)
}

// RemoveFromTask
// удаляет тэг из заданию
// или возвращает ошибки ...
func (ts TagService) RemoveFromTask(ctx context.Context, ids dto.TagAndTaskIDs) error {
	return ts.storage.RemoveFromTask(ctx, ids)
}
