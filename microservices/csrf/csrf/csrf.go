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

	"github.com/google/uuid"
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

var CSRFServiceErrorCodes = map[error]ErrorCode{
	nil:                             ErrorCode_OK,
	apperrors.ErrTokenNotGenerated:  ErrorCode_TOKEN_NOT_GENERATED,
	apperrors.ErrCouldNotBuildQuery: ErrorCode_COULD_NOT_BUILD_QUERY,
	apperrors.ErrCSRFExpired:        ErrorCode_CSRF_EXPIRED,
	apperrors.ErrCSRFNotCreated:     ErrorCode_CSRF_NOT_CREATED,
	apperrors.ErrCSRFNotFound:       ErrorCode_CSRF_NOT_FOUND,
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
func (cs *CSRFService) GetLifetime(ctx context.Context, empty *emptypb.Empty) (*GetLifetimeResponse, error) {
	response := &GetLifetimeResponse{
		Code:     CSRFServiceErrorCodes[nil],
		Response: durationpb.New(cs.sessionDuration),
	}
	return response, nil
}

// SetupCSRF
// возвращает уникальную строку CSRF и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) SetupCSRF(ctx context.Context, request *SetupCSRFRequest) (*SetupCSRFResponse, error) {
	funcName := "CSRFService.SetupCSRF"
	expiresAt := time.Now().Add(cs.sessionDuration)
	response := &SetupCSRFResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	id := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	token, err := generateToken(cs.sessionIDLength)
	if err != nil {
		response.Code = CSRFServiceErrorCodes[apperrors.ErrTokenNotGenerated]
		response.Response = &CSRFData{}
		return response, nil
	}
	cs.logger.DebugFmt("CSRF token generated", requestID.String(), funcName, nodeName)

	csrf := &entities.CSRF{
		Token:          token,
		UserID:         id.Value,
		ExpirationDate: expiresAt,
	}

	err = cs.csrfStorage.Create(sCtx, csrf)
	if err != nil {
		response.Code = CSRFServiceErrorCodes[err]
		response.Response = &CSRFData{}
		return response, nil
	}
	cs.logger.DebugFmt("CSRF session created", requestID.String(), funcName, nodeName)

	response.Code = CSRFServiceErrorCodes[nil]
	response.Response = &CSRFData{
		ID:             token,
		ExpirationDate: timestamppb.New(expiresAt),
	}

	return response, nil
}

// VerifyCSRF
// проверяет состояние CSRF, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500)
func (cs *CSRFService) VerifyCSRF(ctx context.Context, request *VerifyCSRFRequest) (*VerifyCSRFResponse, error) {
	funcName := "CSRFService.VerifyCSRF"
	response := &VerifyCSRFResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	token := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	CSRFObj, err := cs.csrfStorage.Get(sCtx, dto.CSRFToken{Value: token.Value})
	if err != nil {
		response.Code = CSRFServiceErrorCodes[err]
		return response, nil
	}
	cs.logger.DebugFmt("CSRF token found", requestID.String(), funcName, nodeName)

	if CSRFObj.ExpirationDate.Before(time.Now()) {
		cs.logger.DebugFmt("Deleting expired token", requestID.String(), funcName, nodeName)
		deleteCSRFRequest := &DeleteCSRFRequest{
			RequestID: request.RequestID,
			Value:     request.Value,
		}
		for _, err = cs.DeleteCSRF(ctx, deleteCSRFRequest); err != nil; {
			_, err = cs.DeleteCSRF(ctx, deleteCSRFRequest)
		}
		response.Code = CSRFServiceErrorCodes[apperrors.ErrCSRFExpired]
		return response, nil
	}
	cs.logger.DebugFmt("CSRF token is still good", requestID.String(), funcName, nodeName)

	response.Code = CSRFServiceErrorCodes[nil]

	return response, nil
}

// DeleteCSRF
// удаляет CSRF
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (cs *CSRFService) DeleteCSRF(ctx context.Context, request *DeleteCSRFRequest) (*DeleteCSRFResponse, error) {
	response := &DeleteCSRFResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	token := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, cs.logger),
		dto.RequestIDKey, requestID,
	)

	err := cs.csrfStorage.Delete(sCtx, dto.CSRFToken{Value: token.Value})
	response.Code = CSRFServiceErrorCodes[err]

	return response, nil
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
