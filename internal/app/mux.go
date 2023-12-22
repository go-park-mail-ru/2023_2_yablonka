package app

import (
	"net/http"
	_ "server/docs"
	"server/internal/app/handlers"
	"server/internal/app/middleware"
	config "server/internal/config"
	logging "server/internal/logging"

	chi "github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// type Mux interface {
// 	GetMux(handlers.HandlerManager) (http.Handler, error)
// }

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.Handlers, config config.Config, logger logging.ILogger, registry *prometheus.Registry) (http.Handler, error) {
	mux := chi.NewRouter()

	metricsMiddleware := middleware.NewPromMiddleware(logger, registry, nil)

	mux.Use(middleware.SetContext(*config.Server, logger))
	mux.Use(middleware.PanicRecovery)
	mux.Use(middleware.GetCors(*config.CORS, logger))
	mux.Use(middleware.JsonHeader)

	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/verify", metricsMiddleware.WrapHandler(
				"/auth/verify", http.HandlerFunc(manager.AuthHandler.VerifyAuthEndpoint)),
			)
			r.Post("/login/", metricsMiddleware.WrapHandler(
				"/auth/login/", http.HandlerFunc(manager.AuthHandler.LogIn)),
			)
			r.Post("/signup/", metricsMiddleware.WrapHandler(
				"/auth/signup/", http.HandlerFunc(manager.AuthHandler.SignUp)),
			)
			r.Delete("/logout/", metricsMiddleware.WrapHandler(
				"/auth/logout/", http.HandlerFunc(manager.AuthHandler.LogOut)),
			)
		})
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Get("/workspaces", metricsMiddleware.WrapHandler(
				"/user/workspaces", http.HandlerFunc(manager.WorkspaceHandler.GetUserWorkspaces)),
			)
			r.Post("/edit/", metricsMiddleware.WrapHandler(
				"/user/edit/", http.HandlerFunc(manager.UserHandler.ChangeProfile)),
			)
			r.Post("/edit/change_password/", metricsMiddleware.WrapHandler(
				"/user/edit/change_password/", http.HandlerFunc(manager.UserHandler.ChangePassword)),
			)
			r.Post("/edit/change_avatar/", metricsMiddleware.WrapHandler(
				"/user/edit/change_avatar/", http.HandlerFunc(manager.UserHandler.ChangeAvatar)),
			)
			r.Delete("/edit/delete_avatar/", metricsMiddleware.WrapHandler(
				"/user/edit/detele_avatar/", http.HandlerFunc(manager.UserHandler.DeleteAvatar)),
			)
		})
		r.Route("/workspace", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/workspace/create/", http.HandlerFunc(manager.WorkspaceHandler.Create)),
			)
			r.Post("/update/", metricsMiddleware.WrapHandler(
				"/workspace/update/", http.HandlerFunc(manager.WorkspaceHandler.UpdateData)),
			)
			r.Delete("/delete/", metricsMiddleware.WrapHandler(
				"/workspace/delete/", http.HandlerFunc(manager.WorkspaceHandler.Delete)),
			)
		})
		r.Route("/board", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/board/", http.HandlerFunc(manager.BoardHandler.GetFullBoard)),
			)
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/board/create/", http.HandlerFunc(manager.BoardHandler.Create)),
			)
			r.Post("/update/", metricsMiddleware.WrapHandler(
				"/board/update/", http.HandlerFunc(manager.BoardHandler.UpdateData)),
			)
			r.Post("/update/change_thumbnail/", metricsMiddleware.WrapHandler(
				"/board/update/change_thumbnail/", http.HandlerFunc(manager.BoardHandler.UpdateThumbnail)),
			)
			r.Route("/user", func(r chi.Router) {
				r.Post("/add/", metricsMiddleware.WrapHandler(
					"/board/user/add/", http.HandlerFunc(manager.BoardHandler.AddUser)),
				)
				r.Post("/remove/", metricsMiddleware.WrapHandler(
					"/board/user/remove/", http.HandlerFunc(manager.BoardHandler.RemoveUser)),
				)
			})
			r.Delete("/delete/", metricsMiddleware.WrapHandler(
				"/board/delete/", http.HandlerFunc(manager.BoardHandler.Delete)),
			)
		})
		r.Route("/list", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/list/create/", http.HandlerFunc(manager.ListHandler.Create)),
			)
			r.Post("/edit/", metricsMiddleware.WrapHandler(
				"/list/edit/", http.HandlerFunc(manager.ListHandler.Update)),
			)
			r.Delete("/delete/", metricsMiddleware.WrapHandler(
				"/list/delete/", http.HandlerFunc(manager.ListHandler.Delete)),
			)
			r.Post("/reorder/", metricsMiddleware.WrapHandler(
				"/list/reorder/", http.HandlerFunc(manager.ListHandler.UpdateOrder)),
			)
		})
		r.Route("/task", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/task/", http.HandlerFunc(manager.TaskHandler.Read)),
			)
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/task/create/", http.HandlerFunc(manager.TaskHandler.Create)),
			)
			r.Post("/edit/", metricsMiddleware.WrapHandler(
				"/task/edit/", http.HandlerFunc(manager.TaskHandler.Update)),
			)
			r.Post("/move/", metricsMiddleware.WrapHandler(
				"/task/move/", http.HandlerFunc(manager.TaskHandler.Move)),
			)
			r.Route("/user", func(r chi.Router) {
				r.Post("/add/", metricsMiddleware.WrapHandler(
					"/task/user/add/", http.HandlerFunc(manager.TaskHandler.AddUser)),
				)
				r.Post("/remove/", metricsMiddleware.WrapHandler(
					"/task/user/remove/", http.HandlerFunc(manager.TaskHandler.RemoveUser)),
				)
			})
			r.Route("/file", func(r chi.Router) {
				r.Post("/attach/", metricsMiddleware.WrapHandler(
					"/task/file/attach/", http.HandlerFunc(manager.TaskHandler.AttachFile)),
				)
				r.Post("/", metricsMiddleware.WrapHandler(
					"/task/file/", http.HandlerFunc(manager.TaskHandler.GetFileList)),
				)
				r.Delete("/remove/", metricsMiddleware.WrapHandler(
					"/task/file/remove/", http.HandlerFunc(manager.TaskHandler.RemoveFile)),
				)
			})
			r.Delete("/delete/", metricsMiddleware.WrapHandler(
				"/task/delete/", http.HandlerFunc(manager.TaskHandler.Delete)),
			)
		})
		r.Route("/checklist", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/checklist/create/", http.HandlerFunc(manager.ChecklistHandler.Create)),
			)
			r.Post("/edit/", metricsMiddleware.WrapHandler(
				"/checklist/edit/", http.HandlerFunc(manager.ChecklistHandler.Update)),
			)
			r.Delete("/delete/", metricsMiddleware.WrapHandler(
				"/checklist/delete/", http.HandlerFunc(manager.ChecklistHandler.Delete)),
			)
			r.Route("/item", func(r chi.Router) {
				r.Post("/create/", metricsMiddleware.WrapHandler(
					"/checklist/item/create/", http.HandlerFunc(manager.ChecklistItemHandler.Create)),
				)
				r.Post("/edit/", metricsMiddleware.WrapHandler(
					"/checklist/item/edit/", http.HandlerFunc(manager.ChecklistItemHandler.Update)),
				)
				r.Post("/reorder/", metricsMiddleware.WrapHandler(
					"/checklist/item/reorder/", http.HandlerFunc(manager.ChecklistItemHandler.UpdateOrder)),
				)
				r.Delete("/delete/", metricsMiddleware.WrapHandler(
					"/checklist/item/delete/", http.HandlerFunc(manager.ChecklistItemHandler.Delete)),
				)
			})
		})
		r.Route("/comment", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/comment/create/", http.HandlerFunc(manager.CommentHandler.Create)),
			)
		})
		r.Route("/tag", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/tag/create/", http.HandlerFunc(manager.TagHandler.Create)),
			)
			r.Post("/update/", metricsMiddleware.WrapHandler(
				"/tag/update/", http.HandlerFunc(manager.TagHandler.Update)),
			)
			r.Post("/add_to_task/", metricsMiddleware.WrapHandler(
				"/tag/add_to_task/", http.HandlerFunc(manager.TagHandler.AddToTask)),
			)
			r.Delete("/remove_from_task/", metricsMiddleware.WrapHandler(
				"/tag/remove_from_task/", http.HandlerFunc(manager.TagHandler.RemoveFromTask)),
			)
			r.Delete("/delete/", metricsMiddleware.WrapHandler(
				"/tag/delete/", http.HandlerFunc(manager.TagHandler.Delete)),
			)
		})
	})
	mux.Route("/swagger/", func(r chi.Router) {
		r.Get("/*", metricsMiddleware.WrapHandler(
			"swagger", httpSwagger.Handler(
				httpSwagger.URL("swagger/doc.json"))),
		)
	})
	mux.Route("/csat", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
		r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
		r.Route("/question", func(r chi.Router) {
			r.Get("/all", metricsMiddleware.WrapHandler(
				"/csat/question/all", http.HandlerFunc(manager.CSATQuestionHandler.GetQuestions)),
			)
			r.Get("/stats", metricsMiddleware.WrapHandler(
				"/csat/question/stats", http.HandlerFunc(manager.CSATQuestionHandler.GetStats)),
			)
			r.Post("/create/", metricsMiddleware.WrapHandler(
				"/csat/question/create/", http.HandlerFunc(manager.CSATQuestionHandler.Create)),
			)
			r.Post("/edit/", metricsMiddleware.WrapHandler(
				"/csat/question/edit/", http.HandlerFunc(manager.CSATQuestionHandler.Update)),
			)
		})
		r.Post("/answer/", metricsMiddleware.WrapHandler(
			"/csat/answer/create/", http.HandlerFunc(manager.CSATAnswerHandler.Create)),
		)
	})
	return mux, nil
}
