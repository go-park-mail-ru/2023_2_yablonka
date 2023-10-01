package storage

import (
	// apperrors "server/internal/apperrors"

	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	GetBoard(dto.IndividualBoardInfo) (*entities.Board, error)
	UpdateBoard(dto.IndividualBoardInfo) (*entities.Board, error)
	CreateBoard(dto.NewBoardInfo) (*entities.Board, error)
	GetUserBoards(dto.VerifiedAuthInfo) (*[]entities.Board, error)
	GetUserOwnedBoards(dto.VerifiedAuthInfo) (*[]entities.Board, error)
	GetUserGuestBoards(dto.VerifiedAuthInfo) (*[]entities.Board, error)
	DeleteBoard(dto.IndividualBoardInfo) error
}
