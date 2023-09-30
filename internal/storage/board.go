package storage

import (
	// apperrors "server/internal/apperrors"
	apperrors "server/internal/apperrors"
	"server/internal/pkg/datatypes"
	"server/internal/storage/in_memory"
)

type IBoardStorage interface {
	GetBoard(login datatypes.LoginInfo) (*datatypes.Board, error)
	CreateBoard(signup datatypes.SignupInfo) (*datatypes.Board, error)
	GetUserBoards(user datatypes.LoginInfo) (*[]datatypes.Board, error)
}

type LocalBoardStorage struct {
	Storage *in_memory.LocalBoardStorage
}

func NewLocalBoardStorage() IBoardStorage {
	storage := in_memory.NewBoardStorage()
	return &LocalBoardStorage{
		Storage: storage,
	}
}

func (s *LocalBoardStorage) GetHighestID() uint64 {
	if len(s.Storage.BoardDataByUser) == 0 {
		return 0
	}
	var highest uint64 = 0
	userEmails := make([]string, 0, len(s.Storage.BoardDataByUser))
	for _, k := range userEmails {
		for _, Board := range s.Storage.BoardDataByUser[k] {
			if Board.ID > highest {
				highest = Board.ID
			}
		}
	}

	return highest
}

func (s *LocalBoardStorage) GetUserBoards(user datatypes.LoginInfo) (*[]datatypes.Board, error) {
	s.Storage.Mu.Lock()
	boards, ok := s.Storage.BoardDataByUser[user.Email]
	s.Storage.Mu.Unlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	return &boards, nil
}

func (s *LocalBoardStorage) GetBoard(login datatypes.LoginInfo) (*datatypes.Board, error) {
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

func (s *LocalBoardStorage) CreateBoard(signup datatypes.SignupInfo) (*datatypes.Board, error) {
	// TODO Создание борды

	// s.Storage.Mu.Lock()
	// _, ok := s.Storage.BoardData[signup.Email]
	// s.Storage.Mu.Unlock()

	// if ok {
	// 	return nil, apperrors.ErrBoardExists
	// }

	// s.Storage.Mu.Lock()
	// newID := s.GetHighestID() + 1
	// newBoard := datatypes.Board{
	// 	ID:           newID,
	// 	Email:        signup.Email,
	// 	PasswordHash: signup.PasswordHash,
	// }

	// s.Storage.BoardData[signup.Email] = newBoard
	// s.Storage.Mu.Unlock()

	// return &newBoard, nil
	return nil, nil
}