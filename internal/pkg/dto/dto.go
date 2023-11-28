package dto

import (
	"time"
)

type AuthDTO struct {
	VerifiedAuthInfo struct {
		UserID uint64
	}
}

// VerifiedAuthInfo
// DTO, подтверждающее личность на основе сессии, полученных при регистрации
type VerifiedAuthInfo struct {
	UserID uint64
}

// NewCSATAnswerInfo
// DTO с данными ответа на вопрос
type NewCSATAnswerInfo struct {
	ID         uint64 `json:"id"`
	QuestionID uint64 `json:"id_question"`
	Rating     uint64 `json:"rating"`
	//DateCreated time.Time `json:"date_created"`
}

// NewCSATAnswer
// DTO с данными ответа на вопрос и пользователем, ответившим на него
type NewCSATAnswer struct {
	UserID     uint64
	QuestionID uint64
	Rating     uint64
}

// RatingStats
// DTO со статистикой одного рейтинга в вопросе
type RatingStats struct {
	Rating  uint64 `json:"rating"`
	Count   uint64 `json:"count"`
	Average uint64 `json:"average"`
}

// RatingStatsWithQuestionID
// DTO со статистикой одного рейтинга в вопросе
type RatingStatsWithQuestionID struct {
	QuestionID uint64 `json:"question_id"`
	Rating     uint64 `json:"rating"`
	Count      uint64 `json:"count"`
	Average    uint64 `json:"average"`
}

// QuestionWithStats
// DTO со статистикой одного вопроса
type QuestionWithStats struct {
	ID      uint64        `json:"question_id"`
	Content string        `json:"content"`
	Type    string        `json:"type"`
	Stats   []RatingStats `json:"stats"`
}

// CSATQuestionID
// DTO с ID CSAT вопроса
type CSATQuestionID struct {
	Value uint64
}

// CSATQuestionTypeName
// DTO с названием CSAT вопроса
type CSATQuestionTypeName struct {
	Value string `json:"type"`
}

// CSATRatingCheck
// DTO с данными для проверки границ рейтинга
type CSATRatingCheck struct {
	QuestionID uint64 `json:"question_id"`
	MaxRating  uint64 `json:"max_rating"`
}

// CSATRatingCheck
// DTO с данными для проверки границ рейтинга
type CSATQuestionFull struct {
	ID      uint64 `json:"question_id"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

// CSATRatingCheck
// DTO с данными для проверки границ рейтинга
type CSATAnswerFull struct {
	ID      uint64 `json:"answer_id"`
	Content string `json:"rating"`
	Type    string `json:"type"`
}

// NewCSATQuestionInfo
// DTO с новыми данными CSAT вопроса
type NewCSATQuestionInfo struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

// NewCSATQuestion
// DTO с новыми данными CSAT вопроса для БД
type NewCSATQuestion struct {
	Content string `json:"content"`
	TypeID  uint64 `json:"type"`
}

// UpdatedCSATQuestion
// DTO с новыми данными CSAT вопроса для БД
type UpdatedCSATQuestion struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
	Type    uint64 `json:"type"`
}

// UpdatedCSATQuestionInfo
// DTO с новыми данными CSAT вопроса для БД
type UpdatedCSATQuestionInfo struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"`
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
	ID        uint64  `json:"user_id"`
	Email     string  `json:"email" valid:"type(string),email"`
	Name      *string `json:"name"`
	Surname   *string `json:"surname"`
	AvatarURL *string `json:"avatar_url"`
	RoleID    *uint64 `json:"role_id"`
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
	UserID  uint64 `json:"user_id"`
}

// AvatarChangeInfo
// структура для изменения аватарки
type AvatarChangeInfo struct {
	UserID uint64 `json:"-" valid:"-"`
	Avatar []byte `json:"avatar" valid:"-"`
}

// ImageRequest
// структура для изменения аватарки
type ImageUrlInfo struct {
	ID  uint64 `json:"user_id"`
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
	UserID      uint64 `json:"-" valid:"-"`
	OldPassword string `json:"old_password" valid:"type(string),stringlength(8|32)"`
	NewPassword string `json:"new_password" valid:"type(string),stringlength(8|32)"`
}

// UserProfileInfo
// DTO для изменения профиля
type UserProfileInfo struct {
	UserID      uint64 `json:"-" valid:"-"`
	Name        string `json:"name" valid:"stringlength(0|100)"`
	Surname     string `json:"surname,omitempty" valid:"optional,stringlength(0|100)"`
	Description string `json:"description,omitempty" valid:"optional,stringlength(0|256)"`
}

