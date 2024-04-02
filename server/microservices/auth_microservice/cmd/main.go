package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
	"path"
	"server/internal/config"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"server/microservices/auth_microservice/handlers/student_handler"
)

func main() {
	var configPath string
	if t := os.Getenv("config_path"); t != "" {
		configPath = t
	} else {
		dir, _ := os.Getwd()
		configPath = path.Join(dir, "config", "local.yaml")
	}
	cfg := config.MustLoad(configPath)
	log := sl.SetupLogger(cfg.Env).With(slog.String("microservice", "Auth"))

	nc, err := nats.Connect(cfg.NatsConf.Host)
	if err != nil {
		log.Error("Nats is not connected", sl.Err(err))
		panic("Nats is not connected!")
	}

	storage, err := postgres.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBConf.Host, cfg.DBConf.User, cfg.DBConf.Password, cfg.DBConf.DBName, cfg.DBConf.Port))

	if err != nil {
		log.Error("Database is not connected", sl.Err(err))
	}

	nc.Subscribe("post.student", student_handler.AddStudentHandler(log, storage))

	select {}
}
