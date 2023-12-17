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
	"github.com/google/uuid"
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
	funcName := "PostgresTaskStorage.Create"
	errorMessage := "Creating task failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresTaskStorage.Create FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresTaskStorage.Create <<<<<<<<<<<<<<<<<<<")

	sql, args, err := sq.
		Insert("public.task").
		Columns(newTaskFields...).
		Values(info.ListID, info.Name, info.ListPosition).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	task := entities.Task{
		Name:         info.Name,
		ListID:       info.ListID,
		ListPosition: info.ListPosition,
		Users:        []uint64{},
		Checklists:   []uint64{},
		Comments:     []uint64{},
	}
	query := s.db.QueryRow(sql, args...)
	err = query.Scan(&task.ID, &task.DateCreated)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrTaskNotCreated
	}
	logger.DebugFmt("Created task", requestID.String(), funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresTaskStorage.Create SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &task, nil
}

// Read
// находит задание в БД по его id
// или возвращает ошибки ...
func (s *PostgresTaskStorage) Read(ctx context.Context, id dto.TaskID) (*dto.SingleTaskInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetLists"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	query, args, err := sq.
		Select(allTaskFields2...).
		From("public.task").
		Join("public.task_user ON public.task.id = public.task_user.id_task").
		Join("public.comment ON public.task.id = public.comment.id_task").
		Where(sq.Eq{"public.task.id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

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
		(*pq.StringArray)(&task.UserIDs),
		(*pq.StringArray)(&task.CommentIDs),
	); err != nil {
		log.Println("Failed to query DB with error", err.Error())
		return nil, apperrors.ErrCouldNotGetTask
	}
	logger.DebugFmt("Got task from DB", requestID.String(), funcName, nodeName)

	return &task, nil
}

// ReadMany
// находит задание в БД по их id
// или возвращает ошибки ...
func (s *PostgresTaskStorage) ReadMany(ctx context.Context, id dto.TaskIDs) (*[]dto.SingleTaskInfo, error) {
	funcName := "PostgresTaskStorage.ReadMany"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	query, args, err := sq.
		Select(allTaskFields...).
		From("public.task").
		LeftJoin("public.task_user ON public.task.id = public.task_user.id_task").
		LeftJoin("public.comment ON public.task.id = public.comment.id_task").
		LeftJoin("public.checklist ON public.task.id = public.checklist.id_task").
		Where(sq.Eq{"public.task.id": id.Values}).
		GroupBy("public.task.id", "public.task.id_list").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotGetBoardUsers
	}
	defer rows.Close()
	logger.DebugFmt("Got task rows", requestID.String(), funcName, nodeName)

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
			logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotGetBoard
		}
		tasks = append(tasks, task)
	}
	logger.DebugFmt("Got task from DB", requestID.String(), funcName, nodeName)

	return &tasks, nil
}

// Update
// обновляет задание в БД
// или возвращает ошибки ...
func (s PostgresTaskStorage) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	query := sq.Update("public.task")
	/*
		if info.Start != nil {
			query = query.Set("task_start", &info.Start)
		}
		if info.End != nil {
			query = query.Set("task_end", &info.End)
		}
		log.Println(info.Start)
		log.Println(info.End)
	*/

	finalQuery, args, err := query.
		Set("name", info.Name).
		Set("description", info.Description).
		Set("list_position", info.ListPosition).
		Set("task_start", &info.Start).
		Set("task_end", &info.End).
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
	log.Println("Built board query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		log.Println("Failed to exec query with error", err.Error())
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
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	sql, args, err := sq.
		Insert("public.task_user").
		Columns("id_task", "id_user").
		Values(info.TaskID, info.UserID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		logger.DebugFmt("Insert failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotAddTaskUser
	}
	logger.DebugFmt("query executed", requestID.String(), funcName, nodeName)

	return nil
}

// RemoveUser
// удаляет пользователя из карточки
// или возвращает ошибки ...
func (s PostgresTaskStorage) RemoveUser(ctx context.Context, info dto.RemoveTaskUserInfo) error {
	funcName := "PostgreSQLTaskStorage.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	sql, args, err := sq.
		Delete("task_user").
		Where(sq.And{
			sq.Eq{"id_user": info.UserID},
			sq.Eq{"id_task": info.TaskID},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		logger.DebugFmt("Delete failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotRemoveTaskUser
	}
	logger.DebugFmt("query executed", requestID.String(), funcName, nodeName)

	return nil
}

// CheckAccess
// находит пользователя в доске
// или возвращает ошибки ...
func (s *PostgresTaskStorage) CheckAccess(ctx context.Context, info dto.CheckTaskAccessInfo) (bool, error) {
	funcName := "PostgresTaskStorage.CheckAccess"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	listSql, args, err := sq.Select("count(*)").
		From("public.task_user").
		Where(sq.And{
			sq.Eq{"id_task": info.TaskID},
			sq.Eq{"id_user": info.UserID},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+listSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	row := s.db.QueryRow(listSql, args...)
	logger.DebugFmt("Got user row", requestID.String(), funcName, nodeName)

	var count uint64
	if row.Scan(&count) != nil {
		return false, apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("checked database", requestID.String(), funcName, nodeName)

	return count > 0, nil
}

// Move
// переносит задание в другой список
// или возвращает ошибки ...
func (s PostgresTaskStorage) Move(ctx context.Context, taskMoveInfo dto.TaskMoveInfo) error {
	funcName := "PostgreSQLBoardStorage.GetById"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	taskIDs := make(map[uint64]int)
	for i, id := range taskMoveInfo.OldList.TaskIDs {
		taskIDs[id] = i
	}
	for i, id := range taskMoveInfo.NewList.TaskIDs {
		taskIDs[id] = i
	}
	keys := make([]uint64, 0, len(taskIDs))
	for k := range taskIDs {
		keys = append(keys, k)
	}

	log.Println(keys)

	caseBuilder := sq.Case()
	for _, id := range keys {
		caseBuilder = caseBuilder.
			When(sq.Eq{"id": fmt.Sprintf("%v", id)}, fmt.Sprintf("%v", taskIDs[id])).
			Else("list_position")
	}

	updateBuilder := sq.
		Update("public.task").
		Set("list_position", caseBuilder)

	if taskMoveInfo.NewList.ListID != taskMoveInfo.OldList.ListID {
		updateBuilder = updateBuilder.Set("id_list", sq.Case().When(sq.Eq{"id": fmt.Sprintf("%v", taskMoveInfo.TaskID)}, taskMoveInfo.NewList.ListID).Else("id_list"))
	}

	query, args, err := updateBuilder.
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t", requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
	}
	logger.DebugFmt("Commited changes", requestID.String(), funcName, nodeName)

	return nil
}
