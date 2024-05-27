package repository

const (
	insertDelivery = `
		INSERT INTO delivery (name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	insertPayment = `
		INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`
	insertOrder = `
		INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, locale, intersan_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`
	insertItems = `
		INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`
	getOrder = `
		SELECT * FROM orders
		WHERE order_uid = $1
	`
	getDelivery = `
		SELECT * FROM delivery
		WHERE id = $1
	`
	getPayment = `
		SELECT * FROM payment
		WHERE id = $1
	`
	getItems = `
		SELECT * FROM items
		WHERE track_number = $1
	`
	getAllOrders = `
		SELECT order_uid FROM orders
	`
)
