-- name: CreateUser :one
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1 OFFSET $2;

-- name: GetUserByEmail :one
SELECT id, name, email, password
FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
set name = $2,
email = $3,
password = $4
WHERE id = $1
RETURNING *;
