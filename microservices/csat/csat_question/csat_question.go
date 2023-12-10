package csat_microservice

import (
	context "context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CSATQuestionService struct {
	storage storage.ICSATQuestionStorage
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
func NewCSATQuestionService(storage storage.ICSATQuestionStorage) *CSATQuestionService {
	return &CSATQuestionService{
		storage: storage,
	}
}

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (cs CSATQuestionService) CheckRating(ctx context.Context, info *NewCSATAnswerInfo) (*CheckRatingResponse, error) {
	response := &CheckRatingResponse{}

	questionType, err := cs.storage.GetQuestionType(ctx, dto.CSATQuestionID{Value: info.QuestionID})
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
func (cs CSATQuestionService) GetAll(ctx context.Context, empty *emptypb.Empty) (*GetAllResponse, error) {
	response := &GetAllResponse{}

	questionStats, err := cs.storage.GetAll(ctx)
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
func (cs CSATQuestionService) GetStats(ctx context.Context, empty *emptypb.Empty) (*GetStatsResponse, error) {
	response := &GetStatsResponse{}

	stats, err := cs.storage.GetStats(ctx)
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
func (cs CSATQuestionService) Create(ctx context.Context, info *NewCSATQuestionInfo) (*CreateResponse, error) {
	response := &CreateResponse{}

	questionType, err := cs.storage.GetQuestionTypeWithName(ctx, dto.CSATQuestionTypeName{Value: info.Type})
	if err != nil {
		response.Code = CSATQuestionServiceErrorCodes[err]
		response.Response = &CSATQuestionFull{}
		return response, nil
	}
	verifiedInfo := dto.NewCSATQuestion{
		Content: info.Content,
		TypeID:  questionType.ID,
	}
	question, err := cs.storage.Create(ctx, verifiedInfo)
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
func (cs CSATQuestionService) Update(ctx context.Context, info *UpdatedCSATQuestionInfo) (*UpdateResponse, error) {
	response := &UpdateResponse{}

	questionType, err := cs.storage.GetQuestionType(ctx, dto.CSATQuestionID{Value: info.ID})
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
func (cs CSATQuestionService) Delete(ctx context.Context, id *CSATQuestionID) (*DeleteResponse, error) {
	response := &DeleteResponse{}

	convertedID := dto.CSATQuestionID{
		Value: id.Value,
	}
	err := cs.storage.Delete(ctx, convertedID)
	response.Code = CSATQuestionServiceErrorCodes[err]

	return response, nil
}
