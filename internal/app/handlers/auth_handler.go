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
	defer r.Body.Close()

	var login dto.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, "Во время авторизации возникла ошибка")),
		)
		return
	}

	passwordHash := hashFromAuthInfo(login)
	incomingAuth := dto.LoginInfo{
		Email:        login.Email,
		PasswordHash: passwordHash,
	}

	ctx := context.Background()

	user, err := ah.us.GetUser(ctx, incomingAuth)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message)),
		)
		return
	}

	// Return as JSON
	token, expiresAt, err := ah.as.AuthUser(ctx, user)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message),
		))
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

	json, _ := json.Marshal(user)
	response := fmt.Sprintf(`{"body": {"user":%s}}`, string(json))

	w.Write([]byte(response))
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, "Во время регистрации возникла ошибка")),
		)
		return
	}

	passwordHash := hashFromAuthInfo(signup)
	incomingAuth := dto.SignupInfo{
		Email:        signup.Email,
		PasswordHash: passwordHash,
	}

	ctx := context.Background()

	user, err := ah.us.CreateUser(ctx, incomingAuth)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message),
		))
		return
	}

	token, expiresAt, err := ah.as.AuthUser(ctx, user)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message),
		))
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

	json, _ := json.Marshal(user)
	response := fmt.Sprintf(`{"body": {"user":%s}}`, string(json))

	w.Write([]byte(response))
}

func (ah AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}

func (ah AuthHandler) VerifyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		cookie, err := r.Cookie("tabula_user")
		if err != nil {
			w.WriteHeader(apperrors.GenericUnauthorizedResponse.Code)
			w.Write([]byte(
				fmt.Sprintf(`{"body": {
                    "error_response": "%s"
                }}`, apperrors.GenericUnauthorizedResponse.Message),
			))
			return
		}
		token := cookie.Value
		rCtx := r.Context()
		userInfo, err := ah.as.VerifyAuth(rCtx, token)
		if err != nil {
			w.WriteHeader(apperrors.ErrorMap[err].Code)
			w.Write([]byte(
				fmt.Sprintf(`{"body": {
                "error_response": "%s"
            }}`, apperrors.ErrorMap[err].Message),
			))
			return
		}

		userObj, err := ah.us.GetUserByID(rCtx, userInfo.UserID)
		if err == apperrors.ErrUserNotFound {
			w.WriteHeader(apperrors.GenericUnauthorizedResponse.Code)
			w.Write([]byte(
				fmt.Sprintf(`{"body": {
                "error_response": "%s"
            }}`, apperrors.GenericUnauthorizedResponse.Message),
			))
			return
		} else if err != nil {
			w.WriteHeader(apperrors.ErrorMap[err].Code)
			w.Write([]byte(
				fmt.Sprintf(`{"body": {
                "error_response": "%s"
            }}`, apperrors.ErrorMap[err].Message),
			))
			return
		}

		ctx := context.WithValue(rCtx, "userObj", userObj)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ah AuthHandler) VerifyAuthEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		w.WriteHeader(apperrors.GenericUnauthorizedResponse.Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.GenericUnauthorizedResponse.Message),
		))
		return
	}

	ctx := context.Background()
	token := cookie.Value

	userInfo, err := ah.as.VerifyAuth(ctx, token)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message),
		))
		return
	}
	userObj, err := ah.us.GetUserByID(ctx, userInfo.UserID)
	if err == apperrors.ErrUserNotFound {
		w.WriteHeader(apperrors.GenericUnauthorizedResponse.Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.GenericUnauthorizedResponse.Message),
		))
		return
	} else if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message),
		))
		return
	}

	json, _ := json.Marshal(userObj)
	response := fmt.Sprintf(`{"body": {"user":%s}}`, string(json))

	w.Write([]byte(response))
}

// TODO salt
func hashFromAuthInfo(info dto.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
