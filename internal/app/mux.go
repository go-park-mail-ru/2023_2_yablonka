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
func GetChiMux(manager handlers.HandlerManager, config config.BaseServerConfig) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(middleware.JsonHeader)
	mux.Use(middleware.ErrorHandler)
	mux.Use(middleware.Logger)
	mux.Use(middleware.GetCors(config))
	mux.Use(middleware.PanicRecovery)

	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login/", manager.UserHandler.LogIn)
			r.Post("/signup/", manager.UserHandler.SignUp)
			r.Delete("/logout/", manager.UserHandler.LogOut)
			r.Get("/verify", manager.UserHandler.VerifyAuthEndpoint)
		})
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/workspaces/", manager.WorkspaceHandler.GetUserWorkspaces)
			r.Post("/edit", manager.UserHandler.ChangeProfile)
			r.Post("/edit/change_password/", manager.UserHandler.ChangePassword)
			r.Post("/edit/change_avatar", manager.UserHandler.ChangeAvatar)
		})
		r.Route("/workspace", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/create", manager.WorkspaceHandler.Create)
			r.Get("/update/", manager.WorkspaceHandler.UpdateData)
			r.Get("/update/change_thumbnail", manager.WorkspaceHandler.UpdateThumbnail)
			r.Get("/update/change_users", manager.WorkspaceHandler.ChangeGuests)
			r.Get("/delete/", manager.WorkspaceHandler.Delete)
		})
		r.Route("/board", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/", manager.BoardHandler.GetFullBoard)
			r.Get("/create", manager.BoardHandler.Create)
			r.Get("/update/", manager.WorkspaceHandler.UpdateData)
			r.Get("/update/change_thumbnail", manager.WorkspaceHandler.UpdateThumbnail)
			r.Get("/delete", manager.BoardHandler.Delete)
		})
		r.Route("/list", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/create", manager.ListHandler.Create)
			r.Get("/edit", manager.ListHandler.Update)
			r.Get("/delete", manager.ListHandler.Delete)
		})
		r.Route("/task", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/", manager.TaskHandler.Read)
			r.Get("/create", manager.TaskHandler.Create)
			r.Get("/edit", manager.TaskHandler.Update)
			r.Get("/delete", manager.TaskHandler.Delete)
		})
	})
	mux.Route("/swagger/", func(r chi.Router) {
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("swagger/doc.json")))
	})
	return mux, nil
}
