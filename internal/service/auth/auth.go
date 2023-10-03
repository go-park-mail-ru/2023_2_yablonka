package service

import (
	jwt "server/internal/config/jwt"
	session "server/internal/config/session"
	"server/internal/storage"
)

// NewAuthJWTService
// возвращает AuthJWTService с инициализированным JWT-секретом и датой истечения рефреш-токенов
func NewAuthJWTService(config jwt.JWTServerConfig) *AuthJWTService {
	return &AuthJWTService{
		jwtSecret:     []byte(config.JWTSecret),
		tokenLifetime: config.Base.SessionDuration,
	}
}

// NewAuthSessionService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthSessionService(config session.SessionServerConfig, storage storage.IAuthStorage) *AuthSessionService {
	return &AuthSessionService{
		sessionDuration: config.Base.SessionDuration,
		sessionIDLength: config.SessionIDLength,
		storage:         storage,
	}
}
