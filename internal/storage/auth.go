package storage

import (
	"context"
	"server/internal/pkg/entities"
)

type IAuthStorage interface {
	CreateSession(context.Context, *entities.Session) (string, error)
	GetSession(context.Context, string) (*entities.Session, error)
	DeleteSession(context.Context, *entities.User) error
}
