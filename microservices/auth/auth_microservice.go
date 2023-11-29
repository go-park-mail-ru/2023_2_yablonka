package auth_microservice

import (
	"server/internal/config"
	logging "server/internal/logging"
	"server/internal/storage"
	auth "server/microservices/auth/auth"

	"google.golang.org/grpc"
)

const nodeName = "microservice"

func RegisterServices(config *config.Config, storages *storage.Storages, server *grpc.Server, logger *logging.LogrusLogger) {
	funcName := "Auth.RegisterServices"
	authServer := auth.NewAuthService(*config.Session, storages.Auth, logger)
	logger.Debug("Auth GRPC server created", funcName, nodeName)

	auth.RegisterAuthServiceServer(server, authServer)
	logger.Debug("Auth GRPC server registered", funcName, nodeName)
}
