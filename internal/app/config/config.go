package config

import (
	"log"
	"net/http"
	"server/internal/app/handlers"
	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
	"server/internal/storage/in_memory"
)

func ConfigMux(mux *http.ServeMux) {
	userStorage := in_memory.NewUserStorage()
	authStorage, err := in_memory.NewAuthStorage()
	if err != nil {
		log.Fatalf("Failed to initialize auth storage, reason: %s", err.Error())
		return
	}

	authService, err := authservice.NewAuthSessionService(authStorage)
	if err != nil {
		log.Fatalf("Failed to initialize server, reason: %s", err.Error())
		return
	}
	userAuthService := userservice.NewAuthUserService(userStorage)
	// userService := userservice.NewUserService(userStorage)
	authHandler := handlers.NewAuthHandler(authService, userAuthService)

	mux.HandleFunc("/api/v1/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			authHandler.LogIn(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			authHandler.SignUp(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	// TODO getboards (/api/v1/getuserboards/, GET)
}
