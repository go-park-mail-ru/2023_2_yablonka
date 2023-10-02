package main

import (
	"log"
	"net/http"
	"server/internal/app/config"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	serverConfig, err := config.NewEnvConfig("")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("server configured")

	mux := http.NewServeMux()
	err = config.ConfigMux(serverConfig, mux)
	log.Println("mux configured")

	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start server")
	}
	// TODO graceful shutdown
}
