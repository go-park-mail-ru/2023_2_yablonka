package main

import (
	"log"
	"net/http"
	"server/internal/app"
	"server/internal/app/handlers"
	config "server/internal/config/session"
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
	log.Println("server configured")
	handlerManager, err := handlers.NewHandlerManager(*serverConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("handlers configured")

	mux, err := app.GetChiMux(*handlerManager)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("router configured")

	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start server")
	}
	// TODO graceful shutdown
}
