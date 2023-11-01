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
	GetUserOwnedBoards(ctx context.Context, info dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(ctx context.Context, info dto.VerifiedAuthInfo) (*[]entities.Board, error)
	// GetById
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	GetById(ctx context.Context, id int) (*entities.Board, error)
	// Update
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Update(ctx context.Context, id int) (*entities.Board, error)
	// Create
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Create(ctx context.Context, info dto.NewBoardInfo) (*entities.Board, error)
	// Delete
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Delete(ctx context.Context, id int) error
}
