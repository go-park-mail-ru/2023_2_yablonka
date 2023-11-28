package auth_microservice

import (
	"server/internal/config"
	"server/internal/storage"
	auth "server/microservices/auth/auth"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Microservices struct {
	AuthService auth.AuthServiceServer
}

func RegisterServices(config *config.Config, storages *storage.Storages, server *grpc.Server, logger *logrus.Logger) {
	logger.Debug("Creating auth GRPC server")
	authServer := auth.NewAuthService(*config.Session, storages.Auth, logger)
	logger.Debug("Auth GRPC server created")

	auth.RegisterAuthServiceServer(server, authServer)
	logger.Debug("Auth GRPC server registered")
}
