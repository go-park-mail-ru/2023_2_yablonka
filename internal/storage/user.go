package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserStorage interface {
	GetUser(context.Context, dto.LoginInfo) (*entities.User, error)
	GetUserByID(context.Context, uint64) (*entities.User, error)
	CreateUser(context.Context, dto.SignupInfo) (*entities.User, error)
	UpdateUser(context.Context, dto.UpdatedUserInfo) (*entities.User, error)
	DeleteUser(context.Context, uint64) error
}
