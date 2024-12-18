// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	VendorID   uuid.UUID `json:"vendor_id"`
	Date       time.Time `json:"date"`
	TimeSlotID uuid.UUID `json:"time_slot_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Contact struct {
	ID        uuid.UUID      `json:"id"`
	UserID    uuid.UUID      `json:"user_id"`
	Phone     sql.NullString `json:"phone"`
	Address   sql.NullString `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Customer struct {
	ID     uuid.UUID     `json:"id"`
	UserID uuid.NullUUID `json:"user_id"`
}

type Feedback struct {
	ID            uuid.UUID      `json:"id"`
	AppointmentID uuid.NullUUID  `json:"appointment_id"`
	Rating        sql.NullInt32  `json:"rating"`
	Comment       sql.NullString `json:"comment"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Payment struct {
	ID            uuid.UUID      `json:"id"`
	AppointmentID uuid.NullUUID  `json:"appointment_id"`
	CustomerID    uuid.NullUUID  `json:"customer_id"`
	VendorID      uuid.NullUUID  `json:"vendor_id"`
	Amount        sql.NullString `json:"amount"`
	PaymentMethod interface{}    `json:"payment_method"`
	Status        interface{}    `json:"status"`
	PaymentDate   sql.NullTime   `json:"payment_date"`
	TransactionID sql.NullString `json:"transaction_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Service struct {
	ID          uuid.UUID      `json:"id"`
	VendorID    uuid.UUID      `json:"vendor_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       sql.NullString `json:"price"`
	Duration    sql.NullInt64  `json:"duration"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type TimeSlot struct {
	ID         uuid.UUID     `json:"id"`
	VendorID   uuid.NullUUID `json:"vendor_id"`
	StartTime  sql.NullTime  `json:"start_time"`
	EndTime    sql.NullTime  `json:"end_time"`
	IsBooked   sql.NullBool  `json:"is_booked"`
	BufferTime sql.NullInt64 `json:"buffer_time"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vendor struct {
	ID             uuid.UUID      `json:"id"`
	UserID         uuid.NullUUID  `json:"user_id"`
	Biography      sql.NullString `json:"biography"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	Active         sql.NullBool   `json:"active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type VendorAvailability struct {
	ID        uuid.UUID    `json:"id"`
	VendorID  uuid.UUID    `json:"vendor_id"`
	DayOfWeek string       `json:"day_of_week"`
	Date      sql.NullTime `json:"date"`
}
