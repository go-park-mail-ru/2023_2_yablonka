package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
)

// PostgresAuthStorage
// Хранилище данных в PostgreSQL
type PostgresAuthStorage struct {
	db *sql.DB
}

// NewLocalAuthStorage
// возвращает локальное хранилище сессий
func NewAuthStorage(db *sql.DB) *PostgresAuthStorage {
	return &PostgresAuthStorage{
		db: db,
	}
}

// CreateSession
// сохраняет сессию в хранилище, возвращает ID сесссии для куки
// или возвращает ошибку ErrTokenNotGenerated (500), ErrCouldntBuildQuery (500), ErrSessionNotCreated(500)
func (s PostgresAuthStorage) CreateSession(ctx context.Context, session *entities.Session) error {
	funcName := "PostgresAuthStorage.CreateSession"
	errorMessage := "Creating session in failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresAuthStorage.CreateSession FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresAuthStorage.CreateSession <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Insert("public.session").
		Columns("id_user", "expiration_date", "id_session").
		Values(session.UserID, session.ExpiryDate, session.SessionID).
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
		return apperrors.ErrSessionNotCreated
	}
	logger.DebugFmt("Executed query", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresAuthStorage.CreateSession SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}

// GetSession
// находит сессию по строке-токену
// или возвращает ошибку apperrors.ErrSessionNotFound (401), ErrCouldntBuildQuery(500)
func (s PostgresAuthStorage) GetSession(ctx context.Context, token dto.SessionToken) (*entities.Session, error) {
	funcName := "PostgresAuthStorage.GetSession"
	errorMessage := "Getting session in failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresAuthStorage.GetSession FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresAuthStorage.GetSession <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Select(allSessionFields...).
		From("public.session").
		Where(sq.Eq{"id_session": token.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	row := s.db.QueryRow(query, args...)
	session := entities.Session{
		SessionID: token.ID,
	}
	err = row.Scan(
		&session.UserID,
		&session.ExpiryDate,
	)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrSessionNotFound
	}
	logger.DebugFmt("Executed query", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresAuthStorage.GetSession SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &session, nil
}

// DeleteSession
// удаляет сессию по ID из хранилища, если она существует
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (s PostgresAuthStorage) DeleteSession(ctx context.Context, token dto.SessionToken) error {
	funcName := "PostgresAuthStorage.DeleteSession"
	errorMessage := "Deleting session in failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresAuthStorage.DeleteSession FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresAuthStorage.DeleteSession <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Delete("public.session").
		Where(sq.Eq{"id_session": token.ID}).
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
		return apperrors.ErrSessionNotFound
	}
	logger.DebugFmt("Executed query", funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresAuthStorage.DeleteSession SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}
