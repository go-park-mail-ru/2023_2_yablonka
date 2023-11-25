package doc_structs

import (
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// VerifiedAuthInfo
// DTO, подтверждающее личность на основе сессии, полученных при регистрации
type UserResponse struct {
	User dto.UserPublicInfo `json:"user"`
}

type UserBoardsResponse struct {
	Boards []entities.Board `json:"boards"`
}

type AllWorkspacesResponse struct {
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
type WorkspaceResponse struct {
	Workspace entities.Workspace `json:"workspace"`
}

type ListResponse struct {
	List entities.List `json:"list"`
}

type CommentResponse struct {
	Comment entities.Comment `json:"comment"`
}

type QuestionsResponse struct {
	Questions []entities.CSATQuestion `json:"questions"`
}
