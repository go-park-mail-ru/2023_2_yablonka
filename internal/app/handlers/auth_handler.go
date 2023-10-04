package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"server/internal/apperrors"
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

//	@Summary Log user into the system
//	@Description Create new session or continue old one
//
//	@Accept  json
//	@Produce  json

//	@Success 200  body object{} true "User ID and comma separated roles"
//	@Failure 400  {object}  error
//	@Failure 401  {object}  error
//	@Failure 500  {object}  error
//
// @Router /api/v1/auth/login [post]
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

	passwordHash := hashFromAuthInfo(login)
	incomingAuth := dto.LoginInfo{
		Email:        login.Email,
		PasswordHash: passwordHash,
	}

	user, err := ah.us.GetUser(rCtx, incomingAuth)
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
		Path:     "/api/v1/",
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

	r.Body.Close()
	w.Write(jsonResponse)
}

//	@Summary Зарегистрировать нового пользователя
//	@Description Также вводит пользователя в систему
//
//	@Accept  json
//	@Produce  json

//	@Success 200  body object{} true "User ID and comma separated roles"
//	@Failure 400  {object}  error
//	@Failure 404  {object}  error
//	@Failure 500  {object}  error
//
// @Router /api/v1/auth/signup [post]
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

	user, err := ah.us.CreateUser(rCtx, incomingAuth)
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
		Path:     "/api/v1/",
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

//	@Summary Выйти из системы
//	@Description Заканчивает текущую сессию
//
//	@Accept  json
//	@Produce  json

//	@Success 200  body object{} true "User ID and comma separated roles"
//	@Failure 400  {object}  error
//	@Failure 404  {object}  error
//	@Failure 500  {object}  error
//
// @Router /api/v1/auth/logout [post]
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
		Path:     "/api/v1/",
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

//	@Summary Выйти из системы
//	@Description Заканчивает текущую сессию
//
//	@Accept  json
//	@Produce  json

//	@Success 200  body object{} true "User ID and comma separated roles"
//	@Failure 400  {object}  error
//	@Failure 404  {object}  error
//	@Failure 500  {object}  error
//
// @Router /api/v1/auth/logout [post]
func (ah AuthHandler) VerifyAuthEndpoint(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

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
