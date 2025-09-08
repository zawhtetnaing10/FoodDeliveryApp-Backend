-- name: CreateOrder :one
INSERT INTO orders(user_id, delivery_address_id, payment_method_id, total_cost, order_number, created_at, updated_at)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW() AT TIME ZONE 'UTC',
    NOW() AT TIME ZONE 'UTC'
)
RETURNING *;