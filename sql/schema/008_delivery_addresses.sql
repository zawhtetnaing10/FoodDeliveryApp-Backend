-- +goose Up
CREATE TABLE delivery_addresses(
    id BIGSERIAL PRIMARY KEY,
    street_address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE delivery_addresses;