package datatypes

type SignupInfo struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type LoginInfo struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
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
	BoardName    string `json:"board_name"`
	OwnerID      uint64 `json:"owner_id"`
	ThumbnailURL string `json:"thumbnail_url"`
}
