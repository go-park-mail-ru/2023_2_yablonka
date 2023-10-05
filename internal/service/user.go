package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserAuthService interface {
	GetUser(context.Context, dto.LoginInfo) (*entities.User, error)
	CreateUser(context.Context, dto.SignupInfo) (*entities.User, error)
	GetUserByID(context.Context, uint64) (*entities.User, error)
	DeleteUser(context.Context, uint64) error
	// TODO Implement
	// GetUserFromAuthInfo(context.Context, dto.VerifiedAuthInfo) (*entities.User, error)
}

type IUserService interface {
	UpdateUser(context.Context, dto.UpdatedUserInfo) (*entities.User, error)
	IUserAuthService
}
