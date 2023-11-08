package postgresql

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresAuthStorage
// Хранилище данных в PostgreSQL
type PostgresAuthStorage struct {
	db *pgxpool.Pool
}

// NewLocalAuthStorage
// возвращает локальное хранилище сессий
func NewAuthStorage(db *pgxpool.Pool) *PostgresAuthStorage {
	return &PostgresAuthStorage{
		db: db,
	}
}

// CreateSession
// сохраняет сессию в хранилище, возвращает ID сесссии для куки
// или возвращает ошибку ErrTokenNotGenerated (500), ErrCouldntBuildQuery (500), ErrSessionNotCreated(500)
func (s PostgresAuthStorage) CreateSession(ctx context.Context, session *entities.Session) error {
	sql, args, err := sq.
		Insert("public.session").
		Columns("id_user", "expiration_date", "token").
		Values(session.UserID, session.ExpiryDate, session.Token).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrSessionNotCreated
	}

	return nil
}

// GetSession
// находит сессию по строке-токену
// или возвращает ошибку apperrors.ErrSessionNotFound (401), ErrCouldntBuildQuery(500)
func (s PostgresAuthStorage) GetSession(ctx context.Context, token dto.SessionToken) (*entities.Session, error) {
	sql, args, err := sq.
		Select(allSessionFields...).
		From("public.session").
		Where(sq.Eq{"token": token.Value}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	session := entities.Session{}
	if row.Scan(&session) != nil {
		return nil, apperrors.ErrSessionNotFound
	}

	return &session, nil
}

// DeleteSession
// удаляет сессию по ID из хранилища, если она существует
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (s PostgresAuthStorage) DeleteSession(ctx context.Context, token dto.SessionToken) error {
	sql, args, err := sq.
		Delete("public.session").
		Where(sq.Eq{"token": token.Value}).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	if row.Scan() != nil {
		return apperrors.ErrSessionNotFound
	}

	return nil
}
