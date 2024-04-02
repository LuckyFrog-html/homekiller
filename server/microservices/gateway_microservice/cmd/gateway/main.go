package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/nats.go"
	"net/http"
	"os"
	"path"
	"server/internal/config"
	"server/internal/http_server/middlewares"
	"server/internal/lib/logger/sl"
	"server/microservices/gateway_microservice/routes/all_gateways"
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

	conn, err := nats.Connect(cfg.NatsConf.Host)
	if err != nil {
		log.Error(natsIsNotConnectedError)
		return
	}

	defer conn.Drain()
	defer conn.Close()

	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Gateway microservice is started")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middlewares.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat) // Хз, надо ли оно нам

	//router.Post("/auth*", auth.New(log, conn))
	router.Post("/*", all_gateways.NewGatewayPostRoute(log, conn))
	router.Get("/*", all_gateways.NewGatewayGetRoute(log, conn))
	http.ListenAndServe(":8080", router)
}
