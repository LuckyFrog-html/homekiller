package main

import (
	"fmt"
	"os"
	"path"
	"server/internal/config"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func main() {
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "config", "local.yaml")
	cfg := config.MustLoad(configPath)

	log := sl.SetupLogger(cfg.Env)

	//_, err := postgres.New("host=localhost user=postgres password=1234 dbname=homekiller port=5432 sslmode=disable")
	_, err := postgres.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBConf.Host, cfg.DBConf.User, cfg.DBConf.Password, cfg.DBConf.DBName, cfg.DBConf.Port))

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
}
