package embedded_services

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

type CSRFService struct {
	sessionDuration time.Duration
	sessionIDLength uint
	storage         storage.ICSRFStorage
}

const nodeName = "service"

// NewAuthService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewCSRFService(config config.SessionConfig, CSRFStorage storage.ICSRFStorage) *CSRFService {
	return &CSRFService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		storage:         CSRFStorage,
	}
}

// GetLifetime
// возвращает длительность авторизации
func (cs *CSRFService) GetLifetime(ctx context.Context) time.Duration {
	return cs.sessionDuration
}

// SetupCSRF
// возвращает уникальную строку CSRF и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) SetupCSRF(ctx context.Context, id dto.UserID) (dto.CSRFData, error) {
	expiresAt := time.Now().Add(cs.sessionDuration)

	token, err := generateToken(cs.sessionIDLength)
	if err != nil {
		return dto.CSRFData{}, apperrors.ErrTokenNotGenerated
	}

	csrf := &entities.CSRF{
		Token:          token,
		UserID:         id.Value,
		ExpirationDate: expiresAt,
	}

	err = cs.storage.Create(ctx, csrf)
	if err != nil {
		return dto.CSRFData{}, err
	}

	return dto.CSRFData{
		Token:          token,
		ExpirationDate: expiresAt,
	}, nil
}

// VerifyCSRF
// проверяет состояние CSRF, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) VerifyCSRF(ctx context.Context, token dto.CSRFToken) error {
	funcName := "CSRFService.VerifyCSRF"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	CSRFObj, err := cs.storage.Get(ctx, token)
	if err != nil {
		return err
	}

	if CSRFObj.ExpirationDate.Before(time.Now()) {
		logger.Debug("Deleting expired CSRF token", funcName, nodeName)
		for err := cs.DeleteCSRF(ctx, token); err != nil; {
			err = cs.DeleteCSRF(ctx, token)
		}
		return apperrors.ErrSessionExpired
	}

	return nil
}

// DeleteCSRF
// удаляет CSRF
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (cs *CSRFService) DeleteCSRF(ctx context.Context, token dto.CSRFToken) error {
	return cs.storage.Delete(ctx, token)
}

// generateString
// возвращает alphanumeric строку, собранную криптографически безопасным PRNG
func generateToken(n uint) (string, error) {
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
