package app

import (
	"fmt"

	"github.com/ChileKasoka/mis/cmd/api"
	"github.com/ChileKasoka/mis/config"
	"github.com/ChileKasoka/mis/db"
	"github.com/ChileKasoka/mis/internal/repositories"
	"github.com/ChileKasoka/mis/internal/services"
)

type App struct {
	Server *config.Server
}

func Initialize() (*App, error) {

	db, err := db.ConnectDb()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	// Initialize repositories
	appointmentRepo := repositories.NewAppointmentRepository(db)
	//userRepo := repositories.NewUserRepository()

	// Initialize services
	appointmentService := services.NewAppointmentService(appointmentRepo)

	// Initialize handlers
	appointmentHandler := api.NewAppointmentHandler(appointmentService)

	// Initialize server with handlers
	srv := config.NewServer(appointmentHandler)

	return &App{
		Server: srv,
	}, nil
}
