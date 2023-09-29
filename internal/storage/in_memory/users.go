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
				PasswordHash: "$2a$08$t5mReYGdL961vpUe4enrLeHKazp.JLZXIAX38t5vHfPFJZM4TnP2C",
				Name:         "Никита",
				Surname:      "Архаров",
			},
			"email@example.com": {
				ID:           2,
				Email:        "email@example.com",
				PasswordHash: "$2a$08$5vGskE/R50Ju92.4AbbZyeQiBT26Hiiq.4RqoRf5yGOrExfKDCW52",
				Name:         "Даниил",
				Surname:      "Капитанов",
			},
		},
		Mu: &sync.Mutex{},
	}
}
