package storage

import (
	"context"
	"server/internal/pkg/entities"
)

type IAuthStorage interface {
	// CreateSession
	// сохраняет сессию в хранилище, возвращает ID сесссии для куки
	// или возвращает ошибку apperrors.ErrTokenNotGenerated (500)
<<<<<<< Updated upstream
	CreateSession(context.Context, *entities.Session, uint) (string, error)
=======
<<<<<<< Updated upstream
	CreateSession(ctx context.Context, session *entities.Session, sidLength uint) (string, error)
=======
<<<<<<< Updated upstream
	CreateSession(context.Context, *entities.Session, uint) (string, error)
=======
	CreateSession(ctx context.Context, session *entities.Session) error
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	// GetSession
	// находит сессию по строке-токену
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	GetSession(context.Context, string) (*entities.Session, error)
	// DeleteSession
	// удаляет сессию по ID из хранилища, если она существует
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	DeleteSession(context.Context, string) error
}
