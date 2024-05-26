package repository

import (
	"tools/internals/models"

	"github.com/google/uuid"
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
	i_ids := make([]int, len(order.Items))
	err := o.db.QueryRow(insertDelivery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).Scan(&d_id)
	if err != nil {
		return errors.Wrap(err, "repository: could not insert delivery")
	}
	err = o.db.QueryRow(insertPayment, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee).Scan(&p_id)
	if err != nil {
		return errors.Wrap(err, "repository: could not insert payment")
	}
	err = o.db.QueryRow(insertOrder, order.OrderID, order.TrackNumber, order.Entry, d_id, p_id, order.Locale, order.IntersanSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShared).Scan(&o_id)
	if err != nil {
		return errors.Wrap(err, "repository: could not insert order")
	}

	for i := range order.Items {
		err = o.db.QueryRow(insertItems, order.Items[i].ChrtID, order.Items[i].TrackNumber, order.Items[i].Price, order.Items[i].Rid, order.Items[i].Name, order.Items[i].Sale, order.Items[i].Size, order.Items[i].TotalPrice, order.Items[i].NmID, order.Items[i].Brand, order.Items[i].Status).Scan(&(i_ids[i]))
		if err != nil {
			logrus.Warningf("repository: ceased inserting on %d out of %d", i, len(order.Items))
			return errors.Wrap(err, "repository: could not insert item")
		}
	}
	return nil
}

func (o *OrderRepository) OrderInfo(uuid uuid.UUID) (models.Orders, error) {
	var order models.Orders
	err := o.db.QueryRow(getOrder, uuid).Scan(&order.ID, &order.OrderID, &order.TrackNumber, &order.Entry, &order.Delivery.ID, &order.Payment.ID, &order.Locale, &order.IntersanSignature, &order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShared)
	if err != nil {
		return order, errors.Wrap(err, "repository: could not retract order")
	}
	err = o.db.QueryRow(getPayment, order.Payment.ID).Scan(&order.Payment.ID, &order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDT, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee)
	if err != nil {
		return order, errors.Wrap(err, "repository: could not retract payment for order")
	}
	err = o.db.QueryRow(getDelivery, order.Delivery.ID).Scan(&order.Delivery.ID, &order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email)
	if err != nil {
		return order, errors.Wrap(err, "repository: could not retract delivery for order")
	}
	rows, err := o.db.Query(getItems, order.TrackNumber)
	if err != nil {
		return order, errors.Wrap(err, "repository: could not retract items for order")
	}
	defer func() {
		_ = rows.Close()
	}()
	items := make([]models.Items, 0)
	for rows.Next() {
		var item models.Items
		if err = rows.Scan(&item.ID, &item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status); err != nil {
			return order, errors.Wrap(err, "repository: could not scan a row")
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		logrus.Errorf("repository: couldn't execute the OrderInfo query: %v", err)
		return order, err
	}
	order.Items = items
	return order, nil
}
