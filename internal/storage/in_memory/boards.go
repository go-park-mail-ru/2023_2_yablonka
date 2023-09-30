package in_memory

import (
	"server/internal/pkg/datatypes"
	"sync"
)

// LocalUserStorage
// Локальное хранение данных
type LocalBoardStorage struct {
	BoardDataByUser map[string][]datatypes.Board
	Mu              *sync.Mutex
}

// NewLocalBoardStorage
// Возвращает локальное хранилище данных с тестовыми данными
func NewBoardStorage() *LocalBoardStorage {
	return &LocalBoardStorage{
		BoardDataByUser: map[string][]datatypes.Board{
			"test@email.com": {
				datatypes.Board{
					ID:           1,
					Name:         "Проект 1",
					OwnerID:      1,
					ThumbnailURL: "https://media.moddb.com/images/downloads/1/203/202069/missing_textures.png",
				},
				datatypes.Board{
					ID:           2,
					Name:         "Разработка Ведра 2",
					OwnerID:      1,
					ThumbnailURL: "https://nicollelamerichs.files.wordpress.com/2022/05/2022043021483800-9e19570e6059798a45aec175873b4ac1.jpg?w=640",
				},
				datatypes.Board{
					ID:           3,
					Name:         "лучшая вещь",
					OwnerID:      1,
					ThumbnailURL: "https://media.istockphoto.com/id/868643608/photo/thumbs-up-emoji-isolated-on-white-background-emoticon-giving-likes-3d-rendering.jpg?s=612x612&w=0&k=20&c=ulAeL-xm8S-g5VU_28CUlOqzqT-ooGTKuXYe097XEL8=",
				},
			},
			"email@example.com": {},
		},
		Mu: &sync.Mutex{},
	}
}
