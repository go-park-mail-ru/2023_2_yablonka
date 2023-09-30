package service

import (
	"context"
	"server/internal/pkg/datatypes"
)

type IUserAuthService interface {
	GetUser(ctx context.Context, info datatypes.LoginInfo) (*datatypes.User, error)
	CreateUser(ctx context.Context, info datatypes.SignupInfo) (*datatypes.User, error)
}

type IUserService interface {
	UpdateUser()
	IUserAuthService
}
