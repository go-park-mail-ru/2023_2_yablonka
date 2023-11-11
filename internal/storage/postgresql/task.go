package postgresql

import (
	"context"
	"fmt"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresListStorage
// Хранилище данных в PostgreSQL
type PostgresTaskStorage struct {
	db *pgxpool.Pool
}

// NewTaskStorage
// возвращает PostgreSQL хранилище заданий
func NewTaskStorage(db *pgxpool.Pool) *PostgresTaskStorage {
	return &PostgresTaskStorage{
		db: db,
	}
}

// Create
// создает новое задание в БД по данным
// или возвращает ошибки ...
func (s PostgresTaskStorage) Create(ctx context.Context, info dto.NewTaskInfo) (*entities.Task, error) {
	sql, args, err := sq.
		Insert("public.task").
		Columns(newTaskFields...).
		Values(info.ListID, info.Name, info.Description, info.ListPosition, info.Start, info.End).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Formed query\n\t", sql, "\nwith args\n\t", args)

	task := entities.Task{
		Name:         info.Name,
		ListID:       info.ListID,
		Description:  info.Description,
		ListPosition: info.ListPosition,
		Start:        info.Start,
		End:          info.End,
	}

	log.Println("Storage -- Querying DB", err.Error())
	query := s.db.QueryRow(ctx, sql, args...)

	err = query.Scan(&task.ID, &task.DateCreated)

	if err != nil {
		log.Println("Storage -- Task failed to create with error", err.Error())
		return nil, apperrors.ErrTaskNotCreated
	}

	return &task, nil
}

// Read
// находит задание в БД по его id
// или возвращает ошибки ...
func (s *PostgresTaskStorage) Read(ctx context.Context, id dto.TaskID) (*entities.Task, error) {
	query1, args, err := sq.
		Select(allTaskFields...).
		From("public.task").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	task := entities.Task{}

	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Formed query\n\t", query1, "\nwith args\n\t", args)

	row := s.db.QueryRow(context.Background(), query1, args...)
	if err = row.Scan(&task.ID, &task.ListID, &task.DateCreated, &task.Name, &task.Description, &task.ListPosition, &task.Start, &task.End); err != nil {
		log.Println("Failed to query DB with error", err.Error())
		return nil, apperrors.ErrCouldNotGetTask
	}

	query2, args, err := sq.
		Select(allUserFields...).
		From("public.task_user").
		Join("public.user ON public.user.id = public.task_user.id_user").
		Where(sq.Eq{"public.task_user.id_task": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Formed query\n\t", query2, "\nwith args\n\t", args)

	rows, err := s.db.Query(context.Background(), query2, args...)
	if err != nil {
		log.Println("Storage -- Failed to query DB with error", err.Error())
		return nil, apperrors.ErrCouldNotGetUser
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[entities.User])
	if err != nil {
		fmt.Println("Failed collecting users with error:", err.Error())
		return nil, apperrors.ErrCouldNotGetTask
	}

	task.Users = append(task.Users, users...)

	return &task, nil
}

// Update
// обновляет задание в БД
// или возвращает ошибки ...
func (s PostgresTaskStorage) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	sql, args, err := sq.
		Update("public.task").
		Set("name", info.Name).
		Set("description", info.Description).
		Set("start", info.Start).
		Set("end", info.End).
		Set("list_position", info.ListPosition).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrTaskNotUpdated
	}

	return nil
}

// Delete
// удаляет задание в БД по id
// или возвращает ошибки ...
func (s PostgresTaskStorage) Delete(ctx context.Context, id dto.TaskID) error {
	sql, args, err := sq.
		Delete("public.task").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrTaskNotDeleted
	}

	return nil
}
