-- name: CreatePaymentMethod :one
INSERT INTO payment_methods (card_number, expiry_date, cvv, name_on_card, user_id, created_at, updated_at)
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

-- name: GetPaymentMethodsByUser :many
SELECT * FROM payment_methods
WHERE user_id = $1;