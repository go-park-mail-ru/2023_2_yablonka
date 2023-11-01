package postgresql

import (
	"context"
	"log"
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

// GetUserOwnedBoards
// находит все доски, созданные пользователем
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *PostgreSQLBoardStorage) GetUserOwnedBoards(ctx context.Context, userInfo dto.VerifiedAuthInfo) (*[]entities.Board, error) {
	sql, args, err := sq.
		Select(allBoardFields...).
		From("public.Board").
		Join("public.Board_User").
		Where(sq.Eq{"public.Board_User.id_user": userInfo.UserID}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	rows, err := s.db.Query(ctx, sql, args...)

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	defer rows.Close()

	var boards []entities.Board
	for rows.Next() {
		var board entities.Board
		err := rows.Scan(&board.ID, &board.Name, &board.Owner, &board.ThumbnailURL, &board.Guests)
		if err != nil {
			log.Fatal(err)
		}
		boards = append(boards, board)
	}
	if err := rows.Err(); err != nil {
		return nil, apperrors.ErrBoardNotFound
	}

	return &boards, nil
}

// GetUserGuestBoards
// находит все доски, в которых участвует пользователь
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *PostgreSQLBoardStorage) GetUserGuestBoards(ctx context.Context, userInfo dto.VerifiedAuthInfo) (*[]entities.Board, error) {
	return nil, nil
}

func (s *PostgreSQLBoardStorage) GetHighestID() uint64 {
	return 0
}

func (s *PostgreSQLBoardStorage) GetById(ctx context.Context, id uint64) (*entities.Board, error) {
	// TODO Implement error
	// s.mu.RLock()
	// userBoards, ok := s.boardData[board.OwnerEmail]
	// s.mu.RUnlock()

	// if !ok {
	// 	return nil, apperrors.ErrUserNotFound
	// }

	// for i, b := range userBoards {
	// 	if b.ID == board.ID {
	// 		return userBoards[i], nil
	// 	}
	// }
	return nil, nil
}

func (s *PostgreSQLBoardStorage) Create(ctx context.Context, info dto.NewBoardInfo) (*entities.Board, error) {
	query1, args, err := sq.
		Insert("public.Board").
		Columns("id_workspace", "name", "description", "thumbnail_url").
		Values(info.WorkspaceID, info.Name, info.Description, info.ThumbnailURL).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	var boardID int

	row := tx.QueryRow(ctx, query1, args...)
	if row.Scan(&boardID) != nil {
		tx.Rollback(ctx)
		return nil, apperrors.ErrUserNotCreated
	}

	query2, args, err := sq.
		Insert("public.Board_User").
		Columns("id_board", "id_user").
		Values(boardID, info.OwnerID).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldntBuildQuery
	}

	_, err = tx.Exec(ctx, query2, args...)

	if err != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrUserNotCreated
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrUserNotCreated
	}

	return &entities.Board{}, nil
}

func (s *PostgreSQLBoardStorage) Update(ctx context.Context, id uint64) (*entities.Board, error) {
	return nil, nil
}

func (s *PostgreSQLBoardStorage) Delete(ctx context.Context, id uint64) error {
	// TODO Implement later
	// s.mu.RLock()
	// userBoards, ok := s.boardData[board.OwnerEmail]
	// if !ok {
	// 	return apperrors.ErrUserNotFound
	// }
	// s.mu.RUnlock()

	// boardIndex := -1
	// for i, b := range userBoards {
	// 	if b.ID == board.ID {
	// 		boardIndex = i
	// 		break
	// 	}
	// }
	// if boardIndex == -1 {
	// 	return apperrors.ErrBoardNotFound
	// }
	// userBoards[boardIndex] = nil

	// s.mu.Lock()
	// s.boardData[board.OwnerEmail] = userBoards
	// s.mu.Unlock()
	return nil
}
