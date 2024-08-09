package app

import (
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/config"
	"github.com/MovingTowardsADream/SneakerStore-UserService/internal/grpc/grpcserver"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/postgres"
	"log/slog"
)

type App struct {
	Server *grpcserver.Server
	DB     *postgres.Postgres
}

func NewApp(l *slog.Logger, cfg *config.Config) *App {
	// Connect postgres db
	pg, err := postgres.NewPostgres(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		panic("app - Run - postgres.NewPostgresDB: " + err.Error())
	}

	// TODO Migrations

	gRPCServer := grpcserver.New(l, grpcserver.Port(cfg.Port))

	return &App{
		Server: gRPCServer,
		DB:     pg,
	}
}
