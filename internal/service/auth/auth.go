package service

import (
	"server/internal/storage"
	"time"
)

// NewAuthJWTService
// возвращает AuthJWTService с инициализированным JWT-секретом и датой истечения рефреш-токенов
func NewAuthJWTService(JWTSecret string, sessionDuration time.Duration) *AuthJWTService {
	return &AuthJWTService{
		jwtSecret:     []byte(JWTSecret),
		tokenLifetime: sessionDuration,
	}
}

// NewAuthSessionService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthSessionService(sessionIDLength uint, sessionDuration time.Duration, storage storage.IAuthStorage) *AuthSessionService {
	return &AuthSessionService{
		sessionDuration: sessionDuration,
		sessionIDLength: sessionIDLength,
		storage:         storage,
	}
}
