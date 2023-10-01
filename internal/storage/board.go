package storage

import (
	// apperrors "server/internal/apperrors"

	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	GetBoard(login dto.LoginInfo) (*entities.Board, error)
	CreateBoard(signup dto.SignupInfo) (*entities.Board, error)
	GetUserBoards(user entities.User) (*[]entities.Board, error)
}
