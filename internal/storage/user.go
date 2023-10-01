package storage

import (
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IUserStorage interface {
	GetUser(dto.LoginInfo) (*entities.User, error)
	CreateUser(dto.SignupInfo) (*entities.User, error)
	UpdateUser(dto.UpdatedUserInfo) (*entities.User, error)
}
