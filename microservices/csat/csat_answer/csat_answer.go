package csat_microservice

import (
	context "context"
	"server/internal/apperrors"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/storage"

	"github.com/google/uuid"
)

type CSATAnswerService struct {
	storage storage.ICSATAnswerStorage
	logger  *logging.LogrusLogger
	UnimplementedCSATAnswerServiceServer
}

var CSATAnswerServiceErrorCodes = map[error]ErrorCode{
	nil:                             ErrorCode_OK,
	apperrors.ErrCouldNotBuildQuery: ErrorCode_COULD_NOT_BUILD_QUERY,
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewCSATAnswerService(storage storage.ICSATAnswerStorage, logger *logging.LogrusLogger) *CSATAnswerService {
	return &CSATAnswerService{
		storage: storage,
		logger:  logger,
	}
}

// Create
// создает новый ответ CSAT
// или возвращает ошибки ...
func (cs CSATAnswerService) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	convertedInfo := dto.NewCSATAnswer{
		UserID:     info.UserID,
		QuestionID: info.QuestionID,
		Rating:     info.Rating,
	}
	response := &CreateResponse{}

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	err := cs.storage.Create(sCtx, convertedInfo)

	response.Code = CSATAnswerServiceErrorCodes[err]
	return response, nil
}
