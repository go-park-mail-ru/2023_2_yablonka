package postgresql

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

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
	boards := []entities.Board{}
	err := s.db.QueryRow(context.Background(),
		"SELECT id, email, password_hash, name, surname, avatar_url, description FROM user WEHRE id = :id",
	).Scan(&boards)
	return &boards, err
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
