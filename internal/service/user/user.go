package user

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	embedded "server/internal/service/user/embedded"
	micro "server/internal/service/user/microservice"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса пользователей
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IUserService interface {
	// RegisterUser
	// создает нового пользователя по данным
	// или возвращает ошибки ...
	RegisterUser(context.Context, dto.AuthInfo) (*entities.User, error)
	// CheckPassword
	// проверяет пароль пользователя по почте
	// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
	CheckPassword(context.Context, dto.AuthInfo) (*entities.User, error)
	// ReadWithID
	// находит пользователя по его id
	// или возвращает ошибки ...
	GetWithID(context.Context, dto.UserID) (*entities.User, error)
	// UpdatePassword
	// обновляет пароль пользователя
	// или возвращает ошибки ...
	UpdatePassword(context.Context, dto.PasswordChangeInfo) error
	// UpdateProfile
	// обновляет профиль пользователя
	// или возвращает ошибки ...
	UpdateProfile(context.Context, dto.UserProfileInfo) error
	// UpdateAvatar
	// обновляет аватарку пользователя
	// или возвращает ошибки ...
	UpdateAvatar(context.Context, dto.AvatarChangeInfo) (*dto.UrlObj, error)
	// Delete
	// удаляет данного пользователя по id
	// или возвращает ошибки ...
	DeleteUser(context.Context, dto.UserID) error
}

func NewEmbeddedUserService(userStorage storage.IUserStorage) *embedded.UserService {
	return embedded.NewUserService(userStorage)
}

// TODO: User microservice
func NewMicroUserService(userStorage storage.IUserStorage, connection *grpc.ClientConn) *micro.UserService {
	return micro.NewUserService(userStorage)
}
