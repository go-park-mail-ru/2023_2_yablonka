package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"
)

type AuthHandler struct {
	as service.IAuthService
	us service.IUserAuthService
}

// TODO change the default data
// @Summary Log user into the system
// @Description Create new session or continue old one
// @ID login
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} nil
// @Router /api/v1/users/{id} [get]
func (ah AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var login dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.BadRequestResponse))
		return
	}

	passwordHash := hashFromAuthInfo(login)
	incomingAuth := dto.LoginInfo{
		Email:        login.Email,
		PasswordHash: passwordHash,
	}

	user, err := ah.us.GetUser(rCtx, incomingAuth)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}

	token, expiresAt, err := ah.as.AuthUser(rCtx, user)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
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
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}

// TODO change the default data
// @Summary Log user into the system
// @Description Create new session or continue old one
// @ID login
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} nil
// @Router /api/v1/users/{id} [get]
func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var signup dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.BadRequestResponse))
		return
	}

	passwordHash := hashFromAuthInfo(signup)
	incomingAuth := dto.SignupInfo{
		Email:        signup.Email,
		PasswordHash: passwordHash,
	}

	user, err := ah.us.CreateUser(rCtx, incomingAuth)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}

	token, expiresAt, err := ah.as.AuthUser(rCtx, user)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
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
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}

func (ah AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}

func (ah AuthHandler) VerifyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rCtx := r.Context()
		cookie, err := r.Cookie("tabula_user")
		if err != nil {
			*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.GenericUnauthorizedResponse))
			return
		}
		token := cookie.Value
		userInfo, err := ah.as.VerifyAuth(rCtx, token)
		if err != nil {
			*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
			return
		}

		userObj, err := ah.us.GetUserByID(rCtx, userInfo.UserID)
		if err == apperrors.ErrUserNotFound {
			*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.GenericUnauthorizedResponse))
			return
		} else if err != nil {
			*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, "userObj", userObj)))
	})
}

func (ah AuthHandler) VerifyAuthEndpoint(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	cookie, err := r.Cookie("tabula_user")
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.GenericUnauthorizedResponse))
		return
	}

	token := cookie.Value

	userInfo, err := ah.as.VerifyAuth(rCtx, token)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}
	userObj, err := ah.us.GetUserByID(rCtx, userInfo.UserID)
	if err == apperrors.ErrUserNotFound {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.GenericUnauthorizedResponse))
		return
	} else if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": userObj,
		},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.InternalServerErrorResponse))
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
