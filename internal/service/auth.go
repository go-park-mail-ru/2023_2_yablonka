package service

import (
	"context"
	"server/internal/pkg/entities"
	"time"
)

// Интерфейс для аутентификации
type IAuthService interface {
	// AuthUser
	// возвращает уникальную строку авторизации и её длительность
	AuthUser(context.Context, *entities.User) (string, time.Time, error)
	// VerifyAuth
	// проверяет состояние авторизации, возвращает ID авторизированного пользователя
	VerifyAuth(context.Context, string) (uint64, error)
	// GetLifetime
	// возвращает длительность авторизации
	GetLifetime() time.Duration
}
