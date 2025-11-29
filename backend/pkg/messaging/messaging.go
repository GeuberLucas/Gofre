package messaging

import (
	"log"

	"github.com/nats-io/nats.go"
)

func connectToNaTS() (*nats.EncodedConn, error) {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Println("Error connecting to NATS:", err)
		return nil, err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Println("Error connecting to NATS:", err)
		return nil, err
	}
	return ec, nil
}

func PublishMessage(subject string, message []byte) error {
	nc, err := connectToNaTS()
	if err != nil {
		return err
	}
	defer nc.Close()
	return nc.Publish(subject, &message)
}

func SubscribeToSubject(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	nc, err := connectToNaTS()
	if err != nil {
		return nil, err
	}
	return nc.Subscribe(subject, handler)
}
