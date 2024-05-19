package stan

import (
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type ServerConnection struct {
	nc *nats.Conn
}

func Connect() (*ServerConnection, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to NATS")
	}
	return &ServerConnection{nc: nc}, nil
}

func (s *ServerConnection) CloseConnection() {
	s.nc.Close()
}
