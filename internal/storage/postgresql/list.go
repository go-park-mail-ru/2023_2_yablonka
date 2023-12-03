package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

// PostgresListStorage
// Хранилище данных в PostgreSQL
type PostgresListStorage struct {
	db *sql.DB
}

// NewListStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewListStorage(db *sql.DB) *PostgresListStorage {
	return &PostgresListStorage{
		db: db,
	}
}

// Create
// создает новый список в БД по данным
// или возвращает ошибки ...
func (s PostgresListStorage) Create(ctx context.Context, info dto.NewListInfo) (*entities.List, error) {
	sql, args, err := sq.
		Insert("public.list").
		Columns("name", "list_position", "description", "id_board").
		Values(info.Name, info.ListPosition, info.Description, info.BoardID).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built list query\n\t", sql, "\nwith args\n\t", args)

	list := entities.List{
		Name:         info.Name,
		BoardID:      info.BoardID,
		Description:  info.Description,
		ListPosition: info.ListPosition,
		Tasks:        []entities.Task{},
	}

	query := s.db.QueryRow(sql, args...)

	if err := query.Scan(&list.ID); err != nil {
		log.Println("Storage -- Failed to create list")
		return nil, apperrors.ErrListNotCreated
	}

	log.Println("Storage -- List created")

	return &list, nil
}

// Create
// GetWithID новый список задач в БД по данным
// или возвращает ошибки ...
func (s PostgresListStorage) GetTasksWithID(ctx context.Context, ids dto.ListIDs) (*[]dto.SingleTaskInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetTasksWithID"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	taskSql, args, err := sq.Select(allTaskAggFields...).
		From("public.task").
		Where(sq.Eq{"public.task.id_list": ids.Values}).
		LeftJoin("public.task_user ON public.task_user.id_task = public.task.id").
		GroupBy("public.task.id", "public.task.id_list").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+taskSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	taskRows, err := s.db.Query(taskSql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer taskRows.Close()
	logger.Debug("Got task info rows", funcName, nodeName)

	tasks := []dto.SingleTaskInfo{}
	for taskRows.Next() {
		var task dto.SingleTaskInfo

		err = taskRows.Scan(
			&task.ID,
			&task.ListID,
			&task.DateCreated,
			&task.Name,
			&task.Description,
			&task.ListPosition,
			&task.Start,
			&task.End,
			(*pq.StringArray)(&task.UserIDs),
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}
		tasks = append(tasks, task)
	}
	logger.Debug("Collected task info rows", funcName, nodeName)

	return &tasks, nil
}

// Update
// обновляет список в БД
// или возвращает ошибки ...
func (s PostgresListStorage) Update(ctx context.Context, info dto.UpdatedListInfo) error {
	sql, args, err := sq.
		Update("public.list").
		Set("name", info.Name).
		Set("description", info.Description).
		Set("list_position", info.ListPosition).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built list query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		log.Println(err)
		return apperrors.ErrListNotUpdated
	}
	log.Println("list updated")

	return nil
}

// Delete
// удаляет список в БД по id
// или возвращает ошибки ...
func (s PostgresListStorage) Delete(ctx context.Context, id dto.ListID) error {
	sql, args, err := sq.
		Delete("public.list").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrListNotDeleted
	}

	return nil
}
