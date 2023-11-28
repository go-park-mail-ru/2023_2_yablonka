package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service"
	"time"

	"github.com/asaskevich/govalidator"
)

type AuthHandler struct {
	as service.IAuthService
	us service.IUserService
	cs service.ICSRFService
}

func (ah AuthHandler) GetAuthService() service.IAuthService {
	return ah.as
}

func (ah AuthHandler) GetUserService() service.IUserService {
	return ah.us
}

func (ah AuthHandler) GetCSRFService() service.ICSRFService {
	return ah.cs
}

// @Summary Войти в систему
// @Description Для этого использует сессии
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param authData body dto.AuthInfo true "Эл. почта и логин пользователя"
//
// @Success 200  {object}  doc_structs.UserResponse "Объект пользователя"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/login/ [post]
func (ah AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "LogIn"
	nodeName := "handler"
	errorMessage := "Logging user in failed with error: "
	failBorder := "----------------- User Login FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- User Login -----------------")

	var authInfo dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("JSON decoded", funcName, nodeName)

	_, err = govalidator.ValidateStruct(authInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("Login info validated", funcName, nodeName)

	user, err := ah.us.CheckPassword(rCtx, authInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Password checked", funcName, nodeName)

	userID := dto.UserID{
		Value: user.ID,
	}
	session, err := ah.as.AuthUser(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Session created", funcName, nodeName)

	authCookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}
	http.SetCookie(w, authCookie)
	logger.Debug("Cookie set", funcName, nodeName)

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	logger.Debug("CSRF token response header set", funcName, nodeName)

	publicUserInfo := dto.UserPublicInfo{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Description: user.Description,
		AvatarURL:   user.AvatarURL,
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": publicUserInfo,
		},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.Debug("response written", funcName, nodeName)

	logger.Debug("Response written", funcName, nodeName)
	logger.Info("----------------- User Login SUCCESS -----------------")
}

// @Summary Зарегистрировать нового пользователя
// @Description Также вводит пользователя в систему
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param signup body dto.AuthInfo true "Базовые данные пользователя"
//
// @Success 200  {object}  doc_structs.UserResponse "Объект пользователя"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 409  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/signup/ [post]
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "SignUp"
	nodeName := "handler"
	errorMessage := "Signing user up failed with error: "
	failBorder := "----------------- User SignUp FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- User Login -----------------")

	var signup dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("JSON decoded", funcName, nodeName)

	_, err = govalidator.ValidateStruct(signup)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("request struct validated", funcName, nodeName)

	user, err := ah.us.RegisterUser(rCtx, signup)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("User registered", funcName, nodeName)

	userID := dto.UserID{
		Value: user.ID,
	}
	session, err := ah.as.AuthUser(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("User authorized", funcName, nodeName)

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}
	http.SetCookie(w, cookie)
	logger.Debug("Cookie set", funcName, nodeName)

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	logger.Debug("JSON decoded", funcName, nodeName)

	publicUserInfo := dto.UserPublicInfo{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Description: user.Description,
		AvatarURL:   user.AvatarURL,
	}
	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": publicUserInfo,
		},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.Debug("response written", funcName, nodeName)

	logger.Info("----------------- User SignUp SUCCESS -----------------")
}

// @Summary Выйти из системы
// @Description Удаляет текущую сессию пользователя. Может сделать только авторизированный пользователь. Текущая сессия определяется по cookie "tabula_user", в которой лежит строка-токен.
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Success 204  {string} string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/logout/ [delete]
func (ah AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "LogOut"
	nodeName := "handler"
	errorMessage := "Logging user out failed with error: "
	failBorder := "----------------- User LogOut FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- User Logout -----------------")

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.Debug("Cookie found", funcName, nodeName)

	token := dto.SessionToken{
		ID: cookie.Value,
	}

	_, err = ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Session verified", funcName, nodeName)

	err = ah.as.LogOut(rCtx, token)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Session deleted", funcName, nodeName)

	cookie = &http.Cookie{
		Name:     "tabula_user",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Time{},
		Path:     "/api/v2/",
	}
	http.SetCookie(w, cookie)
	logger.Debug("Empty cookie set", funcName, nodeName)

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
	logger.Debug("response written", funcName, nodeName)

	logger.Debug("Response written", funcName, nodeName)
	logger.Info("----------------- User Logout SUCCESS -----------------")
}

// @Summary Подтвердить вход
// @Description Узнать существует ли сессия текущего пользователя. Сессия определяется по cookie "tabula_user", в которой лежит строка-токен.
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.UserResponse "Объект пользователя"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/verify [get]
func (ah AuthHandler) VerifyAuthEndpoint(w http.ResponseWriter, r *http.Request) {
	funcName := "LogOut"
	nodeName := "handler"
	errorMessage := "Logging user out failed with error: "
	failBorder := "----------------- User auth verification FAIL -----------------"

	rCtx := r.Context()
	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- Verifying user authorization -----------------")

	w.Header().Set("X-Csrf-Token", "")
	logger.Debug("Default X-Csrf-Token set", funcName, nodeName)

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.Debug("Cookie found", funcName, nodeName)

	token := dto.SessionToken{
		ID: cookie.Value,
	}
	userID, err := ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Session verified", funcName, nodeName)

	user, err := ah.us.GetWithID(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Got user object", funcName, nodeName)

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	logger.Debug("CSRF set up", funcName, nodeName)

	publicUserInfo := dto.UserPublicInfo{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Description: user.Description,
		AvatarURL:   user.AvatarURL,
	}
	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": publicUserInfo,
		},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.Debug("response written", funcName, nodeName)

	logger.Info("----------------- User SignUp SUCCESS -----------------")
}
