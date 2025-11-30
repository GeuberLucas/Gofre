package messaging

import (
	"log"

	"github.com/nats-io/nats.go"
)

type IMessaging interface {
	ConnectToBroker() error
	PublishMessage(subject string, message []byte) error
	SubscribeToSubject(subject string, handler nats.MsgHandler) (*nats.Subscription, error)
	CloseConnection() error
}

type NATSMessaging struct {
	conn *nats.EncodedConn
}

func NewNATSMessaging() *NATSMessaging {
	return &NATSMessaging{}
}

func (n *NATSMessaging) ConnectToBroker() error {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Println("Error connecting to NATS:", err)
		return err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Println("Error connecting to NATS:", err)
		return err
	}
	n.conn = ec
	return nil
}

func (n *NATSMessaging) PublishMessage(subject string, message []byte) error {
	return n.conn.Publish(subject, &message)
}

func (n *NATSMessaging) SubscribeToSubject(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return n.conn.Subscribe(subject, handler)
}

func (n *NATSMessaging) CloseConnection() error {
	if n.conn != nil {
		n.conn.Close()
	}
	return nil
}
