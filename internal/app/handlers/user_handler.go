package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	apperrors "server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	dto "server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service/user"

	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	us user.IUserService
}

// @Summary Поменять пароль
// @Description Получает старый и новый пароли
// @Tags user
//
// @Accept  json
// @Produce  json
//
// @Param passwords body dto.PasswordChangeInfo true "Старый и новый пароли пользователя"
//
// @Success 200  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/edit/change_password/ [post]
func (uh UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChangePassword"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Changing user password")

	var passwords dto.PasswordChangeInfo

	err := json.NewDecoder(r.Body).Decode(&passwords)
	if err != nil {
		logger.Error("Password change failed")
		handlerDebugLog(logger, funcName, "Changing user password failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON decoded")

	_, err = govalidator.ValidateStruct(passwords)
	if err != nil {
		logger.Error("Password change failed")
		handlerDebugLog(logger, funcName, "Changing user password failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Request data validated")

	userID := rCtx.Value(dto.UserObjKey).(*entities.User).ID
	passwords.UserID = userID

	handlerDebugLog(logger, funcName, fmt.Sprintf("Updating password for user ID %d", passwords.UserID))
	err = uh.us.UpdatePassword(rCtx, passwords)
	if err != nil {
		logger.Error("Password change failed")
		handlerDebugLog(logger, funcName, "Changing user password failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Password updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Password change failed")
		handlerDebugLog(logger, funcName, "Changing user password failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Password change failed")
		handlerDebugLog(logger, funcName, "Changing user password failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished changing user password")
}

// @Summary Поменять данные профиля
// @Description В ответ ничего не шлёт
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param newProfileInfo body dto.UserProfileInfo true "Имя, фамилия, описание пользователя"
//
// @Success 200  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/edit/ [post]
func (uh UserHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChangeProfile"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Changing user info")

	var newProfileInfo dto.UserProfileInfo

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error("User info change failed")
		handlerDebugLog(logger, funcName, "Changing user info failed -- no user passed in context")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	err := json.NewDecoder(r.Body).Decode(&newProfileInfo)
	if err != nil {
		logger.Error("User info change failed")
		handlerDebugLog(logger, funcName, "Changing user info failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON decoded")

	_, err = govalidator.ValidateStruct(newProfileInfo)
	if err != nil {
		logger.Error("User info change failed")
		handlerDebugLog(logger, funcName, "Changing user info failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Request data validated")

	newProfileInfo.UserID = user.ID

	handlerDebugLog(logger, funcName, fmt.Sprintf("Updating profile info for user ID %d", newProfileInfo.UserID))
	err = uh.us.UpdateProfile(rCtx, newProfileInfo)
	if err != nil {
		logger.Error("User info change failed")
		handlerDebugLog(logger, funcName, "Changing user info failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User info updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("User info change failed")
		handlerDebugLog(logger, funcName, "Changing user info failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response generated")

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished changing user info")
}

// @Summary Поменять аватарку
// @Description В ответ шлёт ссылку на файл
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param avatarChangeInfo body dto.AvatarChangeInfo true "id пользователя, изображение"
//
// @Success 200  {object}  doc_structs.AvatarUploadResponse "Объект пользователя"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/edit/change_avatar/ [post]
func (uh UserHandler) ChangeAvatar(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChangeAvatar"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Changing user avatar")

	var avatarChangeInfo dto.AvatarChangeInfo

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error("User avatar change failed")
		handlerDebugLog(logger, funcName, "Changing user avatar failed -- no user passed in context")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	err := json.NewDecoder(r.Body).Decode(&avatarChangeInfo)
	if err != nil {
		logger.Error("User avatar change failed")
		handlerDebugLog(logger, funcName, "Changing user avatar failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	}
	handlerDebugLog(logger, funcName, "JSON decoded")

	_, err = govalidator.ValidateStruct(avatarChangeInfo)
	if err != nil {
		logger.Error("User avatar change failed")
		handlerDebugLog(logger, funcName, "Changing user avatar failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Request data validated")

	avatarChangeInfo.UserID = user.ID

	handlerDebugLog(logger, funcName, fmt.Sprintf("Updating avatar for user ID %d", avatarChangeInfo.UserID))
	url, err := uh.us.UpdateAvatar(rCtx, avatarChangeInfo)
	if err != nil {
		logger.Error("User avatar change failed")
		handlerDebugLog(logger, funcName, "Changing user avatar failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User avatar updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"avatar_url": url,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("User avatar change failed")
		handlerDebugLog(logger, funcName, "Changing user avatar failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("User avatar change failed")
		handlerDebugLog(logger, funcName, "Changing user avatar failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished updating user profile")
}
