package subscriber

import (
	"log"

	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	sc stan.Conn
	ss stan.Subscription
}

func NewSubscriber(sc stan.Conn, subj string) (*Subscriber, error) {
	sub, err := sc.Subscribe(subj, func(m *stan.Msg) {
		logrus.Println("Reading message from subject")
		PrintMsg(m, 0)
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot subscribe to the subject")
	}
	return &Subscriber{sc: sc, ss: sub}, nil
}

func PrintMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
}
