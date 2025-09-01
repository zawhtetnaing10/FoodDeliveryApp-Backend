-- +goose Up
CREATE TABLE restaurants(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    image_url TEXT NOT NULL,
    average_rating DECIMAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


-- +goose Down
DROP TABLE restaurants;