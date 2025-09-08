-- +goose Up
CREATE TABLE orders(
    id BIGSERIAL PRIMARY KEY,
    total_cost NUMERIC NOT NULL,
    order_number TEXT NOT NULL,
    user_id BIGINT REFERENCES users(id) ON DELETE NO ACTION,
    delivery_address_id BIGINT REFERENCES delivery_addresses(id) ON DELETE SET NULL,
    payment_method_id BIGINT REFERENCES payment_methods(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL  
);


-- +goose Down
DROP TABLE orders;