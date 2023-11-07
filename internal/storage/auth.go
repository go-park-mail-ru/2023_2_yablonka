package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IAuthStorage interface {
	// CreateSession
	// сохраняет сессию в хранилище, возвращает ID сесссии для куки
	// или возвращает ошибку apperrors.ErrTokenNotGenerated (500)
	CreateSession(context.Context, *entities.Session) error
	// GetSession
	// находит сессию по строке-токену
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	GetSession(context.Context, dto.SessionToken) (*entities.Session, error)
	// DeleteSession
	// удаляет сессию по ID из хранилища, если она существует
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	DeleteSession(context.Context, dto.SessionToken) error
}
