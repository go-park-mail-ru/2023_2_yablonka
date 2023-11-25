package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для сервиса CSAT вопросов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ICSATQuestionService interface {
	// GetQuestionType
	// возвращает тип CSAT вопроса по его id
	// или возвращает ошибки ...
	GetQuestionType(context.Context, dto.CSATQuestionID) (*entities.QuestionType, error)
	// Create
	// создает новый список
	// или возвращает ошибки ...
	Create(context.Context, dto.NewListInfo) (*entities.CSATQuestion, error)
	// Update
	// обновляет список
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedListInfo) error
	// Delete
	// удаляет список по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.ListID) error
}
