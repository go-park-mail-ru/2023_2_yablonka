package postgresql

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

// LocalUserStorage
// Локальное хранилище данных
type PostgresUserStorage struct {
	db *pgxpool.Pool
}

func NewPostgresUserStorage(db *pgxpool.Pool) *PostgresUserStorage {
	return &PostgresUserStorage{
		db: db,
	}
}

func (s *PostgresUserStorage) GetHighestID() uint64 {
	return 0
}

// GetUserByLogin
// находит пользователя в БД по почте
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *PostgresUserStorage) GetUserByLogin(ctx context.Context, login string) (*entities.User, error) {
	sql, args, err := sq.
		Select(allUserFields...).
		From("user").
		Where(sq.Eq{"login": login}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	user := entities.User{}
	if row.Scan(&user) != nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// GetUserByID
// находит пользователя в БД по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *PostgresUserStorage) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	sql, args, err := sq.
		Select(allUserFields...).
		From("user").
		Where(sq.Eq{"id": uid}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	user := entities.User{}
	if row.Scan(&user) != nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// CreateUser
// создает нового пользователя в БД по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
func (s *PostgresUserStorage) CreateUser(ctx context.Context, signup dto.SignupInfo) (*entities.User, error) {
	sql, args, err := sq.
		Insert("public.user").
		Columns("email", "password_hash").
		Values(signup.Email, signup.PasswordHash).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, email").
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	query := s.db.QueryRow(ctx, sql, args...)

	user := entities.User{}
	if query.Scan(&user.ID, &user.Email) != nil {
		return nil, apperrors.ErrUserNotCreated
	}

	return &user, nil
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
