package app

import (
	"log"
	"net/http"
	"server/internal/app/handlers"
	"server/internal/app/middleware"

	chi "github.com/go-chi/chi/v5"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedHeaders: []string{
			"Acccess-Control-Allow-Origin",
			"Content-Type",
		},
		AllowedOrigins: []string{
			"localhost:8080",
			"213.219.215.40:8080",
			"localhost:8081",
			"213.219.215.40:8081",
			"http://localhost:8081",
			"http://213.219.215.40:8081",
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		Debug:            true,
	})
	log.Println("cors configured")
	mux.Use(c.Handler)
	// mux.Use(middleware.PanicRecovery)

	mux.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login/", manager.AuthHandler.LogIn)
		r.Post("/signup/", manager.AuthHandler.SignUp)
		r.Post("/logout/", manager.AuthHandler.LogOut)
		r.Get("/verify/", manager.AuthHandler.VerifyAuthEndpoint)
	})

	mux.Route("/api/v1/user", func(r chi.Router) {
		r.Use(manager.AuthHandler.VerifyAuthMiddleware)
		r.Get("/boards/", manager.BoardHandler.GetUserBoards)
	})

	return mux, nil
}
