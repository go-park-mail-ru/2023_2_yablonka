package service

import (
	"context"
	"server/internal/apperrors"
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
// возвращает токен JWT, сгенерированный с помощью секрета в []byte и дату+время его истечения для полученного пользователя
func (a *AuthJWTService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	expiresAt := time.Now().Add(a.tokenLifetime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"iss":    "Tabula",
		"exp":    expiresAt,
	})

	str, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", time.Time{}, err
	}
	return str, expiresAt, nil
}

// VerifyAuth
// валидирует токен, возвращает ID пользователя, которому принадлежит токен
func (a *AuthJWTService) VerifyAuth(ctx context.Context, incomingToken string) (uint64, error) {
	token, err := jwt.Parse(incomingToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrJWTWrongMethod
		}
		return a.jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["userID"].(float64)
		if !ok {
			return 0, apperrors.ErrJWTMissingClaim
		}
		return uint64(userId), nil
	}
	return 0, apperrors.ErrJWTInvalidToken
}

// GetLifetime
// возвращает длительность жизни токена
func (a *AuthJWTService) GetLifetime() time.Duration {
	return a.tokenLifetime
}
