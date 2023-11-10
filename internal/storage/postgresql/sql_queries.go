package postgresql

const (
// userDataQuery = "SELECT id, email, password_hash, name, surname, avatar_url, description FROM user"
)

var (
	allUserFields    = []string{"id", "email", "password_hash", "name", "surname", "avatar_url", "description"}
	allBoardFields   = []string{"id", "id_workspace", "name", "description", "date_created", "thumbnail_url"}
	allListFields    = []string{"id", "id_board", "name", "description", "list_position"}
	allTaskFields    = []string{"id", "id_list", "date_created", "name", "description", "list_position", "start", "end"}
	newTaskFields    = []string{"id_list", "name", "description", "list_position", "start", "end"}
	allSessionFields = []string{"id_user", "duration"}
	//allBoardFields             = []string{"id", "id_workspace", "name", "description", "date_created", "thumbnail_url"}
	allWorkspaceAndBoardFields = []string{
		"public.workspace.id", "public.workspace.name", "public.workspace.description", "public.workspace.date_created", "public.workspace.thumbnail_url",
		"public.boards.id", "public.boards.name", "public.boards.description", "public.boards.date_created", "public.boards.thumbnail_url",
	}
	userWorkspaceFields = []string{
		"workspace.id", "workspace.name", "workspace.description",
		"user.id", "user.email",
		"role.id", "role.name", "role.description",
	}
)
