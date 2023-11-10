package postgresql

import (
	"context"
	"fmt"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresListStorage
// Хранилище данных в PostgreSQL
type PostgresTaskStorage struct {
	db *pgxpool.Pool
}

type TaskReturnValue struct {
	Task  entities.Task
	Users []entities.User
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
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	list := entities.Task{
		Name:         info.Name,
		ListID:       info.ListID,
		Description:  info.Description,
		ListPosition: info.ListPosition,
		Start:        info.Start,
		End:          info.End,
	}

	query := s.db.QueryRow(ctx, sql, args...)
	log.Println("Formed query\n\t", sql, "\nwith args\n\t", args)

	err = query.Scan(&list.ID)

	if err != nil {
		log.Println("Storage -- Task failed to create with error", err.Error())
		return nil, apperrors.ErrTaskNotCreated
	}

	return &list, nil
}

// Read
// находит задание в БД по его id
// или возвращает ошибки ...
func (s *PostgresTaskStorage) Read(ctx context.Context, id dto.TaskID) (*entities.Task, error) {
	sql, args, err := sq.
		Select(append(allTaskFields, allUserFields...)...).
		From("public.task").
		Join("public.task_user ON public.task.id = public.task_user.id_task").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetTask
	}
	defer rows.Close()

	// var task dto.TaskReturnValue

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[TaskReturnValue])
	fmt.Println(tasks)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	// var task entities.Task
	// for rows.Next() {
	// 	var user entities.User

	// 	err = rows.Scan(
	// 		&task,
	// 		&user,
	// 	)
	// 	if err != nil {
	// 		return nil, apperrors.ErrCouldNotGetTask
	// 	}

	// 	task.Users = append(task.Users, user)
	// }

	// return &task, nil
	return nil, nil
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
