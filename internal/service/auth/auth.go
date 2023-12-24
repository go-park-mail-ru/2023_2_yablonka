package auth

import (
	"server/internal/config"
	"server/internal/storage"

	micro "server/internal/service/auth/microservice"
	microservice "server/microservices/auth/auth"

	"google.golang.org/grpc"
)

func NewMicroAuthService(authStorage storage.IAuthStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.AuthService {
	client := microservice.NewAuthServiceClient(connection)
	return micro.NewAuthService(config, authStorage, client)
}
