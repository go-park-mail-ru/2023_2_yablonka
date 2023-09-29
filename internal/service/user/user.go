package service

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/datatypes"
	"server/internal/storage"
)

type IAuthUserService interface {
	GetUser(ctx context.Context, info datatypes.LoginInfo) (*datatypes.User, error)
	CreateUser(ctx context.Context, info datatypes.SignupInfo) (*datatypes.User, error)
}

type IUserService interface {
	UpdateUser()
	GetUserBoards()
	IAuthUserService
}

type AuthUserService struct {
	storage storage.IUserStorage
}

type UserService struct {
	storage storage.IUserStorage
}

func NewAuthUserService(storage storage.IUserStorage) AuthUserService {
	return AuthUserService{
		storage: storage,
	}
}

func NewUserService(storage storage.IUserStorage) UserService {
	return UserService{
		storage: storage,
	}
}

func (us AuthUserService) GetUser(ctx context.Context, info datatypes.LoginInfo) (*datatypes.User, error) {
	user, err := us.storage.GetUser(info)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash != info.PasswordHash {
		return nil, apperrors.ErrWrongPassword
	}

	return user, nil
}

func (us AuthUserService) CreateUser(ctx context.Context, info datatypes.SignupInfo) (*datatypes.User, error) {
	user, err := us.storage.CreateUser(info)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us UserService) UpdateUser(ctx context.Context, user *datatypes.User) {
	// TODO Изменение значений у юзера
}

func (us UserService) GetUserBoards(ctx context.Context, user *datatypes.User) []*datatypes.Board {
	// TODO Запрос всех досок пользователя
	return nil
}
