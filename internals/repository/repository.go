package repository

import (
	"tools/internals/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Order: NewOrderRepository(db)}
}

type Order interface {
	NewOrder(order models.Orders) error
}
