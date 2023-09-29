package service

import (
	"context"
	"server/internal/pkg/datatypes"
	authstorage "server/internal/storage/auth"
)

type AuthService struct {
	storage *authstorage.AuthStorage
}

func NewAuthService(storage *authstorage.AuthStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (a AuthService) SignUp(ctx context.Context, auth datatypes.LoginInfo) (*datatypes.User, error) {
	// TODO implement
	return nil, nil
}

func (a AuthService) LogIn(ctx context.Context, auth datatypes.LoginInfo) (*datatypes.User, error) {
	// TODO implement
	return nil, nil
}
