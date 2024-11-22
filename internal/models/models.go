package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Feedback struct {
	ID            uuid.UUID      `json:"id"`
	AppointmentID uuid.NullUUID  `json:"appointment_id"`
	Rating        sql.NullInt32  `json:"rating"`
	Comment       sql.NullString `json:"comment"`
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
}

type Service struct {
	ID          uuid.UUID      `json:"id"`
	VendorID    uuid.NullUUID  `json:"vendor_id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	Price       sql.NullString `json:"price"`
	Duration    sql.NullInt64  `json:"duration"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type TimeSlot struct {
	ID         uuid.UUID     `json:"id"`
	VendorID   uuid.NullUUID `json:"vendor_id"`
	StartTime  sql.NullTime  `json:"start_time"`
	EndTime    sql.NullTime  `json:"end_time"`
	IsBooked   sql.NullBool  `json:"is_booked"`
	BufferTime sql.NullInt64 `json:"buffer_time"`
}

type VendorAvailability struct {
	ID        uuid.UUID     `json:"id"`
	VendorID  uuid.NullUUID `json:"vendor_id"`
	DayOfWeek interface{}   `json:"day_of_week"`
	Date      sql.NullTime  `json:"date"`
}
