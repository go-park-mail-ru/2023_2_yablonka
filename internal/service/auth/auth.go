package service

import (
	"server/internal/pkg/entities"
	"server/internal/storage"
)

// NewAuthJWTService
// возвращает AuthJWTService с инициализированным JWT-секретом и датой истечения рефреш-токенов
func NewAuthJWTService(config *entities.ServerConfig) *AuthJWTService {
	return &AuthJWTService{
		jwtSecret:     []byte(config.JWTSecret),
		tokenLifetime: config.SessionDuration,
	}
}

// NewAuthSessionService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthSessionService(config *entities.ServerConfig, storage storage.IAuthStorage) *AuthSessionService {
	return &AuthSessionService{
		sessionDuration: config.SessionDuration,
		sessionIDLength: config.SessionIDLength,
		storage:         storage,
	}
}
