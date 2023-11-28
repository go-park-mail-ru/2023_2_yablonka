package csat

import (
	"context"
	"server/internal/pkg/dto"
	embedded "server/internal/service/csat/embedded"
	micro "server/internal/service/csat/microservice"
	"server/internal/storage"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса CSAT вопросов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ICSATQuestionService interface {
	// CheckRating
	// проверка рейтинга на соответствие ограничениям типа
	// или возвращает ошибки ...
	CheckRating(context.Context, dto.NewCSATAnswerInfo) error
	// GetAll
	// возвращает все вопросы
	// или возвращает ошибки ...
	GetAll(context.Context) (*[]dto.CSATQuestionFull, error)
	// GetStats
	// возвращает статистику по вопросам
	// или возвращает ошибки ...
	GetStats(context.Context) (*[]dto.QuestionWithStats, error)
	// Create
	// создает новый список
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCSATQuestionInfo) (*dto.CSATQuestionFull, error)
	// Update
	// обновляет список
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedCSATQuestionInfo) error
	// Delete
	// удаляет список по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.CSATQuestionID) error
}

func NewEmbeddedCSATQuestionService(quuestionStorage storage.ICSATQuestionStorage) *embedded.CSATQuestionService {
	return embedded.NewCSATQuestionService(quuestionStorage)
}

func NewMicroCSATQuestionService(quuestionStorage storage.ICSATQuestionStorage, connection *grpc.ClientConn) *micro.CSATQuestionService {
	return micro.NewCSATQuestionService(quuestionStorage, connection)
}
