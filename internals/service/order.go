package service

import (
	cache "tools/internals/cache/middleware"
	"tools/internals/models"
	"tools/internals/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	rp *repository.Repository
	ch *cache.Cache
}

func NewOrderService(rp *repository.Repository, ch *cache.Cache) *OrderService {
	return &OrderService{rp: rp, ch: ch}
}

func (s *OrderService) NewOrder(order models.Orders) error {
	err := s.ch.NewOrder(order)
	if err != nil {
		logrus.Println(err)
	}
	return s.rp.NewOrder(order)
}
func (s *OrderService) OrderInfo(uuid uuid.UUID) (*models.Orders, error) {
	return s.ch.OrderInfo(uuid)
}

func (s *OrderService) RecoverCache() error {
	orders, err := s.rp.GetAllOrders()
	if err != nil {
		return err
	}
	if err = s.ch.Recover(orders); err != nil {
		return err
	}
	logrus.Println("Cache recovered")
	return nil
}
