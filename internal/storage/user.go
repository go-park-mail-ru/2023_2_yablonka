package storage

import (
	apperrors "server/internal/apperrors"
	"server/internal/pkg/datatypes"
	"server/internal/storage/in_memory"
)

type IUserStorage interface {
	GetUser(login datatypes.LoginInfo) (*datatypes.User, error)
	CreateUser(signup datatypes.SignupInfo) (*datatypes.User, error)
}

type LocalUserStorage struct {
	Storage *in_memory.LocalUserStorage
}

func NewLocalUserStorage() IUserStorage {
	storage := in_memory.NewUserStorage()
	return &LocalUserStorage{
		Storage: storage,
	}
}

func (s *LocalUserStorage) GetHighestID() uint64 {
	if len(s.Storage.UserData) == 0 {
		return 0
	}
	var highest uint64 = 0
	for _, user := range s.Storage.UserData {
		if user.ID > highest {
			highest = user.ID
		}
	}
	return highest
}

func (s *LocalUserStorage) GetUser(login datatypes.LoginInfo) (*datatypes.User, error) {
	s.Storage.Mu.Lock()
	user, ok := s.Storage.UserData[login.Email]
	s.Storage.Mu.Unlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

func (s *LocalUserStorage) CreateUser(signup datatypes.SignupInfo) (*datatypes.User, error) {
	s.Storage.Mu.Lock()
	_, ok := s.Storage.UserData[signup.Email]
	s.Storage.Mu.Unlock()

	if ok {
		return nil, apperrors.ErrUserExists
	}

	s.Storage.Mu.Lock()
	newID := s.GetHighestID() + 1
	newUser := datatypes.User{
		ID:           newID,
		Email:        signup.Email,
		PasswordHash: signup.PasswordHash,
	}

	s.Storage.UserData[signup.Email] = newUser
	s.Storage.Mu.Unlock()

	return &newUser, nil
}
