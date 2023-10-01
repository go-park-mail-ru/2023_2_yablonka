package service

import (
	"context"
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

// TODO
func (us BoardService) GetBoard() {

}

// TODO
func (us BoardService) CreateBoard() {

}

// TODO
func (us BoardService) UpdateBoard(ctx context.Context, Board *entities.Board) {

}

// TODO
func (us BoardService) GetBoardUsers() {

}
