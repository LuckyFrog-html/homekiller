package all_gateways

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"server/internal/lib/logger/sl"
	"strings"
	"time"
)

func NewGatewayPostRoute(log *slog.Logger, natsConnection *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var data []byte

		_, err := r.Body.Read(data)
		if err != nil {
			log.Error("Can't read body", sl.Err(err))
			return
		}

		chanelName := "post" + strings.Replace(r.RequestURI, "/", ".", -1)

		log = log.With(slog.String("chanel_name", chanelName))

		msg, err := natsConnection.Request(chanelName,
			data, time.Second*2)
		if err != nil {
			log.Error("Nats request broken", sl.Err(err))
			w.WriteHeader(500)
			return
		}

		_, err = w.Write(msg.Data)
		if err != nil {
			log.Error("Cannot write message", sl.Err(err))
			return
		}
	}
}

func NewGatewayGetRoute(log *slog.Logger, natsConnection *nats.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		msg, err := natsConnection.Request("get."+strings.Replace(r.RequestURI, "/", ".", -1),
			nil, time.Second*2)
		if err != nil {
			log.Error("Nats request broken", sl.Err(err))
			return
		}

		_, err = w.Write(msg.Data)
		if err != nil {
			log.Error("Cannot write message", sl.Err(err))
			return
		}
	}
}
