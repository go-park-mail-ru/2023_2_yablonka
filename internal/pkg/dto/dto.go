package dto

type SignupInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

type AuthInfo struct {
	Email    string
	Password string
}

type VerifiedAuthInfo struct {
	UserID uint64
}

type LoginInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

type IndividualBoardInfo struct {
	ID         uint64 `json:"board_id"`
	OwnerEmail string `json:"owner_email"`
}

type NewBoardInfo struct {
	Name       string `json:"name"`
	OwnerID    uint64 `json:"owner_id"`
	OwnerEmail string `json:"owner_email"`
}

type UserOwnedBoardInfo struct {
	ID           uint64 `json:"board_id"`
	BoardName    string `json:"board_name"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type UserGuestBoardInfo struct {
	OwnerID    uint64 `json:"owner_id"`
	OwnerEmail string `json:"owner_email"`
	UserOwnedBoardInfo
}

type UserInfo struct {
	ID uint64
}

type UpdatedUserInfo struct {
	Email   string `json:"email" valid:"type(string),email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
