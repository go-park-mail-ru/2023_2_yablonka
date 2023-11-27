package service

import (
	"context"
	"server/internal/pkg/dto"
)

// Интерфейс для сервиса ответов CSAT
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ICSATSAnswerService interface {
	// Create
	// создает новый ответ CSAT
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCSATAnswer) error
}
