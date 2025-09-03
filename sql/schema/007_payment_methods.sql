-- +goose Up
CREATE TABLE payment_methods(
    id BIGSERIAL PRIMARY KEY,
    card_number TEXT NOT NULL,
    expiry_date TEXT NOT NULL,
    cvv INTEGER NOT NULL,
    name_on_card TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE payment_methods;