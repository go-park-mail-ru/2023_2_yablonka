package microservice

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/csat/csat_answer"

	"google.golang.org/grpc"
)

type CSATAnswerService struct {
	storage storage.ICSATAnswerStorage
	client  microservice.CSATAnswerServiceClient
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
	_, err := cs.client.Create(ctx, &microservice.NewCSATAnswer{
		UserID:     info.UserID,
		QuestionID: info.QuestionID,
		Rating:     info.Rating,
	})
	return err
}
