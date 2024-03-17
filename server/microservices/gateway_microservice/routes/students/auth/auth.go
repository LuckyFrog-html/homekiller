package auth

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"server/internal/lib/logger/sl"
	"time"
)

func New(log *slog.Logger, natsConnection *nats.Conn) http.HandlerFunc {
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

		msg, err := natsConnection.Request("AddNewStudent", data, time.Second*2)
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
