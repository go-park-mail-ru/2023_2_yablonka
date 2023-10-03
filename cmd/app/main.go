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

	var server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

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
	// TODO graceful shutdown
}
