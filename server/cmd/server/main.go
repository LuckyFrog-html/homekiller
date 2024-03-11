package main

import (
	"log/slog"
	"os"
	"server/internal/config"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	_, err := postgres.New("host=localhost user=postgres password=1234 dbname=homekiller port=5432 sslmode=disable")

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
