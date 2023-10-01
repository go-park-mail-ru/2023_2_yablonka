package service

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// generateJWT
// генерирует токен из entities.User с помощью секрета в []byte, который будет действителен tokenLifetime
func generateJWT(user *entities.User, secret []byte, tokenLifetime time.Duration) (string, time.Time, error) {
	expiresAt := time.Now().Add(tokenLifetime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"iss":    "Tabula",
		"exp":    expiresAt,
	})

	str, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return str, expiresAt, nil
}

// AuthJWTService
// структура сервиса аутентификации с помощью JWT
// содержит секрет для подписи и время действия токена
type AuthJWTService struct {
	jwtSecret     []byte
	tokenLifetime time.Duration
}

// AuthUser
// возвращает JWT и дату+время его истечения для полученного пользователя
func (a *AuthJWTService) AuthUser(ctx context.Context, user *entities.User) (string, time.Time, error) {
	token, expiresAt, err := generateJWT(user, a.jwtSecret, a.tokenLifetime)
	if err != nil {
		return "", time.Time{}, err
	}
	return token, expiresAt, nil
}

// VerifyAuth
// валидирует токен, возвращает ID пользователя, которому принадлежит токен
func (a *AuthJWTService) VerifyAuth(ctx context.Context, incomingToken string) (uint64, error) {
	var userID uint64
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
		uidFloat, ok := claims["userID"].(float64)
		if !ok {
			return 0, apperrors.ErrJWTMissingClaim
		}
		userID = uint64(uidFloat)
	} else {
		return 0, apperrors.ErrJWTInvalidToken
	}
	return userID, nil
}
