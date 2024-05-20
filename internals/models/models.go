package models

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

type Delivery struct {
	ID      big.Int
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	ID           big.Int
	Transaction  string  `json:"transaction"`
	RequestID    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       big.Int `json:"amount"`
	PaymentDT    big.Int `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost big.Int `json:"delivery_cost"`
	GoodsTotal   big.Int `json:"goods_total"`
	CustomFee    big.Int `json:"custom_fee"`
}

type Orders struct {
	ID                big.Int
	OrderID           uuid.UUID `json:"order_id"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	DeliveryID        big.Int   `json:"delivery_id"` //?
	PaymentID         big.Int   `json:"payment_id"`  //?
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Locale            string    `json:"payment_dt"`
	IntersanSignature string    `json:"intersan_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              big.Int   `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShared         string    `json:"oof_shared"`
}

type Items struct {
	ID          big.Int
	ChrtID      big.Int `json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       big.Int `json:"price"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        int     `json:"sale"`
	Size        string  `json:"size"`
	TotalPrice  big.Int `json:"total_price"`
	NmID        big.Int `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int     `json:"status"`
}
