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
	funcName := "PostgresUserStorage.Create"
	errorMessage := "Creating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.Create FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.Create <<<<<<<<<<<<<<<<<<<")

	log.Println("Looking for user with login", login.Value)

	sql, args, err := sq.
		Select(allUserFields...).
		From("public.user").
		Where(sq.Eq{"email": login.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}

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
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

// GetUserByID
// находит пользователя в БД по его id
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetWithID(ctx context.Context, id dto.UserID) (*entities.User, error) {
	funcName := "PostgresUserStorage.GetWithID"
	errorMessage := "Creating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.GetWithID FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.GetWithID <<<<<<<<<<<<<<<<<<<")

	sql, args, err := sq.
		Select(allUserFields...).
		From("public.user").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	row := s.db.QueryRow(sql, args...)
	user := entities.User{}
	if row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Surname, &user.AvatarURL, &user.Description) != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrUserNotFound
	}
	logger.DebugFmt("Got user", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.GetWithID SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &user, nil
}

// GetLoginInfoWithID
// находит данные логина пользователя в БД по id
// или возвращает ошибки ...
func (s *PostgresUserStorage) GetLoginInfoWithID(ctx context.Context, id dto.UserID) (*dto.LoginInfo, error) {
	funcName := "PostgresUserStorage.GetLoginInfoWithID"
	errorMessage := "Creating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.GetLoginInfoWithID FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.GetLoginInfoWithID <<<<<<<<<<<<<<<<<<<")

	sql, args, err := sq.
		Select("email", "password_hash").
		From("public.user").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	row := s.db.QueryRow(sql, args...)
	loginInfo := dto.LoginInfo{}
	err = row.Scan(&loginInfo.Email, &loginInfo.PasswordHash)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrUserNotFound
	}
	logger.DebugFmt("Got user", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.GetLoginInfoWithID SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &loginInfo, nil
}

// Create
// создает нового пользователя в БД по данным
// или возвращает ошибки ...
func (s *PostgresUserStorage) Create(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	funcName := "PostgresUserStorage.Create"
	errorMessage := "Creating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.Create FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.Create <<<<<<<<<<<<<<<<<<<")

	defaultAvatar := "avatar.jpg"

	query, args, err := sq.
		Insert("public.user").
		Columns("email", "password_hash", "avatar_url").
		Values(info.Email, info.PasswordHash, defaultAvatar).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	user := entities.User{
		Email:        info.Email,
		PasswordHash: info.PasswordHash,
		AvatarURL:    &defaultAvatar,
	}
	row := s.db.QueryRow(query, args...)
	err = row.Scan(&user.ID)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrUserNotCreated
	}
	logger.DebugFmt("Got user", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Create SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &user, nil
}

// UpdatePassword
// обновляет пароль пользователя в БД
// или возвращает ошибки ...
func (s *PostgresUserStorage) UpdatePassword(ctx context.Context, info dto.PasswordHashesInfo) error {
	funcName := "PostgresUserStorage.UpdatePassword"
	errorMessage := "Updating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.UpdatePassword FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.UpdatePassword <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Update("public.user").
		Set("password_hash", info.NewPasswordHash).
		Where(sq.Eq{"id": info.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrUserNotUpdated
	}
	logger.DebugFmt("Password updated", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.UpdatePassword SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}

// UpdateProfile
// обновляет профиль пользователя в БД
// или возвращает ошибки ...
func (s *PostgresUserStorage) UpdateProfile(ctx context.Context, info dto.UserProfileInfo) error {
	funcName := "PostgresUserStorage.UpdateProfile"
	errorMessage := "Updating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.UpdateProfile FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.UpdateProfile <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Update("public.user").
		Set("name", info.Name).
		Set("surname", info.Surname).
		Set("description", info.Description).
		Where(sq.Eq{"id": info.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrUserNotUpdated
	}
	logger.DebugFmt("Profile updated", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.UpdateProfile SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}

// UpdateAvatarUrl
// обновляет аватарку пользователя в БД
// или возвращает ошибки ...
func (s *PostgresUserStorage) UpdateAvatarUrl(ctx context.Context, info dto.UserImageUrlInfo) error {
	funcName := "PostgresUserStorage.UpdateAvatarUrl"
	errorMessage := "Updating user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.UpdateAvatarUrl FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.UpdateAvatarUrl <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Update("public.user").
		Set("avatar_url", info.Url).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrUserNotUpdated
	}
	logger.DebugFmt("Avatar updated", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.UpdateAvatarUrl SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}

// Delete
// удаляет данного пользователя в БД по id
// или возвращает ошибки ...
func (s *PostgresUserStorage) Delete(ctx context.Context, id dto.UserID) error {
	funcName := "PostgresUserStorage.Delete"
	errorMessage := "Deleting user failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresUserStorage.Delete FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.Delete <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Delete("public.user").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrUserNotDeleted
	}
	logger.DebugFmt("User deleted", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresUserStorage.Delete SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}
