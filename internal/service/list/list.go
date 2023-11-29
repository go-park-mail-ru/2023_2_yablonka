package list

import (
	"server/internal/storage"

	embedded "server/internal/service/list/embedded"
	micro "server/internal/service/list/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedListService(listStorage storage.IListStorage) *embedded.ListService {
	return embedded.NewListService(listStorage)
}

// TODO: List microservice
func NewMicroListService(listStorage storage.IListStorage, connection *grpc.ClientConn) *micro.ListService {
	return micro.NewListService(listStorage, connection)
}
