package messaging

import (
	"os"

	"github.com/nats-io/nats.go"
)

func connectToNaTS() (*nats.Conn, error) {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	return nc, err
}

func PublishMessage(subject string, message []byte) error {
	nc, err := connectToNaTS()
	if err != nil {
		return err
	}
	defer nc.Close()
	return nc.Publish(subject, message)
}

func SubscribeToSubject(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	nc, err := connectToNaTS()
	if err != nil {
		return nil, err
	}
	return nc.Subscribe(subject, handler)
}
