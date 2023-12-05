package embedded

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

	logger "server/internal/logging"
)

type AuthService struct {
	sessionDuration time.Duration
	sessionIDLength uint
	authStorage     storage.IAuthStorage
}

const nodeName = "service"

// NewAuthService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthService(config config.SessionConfig, authStorage storage.IAuthStorage) *AuthService {
	return &AuthService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		authStorage:     authStorage,
	}
}

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (a *AuthService) AuthUser(ctx context.Context, id dto.UserID) (dto.SessionToken, error) {
	expiresAt := time.Now().Add(a.sessionDuration)

	sessionID, err := generateString(a.sessionIDLength)
	if err != nil {
		return dto.SessionToken{}, apperrors.ErrTokenNotGenerated
	}

	session := &entities.Session{
		SessionID:  sessionID,
		UserID:     id.Value,
		ExpiryDate: expiresAt,
	}

	err = a.authStorage.CreateSession(ctx, session)
	if err != nil {
		return dto.SessionToken{}, err
	}

	return dto.SessionToken{
		ID:             sessionID,
		ExpirationDate: expiresAt,
	}, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthService) VerifyAuth(ctx context.Context, token dto.SessionToken) (dto.UserID, error) {
	funcName := "AuthService.VerifyAuth"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sessionObj, err := a.authStorage.GetSession(ctx, token)
	if err != nil {
		return dto.UserID{}, err
	}
	logger.DebugFmt("Session found", funcName, nodeName)

	if sessionObj.ExpiryDate.Before(time.Now()) {
		logger.DebugFmt("Deleting expired session", funcName, nodeName)
		for err = a.LogOut(ctx, token); err != nil; {
			err = a.LogOut(ctx, token)
		}
		return dto.UserID{}, apperrors.ErrSessionExpired
	}
	logger.DebugFmt("Session is still good", funcName, nodeName)

	return dto.UserID{Value: sessionObj.UserID}, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthService) LogOut(ctx context.Context, token dto.SessionToken) error {
	return a.authStorage.DeleteSession(ctx, token)
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthService) GetLifetime(ctx context.Context) time.Duration {
	return a.sessionDuration
}

// generateString
// возвращает alphanumeric строку, собранную криптографически безопасным PRNG
func generateString(n uint) (string, error) {
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
