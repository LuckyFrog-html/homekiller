package gateway

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"path"
	"server/internal/config"
	"server/internal/lib/logger/sl"
	"time"
)

const (
	natsIsNotConnectedError = "Nats is not connected!"
)

func main() {
	Start()
}

func Start() {
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "config", "local.yaml")
	cfg := config.MustLoad(configPath)
	log := sl.SetupLogger(cfg.Env)

	conn, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Error(natsIsNotConnectedError)
		return
	}

	defer conn.Drain()

	_, err = conn.Subscribe("hello", func(msg *nats.Msg) {
		fmt.Println("Тестим")
		msg.Respond([]byte("Тестим х2"))
	})

	if err != nil {
		log.Error(err.Error())
	}

	req, err := conn.Request("hello", []byte("hi"), time.Second)
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(string(req.Data))
	log.Info("Gateway microservice is started")
}
