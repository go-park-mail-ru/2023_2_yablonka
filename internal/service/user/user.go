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

// NewAuthUserService
// возвращает AuthUserService с инициализированным хранилищем пользователей
// который имплементирует авторизационные методы, привязанные к UserStorage
func NewAuthUserService(storage storage.IUserStorage) *AuthUserService {
	return &AuthUserService{
		storage: storage,
	}
}

// NewUserService
// возвращает UserService с инициализированным хранилищем пользователей
func NewUserService(storage storage.IUserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us AuthUserService) GetUser(ctx context.Context, info dto.LoginInfo) (*entities.User, error) {
	user, err := us.storage.GetUser(ctx, info)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash != info.PasswordHash {
		return nil, apperrors.ErrWrongPassword
	}

	return user, nil
}

func (us AuthUserService) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	return us.storage.GetUserByID(ctx, uid)
}

func (us AuthUserService) CreateUser(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	return us.storage.CreateUser(ctx, info)
}

func (us UserService) UpdateUser(ctx context.Context, info dto.UpdatedUserInfo) (*entities.User, error) {
	return us.storage.UpdateUser(ctx, info)
}

func (us AuthUserService) DeleteUser(ctx context.Context, uid uint64) error {
	return us.storage.DeleteUser(ctx, uid)
}
