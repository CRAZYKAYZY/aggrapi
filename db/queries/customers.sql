-- name: CreateCustomer :one
INSERT INTO customers (id, user_id)
VALUES ($1, $2)
RETURNING id, user_id;

-- name: GetCustomer :one
SELECT id, user_id
FROM customers
WHERE id = $1
LIMIT 1;

-- name: UpdateCustomer :one
UPDATE customers
SET user_id = $2
WHERE id = $1
RETURNING id, user_id;

-- name: DeleteCustomer :one
DELETE FROM customers
WHERE id = $1
RETURNING id;
