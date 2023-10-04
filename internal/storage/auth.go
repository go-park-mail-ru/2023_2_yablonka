package storage

import (
	"context"
	"server/internal/pkg/entities"
)

type IAuthStorage interface {
	CreateSession(context.Context, *entities.Session, uint) (string, error)
	GetSession(context.Context, string) (*entities.Session, error)
	DeleteSession(context.Context, string) error
}
