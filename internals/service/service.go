package service

import (
	cache "tools/internals/cache/middleware"
	"tools/internals/models"
	"tools/internals/repository"

	"github.com/google/uuid"
)

type Service struct {
	Order
}

func NewService(repository *repository.Repository, cache *cache.Cache) *Service {
	return &Service{Order: NewOrderService(repository, cache)}
}

type Order interface {
	NewOrder(order models.Orders) error
	OrderInfo(uuid uuid.UUID) (*models.Orders, error)
}
