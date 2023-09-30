package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"server/internal/app/utils"
	"server/internal/apperrors"
	datatypes "server/internal/pkg/datatypes"
	"server/internal/service"
)

type IAuthHandler interface {
	LogIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	// TODO VerifyAuth
	// TODO LogOut
}

// TODO IUserHandler

type AuthHandler struct {
	as service.IAuthService
	us service.IUserAuthService
}

func NewAuthHandler(as service.IAuthService, us service.IUserAuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
	}
}

func (ah AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var login datatypes.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, `Login error`, http.StatusBadRequest)
		return
	}

	passwordHash := utils.HashFromAuthInfo(login)
	incomingAuth := datatypes.LoginInfo{
		Email:        login.Email,
		PasswordHash: passwordHash,
	}

	ctx := context.Background()

	user, err := ah.us.GetUser(ctx, incomingAuth)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	token, err := ah.as.AuthUser(ctx, user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "user_jwt",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(`{"body": {}}`))
}

func (ah AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var signup datatypes.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		http.Error(w, `Signup error`, http.StatusBadRequest)
		return
	}

	passwordHash := utils.HashFromAuthInfo(signup)
	incomingAuth := datatypes.SignupInfo{
		Email:        signup.Email,
		PasswordHash: passwordHash,
	}

	ctx := context.Background()

	user, err := ah.us.CreateUser(ctx, incomingAuth)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	token, err := ah.as.AuthUser(ctx, user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "user_jwt",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(`{"body": {}}`))
}