// UserProfileInfo
// DTO для изменения профиля
type UserPublicInfo struct {
	ID          uint64  `json:"user_id"`
	Email       string  `json:"email"`
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Description *string `json:"description"`
	AvatarURL   *string `json:"avatar_url"`
}

type UserOwnerInfo struct {
	ID      uint64  `json:"owner_id"`
	Email   string  `json:"owner_email"`
	Name    *string `json:"owner_name"`
	Surname *string `json:"owner_surname"`
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
	PasswordHash string `json:"password_hash" valid:"type(string)"`
}

// IndividualBoardInfo
// DTO для отдельно взятой доски
type IndividualBoardInfo struct {
	ID           uint64  `json:"board_id"`
	OwnerID      uint64  `json:"owner_id"`
	OwnerEmail   string  `json:"owner_email"`
	Name         *string `json:"board_name"`
	ThumbnailURL *string `json:"thumbnail_url"`
}

// AddBoardUserRequest
// DTO для запроса добавления пользователя в доску
type AddBoardUserRequest struct {
	UserEmail string `json:"user_email"`
	BoardID   uint64 `json:"board_id"`
}

// AddBoardUserInfo
// DTO для добавления пользователя в доску
type AddBoardUserInfo struct {
	UserID  uint64 `json:"user_id"`
	BoardID uint64 `json:"board_id"`
}

// RemoveBoardUserInfo
// DTO для удаления пользователя из доски
type RemoveBoardUserInfo struct {
	UserID  uint64 `json:"user_id"`
	BoardID uint64 `json:"board_id"`
}

// NewBoardInfo
// DTO для новой доски
type NewBoardInfo struct {
	Name         string  `json:"name" valid:"type(string)"`
	WorkspaceID  uint64  `json:"workspace_id" valid:"type(int)"`
	OwnerID      uint64  `json:"owner_id" valid:"type(int)"`
	Thumbnail    *[]byte `json:"thumbnail,omitempty" valid:"optional"`
	ThumbnailURL *string
	Users        []UserPublicInfo `json:"user"`
}

type SingleBoardInfo struct {
	ID               uint64    `json:"board_id"`
	Name             string    `json:"name"`
	WorkspaceID      uint64    `json:"workspace_id"`
	WorkspaceOwnerID uint64    `json:"owner_id"`
	ThumbnailURL     *string   `json:"thumbnail_url"`
	DateCreated      time.Time `json:"date_created"`
}

type SingleListInfo struct {
	ID           uint64 `json:"id"`
	BoardID      uint64 `json:"board_id"`
	Name         string `json:"name"`
	ListPosition uint64 `json:"list_position"`
	Tasks        string `json:"cards"`
}

type SingleTaskInfo struct {
	ID           uint64     `json:"id"`
	ListID       uint64     `json:"list_id"`
	DateCreated  time.Time  `json:"date_created"`
	Name         string     `json:"name"`
	Description  *string    `json:"description"`
	ListPosition uint64     `json:"list_position"`
	Start        *time.Time `json:"start"`
	End          *time.Time `json:"end"`
	Users        string     `json:"users"`
	// Commments    []CommentInfo    `json:"comments"`
	// Checklists
}

type FullBoardResult struct {
	Board SingleBoardInfo  `json:"board"`
	Lists []SingleListInfo `json:"lists"`
	Tasks []SingleTaskInfo `json:"cards"`
	Users []UserPublicInfo `json:"users"`
}

// NewBoardRequest
// DTO для запроса новой доски
type NewBoardRequest struct {
	Name        string  `json:"name" valid:"type(string)"`
	WorkspaceID uint64  `json:"workspace_id" valid:"type(uint64)"`
	Thumbnail   *[]byte `json:"thumbnail,omitempty" valid:"optional"`
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
	ID             string    `json:"session_token"`
	ExpirationDate time.Time `json:"expiration_date"`
}

// CSRFToken
// DTO для токена CSRF
type CSRFData struct {
	Token          string    `json:"csrf_token"`
	Expirationdate time.Time `json:"expiration_date"`
}

// CSRFToken
// DTO для токена CSRF
type CSRFToken struct {
	Value string `json:"csrf_token"`
}

// NewWorkspaceInfo
// DTO для нового рабочего пространства
type NewWorkspaceInfo struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	OwnerID     uint64  `json:"owner_id"`
}

