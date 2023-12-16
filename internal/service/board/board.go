package board

import (
	"server/internal/storage"

	micro "server/internal/service/board/microservice"

	"google.golang.org/grpc"
)

// TODO: Board microservice
func NewMicroBoardService(bs storage.IBoardStorage,
	ts storage.ITaskStorage,
	us storage.IUserStorage,
	cs storage.ICommentStorage,
	cls storage.IChecklistStorage,
	clis storage.IChecklistItemStorage,
	connection *grpc.ClientConn,
) *micro.BoardService {
	return micro.NewBoardService(bs, ts, us, cs, cls, clis, connection)
}
