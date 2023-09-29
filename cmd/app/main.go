package main

import (
	"log"
	"net/http"

	handlers "server/internal/app/handlers"
	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
	storage "server/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	// TODO Вынести в отдельный файл
	userStorage := storage.NewLocalUserStorage()

	authService := authservice.NewAuthJWTService()
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
	// TODO verifyauth (/api/v1/verifyauth/, POST)

	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start server")
	}
	// TODO graceful shutdown
}
