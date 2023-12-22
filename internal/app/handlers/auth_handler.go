package handlers

import (
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/mailru/easyjson"
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
	errorMessage := "Logging user in failed with error: "
	failBorder := "---------------------------------- User Login FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- User Login ----------------------------------")

	var authInfo dto.AuthInfo
	err := easyjson.UnmarshalFromReader(r.Body, &authInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON decoded", requestID.String(), funcName, nodeName)

	_, err = govalidator.ValidateStruct(authInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("Login info validated", requestID.String(), funcName, nodeName)

	user, err := ah.us.CheckPassword(rCtx, authInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Password checked", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("Session created", requestID.String(), funcName, nodeName)

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	logger.DebugFmt("CSRF token response header set", requestID.String(), funcName, nodeName)

	authCookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}
	http.SetCookie(w, authCookie)
	logger.DebugFmt("Cookie set", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.DebugFmt("Response written", requestID.String(), funcName, nodeName)
	logger.Info("---------------------------------- User Login SUCCESS ----------------------------------")
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
	errorMessage := "Signing user up failed with error: "
	failBorder := "---------------------------------- User SignUp FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- User Login ----------------------------------")

	var signup dto.AuthInfo
	err := easyjson.UnmarshalFromReader(r.Body, &signup)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON decoded", requestID.String(), funcName, nodeName)

	_, err = govalidator.ValidateStruct(signup)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct validated", requestID.String(), funcName, nodeName)

	user, err := ah.us.RegisterUser(rCtx, signup)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User registered", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("User authorized", requestID.String(), funcName, nodeName)

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	logger.DebugFmt("JSON decoded", requestID.String(), funcName, nodeName)

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}
	http.SetCookie(w, cookie)
	logger.DebugFmt("Cookie set", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- User SignUp SUCCESS ----------------------------------")
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
	errorMessage := "Logging user out failed with error: "
	failBorder := "---------------------------------- User LogOut FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- User Logout ----------------------------------")

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("Cookie found", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("Session verified", requestID.String(), funcName, nodeName)

	err = ah.as.LogOut(rCtx, token)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Session deleted", requestID.String(), funcName, nodeName)

	cookie = &http.Cookie{
		Name:     "tabula_user",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Time{},
		Path:     "/api/v2/",
	}
	http.SetCookie(w, cookie)
	logger.DebugFmt("Empty cookie set", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.DebugFmt("Response written", requestID.String(), funcName, nodeName)
	logger.Info("---------------------------------- User Logout SUCCESS ----------------------------------")
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
	errorMessage := "Logging user out failed with error: "
	failBorder := "---------------------------------- User auth verification FAIL ----------------------------------"

	rCtx := r.Context()
	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Verifying user authorization ----------------------------------")

	w.Header().Set("X-Csrf-Token", "")
	logger.DebugFmt("Default X-Csrf-Token set", requestID.String(), funcName, nodeName)

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("Cookie found", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("Session verified", requestID.String(), funcName, nodeName)

	user, err := ah.us.GetWithID(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Got user object", requestID.String(), funcName, nodeName)

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	logger.DebugFmt("CSRF set up", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- User SignUp SUCCESS ----------------------------------")
}
