package service

import (
	"context"
	"os"
	"server/internal/pkg/datatypes"

	"github.com/dgrijalva/jwt-go"
)

// Приватная функция: не нужна за пределами service/auth
// Генерирует токен из datatypes.User с помощью секрета в []byte
func generateJWT(user *datatypes.User, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"ID":    user.ID,
	})

	str, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return str, nil
}

// Интерфейс для аутентификации (создание и проверка токена)
type IAuthService interface {
	AuthUser(context.Context, *datatypes.User) (string, error)
	VerifyAuth(context.Context, string) (string, error)
}

// Структура сервиса аутентификации с помощью JWT
type AuthJWTService struct {
	jwtSecret []byte
}

// Возвращает AuthJWTService с рабочим JWT-секретом
func NewAuthJWTService() *AuthJWTService {
	secret := os.Getenv("JWT_SECRET")
	return &AuthJWTService{
		jwtSecret: []byte(secret),
	}
}

// Возвращает JWT для полученного пользователя
func (a *AuthJWTService) AuthUser(ctx context.Context, user *datatypes.User) (string, error) {
	token, err := generateJWT(user, a.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthJWTService) VerifyAuth(ctx context.Context, token string) (string, error) {
	// TODO implement
	return "", nil
}
