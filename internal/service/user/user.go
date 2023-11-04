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

// Login
// находит пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us AuthUserService) Login(ctx context.Context, info dto.LoginInfo) (*entities.User, error) {
	user, err := us.storage.GetUserByLogin(ctx, info.Email)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash != info.PasswordHash {
		return nil, apperrors.ErrWrongPassword
	}

	return user, nil
}

// GetUserByID
// находит пользователя по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us AuthUserService) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	return us.storage.GetUserByID(ctx, uid)
}

// RegisterUser
// создает нового пользователя по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
func (us AuthUserService) RegisterUser(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	_, err := us.storage.GetUserByLogin(ctx, info.Email)

	if err == nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

	return us.storage.Create(ctx, info)
}

// UpdateUser
// обновляет пользователя в БД
// или возвращает ошибку apperrors.ErrUserNotFound (409)
<<<<<<< Updated upstream
func (us UserService) UpdateUser(ctx context.Context, info dto.UpdatedUserInfo) (*entities.User, error) {
	return us.storage.UpdateUser(ctx, info)
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
func (us UserService) UpdateUser(ctx context.Context, info dto.UpdatedUserInfo) (*entities.User, error) {
	return us.storage.UpdateUser(ctx, info)
=======
>>>>>>> Stashed changes
func (us UserService) UpdateUser(ctx context.Context, updatedUser entities.User) (*entities.User, error) {
	_, err := us.storage.GetUserByLogin(ctx, updatedUser.Email)

	if err == nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

<<<<<<< Updated upstream
	return us.storage.Update(ctx, updatedUser)
=======
	return &updatedUser, us.storage.Update(ctx, updatedUser)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us AuthUserService) DeleteUser(ctx context.Context, uid uint64) error {
	return us.storage.Delete(ctx, uid)
}
