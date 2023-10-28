package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	apperrors "server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	dto "server/internal/pkg/dto"
	"server/internal/service"

	"github.com/asaskevich/govalidator"
)

type AuthHandler struct {
	as service.IAuthService
	us service.IUserAuthService
}

func (ah AuthHandler) GetAuthService() service.IAuthService {
	return ah.as
}

func (ah AuthHandler) GetUserAuthService() service.IUserAuthService {
	return ah.us
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

	var login dto.AuthInfo

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(login)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	incomingAuth := dto.LoginInfo{
		Email:        login.Email,
		PasswordHash: hashFromAuthInfo(login),
	}

	user, err := ah.us.Login(rCtx, incomingAuth)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	token, expiresAt, err := ah.as.AuthUser(rCtx, user)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expiresAt,
		Path:     "/api/v2/",
	}

	fmt.Println(token)

	http.SetCookie(w, cookie)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": user,
		},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
	w.Write(jsonResponse)
}

// @Summary Зарегистрировать нового пользователя
// @Description Также вводит пользователя в систему
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
// @Failure 409  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/signup/ [post]
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var signup dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(signup)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	passwordHash := hashFromAuthInfo(signup)
	incomingAuth := dto.SignupInfo{
		Email:        signup.Email,
		PasswordHash: passwordHash,
	}

	user, err := ah.us.RegisterUser(rCtx, incomingAuth)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	token, expiresAt, err := ah.as.AuthUser(rCtx, user)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expiresAt,
		Path:     "/api/v2/",
	}

	http.SetCookie(w, cookie)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": user,
		},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}

// @Summary Выйти из системы
// @Description Удаляет текущую сессию пользователя. Может сделать только авторизированный пользователь. Текущая сессия определяется по cookie "tabula_user", в которой лежит строка-токен.
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Success 204 {string} string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/logout/ [delete]
func (ah AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	token := cookie.Value

	_, err = ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	err = ah.as.LogOut(rCtx, token)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	cookie = &http.Cookie{
		Name:     "tabula_user",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Time{},
		Path:     "/api/v2/",
	}
	http.SetCookie(w, cookie)

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
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
	log.Println(r.Cookies())

	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println(cookie.Value)

	token := cookie.Value

	userInfo, err := ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	userObj, err := ah.us.GetUserByID(rCtx, userInfo.UserID)
	if err == apperrors.ErrUserNotFound {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	} else if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": userObj,
		},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}

// TODO salt
func hashFromAuthInfo(info dto.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
