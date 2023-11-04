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
// Хранилище данных в PostgreSQL
type PostgresUserStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *PostgresUserStorage {
	return &PostgresUserStorage{
		db: db,
	}
}

// GetUserByLogin
// находит пользователя в БД по почте
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetUserByLogin(ctx context.Context, login string) (*entities.User, error) {
	sql, args, err := sq.
		Select(allUserFields...).
<<<<<<< Updated upstream
		From("public.User").
=======
		From("public.user").
>>>>>>> Stashed changes
		Where(sq.Eq{"login": login}).
		ToSql()

	if err != nil {
<<<<<<< Updated upstream
		return nil, apperrors.ErrCouldntBuildQuery
=======
		return nil, apperrors.ErrCouldNotBuildQuery
>>>>>>> Stashed changes
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
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	sql, args, err := sq.
		Select(allUserFields...).
<<<<<<< Updated upstream
		From("public.User").
=======
		From("public.user").
>>>>>>> Stashed changes
		Where(sq.Eq{"id": uid}).
		ToSql()

	if err != nil {
<<<<<<< Updated upstream
		return nil, apperrors.ErrCouldntBuildQuery
=======
		return nil, apperrors.ErrCouldNotBuildQuery
>>>>>>> Stashed changes
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
// или возвращает ошибки ...
func (s *PostgresUserStorage) Create(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	sql, args, err := sq.
<<<<<<< Updated upstream
		Insert("public.User").
=======
		Insert("public.user").
>>>>>>> Stashed changes
		Columns("email", "password_hash").
		Values(info.Email, info.PasswordHash).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
<<<<<<< Updated upstream
		return nil, apperrors.ErrCouldntBuildQuery
=======
		return nil, apperrors.ErrCouldNotBuildQuery
>>>>>>> Stashed changes
	}

	query := s.db.QueryRow(ctx, sql, args...)

	user := entities.User{}
	if query.Scan(&user.ID) != nil {
		return nil, apperrors.ErrUserNotCreated
	}

	user.Email = info.Email
	user.PasswordHash = info.PasswordHash

	return &user, nil
}

// Update
// обновляет пользователя в БД
// или возвращает ошибки ...
<<<<<<< Updated upstream
func (s *PostgresUserStorage) Update(ctx context.Context, updatedUser entities.User) (*entities.User, error) {
	sql, args, err := sq.
		Update("public.User").
=======
func (s *PostgresUserStorage) Update(ctx context.Context, updatedUser entities.User) error {
	sql, args, err := sq.
		Update("public.user").
>>>>>>> Stashed changes
		Set("email", updatedUser.Email).
		Set("name", updatedUser.Name).
		Set("surname", updatedUser.Surname).
		Set("avatar_url", updatedUser.AvatarURL).
		Set("password_hash", updatedUser.PasswordHash).
		Where(sq.Eq{"id": updatedUser.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
<<<<<<< Updated upstream
		return nil, apperrors.ErrCouldntBuildQuery
=======
		return apperrors.ErrCouldNotBuildQuery
>>>>>>> Stashed changes
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if query.Scan() != nil {
<<<<<<< Updated upstream
		return nil, apperrors.ErrUserNotUpdated
	}

	return &updatedUser, nil
=======
		return apperrors.ErrUserNotUpdated
	}

	return nil
>>>>>>> Stashed changes
}

// Delete
// удаляет данного пользователя в БД по id
// или возвращает ошибки ...
func (s *PostgresUserStorage) Delete(ctx context.Context, id uint64) error {
	sql, args, err := sq.
<<<<<<< Updated upstream
		Delete("public.User").
=======
		Delete("public.user").
>>>>>>> Stashed changes
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
<<<<<<< Updated upstream
		return apperrors.ErrCouldntBuildQuery
=======
		return apperrors.ErrCouldNotBuildQuery
>>>>>>> Stashed changes
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if query.Scan() != nil {
		return apperrors.ErrUserNotDeleted
	}

	return nil
}
