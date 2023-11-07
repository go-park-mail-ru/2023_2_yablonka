package dto

import "time"

// VerifiedAuthInfo
// DTO, подтверждающее личность на основе сессии, полученных при регистрации
type VerifiedAuthInfo struct {
	UserID uint64
}

// VerifiedAuthInfo
// DTO, подтверждающее личность на основе сессии, полученных при регистрации
type UserEmail struct {
	Email string `json:"email" valid:"type(string),email"`
}

// UserPasswordHash
// структура для хранения хэша пароля
type UserPasswordHash struct {
	Value string
}

// UserInWorkspace
// структура для хранения общих данных о пользователе
type UserInWorkspace struct {
	ID        uint64 `json:"user_id"`
	Email     string `json:"email" valid:"type(string),email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	AvatarURL string `json:"avatar_url"`
	RoleID    uint64 `json:"role_id"`
}

// RoleInWorkspace
// структура для хранения общих данных о роли
type RoleInWorkspace struct {
	ID   uint64 `json:"user_id"`
	Name string `json:"name"`
}

// IndividualBoardRequest
// структура для запроса данных доски
type IndividualBoardRequest struct {
	BoardID uint64 `json:"board_id"`
	UserID  string `json:"user_id"`
}

// AvatarChangeInfo
// структура для изменения аватарки
type AvatarChangeInfo struct {
	UserID string `json:"user_id"`
	Avatar []byte `json:"avatar"`
}

// ImageRequest
// структура для изменения аватарки
type ImageUrlInfo struct {
	ID  string `json:"user_id"`
	Url string `json:"avatar_url"`
}

// ImageRequest
// структура для изменения аватарки
type UrlObj struct {
	Value string `json:"url"`
}

// UsersAndRoles
// структура для хранения пользователей и их ролей
type UsersAndRoles struct {
	Users []UserInWorkspace `json:"users"`
	Roles []RoleInWorkspace `json:"roles"`
}

// AuthInfo
// DTO для обработки данных, полученных при входе
type AuthInfo struct {
	Email    string `json:"email" valid:"type(string),email"`
	Password string `json:"password" valid:"type(string),stringlength(8|32)"`
}

// PasswordChangeInfo
// DTO для смены пароля
type PasswordChangeInfo struct {
	UserID      uint64 `json:"user_id"`
	OldPassword string `json:"old_password" valid:"type(string),stringlength(8|32)"`
	NewPassword string `json:"new_password" valid:"type(string),stringlength(8|32)"`
}

// UserProfileInfo
// DTO для изменения профиля
type UserProfileInfo struct {
	UserID      uint64 `json:"user_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Description string `json:"description"`
}

// PasswordChangeInfo
// DTO для обработки данных, полученных при входе
type PasswordHashesInfo struct {
	UserID          uint64 `json:"user_id"`
	NewPasswordHash string
}

// UserLogin
// DTO для обработки входных данных, идентифицирующих пользователя
type UserLogin struct {
	Value string `json:"email" valid:"type(string),email"`
}

// SignupInfo
// DTO для обработки данных, полученных при регистрации
type SignupInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

// LoginInfo
// DTO для обработки данных, полученных при входе
type LoginInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

// IndividualBoardInfo
// DTO для отдельно взятой доски
type IndividualBoardInfo struct {
	ID           uint64 `json:"board_id"`
	OwnerID      uint64 `json:"owner_id"`
	OwnerEmail   string `json:"owner_email"`
	BoardName    string `json:"board_name"`
	ThumbnailURL string `json:"thumbnail_url"`
}

// NewBoardInfo
// DTO для новой доски
type NewBoardInfo struct {
	Name         string `json:"name"`
	OwnerID      uint64 `json:"owner_id"`
	WorkspaceID  uint64 `json:"workspace_id"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
}

// WorkspaceID
// DTO для id рабочего пространства
type WorkspaceID struct {
	Value uint64 `json:"workspace_id"`
}

// BoardID
// DTO для id доски
type BoardID struct {
	Value uint64 `json:"board_id"`
}

// WorkspaceID
// DTO для id пользователя
type UserID struct {
	Value uint64 `json:"user_id"`
}

// SessionToken
// DTO для токена сессии
type SessionToken struct {
	Value string `json:"session_token"`
}

// NewWorkspaceInfo
// DTO для нового рабочего пространства
type NewWorkspaceInfo struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
	OwnerID      uint64 `json:"owner_id"`
}

// ListID
// DTO для id списка задач
type ListID struct {
	Value string `json:"id"`
}

// TaskID
// DTO для id задач
type TaskID struct {
	Value string `json:"id"`
}

// NewListInfo
// DTO для нового списка задач
type NewListInfo struct {
	BoardID      uint64 `json:"board_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ListPosition uint64 `json:"list_position"`
}

// NewTaskInfo
// DTO для новой задачи
type NewTaskInfo struct {
	ListID       uint64    `json:"list_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ListPosition uint64    `json:"list_position"`
}

// UpdatedTaskInfo
// DTO для обновленной задачи
type UpdatedTaskInfo struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ListPosition string    `json:"list_position"`
}

// UpdatedListInfo
// DTO для обновленного списка задач
type UpdatedListInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ListPosition string `json:"list_position"`
}

// UpdatedWorkspaceInfo
// DTO для обновления данных рабочего пространства
type UpdatedWorkspaceInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AddWorkspaceGuestsInfo
// DTO для изменения списка гостей рабочего пространства
type ChangeWorkspaceGuestsInfo struct {
	ID     string   `json:"id"`
	Guests []UserID `json:"guests"`
}

// ChangeWorkspaceThumbnailInfo
// DTO для изменения картинки рабочего пространства
type ChangeWorkspaceThumbnailInfo struct {
	ID        string `json:"id"`
	Thumbnail []byte `json:"thumbnail"`
}

// UserOwnedBoardInfo
// DTO для доски, принадлежащей конкретному пользователю
type UserOwnedBoardInfo struct {
	ID           uint64 `json:"board_id"`
	BoardName    string `json:"board_name"`
	ThumbnailURL string `json:"thumbnail_url"`
}

// UserGuestBoardInfo
// DTO для доски, в которой пользователь является гостем
type UserGuestBoardInfo struct {
	OwnerID    uint64             `json:"owner_id"`
	OwnerEmail string             `json:"owner_email"`
	BoardInfo  UserOwnedBoardInfo `json:"board_info"`
}

// AllWorkspaces
// DTO, собирающие все доски отдельно взятого пользователя
type AllWorkspaces struct {
	OwnedWorkspaces []UserOwnedBoardInfo `json:"user_owned_boards"`
	GuestWorkspaces []UserGuestBoardInfo `json:"user_guest_boards"`
}

// UserInfo
// DTO базовой информации о пользователе
type UserInfo struct {
	ID    uint64
	Email string
}

// UpdatedUserInfo
// DTO изменённой информации о пользователе
type UpdatedUserInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	AvatarURL    string `json:"avatar_url"`
}

type JSONMap map[string]interface{}

type JSONResponse struct {
	Body interface{} `json:"body"`
}

type key int

const (
	UserObjKey key = iota
	BoardsObjKey
	ErrorKey
	SIDLengthKey
)
