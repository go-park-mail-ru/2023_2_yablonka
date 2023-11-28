package embedded

import (
	"context"
	"crypto/sha256"
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
	funcName := "RegisterUser"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	_, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})

	if err == nil {
		logger.Warn("User exists")
		return nil, apperrors.ErrUserAlreadyExists
	}
	userServiceDebugLog(logger, funcName, "User doesn't exist")

	userServiceDebugLog(logger, funcName, "Creating user")
	return us.storage.Create(ctx, dto.SignupInfo{
		Email:        info.Email,
		PasswordHash: hashFromAuthInfo(info.Email, info.Password),
	})
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	funcName := "RegisterUser"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	user, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})
	if err != nil {
		logger.Warn("User not found")
		return nil, err
	}

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		logger.Warn("Wrong password")
		return nil, apperrors.ErrWrongPassword
	}

	userServiceDebugLog(logger, funcName, "Returning user")
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
	funcName := "UpdatePassword"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	userServiceDebugLog(logger, funcName, fmt.Sprintf("Fetching login info for user ID %d", info.UserID))
	oldLoginInfo, err := us.storage.GetLoginInfoWithID(ctx, dto.UserID{Value: info.UserID})
	if err != nil {
		logger.Warn("User not found")
		return apperrors.ErrUserNotFound
	}
	userServiceDebugLog(logger, funcName, "User found")

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		logger.Warn("Wrong password")
		return apperrors.ErrWrongPassword
	}
	userServiceDebugLog(logger, funcName, "Password verified")

	userServiceDebugLog(logger, funcName, "Updating password")
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
	fileLocation := "img/user_avatars/" + strconv.FormatUint(info.UserID, 10) + ".png"
	log.Println("Service -- File location:", fileLocation)
	avatarUrlInfo := dto.ImageUrlInfo{
		ID:  info.UserID,
		Url: baseURL + fileLocation,
	}
	log.Println("Service -- Full url to file", avatarUrlInfo.Url)
	f, err := os.Create(fileLocation)
	if err != nil {
		log.Println("Service -- Failed to create file with error", err)
		return nil, err
	}

	log.Println("Service -- Writing", len(info.Avatar), "bytes:\n\t", info.Avatar)
	_, err = f.Write(info.Avatar)
	if err != nil {
		log.Println("Service -- Failed to write to file with error", err)
	}

	defer f.Close()

	err = us.storage.UpdateAvatarUrl(ctx, avatarUrlInfo)
	if err != nil {
		errDelete := os.Remove(fileLocation)
		for errDelete != nil {
			log.Println("Service -- Failed to remove file after unsuccessful update with error", err)
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

func userServiceDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "service",
			"function":   function,
		}).
		Debug(message)
}

// func userServiceWarnLog(logger *logrus.Logger, function string, message string) {
// 	logger.
// 		WithFields(logrus.Fields{
// 			"route_node": "service",
// 			"function":   function,
// 		}).
// 		Warn(message)
// }
