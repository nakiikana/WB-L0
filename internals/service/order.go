package service

import (
	"tools/internals/models"
	"tools/internals/repository"
)

type OrderService struct {
	rp *repository.Repository
}

func NewOrderService(rp *repository.Repository) *OrderService {
	return &OrderService{rp: rp}
}

func (s *OrderService) NewOrder(order models.Orders) error {
	return s.rp.NewOrder(order)
}
