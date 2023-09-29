package service

import (
	"context"
	"server/internal/pkg/datatypes"
	authstorage "server/internal/storage/auth"
)

// Remake as interface
type AuthService struct {
	storage authstorage.IAuthStorage
}

func NewAuthService(storage authstorage.IAuthStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (a AuthService) SignUp(ctx context.Context, auth datatypes.LoginInfo) (*datatypes.User, error) {
	// TODO implement
	return nil, nil
}

func (a AuthService) LogIn(ctx context.Context, auth datatypes.LoginInfo) (*datatypes.User, error) {
	// TODO implementa
	return nil, nil
}
