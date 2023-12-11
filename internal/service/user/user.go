package user

import (
	"server/internal/storage"

	micro "server/internal/service/user/microservice"

	"google.golang.org/grpc"
)

// TODO: User microservice
func NewMicroUserService(userStorage storage.IUserStorage, connection *grpc.ClientConn) *micro.UserService {
	return micro.NewUserService(userStorage, connection)
}
