package microservice

import (
	"context"
	"server/internal/apperrors"
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

var CSATQuestionServiceErrors = map[microservice.ErrorCode]error{
	microservice.ErrorCode_OK:                          nil,
	microservice.ErrorCode_COULD_NOT_BUILD_QUERY:       apperrors.ErrCouldNotBuildQuery,
	microservice.ErrorCode_COULD_NOT_CREATE_QUESTION:   apperrors.ErrCouldNotCreateQuestion,
	microservice.ErrorCode_COULD_NOT_GET_QUESTIONS:     apperrors.ErrCouldNotGetQuestions,
	microservice.ErrorCode_COULD_NOT_GET_QUESTION_TYPE: apperrors.ErrCouldNotGetQuestionType,
	microservice.ErrorCode_QUESTION_NOT_UPDATED:        apperrors.ErrQuestionNotUpdated,
	microservice.ErrorCode_QUESTION_NOT_DELETED:        apperrors.ErrQuestionNotDeleted,
	microservice.ErrorCode_ANSWER_RATING_TOO_BIG:       apperrors.ErrAnswerRatingTooBig,
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
	serverResponse, _ := cs.client.CheckRating(ctx, &microservice.NewCSATAnswerInfo{
		QuestionID: info.QuestionID,
		Rating:     info.Rating,
	})

	return CSATQuestionServiceErrors[serverResponse.Code]
}

// GetAll
// возвращает все вопросы CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) GetAll(ctx context.Context) (*[]dto.CSATQuestionFull, error) {
	serverResponse, _ := cs.client.GetAll(ctx, &emptypb.Empty{})
	if serverResponse.Code != microservice.ErrorCode_OK {
		return nil, CSATQuestionServiceErrors[serverResponse.Code]
	}

	convertedQuestions := []dto.CSATQuestionFull{}
	questions := serverResponse.Response
	for _, question := range questions.Questions {
		convertedQuestions = append(convertedQuestions, dto.CSATQuestionFull{
			ID:      question.ID,
			Content: question.Content,
			Type:    question.Type,
		})
	}
	return &convertedQuestions, CSATQuestionServiceErrors[serverResponse.Code]
}

// Create
// создает новый вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, info dto.NewCSATQuestionInfo) (*dto.CSATQuestionFull, error) {
	serverResponse, _ := cs.client.Create(ctx, &microservice.NewCSATQuestionInfo{
		Content: info.Content,
		Type:    info.Type,
	})
	if serverResponse.Code != microservice.ErrorCode_OK {
		return nil, CSATQuestionServiceErrors[serverResponse.Code]
	}

	question := serverResponse.Response

	convertedQuestion := dto.CSATQuestionFull{
		ID:      question.ID,
		Content: question.Content,
		Type:    question.Type,
	}

	return &convertedQuestion, CSATQuestionServiceErrors[serverResponse.Code]
}

// GetStats
// возвращает статистику по вопросам
// или возвращает ошибки ...
func (cs CSATQuestionService) GetStats(ctx context.Context) (*[]dto.QuestionWithStats, error) {
	serverResponse, _ := cs.client.GetStats(ctx, &emptypb.Empty{})
	if serverResponse.Code != microservice.ErrorCode_OK {
		return nil, CSATQuestionServiceErrors[serverResponse.Code]
	}

	convertedQuestions := []dto.QuestionWithStats{}
	allQuestions := serverResponse.Response
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
	return &convertedQuestions, CSATQuestionServiceErrors[serverResponse.Code]
}

// Update
// обновляет вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, info dto.UpdatedCSATQuestionInfo) error {
	serverResponse, _ := cs.client.Update(ctx, &microservice.UpdatedCSATQuestionInfo{
		ID:      info.ID,
		Content: info.Content,
		Type:    info.Type,
	})
	return CSATQuestionServiceErrors[serverResponse.Code]
}

// Delete
// удаляет вопрос CSAT по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	serverResponse, _ := cs.client.Delete(ctx, &microservice.CSATQuestionID{
		Value: id.Value,
	})
	return CSATQuestionServiceErrors[serverResponse.Code]
}
