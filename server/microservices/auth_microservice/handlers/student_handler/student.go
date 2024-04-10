package student_handler

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"server/internal/http_server/network/communication"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/storage/postgres"
)

func AddStudentHandler(logger *slog.Logger, storage *postgres.Storage) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var studentData communicationJson.AddStudentJson
		if err := json.Unmarshal(msg.Data, &studentData); err != nil {
			communication.Error(msg, logger, err, "Can't unmarshal data. BadRequest", http.StatusBadRequest)
			return
		}

		student := storage.AddStudent(studentData.Name, studentData.Stage, studentData.Login, studentData.Password)

		communication.SendReply(msg, logger, communication.NewResponse("Student added", []byte(student.Name)), http.StatusOK)
	}
}

func LoginStudentHandler(logger *slog.Logger, storage *postgres.Storage) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var studentData communicationJson.GetStudentJson
		if err := json.Unmarshal(msg.Data, &studentData); err != nil {
			communication.Error(msg, logger, err, "Can't unmarshal data. BadRequest", http.StatusBadRequest)
			return
		}

		student, err := storage.GetStudentByLogin(studentData.Login, studentData.Password)
		if err != nil {
			communication.Error(msg, logger, err, "Can't find student", http.StatusBadRequest)
			return
		}

		communication.SendReply(msg, logger, communication.NewResponse("Student found", []byte(student.Name)), http.StatusOK)
	}
}