// ListID
// DTO для id списка задач
type ListID struct {
	Value uint64 `json:"id"`
}

// ListID
// DTO для id списов задач
type ListIDs struct {
	Values []uint64 `json:"ids"`
}

// TaskID
// DTO для id задач
type TaskID struct {
	Value uint64 `json:"id"`
}

// CommentID
// DTO для id комментария
type CommentID struct {
	Value uint64 `json:"id"`
}

type ReplyInfo struct {
	OriginalID uint64 `json:"original_id"`
	CommentInfo
}

// CommentInfo
// DTO для данных комментария в задаче
type CommentInfo struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	Text        string    `json:"text"`
	DateCreated time.Time `json:"date_created"`
}

// NewCommentInfo
// DTO для данных нового комментария в задаче
type NewCommentInfo struct {
	UserID uint64 `json:"user_id"`
	TaskID uint64 `json:"task_id"`
	Text   string `json:"text"`
}

// NewListInfo
// DTO для нового списка задач
type NewListInfo struct {
	BoardID      uint64  `json:"board_id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	ListPosition uint64  `json:"list_position"`
}

// NewTaskInfo
// DTO для новой задачи
type NewTaskInfo struct {
	ListID       uint64 `json:"list_id"`
	Name         string `json:"name"`
	ListPosition uint64 `json:"list_position"`
}

// UpdatedTaskInfo
// DTO для обновленной задачи
type UpdatedTaskInfo struct {
	ID           uint64     `json:"id"`
	Name         string     `json:"name"`
	Description  *string    `json:"description,omitempty"`
	Start        *time.Time `json:"start,omitempty"`
	End          *time.Time `json:"end,omitempty"`
	ListPosition string     `json:"list_position"`
}

// UpdatedBoardInfo
// DTO для обновленной доски
type UpdatedBoardInfo struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// UpdatedBoardThumbnailInfo
// DTO для обновленной картинки доски
type UpdatedBoardThumbnailInfo struct {
	ID        uint64 `json:"id"`
	Thumbnail []byte `json:"thumbnail"`
}

// UpdatedListInfo
// DTO для обновленного списка задач
type UpdatedListInfo struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	ListPosition uint64  `json:"list_position"`
}

// UpdatedWorkspaceInfo
// DTO для обновления данных рабочего пространства
type UpdatedWorkspaceInfo struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// AddWorkspaceGuestsInfo
// DTO для изменения списка гостей рабочего пространства
type ChangeWorkspaceGuestsInfo struct {
	WorkspaceID uint64   `json:"id"`
	Guests      []UserID `json:"guests"`
}

// ChangeWorkspaceThumbnailInfo
// DTO для изменения картинки рабочего пространства
type ChangeWorkspaceThumbnailInfo struct {
	ID        uint64 `json:"id"`
	Thumbnail []byte `json:"thumbnail"`
}

// ChangeWorkspaceThumbnailInfo
// DTO для изменения картинки рабочего пространства
type WorkspaceBoardInfo struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	ThumbnailURL *string `json:"thumbnail_url"`
}

// UserOwnedWorkspaceInfo
// DTO для рабочего пространства, принадлежащей конкретному пользователю
type UserOwnedWorkspaceInfo struct {
	ID     uint64               `json:"workspace_id"`
	Name   string               `json:"workspace_name"`
	Boards []WorkspaceBoardInfo `json:"boards"`
}

// UserGuestWorkspaceInfo
// DTO для рабочего пространства, где пользователь гость
type UserGuestWorkspaceInfo struct {
	ID     uint64               `json:"workspace_id"`
	Name   string               `json:"workspace_name"`
	Owner  UserOwnerInfo        `json:"workspace_owner"`
	Boards []WorkspaceBoardInfo `json:"boards"`
}

// AllWorkspaces
// DTO, собирающие все доски отдельно взятого пользователя
type AllWorkspaces struct {
	OwnedWorkspaces []UserOwnedWorkspaceInfo `json:"yourWorkspaces"`
	GuestWorkspaces []UserGuestWorkspaceInfo `json:"guestWorkspaces"`
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
	Email        string  `json:"email" valid:"type(string),email"`
	PasswordHash string  `json:"-"`
	Name         *string `json:"name"`
	Surname      *string `json:"surname"`
	AvatarURL    string  `json:"avatar_url"`
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
	LoggerKey
	BaseURLKey
)
