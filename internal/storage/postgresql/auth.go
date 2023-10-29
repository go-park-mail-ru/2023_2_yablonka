package postgresql

import (
	"context"
	"crypto/rand"
	"math/big"
	"server/internal/apperrors"
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
func (s PostgresAuthStorage) CreateSession(ctx context.Context, session *entities.Session, sidLength uint) (string, error) {
	token, err := generateSessionID(sidLength)
	if err != nil {
		return "", apperrors.ErrTokenNotGenerated
	}

	sql, args, err := sq.
		Insert("public.Session").
		Columns("id_user", "expiration_date", "token").
		Values(session.UserID, session.ExpiryDate, token).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", apperrors.ErrCouldntBuildQuery
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if query.Scan() != nil {
		return "", apperrors.ErrSessionNotCreated
	}

	return token, nil
}

// GetSession
// находит сессию по строке-токену
// или возвращает ошибку apperrors.ErrSessionNotFound (401), ErrCouldntBuildQuery(500)
func (s PostgresAuthStorage) GetSession(ctx context.Context, token string) (*entities.Session, error) {
	sql, args, err := sq.
		Select(allSessionFields...).
		From("Session").
		Where(sq.Eq{"token": token}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
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
func (s PostgresAuthStorage) DeleteSession(ctx context.Context, token string) error {
	sql, args, err := sq.
		Delete("Session").
		Where(sq.Eq{"token": token}).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldntBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)

	if row.Scan() != nil {
		return apperrors.ErrSessionNotFound
	}

	return nil
}

// GenerateSessionID
// возвращает alphanumeric строку, собранную криптографически безопасным PRNG
func generateSessionID(n uint) (string, error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]rune, n)
	for i := range buf {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			return "", err
		}
		buf[i] = letterRunes[j.Int64()]
	}
	return string(buf), nil
}
