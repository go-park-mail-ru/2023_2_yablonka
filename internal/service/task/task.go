package task

import (
	"server/internal/storage"

	embedded "server/internal/service/task/embedded"
	micro "server/internal/service/task/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedTaskService(taskStorage storage.ITaskStorage) *embedded.TaskService {
	return embedded.NewTaskService(taskStorage)
}

// TODO: Task microservice
func NewMicroTaskService(taskStorage storage.ITaskStorage, connection *grpc.ClientConn) *micro.TaskService {
	return micro.NewTaskService(taskStorage)
}
