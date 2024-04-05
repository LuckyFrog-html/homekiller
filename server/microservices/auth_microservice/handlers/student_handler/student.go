package student_handler

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log/slog"
	"server/internal/http_server/network/communication"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

type StudentJson struct {
	Name     string `json:"name"`
	Stage    int64  `json:"stage"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func AddStudentHandler(logger *slog.Logger, storage *postgres.Storage) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var studentData StudentJson
		if err := json.Unmarshal(msg.Data, &studentData); err != nil {
			logger.With(slog.String("request", string(msg.Data))).
				Error("Can't unmarshal data. BadRequest", sl.Err(err))
			reply, err := communication.NewJsonMessage([]byte("Can't unmarshal data. BadRequest"), 400)
			if err != nil {
				logger.Error("Can't marshal message", sl.Err(err))
			}
			msg.Respond(reply)
			return
		}

		student := storage.AddStudent(studentData.Name, studentData.Stage, studentData.Login, studentData.Password)

		marshal, _ := json.Marshal(communication.NewResponse("Student added", []byte(student.Name)))

		reply, err := communication.NewJsonMessage(marshal, 200)
		if err != nil {
			logger.Error("Can't marshal message", sl.Err(err))
		}
		msg.Respond(reply)
	}
}
