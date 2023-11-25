package postgresql

import (
	"context"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresListStorage
// Хранилище данных в PostgreSQL
type PostgresListStorage struct {
	db *pgxpool.Pool
}

// NewListStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewListStorage(db *pgxpool.Pool) *PostgresListStorage {
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
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if err := query.Scan(&list.ID); err != nil {
		log.Println("Storage -- Failed to create list")
		return nil, apperrors.ErrListNotCreated
	}

	log.Println("Storage -- List created")

	return &list, nil
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

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrListNotUpdated
	}

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

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrListNotDeleted
	}

	return nil
}
