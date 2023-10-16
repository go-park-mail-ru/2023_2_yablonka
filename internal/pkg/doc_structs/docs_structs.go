package doc_structs

import (
	"server/internal/pkg/entities"
)

type UserResponse struct {
	User entities.User `json:"user"`
}

type UserBoardsResponse struct {
	User   entities.User    `json:"user"`
	Boards []entities.Board `json:"boards"`
}

type NewBoardResponse struct {
	Board entities.Board `json:"board"`
}
