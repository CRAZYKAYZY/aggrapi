-- name: CreateVendorAvailability :one
INSERT INTO vendor_availability (id, vendor_id, day_of_week, date)
VALUES ($1, $2, $3, $4)
RETURNING id, vendor_id, day_of_week, date;

-- name: GetVendorAvailability :one
SELECT id, vendor_id, day_of_week, date
FROM vendor_availability
WHERE id = $1
LIMIT 1;

-- name: UpdateVendorAvailability :one
UPDATE vendor_availability
SET vendor_id = $2,
    day_of_week = $3,
    date = $4
WHERE id = $1
RETURNING id, vendor_id, day_of_week, date;

-- name: DeleteVendorAvailability :one
DELETE FROM vendor_availability
WHERE id = $1
RETURNING id;
