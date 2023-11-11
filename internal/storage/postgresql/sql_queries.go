package postgresql

const (
// userDataQuery = "SELECT id, email, password_hash, name, surname, avatar_url, description FROM user"
)

var (
	allUserFields    = []string{"id", "email", "password_hash", "name", "surname", "avatar_url", "description"}
	allBoardFields   = []string{"id", "id_workspace", "name", "description", "date_created", "thumbnail_url"}
	allListFields    = []string{"id", "id_board", "name", "description", "list_position"}
	allTaskFields    = []string{"id", "id_list", "date_created", "name", "description", "list_position", "task_start", "task_end"}
	newTaskFields    = []string{"id_list", "name", "description", "list_position", "task_start", "task_end"}
	allSessionFields = []string{"id_user", "expiration_date"}
	//allBoardFields             = []string{"id", "id_workspace", "name", "description", "date_created", "thumbnail_url"}
	// allTaskAndUserFields = []string{
	// 	"public.task.id", "public.task.id_list", "public.task.date_created", "public.task.name",
	// 	"public.task.description", "public.task.list_position", "public.task.task_start", "public.task.task_end",
	// 	"public.user.id", "public.user.email", "public.user.password_hash", "public.user.name", "public.user.surname", "public.user.avatar_url", "public.user.description",
	// }
	allWorkspaceAndBoardFields = []string{
		"public.workspace.id", "public.workspace.name", "public.workspace.description", "public.workspace.date_created", "public.workspace.thumbnail_url",
		"public.boards.id", "public.boards.name", "public.boards.description", "public.boards.date_created", "public.boards.thumbnail_url",
	}
	userOwnedWorkspaceFields = []string{
		"workspace.id", "workspace.name", "workspace.date_created", "workspace.description",
		"user.id", "user.email", "user.name", "user.surname", "user.description", "user.avatar_url",
		"id", "name", "description", "thumbnail_url", "user.id",
	}
	userGuestWorkspaceFields = []string{
		"workspace.id", "workspace.id_creator", "workspace.name", "workspace.date_created", "workspace.description",
		"user.id", "user.email", "user.name", "user.surname", "user.description", "user.avatar_url",
		"id", "name", "description", "thumbnail_url", "user.id",
	}
)
