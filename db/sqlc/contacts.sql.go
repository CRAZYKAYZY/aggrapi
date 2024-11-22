// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: contacts.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createContact = `-- name: CreateContact :one
INSERT INTO contacts (id, user_id, phone, address)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, phone, address
`

type CreateContactParams struct {
	ID      uuid.UUID      `json:"id"`
	UserID  uuid.UUID      `json:"user_id"`
	Phone   sql.NullString `json:"phone"`
	Address sql.NullString `json:"address"`
}

type CreateContactRow struct {
	ID      uuid.UUID      `json:"id"`
	UserID  uuid.UUID      `json:"user_id"`
	Phone   sql.NullString `json:"phone"`
	Address sql.NullString `json:"address"`
}

func (q *Queries) CreateContact(ctx context.Context, arg CreateContactParams) (CreateContactRow, error) {
	row := q.db.QueryRowContext(ctx, createContact,
		arg.ID,
		arg.UserID,
		arg.Phone,
		arg.Address,
	)
	var i CreateContactRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Phone,
		&i.Address,
	)
	return i, err
}

const deleteContact = `-- name: DeleteContact :one
DELETE FROM contacts
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteContact(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, deleteContact, id)
	err := row.Scan(&id)
	return id, err
}

const getContact = `-- name: GetContact :one
SELECT id, user_id, phone, address
FROM contacts
WHERE id = $1
LIMIT 1
`

type GetContactRow struct {
	ID      uuid.UUID      `json:"id"`
	UserID  uuid.UUID      `json:"user_id"`
	Phone   sql.NullString `json:"phone"`
	Address sql.NullString `json:"address"`
}

func (q *Queries) GetContact(ctx context.Context, id uuid.UUID) (GetContactRow, error) {
	row := q.db.QueryRowContext(ctx, getContact, id)
	var i GetContactRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Phone,
		&i.Address,
	)
	return i, err
}

const updateContact = `-- name: UpdateContact :one
UPDATE contacts
SET phone = $2,
    address = $3
WHERE id = $1
RETURNING id, user_id, phone, address
`

type UpdateContactParams struct {
	ID      uuid.UUID      `json:"id"`
	Phone   sql.NullString `json:"phone"`
	Address sql.NullString `json:"address"`
}

type UpdateContactRow struct {
	ID      uuid.UUID      `json:"id"`
	UserID  uuid.UUID      `json:"user_id"`
	Phone   sql.NullString `json:"phone"`
	Address sql.NullString `json:"address"`
}

func (q *Queries) UpdateContact(ctx context.Context, arg UpdateContactParams) (UpdateContactRow, error) {
	row := q.db.QueryRowContext(ctx, updateContact, arg.ID, arg.Phone, arg.Address)
	var i UpdateContactRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Phone,
		&i.Address,
	)
	return i, err
}
