package messaging

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type IMessaging interface {
	PublishMessage(subject string, message []byte) error
	SubscribeToSubject(subject string, handler nats.MsgHandler) (*nats.Subscription, error)
	CloseConnection() error
}

type NATSMessaging struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

func NewNATSMessaging() (IMessaging, error) {
	var urlNats = os.Getenv("NATS_URL")
	nc, err := nats.Connect(urlNats)
	if err != nil {
		log.Println("Error connecting to NATS:", err)
		return nil, err
	}

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Println("Error connecting to NATS:", err)
		return nil, err
	}
	StreamName := "finance"
	cfg := &nats.StreamConfig{
		Name:     StreamName,
		Subjects: []string{"finance.>"},
	}
	stream, err := js.StreamInfo(StreamName)

	if stream == nil {
		log.Printf("Creating stream: %s\n", StreamName)

		_, err = js.AddStream(cfg)
		if err != nil {
			return nil, err
		}
	} else {

		log.Printf("Updating stream: %s\n", StreamName)
		_, err = js.UpdateStream(cfg)
		if err != nil {
			return nil, err
		}
	}
	return &NATSMessaging{js: js, conn: nc}, nil
}

func (n *NATSMessaging) PublishMessage(subject string, message []byte) error {
	_, err := n.js.Publish(subject, message)
	return err
}

func (n *NATSMessaging) SubscribeToSubject(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return n.js.Subscribe(subject, handler)
}

func (n *NATSMessaging) CloseConnection() error {
	if n.conn != nil {
		n.conn.Close()
	}
	return nil
}
