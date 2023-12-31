package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"track_service/adapter"
	"track_service/config"
	"track_service/database"
	"track_service/repositories"
	"track_service/services"
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

	trackRepo := repositories.NewTrackRepository(db)
	trackService := services.NewTrackService(trackRepo, cfg)

	// Set up gRPC server and register the TrackService
	grpcSrv, listener, err := adapter.NewGrpcServer(cfg, trackService)
	if err != nil {
		log.Fatalf("Error creating gRPC server: %v", err)
	}
	go func() {
		if err := grpcSrv.Serve(listener); err != nil {
			log.Fatalf("Error starting gRPC server: %v", err)
		}
	}()

	http.HandleFunc("/", statusHandler)
	go func() {
		if err := http.ListenAndServe(cfg.HTTPPort, nil); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)
	<-quit
	log.Println("Shutting down user server...")

	// Shutdown GRPC server
	grpcSrv.GracefulStop()

	// Close DB connection
	if err := database.CloseDB(db); err != nil {
		log.Fatalf("Failed to close db connection: %v", err)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON := []byte(`{"status": "ok"}`)
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
