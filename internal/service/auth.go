package service

import (
	"context"
	"time"
)

// Интерфейс для аутентификации
type IAuthService interface {
	// AuthUser
	// возвращает уникальную строку авторизации и её длительность
	// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
	// и вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	AuthUser(ctx context.Context, id uint64) (string, time.Time, error)
	// VerifyAuth
	// проверяет состояние авторизации, возвращает ID авторизированного пользователя
	// или вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	VerifyAuth(ctx context.Context, token string) (uint64, error)
	// LogOut
	// удаляет текущую сессию
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	LogOut(ctx context.Context, token string) error
	// GetLifetime
	// возвращает длительность авторизации
	GetLifetime() time.Duration
}
