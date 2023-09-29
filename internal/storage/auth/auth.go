package storage

import (
	"server/internal/storage/in_memory"
)

type IAuthStorage interface {
	VerifyLogin() error
}

func NewLocalAuthStorage() IAuthStorage {
	storage := in_memory.NewLocalTestStorage()
	return storage
}

// func NewPostgresAuthStorage() *AuthStorage {
// 	return &AuthStorage{}
// }
