package service

import (
	"context"
	"log"
	"os"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"

	"github.com/sirupsen/logrus"
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
	funcName := "Create"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	defaultURL := "main_theme.jpg"

	if board.Thumbnail == nil {
		logger.Info("No thumbnail was provided, using the default thumbnail")
		board.ThumbnailURL = &defaultURL
	}
	newBoard, err := bs.storage.Create(ctx, board)

	fileLocation := "img/board_thumbnails/" + strconv.FormatUint(newBoard.ID, 10) + ".png"
	boardServiceDebugLog(logger, funcName, "Thumbnail location: "+fileLocation)
	// thumbnailUrlInfo := dto.ImageUrlInfo{
	// 	ID:  info.ID,
	// 	Url: info.BaseURL + "/" + fileLocation,
	// }
	// f, err := os.Create(thumbnailUrlInfo.Url)
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()
	return newBoard, err
}

// UpdateData
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	return bs.storage.UpdateData(ctx, info)
}

// UpdateThumbnail
// сохраняет картинку доски в папку images/board_thumbnails с названием id доски и сохраняет ссылку на изображение в БД
func (bs BoardService) UpdateThumbnail(ctx context.Context, info dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error) {
	baseURL := ctx.Value(dto.BaseURLKey).(string)
	fileLocation := "img/board_thumbnails/" + strconv.FormatUint(info.ID, 10) + ".png"
	log.Println("Service -- File location:", fileLocation)
	thumbnailUrlInfo := dto.ImageUrlInfo{
		ID:  info.ID,
		Url: baseURL + fileLocation,
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

func boardServiceDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "service",
			"function":   function,
		}).
		Debug(message)
}
