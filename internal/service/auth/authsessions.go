package service

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"time"
)

type AuthSessionService struct {
	sessionDuration time.Duration
	sessionIDLength uint
	storage         storage.IAuthStorage
}

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (a *AuthSessionService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	session := &entities.Session{
		UserID:     user.ID,
		ExpiryDate: time.Now().Add(a.sessionDuration),
	}
	expiresAt := session.ExpiryDate
	sessionId, err := a.storage.CreateSession(ctx, session, a.sessionIDLength)
	if err != nil {
		return "", time.Time{}, err
	}
	return sessionId, expiresAt, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthSessionService) VerifyAuth(ctx context.Context, token string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(ctx, token)
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
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthSessionService) LogOut(ctx context.Context, sessionString string) error {
	_, err := a.storage.GetSession(ctx, sessionString)
	if err != nil {
		return err
	}
	return a.storage.DeleteSession(ctx, sessionString)
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthSessionService) GetLifetime() time.Duration {
	return a.sessionDuration
}
