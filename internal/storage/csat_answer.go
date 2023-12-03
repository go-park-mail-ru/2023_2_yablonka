package storage

import (
	"context"
	"server/internal/pkg/dto"
)

// Интерфейс для хранилища CSAT ответов
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type ICSATAnswerStorage interface {
	// Create
	// создает новый ответ CSAT опроса в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCSATAnswer) error
}
