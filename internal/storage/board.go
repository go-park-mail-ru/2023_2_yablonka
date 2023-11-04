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
<<<<<<< Updated upstream
	GetUserOwnedBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
=======
<<<<<<< Updated upstream
	GetUserOwnedBoards(ctx context.Context, info dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(ctx context.Context, info dto.VerifiedAuthInfo) (*[]entities.Board, error)
=======
<<<<<<< Updated upstream
	GetUserOwnedBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
>>>>>>> Stashed changes
	GetUserGuestBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)

	// TODO Implement
	// GetBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// GetUserBoards(context.Context, dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// DeleteBoard(context.Context, dto.IndividualBoardInfo) error
<<<<<<< Updated upstream
=======
=======
	GetUserOwnedBoards(ctx context.Context, userID uint64) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(ctx context.Context, userID uint64) (*[]entities.Board, error)
>>>>>>> Stashed changes
	// GetById
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	GetById(ctx context.Context, id uint64) (*entities.Board, error)
	// Update
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Update(ctx context.Context, id uint64) (*entities.Board, error)
	// Create
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Create(ctx context.Context, info dto.NewBoardInfo) (*entities.Board, error)
	// Delete
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Delete(ctx context.Context, id uint64) error
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
