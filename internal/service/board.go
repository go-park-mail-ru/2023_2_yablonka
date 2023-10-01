package service

import (
	"context"
	"server/internal/pkg/dto"
)

type IBoardService interface {
	// TODO implement later
	// GetBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	GetUserOwnedBoards(context.Context, dto.VerifiedAuthInfo) (*[]dto.UserOwnedBoardInfo, error)
	GetUserGuestBoards(context.Context, dto.VerifiedAuthInfo) (*[]dto.UserGuestBoardInfo, error)
}
