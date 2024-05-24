package publicher

import (
	"encoding/json"
	"tools/internals/stan/stan-pub/data"

	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tjgq/ticker"
)

type Publisher struct {
	sc stan.Conn
}

func NewPublisher(sc stan.Conn) *Publisher {
	return &Publisher{sc}
}

func (sub *Publisher) SendWithTimeout(done chan bool, tick *ticker.Ticker) error {
	go func() error {
		for {
			data, err := json.Marshal(data.GenerateOrder())
			if err != nil {
				errors.Wrap(err, "error during order marshalig")
				continue
			}
			select {
			case <-done:
				return nil
			case <-tick.C:
				if err := sub.sc.Publish("order", data); err != nil {
					errors.Wrap(err, "error during order publish")
				}
				logrus.Println("New message published")
			}
		}
	}()
	return nil
}
