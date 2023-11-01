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

func (s *PostgreSQLBoardStorage) GetBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
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

func (s *PostgreSQLBoardStorage) CreateBoard(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	// TODO Нужна проверка по количеству доступных пользователю досок, это наверное поле в User

	// s.mu.Lock()
	// newBoard := entities.Board{
	// 	ID:           s.GetHighestID() + 1,
	// 	Name:         board.Name,
	// 	OwnerID:      board.OwnerID,
	// 	ThumbnailURL: "",
	// }

	// s.boardData[board.OwnerEmail] = append(s.boardData[board.OwnerEmail], &newBoard)
	// s.mu.Unlock()

	// return &newBoard, nil
	return nil, nil
}

func (s *PostgreSQLBoardStorage) DeleteBoard(ctx context.Context, board dto.IndividualBoardInfo) error {
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