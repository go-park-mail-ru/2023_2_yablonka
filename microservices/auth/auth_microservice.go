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

func NewMicroServices(config *config.Config, storages *storage.Storages) *Microservices {
	return &Microservices{
		AuthService: auth.NewAuthService(*config.Session, storages.Auth),
	}
}

func RegisterServices(config *config.Config, storages *storage.Storages, server *grpc.Server, logger *logrus.Logger) {
	authServer := auth.NewAuthService(*config.Session, storages.Auth)

	auth.RegisterAuthServiceServer(server, authServer)
}
