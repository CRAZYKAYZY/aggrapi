-- name: CreatePayment :one
INSERT INTO payments (id, appointment_id, customer_id, vendor_id, amount, payment_method, status, payment_date, transaction_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, appointment_id, customer_id, vendor_id, amount, payment_method, status, payment_date, transaction_id;

-- name: GetPayment :one
SELECT id, appointment_id, customer_id, vendor_id, amount, payment_method, status, payment_date, transaction_id
FROM payments
WHERE id = $1
LIMIT 1;

-- name: UpdatePayment :one
UPDATE payments
SET appointment_id = $2,
    customer_id = $3,
    vendor_id = $4,
    amount = $5,
    payment_method = $6,
    status = $7,
    payment_date = $8,
    transaction_id = $9
WHERE id = $1
RETURNING id, appointment_id, customer_id, vendor_id, amount, payment_method, status, payment_date, transaction_id;

-- name: DeletePayment :one
DELETE FROM payments
WHERE id = $1
RETURNING id;
