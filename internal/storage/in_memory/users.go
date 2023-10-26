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
	userData map[string]*entities.User
	mu       *sync.RWMutex
}

func NewLocalUserStorage() *LocalUserStorage {
	return &LocalUserStorage{
		userData: map[string]*entities.User{
			"test@email.com": {
				ID:           1,
				Email:        "test@email.com",
				PasswordHash: "c25596cb96ba3ff939034e762ba447b1493faa2c389b7ccfbd8d532f1c60b9e4",
				Name:         "Никита",
				Surname:      "Архаров",
				ThumbnailURL: "https://sun1-27.userapi.com/s/v1/ig1/cAIfmwiDayww2WxVGPnIr5sHTSgXaf_567nuovSw_X4Cy9XAKrSVsAT2yAmJcJXDPkVOsXPW.jpg?size=50x50&quality=96&crop=351,248,540,540&ava=1",
			},
			"email@example.com": {
				ID:           2,
				Email:        "email@example.com",
				PasswordHash: "eabf42bd681a8eb6fa35c6a3b83d746b9d500587e9c93bd0067c529fa58d6970",
				Name:         "Даниил",
				Surname:      "Капитанов",
				ThumbnailURL: "https://sun1-47.userapi.com/s/v1/ig2/aby-Y8KQ-yfQPLdvO-gq-ZenU63Iiw3ULbNlimdfaqLauSOj1cJ2jLxfBDtBMLpBW5T0UhaLFpyLVxAoYuVZiPB8.jpg?size=50x50&quality=95&crop=0,0,400,400&ava=1",
			},
			"newchallenger@email.com": {
				ID:           3,
				Email:        "newchallenger@email.com",
				PasswordHash: "b677ce792ef38ee2b79422101cd5f05a332daa19d8379ae450841158b6e5a64e",
				Name:         "Major",
				Surname:      "Guile",
				ThumbnailURL: "https://sun1-56.userapi.com/s/v1/ig2/tZBD5-zfhGO0BMktMHmj82YYqRowSPJbj7ZZE2lQ33DO1WzXB7z3fISjgAWaccX0marlKBf6tV_x0ScnO7CdM_ay.jpg?size=50x0&quality=96&crop=192,40,539,539&ava=1",
			},
			"ghostinthem@chi.ne": {
				ID:           4,
				Email:        "ghostinthem@chi.ne",
				PasswordHash: "c4de98763483d00c847e50e1c70a5c35a4fbf35a8e8c0306de609a9429124884",
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

// GetUser
// находит пользователя в БД по почте
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *LocalUserStorage) GetUserByLogin(ctx context.Context, login string) (*entities.User, error) {
	s.mu.RLock()
	user, ok := s.userData[login]
	s.mu.RUnlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	return user, nil
}

// GetUserByID
// находит пользователя в БД по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (s *LocalUserStorage) GetUserByID(ctx context.Context, uid uint64) (*entities.User, error) {
	defer s.mu.RUnlock()

	s.mu.RLock()
	for _, user := range s.userData {
		if user.ID == uid {
			return user, nil
		}
	}

	return nil, apperrors.ErrUserNotFound
}

// CreateUser
// создает нового пользователя в БД по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
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
		ThumbnailURL: "avatar.jpg",
	}

	s.mu.Lock()
	s.userData[signup.Email] = &newUser
	s.mu.Unlock()

	return &newUser, nil
}

// UpdateUser
// обновляет пользователя в БД
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (s *LocalUserStorage) UpdateUser(ctx context.Context, updatedInfo dto.UpdatedUserInfo) (*entities.User, error) {
	s.mu.RLock()
	oldUser, ok := s.userData[updatedInfo.Email]
	s.mu.RUnlock()

	if !ok {
		return nil, apperrors.ErrUserNotFound
	}

	updatedUser := entities.User{
		ID:           oldUser.ID,
		Email:        updatedInfo.Email,
		PasswordHash: oldUser.PasswordHash,
		Name:         updatedInfo.Name,
		Surname:      updatedInfo.Surname,
	}

	s.mu.Lock()
	s.userData[updatedInfo.Email] = &updatedUser
	s.mu.Unlock()

	return &updatedUser, nil
}

// DeleteUser
// удаляет данного пользователя в БД по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (s *LocalUserStorage) DeleteUser(ctx context.Context, id uint64) error {
	var (
		userEmail string
		ok        bool = false
	)
	s.mu.RLock()
	for k := range s.userData {
		if s.userData[k].ID == id {
			userEmail = k
			ok = true
			break
		}
	}
	s.mu.RUnlock()

	if !ok {
		return apperrors.ErrUserNotFound
	}

	s.mu.Lock()
	s.userData[userEmail] = nil
	s.mu.Unlock()

	return nil
}
