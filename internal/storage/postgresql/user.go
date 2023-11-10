package postgresql

import (
	"context"
	"log"
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
func (s *PostgresUserStorage) GetWithLogin(ctx context.Context, login dto.UserLogin) (*entities.User, error) {
	sql, args, err := sq.
		Select(allUserFields...).
		From("public.user").
		Where(sq.Eq{"login": login}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	user := entities.User{}
	err = row.Scan(&user)
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// GetUserByID
// находит пользователя в БД по его id
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetWithID(ctx context.Context, id dto.UserID) (*entities.User, error) {
	sql, args, err := sq.
		Select(allUserFields...).
		From("public.user").
		Where(sq.Eq{"id": id.Value}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	user := entities.User{}
	if row.Scan(&user) != nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// GetLoginInfoWithID
// находит данные логина пользователя в БД по id
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetLoginInfoWithID(ctx context.Context, id dto.UserID) (*dto.LoginInfo, error) {
	sql, args, err := sq.
		Select("email", "password_hash").
		From("public.user").
		Where(sq.Eq{"id": id.Value}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Printf("query created")

	row := s.db.QueryRow(ctx, sql, args...)

	loginInfo := dto.LoginInfo{}
	err = row.Scan(&loginInfo)
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &loginInfo, nil
}

// Create
// создает нового пользователя в БД по данным
// или возвращает ошибки ...
func (s *PostgresUserStorage) Create(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	sql, args, err := sq.
		Insert("public.user").
		Columns("email", "password_hash").
		Values(info.Email, info.PasswordHash).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	user := entities.User{
		Email:        info.Email,
		PasswordHash: info.PasswordHash,
	}

	query := s.db.QueryRow(ctx, sql, args...)
	err = query.Scan(&user.ID)
	if err != nil {
		return nil, apperrors.ErrUserNotCreated
	}

	return &user, nil
}

// UpdatePassword
// обновляет пароль пользователя в БД
// или возвращает ошибки ...
func (s *PostgresUserStorage) UpdatePassword(ctx context.Context, info dto.PasswordHashesInfo) error {
	sql, args, err := sq.
		Update("public.user").
		Set("password_hash", info.NewPasswordHash).
		Where(sq.Eq{"id": info.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrUserNotUpdated
	}

	return nil
}

// UpdateProfile
// обновляет профиль пользователя в БД
// или возвращает ошибки ...
func (s *PostgresUserStorage) UpdateProfile(ctx context.Context, info dto.UserProfileInfo) error {
	sql, args, err := sq.
		Update("public.user").
		Set("name", info.Name).
		Set("surname", info.Surname).
		Set("description", info.Description).
		Where(sq.Eq{"id": info.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrUserNotUpdated
	}

	return nil
}

// UpdateAvatarUrl
// обновляет аватарку пользователя в БД
// или возвращает ошибки ...
func (s *PostgresUserStorage) UpdateAvatarUrl(ctx context.Context, info dto.ImageUrlInfo) error {
	sql, args, err := sq.
		Update("public.user").
		Set("avatar_url", info.Url).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrUserNotUpdated
	}

	return nil
}

// Delete
// удаляет данного пользователя в БД по id
// или возвращает ошибки ...
func (s *PostgresUserStorage) Delete(ctx context.Context, id dto.UserID) error {
	sql, args, err := sq.
		Delete("public.user").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrUserNotDeleted
	}

	return nil
}
