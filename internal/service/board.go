package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardService interface {
	GetBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	CreateBoard(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	UpdateBoard(context.Context, dto.IndividualBoardInfo) (*entities.Board, error)
	GetUserBoards(context.Context, entities.User) (*[]entities.Board, error)
}
