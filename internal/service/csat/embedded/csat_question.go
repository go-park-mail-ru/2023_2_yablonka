package embedded

import (
	"context"
	"server/internal/pkg/dto"
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

	if info.Rating > questionType.MaxRating {
		return err
	}

	return nil
}

// GetAll
// возвращает все вопросы CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) GetAll(ctx context.Context) (*[]dto.CSATQuestionFull, error) {
	return cs.storage.GetAll(ctx)
}

// Create
// создает новый вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, info dto.NewCSATQuestionInfo) (*dto.CSATQuestionFull, error) {
	questionType, err := cs.storage.GetQuestionTypeWithName(ctx, dto.CSATQuestionTypeName{Value: info.Type})
	if err != nil {
		return nil, err
	}
	verifiedInfo := dto.NewCSATQuestion{
		Content: info.Content,
		TypeID:  questionType.ID,
	}
	question, err := cs.storage.Create(ctx, verifiedInfo)
	if err != nil {
		return nil, err
	}
	question.Type = info.Type
	question.Content = info.Content
	return question, nil
}

// GetStats
// возвращает статистику по вопросам
// или возвращает ошибки ...
func (cs CSATQuestionService) GetStats(ctx context.Context) (*[]dto.QuestionWithStats, error) {
	return cs.storage.GetStats(ctx)
}

// Update
// обновляет вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, info dto.UpdatedCSATQuestionInfo) error {
	questionType, err := cs.storage.GetQuestionType(ctx, dto.CSATQuestionID{Value: info.ID})
	if err != nil {
		return nil
	}
	updatedQuestion := dto.UpdatedCSATQuestion{
		ID:      info.ID,
		Content: info.Content,
		Type:    questionType.ID,
	}
	return cs.storage.Update(ctx, updatedQuestion)
}

// Delete
// удаляет вопрос CSAT по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	return cs.storage.Delete(ctx, id)
}
