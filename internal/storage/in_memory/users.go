package in_memory

import (
	"server/internal/apperrors"
	"server/internal/pkg/datatypes"
	"server/internal/storage"
	"sync"
)

// UserStore
// Локальное хранилище данных

type LocalUserStorage struct {
	userData map[string]datatypes.User
	mu       *sync.RWMutex
}

// NewLocalUserStorage
// Возвращает локальное хранилище данных с тестовыми данными
func NewUserStorage() *LocalUserStorage {
	return &LocalUserStorage{
		userData: map[string]datatypes.User{
			"test@email.com": {
				ID:           1,
				Email:        "test@email.com",
				PasswordHash: "8a65f9232aec42190593cebe45067d14ade16eaf9aaefe0c2e9ec425b5b8ca73",
				Name:         "Никита",
				Surname:      "Архаров",
			},
			"email@example.com": {
				ID:           2,
				Email:        "email@example.com",
				PasswordHash: "177e4fd1a8b22992e78145c3ba9c8781124e5c166d03b9c302cf8e100d77ad22",
				Name:         "Даниил",
				Surname:      "Капитанов",
			},
		},
		mu: &sync.RWMutex{},
	}
}

// Убрать в in_memory vvvv
// Заприватить

func NewLocalUserStorage() storage.IUserStorage {
	storage := NewUserStorage()
	return storage
}

func (s *LocalUserStorage) GetHighestID() uint64 {
	if len(s.userData) == 0 {
		return 0
	}
	var highest uint64 = 0
	for _, user := range s.userData {
		if user.ID > highest {
			highest = user.ID
		}
	}
	return highest
}

func (s *LocalUserStorage) GetUser(login datatypes.LoginInfo) (*datatypes.User, error) {
	s.mu.RLock()
	user, ok := s.userData[login.Email]
	s.mu.RUnlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

func (s *LocalUserStorage) CreateUser(signup datatypes.SignupInfo) (*datatypes.User, error) {
	s.mu.Lock()
	_, ok := s.userData[signup.Email]
	s.mu.Unlock()

	if ok {
		return nil, apperrors.ErrUserExists
	}

	s.mu.Lock()
	newID := s.GetHighestID() + 1
	newUser := datatypes.User{
		ID:           newID,
		Email:        signup.Email,
		PasswordHash: signup.PasswordHash,
	}

	s.userData[signup.Email] = newUser
	s.mu.Unlock()

	return &newUser, nil
}

func (s *LocalUserStorage) UpdateUser(updatedInfo datatypes.UpdatedUserInfo) (*datatypes.User, error) {
	s.mu.Lock()
	oldUser, ok := s.userData[updatedInfo.Email]
	s.mu.Unlock()

	if ok {
		return nil, apperrors.ErrUserExists
	}

	s.mu.Lock()
	updatedUser := datatypes.User{
		ID:           oldUser.ID,
		Email:        updatedInfo.Email,
		PasswordHash: oldUser.PasswordHash,
		Name:         updatedInfo.Name,
		Surname:      updatedInfo.Surname,
	}

	s.userData[updatedInfo.Email] = updatedUser
	s.mu.Unlock()

	return &updatedUser, nil
}
