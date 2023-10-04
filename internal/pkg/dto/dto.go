package dto

// VerifiedAuthInfo
// DTO, подтверждающее личность на основе сессии, полученных при регистрации
type VerifiedAuthInfo struct {
	UserID uint64
}

// AuthInfo
// DTO для обработки данных, полученных при входе
type AuthInfo struct {
	Email    string `json:"email" valid:"type(string),email"`
	Password string `json:"password" valid:"type(string),stringlength(8|32)"`
}

// LoginInfo
// DTO для обработки входных данных, идентифицирующих пользователя
type LoginInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

// SignupInfo
// DTO для обработки данных, полученных при регистрации
type SignupInfo struct {
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
	Name       string `json:"name"`
	OwnerID    uint64 `json:"owner_id"`
	OwnerEmail string `json:"owner_email"`
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

// UserTotalBoardInfo
// DTO, собирающие все доски отдельно взятого пользователя
type UserTotalBoardInfo struct {
	OwnedBoards []UserOwnedBoardInfo `json:"user_owned_boards"`
	GuestBoards []UserGuestBoardInfo `json:"user_guest_boards"`
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
	Email   string `json:"email" valid:"type(string),email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
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
