package storage

import (
	"context"
	"server/internal/pkg/entities"
)

type IAuthStorage interface {
	// CreateSession
	// сохраняет сессию в хранилище, возвращает ID сесссии для куки
	// или возвращает ошибку apperrors.ErrTokenNotGenerated (500)
	CreateSession(ctx context.Context, session *entities.Session) error
	// GetSession
	// находит сессию по строке-токену
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	GetSession(ctx context.Context, token string) (*entities.Session, error)
	// DeleteSession
	// удаляет сессию по ID из хранилища, если она существует
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	DeleteSession(ctx context.Context, token string) error
}
