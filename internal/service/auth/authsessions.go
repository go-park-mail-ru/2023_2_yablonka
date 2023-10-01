package service

import (
	"context"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"time"
)

type AuthSessionService struct {
	sessionDuration time.Duration
	storage         storage.IAuthStorage
}

func generateSession(user *entities.User, duration time.Duration) *entities.Session {
	return &entities.Session{
		UserID:     user.ID,
		ExpiryDate: time.Now().Add(duration),
	}
}

func (a *AuthSessionService) GetSessionDuration() time.Duration {
	return a.sessionDuration
}

// Возвращает ID сессии для полученного пользователя
func (a *AuthSessionService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	session := generateSession(user, a.sessionDuration)
	expiresAt := session.ExpiryDate
	sid, err := a.storage.CreateSession(session)
	if err != nil {
		return "", time.Time{}, err
	}
	return sid, expiresAt, nil
}

// VerifyAuth
// возвращает ID пользователя, которому принадлежит сессия
func (a *AuthSessionService) VerifyAuth(ctx context.Context, sessionString string) (uint64, error) {
	sessionObj, err := a.storage.GetSession(sessionString)
	if err != nil {
		return 0, err
	}
	return sessionObj.UserID, nil
}
