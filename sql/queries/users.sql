-- name: CreateUser :one
INSERT INTO users(fullname, email, hashedPassword, created_at, updated_at)
VALUES(
    $1,
    $2,
    $3,
    NOW() AT TIME ZONE 'UTC',
    NOW() AT TIME ZONE 'UTC'
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 
LIMIT 1;

-- name: UpdatePassword :one
UPDATE users
SET hashedPassword = $1
WHERE email = $2
RETURNING *;

