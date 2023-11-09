package postgresql

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgreSQLBoardStorage
// Хранилище досок в PostgreSQL
type PostgreSQLBoardStorage struct {
	db *pgxpool.Pool
}

func NewBoardStorage(db *pgxpool.Pool) *PostgreSQLBoardStorage {
	return &PostgreSQLBoardStorage{
		db: db,
	}
}

// TODO Ограничить количество всего за раз
// GetById
// находит доску и связанные с ней списки и задания по id
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetById(ctx context.Context, id dto.BoardID) (*entities.Board, error) {
	sql, args, err := sq.Select("board.*", "list.*", "task.*").
		From("board").
		Join("list ON board.id = list.board_id").
		Join("task ON list.id = task.list_id").
		Where(sq.Eq{"board.id": id}).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	defer rows.Close()

	var board entities.Board
	for rows.Next() {
		var list entities.List

		err = rows.Scan(
			&board.ID,
			&board.Name,
			&board.Description,
			&list.ID,
			&list.Name,
			&list.Description,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}

		for rows.Next() {
			var task entities.Task

			err = rows.Scan(
				&board.ID,
				&board.Name,
				&board.Description,
				&list.ID,
				&list.Name,
				&list.Description,
				&task.ID,
				&task.Name,
				&task.Description,
			)
			if err != nil {
				return nil, apperrors.ErrCouldNotGetBoard
			}
			list.Tasks = append(list.Tasks, task)
		}

		board.Lists = append(board.Lists, list)
	}

	return &board, nil
}

// GetUsers
// находит пользователей, у которых есть доступ к доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetUsers(ctx context.Context, id dto.BoardID) ([]dto.UserPublicInfo, error) {
	sql, args, err := sq.Select("user.id").
		From("user").
		Join("board_user ON board_user.id_user = user.id").
		Where(sq.Eq{"board_user.id_board": id}).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()

	var users []dto.UserPublicInfo
	for rows.Next() {
		var user dto.UserPublicInfo

		err = rows.Scan(&user)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *PostgreSQLBoardStorage) Create(ctx context.Context, info dto.NewBoardInfo) (*entities.Board, error) {
	query1, args, err := sq.
		Insert("board").
		Columns("id_workspace", "name", "description").
		Values(info.WorkspaceID, info.Name, info.Description).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	var boardID int

	row := tx.QueryRow(ctx, query1, args...)
	if row.Scan(&boardID) != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}

	query2, args, err := sq.
		Insert("board_user").
		Columns("id_board", "id_user").
		Values(boardID, info.OwnerID).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	_, err = tx.Exec(ctx, query2, args...)

	if err != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}

	return &entities.Board{}, nil
}

func (s *PostgreSQLBoardStorage) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	sql, args, err := sq.
		Update("board").
		Set("name", info.Name).
		Set("description", info.Description).
		Where(sq.Eq{"board.id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotUpdated
	}

	return nil
}

func (s *PostgreSQLBoardStorage) UpdateThumbnailUrl(ctx context.Context, info dto.ImageUrlInfo) error {
	sql, args, err := sq.
		Update("board").
		Set("thumbnail_url", info.Url).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if query.Scan() != nil {
		return apperrors.ErrBoardNotUpdated
	}

	return nil
}

func (s *PostgreSQLBoardStorage) Delete(ctx context.Context, id dto.BoardID) error {
	sql, args, err := sq.
		Delete("board").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotDeleted
	}

	return nil
}
