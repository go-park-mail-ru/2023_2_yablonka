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

func (bs BoardService) GetBoardWithListsAndTasks(ctx context.Context, id dto.BoardID) (*entities.Board, error) {
	return bs.storage.GetById(ctx, id)
}

func (bs BoardService) Create(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	return bs.storage.Create(ctx, board)
}

func (bs BoardService) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	return bs.storage.UpdateData(ctx, info)
}

func (bs BoardService) UpdateThumbnail(ctx context.Context, info dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error) {
	thumbnailUrlInfo := dto.ImageUrlInfo{
		ID:  info.ID,
		Url: "images/user_avatars/" + info.ID + ".png",
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

func (bs BoardService) Delete(ctx context.Context, id dto.BoardID) error {
	return bs.storage.Delete(ctx, id)
}
