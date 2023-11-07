package service

import (
	"context"
	"server/internal/pkg/dto"
	"time"
)

// Интерфейс для аутентификации
type IAuthService interface {
	// AuthUser
	// возвращает уникальную строку авторизации и её длительность
	// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
	// и вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	AuthUser(context.Context, dto.UserID) (dto.SessionToken, time.Time, error)
	// VerifyAuth
	// проверяет состояние авторизации, возвращает ID авторизированного пользователя
	// или вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	VerifyAuth(context.Context, dto.SessionToken) (dto.UserID, error)
	// LogOut
	// удаляет текущую сессию
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	LogOut(context.Context, dto.SessionToken) error
	// GetLifetime
	// возвращает длительность авторизации
	GetLifetime() time.Duration
}
