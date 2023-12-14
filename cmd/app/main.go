package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/app"
	"server/internal/app/handlers"
	config "server/internal/config"
	logging "server/internal/logging"
	"server/internal/service"
	"server/internal/storage"
	"server/internal/storage/postgresql"

	"github.com/asaskevich/govalidator"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const configPath string = "config/config.yml"
const envPath string = "config/.env"

// @title LA TABULA API
// @version 2.0
// @description Лучшее и единственно приложение, имитирующее Trello.

// @contact.name Капитанов Даниил
// @contact.url https://vk.com/poophead27
// @contact.email kdanil01@mail.ru

// @license.name None
// @license.url None

// @host localhost:8082
// @BasePath /api/v2
// @query.collection.format multi
func main() {
	config, err := config.LoadConfig(envPath, configPath)
	govalidator.SetFieldsRequiredByDefault(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Config loaded")

	logger, err := logging.NewLogrusLogger(config.Logging)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Logger configured")

	dbConnection, err := postgresql.GetDBConnection(*config.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbConnection.Close()
	logger.Info("Database connection established")

	storages := storage.NewPostgresStorages(dbConnection)
	logger.Info("Storages configured")

	grcpConn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", config.Server.MicroserviceHost, config.Server.MicroservicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Fatal("Failed to connect to the GRPC server as client")
	}
	logger.Info("Connected to GRPC server as client")
	defer grcpConn.Close()

	services := service.NewMicroServices(storages, *config.Session, grcpConn)
	logger.Info("Services configured")

	handlers := handlers.NewHandlers(services)
	logger.Info("Handlers configured")

	mux, err := app.GetChiMux(*handlers, *config, &logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Router configured")

	var server = http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.BackendPort),
		Handler: mux,
	}

	logger.Info("Server is running")

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8012", nil)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logger.Fatal("HTTP server ListenAndServe: " + err.Error())
	}

	<-idleConnsClosed
}
