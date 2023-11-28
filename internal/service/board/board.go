package service

import (
	"context"
	"os"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"
)

type BoardService struct {
	boardStorage storage.IBoardStorage
	userStorage  storage.IUserStorage
	listStorage  storage.IListStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewBoardService(
	bs storage.IBoardStorage,
	ls storage.IListStorage,
	us storage.IUserStorage,
) *BoardService {
	return &BoardService{
		boardStorage: bs,
		listStorage:  ls,
		userStorage:  us,
	}
}

const nodeName string = "service"

// GetFullBoard
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) GetFullBoard(ctx context.Context, info dto.IndividualBoardRequest) (*dto.FullBoardResult, error) {
	funcName := "BoardService.GetFullBoard"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	boardID := dto.BoardID{
		Value: info.BoardID,
	}

	users, err := bs.boardStorage.GetUsers(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.Debug("Got board users", funcName, nodeName)

	userHasAccessToBoard := false
	for _, user := range *users {
		if user.ID == info.UserID {
			userHasAccessToBoard = true
		}
	}
	if !userHasAccessToBoard {
		return nil, apperrors.ErrNoBoardAccess
	}
	logger.Debug("User has access to board", funcName, nodeName)

	board, err := bs.boardStorage.GetById(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.Debug("Got board", funcName, nodeName)

	lists, err := bs.boardStorage.GetLists(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.Debug("Got lists", funcName, nodeName)

	listIDs := dto.ListIDs{}
	for _, list := range *lists {
		listIDs.Values = append(listIDs.Values, list.ID)
	}
	tasks, err := bs.listStorage.GetTasksWithID(ctx, listIDs)
	if err != nil {
		return nil, err
	}
	logger.Debug("Got tasks", funcName, nodeName)

	return &dto.FullBoardResult{
		Users: *users,
		Board: *board,
		Lists: *lists,
		Tasks: *tasks,
	}, nil
}

// Create
// создаёт доску и связь с пользователем-создателем
func (bs BoardService) Create(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	funcName := "BoardService.Create"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	defaultURL := "main_theme.jpg"
	if board.Thumbnail == nil {
		board.ThumbnailURL = &defaultURL
	}
	newBoard, err := bs.boardStorage.Create(ctx, board)
	if err != nil {
		return nil, err
	}
	logger.Debug("Board created", funcName, nodeName)

	// thumbnailUrlInfo := dto.ImageUrlInfo{
	// 	ID:  info.ID,
	// 	Url: info.BaseURL + "/" + fileLocation,
	// }
	// f, err := os.Create(thumbnailUrlInfo.Url)
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()
	fileLocation := "img/board_thumbnails/" + strconv.FormatUint(newBoard.ID, 10) + ".png"
	logger.Debug("Thumbnail location: "+fileLocation, funcName, nodeName)
	return newBoard, nil
}

// UpdateData
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	return bs.boardStorage.UpdateData(ctx, info)
}

// UpdateThumbnail
// сохраняет картинку доски в папку images/board_thumbnails с названием id доски и сохраняет ссылку на изображение в БД
func (bs BoardService) UpdateThumbnail(ctx context.Context, info dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error) {
	funcName := "BoardService.UpdateThumbnail"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	baseURL := ctx.Value(dto.BaseURLKey).(string)
	fileLocation := "img/board_thumbnails/" + strconv.FormatUint(info.ID, 10) + ".png"
	logger.Debug("File location:"+fileLocation, funcName, nodeName)

	thumbnailUrlInfo := dto.ImageUrlInfo{
		ID:  info.ID,
		Url: baseURL + fileLocation,
	}
	f, err := os.Create(thumbnailUrlInfo.Url)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	logger.Debug("File created", funcName, nodeName)

	err = bs.boardStorage.UpdateThumbnailUrl(ctx, thumbnailUrlInfo)
	if err != nil {
		errDelete := os.Remove(thumbnailUrlInfo.Url)
		for errDelete != nil {
			errDelete = os.Remove(thumbnailUrlInfo.Url)
		}
		return nil, err
	}
	logger.Debug("thumbnail url updated", funcName, nodeName)

	return &dto.UrlObj{Value: thumbnailUrlInfo.Url}, nil
}

// Delete
// удаляет доску
func (bs BoardService) Delete(ctx context.Context, id dto.BoardID) error {
	return bs.boardStorage.Delete(ctx, id)
}

// AddUser
// добавляет пользователя на доску
func (bs BoardService) AddUser(ctx context.Context, request dto.AddBoardUserRequest) error {
	funcName := "BoardService.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	requestingUserID := ctx.Value(dto.UserObjKey).(*entities.User).ID

	if !hasAccess(bs.boardStorage, ctx, requestingUserID, request.BoardID) {
		return apperrors.ErrNoBoardAccess
	}
	logger.Debug("user has access to board", funcName, nodeName)

	targetUser, err := bs.userStorage.GetWithLogin(ctx, dto.UserLogin{Value: request.UserEmail})
	if err != nil {
		return err
	}
	logger.Debug("user found", funcName, nodeName)

	info := dto.AddBoardUserInfo{
		UserID:  targetUser.ID,
		BoardID: request.BoardID,
	}
	return bs.boardStorage.AddUser(ctx, info)
}

// RemoveUser
// добавляет пользователя на доску
func (bs BoardService) RemoveUser(ctx context.Context, info dto.RemoveBoardUserInfo) error {
	funcName := "BoardService.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	requestingUserID := ctx.Value(dto.UserObjKey).(*entities.User).ID
	if !hasAccess(bs.boardStorage, ctx, requestingUserID, info.BoardID) {
		return apperrors.ErrNoBoardAccess
	}
	logger.Debug("user has access to board", funcName, nodeName)

	return bs.boardStorage.RemoveUser(ctx, info)
}

func hasAccess(storage storage.IBoardStorage, ctx context.Context, userID uint64, boardID uint64) bool {
	hasAccess := false

	boardUsers, err := storage.GetUsers(ctx, dto.BoardID{Value: boardID})
	if err != nil {
		return hasAccess
	}

	for _, user := range *boardUsers {
		if user.ID == userID {
			hasAccess = true
		}
	}

	return hasAccess
}
