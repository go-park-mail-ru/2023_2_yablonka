package postgresql

import (
	"context"
	"database/sql"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"

	sq "github.com/Masterminds/squirrel"
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
	sql, args, err := sq.
		Insert("public.answer").
		Columns("id_user", "id_question", "score").
		Values(info.UserID, info.QuestionID, info.Rating).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return apperrors.ErrCouldNotCreateAnswer
	}

	return nil
}
