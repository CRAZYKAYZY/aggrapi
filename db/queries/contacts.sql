-- name: CreateContact :one
INSERT INTO contacts (id, user_id, phone, address)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, phone, address;

-- name: GetContact :one
SELECT id, user_id, phone, address
FROM contacts
WHERE id = $1
LIMIT 1;

-- name: UpdateContact :one
UPDATE contacts
SET phone = $2,
    address = $3
WHERE id = $1
RETURNING id, user_id, phone, address;

-- name: DeleteContact :one
DELETE FROM contacts
WHERE id = $1
RETURNING id;
