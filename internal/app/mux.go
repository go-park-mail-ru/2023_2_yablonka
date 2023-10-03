package app

import (
	"log"
	"net/http"
	"server/internal/app/handlers"

	"github.com/go-chi/chi"
)

// type Mux interface {
// 	GetMux(handlers.HandlerManager) (http.Handler, error)
// }

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.HandlerManager) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(logger)
	mux.Use(jsonHeader)

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

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
	})
}

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
}
