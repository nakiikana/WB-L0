package repository

import (
	"tools/internals/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) NewOrder(order models.Orders) error {
	var d_id, p_id, o_id int
	err := o.db.QueryRow(insertDelivery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).Scan(&d_id)
	if err != nil {
		logrus.Errorf("sad: %v", err)
		return errors.Wrap(err, "repository: could not insert delivery")
	}
	err = o.db.QueryRow(insertPayment, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee).Scan(&p_id)
	if err != nil {
		return errors.Wrap(err, "repository: could not insert payment")
	}
	err = o.db.QueryRow(insertOrder, order.OrderID, order.TrackNumber, order.Entry, order.Locale, order.IntersanSignature, order.CustomerID, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShared).Scan(&o_id)
	if err != nil {
		return errors.Wrap(err, "repository: could not insert order")
	}
	return nil
}
