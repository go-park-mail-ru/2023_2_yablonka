package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
)

// PostgresCSATQuestionStorage
// Хранилище данных в PostgreSQL
type PostgresCSATQuestionStorage struct {
	db *sql.DB
}

// NewCSATQuestionStorage
// возвращает PostgreSQL хранилище CSAT вопросов
func NewCSATQuestionStorage(db *sql.DB) *PostgresCSATQuestionStorage {
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

	query := s.db.QueryRow(sql, args...)
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

	rows, err := s.db.Query(sql, args...)
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

// GetStats
// возвращает вопросы со статистикой по их ответам
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) GetStats(ctx context.Context) (*[]dto.QuestionWithStats, error) {
	questionQuery, args, err := sq.
		Select("public.question.id", "public.question.content", "public.question_type.name").
		From("public.question").
		Join("public.question_type ON public.question_type.id = public.question.id_type").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built questions query\n\t", questionQuery, "\nwith args\n\t", args)

	rows, err := s.db.Query(questionQuery, args...)
	if err != nil {
		log.Println("Storage -- DB questions query failed with error", err.Error())
		return nil, err
	}
	log.Println("questions got")
	defer rows.Close()

	questions := map[uint64]dto.QuestionWithStats{}
	var questionsID []uint64
	for rows.Next() {
		var question dto.QuestionWithStats

		err = rows.Scan(
			&question.ID,
			&question.Content,
			&question.Type,
		)
		if err != nil {
			fmt.Println("Scanning failed due to error", err.Error())
			return nil, err
		}
		questions[question.ID] = question
		questionsID = append(questionsID, question.ID)
	}

	statsQuery, args, err := sq.
		Select("id_question", "score", "COUNT(score)", "AVG(score)").
		From("public.answer").
		Join("public.question ON public.answer.id_question = public.question.id").
		Where(sq.Eq{"id_question": questionsID}).
		OrderBy("score").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built boards query\n\t", statsQuery, "\nwith args\n\t", args)

	rows, err = s.db.Query(statsQuery, args...)
	if err != nil {
		log.Println("Storage -- DB stats query failed with error", err.Error())
		return nil, err
	}
	log.Println("Stats got")
	defer rows.Close()

	statRows := []dto.RatingStatsWithQuestionID{}
	for rows.Next() {
		var board dto.RatingStatsWithQuestionID

		err = rows.Scan(
			&board.QuestionID,
			&board.Rating,
			&board.Count,
			&board.Average,
		)
		if err != nil {
			fmt.Println("Scanning failed due to error", err.Error())
			return nil, err
		}
		statRows = append(statRows, board)
	}
	log.Println("Stats collected")

	for _, statRow := range statRows {
		board := dto.RatingStats{
			Rating:  statRow.Rating,
			Count:   statRow.Count,
			Average: statRow.Average,
		}
		question := questions[statRow.QuestionID]
		question.Stats = append(question.Stats, board)
		questions[statRow.QuestionID] = question
	}
	log.Println("Stats appended to workspaces")

	var result []dto.QuestionWithStats
	for _, value := range questions {
		result = append(result, value)
	}

	return &result, nil
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

	if _, err = s.db.Exec(sql, args...); err != nil {
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

	if _, err = s.db.Exec(sql, args...); err != nil {
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

	row := s.db.QueryRow(sql, args...)
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

	row := s.db.QueryRow(sql, args...)
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
