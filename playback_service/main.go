package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"playback_service/adapter"
	"playback_service/config"
	"playback_service/database"
	"playback_service/repositories"
	"playback_service/services"
	"syscall"
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

	playbackRepo := repositories.NewPlaybackRepository(db)
	trackServiceClient, err := services.NewTrackServiceClient(cfg.TrackHost + cfg.TrackGRPCPort)
	if err != nil {
		log.Fatalf("Error creating gRPC server: %v", err)
	}

	playbackService := services.NewPlaybackService(playbackRepo, trackServiceClient, cfg)

	// Set up gRPC server and register the PlaybackService
	grpcSrv, listener, err := adapter.NewGrpcServer(cfg, playbackService)
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
	log.Println("Shutting down playback server...")

	// Shutdown GRPC server
	grpcSrv.GracefulStop()

	// Close TrackService gRPC connection
	if err := trackServiceClient.Close(); err != nil {
		log.Printf("Error closing TrackService gRPC connection: %v", err)
	}

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
