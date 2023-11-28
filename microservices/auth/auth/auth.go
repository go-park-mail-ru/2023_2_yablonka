package auth_microservice

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

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	authStorage     storage.IAuthStorage
	sessionDuration time.Duration
	sessionIDLength uint
	logger          *logrus.Logger
	UnimplementedAuthServiceServer
}

func NewAuthService(config config.SessionConfig, authStorage storage.IAuthStorage, logger *logrus.Logger) *AuthService {
	return &AuthService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		logger:          logger,
		authStorage:     authStorage,
	}
}

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (a *AuthService) AuthUser(ctx context.Context, id *UserID) (*SessionToken, error) {
	funcName := "AuthUser"
	expiresAt := time.Now().Add(a.sessionDuration)

	sessionID, err := generateString(a.sessionIDLength)
	if err != nil {
		return &SessionToken{}, apperrors.ErrTokenNotGenerated
	}
	authServiceDebugLog(a.logger, funcName, "Session ID generated")

	session := &entities.Session{
		SessionID:  sessionID,
		UserID:     id.Value,
		ExpiryDate: expiresAt,
	}

	err = a.authStorage.CreateSession(ctx, session)
	if err != nil {
		return &SessionToken{}, err
	}
	authServiceDebugLog(a.logger, funcName, "Session created")

	return &SessionToken{
		ID:             sessionID,
		ExpirationDate: timestamppb.New(expiresAt),
	}, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthService) VerifyAuth(ctx context.Context, token *SessionToken) (*UserID, error) {
	convertedSession := dto.SessionToken{
		ID:             token.ID,
		ExpirationDate: token.ExpirationDate.AsTime(),
	}

	sessionObj, err := a.authStorage.GetSession(ctx, convertedSession)
	if err != nil {
		return &UserID{}, err
	}

	if sessionObj.ExpiryDate.Before(time.Now()) {
		return &UserID{}, apperrors.ErrSessionExpired
	}
	return &UserID{Value: sessionObj.UserID}, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthService) LogOut(ctx context.Context, token *SessionToken) (*emptypb.Empty, error) {
	convertedSession := dto.SessionToken{
		ID:             token.ID,
		ExpirationDate: token.ExpirationDate.AsTime(),
	}
	return &emptypb.Empty{}, a.authStorage.DeleteSession(ctx, convertedSession)
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthService) GetLifetime(ctx context.Context, empty *emptypb.Empty) (*durationpb.Duration, error) {
	return durationpb.New(a.sessionDuration), nil
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

func authServiceDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "grpc_service",
			"function":   function,
		}).
		Debug(message)
}
