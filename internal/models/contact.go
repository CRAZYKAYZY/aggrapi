package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Contact struct {
	ID      uuid.UUID      `json:"id"`
	UserID  uuid.UUID      `json:"user_id"`
	Phone   sql.NullString `json:"phone"`
	Address sql.NullString `json:"address"`
}

func NewContact(userID, phone, address string) *Contact {
	return &Contact{
		ID:      uuid.New(),
		UserID:  uuid.MustParse(userID),
		Phone:   sql.NullString{String: phone, Valid: phone != ""},
		Address: sql.NullString{String: address, Valid: address != ""},
	}
}
