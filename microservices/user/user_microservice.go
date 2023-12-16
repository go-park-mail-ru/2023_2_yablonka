package user_microservice

import (
	logging "server/internal/logging"
	"server/internal/storage"
	user "server/microservices/user/user"

	"google.golang.org/grpc"
)

const nodeName = "microservice"

func RegisterServices(storages *storage.Storages, server *grpc.Server, logger *logging.LogrusLogger) {
	funcName := "Auth.RegisterServices"
	userServer := user.NewUserService(storages.User, logger)
	logger.DebugRequestlessFmt("Auth GRPC server created", funcName, nodeName)

	user.RegisterUserServiceServer(server, userServer)
	logger.DebugRequestlessFmt("Auth GRPC server registered", funcName, nodeName)
}
