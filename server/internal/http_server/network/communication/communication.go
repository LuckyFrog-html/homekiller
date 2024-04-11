package communication

import (
	"encoding/json"
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
