package board

import (
	"server/internal/storage"

	embedded "server/internal/service/board/embedded"
	micro "server/internal/service/board/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedBoardService(
	bs storage.IBoardStorage,
	ts storage.ITaskStorage,
	us storage.IUserStorage,
	cs storage.ICommentStorage,
	cls storage.IChecklistStorage,
	clis storage.IChecklistItemStorage,
) *embedded.BoardService {
	return embedded.NewBoardService(bs, ts, us, cs, cls, clis)
}

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
