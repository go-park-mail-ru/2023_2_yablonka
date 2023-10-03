package main

import (
	"log"
	"net/http"
	"server/internal/app"
	"server/internal/app/handlers"
	config "server/internal/config/session"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	user "server/internal/service/user"
	"server/internal/storage/in_memory"
)

// @title LA TABULA API
// @version 1.0
// @description haha

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	serverConfig, err := config.NewSessionEnvConfig("")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("config generated")

	userStorage := in_memory.NewUserStorage()
	authStorage := in_memory.NewAuthStorage()
	boardStorage := in_memory.NewBoardStorage()
	log.Println("storages configured")

	log.Println("handlers configured")

	mux, err := app.GetChiMux(*handlers.NewHandlerManager(
		*serverConfig,
		auth.NewAuthSessionService(*serverConfig, authStorage),
		user.NewAuthUserService(userStorage),
		//user.NewUserService(userStorage),
		board.NewBoardService(boardStorage),
	))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("router configured")

	log.Println("server configured")
	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start server")
	}
	// TODO graceful shutdown
}
