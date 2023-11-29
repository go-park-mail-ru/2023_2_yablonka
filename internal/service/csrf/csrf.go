package csrf

import (
	"server/internal/config"
	"server/internal/storage"

	embedded "server/internal/service/csrf/embedded"
	micro "server/internal/service/csrf/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedCSRFService(csrfStorage storage.ICSRFStorage, config config.SessionConfig) *embedded.CSRFService {
	return embedded.NewCSRFService(config, csrfStorage)
}

// TODO: CSRF microservice
func NewMicroCSRFService(csrfStorage storage.ICSRFStorage, config config.SessionConfig, connection *grpc.ClientConn) *micro.CSRFService {
	return micro.NewCSRFService(config, csrfStorage)
}
