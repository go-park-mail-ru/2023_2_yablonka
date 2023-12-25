package app

import (
	"net/http"
	_ "server/docs"
	"server/internal/app/handlers"
	"server/internal/app/middleware"
	"server/internal/app/middleware/parsers"
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

	mux.Route("/api/v3", func(r chi.Router) {
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
			r.Put("/", metricsMiddleware.WrapHandler(
				"/user/", http.HandlerFunc(manager.UserHandler.ChangeProfile)),
			)
			r.Put("/password/", metricsMiddleware.WrapHandler(
				"/user/password/", http.HandlerFunc(manager.UserHandler.ChangePassword)),
			)
			r.Route("/avatar", func(r chi.Router) {
				r.Put("/", metricsMiddleware.WrapHandler(
					"/user/avatar/", http.HandlerFunc(manager.UserHandler.ChangeAvatar)),
				)
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/user/avatar/delete", http.HandlerFunc(manager.UserHandler.DeleteAvatar)),
				)
			})
		})
		r.Route("/workspace", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/workspace/create/", http.HandlerFunc(manager.WorkspaceHandler.Create)),
			)

			r.Route("/{workspaceID}", func(r chi.Router) {
				r.Use(parsers.WorkspaceCtx)
				r.Put("/", metricsMiddleware.WrapHandler(
					"/workspace/update/", http.HandlerFunc(manager.WorkspaceHandler.UpdateData)),
				)
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/workspace/delete/", http.HandlerFunc(manager.WorkspaceHandler.Delete)),
				)
			})
		})
		r.Route("/board", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/board/create/", http.HandlerFunc(manager.BoardHandler.Create)),
			)

			r.Route("/{boardID}", func(r chi.Router) {
				r.Use(parsers.BoardCtx)
				r.Get("/", metricsMiddleware.WrapHandler(
					"/board/", http.HandlerFunc(manager.BoardHandler.GetFullBoard)),
				)
				r.Put("/", metricsMiddleware.WrapHandler(
					"/board/update/", http.HandlerFunc(manager.BoardHandler.UpdateData)),
				)
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/board/delete/", http.HandlerFunc(manager.BoardHandler.Delete)),
				)
				r.Route("/thumbnail", func(r chi.Router) {
					r.Put("/", metricsMiddleware.WrapHandler(
						"/board/update/thumbnail/", http.HandlerFunc(manager.BoardHandler.UpdateThumbnail)),
					)
				})
				r.Route("/user/{userID}", func(r chi.Router) {
					r.Use(parsers.UserCtx)
					r.Put("/", metricsMiddleware.WrapHandler(
						"/board/user/add/", http.HandlerFunc(manager.BoardHandler.AddUser)),
					)
					r.Delete("/", metricsMiddleware.WrapHandler(
						"/board/user/remove/", http.HandlerFunc(manager.BoardHandler.RemoveUser)),
					)
				})
				r.Route("/history", func(r chi.Router) {
					r.Get("/", metricsMiddleware.WrapHandler(
						"/board/history/", http.HandlerFunc(manager.BoardHandler.GetHistory)),
					)
					r.Put("/", metricsMiddleware.WrapHandler(
						"/board/history/submit/", http.HandlerFunc(manager.BoardHandler.SubmitEdit)),
					)
				})
			})
		})
		r.Route("/list", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/list/create/", http.HandlerFunc(manager.ListHandler.Create)),
			)

			r.Route("/{listID}", func(r chi.Router) {
				r.Use(parsers.ListCtx)
				r.Put("/", metricsMiddleware.WrapHandler(
					"/list/edit/", http.HandlerFunc(manager.ListHandler.Update)),
				)
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/list/delete/", http.HandlerFunc(manager.ListHandler.Delete)),
				)
				r.Put("/reorder/", metricsMiddleware.WrapHandler(
					"/list/reorder/", http.HandlerFunc(manager.ListHandler.UpdateOrder)),
				)
			})
		})
		r.Route("/task", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/task/create/", http.HandlerFunc(manager.TaskHandler.Create)),
			)

			r.Route("/{taskID}", func(r chi.Router) {
				r.Use(parsers.TaskCtx)
				r.Get("/", metricsMiddleware.WrapHandler(
					"/task/", http.HandlerFunc(manager.TaskHandler.Read)),
				)
				r.Put("/", metricsMiddleware.WrapHandler(
					"/task/edit/", http.HandlerFunc(manager.TaskHandler.Update)),
				)
				r.Post("/move/", metricsMiddleware.WrapHandler(
					"/task/move/", http.HandlerFunc(manager.TaskHandler.Move)),
				)
				r.Route("/user/{userID}", func(r chi.Router) {
					r.Use(parsers.UserCtx)
					r.Post("/add/", metricsMiddleware.WrapHandler(
						"/task/user/add/", http.HandlerFunc(manager.TaskHandler.AddUser)),
					)
					r.Post("/remove/", metricsMiddleware.WrapHandler(
						"/task/user/remove/", http.HandlerFunc(manager.TaskHandler.RemoveUser)),
					)
				})
				r.Route("/file", func(r chi.Router) {
					r.Post("/", metricsMiddleware.WrapHandler(
						"/task/file/attach/", http.HandlerFunc(manager.TaskHandler.AttachFile)),
					)
					r.Get("/", metricsMiddleware.WrapHandler(
						"/task/file/", http.HandlerFunc(manager.TaskHandler.GetFileList)),
					)
					r.Delete("/", metricsMiddleware.WrapHandler(
						"/task/file/remove/", http.HandlerFunc(manager.TaskHandler.RemoveFile)),
					)
				})
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/task/delete/", http.HandlerFunc(manager.TaskHandler.Delete)),
				)
			})
		})
		r.Route("/checklist", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/checklist/create/", http.HandlerFunc(manager.ChecklistHandler.Create)),
			)

			r.Route("/{checklistID}", func(r chi.Router) {
				r.Use(parsers.ChecklistCtx)
				r.Put("/", metricsMiddleware.WrapHandler(
					"/checklist/edit/", http.HandlerFunc(manager.ChecklistHandler.Update)),
				)
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/checklist/delete/", http.HandlerFunc(manager.ChecklistHandler.Delete)),
				)
				r.Route("/item", func(r chi.Router) {
					r.Post("/", metricsMiddleware.WrapHandler(
						"/checklist/item/create/", http.HandlerFunc(manager.ChecklistItemHandler.Create)),
					)
					r.Route("/{checklistItemID}", func(r chi.Router) {
						r.Use(parsers.ChecklistItemCtx)
						r.Put("/", metricsMiddleware.WrapHandler(
							"/checklist/item/edit/", http.HandlerFunc(manager.ChecklistItemHandler.Update)),
						)
						r.Put("/reorder/", metricsMiddleware.WrapHandler(
							"/checklist/item/reorder/", http.HandlerFunc(manager.ChecklistItemHandler.UpdateOrder)),
						)
						r.Delete("/", metricsMiddleware.WrapHandler(
							"/checklist/item/delete/", http.HandlerFunc(manager.ChecklistItemHandler.Delete)),
						)
					})
				})
			})
		})
		r.Route("/comment", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/comment/create/", http.HandlerFunc(manager.CommentHandler.Create)),
			)
		})
		r.Route("/tag", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserService()))
			r.Use(middleware.CSRFMiddleware(manager.AuthHandler.GetCSRFService()))
			r.Post("/", metricsMiddleware.WrapHandler(
				"/tag/create/", http.HandlerFunc(manager.TagHandler.Create)),
			)

			r.Route("/{tagID}", func(r chi.Router) {
				r.Use(parsers.TagCtx)
				r.Put("/", metricsMiddleware.WrapHandler(
					"/tag/update/", http.HandlerFunc(manager.TagHandler.Update)),
				)
				r.Delete("/", metricsMiddleware.WrapHandler(
					"/tag/delete/", http.HandlerFunc(manager.TagHandler.Delete)),
				)
				r.Route("/task/{taskID}", func(r chi.Router) {
					r.Use(parsers.TagCtx)
					r.Post("/", metricsMiddleware.WrapHandler(
						"/tag/task/add/", http.HandlerFunc(manager.TagHandler.AddToTask)),
					)
					r.Delete("/", metricsMiddleware.WrapHandler(
						"/tag/task/remove", http.HandlerFunc(manager.TagHandler.RemoveFromTask)),
					)
				})
			})
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
