package microservice

import (
	"context"
	"server/internal/apperrors"
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

var AuthServiceErrors = map[microservice.ErrorCode]error{
	microservice.ErrorCode_OK:                    nil,
	microservice.ErrorCode_TOKEN_NOT_GENERATED:   apperrors.ErrTokenNotGenerated,
	microservice.ErrorCode_COULD_NOT_BUILD_QUERY: apperrors.ErrCouldNotBuildQuery,
	microservice.ErrorCode_SESSION_EXPIRED:       apperrors.ErrSessionExpired,
	microservice.ErrorCode_SESSION_NOT_CREATED:   apperrors.ErrSessionNotCreated,
	microservice.ErrorCode_SESSION_NOT_FOUND:     apperrors.ErrSessionNotFound,
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

	logger.DebugFmt("Contacting GRPC server", funcName, nodeName)
	serverResponse, _ := a.client.AuthUser(ctx, &microservice.UserID{Value: id.Value})
	logger.DebugFmt("Response received", funcName, nodeName)

	if serverResponse.Code != microservice.ErrorCode_OK {
		return dto.SessionToken{}, AuthServiceErrors[serverResponse.Code]
	}

	return dto.SessionToken{
		ID:             serverResponse.Response.ID,
		ExpirationDate: serverResponse.Response.ExpirationDate.AsTime(),
	}, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthService) VerifyAuth(ctx context.Context, token dto.SessionToken) (dto.UserID, error) {
	funcName := "AuthService.VerifyAuth"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.DebugFmt("Contacting GRPC server", funcName, nodeName)
	serverResponse, _ := a.client.VerifyAuth(ctx, &microservice.SessionToken{
		ID:             token.ID,
		ExpirationDate: timestamppb.New(token.ExpirationDate),
	})
	logger.DebugFmt("Response received", funcName, nodeName)

	if serverResponse.Code != microservice.ErrorCode_OK {
		return dto.UserID{}, AuthServiceErrors[serverResponse.Code]
	}

	return dto.UserID{Value: serverResponse.Response.Value}, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthService) LogOut(ctx context.Context, token dto.SessionToken) error {
	funcName := "AuthService.LogOut"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.DebugFmt("Contacting GRPC server", funcName, nodeName)
	serverResponse, _ := a.client.LogOut(ctx, &microservice.SessionToken{
		ID:             token.ID,
		ExpirationDate: timestamppb.New(token.ExpirationDate),
	})
	logger.DebugFmt("Response received", funcName, nodeName)

	return AuthServiceErrors[serverResponse.Code]
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthService) GetLifetime(ctx context.Context) time.Duration {
	funcName := "AuthService.GetLifetime"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.DebugFmt("Contacting GRPC server", funcName, nodeName)
	serverResponse, _ := a.client.GetLifetime(ctx, &emptypb.Empty{})
	logger.DebugFmt("Response received", funcName, nodeName)

	if serverResponse.Code != microservice.ErrorCode_OK {
		return 0
	}

	return serverResponse.Response.AsDuration()
}
