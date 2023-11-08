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
			r.Post("/create", manager.WorkspaceHandler.Create)
			r.Post("/edit/", manager.WorkspaceHandler.UpdateData)
			r.Post("/edit/change_thumbnail", manager.WorkspaceHandler.UpdateThumbnail)
			r.Post("/edit/change_users", manager.WorkspaceHandler.ChangeGuests)
			r.Delete("/delete/", manager.WorkspaceHandler.Delete)
		})
		r.Route("/board", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/", manager.BoardHandler.GetFullBoard)
			r.Post("/create", manager.BoardHandler.Create)
			r.Post("/edit/", manager.BoardHandler.UpdateData)
			r.Post("/edit/change_thumbnail/", manager.BoardHandler.UpdateThumbnail)
			r.Delete("/delete", manager.BoardHandler.Delete)
		})
		r.Route("/list", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Post("/create", manager.ListHandler.Create)
			r.Post("/edit", manager.ListHandler.Update)
			r.Delete("/delete", manager.ListHandler.Delete)
		})
		r.Route("/task", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.UserHandler.GetAuthService(), manager.UserHandler.GetUserService()))
			r.Get("/", manager.TaskHandler.Read)
			r.Post("/create", manager.TaskHandler.Create)
			r.Post("/edit", manager.TaskHandler.Update)
			r.Delete("/delete", manager.TaskHandler.Delete)
		})
	})
	mux.Route("/swagger/", func(r chi.Router) {
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("swagger/doc.json")))
	})
	return mux, nil
}
