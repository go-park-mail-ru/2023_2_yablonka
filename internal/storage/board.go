package storage

import (
	// apperrors "server/internal/apperrors"

	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	// GetUserOwnedBoards
	// находит все доски, созданные пользователем
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserOwnedBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)

	// TODO Implement
	// GetBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// GetUserBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// DeleteBoard(context.Context, dto.IndividualBoardInfo) error
}
