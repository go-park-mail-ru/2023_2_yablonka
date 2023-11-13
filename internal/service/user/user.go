package service

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
	log.Println("Service -- Registering")
	_, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})

	if err == nil {
		log.Println("User exists")
		return nil, apperrors.ErrUserAlreadyExists
	}
	log.Println("User doesnt exists")

	return us.storage.Create(ctx, dto.SignupInfo{
		Email:        info.Email,
		PasswordHash: hashFromAuthInfo(info.Email, info.Password),
	})
}

// CheckPassword
// проверяет пароль пользователя по почте
// или возвращает ошибки apperrors.ErrUserNotFound (401), apperrors.ErrWrongPassword (401)
func (us UserService) CheckPassword(ctx context.Context, info dto.AuthInfo) (*entities.User, error) {
	user, err := us.storage.GetWithLogin(ctx, dto.UserLogin{Value: info.Email})
	if err != nil {
		return nil, err
	}

	if user.PasswordHash != hashFromAuthInfo(info.Email, info.Password) {
		return nil, apperrors.ErrWrongPassword
	}

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
	oldLoginInfo, err := us.storage.GetLoginInfoWithID(ctx, dto.UserID{Value: info.UserID})
	if err == nil {
		return apperrors.ErrUserNotFound
	}

	if oldLoginInfo.PasswordHash != hashFromAuthInfo(oldLoginInfo.Email, info.OldPassword) {
		return apperrors.ErrWrongPassword
	}

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
	url := "user_avatars/" + strconv.FormatUint(info.UserID, 10) + ".png"
	fileLocation := "img/" + url
	log.Println("Service -- File location:", fileLocation)
	avatarUrlInfo := dto.ImageUrlInfo{
		ID:  info.UserID,
		Url: url,
	}
	log.Println("Service -- Full url to file", avatarUrlInfo.Url)
	f, err := os.Create("./" + fileLocation)
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
		errDelete := os.Remove("./" + fileLocation)
		for errDelete != nil {
			log.Println("Service -- Failed to remove file after unsuccessful update with error", err)
			errDelete = os.Remove("./" + fileLocation)
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
