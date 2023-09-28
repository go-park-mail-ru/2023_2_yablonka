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
	Email        string `json:"name"`
	PasswordHash string `json:"-"`
	FullName     string `json:"full_name"`
}

type Board struct {
	ID        uint64 `json:"board_id"`
	BoardName string `json:"board_name"`
	OwnerID   uint64 `json:"owner_id"`
}
