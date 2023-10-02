package mux_stuff

import (
	"log"
	"net/http"
	session "server/internal/app/config/session"
	"server/internal/app/handlers"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	user "server/internal/service/user"
	"server/internal/storage/in_memory"
)

// SessionConfigMux
// обвешивает mux приложения хендлерами
func SessionConfigMux(config session.SessionServerConfig, mux *http.ServeMux) error {
	ok, err := config.Validate()
	if !ok {
		return err
	}

	userStorage := in_memory.NewUserStorage()
	authStorage := in_memory.NewAuthStorage()
	boardStorage := in_memory.NewBoardStorage()

	authService := auth.NewAuthSessionService(config.SessionIDLength, config.Base.SessionDuration, authStorage)
	userAuthService := user.NewAuthUserService(userStorage)
	boardService := board.NewBoardService(boardStorage)
	// userService := userservice.NewUserService(userStorage)
	authHandler := handlers.NewAuthHandler(authService, userAuthService)
	boardHandler := handlers.NewBoardHandler(authService, boardService)

	mux.HandleFunc("/api/v1/auth/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			authHandler.LogIn(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/auth/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			authHandler.SignUp(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/user/boards/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			boardHandler.GetUserBoards(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/auth/verify/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			authHandler.VerifyAuth(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})
	return nil
}
