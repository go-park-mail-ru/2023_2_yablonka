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
				ThumbnailURL: "https://sun1-27.userapi.com/s/v1/ig1/cAIfmwiDayww2WxVGPnIr5sHTSgXaf_567nuovSw_X4Cy9XAKrSVsAT2yAmJcJXDPkVOsXPW.jpg?size=50x50&quality=96&crop=351,248,540,540&ava=1",
			},
			"email@example.com": {
				ID:           2,
				Email:        "email@example.com",
				PasswordHash: "177e4fd1a8b22992e78145c3ba9c8781124e5c166d03b9c302cf8e100d77ad22",
				Name:         "Даниил",
				Surname:      "Капитанов",
				ThumbnailURL: "https://sun1-47.userapi.com/s/v1/ig2/aby-Y8KQ-yfQPLdvO-gq-ZenU63Iiw3ULbNlimdfaqLauSOj1cJ2jLxfBDtBMLpBW5T0UhaLFpyLVxAoYuVZiPB8.jpg?size=50x50&quality=95&crop=0,0,400,400&ava=1",
			},
			"newchallenger@email.com": {
				ID:           3,
				Email:        "newchallenger@email.com",
				PasswordHash: "4aeb64424005ea74206c1a8e2c054ae1f34fec181bca5c8899152d8791c5c27f",
				Name:         "Major",
				Surname:      "Guile",
				ThumbnailURL: "https://sun1-56.userapi.com/s/v1/ig2/tZBD5-zfhGO0BMktMHmj82YYqRowSPJbj7ZZE2lQ33DO1WzXB7z3fISjgAWaccX0marlKBf6tV_x0ScnO7CdM_ay.jpg?size=50x0&quality=96&crop=192,40,539,539&ava=1",
			},
			"ghostinthem@chi.ne": {
				ID:           4,
				Email:        "ghostinthem@chi.ne",
				PasswordHash: "cae8f1b32ad474ef1d27bbe387be25e6e6ee848a1726208cc970ef8a08ed3b08",
				Name:         "Lain",
				Surname:      "Iwakura",
				ThumbnailURL: "https://sun1-57.userapi.com/s/v1/ig2/YWU6FC5ElvSSShdkrSEqWC6InGJfDQv2yuWQREEFBpHt9Nsvs_o9qc3rR2yAyVdTGy7MLMaaGbOj1DhY33JWS-b7.jpg?size=50x50&quality=95&crop=439,244,605,605&ava=1",
			},
		},
		mu: &sync.RWMutex{},
	}
}

func (s *LocalUserStorage) GetHighestID() uint64 {
	var highest uint64 = 0

	s.mu.RLock()
	if len(s.userData) != 0 {
		for _, user := range s.userData {
			if user.ID > highest {
				highest = user.ID
			}
		}
	}
	s.mu.RUnlock()

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

func (s *LocalUserStorage) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	defer s.mu.RUnlock()

	s.mu.RLock()
	for _, user := range s.userData {
		if user.ID == uid {
			return &user, nil
		}
	}

	return nil, apperrors.ErrUserNotFound
}

func (s *LocalUserStorage) CreateUser(ctx context.Context, signup dto.SignupInfo) (*entities.User, error) {
	s.mu.RLock()
	_, ok := s.userData[signup.Email]
	s.mu.RUnlock()

	if ok {
		return nil, apperrors.ErrUserAlreadyExists
	}

	newID := s.GetHighestID() + 1
	newUser := entities.User{
		ID:           newID,
		Email:        signup.Email,
		PasswordHash: signup.PasswordHash,
	}

	s.mu.Lock()
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

	updatedUser := entities.User{
		ID:           oldUser.ID,
		Email:        updatedInfo.Email,
		PasswordHash: oldUser.PasswordHash,
		Name:         updatedInfo.Name,
		Surname:      updatedInfo.Surname,
	}

	s.mu.Lock()
	s.userData[updatedInfo.Email] = updatedUser
	s.mu.Unlock()

	return &updatedUser, nil
}
