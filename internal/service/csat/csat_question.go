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

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (cs CSATQuestionService) CheckRating(ctx context.Context, info dto.NewCSATAnswerInfo) error {
	questionType, err := cs.storage.GetQuestionType(ctx, dto.CSATQuestionID{Value: info.QuestionID})
	if err != nil {
		return nil
	}

	if questionType.MaxRating > questionType.MaxRating {
		return err
	}
	return nil
}

// Create
// создает новый список
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, info dto.NewCSATQuestionInfo) (*entities.CSATQuestion, error) {
	verifiedInfo := dto.NewCSATQuestion{
		Content: info.Content,
	}
	return cs.storage.Create(ctx, verifiedInfo)
}

// Update
// обновляет список
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, info dto.UpdatedCSATQuestion) error {
	return cs.storage.Update(ctx, info)
}

// Delete
// удаляет список по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	return cs.storage.Delete(ctx, id)
}
