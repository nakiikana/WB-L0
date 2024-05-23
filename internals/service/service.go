package service

import (
	"tools/internals/models"
	"tools/internals/repository"
)

type Service struct {
	Order
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Order: NewOrderService(repository)}
}

type Order interface {
	NewOrder(order models.Orders) error
}
