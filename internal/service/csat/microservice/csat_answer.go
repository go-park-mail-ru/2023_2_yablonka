package microservice

import (
	"context"
	"server/internal/apperrors"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/csat/csat_answer"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const nodeName = "service"

type CSATAnswerService struct {
	storage storage.ICSATAnswerStorage
	client  microservice.CSATAnswerServiceClient
}

var CSATAnswerServiceErrors = map[microservice.ErrorCode]error{
	microservice.ErrorCode_OK:                    nil,
	microservice.ErrorCode_COULD_NOT_BUILD_QUERY: apperrors.ErrCouldNotBuildQuery,
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewCSATAnswerService(storage storage.ICSATAnswerStorage, connection *grpc.ClientConn) *CSATAnswerService {
	client := microservice.NewCSATAnswerServiceClient(connection)
	return &CSATAnswerService{
		storage: storage,
		client:  client,
	}
}

// Create
// создает новый ответ CSAT
// или возвращает ошибки ...
func (cs CSATAnswerService) Create(ctx context.Context, info dto.NewCSATAnswer) error {
	funcName := "CSATAnswerService.Create"
	logger := ctx.Value(dto.LoggerKey).(logging.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.CreateRequest{
		RequestID: requestID.String(),
		Value: &microservice.NewCSATAnswer{
			UserID:     info.UserID,
			QuestionID: info.QuestionID,
			Rating:     info.Rating,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.Create(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return CSATAnswerServiceErrors[serverResponse.Code]
}
