package server

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// init loads .env file which is located in the project root
func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

// Start runs the server at the address from .env file
func Start() {
	// define location for logging purposes
	wd := "server.Start()"

	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		slog.Error("APP_ADDR environment is missing", "ctx", wd)
		os.Exit(1)
	}

	slog.Info("Server is running", "address", addr)

	if err := newServer().listenAndServe(addr); err != nil {
		slog.Error(err.Error(), "ctx", wd)
		os.Exit(1)
	}
}
