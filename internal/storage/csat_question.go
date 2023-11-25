package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища CSAT вопросов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type ICSATQuestionStorage interface {
	// GetQuestionType
	// возвращает тип CSAT вопроса по его id
	// или возвращает ошибки ...
	GetQuestionType(context.Context, dto.CSATQuestionID) (*entities.QuestionType, error)
	// Create
	// создает новый CSAT вопрос в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCSATQuestion) (*entities.CSATQuestion, error)
	// Update
	// обновляет CSAT вопрос в БД
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedCSATQuestion) error
	// Delete
	// удаляет CSAT вопрос в БД
	// или возвращает ошибки ...
	Delete(context.Context, dto.CSATQuestionID) error
}
