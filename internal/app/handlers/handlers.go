package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/dto"
	"server/internal/service"
)

// Handlers
// объект со всеми хэндлерами приложения
type Handlers struct {
	AuthHandler
	UserHandler
	BoardHandler
	WorkspaceHandler
	CommentHandler
	ListHandler
	TaskHandler
	ChecklistHandler
	ChecklistItemHandler
	CSATAnswerHandler
	CSATQuestionHandler
}

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		AuthHandler:          *NewAuthHandler(services.Auth, services.User, services.CSRF),
		UserHandler:          *NewUserHandler(services.User),
		CommentHandler:       *NewCommentHandler(services.Comment),
		BoardHandler:         *NewBoardHandler(services.Auth, services.Board),
		WorkspaceHandler:     *NewWorkspaceHandler(services.Workspace),
		ListHandler:          *NewListHandler(services.List),
		TaskHandler:          *NewTaskHandler(services.Task),
		ChecklistHandler:     *NewChecklistHandler(services.Checklist),
		ChecklistItemHandler: *NewChecklistItemHandler(services.ChecklistItem),
		CSATAnswerHandler:    *NewCSATAnswerHandler(services.CSATAnswer, services.CSATQuestion),
		CSATQuestionHandler:  *NewCSATQuestionHandler(services.CSATQuestion),
	}
}

// NewAuthHandler
// возвращает AuthHandler с необходимыми сервисами
func NewAuthHandler(as service.IAuthService, us service.IUserService, cs service.ICSRFService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
		cs: cs,
	}
}

// NewChecklistItemHandler
// возвращает NewChecklistItemHandler с необходимыми сервисами
func NewChecklistItemHandler(clis service.IChecklistItemService) *ChecklistItemHandler {
	return &ChecklistItemHandler{
		clis: clis,
	}
}

// NewChecklistHandler
// возвращает NewChecklistHandler с необходимыми сервисами
func NewChecklistHandler(cls service.IChecklistService) *ChecklistHandler {
	return &ChecklistHandler{
		cls: cls,
	}
}

// NewCommentHandler
// возвращает AuthHandler с необходимыми сервисами
func NewCommentHandler(cs service.ICommentService) *CommentHandler {
	return &CommentHandler{
		cs: cs,
	}
}

// NewCSATHandler
// возвращает CSATHandler с необходимыми сервисами
func NewCSATQuestionHandler(qs service.ICSATQuestionService) *CSATQuestionHandler {
	return &CSATQuestionHandler{
		qs: qs,
	}
}

// NewCSATHandler
// возвращает CSATHandler с необходимыми сервисами
func NewCSATAnswerHandler(as service.ICSATSAnswerService, qs service.ICSATQuestionService) *CSATAnswerHandler {
	return &CSATAnswerHandler{
		as: as,
		qs: qs,
	}
}

// NewUserHandler
// возвращает UserHandler с необходимыми сервисами
func NewUserHandler(us service.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

// NewBoardHandler
// возвращает BoardHandler с необходимыми сервисами
func NewBoardHandler(as service.IAuthService, bs service.IBoardService) *BoardHandler {
	return &BoardHandler{
		as: as,
		bs: bs,
	}
}

// NewWorkspaceHandler
// возвращает WorkspaceHandler с необходимыми сервисами
func NewWorkspaceHandler(ws service.IWorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{
		ws: ws,
	}
}

// NewListHandler
// возвращает ListHandler с необходимыми сервисами
func NewListHandler(ls service.IListService) *ListHandler {
	return &ListHandler{
		ls: ls,
	}
}

// NewTaskHandler
// возвращает TaskHandler с необходимыми сервисами
func NewTaskHandler(ts service.ITaskService) *TaskHandler {
	return &TaskHandler{
		ts: ts,
	}
}

// WriteResponse
// формирует и отправляет JSON-ответ клиенту
func WriteResponse(response dto.JSONResponse, w http.ResponseWriter, r *http.Request) error {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	r.Body.Close()

	return nil
}
