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

-- name: CalculateTotalCost :one
SELECT SUM(f.price * t.quantity)::numeric
FROM food_items AS f
CROSS JOIN LATERAL
    jsonb_to_recordset(
        sqlc.arg(items)::jsonb
    ) AS t(id INT, quantity INT)
WHERE 
f.id = t.id;