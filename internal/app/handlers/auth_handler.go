package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/internal/apperrors"
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
	log.Println("--------------LogIn Endpoint START--------------")

	rCtx := r.Context()

	var authInfo dto.AuthInfo

	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	_, err = govalidator.ValidateStruct(authInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println("request struct validated")

	user, err := ah.us.CheckPassword(rCtx, authInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("password checked")

	userID := dto.UserID{
		Value: user.ID,
	}

	session, err := ah.as.AuthUser(rCtx, userID)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("user authorized")

	authCookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}

	http.SetCookie(w, authCookie)
	log.Println("authorization cookie set")

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	w.Header().Set("X-Csrf-Token", csrfToken.Token)

	log.Println("csrf set")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": user,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------LogIn Endpoint SUCCESS--------------")
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
	log.Println("--------------SignUp Endpoint START--------------")
	rCtx := r.Context()

	var signup dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	_, err = govalidator.ValidateStruct(signup)
	if err != nil {
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct validated")

	user, err := ah.us.RegisterUser(rCtx, signup)
	if err != nil {
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("user registered")

	userID := dto.UserID{
		Value: user.ID,
	}

	session, err := ah.as.AuthUser(rCtx, userID)
	if err != nil {
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("user authorized")

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    session.ID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpirationDate,
		Path:     "/api/v2/",
	}

	http.SetCookie(w, cookie)
	log.Println("cookie set")

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	w.Header().Set("X-Csrf-Token", csrfToken.Token)

	log.Println("csrf set")

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
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response generated")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------SignUp Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------SignUp Endpoint SUCCESS--------------")
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
	log.Println("--------------LogOut Endpoint START--------------")

	rCtx := r.Context()

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		log.Println(err)
		log.Println("--------------LogOut Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println("cookie decoded")

	token := dto.SessionToken{
		ID: cookie.Value,
	}

	_, err = ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogOut Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("auth verified")

	err = ah.as.LogOut(rCtx, token)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogOut Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("user logged out")

	cookie = &http.Cookie{
		Name:     "tabula_user",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Time{},
		Path:     "/api/v2/",
	}
	http.SetCookie(w, cookie)
	log.Println("cookie set")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogOut Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response generated")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogOut Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------LogOut Endpoint SUCCESS--------------")
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
	log.Println("--------------VerifyAuthEndpoint Endpoint START--------------")

	rCtx := r.Context()
	log.Println("\tDEBUG cookie list")
	for _, c := range r.Cookies() {
		log.Println("\t", c)
	}

	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		log.Println("cookie not found")
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		w.Header().Set("X-Csrf-Token", "")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println(cookie.Value)

	token := dto.SessionToken{
		ID: cookie.Value,
	}

	userID, err := ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		log.Println("failed to verify auth")
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		w.Header().Set("X-Csrf-Token", "")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("auth verified")

	userObj, err := ah.us.GetWithID(rCtx, userID)
	if err == apperrors.ErrUserNotFound {
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	} else if err != nil {
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("user found")

	csrfToken, err := ah.cs.SetupCSRF(rCtx, userID)
	if err != nil {
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	w.Header().Set("X-Csrf-Token", csrfToken.Token)
	log.Println("set X-Csrf-Token header to", csrfToken.Token)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": userObj,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response generated")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------VerifyAuthEndpoint Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------VerifyAuthEndpoint Endpoint SUCCESS--------------")
}
