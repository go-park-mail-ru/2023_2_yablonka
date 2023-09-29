package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"server/internal/app/utils"
	"server/internal/apperrors"
	datatypes "server/internal/pkg/datatypes"
	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
)

type IAuthHandler interface {
	LogIn(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	as authservice.IAuthService
	us userservice.IAuthUserService
}

func NewAuthHandler(as authservice.IAuthService, us userservice.IAuthUserService) AuthHandler {
	return AuthHandler{
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
