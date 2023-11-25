package postgresql

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresCSATQuestionStorage
// Хранилище данных в PostgreSQL
type PostgresCSATQuestionStorage struct {
	db *pgxpool.Pool
}

// NewCSATQuestionStorage
// возвращает PostgreSQL хранилище CSAT вопросов
func NewCSATQuestionStorage(db *pgxpool.Pool) *PostgresCSATQuestionStorage {
	return &PostgresCSATQuestionStorage{
		db: db,
	}
}

// Create
// создает новый CSAT вопрос в БД по данным
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) Create(ctx context.Context, info dto.NewCSATQuestion) (*dto.CSATQuestionFull, error) {
	return &dto.CSATQuestionFull{}, nil
}

// Create
// создает новый CSAT вопрос в БД по данным
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) GetAll(ctx context.Context) (*[]dto.CSATQuestionFull, error) {
	sql, args, err := sq.
		Select("public.question.id", "public.question.content", "public.question_type.name").
		From("public.question").
		Join("public.question_type ON public.question.id_type = public.question_type.id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetQuestions
	}

	questions := []dto.CSATQuestionFull{}
	for rows.Next() {
		var question dto.CSATQuestionFull

		err = rows.Scan(
			&question.ID,
			&question.Content,
			&question.Type,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetQuestions
		}
		questions = append(questions, question)
	}

	return &questions, nil
}

// Update
// обновляет CSAT вопрос в БД
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) Update(ctx context.Context, info dto.UpdatedCSATQuestion) error {
	return nil
}

// Delete
// удаляет CSAT вопрос в БД
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	return nil
}

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) GetQuestionType(ctx context.Context, id dto.CSATQuestionID) (*entities.QuestionType, error) {
	return &entities.QuestionType{}, nil
}
