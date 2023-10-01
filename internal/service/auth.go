package service

import (
	"context"
	"server/internal/pkg/entities"
	"time"
)

// Интерфейс для аутентификации
type IAuthService interface {
	AuthUser(context.Context, *entities.User) (string, time.Time, error)
	VerifyAuth(context.Context, string) (uint64, error)
}
