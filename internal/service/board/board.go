package service

import (
	"context"
	"server/internal/pkg/dto"
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

// func (us BoardService) GetBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
// 	return us.storage.GetBoard(board)
// }

// func (us BoardService) CreateBoard(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
// 	return us.storage.CreateBoard(board)
// }

// func (us BoardService) UpdateBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
// 	return us.storage.UpdateBoard(board)
// }

func (us BoardService) GetUserOwnedBoards(ctx context.Context, userInfo dto.VerifiedAuthInfo) ([]dto.UserOwnedBoardInfo, error) {
	boards, err := us.storage.GetUserOwnedBoards(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	var boardInfo []dto.UserOwnedBoardInfo
	for _, board := range *boards {
		boardInfo = append(boardInfo, dto.UserOwnedBoardInfo{
			ID:           board.ID,
			BoardName:    board.Name,
			ThumbnailURL: board.ThumbnailURL,
		})
	}
	return boardInfo, nil
}

func (us BoardService) GetUserGuestBoards(ctx context.Context, userInfo dto.VerifiedAuthInfo) ([]dto.UserGuestBoardInfo, error) {
	boards, err := us.storage.GetUserGuestBoards(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	var boardInfo []dto.UserGuestBoardInfo
	for _, board := range *boards {
		boardInfo = append(boardInfo, dto.UserGuestBoardInfo{
			UserOwnedBoardInfo: dto.UserOwnedBoardInfo{
				ID:           board.ID,
				BoardName:    board.Name,
				ThumbnailURL: board.ThumbnailURL,
			},
			OwnerID:    board.Owner.ID,
			OwnerEmail: board.Owner.Email,
		})
	}
	return boardInfo, nil
}
