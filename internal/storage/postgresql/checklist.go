package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

// PostgresChecklistStorage
// Хранилище данных в PostgreSQL
type PostgresChecklistStorage struct {
	db *sql.DB
}

// NewChecklistStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewChecklistStorage(db *sql.DB) *PostgresChecklistStorage {
	return &PostgresChecklistStorage{
		db: db,
	}
}

// Create
// создает новый чеклист в БД по данным
// или возвращает ошибки ...
func (s PostgresChecklistStorage) Create(ctx context.Context, info dto.NewChecklistInfo) (*dto.ChecklistInfo, error) {
	sql, args, err := sq.
		Insert("public.checklist").
		Columns("name", "list_position", "id_task").
		Values(info.Name, info.ListPosition, info.TaskID).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built Checklist query\n\t", sql, "\nwith args\n\t", args)

	checklist := dto.ChecklistInfo{
		Name:         info.Name,
		TaskID:       info.TaskID,
		ListPosition: info.ListPosition,
		Items:        []string{},
	}
	query := s.db.QueryRow(sql, args...)
	if err := query.Scan(&checklist.ID); err != nil {
		log.Println("Storage -- Failed to create Checklist")
		return nil, apperrors.ErrChecklistNotCreated
	}

	log.Println("Storage -- Checklist created")

	return &checklist, nil
}

// Create
// GetWithID новый чеклист в БД по данным
// или возвращает ошибки ...
func (s PostgresChecklistStorage) ReadMany(ctx context.Context, ids dto.ChecklistIDs) (*[]dto.ChecklistInfo, error) {
	funcName := "PostgresChecklistStorage.ReadMany"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	query, args, err := sq.
		Select(allChecklistFields...).
		From("public.checklist").
		LeftJoin("public.checklist_item ON public.checklist.id = public.checklist_item.id_checklist").
		Where(sq.Eq{"public.checklist.id": ids.Values}).
		GroupBy("public.checklist.id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query with error", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		logger.DebugFmt(err.Error(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotCollectRows
	}
	defer rows.Close()
	logger.Debug("Got checklist rows", funcName, nodeName)

	checklists := []dto.ChecklistInfo{}
	for rows.Next() {
		var checklist dto.ChecklistInfo

		err = rows.Scan(
			&checklist.ID,
			&checklist.TaskID,
			&checklist.Name,
			&checklist.ListPosition,
			(*pq.StringArray)(&checklist.Items),
		)
		if err != nil {
			logger.DebugFmt(err.Error(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotGetChecklist
		}
		checklists = append(checklists, checklist)
	}
	logger.DebugFmt("Got checklists from DB", funcName, nodeName)

	return &checklists, nil
}

// Update
// обновляет чеклист в БД
// или возвращает ошибки ...
func (s PostgresChecklistStorage) Update(ctx context.Context, info dto.UpdatedChecklistInfo) error {
	sql, args, err := sq.
		Update("public.checklist").
		Set("name", info.Name).
		Set("list_position", info.ListPosition).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built Checklist query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		log.Println(err)
		return apperrors.ErrChecklistNotUpdated
	}
	log.Println("Checklist updated")

	return nil
}

// Delete
// удаляет чеклист в БД по id
// или возвращает ошибки ...
func (s PostgresChecklistStorage) Delete(ctx context.Context, id dto.ChecklistID) error {
	funcName := "PostgresChecklistStorage.Delete"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sql, args, err := sq.
		Delete("public.checklist").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		logger.DebugFmt("Failed to delete checklist with error "+err.Error(), funcName, nodeName)
		return apperrors.ErrChecklistNotDeleted
	}
	logger.DebugFmt("Checklist deleted", funcName, nodeName)

	return nil
}
