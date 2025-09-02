-- +goose Up
CREATE TABLE food_items(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    image_url TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    restaurant_id BIGINT NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    food_category_id BIGINT NOT NULL REFERENCES food_categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


-- +goose Down
DROP TABLE food_items;