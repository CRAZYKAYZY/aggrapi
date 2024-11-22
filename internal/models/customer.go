package models

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID     uuid.UUID     `json:"id"`
	UserID uuid.NullUUID `json:"user_id"`
}
