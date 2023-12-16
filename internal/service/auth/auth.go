package auth

import (
	"server/internal/config"
	"server/internal/storage"

	micro "server/internal/service/auth/microservice"

	"google.golang.org/grpc"
)

func NewMicroAuthService(authStorage storage.IAuthStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.AuthService {
	return micro.NewAuthService(config, authStorage, connection)
}
