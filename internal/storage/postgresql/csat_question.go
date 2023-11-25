package postgresql

import (
	"context"
	"log"
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
	sql, args, err := sq.
		Insert("public.question").
		Columns("content", "id_type").
		Values(info.Content, info.TypeID).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built list query\n\t", sql, "\nwith args\n\t", args)

	question := dto.CSATQuestionFull{
		Content: info.Content,
	}

	query := s.db.QueryRow(ctx, sql, args...)
	if err := query.Scan(&question.ID); err != nil {
		return nil, apperrors.ErrCouldNotCreateQuestion
	}

	return &question, nil
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
	sql, args, err := sq.
		Update("public.question").
		Set("name", info.Content).
		Set("id_type", info.Type).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	if _, err = s.db.Exec(ctx, sql, args...); err != nil {
		return apperrors.ErrQuestionNotUpdated
	}

	return nil
}

// Delete
// удаляет CSAT вопрос в БД
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) Delete(ctx context.Context, id dto.CSATQuestionID) error {
	sql, args, err := sq.
		Delete("public.question").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	if _, err = s.db.Exec(ctx, sql, args...); err != nil {
		return apperrors.ErrQuestionNotDeleted
	}

	return apperrors.ErrQuestionNotDeleted
}

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) GetQuestionType(ctx context.Context, id dto.CSATQuestionID) (*entities.QuestionType, error) {
	sql, args, err := sq.
		Select("public.question_type.id", "public.question_type.name", "public.question_type.max_score").
		From("public.question_type").
		Join("public.question ON public.question.id_type = public.question_type.id").
		Where(sq.Eq{"public.question.id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetQuestionType
	}

	questionType := entities.QuestionType{}
	err = row.Scan(
		&questionType.ID,
		&questionType.Name,
		&questionType.MaxRating,
	)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetQuestionType
	}

	return &questionType, nil
}

// GetQuestionType
// возвращает тип CSAT вопроса по его названию
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) GetQuestionTypeWithName(ctx context.Context, typeName dto.CSATQuestionTypeName) (*entities.QuestionType, error) {
	sql, args, err := sq.
		Select("id", "name", "max_score").
		From("public.question_type").
		Where(sq.Eq{"name": typeName.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetQuestionType
	}

	questionType := entities.QuestionType{}
	err = row.Scan(
		&questionType.ID,
		&questionType.Name,
		&questionType.MaxRating,
	)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetQuestionType
	}

	return &questionType, nil
}
