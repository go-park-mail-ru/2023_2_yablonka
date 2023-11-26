package postgresql

import (
	"context"
	"database/sql"
	"server/internal/pkg/dto"
)

// PostgresCSATAnswerStorage
// Хранилище данных в PostgreSQL
type PostgresCSATAnswerStorage struct {
	db *sql.DB
}

// NewCSATAnswerStorage
// возвращает PostgreSQL хранилище CSAT ответов
func NewCSATAnswerStorage(db *sql.DB) *PostgresCSATAnswerStorage {
	return &PostgresCSATAnswerStorage{
		db: db,
	}
}

// Create
// создает новый ответ CSAT опроса в БД по данным
// или возвращает ошибки ...
func (s PostgresCSATAnswerStorage) Create(ctx context.Context, info dto.NewCSATAnswer) error {
	return nil
}
