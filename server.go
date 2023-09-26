package main

import (
	"log"
	"net/http"
	"sync"

	handlers "server/internal/app"
	datatypes "server/internal/pkg"
)

func main() {
	handlers := handlers.Handlers{
		Users:  make([]datatypes.User, 0),
		Boards: make([]datatypes.Board, 0),
		Mu:     &sync.Mutex{},
	}

	http.HandleFunc("/api/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleLoginUser(w, r)
			return
		}
	})

	http.HandleFunc("/api/getboards/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodGet {
			handlers.HandleGetBoards(w, r)
			return
		}
	})

	http.HandleFunc("/api/verifyauth/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodGet {
			handlers.HandleVerifyAuth(w, r)
			return
		}
	})

	http.HandleFunc("/api/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleSignupUser(w, r)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
