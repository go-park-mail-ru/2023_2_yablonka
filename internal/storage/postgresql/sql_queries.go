package postgresql

const (
// userDataQuery = "SELECT id, email, password_hash, name, surname, avatar_url, description FROM user"
)

var (
	allUserFields = []string{"id", "email", "password_hash", "name", "surname", "avatar_url", "description"}
	// allWorkspaceFields = []string{"id", "id_creator", "name", "date_created", "description"}
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
		"public.board.id", "public.board.name", "public.board.description", "public.board.date_created", "public.board.thumbnail_url",
	}
	// userOwnedWorkspaceFields = []string{
	// 	"public.workspace.id", "public.workspace.name", "public.workspace.date_created", "public.workspace.description",
	// 	"public.user.id", "public.user.email", "public.user.name", "public.user.surname", "public.user.description", "public.user.avatar_url",
	// 	"public.board.id", "public.board.name", "public.board.description", "public.board.date_created", "public.board.thumbnail_url",
	// }
	userGuestWorkspaceFields = []string{
		"public.workspace.id", "public.workspace.name",
		"public.user.id", "public.user.email", "public.user.name", "public.user.surname",
	}
)
