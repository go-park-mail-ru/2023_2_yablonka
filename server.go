package main

import (
	"log"
	"net/http"

	handlers "server/internal/app/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := handlers.TestApi()

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

	http.HandleFunc("/api/v1/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			api.HandleSignupUser(w, r)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
