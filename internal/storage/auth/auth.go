package storage

import (
	"server/internal/storage/in_memory"
)

type AuthStorage struct {
	storage interface{}
}

func NewLocalAuthStorage() *AuthStorage {
	return &AuthStorage{
		storage: in_memory.NewTestStorage(),
	}
}

func NewPostgresAuthStorage() *AuthStorage {
	return &AuthStorage{}
}

func (a *AuthStorage) GetSession() error {
	return nil
}
