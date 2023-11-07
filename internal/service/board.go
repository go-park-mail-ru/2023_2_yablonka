package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardService interface {
	// GetUserOwnedBoards
	// находит все доски, созданные пользователем
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserOwnedBoards(context.Context, dto.UserID) ([]dto.UserOwnedBoardInfo, error)
	// GetUserGuestBoards
	// находит все доски, в которых участвует пользователь
	// или возвращает ошибку apperrors.ErrUserNotFound (401)
	GetUserGuestBoards(context.Context, dto.UserID) ([]dto.UserGuestBoardInfo, error)
	// GetUsersInBoard
	// возвращает пользователей и их роли в доске
	// или возвращает ошибки ...
	GetBoardWithListsAndTasks(context.Context, dto.BoardID) (*entities.Board, error)
	CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
}
