package handlers

import (
	"encoding/json"
	"net/http"

	apperrors "server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	dto "server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"

	logger "server/internal/logging"

	"github.com/asaskevich/govalidator"
)

type UserHandler struct {
	us service.IUserService
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
	funcName := "UserHandler.ChangePassword"
	nodeName := "handler"
	errorMessage := "Changing user's password failed with error: "
	failBorder := "---------------------------------- Changing user's password FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Changing user's password ----------------------------------")

	var passwords dto.PasswordChangeInfo
	err := json.NewDecoder(r.Body).Decode(&passwords)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON parsed", funcName, nodeName)

	_, err = govalidator.ValidateStruct(passwords)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("Request data validated", funcName, nodeName)

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", funcName, nodeName)

	passwords.UserID = user.ID
	err = uh.us.UpdatePassword(rCtx, passwords)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Password updated", funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", funcName, nodeName)

	logger.Info("---------------------------------- Changing user's password SUCCESS ----------------------------------")
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
	funcName := "UserHandler.ChangeProfile"
	nodeName := "handler"
	errorMessage := "Changing user's profile failed with error: "
	failBorder := "---------------------------------- Changing user's profile FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Changing user's profile ----------------------------------")

	var newProfileInfo dto.UserProfileInfo
	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "No user object found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", funcName, nodeName)

	err := json.NewDecoder(r.Body).Decode(&newProfileInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON decoded", funcName, nodeName)

	_, err = govalidator.ValidateStruct(newProfileInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("Request data validated", funcName, nodeName)

	newProfileInfo.UserID = user.ID
	err = uh.us.UpdateProfile(rCtx, newProfileInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User info updated", funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", funcName, nodeName)

	logger.Info("---------------------------------- Changing user's profile SUCCESS ----------------------------------")
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
	funcName := "UserHandler.ChangeAvatar"
	nodeName := "handler"
	errorMessage := "Changing user's avatar failed with error: "
	failBorder := "---------------------------------- Changing user's avatar FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Changing user's avatar ----------------------------------")

	var avatarChangeInfo dto.AvatarChangeInfo
	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "No user object found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", funcName, nodeName)

	err := json.NewDecoder(r.Body).Decode(&avatarChangeInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON parsed", funcName, nodeName)

	_, err = govalidator.ValidateStruct(avatarChangeInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("Request data validated", funcName, nodeName)

	avatarChangeInfo.UserID = user.ID
	url, err := uh.us.UpdateAvatar(rCtx, avatarChangeInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User avatar updated", funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"avatar_url": url,
		},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", funcName, nodeName)

	logger.Info("---------------------------------- Changing user's avatar SUCCESS ----------------------------------")
}
