package subscriber

import (
	"encoding/json"
	"log"
	"tools/internals/models"
	"tools/internals/service"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	sc      stan.Conn
	service Order
}

func NewSubscriber(sc stan.Conn, subj string, s service.Service) (*Subscriber, error) {
	subsciber := &Subscriber{sc: sc, service: s}
	_, err := subsciber.sc.Subscribe(subj, func(m *stan.Msg) {
		logrus.Println("Reading message from subject")
		err := subsciber.saveMessage(m)
		if err != nil {
			logrus.Errorln(err)
		}
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot subscribe to the subject")
	}

	return subsciber, nil
}

func PrintMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
}

func (s *Subscriber) saveMessage(m *stan.Msg) error {
	var order models.Orders
	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		return errors.Wrap(err, "subscriber: could not unmarshal the order")
	}
	err = s.service.NewOrder(order)
	if err != nil {
		return err
	}
	return nil
}

type Order interface {
	NewOrder(order models.Orders) error
	OrderInfo(uuid uuid.UUID) (models.Orders, error)
}
