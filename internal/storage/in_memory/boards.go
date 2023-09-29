package in_memory

import (
	"server/internal/pkg/datatypes"
	"sync"
)

// LocalUserStorage
// Локальное хранение данных
type LocalBoardStorage struct {
	BoardData map[string]datatypes.Board
	Mu        *sync.Mutex
}

// NewLocalBoardStorage
// Возвращает локальное хранилище данных с тестовыми данными
func NewBoardStorage() *LocalBoardStorage {
	return &LocalBoardStorage{
		BoardData: map[string]datatypes.Board{
			// TODO Заполнить данными
			// "test@email.com": {
			// 	ID:           1,
			// 	Email:        "test@email.com",
			// 	PasswordHash: "8a65f9232aec42190593cebe45067d14ade16eaf9aaefe0c2e9ec425b5b8ca73",
			// 	Name:         "Никита",
			// 	Surname:      "Архаров",
			// },
			// "email@example.com": {
			// 	ID:           2,
			// 	Email:        "email@example.com",
			// 	PasswordHash: "177e4fd1a8b22992e78145c3ba9c8781124e5c166d03b9c302cf8e100d77ad22",
			// 	Name:         "Даниил",
			// 	Surname:      "Капитанов",
			// },
		},
		Mu: &sync.Mutex{},
	}
}
