package service

import (
	"context"
	"server/internal/pkg/dto"
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
	// Create
	// создает новый список
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCSATQuestionInfo) (*dto.CSATQuestionFull, error)
	// Update
	// обновляет список
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedCSATQuestion) error
	// Delete
	// удаляет список по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.CSATQuestionID) error
}
