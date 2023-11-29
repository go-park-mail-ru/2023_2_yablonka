package user

import (
	"server/internal/storage"

	embedded "server/internal/service/user/embedded"
	micro "server/internal/service/user/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedUserService(userStorage storage.IUserStorage) *embedded.UserService {
	return embedded.NewUserService(userStorage)
}

// TODO: User microservice
func NewMicroUserService(userStorage storage.IUserStorage, connection *grpc.ClientConn) *micro.UserService {
	return micro.NewUserService(userStorage)
}
