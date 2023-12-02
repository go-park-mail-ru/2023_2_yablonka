package File

import (
	"server/internal/storage"

	embedded "server/internal/service/file/embedded"
	micro "server/internal/service/file/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedFileService() *embedded.FileService {
	return embedded.NewFileService(FileStorage)
}

// TODO: File microservice
func NewMicroFileService(connection *grpc.ClientConn) *micro.FileService {
	return micro.NewFileService(FileStorage, connection)
}
