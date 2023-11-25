package postgresql

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

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

// GetQuestionType
// возвращает тип CSAT вопроса по его id
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) GetQuestionType(ctx context.Context, id dto.CSATQuestionID) (*entities.QuestionType, error) {
	return &entities.QuestionType{}, nil
}

// Create
// создает новый CSAT вопрос в БД по данным
// или возвращает ошибки ...
func (s PostgresCSATQuestionStorage) Create(ctx context.Context, info dto.NewCSATQuestion) (*entities.CSATQuestion, error) {
	return &entities.CSATQuestion{}, nil
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
