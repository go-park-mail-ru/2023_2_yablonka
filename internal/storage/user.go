package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserStorage interface {
	// GetUser
	// находит пользователя в БД по почте
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUser(context.Context, dto.LoginInfo) (*entities.User, error)
	// GetUserByID
	// находит пользователя в БД по его id
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserByID(context.Context, uint64) (*entities.User, error)
	// CreateUser
	// создает нового пользователя в БД по данным
	// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
	CreateUser(context.Context, dto.SignupInfo) (*entities.User, error)
	// UpdateUser
	// обновляет пользователя в БД
<<<<<<< Updated upstream
	// или возвращает ошибку apperrors.ErrUserNotFound (409)
	UpdateUser(context.Context, dto.UpdatedUserInfo) (*entities.User, error)
	// DeleteUser
=======
<<<<<<< Updated upstream
	// или возвращает ошибки ...
	Update(ctx context.Context, updatedUser entities.User) (*entities.User, error)
	// Delete
=======
<<<<<<< Updated upstream
	// или возвращает ошибку apperrors.ErrUserNotFound (409)
	UpdateUser(context.Context, dto.UpdatedUserInfo) (*entities.User, error)
	// DeleteUser
=======
	// или возвращает ошибки ...
	Update(ctx context.Context, updatedUser entities.User) error
	// Delete
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	// удаляет данного пользователя в БД по id
	// или возвращает ошибку apperrors.ErrUserNotFound (409)
	DeleteUser(context.Context, uint64) error
}
