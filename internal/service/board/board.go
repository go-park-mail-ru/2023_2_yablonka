package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type BoardService struct {
	storage storage.IBoardStorage
}

func NewBoardService(storage storage.IBoardStorage) *BoardService {
	return &BoardService{
		storage: storage,
	}
}

func (us BoardService) GetBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
	return us.storage.GetBoard(board)
}

func (us BoardService) CreateBoard(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	return us.storage.CreateBoard(board)
}

func (us BoardService) UpdateBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
	return us.storage.UpdateBoard(board)
}

func (us BoardService) GetUserBoards(ctx context.Context, user entities.User) (*[]entities.Board, error) {
	return us.storage.GetUserBoards(user)
}
