package csat_microservice

import (
	context "context"
	"server/internal/apperrors"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/storage"

	"github.com/google/uuid"
)

type CSATQuestionService struct {
	storage storage.ICSATQuestionStorage
	logger  *logging.LogrusLogger
	UnimplementedCSATQuestionServiceServer
}

var CSATQuestionServiceErrorCodes = map[error]ErrorCode{
	nil:                                  ErrorCode_OK,
	apperrors.ErrCouldNotBuildQuery:      ErrorCode_COULD_NOT_BUILD_QUERY,
	apperrors.ErrCouldNotCreateQuestion:  ErrorCode_COULD_NOT_CREATE_QUESTION,
	apperrors.ErrCouldNotGetQuestions:    ErrorCode_COULD_NOT_GET_QUESTIONS,
	apperrors.ErrCouldNotGetQuestionType: ErrorCode_COULD_NOT_GET_QUESTION_TYPE,
	apperrors.ErrQuestionNotUpdated:      ErrorCode_QUESTION_NOT_UPDATED,
	apperrors.ErrQuestionNotDeleted:      ErrorCode_QUESTION_NOT_DELETED,
	apperrors.ErrAnswerRatingTooBig:      ErrorCode_ANSWER_RATING_TOO_BIG,
}

// NewCSATQuestionService
// возвращает NewCSATQuestionService с инициализированным хранилищем
func NewCSATQuestionService(storage storage.ICSATQuestionStorage, logger *logging.LogrusLogger) *CSATQuestionService {
	return &CSATQuestionService{
		storage: storage,
		logger:  logger,
	}
}

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (cs CSATQuestionService) CheckRating(ctx context.Context, request *CheckRatingRequest) (*CheckRatingResponse, error) {
	response := &CheckRatingResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	questionType, err := cs.storage.GetQuestionType(sCtx, dto.CSATQuestionID{Value: info.QuestionID})
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		return response, nil
	}

	if info.Rating > questionType.MaxRating {
		response.Code = CSATQuestionServiceErrorCodes[apperrors.ErrAnswerRatingTooBig]
		return response, nil
	}

	response.Code = CSATQuestionServiceErrorCodes[nil]

	return response, nil
}

// GetAll
// возвращает все вопросы CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) GetAll(ctx context.Context, request *GetAllRequest) (*GetAllResponse, error) {
	response := &GetAllResponse{}
	requestID, _ := uuid.Parse(request.RequestID)

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	questionStats, err := cs.storage.GetAll(sCtx)
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		response.Response = &AllQuestionStats{}
		return response, nil
	}

	convertedQuestions := []*CSATQuestionFull{}
	for _, question := range *questionStats {
		convertedQuestions = append(convertedQuestions, &CSATQuestionFull{
			ID:      question.ID,
			Type:    question.Type,
			Content: question.Content,
		})
	}
	response.Code = CSATQuestionServiceErrorCodes[nil]
	response.Response = &AllQuestionStats{
		Questions: convertedQuestions,
	}

	return response, nil
}

// GetStats
// возвращает статистику по вопросам
// или возвращает ошибки ...
func (cs CSATQuestionService) GetStats(ctx context.Context, request *GetStatsRequest) (*GetStatsResponse, error) {
	response := &GetStatsResponse{}
	requestID, _ := uuid.Parse(request.RequestID)

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	stats, err := cs.storage.GetStats(sCtx)
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		response.Response = &AllQuestionsWithStats{}
		return response, nil
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
	response.Response = &AllQuestionsWithStats{
		Questions: convertedQuestions,
	}
	response.Code = CSATQuestionServiceErrorCodes[nil]

	return response, nil
}

// Create
// создает новый вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	response := &CreateResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	questionType, err := cs.storage.GetQuestionTypeWithName(sCtx, dto.CSATQuestionTypeName{Value: info.Type})
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		response.Response = &CSATQuestionFull{}
		return response, nil
	}
	verifiedInfo := dto.NewCSATQuestion{
		Content: info.Content,
		TypeID:  questionType.ID,
	}
	question, err := cs.storage.Create(sCtx, verifiedInfo)
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		response.Response = &CSATQuestionFull{}
		return response, nil
	}

	question.Type = info.Type
	response.Code = CSATQuestionServiceErrorCodes[nil]
	response.Response = &CSATQuestionFull{
		ID:      question.ID,
		Type:    question.Type,
		Content: question.Content,
	}

	return response, nil
}

// Update
// обновляет вопрос CSAT
// или возвращает ошибки ...
func (cs CSATQuestionService) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, error) {
	response := &UpdateResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	questionType, err := cs.storage.GetQuestionType(sCtx, dto.CSATQuestionID{Value: info.ID})
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		return response, nil
	}

	updatedQuestion := dto.UpdatedCSATQuestion{
		ID:      info.ID,
		Content: info.Content,
		Type:    questionType.ID,
	}
	err = cs.storage.Update(ctx, updatedQuestion)
	response.Code = CSATQuestionServiceErrorCodes[err]

	return response, nil
}

// Delete
// удаляет вопрос CSAT по id
// или возвращает ошибки ...
func (cs CSATQuestionService) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	response := &DeleteResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	id := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	convertedID := dto.CSATQuestionID{
		Value: id.Value,
	}
	err := cs.storage.Delete(sCtx, convertedID)
	response.Code = CSATQuestionServiceErrorCodes[err]

	return response, nil
}
