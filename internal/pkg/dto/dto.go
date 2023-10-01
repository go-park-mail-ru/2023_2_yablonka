package dto

type SignupInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

type AuthInfo struct {
	Email    string
	Password string
}

type LoginInfo struct {
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-" valid:"type(string)"`
}

type UserBoardInfo struct {
	ID           uint64 `json:"board_id"`
	BoardName    string `json:"board_name"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type UpdatedUserInfo struct {
	Email   string `json:"email" valid:"type(string),email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
