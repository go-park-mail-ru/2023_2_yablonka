package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	apperrors "server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	dto "server/internal/pkg/dto"
	"server/internal/service"

	"github.com/asaskevich/govalidator"
)

type UserHandler struct {
	as service.IAuthService
	us service.IUserService
}

func (uh UserHandler) GetAuthService() service.IAuthService {
	return uh.as
}

func (uh UserHandler) GetUserService() service.IUserService {
	return uh.us
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
func (uh UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var authInfo dto.AuthInfo

	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(authInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	user, err := uh.us.CheckPassword(rCtx, authInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	token, expiresAt, err := uh.as.AuthUser(rCtx, dto.UserID{Value: user.ID})
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token.Value,
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
	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
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
func (uh UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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

	user, err := uh.us.RegisterUser(rCtx, signup)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	token, expiresAt, err := uh.as.AuthUser(rCtx, dto.UserID{Value: user.ID})
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token.Value,
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

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
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
func (uh UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	token := dto.SessionToken{
		Value: cookie.Value,
	}

	_, err = uh.as.VerifyAuth(rCtx, token)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	err = uh.as.LogOut(rCtx, token)
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

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
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
func (uh UserHandler) VerifyAuthEndpoint(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	log.Println(r.Cookies())

	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println(cookie.Value)

	token := dto.SessionToken{
		Value: cookie.Value,
	}

	userID, err := uh.as.VerifyAuth(rCtx, token)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	userObj, err := uh.us.GetWithID(rCtx, userID)
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

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
}

// @Summary Поменять пароль
// @Description Получает старый и новый пароли, а также id пользователя
// @Tags user
//
// @Accept  json
// @Produce  json
//
// @Param authData body dto.PasswordChangeInfo true "id, старый и новый пароли пользователя"
//
// @Success 200  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/change_password/ [post]
func (uh UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var passwords dto.PasswordChangeInfo

	err := json.NewDecoder(r.Body).Decode(&passwords)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(passwords)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	err = uh.us.UpdatePassword(rCtx, passwords)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
}

// @Summary Поменять данные профиля
// @Description В ответ ничего не шлёт
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param authData body dto.UserProfileInfo true "id пользователя, имя, фамилия, описание пользователя"
//
// @Success 200  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/profile/change/ [post]
func (uh UserHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var newProfileInfo dto.UserProfileInfo

	err := json.NewDecoder(r.Body).Decode(&newProfileInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(newProfileInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	err = uh.us.UpdateProfile(rCtx, newProfileInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
}

// TODO Fix
// @Summary Поменять аватарку
// @Description В ответ шлёт ссылку на файл
// @Tags auth
//
// @Accept  json
// @Produce  json
//
// @Param authData body dto.UserProfileInfo true "id пользователя, имя, фамилия, описание пользователя"
//
// @Success 204  {string} string "no content"
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/profile/change_avatar/ [post]
func (uh UserHandler) ChangeAvatar(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var avatarChangeInfo dto.AvatarChangeInfo

	err := json.NewDecoder(r.Body).Decode(&avatarChangeInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(avatarChangeInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	url, err := uh.us.UpdateAvatar(rCtx, avatarChangeInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"avatar_url": url,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
}
