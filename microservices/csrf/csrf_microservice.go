package csrf_microservice

import (
	"server/internal/config"
	logger "server/internal/logging"
	"server/internal/storage"

	csrf "server/microservices/csrf/csrf"

	"google.golang.org/grpc"
)

const nodeName = "microservice"

func RegisterServices(config *config.Config, storages *storage.Storages, server *grpc.Server, logger *logger.LogrusLogger) {
	funcName := "CSRF.RegisterServices"
	csrfServer := csrf.NewCSRFService(*config.Session, storages.CSRF, logger)
	logger.Debug("CSRF GRPC service created", funcName, nodeName)

	csrf.RegisterCSRFServiceServer(server, csrfServer)
	logger.Debug("CSRF GRPC service registered", funcName, nodeName)
}
