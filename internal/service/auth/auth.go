package auth

import (
	"server/internal/config"
	"server/internal/storage"

	embedded "server/internal/service/auth/embedded"
	micro "server/internal/service/auth/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedAuthService(authStorage storage.IAuthStorage, config config.SessionConfig) *embedded.AuthService {
	return embedded.NewAuthService(config, authStorage)
}

func NewMicroAuthService(authStorage storage.IAuthStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.AuthService {
	return micro.NewAuthService(config, authStorage, connection)
}
