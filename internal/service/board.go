package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardService interface {
	// GetUsersInBoard
	// возвращает пользователей и их роли в доске
	// или возвращает ошибки ...
	GetFullBoard(context.Context, dto.IndividualBoardRequest) (*entities.Board, error)
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	UpdateData(context.Context, dto.UpdatedBoardInfo) error
	UpdateThumbnail(context.Context, dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error)
	Delete(context.Context, dto.BoardID) error
}
