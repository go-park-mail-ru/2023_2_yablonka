package in_memory

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"sync"
)

// LocalUserStorage
// Локальное хранилище данных
type LocalUserStorage struct {
	userData map[string]entities.User
	mu       *sync.RWMutex
}

func NewUserStorage() *LocalUserStorage {
	return &LocalUserStorage{
		userData: map[string]entities.User{
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
			"newchallenger@email.com": {
				ID:           3,
				Email:        "newchallenger@email.com",
				PasswordHash: "4aeb64424005ea74206c1a8e2c054ae1f34fec181bca5c8899152d8791c5c27f",
				Name:         "Major",
				Surname:      "Guile",
			},
			"ghostinthem@chi.ne": {
				ID:           4,
				Email:        "ghostinthem@chi.ne",
				PasswordHash: "cae8f1b32ad474ef1d27bbe387be25e6e6ee848a1726208cc970ef8a08ed3b08",
				Name:         "Lain",
				Surname:      "Iwakura",
			},
		},
		mu: &sync.RWMutex{},
	}
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

func (s *LocalUserStorage) GetUser(ctx context.Context, login dto.LoginInfo) (*entities.User, error) {
	s.mu.RLock()
	user, ok := s.userData[login.Email]
	s.mu.RUnlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	return &user, nil
}

func (s *LocalUserStorage) CreateUser(ctx context.Context, signup dto.SignupInfo) (*entities.User, error) {
	s.mu.RLock()
	_, ok := s.userData[signup.Email]
	s.mu.RUnlock()

	if ok {
		return nil, apperrors.ErrUserAlreadyExists
	}

	s.mu.Lock()
	newID := s.GetHighestID() + 1
	newUser := entities.User{
		ID:           newID,
		Email:        signup.Email,
		PasswordHash: signup.PasswordHash,
	}

	s.userData[signup.Email] = newUser
	s.mu.Unlock()

	return &newUser, nil
}

func (s *LocalUserStorage) UpdateUser(ctx context.Context, updatedInfo dto.UpdatedUserInfo) (*entities.User, error) {
	s.mu.RLock()
	oldUser, ok := s.userData[updatedInfo.Email]
	s.mu.RUnlock()

	if ok {
		return nil, apperrors.ErrUserAlreadyExists
	}

	s.mu.Lock()
	updatedUser := entities.User{
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
