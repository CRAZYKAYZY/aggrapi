-- name: CreateVendor :one
INSERT INTO vendors (id, user_id, biography, profile_picture, active)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, biography, profile_picture, active;

-- name: GetVendor :one
SELECT id, user_id, biography, profile_picture, active
FROM vendors
WHERE id = $1
LIMIT 1;

-- name: UpdateVendor :one
UPDATE vendors
SET biography = $2,
    profile_picture = $3,
    active = $4
WHERE id = $1
RETURNING id, user_id, biography, profile_picture, active;

-- name: DeleteVendor :one
DELETE FROM vendors
WHERE id = $1
RETURNING id;
