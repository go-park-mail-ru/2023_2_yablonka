package main

import (
	"log"
	"net/http"

	handlers "server/internal/app/handlers"
	service "server/internal/service/auth"
	storage "server/internal/storage/auth"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	// TestApi в internal/storage/in_memory
	api := handlers.TestApi()

	// Тесты в internal/app/

	// Отдельный пакет в internal
	http.HandleFunc("/api/v1/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			api.HandleLoginUser(w, r)
			return
		}
	})

	// http.HandleFunc("/api/v1/getboards/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")

	// 	log.Println(r.URL.Path)

	// 	if r.Method == http.MethodGet {
	// 		api.HandleGetUserBoards(w, r)
	// 		return
	// 	}
	// })

	// http.HandleFunc("/api/v1/verifyauth/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")

	// 	log.Println(r.URL.Path)

	// 	if r.Method == http.MethodGet {
	// 		api.HandleVerifyAuth(w, r)
	// 		return
	// 	}
	// })

	// Вынести
	authStorage := storage.NewLocalAuthStorage()
	authService := service.NewAuthService(authStorage)
	authHandler := handlers.NewAuthHandler(authService)

	// Проверять метод внутри mux.HandleFunc()
	// default: method not allowed (405)
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

	// Вместо проверки
	http.HandleFunc("/api/v1/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			api.HandleSignupUser(w, r)
			return
		}
	})

	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start server")
	}
	// graceful shutdown
}
