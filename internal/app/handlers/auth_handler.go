package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
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

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Logging user in")

	var authInfo dto.AuthInfo

	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON decoded")

	_, err = govalidator.ValidateStruct(authInfo)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Login info validated")

	user, err := ah.us.CheckPassword(rCtx, authInfo)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Password checked")

	userID := dto.UserID{
		Value: user.ID,
	}

	handlerDebugLog(logger, funcName, fmt.Sprintf("Creating session for user ID %d", userID.Value))
	session, err := ah.as.AuthUser(rCtx, userID)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Session created")

	authCookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}

	http.SetCookie(w, authCookie)
	handlerDebugLog(logger, funcName, "Cookie set")

	handlerDebugLog(logger, funcName, fmt.Sprintf("Setting up CSRF for user ID %d", userID.Value))
	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "CSRF set up")

	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	handlerDebugLog(logger, funcName, "CSRF token response header set")

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

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Login failed")
		handlerDebugLog(logger, funcName, "Logging user in failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Logged user in")
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

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Signing new user up")

	var signup dto.AuthInfo

	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	_, err = govalidator.ValidateStruct(signup)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Signup data validated")

	user, err := ah.us.RegisterUser(rCtx, signup)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, "SingUp", "User registered")

	userID := dto.UserID{
		Value: user.ID,
	}

	handlerDebugLog(logger, funcName, fmt.Sprintf("Creating session for user ID %d", userID.Value))
	session, err := ah.as.AuthUser(rCtx, userID)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Session created")

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}

	http.SetCookie(w, cookie)
	handlerDebugLog(logger, funcName, "Cookie set")

	handlerDebugLog(logger, funcName, fmt.Sprintf("Setting up CSRF for user ID %d", userID.Value))
	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "CSRF set up")

	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	handlerDebugLog(logger, funcName, "CSRF token response header set")

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

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Signup failed")
		handlerDebugLog(logger, funcName, "Signing user up failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished signing new user up")
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

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Logging user out")

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		logger.Error("Logout failed")
		handlerDebugLog(logger, funcName, "Logging user out failed with error "+err.Error())
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Cookie found")

	token := dto.SessionToken{
		ID: cookie.Value,
	}

	user, err := ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		logger.Error("Logout failed")
		handlerDebugLog(logger, funcName, "Logging user out failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Session verified")

	handlerDebugLog(logger, funcName, fmt.Sprintf("Deleting session for user ID %d", user.Value))
	err = ah.as.LogOut(rCtx, token)
	if err != nil {
		logger.Error("Logout failed")
		handlerDebugLog(logger, funcName, "Logging user out failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Session deleted")

	cookie = &http.Cookie{
		Name:     "tabula_user",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Time{},
		Path:     "/api/v2/",
	}

	http.SetCookie(w, cookie)
	handlerDebugLog(logger, funcName, "Empty cookie set")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Logout failed")
		handlerDebugLog(logger, funcName, "Logging user out failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Logout failed")
		handlerDebugLog(logger, funcName, "Logging user out failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()
	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Logged user out")
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
	rCtx := r.Context()
	funcName := "VerifyAuthEndpoint"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Verifying user's authentication")

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		w.Header().Set("X-Csrf-Token", "")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Cookie found")

	token := dto.SessionToken{
		ID: cookie.Value,
	}

	userID, err := ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		w.Header().Set("X-Csrf-Token", "")
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Session verified")

	handlerDebugLog(logger, funcName, fmt.Sprintf("Getting user info for user ID %d", userID.Value))
	user, err := ah.us.GetWithID(rCtx, userID)
	if err == apperrors.ErrUserNotFound {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		w.Header().Set("X-Csrf-Token", "")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	} else if err != nil {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		w.Header().Set("X-Csrf-Token", "")
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Got user object")

	handlerDebugLog(logger, funcName, fmt.Sprintf("Setting up CSRF for user ID %d", userID.Value))
	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		w.Header().Set("X-Csrf-Token", "")
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "CSRF set up")

	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	handlerDebugLog(logger, funcName, "CSRF token response header set")

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

	handlerDebugLog(logger, funcName, "Marshaling JSON response")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		handlerDebugLog(logger, funcName, "Setting empty CSRF token response header")
		w.Header().Set("X-Csrf-Token", "")
		handlerDebugLog(logger, funcName, "Empty CSRF token response header set")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	handlerDebugLog(logger, funcName, "Writing JSON response")
	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Verification failed")
		handlerDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
		handlerDebugLog(logger, funcName, "Setting empty CSRF token response header")
		w.Header().Set("X-Csrf-Token", "")
		handlerDebugLog(logger, funcName, "Empty CSRF token response header set")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()
	handlerDebugLog(logger, funcName, "Response written")

	logger.Info("Finished verifying user")
}
