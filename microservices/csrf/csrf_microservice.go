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
	funcName := "RegisterServices"
	logger.Debug("Creating CSRF GRPC server", funcName, nodeName)
	csrfServer := csrf.NewCSRFService(*config.Session, storages.CSRF, logger)
	logger.Debug("CSRF GRPC server created", funcName, nodeName)

	csrf.RegisterCSRFServiceServer(server, csrfServer)
	logger.Debug("CSRF GRPC server registered", funcName, nodeName)
}
