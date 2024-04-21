package main

import (
	"fmt"
	"os"
	"path"
	"server/internal/config"
	"server/internal/storage/postgres"
)

func main() {
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "config", "local.yaml")
	cfg := config.MustLoad(configPath)
	storage, err := postgres.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBConf.Host, cfg.DBConf.User, cfg.DBConf.Password, cfg.DBConf.DBName, cfg.DBConf.Port))

	if err != nil {
		fmt.Println(err)
		panic("Database is not connected!")
	}

	_, err = storage.AddTeacher("Андрей", "andrew", "1234")
	if err != nil {
		fmt.Println(err)
	}
}
