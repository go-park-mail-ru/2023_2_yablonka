package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserService interface {
	// RegisterUser
	// создает нового пользователя по данным
	// или возвращает ошибки ...
	RegisterUser(context.Context, dto.AuthInfo) (*entities.User, error)
	// CheckPassword
	// проверяет пароль пользователя по почте
	// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
	CheckPassword(context.Context, dto.AuthInfo) (*entities.User, error)
	// ReadWithID
	// находит пользователя по его id
	// или возвращает ошибки ...
	GetWithID(context.Context, dto.UserID) (*entities.User, error)
	// UpdatePassword
	// обновляет пароль пользователя
	// или возвращает ошибки ...
	UpdatePassword(context.Context, dto.PasswordChangeInfo) error
	// UpdateProfile
	// обновляет профиль пользователя
	// или возвращает ошибки ...
	UpdateProfile(context.Context, dto.UserProfileInfo) error
	// UpdateAvatar
	// обновляет аватарку пользователя
	// или возвращает ошибки ...
	UpdateAvatar(context.Context, dto.AvatarChangeInfo) (*dto.UrlObj, error)
	// Delete
	// удаляет данного пользователя по id
	// или возвращает ошибки ...
	DeleteUser(context.Context, dto.UserID) error
}
