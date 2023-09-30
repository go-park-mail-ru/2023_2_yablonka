package datatypes

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

type User struct {
	ID           uint64 `json:"user_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}

type Board struct {
	ID           uint64 `json:"board_id"`
	Name         string `json:"name"`
	OwnerID      uint64 `json:"owner_id"`
	ThumbnailURL string `json:"thumbnail_url"`
}
