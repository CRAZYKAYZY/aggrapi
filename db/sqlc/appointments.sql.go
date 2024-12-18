// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: appointments.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const checkConfirmedAppointment = `-- name: CheckConfirmedAppointment :one
SELECT EXISTS (
    SELECT 1 
    FROM appointments 
    WHERE vendor_id = $1 
    AND time_slot_id = $2 
    AND date = $3 
    AND status = 'confirmed'
) AS exists
`

type CheckConfirmedAppointmentParams struct {
	VendorID   uuid.UUID `json:"vendor_id"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Date       time.Time `json:"date"`
}

func (q *Queries) CheckConfirmedAppointment(ctx context.Context, arg CheckConfirmedAppointmentParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkConfirmedAppointment, arg.VendorID, arg.TimeSlotID, arg.Date)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createAppointment = `-- name: CreateAppointment :one
INSERT INTO appointments (id, customer_id, vendor_id, date, time_slot_id, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, customer_id, vendor_id, date, time_slot_id, status
`

type CreateAppointmentParams struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	VendorID   uuid.UUID `json:"vendor_id"`
	Date       time.Time `json:"date"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Status     string    `json:"status"`
}

type CreateAppointmentRow struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	VendorID   uuid.UUID `json:"vendor_id"`
	Date       time.Time `json:"date"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Status     string    `json:"status"`
}

func (q *Queries) CreateAppointment(ctx context.Context, arg CreateAppointmentParams) (CreateAppointmentRow, error) {
	row := q.db.QueryRowContext(ctx, createAppointment,
		arg.ID,
		arg.CustomerID,
		arg.VendorID,
		arg.Date,
		arg.TimeSlotID,
		arg.Status,
	)
	var i CreateAppointmentRow
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.VendorID,
		&i.Date,
		&i.TimeSlotID,
		&i.Status,
	)
	return i, err
}

const deleteAppointment = `-- name: DeleteAppointment :one
DELETE FROM appointments
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteAppointment(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, deleteAppointment, id)
	err := row.Scan(&id)
	return id, err
}

const getAllAppointments = `-- name: GetAllAppointments :many
SELECT id, customer_id, vendor_id, date, time_slot_id, status, created_at, updated_at from appointments
`

func (q *Queries) GetAllAppointments(ctx context.Context) ([]Appointment, error) {
	rows, err := q.db.QueryContext(ctx, getAllAppointments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Appointment{}
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.VendorID,
			&i.Date,
			&i.TimeSlotID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAppointmentById = `-- name: GetAppointmentById :one
SELECT id, customer_id, vendor_id, date, time_slot_id, status
FROM appointments
WHERE id = $1
LIMIT 1
`

type GetAppointmentByIdRow struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	VendorID   uuid.UUID `json:"vendor_id"`
	Date       time.Time `json:"date"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Status     string    `json:"status"`
}

func (q *Queries) GetAppointmentById(ctx context.Context, id uuid.UUID) (GetAppointmentByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getAppointmentById, id)
	var i GetAppointmentByIdRow
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.VendorID,
		&i.Date,
		&i.TimeSlotID,
		&i.Status,
	)
	return i, err
}

const updateAppointment = `-- name: UpdateAppointment :one
UPDATE appointments
SET customer_id = $2,
    vendor_id = $3,
    date = $4,
    time_slot_id = $5,
    status = $6
WHERE id = $1
RETURNING id, customer_id, vendor_id, date, time_slot_id, status
`

type UpdateAppointmentParams struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	VendorID   uuid.UUID `json:"vendor_id"`
	Date       time.Time `json:"date"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Status     string    `json:"status"`
}

type UpdateAppointmentRow struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	VendorID   uuid.UUID `json:"vendor_id"`
	Date       time.Time `json:"date"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Status     string    `json:"status"`
}

func (q *Queries) UpdateAppointment(ctx context.Context, arg UpdateAppointmentParams) (UpdateAppointmentRow, error) {
	row := q.db.QueryRowContext(ctx, updateAppointment,
		arg.ID,
		arg.CustomerID,
		arg.VendorID,
		arg.Date,
		arg.TimeSlotID,
		arg.Status,
	)
	var i UpdateAppointmentRow
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.VendorID,
		&i.Date,
		&i.TimeSlotID,
		&i.Status,
	)
	return i, err
}
