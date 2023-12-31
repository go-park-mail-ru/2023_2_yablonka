package service

import (
	"context"
	"server/internal/apperrors"
	config "server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/csrf/csrf"
	"time"

	logger "server/internal/logging"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type CSRFService struct {
	sessionDuration time.Duration
	sessionIDLength uint
	storage         storage.ICSRFStorage
	client          microservice.CSRFServiceClient
}

var CSRFServiceErrors = map[microservice.ErrorCode]error{
	microservice.ErrorCode_OK:                    nil,
	microservice.ErrorCode_TOKEN_NOT_GENERATED:   apperrors.ErrTokenNotGenerated,
	microservice.ErrorCode_COULD_NOT_BUILD_QUERY: apperrors.ErrCouldNotBuildQuery,
	microservice.ErrorCode_CSRF_EXPIRED:          apperrors.ErrCSRFExpired,
	microservice.ErrorCode_CSRF_NOT_CREATED:      apperrors.ErrCSRFNotCreated,
	microservice.ErrorCode_CSRF_NOT_FOUND:        apperrors.ErrCSRFNotFound,
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
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.SetupCSRFRequest{
		RequestID: requestID.String(),
		Value:     &microservice.UserID{Value: id.Value},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.SetupCSRF(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	if serverResponse.Code != microservice.ErrorCode_OK {
		return dto.CSRFData{}, CSRFServiceErrors[serverResponse.Code]
	}

	csrf := serverResponse.Response

	return dto.CSRFData{
		Token:          csrf.ID,
		ExpirationDate: csrf.ExpirationDate.AsTime(),
	}, nil
}

// VerifyCSRF
// проверяет состояние CSRF, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) VerifyCSRF(ctx context.Context, token dto.CSRFToken) error {
	funcName := "CSRFService.VerifyCSRF"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.VerifyCSRFRequest{
		RequestID: requestID.String(),
		Value:     &microservice.CSRFToken{Value: token.Value},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.VerifyCSRF(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return CSRFServiceErrors[serverResponse.Code]
}

// DeleteCSRF
// удаляет CSRF
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (cs *CSRFService) DeleteCSRF(ctx context.Context, token dto.CSRFToken) error {
	funcName := "CSRFService.DeleteCSRF"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.DeleteCSRFRequest{
		RequestID: requestID.String(),
		Value:     &microservice.CSRFToken{Value: token.Value},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := cs.client.DeleteCSRF(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return CSRFServiceErrors[serverResponse.Code]
}
