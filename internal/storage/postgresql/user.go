package in_memory

import (
	"context"
	"database/sql"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// LocalUserStorage
// Локальное хранилище данных
type PostgresUserStorage struct {
	db sql.DB
}

func NewUserStorage() *PostgresUserStorage {
	return &PostgresUserStorage{}
}

func (s *PostgresUserStorage) GetHighestID() uint64 {
	return 0
}

// GetUser
// находит пользователя в БД по почте
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *PostgresUserStorage) GetUser(ctx context.Context, login dto.LoginInfo) (*entities.User, error) {
	user := entities.User{}
	err := s.db.QueryRow("SELECT id, email, password_hash, name, surname, avatar_url, description FROM user WEHRE id = :id",
		map[string]interface{}{
			"id": login.Email,
		}).
		Scan(&user.ID, &user)
	return &user, err
}

// GetUserByID
// находит пользователя в БД по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *PostgresUserStorage) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	return nil, apperrors.ErrUserNotFound
}

// CreateUser
// создает нового пользователя в БД по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
func (s *PostgresUserStorage) CreateUser(ctx context.Context, signup dto.SignupInfo) (*entities.User, error) {
	return nil, apperrors.ErrUserAlreadyExists
}

// UpdateUser
// обновляет пользователя в БД
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (s *PostgresUserStorage) UpdateUser(ctx context.Context, updatedInfo dto.UpdatedUserInfo) (*entities.User, error) {
	return nil, apperrors.ErrUserNotFound
}

// DeleteUser
// удаляет данного пользователя в БД по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (s *PostgresUserStorage) DeleteUser(ctx context.Context, id uint64) error {
	return apperrors.ErrUserNotFound
}
