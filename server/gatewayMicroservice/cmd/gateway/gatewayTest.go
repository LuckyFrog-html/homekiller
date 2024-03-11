package gateway

import (
	"github.com/nats-io/nats.go"
	"log"
)

func TestConnections() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("asfasfasfas")
		return
	}
	conn.Request("hello", []byte("hello all"), 1)
}
