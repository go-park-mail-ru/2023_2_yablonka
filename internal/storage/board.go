package storage

import (
	// apperrors "server/internal/apperrors"

	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	GetBoard(dto.LoginInfo) (*entities.Board, error)
	CreateBoard(dto.SignupInfo) (*entities.Board, error)
	GetUserBoards(entities.User) (*[]entities.Board, error)
}
