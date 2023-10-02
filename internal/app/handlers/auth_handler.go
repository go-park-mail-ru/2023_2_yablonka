package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"server/internal/app/utils"
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
	defer r.Body.Close()

	var login dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, `Login error`, http.StatusBadRequest)
		return
	}

	passwordHash := utils.HashFromAuthInfo(login)
	incomingAuth := dto.LoginInfo{
		Email:        login.Email,
		PasswordHash: passwordHash,
	}

	ctx := context.Background()

	user, err := ah.us.GetUser(ctx, incomingAuth)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	// Return as JSON
	token, expiresAt, err := ah.as.AuthUser(ctx, user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
		Expires:  expiresAt,
	}

	http.SetCookie(w, cookie)

	// Return user info
	w.Write([]byte(`{"body": {}}`))
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
	defer r.Body.Close()

	var signup dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		http.Error(w, `Signup error`, http.StatusBadRequest)
		return
	}

	passwordHash := utils.HashFromAuthInfo(signup)
	incomingAuth := dto.SignupInfo{
		Email:        signup.Email,
		PasswordHash: passwordHash,
	}

	ctx := context.Background()

	user, err := ah.us.CreateUser(ctx, incomingAuth)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	token, expiresAt, err := ah.as.AuthUser(ctx, user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "tabula_user",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
		Expires:  expiresAt,
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(`{"body": {}}`))
}

func (ah AuthHandler) VerifyAuth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	token := cookie.Value

	userInfo, err := ah.as.VerifyAuth(ctx, token)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}
	userObj, err := ah.us.GetUserByID(ctx, userInfo.UserID)
	if err == apperrors.ErrUserNotFound {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	json, _ := json.Marshal(userObj)
	response := fmt.Sprintf(`{"body": %s}`, string(json))

	w.Write([]byte(response))
}
