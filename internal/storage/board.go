package storage

import (
	// apperrors "server/internal/apperrors"

	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	GetBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	GetUserBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	GetUserOwnedBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	GetUserGuestBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	DeleteBoard(context.Context, dto.IndividualBoardInfo) error
}
