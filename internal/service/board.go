package service

import (
	"context"
	"server/internal/pkg/dto"
)

type IBoardService interface {
	// GetUserOwnedBoards
	// находит все доски, созданные пользователем
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserOwnedBoards(context.Context, dto.VerifiedAuthInfo) ([]dto.UserOwnedBoardInfo, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(context.Context, dto.VerifiedAuthInfo) ([]dto.UserGuestBoardInfo, error)

	// TODO Implement
	// GetBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
}
