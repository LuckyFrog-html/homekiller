package all_gateways

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/nats.go"
	"io"
	"log/slog"
	"net/http"
	"server/internal/http_server/network/communication"
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
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read body", sl.Err(err))
			w.WriteHeader(400)
			w.Write([]byte("Can't read body"))
			return
		}

		chanelName := "post" + strings.Replace(r.RequestURI, "/", ".", -1)

		log = log.With(slog.String("chanel_name", chanelName))

		msg, err := natsConnection.Request(chanelName,
			data, time.Second*2)
		if err != nil {
			log.Error("Nats request broken", sl.Err(err))
			w.WriteHeader(500)
			w.Write([]byte("Nats request broken"))
			return
		}

		reply, err := communication.MessageFromJson(msg.Data)
		if err != nil {
			log.Error("Cannot parse message", sl.Err(err))
			w.WriteHeader(500)
			w.Write([]byte("Cannot parse message"))
			return
		}

		w.WriteHeader(reply.StatusCode)
		_, err = w.Write(reply.Data)
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

		reply, err := communication.MessageFromJson(msg.Data)
		if err != nil {
			log.Error("Cannot parse message", sl.Err(err))
			w.Write([]byte("Cannot parse message"))
			return
		}

		_, err = w.Write(reply.Data)
		w.WriteHeader(reply.StatusCode)
		if err != nil {
			log.Error("Cannot write message", sl.Err(err))
			return
		}
	}
}
