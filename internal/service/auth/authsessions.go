package service

import (
	"context"
	"server/internal/pkg/dto"
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

// AuthUser
// возвращает ID сессии и срок её годности для полученного пользователя
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
func (a *AuthSessionService) VerifyAuth(ctx context.Context, sessionString string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(sessionString)
	if err != nil {
		return nil, err
	}
	return &dto.VerifiedAuthInfo{
		UserID: sessionObj.UserID,
	}, nil
}
