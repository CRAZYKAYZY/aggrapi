package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Vendor struct {
	ID             uuid.UUID      `json:"id"`
	UserID         uuid.NullUUID  `json:"user_id"`
	Biography      sql.NullString `json:"biography"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	Active         sql.NullBool   `json:"active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
