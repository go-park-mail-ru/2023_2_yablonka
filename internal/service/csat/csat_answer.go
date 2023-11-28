package csat

import (
	"context"
	"server/internal/pkg/dto"
	embedded "server/internal/service/csat/embedded"
	micro "server/internal/service/csat/microservice"
	"server/internal/storage"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса ответов CSAT
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ICSATAnswerService interface {
	// Create
	// создает новый ответ CSAT
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCSATAnswer) error
}

func NewEmbeddedCSATAnswerService(answerStorage storage.ICSATAnswerStorage) *embedded.CSATAnswerService {
	return embedded.NewCSATAnswerService(answerStorage)
}

func NewMicroCSATAnswerService(answerStorage storage.ICSATAnswerStorage, connection *grpc.ClientConn) *micro.CSATAnswerService {
	return micro.NewCSATAnswerService(answerStorage, connection)
}
