package in_memory

import (
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"sync"
)

// LocalUserStorage
// Локальное хранение данных
type LocalBoardStorage struct {
	boardData map[uint64]entities.Board
	mu        *sync.RWMutex
}

func NewBoardStorage() *LocalBoardStorage {
	return &LocalBoardStorage{
		boardData: map[uint64]entities.Board{
			1: entities.Board{
				ID:           1,
				Name:         "Проект 1",
				OwnerID:      1,
				ThumbnailURL: "https://media.moddb.com/images/downloads/1/203/202069/missing_textures.png",
			},
			2: entities.Board{
				ID:           2,
				Name:         "Разработка Ведра 2",
				OwnerID:      1,
				ThumbnailURL: "https://nicollelamerichs.files.wordpress.com/2022/05/2022043021483800-9e19570e6059798a45aec175873b4ac1.jpg?w=640",
			},
			3: entities.Board{
				ID:           3,
				Name:         "лучшая вещь",
				OwnerID:      1,
				ThumbnailURL: "https://media.istockphoto.com/id/868643608/photo/thumbs-up-emoji-isolated-on-white-background-emoticon-giving-likes-3d-rendering.jpg?s=612x612&w=0&k=20&c=ulAeL-xm8S-g5VU_28CUlOqzqT-ooGTKuXYe097XEL8=",
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
	userEmails := make([]string, 0, len(s.boardData))
	for _, k := range userEmails {
		for _, Board := range s.boardData[k] {
			if Board.ID > highest {
				highest = Board.ID
			}
		}
	}

	return highest
}

func (s *LocalBoardStorage) GetUserBoards(user entities.User) (*[]*entities.Board, error) {
	s.mu.RLock()
	boards, ok := s.boardData[user.Email]
	s.mu.Unlock()

	if !ok {
		return nil, apperrors.ErrBoardNotFound
	}

	return &boards, nil
}

func (s *LocalBoardStorage) GetBoard(board dto.IndividualBoardInfo) (*entities.Board, error) {
	s.mu.RLock()
	userBoards, ok := s.boardData[board.OwnerEmail]
	s.mu.RUnlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	for i, b := range userBoards {
		if b.ID == board.ID {
			return userBoards[i], nil
		}
	}
	return nil, apperrors.ErrBoardNotFound
}

func (s *LocalBoardStorage) CreateBoard(board dto.NewBoardInfo) (*entities.Board, error) {
	// TODO Нужна проверка по количеству доступных пользователю досок, это наверное поле в User

	s.mu.Lock()
	newBoard := entities.Board{
		ID:           s.GetHighestID() + 1,
		Name:         board.Name,
		OwnerID:      board.OwnerID,
		ThumbnailURL: "",
	}

	s.boardData[board.OwnerEmail] = append(s.boardData[board.OwnerEmail], &newBoard)
	s.mu.Unlock()

	return &newBoard, nil
}

func (s *LocalBoardStorage) DeleteBoard(board dto.IndividualBoardInfo) error {
	s.mu.RLock()
	userBoards, ok := s.boardData[board.OwnerEmail]
	if !ok {
		return apperrors.ErrUserNotFound
	}
	s.mu.RUnlock()

	boardIndex := -1
	for i, b := range userBoards {
		if b.ID == board.ID {
			boardIndex = i
			break
		}
	}
	if boardIndex == -1 {
		return apperrors.ErrBoardNotFound
	}
	userBoards[boardIndex] = nil

	s.mu.Lock()
	s.boardData[board.OwnerEmail] = userBoards
	s.mu.Unlock()
	return nil
}
