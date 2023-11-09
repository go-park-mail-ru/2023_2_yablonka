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
	ThumbnailUrl dto.UrlObj `json:"thumbnail_url"`
}

type AvatarUploadResponse struct {
	AvatarUrl dto.UrlObj `json:"avatar_url"`
}

type TaskResponse struct {
	Task entities.Task `json:"task"`
}

type BoardResponse struct {
	Board entities.Board `json:"board"`
}

type ListResponse struct {
	Task entities.List `json:"list"`
}
