package task

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const nodeName = "service"

type TaskService struct {
	storage     storage.ITaskStorage
	userStorage storage.IUserStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewTaskService(ts storage.ITaskStorage, us storage.IUserStorage) *TaskService {
	return &TaskService{
		storage:     ts,
		userStorage: us,
	}
}

// Read
// возвращает задание
// или возвращает ошибки ...
func (ts TaskService) Read(ctx context.Context, id dto.TaskID) (*dto.SingleTaskInfo, error) {
	return ts.storage.Read(ctx, id)
}

// Create
// создает новое задание
// или возвращает ошибки ...
func (ts TaskService) Create(ctx context.Context, info dto.NewTaskInfo) (*entities.Task, error) {
	return ts.storage.Create(ctx, info)
}

// Update
// обновляет задание
// или возвращает ошибки ...
func (ts TaskService) Update(ctx context.Context, info dto.UpdatedTaskInfo) error {
	return ts.storage.Update(ctx, info)
}

// Delete
// удаляет задание
// или возвращает ошибки ...
func (ts TaskService) Delete(ctx context.Context, id dto.TaskID) error {
	return ts.storage.Delete(ctx, id)
}

// AddUser
// добавляет пользователя в карточку
// или возвращает ошибки ...
func (ts TaskService) AddUser(ctx context.Context, info dto.AddTaskUserInfo) error {
	funcName := "TaskService.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	userAccess, err := ts.storage.CheckAccess(ctx, (dto.CheckTaskAccessInfo)(info))
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if userAccess {
		return apperrors.ErrUserAlreadyInTask
	}
	logger.DebugFmt("user not in task", requestID.String(), funcName, nodeName)

	return ts.storage.AddUser(ctx, info)
}

// RemoveUser
// удаляет пользователя из карточки
// или возвращает ошибки ...
func (ts TaskService) RemoveUser(ctx context.Context, info dto.RemoveTaskUserInfo) error {
	funcName := "TaskService.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	userAccess, err := ts.storage.CheckAccess(ctx, (dto.CheckTaskAccessInfo)(info))
	if err != nil {
		return apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("got user", requestID.String(), funcName, nodeName)

	if !userAccess {
		return apperrors.ErrUserNotInTask
	}
	logger.DebugFmt("user not in task", requestID.String(), funcName, nodeName)

	return ts.storage.RemoveUser(ctx, info)
}

// Move
// переносит задание в другой список
// или возвращает ошибки ...
func (ts TaskService) Move(ctx context.Context, taskMoveInfo dto.TaskMoveInfo) error {
	return ts.storage.Move(ctx, taskMoveInfo)
}

// GetFileList
// добавляет файл в задание
// или возвращает ошибки ...
func (ts TaskService) GetFileList(ctx context.Context, id dto.TaskID) (*[]dto.AttachedFileInfo, error) {
	return ts.storage.GetFileList(ctx, id)
}

// Attach
// добавляет файл в задание
// или возвращает ошибки ...
func (ts TaskService) Attach(ctx context.Context, info dto.NewFileInfo) (*dto.AttachedFileInfo, error) {
	funcName := "TaskService.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	fileName := hashFromFileInfo(info.Filename, info.Mimetype,
		strconv.FormatUint(info.UserID, 10), strconv.FormatUint(info.TaskID, 10))
	extension := path.Ext(info.Filename)

	if err := os.MkdirAll("attachments/task/"+strconv.FormatUint(info.TaskID, 10), 0755); err != nil {
		logger.DebugFmt("Failed to create directory with error: "+err.Error(), requestID.String(), funcName, nodeName)
		return &dto.AttachedFileInfo{}, apperrors.ErrFailedToCreateFile
	}

	fileLocation := "attachments/task/" + strconv.FormatUint(info.TaskID, 10) + "/" + fileName + extension

	f, err := os.Create(fileLocation)
	if err != nil {
		logger.DebugFmt("Failed to create file with error: "+err.Error(), requestID.String(), funcName, nodeName)
		return &dto.AttachedFileInfo{}, apperrors.ErrFailedToCreateFile
	}

	logger.DebugFmt(fmt.Sprintf("Writing %v bytes", len(info.File)), requestID.String(), funcName, nodeName)
	_, err = f.Write(info.File)
	if err != nil {
		logger.DebugFmt("Failed to create file with error: "+err.Error(), requestID.String(), funcName, nodeName)
		return &dto.AttachedFileInfo{}, apperrors.ErrFailedToSaveFile
	}
	createdAt := time.Now()

	defer f.Close()

	fileInfo := &dto.AttachedFileInfo{
		TaskID:       info.TaskID,
		OriginalName: info.Filename,
		FilePath:     fileLocation,
		DateCreated:  createdAt,
	}

	err = ts.storage.AttachFile(ctx, *fileInfo)
	if err != nil {
		errDelete := os.Remove(fileLocation)
		if errDelete != nil {
			logger.DebugFmt("Failed to remove file after unsuccessful update with error: "+err.Error(), requestID.String(), funcName, nodeName)
			return &dto.AttachedFileInfo{}, apperrors.ErrFailedToDeleteFile
		}
		return &dto.AttachedFileInfo{}, err
	}

	return fileInfo, nil
}

func hashFromFileInfo(strs ...string) string {
	hasher := sha256.New()
	hasher.Write([]byte(strings.Join(strs, "")))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
