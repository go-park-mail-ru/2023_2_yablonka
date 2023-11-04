package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"server/internal/apperrors"
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
func (a *AuthSessionService) AuthUser(ctx context.Context, id uint64) (string, time.Time, error) {
	expiresAt := time.Now().Add(a.sessionDuration)

	token, err := generateSessionID(a.sessionIDLength)
	if err != nil {
		return "", time.Time{}, apperrors.ErrTokenNotGenerated
	}

	session := &entities.Session{
		Token:      token,
		UserID:     id,
		ExpiryDate: expiresAt,
	}

	err = a.storage.CreateSession(ctx, session)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiresAt, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
<<<<<<< Updated upstream
func (a *AuthSessionService) VerifyAuth(ctx context.Context, sessionString string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(ctx, sessionString)
=======
<<<<<<< Updated upstream
func (a *AuthSessionService) VerifyAuth(ctx context.Context, token string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(ctx, token)
=======
<<<<<<< Updated upstream
func (a *AuthSessionService) VerifyAuth(ctx context.Context, sessionString string) (*dto.VerifiedAuthInfo, error) {
	sessionObj, err := a.storage.GetSession(ctx, sessionString)
=======
func (a *AuthSessionService) VerifyAuth(ctx context.Context, token string) (uint64, error) {
	sessionObj, err := a.storage.GetSession(ctx, token)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	if err != nil {
		return 0, err
	}
	if sessionObj.ExpiryDate.Before(time.Now()) {
		return 0, apperrors.ErrSessionExpired
	}
	return sessionObj.UserID, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthSessionService) LogOut(ctx context.Context, token string) error {
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
