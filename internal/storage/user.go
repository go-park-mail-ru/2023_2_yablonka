package storage

import (
	"server/internal/pkg/datatypes"
)

type IUserStorage interface {
	GetUser(login datatypes.LoginInfo) (*datatypes.User, error)
	CreateUser(signup datatypes.SignupInfo) (*datatypes.User, error)
	UpdateUser(updatedInfo datatypes.UpdatedUserInfo) (*datatypes.User, error)
}
