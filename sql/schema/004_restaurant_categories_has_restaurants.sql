-- +goose Up
CREATE TABLE restaurant_categories_has_restaurants(
    id BIGSERIAL PRIMARY KEY,
    restaurant_category_id BIGINT NOT NULL  REFERENCES restaurant_categories(id) ON DELETE CASCADE,
    restaurant_id BIGINT NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE(restaurant_category_id, restaurant_id)
);


-- +goose Down
DROP TABLE restaurant_categories_has_restaurants;