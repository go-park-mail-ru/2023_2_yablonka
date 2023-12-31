package postgresql

const (
// userDataQuery = "SELECT id, email, password_hash, name, surname, avatar_url, description FROM user"
)

var (
	allUserFields       = []string{"id", "email", "password_hash", "name", "surname", "avatar_url", "description"}
	allPublicUserFields = []string{"public.user.id", "public.user.email", "public.user.name", "public.user.surname", "public.user.description", "public.user.avatar_url"}

	// allWorkspaceFields = []string{"id", "id_creator", "name", "date_created", "description"}

	// allBoardFields     = []string{"id", "id_workspace", "name", "date_created", "thumbnail_url"}
	allBoardFields = []string{"public.board.id", "public.board.id_workspace", "public.workspace.id_creator",
		"public.board.name", "public.board.date_created", "public.board.thumbnail_url"}

	// allListFields     = []string{"id", "id_board", "name", "list_position"}
	allListTaskAggFields = []string{"public.list.id", "public.list.id_board", "public.list.name", "public.list.list_position",
		"array_remove(array_agg(public.task.id ORDER BY public.task.list_position), NULL)"}

	allTagFields = []string{"public.tag.id", "public.tag.name", "public.tag.color"}

	allTaskFields = []string{"public.task.id", "public.task.id_list", "public.task.date_created",
		"public.task.name", "public.task.description", "public.task.list_position", "public.task.task_start", "public.task.task_end",
		"array_remove(array_agg(public.task_user.id_user ORDER BY public.task_user.id_user), NULL)",
		"array_remove(array_agg(public.comment.id ORDER BY public.comment.date_created), NULL)",
		"array_remove(array_agg(public.checklist.id ORDER BY public.checklist.list_position), NULL)",
		"array_remove(array_agg(public.tag_task.id_tag ORDER BY public.tag_task.id_tag), NULL)",
	}
	allTaskFields2 = []string{
		"public.task.id",
		"public.task.id_list",
		"public.task.date_created",
		"public.task.name",
		"public.task.description",
		"public.task.list_position",
		"public.task.task_start",
		"public.task.task_end",
		"array_remove(array_agg(public.task_user.id_user ORDER BY public.task_user.id_user), NULL)",
		"array_remove(array_agg(public.comment.id ORDER BY public.comment.date_created), NULL)",
		"array_remove(array_agg(public.tag_task.id_tag ORDER BY public.tag_task.id_tag), NULL)",
	}
	allChecklistFields = []string{"public.checklist.id", "public.checklist.id_task", "public.checklist.name", "public.checklist.list_position",
		"array_remove(array_agg(public.checklist_item.id ORDER BY public.checklist_item.list_position), NULL)",
	}
	allTaskAggFields = []string{"public.task.id", "public.task.id_list", "public.task.date_created", "public.task.name",
		"public.task.description", "public.task.list_position", "public.task.task_start", "public.task.task_end",
		"array_remove(array_agg(public.task_user.id_user ORDER BY public.task_user.id_user), NULL)",
	}
	newTaskFields    = []string{"id_list", "name", "list_position"}
	allSessionFields = []string{"id_user", "expiration_date"}

	// allWorkspaceAndBoardFields = []string{
	// 	"public.workspace.id", "public.workspace.name", "public.workspace.description", "public.workspace.date_created",
	// 	"public.board.id", "public.board.name", "public.board.description", "public.board.date_created", "public.board.thumbnail_url",
	// }

	// userOwnedWorkspaceFields = []string{
	// 	"public.workspace.id", "public.workspace.name", "public.workspace.date_created", "public.workspace.description",
	// 	"public.user.id", "public.user.email", "public.user.name", "public.user.surname", "public.user.description", "public.user.avatar_url",
	// 	"public.board.id", "public.board.name", "public.board.description", "public.board.date_created", "public.board.thumbnail_url",
	// }

	userGuestWorkspaceFields = []string{
		"public.workspace.id", "public.workspace.name", "public.workspace.date_created",
		"public.user.id", "public.user.email", "public.user.name", "public.user.surname",
	}
	guestBoardFields = []string{"public.board.id_workspace", "public.board.id", "public.board.name", "public.board.description", "public.board.thumbnail_url"}

	newCommentFields = []string{
		"id_task", "id_user", "content",
	}

	allCommentFields = []string{
		"id", "id_user", "content", "date_created",
	}

	allFileInfoFields = []string{
		"name", "filepath", "date_created",
	}

	allPublicFileInfoFields = []string{
		"public.file.name", "public.file.filepath", "public.file.date_created",
	}

	allHistoryEntryFields = []string{
		"public.edit_history.edit_date", "public.edit_history.edit_summary",
	}

	// taskUserFields = []string{
	// 	"id_task", "id_user",
	// }
)
