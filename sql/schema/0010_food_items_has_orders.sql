-- +goose Up
CREATE TABLE food_items_has_orders(
    id BIGSERIAL PRIMARY KEY,
    food_item_id BIGINT NOT NULL REFERENCES food_items(id) ON DELETE SET NULL,
    order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE SET NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(food_item_id, order_id)    
);

-- +goose Down
DROP TABLE food_items_has_orders;