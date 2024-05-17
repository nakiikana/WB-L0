CREATE TABLE IF NOT EXISTS orders (
    id bigserial primary key NOT NULL,
    order_uid text NOT NULL, 
    track_number text NOT NULL, 
    entry text NOT NULL, 
    delivery_id bigint foreign key REFERENCES delivery(id), 
    payment_id bigint foreign key REFERENCES payment(id), 
    locale text NOT NULL, 
    intersan_signature text NOT NULL, 
    customer_id text NOT NULL, 
    delivery_service text NOT NULL, 
    shardkey text NOT NULL,
    sm_id bigint NOT NULL, 
    date_created timestamp DEFAULT NOW(),
    oof_shard text NOT NULL
);

CREATE TABLE IF NOT EXISTS delivery (
    id bigserial primary key NOT NULL, 
    name text NOT NULL, 
    phone text NOT NULL, 
    zip text NOT NULL,
    city text NOT NULL, 
    address text NOT NULL,
    region text NOT NULL, 
    email text NOT NULL 
);

CREATE TABLE IF NOT EXISTS payment (
    id bigserial primary key NOT NULL, 
    transaction text NOT NULL, 
    request_id text NOT NULL, 
    currency text NOT NULL, 
    provider text NOT NULL, 
    amount bigint NOT NULL, 
    payment_dt bigint NOT NULL, 
    bank text NOT NULL, 
    delivery_cost bigint NOT NULL, 
    goods_total bigint NOT NULL, 
    custom_fee bigint NOT NULL
); 

CREATE TABLE IF NOT EXISTS items (
    id bigserial primary key NOT NULL, 
    chrt_id bigint NOT NULL, 
    track_number text foreign key REFERENCES order(track_number) NOT NULL, 
    price bigint NOT NULL,
    rid text NOT NULL,
    name text NOT NULL, 
    sale int NOT NULL, 
    total_price bigint NOT NULL, 
    nm_id bigint NOT NULL, 
    brand text NOT NULL, 
    status int NOT NULL
);
