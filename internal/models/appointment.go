package models

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentStatus string

const (
	StatusPending   AppointmentStatus = "pending"
	StatusConfirmed AppointmentStatus = "confirmed"
	StatusCompleted AppointmentStatus = "completed"
	StatusCanceled  AppointmentStatus = "canceled"
)

type Appointment struct {
	ID         uuid.UUID         `json:"id"`
	CustomerID uuid.UUID         `json:"customer_id"`
	VendorID   uuid.UUID         `json:"vendor_id"`
	Date       time.Time         `json:"date"`
	TimeSlotID uuid.UUID         `json:"time_slot_id"`
	Status     AppointmentStatus `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// NewAppointment initializes a new Appointment with the provided values.
func NewAppointment(customerID, vendorID, timeSlotID uuid.UUID, date time.Time, status AppointmentStatus) *Appointment {
	// Default to "pending" status if none provided
	if status == "" {
		status = StatusPending
	}

	return &Appointment{
		ID:         uuid.New(),
		CustomerID: customerID,
		VendorID:   vendorID,
		Date:       date,
		TimeSlotID: timeSlotID,
		Status:     status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
