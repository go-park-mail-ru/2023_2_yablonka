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
		Users: make([]datatypes.User, 0),
		Mu:    &sync.Mutex{},
	}

	http.HandleFunc("/api/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	})

	http.HandleFunc("/api/getboards/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	})

	http.HandleFunc("/api/verifyauth/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
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
