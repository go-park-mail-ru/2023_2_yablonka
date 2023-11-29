package board

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	embedded "server/internal/service/board/embedded"
	micro "server/internal/service/board/microservice"

	"google.golang.org/grpc"
)

//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IBoardService interface {
	// GetFullBoard
	// возвращает доску со связанными пользователями, списками и заданиями
	GetFullBoard(context.Context, dto.IndividualBoardRequest) (*dto.FullBoardResult, error)
	// Create
	// создаёт доску и связь с пользователем-создателем
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// UpdateData
	// возвращает доску со связанными пользователями, списками и заданиями
	UpdateData(context.Context, dto.UpdatedBoardInfo) error
	// UpdateThumbnail
	// сохраняет картинку доски в папку images/board_thumbnails с названием id доски и сохраняет ссылку на изображение в БД
	UpdateThumbnail(context.Context, dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error)
	// Delete
	// удаляет доску
	Delete(context.Context, dto.BoardID) error
	// AddUser
	// добавляет пользователя на доску
	AddUser(context.Context, dto.AddBoardUserRequest) error
	// RemoveUser
	// удаляет пользователя с доски
	RemoveUser(context.Context, dto.RemoveBoardUserInfo) error
}

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
