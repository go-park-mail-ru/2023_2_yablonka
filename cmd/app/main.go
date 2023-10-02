package main

import (
	"log"
	"net/http"
	"server/internal/app/config"
)

func main() {
	serverConfig, err := config.NewEnvConfig("")
	if err != nil {
		log.Fatal(err.Error())
	}

	mux := http.NewServeMux()

	err = config.ConfigMux(serverConfig, mux)

	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start server")
	}
	// TODO graceful shutdown

}
