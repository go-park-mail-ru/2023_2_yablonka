package service

import (
	"os"
	"server/internal/app/utils"
	"server/internal/apperrors"
	"server/internal/storage"
)

// NewAuthJWTService
// возвращает AuthJWTService с инициализированным JWT-секретом и датой истечения рефреш-токенов
func NewAuthJWTService() (*AuthJWTService, error) {
	tokenLifetime, err := utils.BuildSessionDuration()
	if err != nil {
		return nil, err
	}
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, apperrors.ErrJWTSecretMissing
	}
	return &AuthJWTService{
		jwtSecret:     []byte(secret),
		tokenLifetime: tokenLifetime,
	}, nil
}

// NewAuthSessionService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthSessionService(storage storage.IAuthStorage) (*AuthSessionService, error) {
	sessionDuration, err := utils.BuildSessionDuration()
	if err != nil {
		return nil, err
	}
	return &AuthSessionService{
		sessionDuration: sessionDuration,
		storage:         storage,
	}, nil
}
