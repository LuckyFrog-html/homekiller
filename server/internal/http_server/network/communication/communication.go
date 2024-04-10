package communication

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log/slog"
	"server/internal/lib/logger/sl"
)

type Message struct {
	Data       []byte
	StatusCode int
}

func NewMessage(data []byte, statusCode int) Message {
	return Message{Data: data, StatusCode: statusCode}
}

func NewJsonMessage(data []byte, statusCode int) ([]byte, error) {
	jsonData, err := json.Marshal(NewMessage(data, statusCode))
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func MessageFromJson(data []byte) (Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return Message{}, err
	}
	return message, nil
}

type Response struct {
	Msg  string `json:"message"`
	Data []byte `json:"data,omitempty"`
}

func NewResponse(msg string, data []byte) Response {
	return Response{Msg: msg, Data: data}
}

func Error(msg *nats.Msg, logger *slog.Logger, err error, errorMessage string, statusCode int) {
	logger.With(slog.String("request", string(msg.Data))).
		Error(errorMessage, sl.Err(err))
	reply, err := NewJsonMessage([]byte(errorMessage), statusCode)
	if err != nil {
		logger.Error("Can't marshal message", sl.Err(err))
	}
	msg.Respond(reply)
}

func SendReply(msg *nats.Msg, logger *slog.Logger, response Response, statusCode int) {
	marshal, _ := json.Marshal(response)
	reply, err := NewJsonMessage(marshal, statusCode)
	if err != nil {
		logger.Error("Can't marshal message", sl.Err(err))
	}
	msg.Respond(reply)
}
