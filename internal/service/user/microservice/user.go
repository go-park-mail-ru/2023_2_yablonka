package microservice

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	microservice "server/microservices/user/user"

	logger "server/internal/logging"

	"google.golang.org/grpc"
)

type UserService struct {
	storage storage.IUserStorage
	client  microservice.UserServiceClient
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

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	user, err := us.client.RegisterUser(ctx, &microservice.AuthInfo{
		Email:    info.Email,
		Password: info.Password,
	})
	if handledErr := apperrors.HandleGRPCError(err); handledErr != nil {
		return &entities.User{}, handledErr
	}
	logger.Debug("Info received", funcName, nodeName)

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

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	user, err := us.client.CheckPassword(ctx, &microservice.AuthInfo{
		Email:    info.Email,
		Password: info.Password,
	})
	if handledErr := apperrors.HandleGRPCError(err); handledErr != nil {
		return &entities.User{}, handledErr
	}
	logger.Debug("Info received", funcName, nodeName)

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

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	user, err := us.client.GetWithID(ctx, &microservice.UserID{Value: id.Value})
	if handledErr := apperrors.HandleGRPCError(err); handledErr != nil {
		return &entities.User{}, handledErr
	}
	logger.Debug("Info received", funcName, nodeName)

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

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	_, err := us.client.UpdatePassword(ctx, &microservice.PasswordChangeInfo{
		UserID:      info.UserID,
		OldPassword: info.OldPassword,
		NewPassword: info.NewPassword,
	})
	logger.Debug("Response received", funcName, nodeName)

	return apperrors.HandleGRPCError(err)
}

// UpdateProfile
// обновляет профиль пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateProfile(ctx context.Context, info dto.UserProfileInfo) error {
	funcName := "UserService.UpdateProfile"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	_, err := us.client.UpdateProfile(ctx, &microservice.UserProfileInfo{
		UserID:      info.UserID,
		Name:        info.Name,
		Surname:     info.Surname,
		Description: info.Description,
	})
	logger.Debug("Response received", funcName, nodeName)

	return apperrors.HandleGRPCError(err)
}

// UpdateProfile
// обновляет аватарку пользователя
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) UpdateAvatar(ctx context.Context, info dto.AvatarChangeInfo) (*dto.UrlObj, error) {
	funcName := "UserService.UpdateAvatar"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	urlObj, err := us.client.UpdateAvatar(ctx, &microservice.AvatarChangeInfo{
		UserID:  info.UserID,
		Avatar:  info.Avatar,
		BaseURL: ctx.Value(dto.BaseURLKey).(string),
	})
	logger.Debug("Response received", funcName, nodeName)
	if handledErr := apperrors.HandleGRPCError(err); handledErr != nil {
		return &dto.UrlObj{}, handledErr
	}

	return &dto.UrlObj{Value: urlObj.Value}, nil
}

// DeleteUser
// удаляет данного пользователя по id
// или возвращает ошибку apperrors.ErrUserNotFound (409)
func (us UserService) DeleteUser(ctx context.Context, id dto.UserID) error {
	funcName := "UserService.UpdateProfile"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Debug("Contacting GRPC server", funcName, nodeName)
	_, err := us.client.DeleteUser(ctx, &microservice.UserID{Value: id.Value})
	logger.Debug("Response received", funcName, nodeName)

	return apperrors.HandleGRPCError(err)
}
