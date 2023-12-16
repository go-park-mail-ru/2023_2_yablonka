package microservice

import (
	"context"
	"server/internal/apperrors"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/csat/csat_question"

	"github.com/google/uuid"
	"google.golang.org/grpc"
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
	funcName := "CSATQuestionService.CheckRating"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.CheckRatingRequest{
		RequestID: requestID.String(),
		Value: &microservice.NewCSATAnswerInfo{
			QuestionID: info.QuestionID,
			Rating:     info.Rating,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.CheckRating(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return CSATQuestionServiceErrors[serverResponse.Code]
}

// GetAll
// возвращает все вопросы CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) GetAll(ctx context.Context) (*[]dto.CSATQuestionFull, error) {
	funcName := "CSATQuestionService.GetAll"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.GetAllRequest{
		RequestID: requestID.String(),
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.GetAll(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

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
	funcName := "CSATQuestionService.Create"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.CreateRequest{
		RequestID: requestID.String(),
		Value: &microservice.NewCSATQuestionInfo{
			Content: info.Content,
			Type:    info.Type,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.Create(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

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
	funcName := "CSATQuestionService.GetStats"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.GetStatsRequest{
		RequestID: requestID.String(),
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.GetStats(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

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
	funcName := "CSATQuestionService.Update"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.UpdateRequest{
		RequestID: requestID.String(),
		Value: &microservice.UpdatedCSATQuestionInfo{
			ID:      info.ID,
			Content: info.Content,
			Type:    info.Type,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.Update(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return CSATQuestionServiceErrors[serverResponse.Code]
}

// Delete
// удаляет вопрос CSAT по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	funcName := "CSATQuestionService.Delete"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.DeleteRequest{
		RequestID: requestID.String(),
		Value: &microservice.CSATQuestionID{
			Value: id.Value,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.Delete(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return CSATQuestionServiceErrors[serverResponse.Code]
}
