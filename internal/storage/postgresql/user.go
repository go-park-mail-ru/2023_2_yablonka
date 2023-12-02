package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
)

// LocalUserStorage
// Хранилище данных в PostgreSQL
type PostgresUserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *PostgresUserStorage {
	return &PostgresUserStorage{
		db: db,
	}
}

// GetUserByLogin
// находит пользователя в БД по почте
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetWithLogin(ctx context.Context, login dto.UserLogin) (*entities.User, error) {
	log.Println("Looking for user with login", login.Value)

	sql, args, err := sq.
		Select(allUserFields...).
		From("public.user").
		Where(sq.Eq{"email": login.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built query:", sql, "\nwith args:", args)

	row := s.db.QueryRow(sql, args...)

	user := entities.User{}
	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.Surname,
		&user.AvatarURL,
		&user.Description,
	)
	log.Println(user)
	log.Println(err)

	if err != nil {
		fmt.Println("Error", err.Error())
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
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(sql, args...)

	user := entities.User{}
	if row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Surname, &user.AvatarURL, &user.Description) != nil {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// GetLoginInfoWithID
// находит данные логина пользователя в БД по id
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetLoginInfoWithID(ctx context.Context, id dto.UserID) (*dto.LoginInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetTasksWithID"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sql, args, err := sq.
		Select("email", "password_hash").
		From("public.user").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	row := s.db.QueryRow(sql, args...)

	loginInfo := dto.LoginInfo{}
	err = row.Scan(&loginInfo.Email, &loginInfo.PasswordHash)
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}
	logger.Debug("Parsed result", funcName, nodeName)

	return &loginInfo, nil
}

// Create
// создает нового пользователя в БД по данным
// или возвращает ошибки ...
func (s *PostgresUserStorage) Create(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	defaultAvatar := "avatar.jpg"

	sql, args, err := sq.
		Insert("public.user").
		Columns("email", "password_hash", "avatar_url").
		Values(info.Email, info.PasswordHash, defaultAvatar).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	user := entities.User{
		Email:        info.Email,
		PasswordHash: info.PasswordHash,
		AvatarURL:    &defaultAvatar,
	}

	query := s.db.QueryRow(sql, args...)
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

	_, err = s.db.Exec(sql, args...)

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

	log.Println("Built query:", sql, "\nwith args:", args)

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		log.Println("Storage -- Failed to execute query with error", err.Error())
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

	_, err = s.db.Exec(sql, args...)

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

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrUserNotDeleted
	}

	return nil
}
