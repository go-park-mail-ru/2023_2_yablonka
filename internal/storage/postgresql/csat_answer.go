package postgresql

import (
	"context"
	"server/internal/pkg/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresCSATAnswerStorage
// Хранилище данных в PostgreSQL
type PostgresCSATAnswerStorage struct {
	db *pgxpool.Pool
}

// NewCSATAnswerStorage
// возвращает PostgreSQL хранилище CSAT ответов
func NewCSATAnswerStorage(db *pgxpool.Pool) *PostgresCSATAnswerStorage {
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
