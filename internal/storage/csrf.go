package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища CSRF
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type ICSRFStorage interface {
	// Create
	// сохраняет CSRF в хранилище, возвращает токен
	// или возвращает ошибку apperrors.ErrTokenNotGenerated (500)
	Create(context.Context, *entities.CSRF) error
	// Get
	// находит CSRF по токену
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	Get(context.Context, dto.CSRFToken) (*entities.CSRF, error)
	// Delete
	// удаляет CSRF по токену из хранилища, если она существует
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	Delete(context.Context, dto.CSRFToken) error
}
