package services

import (
	models "github.com/ChileKasoka/mis/db/sqlc"
	"github.com/ChileKasoka/mis/internal/repositories"
)

// AppointmentService defines the business logic layer for managing appointments
type AppointmentService interface {
	GetAllAppointments() ([]models.Appointment, error)
	GetAppointmentByID(id string) (*models.Appointment, error)
	CreateAppointment(appointment models.Appointment) (models.Appointment, error)
}

type appointmentServiceImpl struct {
	repository repositories.AppointmentRepository
}

// NewAppointmentService creates a new instance of AppointmentService
func NewAppointmentService(repository repositories.AppointmentRepository) AppointmentService {
	return &appointmentServiceImpl{repository: repository}
}

func (s *appointmentServiceImpl) CreateAppointment(appointment models.Appointment) (models.Appointment, error) {
	// You can add any business logic here, like checking if the appointment is valid
	createdAppointment, err := s.repository.CreateAppointment(appointment)
	if err != nil {
		return models.Appointment{}, err
	}
	return createdAppointment, nil
}

// GetAllAppointments retrieves all appointments from the repository
func (s *appointmentServiceImpl) GetAllAppointments() ([]models.Appointment, error) {
	appointments, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

// GetAppointmentByID retrieves a specific appointment by its ID
func (s *appointmentServiceImpl) GetAppointmentByID(id string) (*models.Appointment, error) {
	appointment, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}
