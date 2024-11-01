-- name: CreateFeedback :one
INSERT INTO feedback (id, appointment_id, rating, comment)
VALUES ($1, $2, $3, $4)
RETURNING id, appointment_id, rating, comment;

-- name: GetFeedback :one
SELECT id, appointment_id, rating, comment
FROM feedback
WHERE id = $1
LIMIT 1;

-- name: UpdateFeedback :one
UPDATE feedback
SET appointment_id = $2,
    rating = $3,
    comment = $4
WHERE id = $1
RETURNING id, appointment_id, rating, comment;

-- name: DeleteFeedback :one
DELETE FROM feedback
WHERE id = $1
RETURNING id;
