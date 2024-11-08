package repositories

import (
	"context"
	"database/sql"

	sqlc "github.com/ChileKasoka/mis/db/sqlc"
	models "github.com/ChileKasoka/mis/internal/models"
	"github.com/google/uuid"
)

type AppointmentRepository interface {
	FindAll() ([]models.Appointment, error)
	FindById(id string) (*models.Appointment, error)
	CreateAppointment(appointment models.Appointment) (models.Appointment, error)
}

type appointmentRepositoryImpl struct {
	Queries *sqlc.Queries
}

func NewAppointmentRepository(db *sql.DB) AppointmentRepository {
	return &appointmentRepositoryImpl{Queries: sqlc.New(db)}
}

func (r *appointmentRepositoryImpl) CreateAppointment(appointment models.Appointment) (models.Appointment, error) {
	arg := sqlc.CreateAppointmentParams{
		ID:         uuid.New(),
		CustomerID: appointment.CustomerID,
		VendorID:   appointment.VendorID,
		Date:       appointment.Date,
		TimeSlotID: appointment.TimeSlotID,
		Status:     appointment.Status,
	}

	createdAppointment, err := r.Queries.CreateAppointment(context.TODO(), arg)
	if err != nil {
		return models.Appointment{}, err
	}

	result := models.Appointment{
		ID:         createdAppointment.ID,
		CustomerID: createdAppointment.CustomerID,
		VendorID:   createdAppointment.VendorID,
		Date:       createdAppointment.Date,
		TimeSlotID: createdAppointment.TimeSlotID,
		Status:     createdAppointment.Status,
	}

	return result, nil
}

func (r *appointmentRepositoryImpl) FindAll() ([]models.Appointment, error) {
	rows, err := r.Queries.GetAllAppointments(context.TODO())
	if err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	for _, row := range rows {
		appointments = append(appointments, models.Appointment{
			ID:         row.ID,
			CustomerID: row.CustomerID,
			VendorID:   row.VendorID,
			Date:       row.Date,
			TimeSlotID: row.TimeSlotID,
			Status:     row.Status,
		})
	}

	return appointments, nil
}

func (r *appointmentRepositoryImpl) FindById(id string) (*models.Appointment, error) {
	// Convert the string ID to a UUID (assuming ID is stored as UUID)
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Fetch the appointment using sqlc's generated method
	appointment, err := r.Queries.GetAppointmentById(context.TODO(), uuidID)
	if err != nil {
		return nil, err
	}

	// Map sqlc's Appointment result to your models.Appointment (if needed)
	result := &models.Appointment{
		ID:         appointment.ID,
		CustomerID: appointment.CustomerID,
		VendorID:   appointment.VendorID,
		Date:       appointment.Date,
		TimeSlotID: appointment.TimeSlotID,
		Status:     appointment.Status,
	}

	return result, nil
}
