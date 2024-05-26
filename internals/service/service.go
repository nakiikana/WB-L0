package service

import (
	"tools/internals/models"
	"tools/internals/repository"

	"github.com/google/uuid"
)

type Service struct {
	Order
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Order: NewOrderService(repository)}
}

type Order interface {
	NewOrder(order models.Orders) error
	OrderInfo(uuid uuid.UUID) (models.Orders, error)
}
