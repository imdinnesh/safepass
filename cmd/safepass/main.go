package main

import (
	"log"
	"net/http"
	"os"

	"github.com/imdinnesh/safepass/internal/config"
	"github.com/imdinnesh/safepass/internal/proxy"
	"github.com/imdinnesh/safepass/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.Load("safepass.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize global logger
	logger.Init(cfg.Debug)

	// Setup reverse proxy server
	handler, err := proxy.NewGateway(cfg)
	if err != nil {
		log.Fatalf("failed to start gateway: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Safepass is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
