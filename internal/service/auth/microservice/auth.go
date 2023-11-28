package microservice

import (
	"context"
	"crypto/rand"
	"math/big"
	config "server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	microservice "server/microservices/auth/auth"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	sessionDuration time.Duration
	client          microservice.AuthServiceClient
	sessionIDLength uint
	authStorage     storage.IAuthStorage
}

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
	sessionpb, err := a.client.AuthUser(ctx, &microservice.UserID{Value: id.Value})
	if err != nil {
		return dto.SessionToken{}, nil
	}

	return dto.SessionToken{
		ID:             sessionpb.ID,
		ExpirationDate: sessionpb.ExpirationDate.AsTime(),
	}, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthService) VerifyAuth(ctx context.Context, token dto.SessionToken) (dto.UserID, error) {
	uidpb, err := a.client.VerifyAuth(ctx, &microservice.SessionToken{
		ID:             token.ID,
		ExpirationDate: timestamppb.New(token.ExpirationDate),
	})
	if err != nil {
		return dto.UserID{}, err
	}

	return dto.UserID{Value: uidpb.Value}, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthService) LogOut(ctx context.Context, token dto.SessionToken) error {
	_, err := a.client.LogOut(ctx, &microservice.SessionToken{
		ID:             token.ID,
		ExpirationDate: timestamppb.New(token.ExpirationDate),
	})

	return err
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthService) GetLifetime(ctx context.Context) time.Duration {
	lifetimepb, err := a.client.GetLifetime(ctx, nil)
	if err != nil {
		return 0
	}
	return lifetimepb.AsDuration()
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
