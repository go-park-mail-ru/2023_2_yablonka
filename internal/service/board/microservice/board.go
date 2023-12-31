package microservice

import (
	"context"
	"os"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type BoardService struct {
	boardStorage         storage.IBoardStorage
	userStorage          storage.IUserStorage
	taskStorage          storage.ITaskStorage
	commentStorage       storage.ICommentStorage
	checklistStorage     storage.IChecklistStorage
	checklistItemStorage storage.IChecklistItemStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewBoardService(
	bs storage.IBoardStorage,
	ts storage.ITaskStorage,
	us storage.IUserStorage,
	cs storage.ICommentStorage,
	cls storage.IChecklistStorage,
	clis storage.IChecklistItemStorage,
	conn *grpc.ClientConn,
) *BoardService {
	return &BoardService{
		boardStorage:         bs,
		taskStorage:          ts,
		userStorage:          us,
		commentStorage:       cs,
		checklistStorage:     cls,
		checklistItemStorage: clis,
	}
}

const nodeName = "service"

// GetFullBoard
// возвращает доску со связанными пользователями, списками и заданиями
func (bs BoardService) GetFullBoard(ctx context.Context, info dto.IndividualBoardRequest) (*dto.FullBoardResult, error) {
	funcName := "BoardService.GetFullBoard"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	boardID := dto.BoardID{
		Value: info.BoardID,
	}

	users, err := bs.boardStorage.GetUsers(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got board users", requestID.String(), funcName, nodeName)

	userHasAccessToBoard := false
	for _, user := range *users {
		if user.ID == info.UserID {
			userHasAccessToBoard = true
		}
	}
	if !userHasAccessToBoard {
		return nil, apperrors.ErrNoBoardAccess
	}
	logger.DebugFmt("User has access to board", requestID.String(), funcName, nodeName)

	board, err := bs.boardStorage.GetById(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got board", requestID.String(), funcName, nodeName)

	lists, err := bs.boardStorage.GetLists(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got lists", requestID.String(), funcName, nodeName)

	taskIDs := dto.TaskIDs{}
	for _, list := range *lists {
		taskIDs.Values = append(taskIDs.Values, list.TaskIDs...)
	}
	tasks, err := bs.taskStorage.ReadMany(ctx, taskIDs)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got tasks", requestID.String(), funcName, nodeName)

	commentIDs := dto.CommentIDs{}
	checklistIDs := dto.ChecklistIDs{}
	for _, task := range *tasks {
		commentIDs.Values = append(commentIDs.Values, task.CommentIDs...)
		checklistIDs.Values = append(checklistIDs.Values, task.ChecklistIDs...)
	}
	logger.DebugFmt("Got comment ids", requestID.String(), funcName, nodeName)

	comments, err := bs.commentStorage.ReadMany(ctx, commentIDs)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got comments", requestID.String(), funcName, nodeName)

	checklists, err := bs.checklistStorage.ReadMany(ctx, checklistIDs)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got checklists", requestID.String(), funcName, nodeName)

	checklistItemIDs := dto.ChecklistItemStringIDs{}
	for _, checklist := range *checklists {
		checklistItemIDs.Values = append(checklistItemIDs.Values, checklist.Items...)
	}
	checklistItems, err := bs.checklistItemStorage.ReadMany(ctx, checklistItemIDs)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got checklist items", requestID.String(), funcName, nodeName)

	tags, err := bs.boardStorage.GetTags(ctx, boardID)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Got tags", requestID.String(), funcName, nodeName)

	return &dto.FullBoardResult{
		Users:          *users,
		Board:          *board,
		Lists:          *lists,
		Tags:           *tags,
		Tasks:          *tasks,
		Comments:       *comments,
		Checklists:     *checklists,
		ChecklistItems: *checklistItems,
	}, nil
}

// Create
// создаёт доску и связь с пользователем-создателем
func (bs BoardService) Create(ctx context.Context, board dto.NewBoardInfo) (*entities.Board, error) {
	funcName := "BoardService.Create"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	defaultURL := "main_theme.jpg"
	if board.Thumbnail == nil {
		board.ThumbnailURL = &defaultURL
	}
	newBoard, err := bs.boardStorage.Create(ctx, board)
	if err != nil {
		return nil, err
	}
	logger.DebugFmt("Board created", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("Thumbnail location: "+fileLocation, requestID.String(), funcName, nodeName)
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
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	fileLocation := "img/board_thumbnails/" + strconv.FormatUint(info.ID, 10) + ".png"
	logger.DebugFmt("File location:"+fileLocation, requestID.String(), funcName, nodeName)

	thumbnailUrlInfo := dto.BoardImageUrlInfo{
		ID:  info.ID,
		Url: fileLocation,
	}
	f, err := os.Create(thumbnailUrlInfo.Url)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	logger.DebugFmt("File created", requestID.String(), funcName, nodeName)

	err = bs.boardStorage.UpdateThumbnailUrl(ctx, thumbnailUrlInfo)
	if err != nil {
		errDelete := os.Remove(thumbnailUrlInfo.Url)
		for errDelete != nil {
			errDelete = os.Remove(thumbnailUrlInfo.Url)
		}
		return nil, err
	}
	logger.DebugFmt("thumbnail url updated", requestID.String(), funcName, nodeName)

	return &dto.UrlObj{Value: thumbnailUrlInfo.Url}, nil
}

// Delete
// удаляет доску
func (bs BoardService) Delete(ctx context.Context, info dto.BoardDeleteRequest) error {
	return bs.boardStorage.Delete(ctx, info)
}

// AddUser
// добавляет пользователя на доску
func (bs BoardService) AddUser(ctx context.Context, request dto.AddBoardUserRequest) (dto.UserPublicInfo, error) {
	funcName := "BoardService.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	accessInfo := dto.CheckBoardAccessInfo{
		UserID:  ctx.Value(dto.UserObjKey).(*entities.User).ID,
		BoardID: request.BoardID,
	}
	requestingUserAccess, err := bs.boardStorage.CheckAccess(ctx, accessInfo)
	if err != nil {
		return dto.UserPublicInfo{}, apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if !requestingUserAccess {
		return dto.UserPublicInfo{}, apperrors.ErrNoBoardAccess
	}
	logger.DebugFmt("Requesting user has access to board", requestID.String(), funcName, nodeName)

	targetUser, err := bs.userStorage.GetWithLogin(ctx, dto.UserLogin{Value: request.UserEmail})
	if err != nil {
		return dto.UserPublicInfo{}, apperrors.ErrUserNotFound
	}
	logger.DebugFmt("user found", requestID.String(), funcName, nodeName)

	accessInfo.UserID = targetUser.ID
	userAccess, err := bs.boardStorage.CheckAccess(ctx, accessInfo)
	if err != nil {
		return dto.UserPublicInfo{}, apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if userAccess {
		return dto.UserPublicInfo{}, apperrors.ErrUserAlreadyInBoard
	}
	logger.DebugFmt("user not in board", requestID.String(), funcName, nodeName)

	info := dto.AddBoardUserInfo{
		UserID:      targetUser.ID,
		BoardID:     request.BoardID,
		WorkspaceID: request.WorkspaceID,
	}
	return bs.boardStorage.AddUser(ctx, info)
}

// RemoveUser
// добавляет пользователя на доску
func (bs BoardService) RemoveUser(ctx context.Context, info dto.RemoveBoardUserInfo) error {
	funcName := "BoardService.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	accessInfo := dto.CheckBoardAccessInfo{
		UserID:  ctx.Value(dto.UserObjKey).(*entities.User).ID,
		BoardID: info.BoardID,
	}
	requestingUserAccess, err := bs.boardStorage.CheckAccess(ctx, accessInfo)
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if !requestingUserAccess {
		return apperrors.ErrNoBoardAccess
	}
	logger.DebugFmt("user has access to board", requestID.String(), funcName, nodeName)

	accessInfo.UserID = info.UserID
	userAccess, err := bs.boardStorage.CheckAccess(ctx, accessInfo)
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if !userAccess {
		return apperrors.ErrUserNotInBoard
	}
	logger.DebugFmt("user in board", requestID.String(), funcName, nodeName)

	return bs.boardStorage.RemoveUser(ctx, info)
}

// GetHistory
// возвращает историю изменения доски
func (bs BoardService) GetHistory(ctx context.Context, id dto.BoardID) (*[]dto.BoardHistoryEntry, error) {
	return bs.boardStorage.GetHistory(ctx, id)
}

// SubmitEdit
// записывает изменение доски в историю
func (bs BoardService) SubmitEdit(ctx context.Context, entry dto.NewHistoryEntry) error {
	return bs.boardStorage.SubmitEdit(ctx, entry)
}
