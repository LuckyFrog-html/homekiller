package main

import (
	"github.com/nats-io/nats.go"
	"os"
	"path"
	"server/internal/config"
	"server/internal/lib/logger/sl"
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
	log := sl.SetupLogger(cfg.Env)

	nc, err := nats.Connect(cfg.NatsConf.Host)
	if err != nil {
		log.Error("Nats is not connected", sl.Err(err))
	}

	nc.Subscribe("AddNewStudent", func(msg *nats.Msg) {
		log.Info("Received a message")
		msg.Respond([]byte("Hello"))
	})
	select {}
}
