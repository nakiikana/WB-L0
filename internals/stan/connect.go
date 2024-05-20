package stan

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type ServerConnection struct {
	nc  *nats.Conn
	URL string
}

func Connect() (*ServerConnection, error) {
	fmt.Println("url: ", nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to NATS")
	}
	return &ServerConnection{nc: nc, URL: nats.DefaultURL}, nil
}

func (s *ServerConnection) CloseConnection() {
	s.nc.Close()
}
