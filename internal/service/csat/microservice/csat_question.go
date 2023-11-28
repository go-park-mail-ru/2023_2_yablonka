package microservice

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/csat/csat_question"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CSATQuestionService struct {
	storage storage.ICSATQuestionStorage
	client  microservice.CSATQuestionServiceClient
}

// NewCSATQuestionService
// возвращает NewCSATQuestionService с инициализированным хранилищем
func NewCSATQuestionService(storage storage.ICSATQuestionStorage, connection *grpc.ClientConn) *CSATQuestionService {
	client := microservice.NewCSATQuestionServiceClient(connection)
	return &CSATQuestionService{
		storage: storage,
		client:  client,
	}
}

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (cs CSATQuestionService) CheckRating(ctx context.Context, info dto.NewCSATAnswerInfo) error {
	_, err := cs.client.CheckRating(ctx, &microservice.NewCSATAnswerInfo{
		QuestionID: info.QuestionID,
		Rating:     info.Rating,
	})

	return err
}

// GetAll
// возвращает все вопросы CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) GetAll(ctx context.Context) (*[]dto.CSATQuestionFull, error) {
	questions, err := cs.client.GetAll(ctx, &emptypb.Empty{})
	convertedQuestions := []dto.CSATQuestionFull{}
	for _, question := range questions.Questions {
		convertedQuestions = append(convertedQuestions, dto.CSATQuestionFull{
			ID:      question.ID,
			Content: question.Content,
			Type:    question.Type,
		})
	}
	return &convertedQuestions, err
}

// Create
// создает новый вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, info dto.NewCSATQuestionInfo) (*dto.CSATQuestionFull, error) {
	question, err := cs.client.Create(ctx, &microservice.NewCSATQuestionInfo{
		Content: info.Content,
		Type:    info.Type,
	})
	if err != nil {
		return nil, err
	}

	convertedQuestion := dto.CSATQuestionFull{
		ID:      question.ID,
		Content: question.Content,
		Type:    question.Type,
	}

	return &convertedQuestion, nil
}

// GetStats
// возвращает статистику по вопросам
// или возвращает ошибки ...
func (cs CSATQuestionService) GetStats(ctx context.Context) (*[]dto.QuestionWithStats, error) {
	allQuestions, err := cs.client.GetStats(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	convertedQuestions := []dto.QuestionWithStats{}
	for _, question := range allQuestions.Questions {
		stats := []dto.RatingStats{}
		for _, stat := range question.Stats {
			stats = append(stats, dto.RatingStats{
				Rating:  stat.Rating,
				Count:   stat.Count,
				Average: stat.Average,
			})
		}
		convertedQuestions = append(convertedQuestions, dto.QuestionWithStats{
			ID:      question.ID,
			Content: question.Content,
			Type:    question.Type,
			Stats:   stats,
		})
	}
	return &convertedQuestions, err
}

// Update
// обновляет вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, info dto.UpdatedCSATQuestionInfo) error {
	_, err := cs.client.Update(ctx, &microservice.UpdatedCSATQuestionInfo{
		ID:      info.ID,
		Content: info.Content,
		Type:    info.Type,
	})
	return err
}

// Delete
// удаляет вопрос CSAT по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	_, err := cs.client.Delete(ctx, &microservice.CSATQuestionID{
		Value: id.Value,
	})
	return err
}
