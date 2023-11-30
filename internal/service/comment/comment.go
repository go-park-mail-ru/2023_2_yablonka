package comment

import (
	"server/internal/storage"

	embedded "server/internal/service/comment/embedded"
	micro "server/internal/service/comment/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedCommentService(commentStorage storage.ICommentStorage) *embedded.CommentService {
	return embedded.NewCommentService(commentStorage)
}

// TODO: Board microservice
func NewMicroCommentService(commentStorage storage.ICommentStorage, connection *grpc.ClientConn) *micro.CommentService {
	return micro.NewCommentService(commentStorage)
}
