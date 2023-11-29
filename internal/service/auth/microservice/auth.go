package microservice

import (
	"context"
	config "server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/auth/auth"
	"time"

	logger "server/internal/logging"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	sessionDuration time.Duration
	client          microservice.AuthServiceClient
	sessionIDLength uint
	authStorage     storage.IAuthStorage
}

const nodeName = "service"

// NewAuthService
// возвращает AuthSessionService с инициализированной датой истечения сессии и хранилищем сессий
func NewAuthService(config config.SessionConfig, authStorage storage.IAuthStorage, connection *grpc.ClientConn) *AuthService {
	client := microservice.NewAuthServiceClient(connection)
	return &AuthService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		authStorage:     authStorage,
		client:          client,
	}
}

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (a *AuthService) AuthUser(ctx context.Context, id dto.UserID) (dto.SessionToken, error) {
	funcName := "AuthService.AuthUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	sessionpb, err := a.client.AuthUser(ctx, &microservice.UserID{Value: id.Value})
	if err != nil {
		return dto.SessionToken{}, err
	}
	logger.Debug("Info received", funcName, nodeName)

	return dto.SessionToken{
		ID:             sessionpb.ID,
		ExpirationDate: sessionpb.ExpirationDate.AsTime(),
	}, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthService) VerifyAuth(ctx context.Context, token dto.SessionToken) (dto.UserID, error) {
	funcName := "AuthService.VerifyAuth"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	uidpb, err := a.client.VerifyAuth(ctx, &microservice.SessionToken{
		ID:             token.ID,
		ExpirationDate: timestamppb.New(token.ExpirationDate),
	})
	if err != nil {
		return dto.UserID{}, err
	}
	logger.Debug("Info received", funcName, nodeName)

	return dto.UserID{Value: uidpb.Value}, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthService) LogOut(ctx context.Context, token dto.SessionToken) error {
	funcName := "AuthService.LogOut"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	_, err := a.client.LogOut(ctx, &microservice.SessionToken{
		ID:             token.ID,
		ExpirationDate: timestamppb.New(token.ExpirationDate),
	})
	logger.Debug("Info received", funcName, nodeName)

	return err
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthService) GetLifetime(ctx context.Context) time.Duration {
	lifetimepb, err := a.client.GetLifetime(ctx, &emptypb.Empty{})
	if err != nil {
		return 0
	}
	return lifetimepb.AsDuration()
}
