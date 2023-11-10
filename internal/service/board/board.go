package service

import (
	"context"
	"os"
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

// GetFullBoard
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) GetFullBoard(ctx context.Context, info dto.IndividualBoardRequest) (*entities.Board, error) {
	boardID := dto.BoardID{
		Value: info.BoardID,
	}

	boardUsers, err := bs.storage.GetUsers(ctx, boardID)
	if err != nil {
		return nil, err
	}

	userHasAccessToBoard := false
	for _, user := range *boardUsers {
		if user.ID == info.UserID {
			userHasAccessToBoard = true
		}
	}

	if !userHasAccessToBoard {
		return nil, err
	}

	board, err := bs.storage.GetById(ctx, boardID)
	if err != nil {
		return nil, err
	}

	board.Users = *boardUsers

	return board, nil
}

// Create
// создаёт доску и связь с пользователем-создателем
func (bs BoardService) Create(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	return bs.storage.Create(ctx, board)
}

// UpdateData
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	return bs.storage.UpdateData(ctx, info)
}

// UpdateThumbnail
// сохраняет картинку доски в папку images/board_thumbnails с названием id доски и сохраняет ссылку на изображение в БД
func (bs BoardService) UpdateThumbnail(ctx context.Context, info dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error) {
	thumbnailUrlInfo := dto.ImageUrlInfo{
		ID:  info.ID,
		Url: "images/board_thumbnails/" + string(info.ID) + ".png",
	}
	f, err := os.Create(thumbnailUrlInfo.Url)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = bs.storage.UpdateThumbnailUrl(ctx, thumbnailUrlInfo)
	if err != nil {
		errDelete := os.Remove(thumbnailUrlInfo.Url)
		for errDelete != nil {
			errDelete = os.Remove(thumbnailUrlInfo.Url)
		}
		return nil, err
	}

	return &dto.UrlObj{Value: thumbnailUrlInfo.Url}, nil
}

// Delete
// удаляет доску
func (bs BoardService) Delete(ctx context.Context, id dto.BoardID) error {
	return bs.storage.Delete(ctx, id)
}
