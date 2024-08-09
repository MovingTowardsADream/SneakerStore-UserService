package main

import (
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/app"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/config"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Log.Level)

	application := app.NewApp(log, cfg)

	// Run servers
	go func() {
		application.Server.Run()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
	}

	log.Info("Starting graceful shutdown")

	application.Server.Shutdown()

	application.DB.Close()

	log.Info("Gracefully stopped")
}
