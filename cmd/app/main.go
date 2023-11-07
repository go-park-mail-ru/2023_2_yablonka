package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/app"
	"server/internal/app/handlers"
	config "server/internal/config/session"
	"server/internal/storage/postgresql"

	"github.com/asaskevich/govalidator"
)

const configPath string = "internal/config/config.yml"
const envPath string = "internal/config/.env"

// @title LA TABULA API
// @version 2.0
// @description Лучшее и единственно приложение, имитирующее Trello.

// @contact.name Капитанов Даниил
// @contact.url https://vk.com/poophead27
// @contact.email kdanil01@mail.ru

// @license.name None
// @license.url None

// @host localhost:8080
// @BasePath /api/v2
// @query.collection.format multi
func main() {
	config, err := config.NewSessionEnvConfig(envPath, configPath)
	govalidator.SetFieldsRequiredByDefault(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("config generated")

	dbConnection, err := postgresql.GetDBConnection(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbConnection.Close()
	log.Println("database connected")

	handlerManager := handlers.NewHandlerManager(dbConnection, config)
	log.Println("handlers configured")

	mux, err := app.GetChiMux(*handlerManager, config.Base)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("router configured")

	var server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server is running")

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
