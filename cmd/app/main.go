package main

import (
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/config"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Log.Level)

	_ = log
}
