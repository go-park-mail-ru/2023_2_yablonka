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

// PostgresListStorage
// Хранилище данных в PostgreSQL
type PostgresTaskStorage struct {
	db *sql.DB
}

// NewTaskStorage
// возвращает PostgreSQL хранилище заданий
func NewTaskStorage(db *sql.DB) *PostgresTaskStorage {
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
		Values(info.ListID, info.Name, info.ListPosition).
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
		ListPosition: info.ListPosition,
		Users:        []uint64{},
		Checklists:   []uint64{},
		Comments:     []uint64{},
	}

	log.Println("Storage -- Querying DB")
	query := s.db.QueryRow(sql, args...)

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
		Join("public.task_user ON public.task.id = public.task_user.id_task").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	task := entities.Task{}

	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Formed query\n\t", query1, "\nwith args\n\t", args)

	row := s.db.QueryRow(query1, args...)
	if err = row.Scan(
		&task.ID,
		&task.ListID,
		&task.DateCreated,
		&task.Name,
		&task.Description,
		&task.ListPosition,
		&task.Start,
		&task.End,
		&task.Users); err != nil {
		log.Println("Failed to query DB with error", err.Error())
		return nil, apperrors.ErrCouldNotGetTask
	}

	return &task, nil
}

// Update
// обновляет задание в БД
// или возвращает ошибки ...
func (s PostgresTaskStorage) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	query := sq.Update("public.task")
	if info.Start != nil {
		query = query.Set("start", &info.Start)
	}
	if info.End != nil {
		query = query.Set("end", &info.End)
	}

	finalQuery, args, err := query.
		Set("name", info.Name).
		Set("description", info.Description).
		Set("list_position", info.ListPosition).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Formed query\n\t", finalQuery, "\nwith args\n\t", args)

	_, err = s.db.Exec(finalQuery, args...)

	if err != nil {
		log.Println(err)
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

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrTaskNotDeleted
	}

	return nil
}
