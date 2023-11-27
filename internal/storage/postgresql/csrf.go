package postgresql

import (
	"context"
	"database/sql"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
)

// PostgresAuthStorage
// Хранилище данных в PostgreSQL
type PostgresCSRFStorage struct {
	db *sql.DB
}

// NewCSRFStorage
// возвращает локальное хранилище CSRF
func NewCSRFStorage(db *sql.DB) *PostgresCSRFStorage {
	return &PostgresCSRFStorage{
		db: db,
	}
}

// Create
// сохраняет CSRF в хранилище, возвращает токен
// или возвращает ошибку apperrors.ErrTokenNotGenerated (500)
func (s PostgresCSRFStorage) Create(ctx context.Context, csrf *entities.CSRF) error {
	log.Println("Storage -- saving session")

	sql, args, err := sq.
		Insert("public.csrf").
		Columns("id_user", "expiration_date", "token").
		Values(csrf.UserID, csrf.ExpirationDate, csrf.Token).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Failed to build query")
		return apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Query formed:", sql)

	_, err = s.db.Exec(sql, args...)

	log.Println("Queried DB")

	if err != nil {
		log.Println("Failed to store CSRF")
		log.Println("Returned error", err.Error())
		return apperrors.ErrSessionNotCreated
	}

	log.Println("Stored CSRF")

	return nil
}

// Get
// находит CSRF по токену
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (s PostgresCSRFStorage) Get(ctx context.Context, token dto.CSRFToken) (*entities.CSRF, error) {
	sql, args, err := sq.
		Select("id_user", "expiration_date").
		From("public.csrf").
		Where(sq.Eq{"token": token.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(sql, args...)

	csrf := entities.CSRF{}
	err = row.Scan(
		&csrf.UserID,
		&csrf.ExpirationDate,
	)
	if err != nil {
		log.Println("")
		return nil, apperrors.ErrCSRFNotFound
	}

	return &csrf, nil
}

// Delete
// удаляет CSRF по токену из хранилища, если она существует
// или возвращает ошибку apperrors.ErrSessionNotFound (401)
func (s PostgresCSRFStorage) Delete(ctx context.Context, token dto.CSRFToken) error {
	sql, args, err := sq.
		Delete("public.csrf").
		Where(sq.Eq{"token": token.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrSessionNotFound
	}

	return nil
}
