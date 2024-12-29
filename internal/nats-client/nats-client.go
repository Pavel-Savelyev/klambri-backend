package natsclient

import (
	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func Init() error {
	var err error
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	if nc != nil {
		nc.Close()
	}
}

func Publish(subject string, msg []byte) error {
	if nc == nil {
		return nats.ErrConnectionClosed
	}
	return nc.Publish(subject, msg)
}
