package embedded

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"

	logger "server/internal/logging"
)

const nodeName = "service"

type UserService struct {
	storage storage.IUserStorage
}

// NewUserService
// возвращает UserService с инициализированным хранилищем пользователей
func NewUserService(storage storage.IUserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

// RegisterUser
// создает нового пользователя по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
func (us UserService) RegisterUser(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	funcName := "UserService.RegisterUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	_, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})

	if err == nil {
		return nil, apperrors.ErrUserAlreadyExists
	}
	logger.DebugFmt("User doesn't exist", funcName, nodeName)

	logger.DebugFmt("Creating user", funcName, nodeName)
	return us.storage.Create(ctx, dto.SignupInfo{
		Email:        info.Email,
		PasswordHash: hashFromAuthInfo(info.Email, info.Password),
	})
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	funcName := "UserService.CheckPassword"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	user, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})
	if err != nil {
		logger.DebugFmt("User not found", funcName, nodeName)
		return nil, err
	}
	logger.DebugFmt("User found", funcName, nodeName)

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		return nil, apperrors.ErrWrongPassword
	}
	logger.DebugFmt("Password match", funcName, nodeName)

	return user, nil
}

// GetWithID
// находит пользователя по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us UserService) GetWithID(ctx context.Context, id dto.UserID) (*entities.User, error) {
	return us.storage.GetWithID(ctx, id)
}

// UpdatePassword
// меняет пароль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdatePassword(ctx context.Context, info dto.PasswordChangeInfo) error {
	funcName := "UserService.UpdatePassword"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	oldLoginInfo, err := us.storage.GetLoginInfoWithID(ctx, dto.UserID{Value: info.UserID})
	if err != nil {
		return apperrors.ErrUserNotFound
	}
	logger.DebugFmt("User found", funcName, nodeName)

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		return apperrors.ErrWrongPassword
	}
	logger.DebugFmt("Old verified", funcName, nodeName)

	return us.storage.UpdatePassword(ctx, dto.PasswordHashesInfo{
		UserID:          info.UserID,
		NewPasswordHash: hashFromAuthInfo(oldLoginInfo.Email, info.NewPassword),
	})
}

// UpdateProfile
// обновляет профиль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateProfile(ctx context.Context, info dto.UserProfileInfo) error {
	return us.storage.UpdateProfile(ctx, info)
}

// UpdateProfile
// обновляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateAvatar(ctx context.Context, info dto.AvatarChangeInfo) (*dto.UrlObj, error) {
	baseURL := ctx.Value(dto.BaseURLKey).(string)
	funcName := "UserService.UpdateAvatar"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	cwd, _ := os.Getwd()
	fileLocation := "img/user_avatars/" + strconv.FormatUint(info.UserID, 10) + ".png"
	logger.DebugFmt("Relative path: "+fileLocation, funcName, nodeName)
	logger.DebugFmt("CWD: "+cwd, funcName, nodeName)
	avatarUrlInfo := dto.ImageUrlInfo{
		ID:  info.UserID,
		Url: baseURL + fileLocation,
	}
	logger.DebugFmt("Full URL: "+avatarUrlInfo.Url, funcName, nodeName)
	f, err := os.Create(fileLocation)
	if err != nil {
		logger.DebugFmt("Failed to create file with error: "+err.Error(), funcName, nodeName)
		return nil, err
	}

	logger.DebugFmt(fmt.Sprintf("Writing %v bytes", len(info.Avatar)), funcName, nodeName)
	_, err = f.Write(info.Avatar)
	if err != nil {
		logger.DebugFmt("Failed to write to file with error: "+err.Error(), funcName, nodeName)
	}

	defer f.Close()

	err = us.storage.UpdateAvatarUrl(ctx, avatarUrlInfo)
	if err != nil {
		errDelete := os.Remove(fileLocation)
		for errDelete != nil {
			logger.DebugFmt("Failed to remove file after unsuccessful update with error: "+err.Error(), funcName, nodeName)
			errDelete = os.Remove(fileLocation)
		}
		return nil, err
	}

	return &dto.UrlObj{Value: avatarUrlInfo.Url}, nil
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteUser(ctx context.Context, id dto.UserID) error {
	return us.storage.Delete(ctx, id)
}

// TODO salt
func hashFromAuthInfo(email string, password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email + password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
