package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"time"
)

// AuthJWTService
// структура сервиса аутентификации с помощью создания сессий
// содержит интерфейс для работы с БД и длительность сессии
type AuthSessionService struct {
	sessionDuration time.Duration
	storage         storage.IAuthStorage
}

// AuthUser
// Возвращает ID сессии её длительность
func (a *AuthSessionService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	session := &entities.Session{
		UserID:     user.ID,
		ExpiryDate: time.Now().Add(a.sessionDuration),
	}
	expiresAt := session.ExpiryDate
	sessionId, err := a.storage.CreateSession(session)
	if err != nil {
		return "", time.Time{}, err
	}
	return sessionId, expiresAt, nil
}

// VerifyAuth
// возвращает ID пользователя, которому принадлежит сессия
func (a *AuthSessionService) VerifyAuth(ctx context.Context, sessionString string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(sessionString)
	if err != nil {
		return nil, err
	}
	return &dto.VerifiedAuthInfo{
		UserID: sessionObj.UserID,
	}, nil
}

// GetLifetime
// возвращает длительность сессии
func (a *AuthSessionService) GetLifetime() time.Duration {
	return a.sessionDuration
}
