package in_memory

import (
	"server/internal/pkg/datatypes"
	"sync"
)

// LocalUserStorage
// Локальное хранение данных
type LocalUserStorage struct {
	UserData map[string]datatypes.User
	Mu       *sync.Mutex
}

// NewLocalUserStorage
// Возвращает локальное хранилище данных с тестовыми данными
func NewUserStorage() *LocalUserStorage {
	return &LocalUserStorage{
		UserData: map[string]datatypes.User{
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
		Mu: &sync.Mutex{},
	}
}
