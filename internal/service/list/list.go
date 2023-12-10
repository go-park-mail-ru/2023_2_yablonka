package list

import (
	"server/internal/storage"

	micro "server/internal/service/list/microservice"

	"google.golang.org/grpc"
)

// TODO: List microservice
func NewMicroListService(listStorage storage.IListStorage, connection *grpc.ClientConn) *micro.ListService {
	return micro.NewListService(listStorage, connection)
}
