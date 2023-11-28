package main

import (
	"fmt"
	"log"
	"net"
	"server/internal/config"
	logging "server/internal/logging"
	"server/internal/service"
	microservice "server/internal/service/msvc"
	"server/internal/storage"
	"server/internal/storage/postgresql"

	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc"
)

const configPath string = "config/config.yml"
const envPath string = "config/.env"

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

	server := grpc.NewServer()

	microservices := service.NewMicroServices(storages)

	lstn, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.MicroservicePort))
	if err != nil {
		logger.Fatal("Can't listen to port, " + err.Error())
	}

	microservice.RegisterCSATSAnswerServiceServer(server, microservices.CSATAnswer)
	microservice.RegisterCSATQuestionServiceServer(server, microservices.CSATQuestion)

	server.Serve(lstn)
}
