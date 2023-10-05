package main

import (
	"log"
	"playback_service/config"
	"playback_service/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	//db.AutoMigrate(&models.User{})

	// Close DB connection
	if err := database.CloseDB(db); err != nil {
		log.Fatalf("Failed to close db connection: %v", err)
	}
}
