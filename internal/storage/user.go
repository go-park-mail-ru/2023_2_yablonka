package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserStorage interface {
	// GetUserByLogin
	// находит пользователя в БД по почте
	// или возвращает ошибки ...
	GetUserByLogin(ctx context.Context, login string) (*entities.User, error)
	// GetUserByID
	// находит пользователя в БД по его id
	// или возвращает ошибки ...
	GetUserByID(ctx context.Context, id uint64) (*entities.User, error)
	// Create
	// создает нового пользователя в БД по данным
	// или возвращает ошибки ...
	Create(ctx context.Context, info dto.SignupInfo) (*entities.User, error)
	// Update
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
	// или возвращает ошибки ...
	Delete(ctx context.Context, id uint64) error
}
