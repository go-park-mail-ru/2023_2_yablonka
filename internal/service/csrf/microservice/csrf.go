package service

import (
	"context"
	config "server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/csrf/csrf"
	"time"

	logger "server/internal/logging"

	"google.golang.org/grpc"
)

type CSRFService struct {
	sessionDuration time.Duration
	sessionIDLength uint
	storage         storage.ICSRFStorage
	client          microservice.CSRFServiceClient
}

const nodeName = "service"

// NewAuthService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewCSRFService(config config.SessionConfig, CSRFStorage storage.ICSRFStorage, conn *grpc.ClientConn) *CSRFService {
	client := microservice.NewCSRFServiceClient(conn)
	return &CSRFService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		storage:         CSRFStorage,
		client:          client,
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
	funcName := "CSRFService.SetupCSRF"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	csrfpb, err := cs.client.SetupCSRF(ctx, &microservice.UserID{Value: id.Value})
	if err != nil {
		return dto.CSRFData{}, err
	}
	logger.Debug("Info received", funcName, nodeName)

	return dto.CSRFData{
		Token:          csrfpb.ID,
		ExpirationDate: csrfpb.ExpirationDate.AsTime(),
	}, nil
}

// VerifyCSRF
// проверяет состояние CSRF, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) VerifyCSRF(ctx context.Context, token dto.CSRFToken) error {
	funcName := "CSRFService.VerifyCSRF"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	_, err := cs.client.VerifyCSRF(ctx, &microservice.CSRFToken{
		Value: token.Value,
	})
	logger.Debug("Info received", funcName, nodeName)

	return err
}

// DeleteCSRF
// удаляет CSRF
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (cs *CSRFService) DeleteCSRF(ctx context.Context, token dto.CSRFToken) error {
	funcName := "CSRFService.DeleteCSRF"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	_, err := cs.client.DeleteCSRF(ctx, &microservice.CSRFToken{
		Value: token.Value,
	})
	logger.Debug("Info received", funcName, nodeName)

	return err
}
