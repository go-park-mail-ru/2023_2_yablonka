package csat_microservice

import (
	context "context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/storage"
)

type CSATAnswerService struct {
	storage storage.ICSATAnswerStorage
	UnimplementedCSATAnswerServiceServer
}

var CSATAnswerServiceErrorCodes = map[error]ErrorCode{
	nil:                             ErrorCode_OK,
	apperrors.ErrCouldNotBuildQuery: ErrorCode_COULD_NOT_BUILD_QUERY,
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
func (cs CSATAnswerService) Create(ctx context.Context, info *NewCSATAnswer) (*CreateResponse, error) {
	convertedInfo := dto.NewCSATAnswer{
		UserID:     info.UserID,
		QuestionID: info.QuestionID,
		Rating:     info.Rating,
	}
	response := &CreateResponse{}

	err := cs.storage.Create(ctx, convertedInfo)

	response.Code = CSATAnswerServiceErrorCodes[err]
	return response, nil
}
