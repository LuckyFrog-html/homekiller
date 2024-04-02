package student_handler

import (
	"github.com/nats-io/nats.go"
	"log/slog"
	"server/internal/storage/postgres"
)

func AddStudentHandler(logger *slog.Logger, storage *postgres.Storage) nats.MsgHandler {
	return func(msg *nats.Msg) {
		msg.Respond([]byte("Lol kek!"))
	}
}
