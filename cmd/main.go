package main

import (
	"balancer/internal/config"
	"balancer/internal/grpc"
	"balancer/pkg/logger"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatalf("could not create config: %v", err)
	}

	logger.InitLogger(cfg.Logger)
	logger := slog.Default()

	grpcSrv := grpc.NewServer(cfg)

	listener, err := net.Listen("tcp", cfg.APP_PORT)
	if err != nil {
		logger.Info("Error creating TCP listener", "error", err.Error())
	}

	logger.Info("gRPC server started", "port", cfg.APP_PORT)

	go func() {
		if err = grpcSrv.Serve(listener); err != nil {
			logger.Error("Error running gRPC server", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Stopping gRPC server...")
	grpcSrv.GracefulStop()
	logger.Info("Server stopped")
}
