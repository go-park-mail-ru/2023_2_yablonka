package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"
)

const nodeName string = "service"

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
	funcName := "BoardService.UpdateThumbnail"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	_, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})

	if err == nil {
		return nil, apperrors.ErrUserAlreadyExists
	}
	logger.Debug("User doesn't exist", funcName, nodeName)

	return us.storage.Create(ctx, dto.SignupInfo{
		Email:        info.Email,
		PasswordHash: hashFromAuthInfo(info.Email, info.Password),
	})
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	funcName := "BoardService.UpdateThumbnail"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	user, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})
	if err != nil {
		return nil, err
	}
	logger.Debug("got user", funcName, nodeName)

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		return nil, apperrors.ErrWrongPassword
	}
	logger.Debug("passwords match", funcName, nodeName)

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
	funcName := "BoardService.UpdateThumbnail"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	oldLoginInfo, err := us.storage.GetLoginInfoWithID(ctx, dto.UserID{Value: info.UserID})
	if err != nil {
		return apperrors.ErrUserNotFound
	}
	logger.Debug("got user", funcName, nodeName)

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		return apperrors.ErrWrongPassword
	}
	logger.Debug("passwords match", funcName, nodeName)

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
	funcName := "BoardService.UpdateThumbnail"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	baseURL := ctx.Value(dto.BaseURLKey).(string)
	fileLocation := "img/user_avatars/" + strconv.FormatUint(info.UserID, 10) + ".png"
	avatarUrlInfo := dto.ImageUrlInfo{
		ID:  info.UserID,
		Url: baseURL + fileLocation,
	}

	f, err := os.Create(fileLocation)
	if err != nil {
		return nil, err
	}
	logger.Debug("file created", funcName, nodeName)

	_, err = f.Write(info.Avatar)
	if err != nil {
		return nil, err
	}
	logger.Debug("avatar written", funcName, nodeName)

	defer f.Close()

	err = us.storage.UpdateAvatarUrl(ctx, avatarUrlInfo)
	if err != nil {
		errDelete := os.Remove(fileLocation)
		for errDelete != nil {
			errDelete = os.Remove(fileLocation)
		}
		return nil, err
	}
	logger.Debug("avatar url updated", funcName, nodeName)

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
