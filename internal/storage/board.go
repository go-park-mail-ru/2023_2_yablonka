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
	GetUserOwnedBoards(context.Context, dto.UserID) (*[]entities.Board, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(context.Context, dto.UserID) (*[]entities.Board, error)
	// GetById
	// находит доску и связанные с ней списки и задания по id
	// или возвращает ошибки ...
	GetById(context.Context, dto.BoardID) (*entities.Board, error)
	// Update
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Update(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	// Create
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// Delete
	// находит все доски, созданные пользователем
	// или возвращает ошибки ...
	Delete(context.Context, dto.BoardID) error
}
