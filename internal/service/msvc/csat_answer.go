package microservice

import (
	context "context"
	"server/internal/pkg/dto"
	"server/internal/storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CSATAnswerService struct {
	storage storage.ICSATAnswerStorage
	UnimplementedCSATSAnswerServiceServer
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewCSATAnswerService(storage storage.ICSATAnswerStorage) *CSATAnswerService {
	return &CSATAnswerService{
		storage: storage,
	}
}

// Create
// создает новый ответ CSAT
// или возвращает ошибки ...
func (cs CSATAnswerService) Create(ctx context.Context, info *NewCSATAnswer) (*emptypb.Empty, error) {
	convertedInfo := dto.NewCSATAnswer{
		UserID:     info.UserID,
		QuestionID: info.QuestionID,
		Rating:     info.Rating,
	}
	return nil, cs.storage.Create(ctx, convertedInfo)
}
