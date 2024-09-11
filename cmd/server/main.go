package main

import (
	"log"
	"net/http"

	"file-storage-server/internal/api"
	"file-storage-server/internal/config"
	"file-storage-server/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := api.SetupRoutes(db)

	log.Printf("Server starting on %s", cfg.ServerAddr)
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, router))
}
