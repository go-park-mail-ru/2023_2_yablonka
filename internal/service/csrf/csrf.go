package csrf

import (
	"server/internal/config"
	"server/internal/storage"

	micro "server/internal/service/csrf/microservice"

	"google.golang.org/grpc"
)

// TODO: CSRF microservice
func NewMicroCSRFService(csrfStorage storage.ICSRFStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.CSRFService {
	return micro.NewCSRFService(config, csrfStorage, connection)
}
