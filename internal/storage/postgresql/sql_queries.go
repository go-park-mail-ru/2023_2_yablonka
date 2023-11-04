package postgresql

const (
// userDataQuery = "SELECT id, email, password_hash, name, surname, avatar_url, description FROM user"
)

var (
<<<<<<< Updated upstream
	allUserFields    = []string{"id", "email", "password_hash", "name", "surname", "avatar_url", "description"}
	allSessionFields = []string{"id_user", "duration"}
	allBoardFields   = []string{"id", "name", "description", "date_created", "thumbnail_url"}
=======
	allUserFields      = []string{"id", "email", "password_hash", "name", "surname", "avatar_url", "description"}
	allSessionFields   = []string{"id_user", "duration"}
	allBoardFields     = []string{"id", "id_workspace", "name", "description", "date_created", "thumbnail_url"}
	allWorkspaceFields = []string{"id", "name", "description", "date_created", "thumbnail_url"}
>>>>>>> Stashed changes
)
