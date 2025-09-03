-- name: CreateDeliveryAddress :one
INSERT INTO delivery_addresses(street_address, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    NOW() AT TIME ZONE 'UTC',
    NOW() AT TIME ZONE 'UTC'
)
RETURNING *;

-- name: GetDeliveryAddressesForUser :many
SELECT * FROM delivery_addresses
WHERE user_id = $1;