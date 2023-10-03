package app

import (
	"net/http"
	"server/internal/app/handlers"
	"server/internal/app/middleware"

	"github.com/go-chi/chi"
)

// type Mux interface {
// 	GetMux(handlers.HandlerManager) (http.Handler, error)
// }

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.HandlerManager) (http.Handler, error) {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.JsonHeader)

	mux.Route("/api/v1/auth", func(r chi.Router) {
		mux.Post("/login", manager.AuthHandler.LogIn)
		mux.Post("/signup", manager.AuthHandler.SignUp)
		mux.Get("/verify", manager.AuthHandler.VerifyAuth)
	})

	mux.Route("/api/v1/user", func(r chi.Router) {
		mux.Get("/boards", manager.BoardHandler.GetUserBoards)
	})

	return mux, nil
}
