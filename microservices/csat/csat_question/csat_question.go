package csat_microservice

import (
	context "context"
	"server/internal/pkg/dto"
	"server/internal/storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CSATQuestionService struct {
	storage storage.ICSATQuestionStorage
	UnimplementedCSATQuestionServiceServer
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
func (cs CSATQuestionService) CheckRating(ctx context.Context, info *NewCSATAnswerInfo) (*emptypb.Empty, error) {
	questionType, err := cs.storage.GetQuestionType(ctx, dto.CSATQuestionID{Value: info.QuestionID})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if info.Rating > questionType.MaxRating {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

// GetAll
// возвращает все вопросы CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) GetAll(ctx context.Context, empty *emptypb.Empty) (*AllQuestionStats, error) {
	questionStats, err := cs.storage.GetAll(ctx)
	if err != nil {
		return &AllQuestionStats{}, err
	}

	convertedQuestions := []*CSATQuestionFull{}
	for _, question := range *questionStats {
		convertedQuestions = append(convertedQuestions, &CSATQuestionFull{
			ID:      question.ID,
			Type:    question.Type,
			Content: question.Content,
		})
	}
	convertedStats := &AllQuestionStats{
		Questions: convertedQuestions,
	}

	return convertedStats, nil
}

// GetStats
// возвращает статистику по вопросам
// или возвращает ошибки ...
func (cs CSATQuestionService) GetStats(ctx context.Context, empty *emptypb.Empty) (*AllQuestionsWithStats, error) {
	stats, err := cs.storage.GetStats(ctx)
	if err != nil {
		return &AllQuestionsWithStats{}, err
	}

	convertedQuestions := []*QuestionWithStats{}
	for _, question := range *stats {
		convertedRatings := []*RatingStats{}
		for _, rating := range question.Stats {
			convertedRatings = append(convertedRatings, &RatingStats{
				Rating:  rating.Rating,
				Average: rating.Average,
				Count:   rating.Count,
			})
		}
		convertedQuestions = append(convertedQuestions, &QuestionWithStats{
			ID:      question.ID,
			Type:    question.Type,
			Content: question.Content,
			Stats:   convertedRatings,
		})
	}
	convertedStats := AllQuestionsWithStats{
		Questions: convertedQuestions,
	}

	return &convertedStats, nil
}

// Create
// создает новый вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, info *NewCSATQuestionInfo) (*CSATQuestionFull, error) {
	questionType, err := cs.storage.GetQuestionTypeWithName(ctx, dto.CSATQuestionTypeName{Value: info.Type})
	if err != nil {
		return &CSATQuestionFull{}, err
	}
	verifiedInfo := dto.NewCSATQuestion{
		Content: info.Content,
		TypeID:  questionType.ID,
	}
	question, err := cs.storage.Create(ctx, verifiedInfo)
	if err != nil {
		return &CSATQuestionFull{}, err
	}
	question.Type = info.Type
	convertedQuestion := &CSATQuestionFull{
		ID:      question.ID,
		Type:    question.Type,
		Content: question.Content,
	}
	return convertedQuestion, nil
}

// Update
// обновляет вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, info *UpdatedCSATQuestionInfo) (*emptypb.Empty, error) {
	questionType, err := cs.storage.GetQuestionType(ctx, dto.CSATQuestionID{Value: info.ID})
	if err != nil {
		return &emptypb.Empty{}, nil
	}
	updatedQuestion := dto.UpdatedCSATQuestion{
		ID:      info.ID,
		Content: info.Content,
		Type:    questionType.ID,
	}
	return &emptypb.Empty{}, cs.storage.Update(ctx, updatedQuestion)
}

// Delete
// удаляет вопрос CSAT по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id *CSATQuestionID) (*emptypb.Empty, error) {
	convertedID := dto.CSATQuestionID{
		Value: id.Value,
	}
	return &emptypb.Empty{}, cs.storage.Delete(ctx, convertedID)
}
