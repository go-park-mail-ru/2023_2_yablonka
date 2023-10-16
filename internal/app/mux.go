package app

import (
	"fmt"
	"net/http"
	_ "server/docs"
	"server/internal/app/handlers"
	"server/internal/app/middleware"

	chi "github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// type Mux interface {
// 	GetMux(handlers.HandlerManager) (http.Handler, error)
// }

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.HandlerManager) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(middleware.JsonHeader)
	mux.Use(middleware.ErrorHandler)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Cors)
	mux.Use(middleware.PanicRecovery)

	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login/", manager.AuthHandler.LogIn)
			r.Post("/signup/", manager.AuthHandler.SignUp)
			r.Delete("/logout/", manager.AuthHandler.LogOut)
			r.Get("/verify/", manager.AuthHandler.VerifyAuthEndpoint)
		})
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(manager.AuthHandler.GetAuthService(), manager.AuthHandler.GetUserAuthService()))
			r.Get("/boards/", manager.BoardHandler.GetUserBoards)
		})
		r.Route("/swagger/*", func(r chi.Router) {
			r.Get("/", httpSwagger.Handler(
				httpSwagger.URL("swagger/doc.json")))
		})
	})
	return mux, nil
}

// GetChiMux
// возвращает mux Сваггера, реализованный с помощью модуля chi
func GetSwaggerMux(host string, port string) (http.Handler, error) {
	mux := chi.NewRouter()

	url := fmt.Sprint("http://", host, ":", port, "/swagger/doc.json")
	fmt.Println(url)

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(url),
	))

	return mux, nil
}
