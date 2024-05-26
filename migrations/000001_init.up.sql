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
    amount int NOT NULL, 
    payment_dt bigint NOT NULL, 
    bank text NOT NULL, 
    delivery_cost int NOT NULL, 
    goods_total int NOT NULL, 
    custom_fee int NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id bigserial primary key NOT NULL,
    order_uid uuid NOT NULL UNIQUE, 
    track_number text NOT NULL UNIQUE, 
    entry text NOT NULL, 
    delivery_id bigint NOT NULL, 
    payment_id bigint NOT NULL, 
    foreign key (delivery_id) REFERENCES delivery (id) on delete cascade, 
    foreign key (payment_id) REFERENCES payment(id) on delete cascade, 
    locale text NOT NULL, 
    intersan_signature text NOT NULL, 
    customer_id text NOT NULL, 
    delivery_service text NOT NULL, 
    shardkey text NOT NULL,
    sm_id bigint NOT NULL, 
    date_created timestamp DEFAULT NOW(),
    oof_shard text NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
    id bigserial primary key NOT NULL, 
    chrt_id bigint NOT NULL,
    track_number text NOT NULL,
    foreign key (track_number) REFERENCES orders (track_number) on delete cascade, 
    price int NOT NULL,
    rid text NOT NULL,
    name text NOT NULL, 
    sale int NOT NULL, 
    size text NOT NULL, 
    total_price int NOT NULL, 
    nm_id bigint NOT NULL, 
    brand text NOT NULL, 
    status int NOT NULL
);
