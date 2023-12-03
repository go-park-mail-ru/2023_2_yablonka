package csrf_microservice

import (
	context "context"
	"crypto/rand"
	"math/big"
	"server/internal/apperrors"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"time"

	logger "server/internal/logging"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type CSRFService struct {
	csrfStorage     storage.ICSRFStorage
	sessionDuration time.Duration
	sessionIDLength uint
	logger          *logger.LogrusLogger
	UnimplementedCSRFServiceServer
}

const nodeName = "microservice_server"

// NewAuthService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewCSRFService(config config.SessionConfig, CSRFStorage storage.ICSRFStorage, logger *logger.LogrusLogger) *CSRFService {
	return &CSRFService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		logger:          logger,
		csrfStorage:     CSRFStorage,
	}
}

// GetLifetime
// возвращает длительность авторизации
func (cs *CSRFService) GetLifetime(ctx context.Context, empty *emptypb.Empty) (*durationpb.Duration, error) {
	return durationpb.New(cs.sessionDuration), nil
}

// SetupCSRF
// возвращает уникальную строку CSRF и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) SetupCSRF(ctx context.Context, id *UserID) (*CSRFData, error) {
	funcName := "CSRFService.SetupCSRF"
	expiresAt := time.Now().Add(cs.sessionDuration)

	token, err := generateToken(cs.sessionIDLength)
	if err != nil {
		return &CSRFData{}, apperrors.MakeGRPCError(apperrors.ErrTokenNotGenerated)
	}
	cs.logger.DebugFmt("CSRF token generated", funcName, nodeName)

	csrf := &entities.CSRF{
		Token:          token,
		UserID:         id.Value,
		ExpirationDate: expiresAt,
	}

	err = cs.csrfStorage.Create(ctx, csrf)
	if err != nil {
		return &CSRFData{}, apperrors.MakeGRPCError(err)
	}
	cs.logger.DebugFmt("CSRF session created", funcName, nodeName)

	return &CSRFData{
		ID:             token,
		ExpirationDate: timestamppb.New(expiresAt),
	}, nil
}

// VerifyCSRF
// проверяет состояние CSRF, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) VerifyCSRF(ctx context.Context, token *CSRFToken) (*emptypb.Empty, error) {
	funcName := "CSRFService.VerifyCSRF"
	CSRFObj, err := cs.csrfStorage.Get(ctx, dto.CSRFToken{Value: token.Value})
	if err != nil {
		return &emptypb.Empty{}, apperrors.MakeGRPCError(err)
	}
	cs.logger.DebugFmt("CSRF token found", funcName, nodeName)

	if CSRFObj.ExpirationDate.Before(time.Now()) {
		cs.logger.DebugFmt("Deleting expired token", funcName, nodeName)
		for _, err = cs.DeleteCSRF(ctx, token); err != nil; {
			_, err = cs.DeleteCSRF(ctx, token)
		}
		return &emptypb.Empty{}, apperrors.MakeGRPCError(apperrors.ErrSessionExpired)
	}
	cs.logger.DebugFmt("CSRF token is still good", funcName, nodeName)

	return &emptypb.Empty{}, nil
}

// DeleteCSRF
// удаляет CSRF
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (cs *CSRFService) DeleteCSRF(ctx context.Context, token *CSRFToken) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, apperrors.MakeGRPCError(cs.csrfStorage.Delete(ctx, dto.CSRFToken{Value: token.Value}))
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
