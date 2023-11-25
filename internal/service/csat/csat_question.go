package list

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type CSATQuestionService struct {
	storage storage.ICSATQuestionStorage
}

// NewCSATQuestionService
// возвращает NewCSATQuestionService с инициализированным хранилищем
func NewCSATQuestionService(storage storage.ICSATQuestionStorage) *CSATQuestionService {
	return &CSATQuestionService{
		storage: storage,
	}
}

// Create
// создает новый список
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, info dto.NewCSATQuestionInfo) (*entities.CSAT, error) {
	verifiedInfo := dto.NewCSATQuestion{
		Content: info.Content,
	}
	return cs.storage.Create(ctx, verifiedInfo)
}

// Update
// обновляет список
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, info dto.UpdatedCSATInfo) error {
	return cs.storage.Update(ctx, info)
}

// Delete
// удаляет список по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id dto.CSATID) error {
	return cs.storage.Delete(ctx, id)
}
