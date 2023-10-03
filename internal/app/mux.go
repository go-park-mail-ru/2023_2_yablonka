package app

import (
	"log"
	"net/http"
	"server/internal/app/handlers"
	session "server/internal/config/session"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	user "server/internal/service/user"
	"server/internal/storage/in_memory"

	"github.com/go-chi/chi"
)

// SessionConfigMux
// обвешивает mux приложения хендлерами
func SessionConfigMux(config session.SessionServerConfig) (http.Handler, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	userStorage := in_memory.NewUserStorage()
	authStorage := in_memory.NewAuthStorage()
	boardStorage := in_memory.NewBoardStorage()

	authService := auth.NewAuthSessionService(config, authStorage)
	userAuthService := user.NewAuthUserService(userStorage)
	//userUserService := user.NewUserService(userStorage)
	boardService := board.NewBoardService(boardStorage)
	authHandler := handlers.NewAuthHandler(authService, userAuthService)
	boardHandler := handlers.NewBoardHandler(authService, boardService)

	router := chi.NewRouter()

	router.Use(logger)
	router.Use(jsonHeader)

	router.Route("/api/v1/auth", func(r chi.Router) {
		router.Post("/login", authHandler.LogIn)
		router.Post("/signup", authHandler.SignUp)
		router.Get("/verify", authHandler.VerifyAuth)
	})

	router.Route("/api/v1/user", func(r chi.Router) {
		router.Get("/boards", boardHandler.GetUserBoards)
	})

	return router, nil
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
