package doc_structs

import (
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// VerifiedAuthInfo
// DTO, подтверждающее личность на основе сессии, полученных при регистрации
type UserResponse struct {
	User entities.User `json:"user"`
}

type UserBoardsResponse struct {
	User   entities.User    `json:"user"`
	Boards []entities.Board `json:"boards"`
}

type AllWorkspacesResponse struct {
	User       entities.User     `json:"user"`
	Workspaces dto.AllWorkspaces `json:"workspaces"`
}

type ThumbnailUploadResponse struct {
	ThumbnailUrl dto.AllWorkspaces `json:"thumbnail_url"`
}
