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

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewBoardService(storage storage.IBoardStorage) *BoardService {
	return &BoardService{
		storage: storage,
	}
}

// GetUserOwnedBoards
// находит все доски, созданные пользователем
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us BoardService) GetUserOwnedBoards(ctx context.Context, id dto.UserID) ([]dto.UserOwnedBoardInfo, error) {
	boards, err := us.storage.GetUserOwnedBoards(ctx, id)
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

// GetUserGuestBoards
// находит все доски, в которых участвует пользователь
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us BoardService) GetUserGuestBoards(ctx context.Context, id dto.UserID) ([]dto.UserGuestBoardInfo, error) {
	boards, err := us.storage.GetUserGuestBoards(ctx, id)
	if err != nil {
		return nil, err
	}

	var boardInfo []dto.UserGuestBoardInfo
	for _, board := range *boards {
		boardInfo = append(boardInfo, dto.UserGuestBoardInfo{
			BoardInfo: dto.UserOwnedBoardInfo{
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

func (bs BoardService) GetBoardWithListsAndTasks(ctx context.Context, board dto.BoardID) (*entities.Board, error) {
	return bs.storage.GetById(ctx, board)
}

func (us BoardService) CreateBoard(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	return us.storage.Create(ctx, board)
}

func (us BoardService) UpdateBoard(ctx context.Context, board dto.IndividualBoardInfo) (*entities.Board, error) {
	return us.storage.Update(ctx, board)
}
