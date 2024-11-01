-- name: CreateService :one
INSERT INTO services (id, vendor_id, name, description, price, duration, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, vendor_id, name, description, price, duration, created_at, updated_at;

-- name: GetService :one
SELECT id, vendor_id, name, description, price, duration, created_at, updated_at
FROM services
WHERE id = $1
LIMIT 1;

-- name: UpdateService :one
UPDATE services
SET vendor_id = $2,
    name = $3,
    description = $4,
    price = $5,
    duration = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, vendor_id, name, description, price, duration, updated_at;

-- name: DeleteService :one
DELETE FROM services
WHERE id = $1
RETURNING id;
