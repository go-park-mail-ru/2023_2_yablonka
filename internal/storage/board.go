package storage

import (
	// apperrors "server/internal/apperrors"

	"server/internal/pkg/datatypes"
)

type IBoardStorage interface {
	GetBoard(login datatypes.LoginInfo) (*datatypes.Board, error)
	CreateBoard(signup datatypes.SignupInfo) (*datatypes.Board, error)
	GetUserBoards(user datatypes.User) (*[]datatypes.Board, error)
}
