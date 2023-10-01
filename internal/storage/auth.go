package storage

import "server/internal/pkg/entities"

type IAuthStorage interface {
	CreateSession(*entities.Session) (string, error)
	GetSession(string) (*entities.Session, error)
	DeleteSession(*entities.User) error
}
