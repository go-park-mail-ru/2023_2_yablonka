package Tag

import (
	"server/internal/storage"

	micro "server/internal/service/tag/microservice"

	"google.golang.org/grpc"
)

// TODO: Tag microservice
func NewMicroTagService(tagStorage storage.ITagStorage, connection *grpc.ClientConn) *micro.TagService {
	return micro.NewTagService(tagStorage)
}
