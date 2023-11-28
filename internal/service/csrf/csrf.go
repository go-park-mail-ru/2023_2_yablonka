package csrf

import (
	"context"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	"time"

	embedded "server/internal/service/csrf/embedded"
	micro "server/internal/service/csrf/microservice"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса проверки на CSRF
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ICSRFService interface {
	// AuthUser
	// возвращает уникальную строку CSRF и её длительность
	// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
	// и вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	SetupCSRF(context.Context, dto.UserID) (dto.CSRFData, error)
	// VerifyCSRF
	// проверяет состояние CSRF, возвращает ID авторизированного пользователя
	// или вариации GenericUnauthorizedResponse (401) в зависимости от имплементации
	VerifyCSRF(context.Context, dto.CSRFToken) error
	// LogOut
	// удаляет текущую сессию CSRF
	// или возвращает ошибку apperrors.ErrSessionNotFound (401)
	DeleteCSRF(context.Context, dto.CSRFToken) error
	// GetLifetime
	// возвращает длительность авторизации
	GetLifetime(context.Context) time.Duration
}

func NewEmbeddedCSRFService(csrfStorage storage.ICSRFStorage, config config.SessionConfig) *embedded.CSRFService {
	return embedded.NewCSRFService(config, csrfStorage)
}

// TODO: Board microservice
func NewMicroCSRFService(csrfStorage storage.ICSRFStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.CSRFService {
	return micro.NewCSRFService(config, csrfStorage)
}
