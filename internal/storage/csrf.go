package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type ICSRFStorage interface {
	// CreateCSRF
	// сохраняет сессию в хранилище, возвращает токен
	// или возвращает ошибку apperrors.ErrTokenNotGenerated (500)
	Create(context.Context, *entities.CSRF) error
	// GetSession
	// находит сессию по токену
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	Get(context.Context, dto.CSRFToken) (*entities.CSRF, error)
	// DeleteSession
	// удаляет сессию по токену из хранилища, если она существует
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	Delete(context.Context, dto.CSRFToken) error
}
