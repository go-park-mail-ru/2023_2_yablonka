package auth_microservice

import (
	"server/internal/config"
	"server/internal/storage"
	"time"
)

type AuthService struct {
	authStorage     storage.IAuthStorage
	sessionDuration time.Duration
	sessionIDLength uint
	UnimplementedAuthServiceServer
}

func NewAuthService(config config.SessionConfig, authStorage storage.IAuthStorage) *AuthService {
	return &AuthService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		authStorage:     authStorage,
	}
}
