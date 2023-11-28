package auth

import (
	"context"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	"time"

	embedded "server/internal/service/auth/embedded"
	micro "server/internal/service/auth/microservice"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса аутентификации
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IAuthService interface {
	// AuthUser
	// возвращает уникальную строку авторизации и её длительность
	// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
	// и вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	AuthUser(context.Context, dto.UserID) (dto.SessionToken, error)
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

func NewEmbeddedAuthService(authStorage storage.IAuthStorage, config config.SessionConfig) *embedded.AuthService {
	return embedded.NewAuthService(config, authStorage)
}

func NewMicroAuthService(authStorage storage.IAuthStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.AuthService {
	return micro.NewAuthService(config, authStorage, connection)
}
