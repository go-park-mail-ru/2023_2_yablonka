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
)

const nodeName = "service"

type UserService struct {
	storage storage.IUserStorage
	logger  *logger.LogrusLogger
	UnimplementedUserServiceServer
}

var UserServiceErrorCodes = map[error]ErrorCode{
	nil:                             ErrorCode_OK,
	apperrors.ErrCouldNotBuildQuery: ErrorCode_COULD_NOT_BUILD_QUERY,
	apperrors.ErrUserNotFound:       ErrorCode_USER_NOT_FOUND,
	apperrors.ErrWrongPassword:      ErrorCode_WRONG_PASSWORD,
	apperrors.ErrUserAlreadyExists:  ErrorCode_USER_ALREADY_EXISTS,
	apperrors.ErrUserNotCreated:     ErrorCode_USER_NOT_CREATED,
	apperrors.ErrUserNotUpdated:     ErrorCode_USER_NOT_UPDATED,
	apperrors.ErrUserNotDeleted:     ErrorCode_USER_NOT_DELETED,
	apperrors.ErrCouldNotGetUser:    ErrorCode_COULD_NOT_GET_USER,
	apperrors.ErrFailedToCreateFile: ErrorCode_FAILED_TO_CREATE_FILE,
	apperrors.ErrFailedToSaveFile:   ErrorCode_FAILED_TO_SAVE_FILE,
	apperrors.ErrFailedToDeleteFile: ErrorCode_FAILED_TO_DELETE_FILE,
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
func (us UserService) RegisterUser(ctx context.Context, info *AuthInfo) (*RegisterUserResponse, error) {
	funcName := "UserService.RegisterUser"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &RegisterUserResponse{}

	_, err := us.storage.GetWithLogin(sCtx, dto.UserLogin{Value: info.Email})

	if err == nil {
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}
	us.logger.Debug("User doesn't exist", funcName, nodeName)

	us.logger.Debug("Creating user", funcName, nodeName)
	user, err := us.storage.Create(sCtx, dto.SignupInfo{
		Email:        info.Email,
		PasswordHash: hashFromAuthInfo(info.Email, info.Password),
	})
	if err != nil {
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}

	response.Code = UserServiceErrorCodes[nil]
	response.Response = convertUser(user)

	return response, nil
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info *AuthInfo) (*CheckPasswordResponse, error) {
	funcName := "UserService.CheckPassword"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &CheckPasswordResponse{}

	user, err := us.storage.GetWithLogin(sCtx, dto.UserLogin{Value: info.Email})
	if err != nil {
		us.logger.Debug("User not found", funcName, nodeName)
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}
	us.logger.Debug("User found", funcName, nodeName)

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		response.Code = UserServiceErrorCodes[apperrors.ErrWrongPassword]
		response.Response = &User{}
		return response, nil
	}
	us.logger.Debug("Password match", funcName, nodeName)

	response.Code = UserServiceErrorCodes[nil]
	response.Response = convertUser(user)

	return response, nil
}

// GetWithID
// находит пользователя по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us UserService) GetWithID(ctx context.Context, id *UserID) (*GetWithIDResponse, error) {
	funcName := "UserService.GetWithID"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &GetWithIDResponse{}

	user, err := us.storage.GetWithID(sCtx, dto.UserID{Value: id.Value})
	if err != nil {
		us.logger.Debug("Failed to get user with error "+err.Error(), funcName, nodeName)
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}
	us.logger.Debug("Got user", funcName, nodeName)

	response.Code = UserServiceErrorCodes[nil]
	response.Response = convertUser(user)

	return response, nil
}

// UpdatePassword
// меняет пароль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdatePassword(ctx context.Context, info *PasswordChangeInfo) (*UpdatePasswordResponse, error) {
	funcName := "UserService.UpdatePassword"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &UpdatePasswordResponse{}

	oldLoginInfo, err := us.storage.GetLoginInfoWithID(sCtx, dto.UserID{Value: info.UserID})
	if err != nil {
		response.Code = UserServiceErrorCodes[err]
		return response, nil
	}
	us.logger.Debug("User found", funcName, nodeName)

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		response.Code = UserServiceErrorCodes[apperrors.ErrWrongPassword]
		return response, nil
	}
	us.logger.Debug("Old password verified", funcName, nodeName)

	err = us.storage.UpdatePassword(sCtx, dto.PasswordHashesInfo{
		UserID:          info.UserID,
		NewPasswordHash: hashFromAuthInfo(oldLoginInfo.Email, info.NewPassword),
	})
	response.Code = UserServiceErrorCodes[err]

	return response, nil
}

// UpdateProfile
// обновляет профиль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateProfile(ctx context.Context, info *UserProfileInfo) (*UpdateProfileResponse, error) {
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &UpdateProfileResponse{}

	err := us.storage.UpdateProfile(sCtx, dto.UserProfileInfo{
		UserID:      info.UserID,
		Name:        info.Name,
		Surname:     info.Surname,
		Description: info.Description,
	})
	response.Code = UserServiceErrorCodes[err]

	return response, nil
}

// UpdateProfile
// обновляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateAvatar(ctx context.Context, info *AvatarChangeInfo) (*UpdateAvatarResponse, error) {
	funcName := "UserService.UpdateAvatar"
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &UpdateAvatarResponse{}

	fileName := hashFromFileInfo(info.Filename, strconv.FormatUint(info.UserID, 10), info.Mimetype)

	fileLocation := "img/user_avatars/" + fileName + ".png"
	us.logger.Debug("Relative path: "+fileLocation, funcName, nodeName)
	avatarUrlInfo := dto.ImageUrlInfo{
		ID:  info.UserID,
		Url: fileLocation,
	}
	f, err := os.Create(fileLocation)
	if err != nil {
		us.logger.Debug("Failed to create file with error: "+err.Error(), funcName, nodeName)
		response.Code = UserServiceErrorCodes[apperrors.ErrFailedToCreateFile]
		response.Response = &UrlObj{}
		return response, nil
	}

	us.logger.Debug(fmt.Sprintf("Writing %v bytes", len(info.Avatar)), funcName, nodeName)
	_, err = f.Write(info.Avatar)
	if err != nil {
		response.Code = UserServiceErrorCodes[apperrors.ErrFailedToSaveFile]
		response.Response = &UrlObj{}
		return response, nil
	}

	defer f.Close()

	err = us.storage.UpdateAvatarUrl(sCtx, avatarUrlInfo)
	if err != nil {
		errDelete := os.Remove(fileLocation)
		if errDelete != nil {
			us.logger.Debug("Failed to remove file after unsuccessful update with error: "+err.Error(), funcName, nodeName)
			response.Code = UserServiceErrorCodes[apperrors.ErrFailedToDeleteFile]
			response.Response = &UrlObj{}
			return response, nil
		}
		response.Code = UserServiceErrorCodes[err]
		response.Response = &UrlObj{}
		return response, nil
	}

	response.Code = UserServiceErrorCodes[nil]
	response.Response = &UrlObj{Value: avatarUrlInfo.Url}

	return response, nil
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteUser(ctx context.Context, id *UserID) (*DeleteUserResponse, error) {
	sCtx := context.WithValue(ctx, dto.LoggerKey, us.logger)
	response := &DeleteUserResponse{}

	err := us.storage.Delete(sCtx, dto.UserID{Value: id.Value})
	response.Code = UserServiceErrorCodes[err]

	return response, nil
}

// TODO salt
func hashFromAuthInfo(email string, password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email + password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func hashFromFileInfo(filename string, id string, mimetype string) string {
	hasher := sha256.New()
	hasher.Write([]byte(filename + id + mimetype))
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
