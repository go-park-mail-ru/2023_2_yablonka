package service

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type AuthUserService struct {
	storage storage.IUserStorage
}

type UserService struct {
	storage storage.IUserStorage
}

func NewAuthUserService(storage storage.IUserStorage) *AuthUserService {
	return &AuthUserService{
		storage: storage,
	}
}

func NewUserService(storage storage.IUserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us AuthUserService) GetUser(ctx context.Context, info dto.LoginInfo) (*entities.User, error) {
	user, err := us.storage.GetUser(info)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash != info.PasswordHash {
		return nil, apperrors.ErrWrongPassword
	}

	return user, nil
}

func (us AuthUserService) CreateUser(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	user, err := us.storage.CreateUser(info)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us UserService) UpdateUser(ctx context.Context, info dto.UpdatedUserInfo) (*entities.User, error) {
	// TODO Временно можно обновлять и мэйл, надо сделать какую-то проверку для этого
	user, err := us.storage.UpdateUser(info)
	if err != nil {
		return nil, err
	}
	return user, nil
}
