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

	logging "server/internal/logging"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	authStorage     storage.IAuthStorage
	sessionDuration time.Duration
	sessionIDLength uint
	logger          *logging.LogrusLogger
	UnimplementedAuthServiceServer
}

var AuthServiceErrorCodes = map[error]ErrorCode{
	nil:                             ErrorCode_OK,
	apperrors.ErrTokenNotGenerated:  ErrorCode_TOKEN_NOT_GENERATED,
	apperrors.ErrCouldNotBuildQuery: ErrorCode_COULD_NOT_BUILD_QUERY,
	apperrors.ErrSessionExpired:     ErrorCode_SESSION_EXPIRED,
	apperrors.ErrSessionNotCreated:  ErrorCode_SESSION_NOT_CREATED,
	apperrors.ErrSessionNotFound:    ErrorCode_SESSION_NOT_FOUND,
}

func NewAuthService(config config.SessionConfig, authStorage storage.IAuthStorage, logger *logging.LogrusLogger) *AuthService {
	return &AuthService{
		sessionDuration: config.Duration,
		sessionIDLength: config.IDLength,
		logger:          logger,
		authStorage:     authStorage,
	}
}

const nodeName = "microservice_server"

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (a *AuthService) AuthUser(ctx context.Context, id *UserID) (*AuthUserResponse, error) {
	funcName := "AuthService.AuthUser"
	expiresAt := time.Now().Add(a.sessionDuration)
	response := &AuthUserResponse{}

	sessionID, err := generateString(a.sessionIDLength)
	if err != nil {
		response.Code = AuthServiceErrorCodes[apperrors.ErrTokenNotGenerated]
		response.Response = &SessionToken{}
		return response, nil
	}
	a.logger.DebugFmt("Session ID generated", funcName, nodeName)

	session := &entities.Session{
		SessionID:  sessionID,
		UserID:     id.Value,
		ExpiryDate: expiresAt,
	}

	sCtx := context.WithValue(ctx, dto.LoggerKey, a.logger)

	err = a.authStorage.CreateSession(sCtx, session)
	if err != nil {
		response.Code = AuthServiceErrorCodes[err]
		response.Response = &SessionToken{}
		return response, nil
	}
	a.logger.DebugFmt("Session created", funcName, nodeName)

	response.Code = AuthServiceErrorCodes[nil]
	response.Response = &SessionToken{
		ID:             sessionID,
		ExpirationDate: timestamppb.New(expiresAt),
	}
	return response, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrSessionNotFound (401)
func (a *AuthService) VerifyAuth(ctx context.Context, token *SessionToken) (*VerifyAuthResponse, error) {
	funcName := "AuthService.VerifyAuth"
	convertedSession := dto.SessionToken{
		ID:             token.ID,
		ExpirationDate: token.ExpirationDate.AsTime(),
	}
	response := &VerifyAuthResponse{}

	sCtx := context.WithValue(ctx, dto.LoggerKey, a.logger)

	sessionObj, err := a.authStorage.GetSession(sCtx, convertedSession)
	if err != nil {
		a.logger.DebugFmt("Session not found", funcName, nodeName)
		response.Code = AuthServiceErrorCodes[err]
		response.Response = &UserID{}
		return response, nil
	}
	a.logger.DebugFmt("Found session", funcName, nodeName)

	if sessionObj.ExpiryDate.Before(time.Now()) {
		a.logger.DebugFmt("Deleting expired session", funcName, nodeName)
		for _, err = a.LogOut(sCtx, token); err != nil; {
			_, err = a.LogOut(sCtx, token)
		}
		response.Code = AuthServiceErrorCodes[apperrors.ErrSessionExpired]
		response.Response = &UserID{}
		return response, nil
	}

	response.Code = AuthServiceErrorCodes[nil]
	response.Response = &UserID{Value: sessionObj.UserID}
	return response, nil
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthService) LogOut(ctx context.Context, token *SessionToken) (*LogOutResponse, error) {
	funcName := "AuthService.LogOut"
	convertedSession := dto.SessionToken{
		ID:             token.ID,
		ExpirationDate: token.ExpirationDate.AsTime(),
	}
	response := &LogOutResponse{}

	sCtx := context.WithValue(ctx, dto.LoggerKey, a.logger)

	err := a.authStorage.DeleteSession(sCtx, convertedSession)
	if err != nil {
		a.logger.DebugFmt("Failed to delete session with error "+err.Error(), funcName, nodeName)
	} else {
		a.logger.DebugFmt("Session deleted", funcName, nodeName)
	}
	response.Code = AuthServiceErrorCodes[err]

	return response, nil
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthService) GetLifetime(ctx context.Context, empty *emptypb.Empty) (*GetLifetimeResponse, error) {
	response := &GetLifetimeResponse{
		Code:     AuthServiceErrorCodes[nil],
		Response: durationpb.New(a.sessionDuration),
	}
	return response, nil
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
