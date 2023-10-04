package service

import (
	"context"
	"server/internal/apperrors"
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
	sessionIDLength uint
	storage         storage.IAuthStorage
}

// AuthUser
// возвращает ID сессии её длительность
func (a *AuthSessionService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	session := &entities.Session{
		UserID:     user.ID,
		ExpiryDate: time.Now().Add(a.sessionDuration),
	}
	authCtx := context.WithValue(ctx, dto.SIDLengthKey, a.sessionIDLength)
	expiresAt := session.ExpiryDate
	sessionId, err := a.storage.CreateSession(authCtx, session)
	if err != nil {
		return "", time.Time{}, err
	}
	return sessionId, expiresAt, nil
}

// VerifyAuth
// возвращает ID пользователя, которому принадлежит сессия
func (a *AuthSessionService) VerifyAuth(ctx context.Context, sessionString string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(ctx, sessionString)
	if err != nil {
		return nil, err
	}
	if sessionObj.ExpiryDate.Before(time.Now()) {
		return nil, apperrors.ErrSessionExpired
	}
	return &dto.VerifiedAuthInfo{
		UserID: sessionObj.UserID,
	}, nil
}

// LogOut
// удаляет сессию пользователя из хранилища, если она существует
func (a *AuthSessionService) LogOut(ctx context.Context, sessionString string) error {
	_, err := a.storage.GetSession(ctx, sessionString)
	if err != nil {
		return err
	}
	return a.storage.DeleteSession(ctx, sessionString)
}

// GetLifetime
// возвращает длительность сессии
func (a *AuthSessionService) GetLifetime() time.Duration {
	return a.sessionDuration
}
