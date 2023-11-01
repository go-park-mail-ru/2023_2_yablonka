package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserAuthService interface {
	// Login
	// находит пользователя по почте
	// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
	Login(context.Context, dto.LoginInfo) (*entities.User, error)
	// GetUserByID
	// находит пользователя по его id
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserByID(context.Context, uint64) (*entities.User, error)
	// RegisterUser
	// создает нового пользователя по данным
	// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
	RegisterUser(context.Context, dto.SignupInfo) (*entities.User, error)
	// DeleteUser
	// удаляет данного пользователя по id
	// или возвращает ошибку apperrors.ErrUserNotFound (409)
	DeleteUser(context.Context, uint64) error

	// TODO Implement
	// GetUserFromAuthInfo(context.Context, dto.VerifiedAuthInfo) (*entities.User, error)
}

type IUserService interface {
	IUserAuthService
	// UpdateUser
	// обновляет пользователя в БД
	// или возвращает ошибку apperrors.ErrUserNotFound (409)
	UpdateUser(context.Context, entities.User) (*entities.User, error)
}
