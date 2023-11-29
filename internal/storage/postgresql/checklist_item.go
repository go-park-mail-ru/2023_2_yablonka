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
)

// PostgresChecklistItemStorage
// Хранилище данных в PostgreSQL
type PostgresChecklistItemStorage struct {
	db *sql.DB
}

// NewChecklistItemStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewChecklistItemStorage(db *sql.DB) *PostgresChecklistItemStorage {
	return &PostgresChecklistItemStorage{
		db: db,
	}
}

// Create
// создает новый чеклист в БД по данным
// или возвращает ошибки ...
func (s PostgresChecklistItemStorage) Create(ctx context.Context, info dto.NewChecklistItemInfo) (*dto.ChecklistItemInfo, error) {
	sql, args, err := sq.
		Insert("public.checklist_item").
		Columns("name", "list_position", "id_checklist", "done").
		Values(info.Name, info.ListPosition, info.ChecklistID, info.Done).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built ChecklistItem query\n\t", sql, "\nwith args\n\t", args)

	checklistItem := dto.ChecklistItemInfo{
		Name:         info.Name,
		ChecklistID:  info.ChecklistID,
		ListPosition: info.ListPosition,
		Done:         info.Done,
	}
	query := s.db.QueryRow(sql, args...)
	if err := query.Scan(&checklistItem.ID); err != nil {
		log.Println("Storage -- Failed to create ChecklistItem")
		return nil, apperrors.ErrChecklistItemNotCreated
	}

	log.Println("Storage -- ChecklistItem created")

	return &checklistItem, nil
}

// ReadMany
// дает много элементов чеклистов
// или возвращает ошибки ...
func (s PostgresChecklistItemStorage) ReadMany(ctx context.Context, ids dto.ChecklistItemIDs) (*[]dto.ChecklistItemInfo, error) {
	funcName := "PostgresChecklistItemStorage.ReadMany"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	query, args, err := sq.
		Select("id", "name", "list_position", "id_checklist", "done").
		From("public.checklist_item").
		Where(sq.Eq{"id": ids.Values}).
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
		return nil, apperrors.ErrCouldNotCollectRows
	}
	defer rows.Close()
	logger.Debug("Got checklist item rows", funcName, nodeName)

	checklistItems := []dto.ChecklistItemInfo{}
	for rows.Next() {
		var checklistItem dto.ChecklistItemInfo

		err = rows.Scan(
			&checklistItem.ID,
			&checklistItem.Name,
			&checklistItem.ListPosition,
			&checklistItem.ChecklistID,
			&checklistItem.Done,
		)
		if err != nil {
			logger.Debug(err.Error(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotGetChecklistItem
		}
		checklistItems = append(checklistItems, checklistItem)
	}
	logger.Debug("Got checklistItems from DB", funcName, nodeName)

	return &checklistItems, nil
}

// Update
// обновляет чеклист в БД
// или возвращает ошибки ...
func (s PostgresChecklistItemStorage) Update(ctx context.Context, info dto.UpdatedChecklistItemInfo) error {
	sql, args, err := sq.
		Update("public.checklist_item").
		Set("name", info.Name).
		Set("done", info.Done).
		Set("list_position", info.ListPosition).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built ChecklistItem query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		log.Println(err)
		return apperrors.ErrChecklistItemNotUpdated
	}
	log.Println("ChecklistItem updated")

	return nil
}

// Delete
// удаляет чеклист в БД по id
// или возвращает ошибки ...
func (s PostgresChecklistItemStorage) Delete(ctx context.Context, id dto.ChecklistItemID) error {
	sql, args, err := sq.
		Delete("public.checklist_item").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return apperrors.ErrChecklistItemNotDeleted
	}

	return nil
}
