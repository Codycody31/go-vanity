// main.go

package main

import (
	"flag"
	"log"
	"net/http"

	"go.codycody31.dev/go-vanity/config"
	"go.codycody31.dev/go-vanity/server"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	router := server.NewRouter(cfg)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
