package service

import (
	"context"
	"server/internal/pkg/datatypes"
)

// Интерфейс для аутентификации (создание и проверка токена)
type IAuthService interface {
	AuthUser(context.Context, *datatypes.User) (string, error)
	VerifyAuth(context.Context, string) (string, error)
}
