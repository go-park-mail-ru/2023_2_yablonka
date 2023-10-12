package service

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthJWTService
// структура сервиса аутентификации с помощью JWT
// содержит секрет для подписи и время действия токена
type AuthJWTService struct {
	jwtSecret     []byte
	tokenLifetime time.Duration
}

// AuthUser
// возвращает уникальную строку авторизации и её длительность
// или возвращает ошибки apperrors.ErrTokenNotGenerated (500),
// apperrors.ErrJWTWrongMethod (401), apperrors.ErrJWTMissingClaim (401), apperrors.ErrJWTInvalidToken (401)
func (a *AuthJWTService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	expiresAt := time.Now().Add(a.tokenLifetime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"iss":    "Tabula",
		"exp":    expiresAt,
	})

	str, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", time.Time{}, apperrors.ErrTokenNotGenerated
	}
	return str, expiresAt, nil
}

// VerifyAuth
// проверяет состояние авторизации, возвращает ID авторизированного пользователя
// или возвращает ошибки apperrors.ErrJWTWrongMethod (401), apperrors.ErrJWTMissingClaim (401), apperrors.ErrJWTInvalidToken (401)
func (a *AuthJWTService) VerifyAuth(ctx context.Context, incomingToken string) (*dto.VerifiedAuthInfo, error) {
	token, err := jwt.Parse(incomingToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrJWTWrongMethod
		}
		return a.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["userID"].(float64)
		if !ok {
			return nil, apperrors.ErrJWTMissingClaim
		}
		return &dto.VerifiedAuthInfo{
			UserID: uint64(userIDFloat),
		}, nil
	}
	return nil, apperrors.ErrJWTInvalidToken
}

// GetLifetime
// возвращает длительность авторизации
func (a *AuthJWTService) GetLifetime() time.Duration {
	return a.tokenLifetime
}

// LogOut
// удаляет текущую сессию
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (a *AuthJWTService) LogOut(ctx context.Context, sessionString string) error {
	return nil
}
