package task

import (
	"server/internal/storage"

	micro "server/internal/service/task/microservice"

	"google.golang.org/grpc"
)

// TODO: Task microservice
func NewMicroTaskService(taskStorage storage.ITaskStorage, userStorage storage.IUserStorage, connection *grpc.ClientConn) *micro.TaskService {
	return micro.NewTaskService(taskStorage, userStorage)
}
