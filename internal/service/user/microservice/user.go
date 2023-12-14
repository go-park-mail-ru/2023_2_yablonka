package microservice

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	microservice "server/microservices/user/user"

	logger "server/internal/logging"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type UserService struct {
	storage storage.IUserStorage
	client  microservice.UserServiceClient
}

var UserServiceErrors = map[microservice.ErrorCode]error{
	microservice.ErrorCode_OK:                    nil,
	microservice.ErrorCode_COULD_NOT_BUILD_QUERY: apperrors.ErrCouldNotBuildQuery,
	microservice.ErrorCode_USER_NOT_FOUND:        apperrors.ErrUserNotFound,
	microservice.ErrorCode_WRONG_PASSWORD:        apperrors.ErrWrongPassword,
	microservice.ErrorCode_USER_ALREADY_EXISTS:   apperrors.ErrUserAlreadyExists,
	microservice.ErrorCode_USER_NOT_CREATED:      apperrors.ErrUserNotCreated,
	microservice.ErrorCode_USER_NOT_UPDATED:      apperrors.ErrUserNotUpdated,
	microservice.ErrorCode_USER_NOT_DELETED:      apperrors.ErrUserNotDeleted,
	microservice.ErrorCode_COULD_NOT_GET_USER:    apperrors.ErrCouldNotGetUser,
	microservice.ErrorCode_FAILED_TO_CREATE_FILE: apperrors.ErrFailedToCreateFile,
	microservice.ErrorCode_FAILED_TO_SAVE_FILE:   apperrors.ErrFailedToSaveFile,
	microservice.ErrorCode_FAILED_TO_DELETE_FILE: apperrors.ErrFailedToDeleteFile,
}

const nodeName = "service"

// NewUserService
// возвращает UserService с инициализированным хранилищем пользователей
func NewUserService(storage storage.IUserStorage, conn *grpc.ClientConn) *UserService {
	client := microservice.NewUserServiceClient(conn)
	return &UserService{
		storage: storage,
		client:  client,
	}
}

// RegisterUser
// создает нового пользователя по данным
// или возвращает ошибку apperrors.ErrUserAlreadyExists (409)
func (us UserService) RegisterUser(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	funcName := "UserService.RegisterUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.RegisterUserRequest{
		RequestID: requestID.String(),
		Value: &microservice.AuthInfo{
			Email:    info.Email,
			Password: info.Password,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.RegisterUser(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	if serverResponse.Code != microservice.ErrorCode_OK {
		return &entities.User{}, UserServiceErrors[serverResponse.Code]
	}

	user := serverResponse.Response

	return &entities.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         &user.Name,
		Surname:      &user.Surname,
		Description:  &user.Description,
		AvatarURL:    &user.AvatarURL,
	}, nil
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	funcName := "UserService.CheckPassword"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.CheckPasswordRequest{
		RequestID: requestID.String(),
		Value: &microservice.AuthInfo{
			Email:    info.Email,
			Password: info.Password,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.CheckPassword(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)
	if serverResponse.Code != microservice.ErrorCode_OK {
		return &entities.User{}, UserServiceErrors[serverResponse.Code]
	}

	user := serverResponse.Response

	return &entities.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         &user.Name,
		Surname:      &user.Surname,
		Description:  &user.Description,
		AvatarURL:    &user.AvatarURL,
	}, nil
}

// GetWithID
// находит пользователя по его id
// или возвращает ошибку apperrors.ErrUserNotFound (401)
func (us UserService) GetWithID(ctx context.Context, id dto.UserID) (*entities.User, error) {
	funcName := "UserService.GetWithID"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.GetWithIDRequest{
		RequestID: requestID.String(),
		Value:     &microservice.UserID{Value: id.Value},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.GetWithID(ctx, grpcRequest)
	if serverResponse.Code != microservice.ErrorCode_OK {
		return &entities.User{}, UserServiceErrors[serverResponse.Code]
	}
	logger.DebugFmt("Info received", requestID.String(), funcName, nodeName)

	user := serverResponse.Response

	return &entities.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         &user.Name,
		Surname:      &user.Surname,
		Description:  &user.Description,
		AvatarURL:    &user.AvatarURL,
	}, nil
}

// UpdatePassword
// меняет пароль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdatePassword(ctx context.Context, info dto.PasswordChangeInfo) error {
	funcName := "UserService.UpdatePassword"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.UpdatePasswordRequest{
		RequestID: requestID.String(),
		Value: &microservice.PasswordChangeInfo{
			UserID:      info.UserID,
			OldPassword: info.OldPassword,
			NewPassword: info.NewPassword,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.UpdatePassword(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return UserServiceErrors[serverResponse.Code]
}

// UpdateProfile
// обновляет профиль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateProfile(ctx context.Context, info dto.UserProfileInfo) error {
	funcName := "UserService.UpdateProfile"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.UpdateProfileRequest{
		RequestID: requestID.String(),
		Value: &microservice.UserProfileInfo{
			UserID:      info.UserID,
			Name:        info.Name,
			Surname:     info.Surname,
			Description: info.Description,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.UpdateProfile(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return UserServiceErrors[serverResponse.Code]
}

// UpdateProfile
// обновляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateAvatar(ctx context.Context, info dto.AvatarChangeInfo) (*dto.UrlObj, error) {
	funcName := "UserService.UpdateAvatar"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.UpdateAvatarRequest{
		RequestID: requestID.String(),
		Value: &microservice.AvatarChangeInfo{
			UserID:   info.UserID,
			Avatar:   info.Avatar,
			Filename: info.Filename,
			Mimetype: info.Mimetype,
		},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.UpdateAvatar(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)
	if serverResponse.Code != microservice.ErrorCode_OK {
		return &dto.UrlObj{}, UserServiceErrors[serverResponse.Code]
	}

	urlObj := serverResponse.Response

	return &dto.UrlObj{Value: urlObj.Value}, nil
}

// Delete
// удаляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteAvatar(ctx context.Context, info dto.AvatarRemovalInfo) (*dto.UrlObj, error) {
	funcName := "UserService.DeleteAvatar"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.DebugFmt("Contacting GRPC server", funcName, nodeName)
	serverResponse, _ := us.client.DeleteAvatar(ctx, &microservice.AvatarRemovalInfo{
		UserID:   info.UserID,
		Filename: info.AvatarUrl,
	})
	logger.DebugFmt("Response received", funcName, nodeName)

	if serverResponse.Code != microservice.ErrorCode_OK {
		return &dto.UrlObj{}, UserServiceErrors[serverResponse.Code]
	}

	return &dto.UrlObj{
		Value: serverResponse.Response.Value,
	}, nil
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteUser(ctx context.Context, id dto.UserID) error {
	funcName := "UserService.UpdateProfile"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	grpcRequest := &microservice.DeleteUserRequest{
		RequestID: requestID.String(),
		Value:     &microservice.UserID{Value: id.Value},
	}

	logger.DebugFmt("Contacting GRPC server", requestID.String(), funcName, nodeName)
	serverResponse, _ := us.client.DeleteUser(ctx, grpcRequest)
	logger.DebugFmt("Response received", requestID.String(), funcName, nodeName)

	return UserServiceErrors[serverResponse.Code]
}
