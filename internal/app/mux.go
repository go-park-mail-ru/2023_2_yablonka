package app

import (
	"net/http"
	_ "server/docs"
	"server/internal/app/handlers"
	"server/internal/app/middleware"
	config "server/internal/config"

	chi "github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// type Mux interface {
// 	GetMux(handlers.HandlerManager) (http.Handler, error)
// }

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.Handlers, config config.Config) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(middleware.SetLogger(*config.Logging, *config.Server))
	mux.Use(middleware.PanicRecovery)
	mux.Use(middleware.GetCors(*config.CORS, *config.Logging))
	mux.Use(middleware.JsonHeader)

	// Testing in-place error handling
	// mux.Use(middleware.ErrorHandler)

	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/verify", manager.AuthHandler.VerifyAuthEndpoint)
			r.Post("/login/", manager.AuthHandler.LogIn)
			r.Post("/signup/", manager.AuthHandler.SignUp)
			r.Delete("/logout/", manager.AuthHandler.LogOut)
		})
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Get("/workspaces", manager.WorkspaceHandler.GetUserWorkspaces)
			r.Post("/edit/", manager.UserHandler.ChangeProfile)
			r.Post("/edit/change_password/", manager.UserHandler.ChangePassword)
			r.Post("/edit/change_avatar/", manager.UserHandler.ChangeAvatar)
		})
		r.Route("/workspace", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", manager.WorkspaceHandler.Create)
			r.Post("/update/", manager.WorkspaceHandler.UpdateData)
			r.Post("/update/change_users/", manager.WorkspaceHandler.ChangeGuests)
			r.Delete("/delete/", manager.WorkspaceHandler.Delete)
		})
		r.Route("/board", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", manager.BoardHandler.GetFullBoard)
			r.Post("/create/", manager.BoardHandler.Create)
			r.Post("/update/", manager.BoardHandler.UpdateData)
			r.Post("/update/change_thumbnail/", manager.BoardHandler.UpdateThumbnail)
			r.Route("/user", func(r chi.Router) {
				r.Post("/add/", manager.BoardHandler.AddUser)
				r.Post("/remove/", manager.BoardHandler.RemoveUser)
			})
			r.Delete("/delete/", manager.BoardHandler.Delete)
		})
		r.Route("/list", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", manager.ListHandler.Create)
			r.Post("/edit/", manager.ListHandler.Update)
			r.Delete("/delete/", manager.ListHandler.Delete)
		})
		r.Route("/task", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", manager.TaskHandler.Read)
			r.Post("/create/", manager.TaskHandler.Create)
			r.Post("/edit/", manager.TaskHandler.Update)
			r.Delete("/delete/", manager.TaskHandler.Delete)
		})
	})
	mux.Route("/swagger/", func(r chi.Router) {
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("swagger/doc.json")))
	})
	mux.Route("/csat", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
		r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
		r.Route("/question", func(r chi.Router) {
			r.Get("/all", manager.CSATQuestionHandler.GetQuestions)
			r.Get("/stats", manager.CSATQuestionHandler.GetStats)
			r.Post("/create/", manager.CSATQuestionHandler.Create)
			r.Post("/edit/", manager.CSATQuestionHandler.Update)
		})
		r.Post("/answer", manager.CSATAnswerHandler.Create)
	})
	return mux, nil
}
