package app

import (
	"net/http"
	"server/internal/app/handlers"
	"server/internal/app/middleware"

	chi "github.com/go-chi/chi/v5"
)

// type Mux interface {
// 	GetMux(handlers.HandlerManager) (http.Handler, error)
// }

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.HandlerManager) (http.Handler, error) {
	mux := chi.NewRouter()
	mux.Use(middleware.JsonHeader)
	mux.Use(middleware.ErrorHandler)
	mux.Use(middleware.Logger)
	mux.Use(middleware.CorsNew)
	mux.Use(middleware.PanicRecovery)

	mux.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login/", manager.AuthHandler.LogIn)
		r.Post("/signup/", manager.AuthHandler.SignUp)
		r.Post("/logout/", manager.AuthHandler.LogOut)
		r.Get("/verify/", manager.AuthHandler.VerifyAuthEndpoint)
	})

	mux.Route("/api/v1/user", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserAuthService()))
		r.Get("/boards/", manager.BoardHandler.GetUserBoards)
	})

	return mux, nil
}
