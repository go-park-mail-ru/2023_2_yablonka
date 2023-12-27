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

	"github.com/google/uuid"
)

const nodeName = "microservice"

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
func (us UserService) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	funcName := "UserService.RegisterUser"
	response := &RegisterUserResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	_, err := us.storage.GetWithLogin(sCtx, dto.UserLogin{Value: info.Email})

	if err == nil {
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}
	us.logger.DebugFmt("User doesn't exist", requestID.String(), funcName, nodeName)

	us.logger.DebugFmt("Creating user", requestID.String(), funcName, nodeName)
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
func (us UserService) CheckPassword(ctx context.Context, request *CheckPasswordRequest) (*CheckPasswordResponse, error) {
	funcName := "UserService.CheckPassword"
	response := &CheckPasswordResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	user, err := us.storage.GetWithLogin(sCtx, dto.UserLogin{Value: info.Email})
	if err != nil {
		us.logger.DebugFmt("User not found", requestID.String(), funcName, nodeName)
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}
	us.logger.DebugFmt("User found", requestID.String(), funcName, nodeName)

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		response.Code = UserServiceErrorCodes[apperrors.ErrWrongPassword]
		response.Response = &User{}
		return response, nil
	}
	us.logger.DebugFmt("Password match", requestID.String(), funcName, nodeName)

	response.Code = UserServiceErrorCodes[nil]
	response.Response = convertUser(user)

	return response, nil
}

// GetWithID
// находит пользователя по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us UserService) GetWithID(ctx context.Context, request *GetWithIDRequest) (*GetWithIDResponse, error) {
	funcName := "UserService.GetWithID"
	response := &GetWithIDResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	id := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	user, err := us.storage.GetWithID(sCtx, dto.UserID{Value: id.Value})
	if err != nil {
		us.logger.DebugFmt("Failed to get user with error "+err.Error(), requestID.String(), funcName, nodeName)
		response.Code = UserServiceErrorCodes[err]
		response.Response = &User{}
		return response, nil
	}
	us.logger.DebugFmt("Got user", requestID.String(), funcName, nodeName)

	response.Code = UserServiceErrorCodes[nil]
	response.Response = convertUser(user)

	return response, nil
}

// UpdatePassword
// меняет пароль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdatePassword(ctx context.Context, request *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	funcName := "UserService.UpdatePassword"
	response := &UpdatePasswordResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	oldLoginInfo, err := us.storage.GetLoginInfoWithID(sCtx, dto.UserID{Value: info.UserID})
	if err != nil {
		response.Code = UserServiceErrorCodes[err]
		return response, nil
	}
	us.logger.DebugFmt("User found", requestID.String(), funcName, nodeName)

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		response.Code = UserServiceErrorCodes[apperrors.ErrWrongPassword]
		return response, nil
	}
	us.logger.DebugFmt("Old password verified", requestID.String(), funcName, nodeName)

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
func (us UserService) UpdateProfile(ctx context.Context, request *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	response := &UpdateProfileResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	err := us.storage.UpdateProfile(sCtx, dto.UserProfileInfo{
		UserID:      info.UserID,
		Name:        info.Name,
		Surname:     info.Surname,
		Description: info.Description,
	})
	response.Code = UserServiceErrorCodes[err]

	return response, nil
}

// UpdateAvatar
// обновляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateAvatar(ctx context.Context, request *UpdateAvatarRequest) (*UpdateAvatarResponse, error) {
	funcName := "UserService.UpdateAvatar"
	response := &UpdateAvatarResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	fileName := hashFromFileInfo(info.Filename, strconv.FormatUint(info.UserID, 10), info.Mimetype)

	fileLocation := "img/user_avatars/" + fileName + ".png"
	cwd, _ := os.Getwd()
	us.logger.DebugFmt("Relative path: "+fileLocation, requestID.String(), funcName, nodeName)
	us.logger.DebugFmt("CWD: "+cwd, requestID.String(), funcName, nodeName)
	avatarUrlInfo := dto.UserImageUrlInfo{
		ID:  info.UserID,
		Url: fileLocation,
	}
	us.logger.DebugFmt("Full URL: "+avatarUrlInfo.Url, requestID.String(), funcName, nodeName)
	f, err := os.Create(fileLocation)
	if err != nil {
		us.logger.DebugFmt("Failed to create file with error: "+err.Error(), requestID.String(), funcName, nodeName)
		response.Code = UserServiceErrorCodes[apperrors.ErrFailedToCreateFile]
		response.Response = &UrlObj{}
		return response, nil
	}

	us.logger.DebugFmt(fmt.Sprintf("Writing %v bytes", len(info.Avatar)), requestID.String(), funcName, nodeName)
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
			us.logger.DebugFmt("Failed to remove file after unsuccessful update with error: "+err.Error(), requestID.String(), funcName, nodeName)
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

// DeleteAvatar
// удаляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteAvatar(ctx context.Context, request *DeleteAvatarRequest) (*DeleteAvatarResponse, error) {
	funcName := "UserService.DeleteAvatar"
	requestID, _ := uuid.Parse(request.RequestID)
	info := request.Value
	response := &DeleteAvatarResponse{}

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

	err := os.Remove(info.Filename)
	if err != nil {
		us.logger.DebugFmt("Failed to remove file after unsuccessful update with error: "+err.Error(), request.RequestID, funcName, nodeName)
		response.Code = UserServiceErrorCodes[apperrors.ErrFailedToDeleteFile]
		response.Response = &UrlObj{}
		return response, nil
	}

	url := dto.UserImageUrlInfo{
		ID:  info.UserID,
		Url: "img/user_avatars/avatar.jpg",
	}

	err = us.storage.UpdateAvatarUrl(sCtx, url)
	if err != nil {
		response.Code = UserServiceErrorCodes[err]
		response.Response = &UrlObj{}
		return response, nil
	}

	response.Code = UserServiceErrorCodes[nil]
	response.Response = &UrlObj{Value: url.Url}

	return response, nil
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	response := &DeleteUserResponse{}
	requestID, _ := uuid.Parse(request.RequestID)
	id := request.Value

	sCtx := context.WithValue(
		context.WithValue(ctx, dto.LoggerKey, us.logger),
		dto.RequestIDKey, requestID,
	)

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
