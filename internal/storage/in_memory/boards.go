package in_memory

import (
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"sync"
)

// LocalUserStorage
// Локальное хранение данных
type LocalBoardStorage struct {
	boardData map[uint64]entities.Board
	mu        *sync.RWMutex
}

// NewLocalBoardStorage
// Возвращает локальное хранилище данных с тестовыми данными
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

func NewLocalBoardStorage() storage.IBoardStorage {
	storage := NewBoardStorage()
	return storage
}

func (s *LocalBoardStorage) GetHighestID() uint64 {
	if len(s.boardDataByUser) == 0 {
		return 0
	}
	var highest uint64 = 0
	userEmails := make([]string, 0, len(s.boardDataByUser))
	for _, k := range userEmails {
		for _, Board := range s.boardDataByUser[k] {
			if Board.ID > highest {
				highest = Board.ID
			}
		}
	}

	return highest
}

func (s *LocalBoardStorage) GetUserBoards(user entities.User) (*[]entities.Board, error) {
	s.mu.RLock()
	boards, ok := s.boardDataByUser[user.Email]
	s.mu.Unlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	return &boards, nil
}

func (s *LocalBoardStorage) GetBoard(login dto.LoginInfo) (*entities.Board, error) {
	// TODO Получение борды

	// s.Storage.Mu.Lock()
	// Board, ok := s.Storage.BoardData[login.Email]
	// s.Storage.Mu.Unlock()

	// if !ok {
	// 	return nil, apperrors.ErrBoardNotFound
	// }

	// return &Board, nilr
	return nil, nil
}

func (s *LocalBoardStorage) CreateBoard(signup dto.SignupInfo) (*entities.Board, error) {
	// TODO Создание борды

	// s.Storage.Mu.Lock()
	// _, ok := s.Storage.BoardData[signup.Email]
	// s.Storage.Mu.Unlock()

	// if ok {
	// 	return nil, apperrors.ErrBoardExists
	// }

	// s.Storage.Mu.Lock()
	// newID := s.GetHighestID() + 1
	// newBoard := entities.Board{
	// 	ID:           newID,
	// 	Email:        signup.Email,
	// 	PasswordHash: signup.PasswordHash,
	// }

	// s.Storage.BoardData[signup.Email] = newBoard
	// s.Storage.Mu.Unlock()

	// return &newBoard, nil
	return nil, nil
}
