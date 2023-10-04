package in_memory

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"sync"
)

// LocalUserStorage
// Локальное хранение данных
type LocalBoardStorage struct {
	boardData []entities.Board
	mu        *sync.RWMutex
}

func NewBoardStorage() *LocalBoardStorage {
	return &LocalBoardStorage{
		boardData: []entities.Board{
			{
				ID:   1,
				Name: "Проект 1",
				Owner: dto.UserInfo{
					ID:    1,
					Email: "test@example.com",
				},
				ThumbnailURL: "https://images.freecreatives.com/wp-content/uploads/2016/04/Download-Pastel-HD-Wallpaper.jpg",
				Guests: []dto.UserInfo{
					{
						ID:    3,
						Email: "newchallenger@example.com",
					},
				},
			},
			{
				ID:   2,
				Name: "Проект 2",
				Owner: dto.UserInfo{
					ID:    1,
					Email: "test@example.com",
				},
				ThumbnailURL: "https://images.freecreatives.com/wp-content/uploads/2016/04/Pastel-HD-Widescreen-Wallpaper.jpg",
			},
			{
				ID:   3,
				Name: "Проект 3",
				Owner: dto.UserInfo{
					ID:    2,
					Email: "example@email.com",
				},
				ThumbnailURL: "https://images.freecreatives.com/wp-content/uploads/2016/04/Soft-Pastels-Wallpaper.jpg",
				Guests: []dto.UserInfo{
					{
						ID:    1,
						Email: "test@example.com",
					},
					{
						ID:    3,
						Email: "newchallenger@example.com",
					},
				},
			},
		},
		mu: &sync.RWMutex{},
	}
}

func (s *LocalBoardStorage) GetHighestID() uint64 {
	if len(s.boardData) == 0 {
		return 0
	}
	var highest uint64 = 0
	for _, board := range s.boardData {
		if board.ID > highest {
			highest = board.ID
		}
	}

	return highest
}

func (s *LocalBoardStorage) GetUserOwnedBoards(ctx context.Context, userInfo dto.VerifiedAuthInfo) (*[]entities.Board, error) {
	var boards []entities.Board
	s.mu.RLock()
	for _, board := range s.boardData {
		if board.Owner.ID == userInfo.UserID {
			boards = append(boards, board)
		}
	}
	s.mu.RUnlock()

	return &boards, nil
}

func (s *LocalBoardStorage) GetUserGuestBoards(ctx context.Context, userInfo dto.VerifiedAuthInfo) (*[]entities.Board, error) {
	var boards []entities.Board
	s.mu.RLock()
	for _, board := range s.boardData {
		for _, guest := range board.Guests {
			if guest.ID == userInfo.UserID {
				boards = append(boards, board)
			}
		}
	}
	s.mu.RUnlock()

	return &boards, nil
}

func (s *LocalBoardStorage) GetBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
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

func (s *LocalBoardStorage) CreateBoard(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
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

func (s *LocalBoardStorage) DeleteBoard(ctx context.Context, board dto.IndividualBoardInfo) error {
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
