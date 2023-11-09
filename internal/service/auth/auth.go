package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"server/internal/apperrors"
	config "server/internal/config"
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

// NewAuthService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthService(config config.SessionConfig, storage storage.IAuthStorage) *AuthSessionService {
	return &AuthSessionService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		storage:         storage,
	}
}

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (a *AuthSessionService) AuthUser(ctx context.Context, id dto.UserID) (dto.SessionToken, time.Time, error) {
	expiresAt := time.Now().Add(a.sessionDuration)

	token, err := generateSessionID(a.sessionIDLength)
	if err != nil {
		return dto.SessionToken{}, time.Time{}, apperrors.ErrTokenNotGenerated
	}

	session := &entities.Session{
		Token:      token,
		UserID:     id.Value,
		ExpiryDate: expiresAt,
	}

	err = a.storage.CreateSession(ctx, session)
	if err != nil {
		return dto.SessionToken{}, time.Time{}, err
	}

	return dto.SessionToken{Value: token}, expiresAt, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthSessionService) VerifyAuth(ctx context.Context, token dto.SessionToken) (dto.UserID, error) {
	sessionObj, err := a.storage.GetSession(ctx, token)
	if err != nil {
		return dto.UserID{}, err
	}

	if sessionObj.ExpiryDate.Before(time.Now()) {
		return dto.UserID{}, apperrors.ErrSessionExpired
	}
	return dto.UserID{Value: sessionObj.UserID}, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthSessionService) LogOut(ctx context.Context, token dto.SessionToken) error {
	return a.storage.DeleteSession(ctx, token)
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthSessionService) GetLifetime() time.Duration {
	return a.sessionDuration
}

// GenerateSessionID
// возвращает alphanumeric строку, собранную криптографически безопасным PRNG
func generateSessionID(n uint) (string, error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]rune, n)
	for i := range buf {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			return "", err
		}
		buf[i] = letterRunes[j.Int64()]
	}
	return string(buf), nil
}
