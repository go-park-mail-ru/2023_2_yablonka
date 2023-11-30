package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища пользователей
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type IUserStorage interface {
	// GetWithLogin
	// находит пользователя в БД по почте
	// или возвращает ошибки ...
	GetWithLogin(context.Context, dto.UserLogin) (*entities.User, error)
	// GetWithID
	// находит пользователя в БД по его id
	// или возвращает ошибки ...
	GetWithID(context.Context, dto.UserID) (*entities.User, error)
	// GetLoginInfoWithID
	// находит данные логина пользователя в БД по id
	// или возвращает ошибки ...
	GetLoginInfoWithID(context.Context, dto.UserID) (*dto.LoginInfo, error)
	// Create
	// создает нового пользователя в БД по данным
	// или возвращает ошибки ...
	Create(context.Context, dto.SignupInfo) (*entities.User, error)
	// UpdatePassword
	// обновляет пароль пользователя в БД
	// или возвращает ошибки ...
	UpdatePassword(context.Context, dto.PasswordHashesInfo) error
	// UpdateProfile
	// обновляет профиль пользователя в БД
	// или возвращает ошибки ...
	UpdateProfile(context.Context, dto.UserProfileInfo) error
	// UpdateAvatar
	// обновляет аватарку пользователя в БД
	// или возвращает ошибки ...
	UpdateAvatarUrl(context.Context, dto.ImageUrlInfo) error
	// Delete
	// удаляет данного пользователя в БД по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.UserID) error
}
