package embedded

import (
	"context"
	"fmt"
	"log"
	"os"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"

	"github.com/sirupsen/logrus"
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

// GetFullBoard
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) GetFullBoard(ctx context.Context, info dto.IndividualBoardRequest) (*dto.FullBoardResult, error) {
	funcName := "GetFullBoard"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	boardID := dto.BoardID{
		Value: info.BoardID,
	}

	users, err := bs.boardStorage.GetUsers(ctx, boardID)
	if err != nil {
		boardServiceDebugLog(logger, funcName, "Failed to get board users with error "+err.Error())
		return nil, err
	}
	boardServiceDebugLog(logger, funcName, "Got board users")

	userHasAccessToBoard := false
	for _, user := range *users {
		if user.ID == info.UserID {
			userHasAccessToBoard = true
		}
	}
	if !userHasAccessToBoard {
		logger.Warn(fmt.Sprintf("Requesting user (ID %d) doesn't have access to the board (ID %d)", info.UserID, info.BoardID))
		return nil, apperrors.ErrNoBoardAccess
	}
	boardServiceDebugLog(logger, funcName, "User has access to board")

	board, err := bs.boardStorage.GetById(ctx, boardID)
	if err != nil {
		boardServiceDebugLog(logger, funcName, "Failed to get board from storage with error "+err.Error())
		return nil, err
	}
	boardServiceDebugLog(logger, funcName, "Got board")

	lists, err := bs.boardStorage.GetLists(ctx, boardID)
	if err != nil {
		boardServiceDebugLog(logger, funcName, "Failed to get board from storage with error "+err.Error())
		return nil, err
	}
	boardServiceDebugLog(logger, funcName, "Got lists")

	listIDs := dto.ListIDs{}
	for _, list := range *lists {
		listIDs.Values = append(listIDs.Values, list.ID)
	}

	tasks, err := bs.listStorage.GetTasksWithID(ctx, listIDs)
	if err != nil {
		boardServiceDebugLog(logger, funcName, "Failed to get board from storage with error "+err.Error())
		return nil, err
	}
	boardServiceDebugLog(logger, funcName, "Got tasks")

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
	funcName := "Create"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	defaultURL := "main_theme.jpg"

	if board.Thumbnail == nil {
		logger.Info("No thumbnail was provided, using the default thumbnail")
		board.ThumbnailURL = &defaultURL
	}
	newBoard, err := bs.boardStorage.Create(ctx, board)

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
	return bs.boardStorage.UpdateData(ctx, info)
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

	err = bs.boardStorage.UpdateThumbnailUrl(ctx, thumbnailUrlInfo)
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
	return bs.boardStorage.Delete(ctx, id)
}

// AddUser
// добавляет пользователя на доску
func (bs BoardService) AddUser(ctx context.Context, request dto.AddBoardUserRequest) error {
	funcName := "AddUser"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	requestingUserID := ctx.Value(dto.UserObjKey).(*entities.User).ID

	if !hasAccess(bs.boardStorage, ctx, requestingUserID, request.BoardID) {
		logger.Warn(fmt.Sprintf("Requesting user (ID %d) doesn't have access to the board (ID %d)", requestingUserID, request.BoardID))
		return apperrors.ErrNoBoardAccess
	}

	targetUser, err := bs.userStorage.GetWithLogin(ctx, dto.UserLogin{Value: request.UserEmail})
	if err != nil {
		return err
	}

	info := dto.AddBoardUserInfo{
		UserID:  targetUser.ID,
		BoardID: request.BoardID,
	}

	boardServiceDebugLog(logger, funcName, "Adding user")
	return bs.boardStorage.AddUser(ctx, info)
}

// RemoveUser
// добавляет пользователя на доску
func (bs BoardService) RemoveUser(ctx context.Context, info dto.RemoveBoardUserInfo) error {
	funcName := "RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	requestingUserID := ctx.Value(dto.UserObjKey).(*entities.User).ID

	if !hasAccess(bs.boardStorage, ctx, requestingUserID, info.BoardID) {
		logger.Warn(fmt.Sprintf("Requesting user (ID %d) doesn't have access to the board (ID %d)", requestingUserID, info.BoardID))
		return apperrors.ErrNoBoardAccess
	}

	boardServiceDebugLog(logger, funcName, "Removing user")
	return bs.boardStorage.RemoveUser(ctx, info)
}

func boardServiceDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "service",
			"function":   function,
		}).
		Debug(message)
}

func hasAccess(storage storage.IBoardStorage, ctx context.Context, userID uint64, boardID uint64) bool {
	hasAccess := false

	funcName := "hasAccess"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	boardUsers, err := storage.GetUsers(ctx, dto.BoardID{Value: boardID})
	if err != nil {
		boardServiceDebugLog(logger, funcName, "Failed to get board users with error "+err.Error())
		return hasAccess
	}
	boardServiceDebugLog(logger, funcName, fmt.Sprintf("Got %d board users", len(*boardUsers)))

	for _, user := range *boardUsers {
		boardServiceDebugLog(logger, funcName, fmt.Sprintf("User ID %v", user.ID))
		if user.ID == userID {
			boardServiceDebugLog(logger, funcName, "Match")
			hasAccess = true
		}
	}

	return hasAccess
}
