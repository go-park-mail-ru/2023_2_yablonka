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
)

// PostgresTagStorage
// Хранилище данных в PostgreSQL
type PostgresTagStorage struct {
	db *sql.DB
}

// NewTagStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewTagStorage(db *sql.DB) *PostgresTagStorage {
	return &PostgresTagStorage{
		db: db,
	}
}

// Create
// создает новый чеклист в БД по данным
// или возвращает ошибки ...
func (s PostgresTagStorage) Create(ctx context.Context, info dto.NewTagInfo) (*entities.Tag, error) {
	funcName := "PostgresTagStorage.Create"
	logger, ok := ctx.Value(dto.LoggerKey).(logger.ILogger)
	if !ok {
		return nil, apperrors.ErrNoLoggerFound
	}
	requestID, ok := ctx.Value(dto.RequestIDKey).(uuid.UUID)
	if !ok {
		logger.DebugFmt("No request ID found", requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrNoRequestIDFound
	}

	query1, args, err := sq.
		Insert("public.tag").
		Columns("name", "color").
		Values(info.Name, info.Color).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query1+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.DebugFmt("Failed to start transaction with error "+err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotBeginTransaction
	}
	logger.DebugFmt("Transaction started", requestID.String(), funcName, nodeName)

	newTag := &entities.Tag{
		Name:    info.Name,
		Color:   info.Color,
		TaskID:  info.TaskID,
		BoardID: info.BoardID,
	}
	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&newTag.ID); err != nil {
		logger.DebugFmt("Creating tag failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotScanRows
	}
	logger.DebugFmt("Tag created", requestID.String(), funcName, nodeName)

	query2, args, err := sq.
		Insert("public.tag_board").
		Columns("id_tag", "id_board").
		Values(newTag.ID, newTag.BoardID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query2+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = tx.Exec(query2, args...)
	if err != nil {
		logger.DebugFmt("Failed to execute query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotExecuteQuery
	}
	logger.DebugFmt("Connection between board and tag created", requestID.String(), funcName, nodeName)

	query3, args, err := sq.
		Insert("public.tag_task").
		Columns("id_tag", "id_task").
		Values(newTag.ID, newTag.TaskID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query3+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = tx.Exec(query3, args...)
	if err != nil {
		logger.DebugFmt("Failed to execute query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotExecuteQuery
	}
	logger.DebugFmt("Board linked to creator", requestID.String(), funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt("Failed to commit changes with:"+err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotCommit
	}
	logger.DebugFmt("Changes commited", requestID.String(), funcName, nodeName)

	return newTag, nil
}

// Update
// обновляет тэг в БД
// или возвращает ошибки ...
func (s PostgresTagStorage) Update(ctx context.Context, info dto.UpdatedTagInfo) error {
	funcName := "PostgresTagStorage.Update"
	logger, ok := ctx.Value(dto.LoggerKey).(logger.ILogger)
	if !ok {
		return apperrors.ErrNoLoggerFound
	}
	requestID, ok := ctx.Value(dto.RequestIDKey).(uuid.UUID)
	if !ok {
		logger.DebugFmt("No request ID found", requestID.String(), funcName, nodeName)
		return apperrors.ErrNoRequestIDFound
	}

	query, args, err := sq.
		Update("public.tag").
		Set("name", info.Name).
		Set("color", info.Color).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt("Failed to update tag with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotExecuteQuery
	}
	logger.DebugFmt("Tag updated", requestID.String(), funcName, nodeName)

	return nil
}

// AddToTask
// добавляет тэг в задание в БД
// или возвращает ошибки ...
func (s PostgresTagStorage) AddToTask(ctx context.Context, ids dto.TagAndTaskIDs) error {
	funcName := "PostgresTagStorage.AddToTask"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	query, args, err := sq.
		Insert("public.tag_task").
		Columns("id_tag", "id_task").
		Values(ids.TagID, ids.TaskID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt("Failed to update tag_task with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrTagNotAddedToTask
	}
	log.Println("Tag added to task")

	return nil
}

// RemoveFromTask
// удаляет тэг из задания в БД
// или возвращает ошибки ...
func (s PostgresTagStorage) RemoveFromTask(ctx context.Context, ids dto.TagAndTaskIDs) error {
	funcName := "PostgresTagStorage.RemoveFromTask"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	query, args, err := sq.
		Delete("public.tag_task").
		Where(sq.And{
			sq.Eq{"id_task": ids.TaskID},
			sq.Eq{"id_tag": ids.TagID},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt("Failed to update tag_task with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrTagNotRemovedFromTask
	}
	log.Println("Tag removed from task")

	return nil
}

// Delete
// удаляет чеклист в БД по id
// или возвращает ошибки ...
func (s PostgresTagStorage) Delete(ctx context.Context, id dto.TagID) error {
	funcName := "PostgresTagStorage.Delete"
	logger, ok := ctx.Value(dto.LoggerKey).(logger.ILogger)
	if !ok {
		return apperrors.ErrNoLoggerFound
	}
	requestID, ok := ctx.Value(dto.RequestIDKey).(uuid.UUID)
	if !ok {
		logger.DebugFmt("No request ID found", requestID.String(), funcName, nodeName)
		return apperrors.ErrNoRequestIDFound
	}

	query, args, err := sq.
		Delete("public.tag").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt("Failed to update tag with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotExecuteQuery
	}
	log.Println("Tag deleted")

	return nil
}
