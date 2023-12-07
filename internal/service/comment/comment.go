package comment

import (
	"server/internal/storage"

	micro "server/internal/service/comment/microservice"

	"google.golang.org/grpc"
)

// TODO: Board microservice
func NewMicroCommentService(commentStorage storage.ICommentStorage, connection *grpc.ClientConn) *micro.CommentService {
	return micro.NewCommentService(commentStorage)
}
