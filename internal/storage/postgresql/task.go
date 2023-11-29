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
	funcName := "PostgreSQLTaskStorage.Create"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sql, args, err := sq.
		Insert("public.task").
		Columns(newTaskFields...).
		Values(info.ListID, info.Name, info.ListPosition).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	logger.Debug("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

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
func (s *PostgresTaskStorage) Read(ctx context.Context, id dto.TaskID) (*dto.SingleTaskInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetLists"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	query, args, err := sq.
		Select(allTaskFields...).
		From("public.task").
		Join("public.task_user ON public.task.id = public.task_user.id_task").
		Join("public.comment ON public.task.id = public.comment.id_task").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	task := dto.SingleTaskInfo{}
	row := s.db.QueryRow(query, args...)
	if err = row.Scan(
		&task.ID,
		&task.ListID,
		&task.DateCreated,
		&task.Name,
		&task.Description,
		&task.ListPosition,
		&task.Start,
		&task.End,
		&task.UserIDs,
		&task.CommentIDs,
	); err != nil {
		log.Println("Failed to query DB with error", err.Error())
		return nil, apperrors.ErrCouldNotGetTask
	}
	logger.Debug("Got task from DB", funcName, nodeName)

	return &task, nil
}

// ReadMany
// находит задание в БД по их id
// или возвращает ошибки ...
func (s *PostgresTaskStorage) ReadMany(ctx context.Context, id dto.TaskIDs) (*[]dto.SingleTaskInfo, error) {
	funcName := "PostgresTaskStorage.ReadMany"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	query, args, err := sq.
		Select(allTaskFields...).
		From("public.task").
		LeftJoin("public.task_user ON public.task.id = public.task_user.id_task").
		LeftJoin("public.comment ON public.task.id = public.comment.id_task").
		LeftJoin("public.checklist ON public.task.id = public.checklist.id_task").
		Where(sq.Eq{"public.task.id": id.Values}).
		GroupBy("public.task.id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		logger.Debug(err.Error(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotGetBoardUsers
	}
	defer rows.Close()
	logger.Debug("Got task rows", funcName, nodeName)

	tasks := []dto.SingleTaskInfo{}
	for rows.Next() {
		var task dto.SingleTaskInfo

		err = rows.Scan(
			&task.ID,
			&task.ListID,
			&task.DateCreated,
			&task.Name,
			&task.Description,
			&task.ListPosition,
			&task.Start,
			&task.End,
			(*pq.StringArray)(&task.UserIDs),
			(*pq.StringArray)(&task.CommentIDs),
			(*pq.StringArray)(&task.ChecklistIDs),
		)
		if err != nil {
			logger.Debug(err.Error(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotGetBoard
		}
		tasks = append(tasks, task)
	}
	logger.Debug("Got task from DB", funcName, nodeName)

	return &tasks, nil
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

// AddUser
// добавляет пользователя в карточку
// или возвращает ошибки ...
func (s PostgresTaskStorage) AddUser(ctx context.Context, info dto.AddTaskUserInfo) error {
	funcName := "PostgreSQLTaskStorage.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sql, args, err := sq.
		Insert("task_user").
		Columns(taskUserFields...).
		Values(info.UserID, info.TaskID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		logger.Debug("Insert failed with error "+err.Error(), funcName, nodeName)
		return apperrors.ErrCouldNotAddTaskUser
	}
	logger.Debug("query executed", funcName, nodeName)

	return nil
}

// RemoveUser
// удаляет пользователя из карточки
// или возвращает ошибки ...
func (s PostgresTaskStorage) RemoveUser(ctx context.Context, info dto.RemoveTaskUserInfo) error {
	funcName := "PostgreSQLTaskStorage.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sql, args, err := sq.
		Delete("task_user").
		Where(sq.And{
			sq.Eq{"id_user": info.UserID},
			sq.Eq{"id_task": info.UserID},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		logger.Debug("Delete failed with error "+err.Error(), funcName, nodeName)
		return apperrors.ErrCouldNotRemoveTaskUser
	}
	logger.Debug("query executed", funcName, nodeName)

	return nil
}
