package user_microservice

import (
	context "context"
	"crypto/sha256"
	"fmt"
	"os"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"strconv"

	"google.golang.org/protobuf/types/known/emptypb"
)

const nodeName = "service"

type UserService struct {
	storage storage.IUserStorage
	logger  *logger.LogrusLogger
	UnimplementedUserServiceServer
}

// NewUserService
// возвращает UserService с инициализированным хранилищем пользователей
func NewUserService(storage storage.IUserStorage, logger *logger.LogrusLogger) *UserService {
	return &UserService{
		storage: storage,
		logger:  logger,
	}
}

// RegisterUser
// создает нового пользователя по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
func (us UserService) RegisterUser(ctx context.Context, info *AuthInfo) (*User, error) {
	funcName := "UserService.RegisterUser"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)

	_, err := us.storage.GetWithLogin(sCtx, dto.UserLogin{Value: info.Email})

	if err == nil {
		return &User{}, apperrors.MakeGRPCError(apperrors.ErrUserAlreadyExists)
	}
	us.logger.Debug("User doesn't exist", funcName, nodeName)

	us.logger.Debug("Creating user", funcName, nodeName)
	user, err := us.storage.Create(sCtx, dto.SignupInfo{
		Email:        info.Email,
		PasswordHash: hashFromAuthInfo(info.Email, info.Password),
	})
	if err != nil {
		return &User{}, apperrors.MakeGRPCError(err)
	}

	convertedUser := convertUser(user)

	return convertedUser, nil
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info *AuthInfo) (*User, error) {
	funcName := "UserService.CheckPassword"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)

	user, err := us.storage.GetWithLogin(sCtx, dto.UserLogin{Value: info.Email})
	if err != nil {
		us.logger.Debug("User not found", funcName, nodeName)
		return &User{}, apperrors.MakeGRPCError(err)
	}
	us.logger.Debug("User found", funcName, nodeName)

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		return &User{}, apperrors.MakeGRPCError(apperrors.ErrWrongPassword)
	}
	us.logger.Debug("Password match", funcName, nodeName)

	convertedUser := convertUser(user)

	return convertedUser, nil
}

// GetWithID
// находит пользователя по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us UserService) GetWithID(ctx context.Context, id *UserID) (*User, error) {
	funcName := "UserService.GetWithID"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)

	user, err := us.storage.GetWithID(sCtx, dto.UserID{Value: id.Value})
	if err != nil {
		us.logger.Debug("Failed to get user with error "+err.Error(), funcName, nodeName)
		return &User{}, apperrors.MakeGRPCError(err)
	}
	us.logger.Debug("Got user", funcName, nodeName)

	convertedUser := convertUser(user)

	return convertedUser, nil
}

// UpdatePassword
// меняет пароль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdatePassword(ctx context.Context, info *PasswordChangeInfo) (*emptypb.Empty, error) {
	funcName := "UserService.UpdatePassword"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)

	oldLoginInfo, err := us.storage.GetLoginInfoWithID(sCtx, dto.UserID{Value: info.UserID})
	if err != nil {
		return &emptypb.Empty{}, apperrors.MakeGRPCError(apperrors.ErrUserNotFound)
	}
	us.logger.Debug("User found", funcName, nodeName)

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		return &emptypb.Empty{}, apperrors.MakeGRPCError(apperrors.ErrWrongPassword)
	}
	us.logger.Debug("Old verified", funcName, nodeName)

	return &emptypb.Empty{}, apperrors.MakeGRPCError(
		us.storage.UpdatePassword(sCtx, dto.PasswordHashesInfo{
			UserID:          info.UserID,
			NewPasswordHash: hashFromAuthInfo(oldLoginInfo.Email, info.NewPassword),
		}))
}

// UpdateProfile
// обновляет профиль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateProfile(ctx context.Context, info *UserProfileInfo) (*emptypb.Empty, error) {
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)

	return &emptypb.Empty{}, apperrors.MakeGRPCError(
		us.storage.UpdateProfile(sCtx, dto.UserProfileInfo{
			UserID:      info.UserID,
			Name:        info.Name,
			Surname:     info.Surname,
			Description: info.Description,
		}))
}

// UpdateProfile
// обновляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateAvatar(ctx context.Context, info *AvatarChangeInfo) (*UrlObj, error) {
	baseURL := info.BaseURL
	funcName := "UserService.UpdateAvatar"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)

	cwd, _ := os.Getwd()
	fileLocation := "img/user_avatars/" + strconv.FormatUint(info.UserID, 10) + ".png"
	us.logger.Debug("Relative path: "+fileLocation, funcName, nodeName)
	us.logger.Debug("CWD: "+cwd, funcName, nodeName)
	avatarUrlInfo := dto.ImageUrlInfo{
		ID:  info.UserID,
		Url: baseURL + fileLocation,
	}
	us.logger.Debug("Full URL: "+avatarUrlInfo.Url, funcName, nodeName)
	f, err := os.Create(fileLocation)
	if err != nil {
		us.logger.Debug("Failed to create file with error: "+err.Error(), funcName, nodeName)
		return nil, apperrors.MakeGRPCError(err)
	}

	us.logger.Debug(fmt.Sprintf("Writing %v bytes", len(info.Avatar)), funcName, nodeName)
	_, err = f.Write(info.Avatar)
	if err != nil {
		us.logger.Debug("Failed to write to file with error: "+err.Error(), funcName, nodeName)
		return nil, apperrors.MakeGRPCError(err)
	}

	defer f.Close()

	err = us.storage.UpdateAvatarUrl(sCtx, avatarUrlInfo)
	if err != nil {
		errDelete := os.Remove(fileLocation)
		for errDelete != nil {
			us.logger.Debug("Failed to remove file after unsuccessful update with error: "+err.Error(), funcName, nodeName)
			errDelete = os.Remove(fileLocation)
		}
		return nil, apperrors.MakeGRPCError(err)
	}

	return &UrlObj{Value: avatarUrlInfo.Url}, nil
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteUser(ctx context.Context, id *UserID) (*emptypb.Empty, error) {
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	return &emptypb.Empty{}, apperrors.MakeGRPCError(us.storage.Delete(sCtx, dto.UserID{Value: id.Value}))
}

// TODO salt
func hashFromAuthInfo(email string, password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email + password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func convertUser(user *entities.User) *User {
	convertedUser := User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
	if user.Name == nil {
		convertedUser.Name = ""
	} else {
		convertedUser.Name = *user.Name
	}
	if user.Surname == nil {
		convertedUser.Surname = ""
	} else {
		convertedUser.Surname = *user.Surname
	}
	if user.Description == nil {
		convertedUser.Description = ""
	} else {
		convertedUser.Description = *user.Description
	}
	if user.AvatarURL == nil {
		convertedUser.AvatarURL = ""
	} else {
		convertedUser.AvatarURL = *user.AvatarURL
	}
	return &convertedUser
}
