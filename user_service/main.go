package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user_service/adapter"
	"user_service/config"
	"user_service/database"
	"user_service/repositories"
	"user_service/services"
)

func main() {
	// Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config.")
	}

	// Logger
	logger := zerolog.New(os.Stderr)

	// Database
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database.")
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, cfg)

	// Set up gRPC server and register the UserService
	grpcSrv, listener, reg, err := adapter.NewGrpcServer(cfg, userService, logger)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating gRPC server")
	}
	go func() {
		if err := grpcSrv.Serve(listener); err != nil {
			log.Fatal().Err(err).Msg("Error starting gRPC server")
		}
	}()

	// Start HTTP Server
	http.HandleFunc("/", statusHandler)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go func() {
		if err := http.ListenAndServe(cfg.HTTPPort, nil); err != nil {
			log.Fatal().Err(err).Msg("Error starting HTTP server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)
	<-quit
	log.Info().Msg("Shutting down user server...")

	// Shutdown GRPC server
	grpcSrv.GracefulStop()

	// Close DB connection
	if err := database.CloseDB(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to close db connection")
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON := []byte(`{"status": "ok"}`)
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
