package embedded

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/storage"
)

type CSATAnswerService struct {
	storage storage.ICSATAnswerStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewCSATAnswerService(storage storage.ICSATAnswerStorage) *CSATAnswerService {
	return &CSATAnswerService{
		storage: storage,
	}
}

// Create
// создает новый ответ CSAT
// или возвращает ошибки ...
func (cs CSATAnswerService) Create(ctx context.Context, info dto.NewCSATAnswer) error {
	return cs.storage.Create(ctx, info)
}
